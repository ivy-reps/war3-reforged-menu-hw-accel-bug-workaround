package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/ncruces/zenity"
	"github.com/shirou/gopsutil/process"
)

const title = "Warcraft III"
const launcherName = "Retrying to start Warcraft III with HW Acceleration support"

var startTime = time.Now()

func findProcessByExeNameAndCmdLineSubStr(exeName string, cmdLine string) *process.Process {
	processes, e := process.Processes()
	if e != nil {

	}

	for _, prcs := range processes {
		exe, e := prcs.Exe()
		if e != nil {

		}

		if strings.Contains(exe, exeName) {
			processCmdLine, e := prcs.Cmdline()
			if e != nil {

			}
			if cmdLine == "" {
				return prcs
			}
			if strings.Contains(processCmdLine, cmdLine) {
				return prcs
			}
		}
	}
	return nil
}

var lastLineWidth int
var gameIsLoadedSuccessfully = false

func writeLine(strs ...string) {
	line := strings.Join(strs, ",")

	if lastLineWidth > 0 {
		fmt.Printf(strings.Repeat(" ", lastLineWidth))
	}
	lastLineWidth = len(line)
	fmt.Printf("\r" + line)

}

func findProcessPidByExeNameAndCmdLineSubStr(needleName, needleCmd string, prcs *process.Process) int32 {
	childrens, err := prcs.Children()
	if err != nil {
		log.Fatal(err)
	}
	for _, children := range childrens {
		cmdline, err := children.Cmdline()
		if err != nil {
			log.Fatal(err)
		}

		name, err := children.Name()
		if err != nil {
			log.Fatal(err)
		}

		if name == needleName && strings.Contains(cmdline, needleCmd) {
			return children.Pid
		}

		if pid := findProcessPidByExeNameAndCmdLineSubStr(needleName, needleCmd, children); pid > 0 {
			return pid
		}

	}

	return 0
}

var executionAttempts = 1

func killWarcraftIIIAndIncrease(process *process.Process) {

	if err := process.Kill(); err != nil {
		log.Fatal(err)
	}
	executionAttempts++

}

func main() {
	var warcraftProcess *process.Process

	dlg, err := zenity.Progress(
		zenity.Title(launcherName),
		zenity.Pulsate(), zenity.NoCancel(),
		zenity.OKLabel("Stop it!"),
	)
	if err != nil {
		return
	}
	defer dlg.Close()

	dlg.Text("Checking if Warcraft III started already")
	dlg.Value(0)

	dlgDone := dlg.Done()

	go func() {
		if _, ok := <-dlgDone; warcraftProcess != nil && !ok && !gameIsLoadedSuccessfully {
			isRunning, err := warcraftProcess.IsRunning()
			if err != nil {
				log.Fatal(err)
			}

			if isRunning {
				err = warcraftProcess.Kill()
				if err != nil {
					log.Fatal(err)
				}
			}

			os.Exit(0)
		}
	}()

	for {

		time.Sleep(time.Second)
		var blizzardBrowserPID int32

		warcraftProcessIsRunning := false
		if warcraftProcess == nil {
			warcraftProcess = findProcessByExeNameAndCmdLineSubStr("Warcraft III.exe", "")
		}

		if warcraftProcess != nil {
			isRunning, err := warcraftProcess.IsRunning()
			if err != nil {
				log.Fatal(err)
			}
			warcraftProcessIsRunning = isRunning
		}

		if !warcraftProcessIsRunning {
			if err := dlg.Text(fmt.Sprintf("Attempt %v: Starting Warcraft III...", executionAttempts)); err != nil {
				log.Fatal(err)
			}
			cmd := exec.Command(`Warcraft III.exe`, "-launch", "-windowmode", "windowed")
			cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000, HideWindow: true}
			err = cmd.Start()

			if err != nil {
				log.Fatal(err)
			}

			warcraftProcess, err = process.NewProcess(int32(cmd.Process.Pid))
			if err != nil {
				log.Fatal(err)
			}
			if isRunning, _ := warcraftProcess.IsRunning(); isRunning && warcraftProcess.Pid != 0 {

				if err := dlg.Text(fmt.Sprintf("Attempt %v: Warcraft III started hiddenly....", executionAttempts)); err != nil {
					log.Fatal(err)
				}

				if err := dlg.Value(20); err != nil {
					log.Fatal(err)
				}

				time.Sleep(time.Second)

				if err := dlg.Text(fmt.Sprintf("Attempt %v: Awaiting for first BlizzardBrowser process to be loaded...", executionAttempts)); err != nil {
					log.Fatal(err)
				}
			}
			continue
		}

		blizzardBrowserPID = findProcessPidByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "width", warcraftProcess)
		if blizzardBrowserPID == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		err = dlg.Text(fmt.Sprintf("Attempt %v: BlizzardBrowser loaded... Awaiting for Browser renderer subprocess...", executionAttempts))
		if err != nil {
			log.Fatal(err)
		}

		err = dlg.Value(30)
		if err != nil {
			log.Fatal(err)
		}

		blizzardRenderPID := findProcessPidByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=ren", warcraftProcess)
		if blizzardRenderPID == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		err = dlg.Text(fmt.Sprintf("Attempt %v: BlizzardBrowser renderer loaded... Awaiting for Browser GPU subprocess...", executionAttempts))
		if err != nil {
			log.Fatal(err)
		}
		err = dlg.Value(40)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(2 * time.Second)
		blizzard3DPID := findProcessPidByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu", warcraftProcess)
		if blizzard3DPID == 0 {
			err = dlg.Text(fmt.Sprintf("Attempt %v: Browser GPU subprocess was not loaded. Rechecking in 2 seconds...", executionAttempts))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(2 * time.Second)
			blizzard3DPID2ndAttempt := findProcessPidByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu", warcraftProcess)
			if blizzard3DPID2ndAttempt == 0 {
				err = dlg.Text(fmt.Sprintf("Attempt %v: Browser GPU subprocess was not loaded. Let's repeat it all over again Warcraft III", executionAttempts))
				if err != nil {
					log.Fatal(err)
				}
				time.Sleep(100 * time.Millisecond)

				killWarcraftIIIAndIncrease(warcraftProcess)
				continue
			}
		}
		err = dlg.Text(fmt.Sprintf("Attempt %v: Browser GPU subprocess is loaded. Checking that it will not be crashed in 2 seconds", executionAttempts))
		if err != nil {
			log.Fatal(err)
		}
		err = dlg.Value(50)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(2 * time.Second)

		if blizzard3DPID2ndAttempt := findProcessPidByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu", warcraftProcess); blizzard3DPID2ndAttempt != 0 {
			break
		}

		if err := dlg.Text(fmt.Sprintf("Attempt %v: Browser GPU subprocess was not loaded. Let's repeat it all over again Warcraft III", executionAttempts)); err != nil {
			log.Fatal(err)
		}
		killWarcraftIIIAndIncrease(warcraftProcess)

		continue
	}

	if err := dlg.Text(fmt.Sprintf("Warcraft III loaded in %v attempts. Please, have fun!", executionAttempts)); err != nil {
		log.Fatal(err)
	}

	if err := dlg.Complete(); err != nil {
		log.Fatal(err)
	}

	gameIsLoadedSuccessfully = true
	var endTime = time.Now()

	timeSpent := endTime.Sub(startTime)
	minutesSpent := math.Round(timeSpent.Seconds() / 60)
	go func() {
		if err := zenity.Notify(
			fmt.Sprintf("Game is ready to play in %v attempts and %v minutes! Just give me Go play it! to try to go to windowed fullscreen and to close the app then", executionAttempts, minutesSpent),
			zenity.Title(launcherName),
		); err != nil {
			if err.Error() != "Invalid window handle." {
				log.Println(fmt.Errorf("Error on notification showing: %v", err))
			} else {
				log.Println(fmt.Errorf("Error on notification showing: %v.", err))
				log.Println("Probably that one or any other one notification had shown before, but in Windows 7 there are bug, when after to be shown, notification icon in tray not dissappears, so the place is busy, that or related thing is the cause that error.")
			}
		}
	}()

	writeLine(fmt.Sprintf("%v attempts of game execution was made, to wait until gpu renderer will not crash.\n", executionAttempts))
	time.Sleep(3 * time.Second)

	dlg.Close()

	dlgWhenFinished, err := zenity.Progress(
		zenity.Title(launcherName),
		zenity.Pulsate(), zenity.NoCancel(),
		zenity.MaxValue(100),
		zenity.OKLabel("Go play it!"),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer dlgWhenFinished.Close()

	dlgWhenFinished.Text(fmt.Sprintf("Game is ready to play in %v attempts and %v minutes! Just give me Go play it! to try to go to windowed fullscreen and to close the app then", executionAttempts, minutesSpent))

	dlgWhenFinished.Complete()

	var wg sync.WaitGroup
	wg.Add(1)
	done := dlgWhenFinished.Done()
	go func() {
		if _, ok := <-done; !ok {
			wg.Done()
		}
	}()
	wg.Wait()

	if err := dlgWhenFinished.Text("Showing and maximizing Warcraft III..."); err != nil && err.Error() != "dialog canceled" {
		log.Println("Error on Showing and maximizing Warcraft III:", err)
	}

	running, err := warcraftProcess.IsRunning()
	if err != nil {
		log.Fatal(err)
	}

	if !running {
		if err := dlgWhenFinished.Text("Warcraft III has disappeared. Don't know why..."); err != nil && err.Error() != "dialog canceled" {
			log.Println("Error on changing text of dialog: Warcraft III has disappeared. Don't know why:", err)
		}
		log.Fatal("Error on finding Warcraft III:" + "Warcraft III has disappeared. Don't know why...")
	}

	robotgo.MinWindow(warcraftProcess.Pid)
	robotgo.MaxWindow(warcraftProcess.Pid)

	if err := dlgWhenFinished.Text("Switching to Warcraft III and sending alt + enter the window to turn fullcreen windowed mode on..."); err != nil && err.Error() != "dialog canceled" {
		log.Println("Error on changing text of dialog: Switching to Warcraft III and ...:", err)
	}

	robotgo.KeyDown("alt")
	robotgo.KeyDown("enter")
	robotgo.KeyUp("alt")
	robotgo.KeyUp("enter")

	os.Exit(0)
	os.Exit(0)
}
