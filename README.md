# Workaround of Warcraft III: Reforged menu freezing/hanging
Warcraft III: Reforged have some kind of the bug, that turns to freezing in main menu and freezing while game play. Somehow, while trying to start the game with other video preferences, to find the best mode to have some fun, I just find that at some starts the game starting without lags. After a day of reruns of the client, I just finally realized the though, that may be it will be possible to rerun the game allover again until it will be initialized with support of 3D Hardware Acceleration, because unsupportion of that thing, is the cause why you have menu with lags.

So, I just have application written. I've written it on Golang. Just put it on folder with Warcraft III.exe executable folder, and execute. It will rerun the game until BlizzardBrowser.exe process will not be started with subprocess which are providing of hardware accelerated rendering.

Feel free to build it by yourself. I made binaries on my system. Also feel free to test it on VirusTotal.com ( [latest check](https://www.virustotal.com/gui/file/8d60b7ab0f2cc661518ad7cec8d1f8b80915752a591fa05054491650af5b482d?nocache=1) ) or with any other antivirus.

# How to use it
Just put that file in subdirectory *_retail_\x86_64* of the game main directory, and run.
For example Warcraft III: Reforged have installed into *C:\Program Files\Warcraft III - Reforged*, then you have to open that path, search for *_retail_* (C:\Program Files\Warcraft III - Reforged\_retail_) folder there and then for *x86_64* (C:\Program Files\Warcraft III - Reforged\_retail_\x86_64) in it. Place exe file from release sections on that folder and run. You have to run it every time you want to play the game, until Blizzard will not fix that bug.

## cc1.exe: sorry, unimplemented: 64-bit mode not compiled in
If you compiling in Windows, you have to install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) and specify path to that compilers bin folder right before command will be specified, or in Windows Environment Variables configuration like this: 
*PATH=C:\TDM-GCC-64\bin;%PATH%*

Libraries used:
* [Golang Desktop Automation. Control the mouse, keyboard, bitmap and image, read the screen, process, Window Handle and global event listener.](github.com/go-vgo/robotgo) - Used to minimize Warcraft II while restarting
* [gopsutil: psutil for golang](https://github.com/shirou/gopsutil/process) - Provide processes iteration and easier access to processes properties (process pid, process start command line)

