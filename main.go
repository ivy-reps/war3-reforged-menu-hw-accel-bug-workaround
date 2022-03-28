package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"


	"github.com/go-vgo/robotgo"
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

func writeLine(strs ...string) {
	fmt.Printf("\r"+strings.Join(strs, ","))
}

func main() {

	var executionAttempts int
	for {
		time.Sleep(1 * time.Second)
		warcraftProcessPID := findProcessByExeNameAndCmdLineSubStr("Warcraft III.exe", "")

		var blizzardBrowserPID int32
		if warcraftProcessPID == 0 {
			writeLine("WarcraftProcessPID not found")
			cmd := exec.Command(`Q:\Games\Blizzard Entertainment\Warcraft III\_retail_\x86_64\Warcraft III.exe`, "-launch")
			err := cmd.Start()
			//cmd.Process.Pid
			if err != nil {
				log.Fatal(err)
			}

			prcss, _ := process.NewProcess( int32(cmd.Process.Pid) )
			executionAttempts++

			if isRunning, _ := prcss.IsRunning(); isRunning && prcss.Pid != 0 {
				go func() {
					for {
						writeLine("Warcraft III process window should be minimized now")
						robotgo.MinWindow( prcss.Pid )
						if blizzardBrowserPID > 0 {
							break
						}
						time.Sleep(50 * time.Millisecond)
					}
				}()
			}
			continue
		}
		writeLine("warcraftProcessPID found")

		blizzardBrowserPID = findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "width")
		if blizzardBrowserPID == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		writeLine("blizzardBrowserPID found")

		blizzardRenderPID := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=ren")
		if blizzardRenderPID == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		writeLine("blizzardRenderPID found")

		time.Sleep(2 * time.Second)
		blizzard3DPID := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu")
		if blizzard3DPID == 0 {
			writeLine("blizzard3DPID not found, rechecking")
			time.Sleep(2 * time.Second)
			blizzard3DPID2ndAttempt := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu")
			if blizzard3DPID2ndAttempt == 0 {
				writeLine("blizzard3DPID not found, finishing BlizzardBrowser process")
				blizzardBrowserProcess, _ := process.NewProcess(warcraftProcessPID)
				blizzardBrowserProcess.Kill()
				continue
			}
		}
		writeLine("blizzard3DPID found, rechecking")
		time.Sleep(1 * time.Second)
		blizzard3DPID2ndAttempt := findProcessByExeNameAndCmdLineSubStr("BlizzardBrowser.exe", "type=gpu")
		if blizzard3DPID2ndAttempt == 0 {
			writeLine("blizzard3DPID not found killing process")
			blizzardBrowserProcess, _ := process.NewProcess(warcraftProcessPID)
			blizzardBrowserProcess.Kill()
			continue
		}

		writeLine(fmt.Sprintf("%v attempts of game execution was made, to wait until gpu renderer will not crash.", executionAttempts))
		os.Exit(0)

	}
}
