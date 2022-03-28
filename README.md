# Workaround of Warcraft III: Reforged menu freezing/hanging
Warcraft III: Reforged have some kind of the bug, that turns to freezing in main menu and freezing while game play. Somehow, while trying to start the game with other video preferences, to find the best mode to have some fun, I just find that at some starts the game starting without lags. After a day of reruns of the client, I just finally realized the though, that may be it will be possible to rerun the game allover again until it will be initialized with support of 3D Hardware Acceleration, because unsupportion of that thing, is the cause why you have menu with lags.

So, I just have application written. I've written it on Golang. Just put it on folder with Warcraft III.exe executable folder, and execute. It will rerun the game until BlizzardBrowser.exe process will not be started with subprocess which are providing of hardware accelerated rendering.

Feel free to build it by yourself. I made binaries on my system. Also feel free to test it on VirusTotal.com ( [latest check](https://www.virustotal.com/gui/file/8d60b7ab0f2cc661518ad7cec8d1f8b80915752a591fa05054491650af5b482d?nocache=1) ) or with any other antivirus.

Libraries used:
* [Golang Desktop Automation. Control the mouse, keyboard, bitmap and image, read the screen, process, Window Handle and global event listener.](github.com/go-vgo/robotgo) - Used to minimize Warcraft II while restarting
* [gopsutil: psutil for golang](https://github.com/shirou/gopsutil/process) - Provide processes iteration and easier access to processes properties (process pid, process start command line)

