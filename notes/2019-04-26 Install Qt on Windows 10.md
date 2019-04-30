## 2019-04-26 Install Qt on Windows 10

### Background

- I want to be able to develop cross-platform GUIs for Go applications.
- I can now install the Go build tools on Windows via Ansible.
- I will try to install Qt and its bindings to Go and exercise with a minimal application.
- I will try to automate the installation and configuration of Qt and its bindings to Go.

### Install Qt on Windows 10 with Visual Studio 2019 preinstalled

#### Downloaded Qt offline installer

- Read about online vs offline installations of Qt at [Getting Started with Qt](https://doc.qt.io/qt-5/gettingstarted.html).  Offline installations involve downloading a single installer and running it, and I expect this will be more conducive to automation.
- Found the Windows offline installer for Qt 5.12.3 (3.7 GB) at [Offline Qt Downloads](https://www.qt.io/offline-installers):
http://download.qt.io/official_releases/qt/5.12/5.12.3/qt-opensource-windows-x86-5.12.3.exe
- Tried downloading the Windows installer from PowerShell using Invoke-Request:

    ```console
    PS C:\temp> Invoke-WebRequest http://download.qt.io/official_releases/qt/5.12/5.12.3/qt-opensource-windows-x86-5.12.3.exe
    ```
- However, the above download is far too slow. Aborted that download and tried a nearby mirror at http://mirrors.ocf.berkeley.edu/qt/archive/qt/5.12/5.12.3/qt-opensource-windows-x86-5.12.3.exe, but this was slow, too.

- Aborted the second download and tried System.Net.WebClient.DownloadFile instad of Invoke-WebRequest:

	```console
	PS C:\temp> (New-Object System.Net.WebClient).DownloadFile("http://mirrors.ocf.berkeley.edu/qt/archive/qt/5.12/5.12.3/qt-opensource-windows-x86-5.12.3.exe", "qt-opensource-windows-x86-5.12.3.exe")
	```
- The above download completed in a few second, but strangely left the file in my Windows home directory.
- However, providing a fully qualified path as the second argument left the installer in C:\Temp in 182 seconds:

    ```console
    PS C:\temp> Measure-Command { (New-Object System.Net.WebClient).DownloadFile("http://mirrors.ocf.berkeley.edu/qt/archive/qt/5.12/5.12.3/qt-opensource-windows-x86-5.12.3.exe", "C:\Temp\qt-opensource-windows-x86-5.12.3.exe") }
    Days              : 0
    Hours             : 0
    Minutes           : 3
    Seconds           : 2
    Milliseconds      : 36
    Ticks             : 1820369026
    TotalDays         : 0.00210690859490741
    TotalHours        : 0.0505658062777778
    TotalMinutes      : 3.03394837666667
    TotalSeconds      : 182.0369026
    TotalMilliseconds : 182036.9026
    
    PS C:\temp> dir
    
        Directory: C:\temp
    
    Mode                LastWriteTime         Length Name
    ----                -------------         ------ ----
    -a----        4/26/2019   8:03 PM     3933849864 qt-opensource-windows-x86-5.12.3.exe
    ```

### Ran the Qt installer
- Executed the installer:

    ```console
    PS C:\Temp> .\qt-opensource-windows-x86-5.12.3.exe
    ```

- Accepted default location for installation:  
`C:\Qt\Qt5.12.3`

- Selected three components for installation:
	- Qt -> Qt 5.12.3
		- MinGW 7.3.0 64 bit (prebuilt components)
	- Qt -> Developer and Designer Tools
		- Qt Creator 4.9.0 (IDE for application development)
		- MinGW 7.3.0 64-bit (MinGW-builds 7.3.0 64-bit toolchain)

