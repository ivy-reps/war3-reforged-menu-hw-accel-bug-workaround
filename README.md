# Workaround of Warcraft III: Reforged menu freezing/hanging
Warcraft III: Reforged have some kind of the bug, that turns to freezing in main menu and freezing while game play. Somehow, while trying to start the game with other video preferences, to find the best mode to have some fun, I just find that at some starts the game starting without lags. After a day of reruns of the client, I just finally realized the though, that may be it will be possible to rerun the game allover again until it will be initialized with support of 3D Hardware Acceleration, because unsupportion of that thing, is the cause why you have menu with lags.

So, I just have application written. I've written it on Golang. Just put it on folder with Warcraft III.exe executable folder, and execute. It will rerun the game until BlizzardBrowser.exe process will not be started with subprocess which are providing of hardware accelerated rendering.

Feel free to build it by yourself. I made binaries on my system. Also feel free to test it on VirusTotal.com or wuth any other antivirus.
