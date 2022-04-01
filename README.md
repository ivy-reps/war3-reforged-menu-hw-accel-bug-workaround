# Workaround of Warcraft III: Reforged menu freezing/hanging
Warcraft III: Reforged have some kind of the bug, that turns to freezing in main menu and freezing while game play. Somehow, while trying to start the game with other video preferences, to find the best mode to have some fun, I just find that at some starts the game starting without lags. Can't imagine why I've spent so much time with that. But that testing have moved me to conclusion, that sometimes the game starting with normally working GPU rendering of Chromium which integrated in Warcraft III - Reforged, and renders the scene and menu. After a day of reruns of the client, I just finally realized the though, that may be it will be possible to rerun the game allover again until it will be initialized with support of 3D Hardware Acceleration, because unsupportion of that thing, is the cause why you have menu with lags.

So, I just have application written. I've written it on Golang. Just put it on folder with Warcraft III.exe executable folder, and execute. It will rerun the game until BlizzardBrowser.exe process will not be started with subprocess which are providing of hardware accelerated rendering.

Feel free to build it by yourself. I made binaries on my system. Also feel free to test it on VirusTotal.com ( [latest check](https://www.virustotal.com/gui/file/7e66e117ef4dab8cb864d277f90616f0f1d71b168c7e3c93fb39a17167bbfffb?nocache=1) ) or with any other antivirus.

## What I did to understand what prevents the game from starting without fps slowed down
I tried to kill BlizzardBrowser process many times while Warcraft III main process working because I was thinking that the problem in BlizzardBrowser process exactly and I was thinking that if it will be restarted then at some restart BlizzardBrowser process with type=gpu-process which provides hardware accelerated rendereing will not be crashed and executed successfully (If you kill BlizzardBrowser process while Warcraft is working, Warcraft will restart Blizzard Browser)
It is all was invain.

Then I just applied that practice to Warcraft III executable. I've written application in Golang programming language, which rerun Warcraft allover again, if after BlizzardBrowser subprocess start will be no BlizzardBrowser subprocess which provides gpu rendering. And it is working! Sometimes it needs 6 attempts to restart the game, sometimes it needs 20 attempts of Warcraft to be restarted to make that BlizzardBrowser and Warcraft tondem to work normally.
Also I'm thinking that the problem is not in BlizzardBrowser process, but exactly in Warcraft executable. I'm thinking so, because you can terminate BlizzardBrowser alloveragain in your task manager and it never will be restarted with hardware acceleration support, but if you restarting Warcraft itself, at some rerun, BlizzardBrowser will be started with Hardware acceleration usage, because BlizzardBrowser subprocess will not crash.

## Message on Blizzard forum
I placed my report about such way to run game normally here https://us.forums.blizzard.com/en/warcraft3/t/extremely-low-fps-on-menus/19685/176, but that message was removed without an explanation and notification.

## A kind of disclaimer
1. **Be ready that while reruns the game window will flash repeatedly. I advise people with epilepsy to open something dark before starting that utility so that the disappearing of the game black window will does not contrast with a white or other bright background behind the game frame.**
2. There is no guarantee, that the game will work stable at any case, with or without hardware acceleration supported BlizzardBrowser execution.
3. There is no guarantee that my approach and my app will work in your case.
4. Also, I can't give you a guarantee that the system will not hang or will not go to BSOD (blue screen of death) after the game will be restart for so many times in a short period of time. For example from my point of view, nvidia drivers the most instable last years (IMHO). Maybe I will be wrong in this statement.
5. Please feel free to check my file with your antivirus, be with actual antivirul databases. Also you can use virustotal service. You can search VirusTotal by your own in the google for example.
Here are the results of latest check of binary in virustotal. Sadly link to VirusTotal aren't allowed here, so I can share only part of the link with you guys:
*https://www.virustotal.com/gui/file/7e66e117ef4dab8cb864d277f90616f0f1d71b168c7e3c93fb39a17167bbfffb*
6. **Blizzard, please feel free to check my code. There is no nothing to hide.**
7. Please keep in mind that antiviral products presented on VirusTotal works in paranoid mode, which means that they can treat almost any application as threat. From my point of view best approach will be to look on what the most antiviral products are saying, like Kaspersky, DrWeb, Symantec, McAffe.
8. You can compile that executable by your own anyway, or ask a friend to do that.
You can freely fork it if you need. I hope I will be ready to mke fixes. To request them, please use Issues tab of repository.
																																						   
This is my first GitHub repository shown to the peoples and I'm sorry if it is prepared badly. That repositiory contains the source code of that utility.
Here is the link to executable that I've compiled and placed there on the GitHub:
**[https://github.com/ivy-reps/war3-reforged-menu-hw-accel-bug-workaround/releases/download/v.0.3-pre-release/Try-to-start-Warcraft-III-with-Hardware-Acceleration.exe](https://github.com/ivy-reps/war3-reforged-menu-hw-accel-bug-workaround/releases/download/v.0.3-pre-release/Try-to-start-Warcraft-III-with-Hardware-Acceleration.exe)**
Executable have compiled in Windows 7 environment. I can't say whether that will work in Windows 8 and Windows 10 or not.

## What you can do to start the game normally
As I've said I've written utility/application in Golang. That utility have no reverse engineering of the game itself and have no anything bad made to Blizzard licence agreement. No hacking. No game process memory interruption/intrusion. It gathers no information. It just starts Warcraft executable and restarts it if no hardware acceration BlizzardBrowser subprocess after few seconds of Warcraft start.

Just put that file in subdirectory `_retail_\x86_64` of the game main directory, and run.
That executable should be placed to the game executable folder
For example the game installed here:
`Q:\Games\Blizzard Entertainment\Warcraft III`

Then the file should be placed here:
`Q:\Games\Blizzard Entertainment\Warcraft III\_retail_\x86_64`

Then just open *Try-to-start-Warcraft-III-with-Hardware-Acceleration.exe* and wait until app will be rerunning the game allover again until it will start normally with no freeze.

At my case, my system is Windows 7. And I have almost obsolete PC of 2008 year, but the game works fine except that lighting effect. If lighting effect have turned on, the game became incredible laggy, but I think, with now days video devices will be no such problem as it happend in my case.
I think there are no need to make them, because the game starts at some point as it expected, and it's enought even if I made code it dirty and roughly. I'm ready to fix something anyway.

## cc1.exe: sorry, unimplemented: 64-bit mode not compiled in
If you compiling in Windows, you have to install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) and specify path to that compilers bin folder right before command will be specified, or in Windows Environment Variables configuration like this: 
*PATH=C:\TDM-GCC-64\bin;%PATH%*

Libraries used:
* [Golang Desktop Automation. Control the mouse, keyboard, bitmap and image, read the screen, process, Window Handle and global event listener.](github.com/go-vgo/robotgo) - Used to minimize Warcraft II while restarting
* [gopsutil: psutil for golang](https://github.com/shirou/gopsutil/process) - Provide processes iteration and easier access to processes properties (process pid, process start command line)

