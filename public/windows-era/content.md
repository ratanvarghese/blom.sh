![Windows logo](/windows-era/attachments/windows.svg) [Image credit: Microsoft](class:credit)

I used Windows uncomfortably often during the last year, much more than in the prior ten. It turns out, a lot of my expectations were obsolete. Windows has changed in many tiny ways over the past decade, but still finds ways to frustrate and confound.

### Windows Subsystem For Linux (WSL)

According to [Microsoft itself](https://docs.microsoft.com/en-us/windows/wsl/):

> The Windows Subsystem for Linux lets developers run a GNU/Linux environment -- including most command-line tools, utilities, and applications -- directly on Windows, unmodified, without the overhead of a traditional virtual machine or dual-boot setup.

In my months of experience, I found this statement to be mostly accurate. Obviously there isn't much hardware access, but there is at least file system and network access. A lot of my gripes with WSL are being actively worked on.

There are some interesting subtleties, even in Microsoft's official statement. Consider the word "developers". WSL isn't meant for most Windows "users", it is meant for "developers" to either work on Linux applications or use a Linux toolchain to cross-compile a Windows app. Creating an application that needs WSL on a user's machine is probably a hack.

There is only one Windows app I have used that literally requires WSL: a recent version of Docker. Docker's entire purpose is running Linux containers, so it is sensible for them to avoid duplicating WSL's functionality. Not so sensible for other applications, however.

Microsoft also decided to refer to the Linux userspace as "GNU/Linux", which is a bit odd considering that [Alpine Linux](https://www.reddit.com/r/gnu/comments/bqszff/is_alpine_gnulinux/) is one of the [officially supported options](https://www.microsoft.com/en-us/p/alpine-wsl/9p804crf0395?activetab=pivot:overviewtab). However, if you're using WSL because the Windows Command Prompt is missing features you want, then Alpine is an annoyingly minimalist choice. I ended up using Ubuntu WSL more often than Alpine WSL.

WSL doesn't just enable "GNU/Linux" applications to run on Windows. It also allows applications in your Windows PATH to run in WSL. [For example you can type `explorer.exe .` to run Windows Explorer from inside WSL.](https://www.howtogeek.com/426749/how-to-access-your-linux-wsl-files-in-windows-10/) Of course not all Windows applications were designed to be run this way.

### Changing the Windows PATH

![Windows interface for changing PATH](/windows-era/attachments/path.jpg) [Image credit: Ratan Abraham Varghese](class:credit)

The interface for changing the Windows PATH is actually really nifty. I wish more Unix-based systems had a similar window. However the obvious problem is that there isn't just one PATH, is there? Each shell and application might be making alterations its PATH while it is running, even on Windows.

Also, the only people who even care about the PATH are command-line users, who should be okay with reading and writing text files anyway. Or should they? Sometimes I wonder if these sorts of assumptions are holding back the possible capabilities of our software.

### Shells, Consoles, Terminals

![Windows Terminal with 13 shells](/windows-era/attachments/shell.jpg) [Image credit: Ratan Abraham Varghese](class:credit)

If you use command-lines often, you likely understand that a modern command-line involves a number of applications working together. Details such as font size and color combinations are controlled seperately from details such as how the computer responds to `echo Hello`.

Most articles on this topic use the terms "shell", "console" and "terminal". To quote [this answer on SuperUser](https://superuser.com/a/144668):

> The **shell** is the program which actually processes commands and returns output. Most shells also manage foreground and background processes, command history and command line editing...

> A **terminal** refers to a wrapper program which runs a shell. Decades ago, this was a physical device consisting of little more than a monitor and keyboard. As unix/linux systems added better multiprocessing and windowing systems, this terminal concept was abstracted into software.

Typical shells on a Linux system include `bash`, `zsh`, `fish` and `ash`. Typical terminals include Gnome Terminal, `xterm`, and `urxvt`.

So what is the shell/terminal situation on Windows like?

The default shell is of course the [Command Prompt](https://en.wikipedia.org/wiki/Cmd.exe). Like many Microsoft products, it is designed to be compatible with ancient software that someone, somewhere cannot bear to abandon. Apart from being available on all Windows computers since the dawn of the NT kernel, I cannot think of a redeeming quality of Command Prompt compared to alternatives.

There is also the much more interesting [Powershell](https://en.wikipedia.org/wiki/PowerShell). Unlike typical Unix shells where every command inputs and outputs text, in PowerShell every command inputs and outputs an object. This is pretty cool, but there are several caveats.

For one thing, PowerShell is entirely different from typical Unix shells. This in itself is problem for a lot of people. Interestingly, Microsoft has ported "PowerShell Core" to non-Windows systems. Perhaps PowerShell could become someone's default Unix shell someday.

That would be deeply ironic because on Windows, PowerShell cannot execute scripts by default. Scripting must be explicitly enabled by the user. [This is an entirely intentional feature.](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_execution_policies?view=powershell-7.1) Unfortunately this means as an application developer, you cannot rely on PowerShell always being available on a user's Windows computer.

There are also "Developer Command Prompts" available with Visual Studio: one for every combination of target and tool CPU architecture. No, I'm serious, check out [this page](https://docs.microsoft.com/en-us/cpp/build/building-on-the-command-line?view=msvc-160) which lists the following:

> + Developer Command Prompt - Sets the environment to use 32-bit, x86-native tools to build 32-bit, x86-native code.
> + x86 Native Tools Command Prompt - Sets the environment to use 32-bit, x86-native tools to build 32-bit, x86-native code.
> + x64 Native Tools Command Prompt - Sets the environment to use 64-bit, x64-native tools to build 64-bit, x64-native code.
> + x86_x64 Cross Tools Command Prompt - Sets the environment to use 32-bit, x86-native tools to build 64-bit, x64-native code.
> + x64_x86 Cross Tools Command Prompt - Sets the environment to use 64-bit, x64-native tools to build 32-bit, x86-native code.

And to add insult to injury, all five of those are based on Command Prompt instead of PowerShell. The Developer PowerShell is [seperate](https://docs.microsoft.com/en-us/visualstudio/ide/reference/command-prompt-powershell?view=vs-2019).

Other shells are available from third parties. For example, Git is typically installed with [Git Bash](https://gitforwindows.org/), bringing the Unix `bash` shell to Windows. [MSYS2](https://www.msys2.org/) really changed the game by bringing the Unix `bash` shell to Windows. But [Cygwin](https://cygwin.com/index.html) also still exists, providing the ability to run `bash` on Windows.

Oh, and of course WSL instances allow "developers" to run `bash` on Windows.

The situation with terminals on Windows also strange. [Windows Console](https://en.wikipedia.org/wiki/Windows_Console) is the default terminal for Command Prompt, PowerShell and numerous other applications. Luckily, I didn't have to use it long enough to hate it.

The newer [Windows Terminal](https://en.wikipedia.org/wiki/Windows_Terminal) adds a lot of missing features, but with a twist: there isn't a graphical settings dialog. [Windows Terminal is currently configured by editing a JSON file](https://garrytrinder.github.io/2020/12/my-wsl2-windows-terminal-setup), which isn't terrible in and of itself.

Also, Windows Terminal isn't shipped with Windows, but provided as a download from the Microsoft Store. By contrast, the Windows Console will continue shipping for ["decades to come"](https://devblogs.microsoft.com/commandline/introducing-windows-terminal/). This means that much like PowerShell, application developers cannot rely on Windows Terminal existing on a user's machine. So of course they are all still defaulting to using Windows Console.

The end result is that if *I* want to use Windows Terminal for all my shells, *I* must edit the Windows Terminal JSON whenever I install a new shell that isn't "blessed" by the Terminal team, such as Git Bash and the Visual Studio shell prompts...

This is frankly a bit ridiculous. Can't all the teams at Microsoft get together and agree on *one* shell in *one* terminal with all the necessary features?

### Ports of Unix Favourites

Let's say you want to play [*NetHack*](https://nethack.org/) on a [public server](https://alt.org/nethack/) using SSH, but your home computer runs Windows.

Conventional knowledge might lead you to install [PuTTY](https://www.chiark.greenend.org.uk/~sgtatham/putty/). Or perhaps since you've kept up with all the latest development trends, you could opt to use SSH inside of WSL instead.

Or... you could just type `ssh nethack@alt.org` into the default Command Prompt just like on Linux.

![SSH in Command Prompt](/windows-era/attachments/ssh.jpg) [Image credit: Ratan Abraham Varghese](class:credit)

Yes, that's right. You can even copy your `.ssh` folder directly from Linux and plop it in your Windows user folder, because [since about 2018, Microsoft has been bundling OpenSSH with Windows 10](https://www.howtogeek.com/336775/how-to-enable-and-use-windows-10s-built-in-ssh-commands/).

They also started [bundling `tar` and `curl`](https://techcommunity.microsoft.com/t5/containers/tar-and-curl-come-to-windows/ba-p/382409). This is a major convenience: before Microsoft started bundling `curl`, you might have needed to download files using [this strange thing](https://docs.microsoft.com/en-us/windows/win32/bits/background-intelligent-transfer-service-portal), or at least its various PowerShell wrappers. As for `tar`, opening its file format used to require third-party tools like 7-Zip.

Is there any more Unix software being discreetly bundled into recent releases of Windows? I could trawl through Microsoft's entire development blog, actually read the release notes of every Windows update and ... on second thoughts I'll just search for tools when I need them.

### WinGet Package Manager

Microsoft is finally making a package manager for Windows, known as `winget`. [Uninstalling a package is still an experimental feature](https://beebom.com/how-uninstall-apps-using-winget-windows-package-manager/). 

For a home user who isn't installing applications on a hundred Windows machines, I don't think `winget` (or other Windows package managers such as [Chocolately](https://chocolatey.org/)) are as useful as their Unix equivalents. In most cases Windows applications do not depend on other Windows applications, and are bundled with all their libraries.

Of course the situation is different when considering developer-focused package managers like [`nuget`](https://www.nuget.org/), since developers do need to keep track of dependencies in their projects.

Also, if Windows Terminal, WSL, `ssh`, `curl` and the like were all being managed by `winget`, then `winget` could treat those as dependencies for some larger utility program.

### Reset Your PC

It's pretty easy to reset or reinstall Windows these days, which is great if there are a lot of useless files and programs to be removed. However there is a catch: all software bundled with the PC is reinstalled.

Yes, [both the "Refresh your PC" and "Reset your PC" options keeps the apps that came with the PC](https://support.microsoft.com/en-us/windows/how-to-refresh-reset-or-restore-your-pc-51391d9a-eb0a-84a7-69e4-c2c1fbceb8dd). Did your PC come with the Netflix Windows app, random games, and pointless OEM shovelware preinstalled? Say hello to them again!

That said this is still a convenient feature. If you have a backup ready to go, you can return your PC to factory settings with ease.

### Backing Up

[Oh of course Windows supplies two seperate backup utilities.](https://superuser.com/questions/998785/windows-10-difference-between-file-history-and-backup-and-restore)

[Neither of which are all that great.](https://superuser.com/questions/1133653/)

Honestly I haven't been keeping many backups, for reasons independent of software. The external hard disk that I would normally use for backups was extremely unreliable. It would just disconnect at the slightest disturbance. I had to basically cradle the hard disk like a baby for it to actually maintain a connection to the computer. Recently I got an external SSD which will hopefully be more convenient at this critical task.

### Microsoft PowerToys

One of the features of Linux I was missing was the ability to install a [tiling window manager](https://en.wikipedia.org/wiki/Tiling_window_manager) such as `xmonad`, `i3` or perhaps `sway`. The search for a Windows alternative is what led me to the [PowerToys](https://docs.microsoft.com/en-us/windows/powertoys/) utilties in general, and [FancyZones](https://docs.microsoft.com/en-us/windows/powertoys/fancyzones) in particular.

However, by default FancyZones doesn't force new windows into zones like a tiling window manager would. Instead the user must drag a window into a zone. It's easy to fall out of the habit. FancyZones also has a bit of trouble with applications running in Administrator mode.

There was another PowerToy that I found pretty handy: the [PowerRename](https://docs.microsoft.com/en-us/windows/powertoys/powerrename) utility. In a recent biology course, I didn't like the video filenames the professor chose, and PowerRename helped me rename them in bulk. On Linux, of course, there is the [`rename` command](https://man7.org/linux/man-pages/man1/rename.1.html) which provides all the same features, minus the graphical interface.

There are many other PowerToys that I never ended up using.

### Leaving Microsoft Windows

![Screenshot of Manjaro with i3](/windows-era/attachments/i3_screenshot.jpg) [Image credit: Ratan Abraham Varghese](class:credit)

It was the lack of a tiling window manager that really tempted me away from Windows. Moving around applications individually is a nuisance. Switching back to Linux during the middle of a school term seemed like a bad idea, so I bided my time.

However even on Linux, tiling window managers aren't *that* popular. I eventually chose the [`i3` community edition of Manjaro Linux](https://www.manjaro.org/downloads/community/i3/). But of course, that was a "community" edition, lacking "official" support. Eventually when [Wayland and pals](https://wiki.archlinux.org/title/wayland) are finally ready to dethrone X11, I might need to switch from `i3` to `sway`, which is of course even more obscure.

Before installing Manjaro, I reinstalled Windows and kept a tiny partition for it, just in case.

### Conclusions?

So is there an actual point to this list of links and gripes?

Clearly Microsoft is incrementally moving towards making their system more attractive to develop on. It is becoming possible to develop a Linux userland application on Windows without any third-party tools at all. People who primarily use desktop Linux to develop web or server applications might be tempted to just use Windows.

There are still plenty of reasons to use desktop Linux. The most industry-relevant reason is to develop a Linux application that directly communicates with hardware.

Desktop Linux also provides the freedom to customize your Linux kernel and install alternative user interfaces (like tiling window managers). It allows you to really customize every tiny detail of your system, but that freedom is not really important or even accessible for most people. Is being able to edit the source code of Linux `rename` *that* compelling compared to actually knowing what you're doing with Microsoft PowerRename?

I can only hope that these moves by Microsoft encourage more creativity and innovation in Linux's graphical interfaces. Creating (more?) replicas of the Windows user interface would be a really poor way for desktop Linux to compete with WSL.

Inter-team confusion continues to impede Microsoft's progress. Did the Developer Command Prompt people even know that Microsoft Terminal was being developed? [Did the WSL team ever meet the people bundling `ssh` into Windows?](https://devblogs.microsoft.com/commandline/sharing-ssh-keys-between-windows-and-wsl-2/) Will the PowerToys team ever be able to *really* change the Windows interface with a future version of FancyZones?

I suppose that from now on, I'll be keeping a closer eye on Windows, even if I avoid using it at home.