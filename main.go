package main

import (
	"fmt"
	"log"
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

func findProcessByExeNameAndCmdLineSubStr(exeName string, cmdLine string) int32 {
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
				return prcs.Pid
			}
			if strings.Contains(processCmdLine, cmdLine) {
				return prcs.Pid
			}
		}
	}
	return 0
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

func main() {
	executionAttempts := 0
	var warcraftProcessPID int32

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
		if _, ok := <-dlgDone; !ok {
			if !gameIsLoadedSuccessfully {
				if warcraftProcessPID := findProcessByExeNameAndCmdLineSubStr("Warcraft III.exe", ""); warcraftProcessPID > 0 {
					prcss, err := process.NewProcess(warcraftProcessPID)
					if err != nil {
						log.Fatal(err)
					}
					err = prcss.Kill()
					if err != nil {
						log.Fatal(err)
					}
				}

				os.Exit(0)
			}
		}
	}()

	for {

		time.Sleep(time.Second)
		var blizzardBrowserPID int32

		warcraftProcessPID = findProcessByExeNameAndCmdLineSubStr("Warcraft III.exe", "")
		if warcraftProcessPID == 0 {
			executionAttempts++

			if err := dlg.Text(fmt.Sprintf("Attempt %v: Starting Warcraft III...", executionAttempts)); err != nil {
				log.Fatal(err)
			}
			cmd := exec.Command(`Warcraft III.exe`, "-launch", "-windowmode", "windowed")
			cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000, HideWindow: true}
			err = cmd.Start()
			//cmd.Process.Pid
			if err != nil {
				log.Fatal(err)
			}

			prcss, err := process.NewProcess(int32(cmd.Process.Pid))
			if err != nil {
				log.Fatal(err)
			}
			//prcss.ForegroundWithContext()

			if isRunning, _ := prcss.IsRunning(); isRunning && prcss.Pid != 0 {

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

		blizzardBrowserPID = findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "width")
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

		blizzardRenderPID := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=ren")
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
		blizzard3DPID := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu")
		if blizzard3DPID == 0 {
			err = dlg.Text(fmt.Sprintf("Attempt %v: Browser GPU subprocess was not loaded. Rechecking in 2 seconds...", executionAttempts))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(2 * time.Second)
			blizzard3DPID2ndAttempt := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu")
			if blizzard3DPID2ndAttempt == 0 {
				err = dlg.Text(fmt.Sprintf("Attempt %v: Browser GPU subprocess was not loaded. Let's repeat it all over again Warcraft III", executionAttempts))
				if err != nil {
					log.Fatal(err)
				}
				time.Sleep(100 * time.Millisecond)

				blizzardBrowserProcess, err := process.NewProcess(warcraftProcessPID)
				if err != nil {
					log.Fatal(err)
				}

				if err = blizzardBrowserProcess.Kill(); err != nil {
					log.Fatal(err)
				}
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

		if blizzard3DPID2ndAttempt := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu"); blizzard3DPID2ndAttempt != 0 {
			break
		}

		if err := dlg.Text(fmt.Sprintf("Attempt %v: Browser GPU subprocess was not loaded. Let's repeat it all over again Warcraft III", executionAttempts)); err != nil {
			log.Fatal(err)
		}

		blizzardBrowserProcess, err := process.NewProcess(warcraftProcessPID)
		if err != nil {
			log.Fatal(err)
		}

		if err := blizzardBrowserProcess.Kill(); err != nil {
			log.Fatal(err)
		}
		continue
	}

	if err := dlg.Text(fmt.Sprintf("Warcraft III loaded in %v attempts. Please, have fun!", executionAttempts)); err != nil {
		log.Fatal(err)
	}

	if err := dlg.Complete(); err != nil {
		log.Fatal(err)
	}

	gameIsLoadedSuccessfully = true
	go func() {
		if err := zenity.Notify(
			fmt.Sprintf("Game is ready to play in %v attempts! Just give me Go play it! to try to go to windowed fullscreen and to close the app then", executionAttempts),
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

	dlgWhenFinished.Text(fmt.Sprintf("Game is ready to play in %v attempts! Just give me Go play it! to try to go to windowed fullscreen and to close the app then", executionAttempts))

	dlgWhenFinished.Complete()

	var wg sync.WaitGroup
	wg.Add(1)
	done := dlgWhenFinished.Done()
	go func() {
		if _, ok := <-done; !ok {
			if err := dlgWhenFinished.Text("Showing and maximizing Warcraft III..."); err != nil {
				log.Fatal(err)
			}

			warcraftProcessPID = findProcessByExeNameAndCmdLineSubStr("Warcraft III.exe", "")
			if warcraftProcessPID == 0 {
				if err := dlgWhenFinished.Text("Warcraft III has disappeared. Don't know why..."); err != nil {
					log.Fatal(err)
				}
				log.Fatal("Warcraft III has disappeared. Don't know why...")
			}

			robotgo.MinWindow(warcraftProcessPID)
			robotgo.MaxWindow(warcraftProcessPID)

			if err := dlgWhenFinished.Text("Switching to Warcraft III and sending alt + enter the window to turn fullcreen windowed mode on..."); err != nil {
				log.Fatal(err)
			}

			robotgo.KeyDown("alt")
			robotgo.KeyDown("enter")
			robotgo.KeyUp("alt")
			robotgo.KeyUp("enter")

			os.Exit(0)
		}
	}()
	wg.Wait()
	os.Exit(0)
}
