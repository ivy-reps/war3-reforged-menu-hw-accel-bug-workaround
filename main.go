package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/ncruces/zenity"
	"github.com/shirou/gopsutil/process"
)

func findProcessByExeNameAndCmdLineSubStr(exeName string, cmdLine string) int32 {
	processes, e := process.Processes()
	if e != nil {

	}

	for _, process := range processes {
		exe, e := process.Exe()
		if e != nil {

		}

		if strings.Contains(exe, exeName) {
			processCmdLine, e := process.Cmdline()
			if e != nil {

			}
			if cmdLine == "" {
				return process.Pid
			}
			if strings.Contains(processCmdLine, cmdLine) {
				return process.Pid
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
	dlg, err := zenity.Progress(
		zenity.Title("Retrying to start Warcraft III with HW Acceleration support"),
		zenity.Pulsate(), zenity.NoCancel(),
		zenity.OKLabel("Stop it!"),
	)
	if err != nil {
		return
	}
	defer dlg.Close()

	dlg.Text("Checking if Warcraft III started already")
	dlg.Value(0)

	var wg sync.WaitGroup
	wg.Add(1)

	done := dlg.Done()
	go func() {
		if _, ok := <-done; !ok {
			if !gameIsLoadedSuccessfully {
				if warcraftProcessPID := findProcessByExeNameAndCmdLineSubStr("Warcraft III.exe", ""); warcraftProcessPID > 0 {
					prcss, _ := process.NewProcess(warcraftProcessPID)
					prcss.Kill()
				}
			}
			os.Exit(0)
		}
	}()

	var executionAttempts int
	for {

		time.Sleep(time.Second)
		var blizzardBrowserPID int32

		warcraftProcessPID := findProcessByExeNameAndCmdLineSubStr("Warcraft III.exe", "")
		if warcraftProcessPID == 0 {
			dlg.Value(10)

			dlg.Text("Starting Warcraft III...")
			cmd := exec.Command(`Warcraft III.exe`, "-launch")
			err := cmd.Start()
			//cmd.Process.Pid
			if err != nil {
				log.Fatal(err)
			}

			prcss, _ := process.NewProcess(int32(cmd.Process.Pid))
			executionAttempts++

			if isRunning, _ := prcss.IsRunning(); isRunning && prcss.Pid != 0 {
				dlg.Text("Warcraft III started.")
				dlg.Value(20)

				time.Sleep(time.Second)

				dlg.Text("Awaiting for BlizzardBrowser to be loaded... Meanwhile minimizing Warcraft III...")
				go func() {
					for {
						robotgo.MinWindow(prcss.Pid)
						if blizzardBrowserPID > 0 {
							break
						}
						time.Sleep(10 * time.Millisecond)
					}
				}()
			}
			continue
		}

		blizzardBrowserPID = findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "width")
		if blizzardBrowserPID == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		dlg.Text("BlizzardBrowser loaded... Awaiting for Browser renderer subprocess...")
		dlg.Value(30)

		blizzardRenderPID := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=ren")
		if blizzardRenderPID == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		dlg.Text("BlizzardBrowser renderer loaded... Awaiting for Browser GPU subprocess...")
		dlg.Value(40)

		time.Sleep(2 * time.Second)
		blizzard3DPID := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu")
		if blizzard3DPID == 0 {
			dlg.Text("Browser GPU subprocess was not loaded. Rechecking in 2 seconds...")
			time.Sleep(2 * time.Second)
			blizzard3DPID2ndAttempt := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu")
			if blizzard3DPID2ndAttempt == 0 {
				dlg.Text("Browser GPU subprocess was not loaded. Let's repeat it all over again Warcraft III")
				time.Sleep(100 * time.Millisecond)
				blizzardBrowserProcess, _ := process.NewProcess(warcraftProcessPID)
				blizzardBrowserProcess.Kill()
				continue
			}
		}
		dlg.Text("Browser GPU subprocess is loaded. Checking that it will not be crashed in 2 seconds")
		dlg.Value(50)
		time.Sleep(2 * time.Second)
		blizzard3DPID2ndAttempt := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu")
		if blizzard3DPID2ndAttempt == 0 {
			dlg.Text("Browser GPU subprocess was not loaded. Let's repeat it all over again Warcraft III")
			blizzardBrowserProcess, _ := process.NewProcess(warcraftProcessPID)
			blizzardBrowserProcess.Kill()
			continue
		}

		dlg.Text(fmt.Sprintf("Warcraft III loaded in %v attempts. Please, have fun!", executionAttempts))
		dlg.Complete()
		robotgo.MaxWindow(warcraftProcessPID)
		writeLine(fmt.Sprintf("%v attempts of game execution was made, to wait until gpu renderer will not crash.", executionAttempts))
		break
	}
	gameIsLoadedSuccessfully = true
	wg.Wait()
	os.Exit(0)
}
