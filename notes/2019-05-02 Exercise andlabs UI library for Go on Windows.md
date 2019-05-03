## 2019-05-02 Exercise andlabs UI library for Go on Windows

### Background

- Still looking for a lightweight, cross-platform GUI toolkit for Go.
- Would like to distribute static binaries go GUI apps, and cross-compile for all platforms from any particular platform
- The [andlabs/ui](https://github.com/andlabs/ui) library for Go apparently provides these capabilities.  
- It is a Go wrapper for the [andlabs/libui](https://github.com/andlabs/libui) C++ library.

### Installed and andlabs/ui

- Followed installation instructions:

    ```console
    PS C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs> go get github.com/andlabs/ui/...
    ```
- Noted that gcc is required for using andlabs/ui, but gcc is already available via the PATH variable in PowerShell because Qt was installed previously.

### Built and ran Hello World application
- Found [code for app](https://learning.oreilly.com/library/view/hands-on-gui-application/9781789138412/2dcffd1b-bff9-42b9-8848-6d7ac05cfc8b.xhtml) in *Hand-On GUI Application Development in Go*:

    ```go
    package main
    
    import _ "github.com/andlabs/ui/winmanifest"
    
    import (
    	"github.com/andlabs/ui"
    )
    
    func main() {
    
    	err := ui.Main(func() {
    
    		window := ui.NewWindow("Hello", 100, 50, false)
    		window.SetMargined(true)
    		window.OnClosing(func(*ui.Window) bool {
    			ui.Quit()
    			return true
    		})
    
    		button := ui.NewButton("Quit")
    		button.OnClicked(func(*ui.Button) {
    			ui.Quit()
    		})
    
    		box := ui.NewVerticalBox()
    		box.Append(ui.NewLabel("Hello World!"), false)
    		box.Append(button, true)
    
    		window.SetChild(box)
    		window.Show()
    	})
    
    	if err != nil {
    		panic(err)
    	}
    }
    ```
- Noted that the import of `github.com/andlabs/ui/winmanifest` is required for the application to do anything on Windows.  Without it, the program exits when using `go run`:

    ```console
    PS C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs> go run .\hello_andlabs.go
    exit status 3221225785
    ```

-  Using `go build` without including the above import produces an executable without any error or warning messages, but the program exits as soon the executable is run:

    ```console
    PS C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs> dir *.exe
    
        Directory: C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs
    
    Mode                LastWriteTime         Length Name
    ----                -------------         ------ ----
    -a----         5/2/2019  10:31 PM       13829643 hello_andlabs.exe
    
    PS C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs> . .\hello_andlabs.exe
    PS C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs>
    ```
- With the import present, the program runs and behaves as expected.


### Failed to install andlab/ui and build app on Windows system without gcc

- Attempted to build the above app on Metis and got error message that andlabs/ui cannot be found:

    ```console
    PS C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs> go build .\hello_andlabs.go
    hello_andlabs.go:6:2: cannot find package "github.com/andlabs/ui" in any of:
            C:\Go\src\github.com\andlabs\ui (from $GOROOT)
            C:\Users\tmcphill\go\src\github.com\andlabs\ui (from $GOPATH)
    ```

- Installed andlabs/ui but during install got error stating that gcc was not installed:

    ```console
    PS C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs> go get github.com/andlabs/ui/...
    # github.com/andlabs/ui
    exec: "gcc": executable file not found in %PATH%-
	```
- Attempting to build the app now also gives the same error message:
    ```console
    PS C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs> go build .\hello_andlabs.go
    # github.com/andlabs/ui
    exec: "gcc": executable file not found in %PATH%
    ```
### Installed gcc for Windows by installing MinGW

- Installed MinGW via Chocolatey:

    ```console
    PS C:\WINDOWS\system32> choco install mingw
    Chocolatey v0.10.13
    2 validations performed. 1 success(es), 1 warning(s), and 0 error(s).
    
    Validation Warnings:
     - A pending system reboot request has been detected, however, this is
       being ignored due to the current Chocolatey configuration.  If you
       want to halt when this occurs, then either set the global feature
       using:
         choco feature enable -name=exitOnRebootDetected
       or pass the option --exit-when-reboot-detected.
    
    Installing the following packages:
    mingw
    By installing you accept licenses for the packages.
    Progress: Downloading mingw 8.1.0... 100%
    
    mingw v8.1.0 [Approved]
    mingw package files install completed. Performing other installation steps.
    The package mingw wants to run 'chocolateyinstall.ps1'.
    Note: If you don't run this script, the installation will fail.
    Note: To confirm automatically next time, use '-y' or consider:
    choco feature enable -n allowGlobalConfirmation
    Do you want to run the script?([Y]es/[N]o/[P]rint): y
    
    Downloading mingw 64 bit
      from 'https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win64/Personal%20Builds/mingw-builds/8.1.0/threads-posix/seh/x86_64-8.1.0-release-posix-seh-rt_v6-rev0.7z/download'
    Progress: 100% - Completed download of C:\Users\tmcphill\AppData\Local\Temp\chocolatey\mingw\8.1.0\x86_64-8.1.0-release-posix-seh-rt_v6-rev0.7z (47.08 MB).
    Download of x86_64-8.1.0-release-posix-seh-rt_v6-rev0.7z (47.08 MB) completed.
    Hashes match.
    Extracting C:\Users\tmcphill\AppData\Local\Temp\chocolatey\mingw\8.1.0\x86_64-8.1.0-release-posix-seh-rt_v6-rev0.7z to C:\ProgramData\chocolatey\lib\mingw\tools\install...
    C:\ProgramData\chocolatey\lib\mingw\tools\install
    PATH environment variable does not have C:\ProgramData\chocolatey\lib\mingw\tools\install\mingw64\bin in it. Adding...
    Environment Vars (like PATH) have changed. Close/reopen your shell to
     see the changes (or in powershell/cmd.exe just type `refreshenv`).
     ShimGen has successfully created a shim for addr2line.exe
     ShimGen has successfully created a shim for ar.exe
     ShimGen has successfully created a shim for as.exe
     ShimGen has successfully created a shim for c++.exe
     ShimGen has successfully created a shim for c++filt.exe
     ShimGen has successfully created a shim for cpp.exe
     ShimGen has successfully created a shim for dlltool.exe
     ShimGen has successfully created a shim for dllwrap.exe
     ShimGen has successfully created a shim for dwp.exe
     ShimGen has successfully created a shim for elfedit.exe
     ShimGen has successfully created a shim for g++.exe
     ShimGen has successfully created a shim for gcc-ar.exe
     ShimGen has successfully created a shim for gcc-nm.exe
     ShimGen has successfully created a shim for gcc-ranlib.exe
     ShimGen has successfully created a shim for gcc.exe
     ShimGen has successfully created a shim for gcov-dump.exe
     ShimGen has successfully created a shim for gcov-tool.exe
     ShimGen has successfully created a shim for gcov.exe
     ShimGen has successfully created a shim for gdb.exe
     ShimGen has successfully created a shim for gdborig.exe
     ShimGen has successfully created a shim for gdbserver.exe
     ShimGen has successfully created a shim for gendef.exe
     ShimGen has successfully created a shim for genidl.exe
     ShimGen has successfully created a shim for genpeimg.exe
     ShimGen has successfully created a shim for gfortran.exe
     ShimGen has successfully created a shim for gprof.exe
     ShimGen has successfully created a shim for ld.bfd.exe
     ShimGen has successfully created a shim for ld.exe
     ShimGen has successfully created a shim for ld.gold.exe
     ShimGen has successfully created a shim for mingw32-make.exe
     ShimGen has successfully created a shim for nm.exe
     ShimGen has successfully created a shim for objcopy.exe
     ShimGen has successfully created a shim for objdump.exe
     ShimGen has successfully created a shim for ranlib.exe
     ShimGen has successfully created a shim for readelf.exe
     ShimGen has successfully created a shim for size.exe
     ShimGen has successfully created a shim for strings.exe
     ShimGen has successfully created a shim for strip.exe
     ShimGen has successfully created a shim for widl.exe
     ShimGen has successfully created a shim for windmc.exe
     ShimGen has successfully created a shim for windres.exe
     ShimGen has successfully created a shim for x86_64-w64-mingw32-c++.exe
     ShimGen has successfully created a shim for x86_64-w64-mingw32-g++.exe
     ShimGen has successfully created a shim for x86_64-w64-mingw32-gcc-8.1.0.exe
     ShimGen has successfully created a shim for x86_64-w64-mingw32-gcc-ar.exe
     ShimGen has successfully created a shim for x86_64-w64-mingw32-gcc-nm.exe
     ShimGen has successfully created a shim for x86_64-w64-mingw32-gcc-ranlib.exe
     ShimGen has successfully created a shim for x86_64-w64-mingw32-gcc.exe
     ShimGen has successfully created a shim for x86_64-w64-mingw32-gfortran.exe
     ShimGen has successfully created a shim for cc1.exe
     ShimGen has successfully created a shim for cc1plus.exe
     ShimGen has successfully created a shim for collect2.exe
     ShimGen has successfully created a shim for f951.exe
     ShimGen has successfully created a shim for lto-wrapper.exe
     ShimGen has successfully created a shim for lto1.exe
     ShimGen has successfully created a shim for fixincl.exe
     ShimGen has successfully created a shim for gdbmtool.exe
     ShimGen has successfully created a shim for gdbm_dump.exe
     ShimGen has successfully created a shim for gdbm_load.exe
     ShimGen has successfully created a shim for python.exe
     ShimGen has successfully created a shim for python2.7.exe
     ShimGen has successfully created a shim for python2.exe
     ShimGen has successfully created a shim for wininst-6.0.exe
     ShimGen has successfully created a shim for wininst-7.1.exe
     ShimGen has successfully created a shim for wininst-8.0.exe
     ShimGen has successfully created a shim for wininst-9.0-amd64.exe
     ShimGen has successfully created a shim for wininst-9.0.exe
     ShimGen has successfully created a shim for ar.exe
     ShimGen has successfully created a shim for as.exe
     ShimGen has successfully created a shim for dlltool.exe
     ShimGen has successfully created a shim for ld.bfd.exe
     ShimGen has successfully created a shim for ld.exe
     ShimGen has successfully created a shim for ld.gold.exe
     ShimGen has successfully created a shim for nm.exe
     ShimGen has successfully created a shim for objcopy.exe
     ShimGen has successfully created a shim for objdump.exe
     ShimGen has successfully created a shim for ranlib.exe
     ShimGen has successfully created a shim for readelf.exe
     ShimGen has successfully created a shim for strip.exe
     The install of mingw was successful.
      Software installed to 'C:\ProgramData\chocolatey\lib\mingw\tools\install'
    
    Chocolatey installed 1/1 packages.
     See the log for details (C:\ProgramData\chocolatey\logs\chocolatey.log).
     ```

- Checked version of gcc just installed:

    ```console
    PS C:\WINDOWS\system32> gcc --version
    gcc.exe (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0
    Copyright (C) 2018 Free Software Foundation, Inc.
    This is free software; see the source for copying conditions.  There is NO
    warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
    ```

### Successfully built and ran andlabs app

- Installed andlabs/ui without error:

    ```console
    PS C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs> go get github.com/andlabs/ui/...
    ```
- Built and ran hello app:

    ```console
    PS C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs> go build .\hello_andlabs.go
    PS C:\Users\tmcphill\go\src\go-demos\05_hello_andlabs> .\hello_andlabs.exe
    ```

