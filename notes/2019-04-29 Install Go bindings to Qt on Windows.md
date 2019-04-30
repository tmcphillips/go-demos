## 2019-04-29 Install Go bindings to Qt on Windows

### Background
- The most commonly used way to use Qt with Go is [therecipe/qt](https://github.com/therecipe/qt).
- Installation instructions for Windows: 
https://github.com/therecipe/qt/wiki/Installation-on-Windows
Go version referenced is 1.12.4.

### Upgraded Go to match version in therecipe/qt instructions
- Version currently installed on circe-win10 is  1.12.1

    ```console
    PS C:\Users\tmcphill\go\src\github.com\therecipe\qt\cmd\qtsetup> go version
    go version go1.12.1 windows/amd64
    ```
- Updated the golang role to install latest version:

    ```yaml
    ---
    
    - name: download and install the Go MSI package using Chocolatey
      win_chocolatey:
        name: golang
        state: latest
    
    - name: create go workspace directories
      win_file:
        path: "{{ ansible_env.GOPATH }}/src"
        state: directory
    ```
- Ran the go-dev playbook to do the upgrade:

    ```console
    (.venv-ansible-playbooks) tmcphill@circe-win10:~/GitRepos/ansible-playbooks/windows$ ansible-playbook playbooks/go-dev.yml
    
    PLAY [install go development environment] **************************************************************************************************************************************************
    
    TASK [Gathering Facts] *********************************************************************************************************************************************************************
    ok: [127.0.0.1]
    
    TASK [git : download and install the Go MSI package using Chocolatey] **********************************************************************************************************************
    ok: [127.0.0.1]
    
    TASK [git : configure global git settings] *************************************************************************************************************************************************
    changed: [127.0.0.1]
    
    TASK [golang : download and install the Go MSI package using Chocolatey] *******************************************************************************************************************
    
    changed: [127.0.0.1]
    
    TASK [golang : create go workspace directories] ********************************************************************************************************************************************
    ok: [127.0.0.1]
    
    TASK [vscode : install Visual Studio Code] *************************************************************************************************************************************************
    ok: [127.0.0.1]
    
    PLAY RECAP *********************************************************************************************************************************************************************************
    127.0.0.1                  : ok=6    changed=2    unreachable=0    failed=0
    ```
- Confirmed that installed version is no 1.12.4:

    ```console
    PS C:\Users\tmcphill\go\src\github.com\therecipe\qt\cmd\qtsetup> go version
    go version go1.12.4 windows/amd64
    ```
- Examined Go environment:

    ```console
    PS C:\Users\tmcphill\go\src\github.com\therecipe\qt\cmd\qtsetup> go env
    set GOARCH=amd64
    set GOBIN=
    set GOCACHE=C:\Users\tmcphill\AppData\Local\go-build
    set GOEXE=.exe
    set GOFLAGS=
    set GOHOSTARCH=amd64
    set GOHOSTOS=windows
    set GOOS=windows
    set GOPATH=C:\Users\tmcphill\go
    set GOPROXY=
    set GORACE=
    set GOROOT=C:\Go
    set GOTMPDIR=
    set GOTOOLDIR=C:\Go\pkg\tool\windows_amd64
    set GCCGO=gccgo
    set CC=gcc
    set CXX=g++
    set CGO_ENABLED=1
    set GOMOD=
    set CGO_CFLAGS=-g -O2
    set CGO_CPPFLAGS=
    set CGO_CXXFLAGS=-g -O2
    set CGO_FFLAGS=-g -O2
    set CGO_LDFLAGS=-g -O2
    set PKG_CONFIG=pkg-config
    set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\tmcphill\AppData\Local\Temp\go-build087651266=/tmp/go-build -gno-record-gcc-switches
    
    PS C:\Users\tmcphill\go\src\github.com\therecipe\qt\cmd\qtsetup> dir env:GOPATH
    
    Name                           Value
    ----                           -----
    GOPATH                         C:\Users\tmcphill\go
    ```

### Tried installing therecipe/qt Go package but got errors running qtsetup:

- Cloned the repo:
	```console
	PS C:\Users\tmcphill\go> go get -u -v -tags=no_env github.com/therecipe/qt/cmd/...
	github.com/therecipe/qt (download)
	github.com/gopherjs/gopherjs (download)
	github.com/sirupsen/logrus (download)
	github.com/konsorten/go-windows-terminal-sequences (download)
	Fetching https://golang.org/x/tools/imports?go-get=1
	Parsing meta tags from https://golang.org/x/tools/imports?go-get=1 (status code 200)
	get "golang.org/x/tools/imports": found meta tag get.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/imports?go-get=1
	get "golang.org/x/tools/imports": verifying non-authoritative meta tag
	Fetching https://golang.org/x/tools?go-get=1
	Parsing meta tags from https://golang.org/x/tools?go-get=1 (status code 200)
	golang.org/x/tools (download)
	Fetching https://golang.org/x/tools/go/ast/astutil?go-get=1
	Parsing meta tags from https://golang.org/x/tools/go/ast/astutil?go-get=1 (status code 200)
	get "golang.org/x/tools/go/ast/astutil": found meta tag get.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/go/ast/astutil?go-get=1
	get "golang.org/x/tools/go/ast/astutil": verifying non-authoritative meta tag
	Fetching https://golang.org/x/tools/go/packages?go-get=1
	Parsing meta tags from https://golang.org/x/tools/go/packages?go-get=1 (status code 200)
	get "golang.org/x/tools/go/packages": found meta tag get.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/go/packages?go-get=1
	get "golang.org/x/tools/go/packages": verifying non-authoritative meta tag
	Fetching https://golang.org/x/tools/go/gcexportdata?go-get=1
	Parsing meta tags from https://golang.org/x/tools/go/gcexportdata?go-get=1 (status code 200)
	get "golang.org/x/tools/go/gcexportdata": found meta tag get.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/go/gcexportdata?go-get=1
	get "golang.org/x/tools/go/gcexportdata": verifying non-authoritative meta tag
	Fetching https://golang.org/x/tools/go/internal/gcimporter?go-get=1
	Parsing meta tags from https://golang.org/x/tools/go/internal/gcimporter?go-get=1 (status code 200)
	get "golang.org/x/tools/go/internal/gcimporter": found meta tag get.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/go/internal/gcimporter?go-get=1
	get "golang.org/x/tools/go/internal/gcimporter": verifying non-authoritative meta tag
	Fetching https://golang.org/x/tools/go/internal/packagesdriver?go-get=1
	Parsing meta tags from https://golang.org/x/tools/go/internal/packagesdriver?go-get=1 (status code 200)
	get "golang.org/x/tools/go/internal/packagesdriver": found meta tag get.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/go/internal/packagesdriver?go-get=1
	get "golang.org/x/tools/go/internal/packagesdriver": verifying non-authoritative meta tag
	Fetching https://golang.org/x/tools/internal/gopathwalk?go-get=1
	Parsing meta tags from https://golang.org/x/tools/internal/gopathwalk?go-get=1 (status code 200)
	get "golang.org/x/tools/internal/gopathwalk": found meta tag get.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/internal/gopathwalk?go-get=1
	get "golang.org/x/tools/internal/gopathwalk": verifying non-authoritative meta tag
	Fetching https://golang.org/x/tools/internal/fastwalk?go-get=1
	Parsing meta tags from https://golang.org/x/tools/internal/fastwalk?go-get=1 (status code 200)
	get "golang.org/x/tools/internal/fastwalk": found meta tag get.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/internal/fastwalk?go-get=1
	get "golang.org/x/tools/internal/fastwalk": verifying non-authoritative meta tag
	Fetching https://golang.org/x/tools/internal/semver?go-get=1
	Parsing meta tags from https://golang.org/x/tools/internal/semver?go-get=1 (status code 200)
	get "golang.org/x/tools/internal/semver": found meta tag get.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/internal/semver?go-get=1
	get "golang.org/x/tools/internal/semver": verifying non-authoritative meta tag
	Fetching https://golang.org/x/tools/internal/module?go-get=1
	Parsing meta tags from https://golang.org/x/tools/internal/module?go-get=1 (status code 200)
	get "golang.org/x/tools/internal/module": found meta tag get.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/internal/module?go-get=1
	get "golang.org/x/tools/internal/module": verifying non-authoritative meta tag
	Fetching https://golang.org/x/crypto/ssh?go-get=1
	Parsing meta tags from https://golang.org/x/crypto/ssh?go-get=1 (status code 200)
	get "golang.org/x/crypto/ssh": found meta tag get.metaImport{Prefix:"golang.org/x/crypto", VCS:"git", RepoRoot:"https://go.googlesource.com/crypto"} at https://golang.org/x/crypto/ssh?go-get=1
	get "golang.org/x/crypto/ssh": verifying non-authoritative meta tag
	Fetching https://golang.org/x/crypto?go-get=1
	Parsing meta tags from https://golang.org/x/crypto?go-get=1 (status code 200)
	golang.org/x/crypto (download)
	Fetching https://golang.org/x/crypto/curve25519?go-get=1
	Parsing meta tags from https://golang.org/x/crypto/curve25519?go-get=1 (status code 200)
	get "golang.org/x/crypto/curve25519": found meta tag get.metaImport{Prefix:"golang.org/x/crypto", VCS:"git", RepoRoot:"https://go.googlesource.com/crypto"} at https://golang.org/x/crypto/curve25519?go-get=1
	get "golang.org/x/crypto/curve25519": verifying non-authoritative meta tag
	Fetching https://golang.org/x/crypto/ed25519?go-get=1
	Parsing meta tags from https://golang.org/x/crypto/ed25519?go-get=1 (status code 200)
	get "golang.org/x/crypto/ed25519": found meta tag get.metaImport{Prefix:"golang.org/x/crypto", VCS:"git", RepoRoot:"https://go.googlesource.com/crypto"} at https://golang.org/x/crypto/ed25519?go-get=1
	get "golang.org/x/crypto/ed25519": verifying non-authoritative meta tag
	Fetching https://golang.org/x/crypto/ed25519/internal/edwards25519?go-get=1
	Parsing meta tags from https://golang.org/x/crypto/ed25519/internal/edwards25519?go-get=1 (status code 200)
	get "golang.org/x/crypto/ed25519/internal/edwards25519": found meta tag get.metaImport{Prefix:"golang.org/x/crypto", VCS:"git", RepoRoot:"https://go.googlesource.com/crypto"} at https://golang.org/x/crypto/ed25519/internal/edwards25519?go-get=1
	get "golang.org/x/crypto/ed25519/internal/edwards25519": verifying non-authoritative meta tag
	Fetching https://golang.org/x/crypto/internal/chacha20?go-get=1
	Parsing meta tags from https://golang.org/x/crypto/internal/chacha20?go-get=1 (status code 200)
	get "golang.org/x/crypto/internal/chacha20": found meta tag get.metaImport{Prefix:"golang.org/x/crypto", VCS:"git", RepoRoot:"https://go.googlesource.com/crypto"} at https://golang.org/x/crypto/internal/chacha20?go-get=1
	get "golang.org/x/crypto/internal/chacha20": verifying non-authoritative meta tag
	Fetching https://golang.org/x/crypto/internal/subtle?go-get=1
	Parsing meta tags from https://golang.org/x/crypto/internal/subtle?go-get=1 (status code 200)
	get "golang.org/x/crypto/internal/subtle": found meta tag get.metaImport{Prefix:"golang.org/x/crypto", VCS:"git", RepoRoot:"https://go.googlesource.com/crypto"} at https://golang.org/x/crypto/internal/subtle?go-get=1
	get "golang.org/x/crypto/internal/subtle": verifying non-authoritative meta tag
	Fetching https://golang.org/x/crypto/poly1305?go-get=1
	Parsing meta tags from https://golang.org/x/crypto/poly1305?go-get=1 (status code 200)
	get "golang.org/x/crypto/poly1305": found meta tag get.metaImport{Prefix:"golang.org/x/crypto", VCS:"git", RepoRoot:"https://go.googlesource.com/crypto"} at https://golang.org/x/crypto/poly1305?go-get=1
	get "golang.org/x/crypto/poly1305": verifying non-authoritative meta tag
	github.com/gopherjs/gopherjs/js
	github.com/therecipe/qt/internal/binding/files/docs/5.10.0
	github.com/therecipe/qt/internal/binding/files/docs/5.11.1
	github.com/therecipe/qt/internal/binding/files/docs/5.12.0
	github.com/therecipe/qt/internal/binding/files/docs/5.6.3
	github.com/therecipe/qt/internal/binding/files/docs/5.7.0
	github.com/therecipe/qt/internal/binding/files/docs/5.7.1
	github.com/therecipe/qt/internal/binding/files/docs/5.8.0
	github.com/therecipe/qt/internal/binding/files/docs/5.9.0
	golang.org/x/tools/internal/semver
	github.com/konsorten/go-windows-terminal-sequences
	golang.org/x/tools/internal/fastwalk
	github.com/therecipe/qt/internal/binding/files/docs
	golang.org/x/tools/go/ast/astutil
	golang.org/x/tools/go/internal/packagesdriver
	golang.org/x/tools/go/internal/gcimporter
	golang.org/x/crypto/curve25519
	golang.org/x/tools/internal/module
	github.com/sirupsen/logrus
	golang.org/x/crypto/ed25519/internal/edwards25519
	golang.org/x/tools/internal/gopathwalk
	golang.org/x/crypto/internal/subtle
	golang.org/x/crypto/poly1305
	github.com/therecipe/qt/internal/utils
	golang.org/x/crypto/ed25519
	golang.org/x/crypto/internal/chacha20
	golang.org/x/tools/go/gcexportdata
	golang.org/x/tools/go/packages
	github.com/therecipe/qt/internal/cmd
	github.com/therecipe/qt/internal/binding/parser
	golang.org/x/crypto/ssh
	golang.org/x/tools/imports
	github.com/therecipe/qt/internal/binding/converter
	github.com/therecipe/qt/internal/binding/templater
	github.com/therecipe/qt/internal/cmd/rcc
	github.com/therecipe/qt/internal/cmd/moc
	github.com/therecipe/qt/internal/cmd/minimal
	github.com/therecipe/qt/cmd/qtrcc
	github.com/therecipe/qt/cmd/qtmoc
	github.com/therecipe/qt/cmd/qtminimal
	github.com/therecipe/qt/internal/cmd/deploy
	github.com/therecipe/qt/cmd/qtdeploy
	github.com/therecipe/qt/internal/cmd/setup
	github.com/therecipe/qt/cmd/qtsetup
	```
- Found new binaries that were installed in the above process:
	```console
	PS C:\Users\tmcphill\go\bin> dir qt*

	    Directory: C:\Users\tmcphill\go\bin

	Mode                LastWriteTime         Length Name
	----                -------------         ------ ----
	-a----        4/29/2019   2:53 PM       12652544 qtdeploy.exe
	-a----        4/29/2019   2:53 PM        5692928 qtminimal.exe
	-a----        4/29/2019   2:53 PM       10128896 qtmoc.exe
	-a----        4/29/2019   2:53 PM        4120064 qtrcc.exe
	-a----        4/29/2019   2:53 PM       12767744 qtsetup.exe
	```
- Running qtsetup.exe via PowerShell failed because g++ cannot be found:
	```console
	PS C:\Users\tmcphill\go\bin> .\qtsetup.exe
	time="2019-04-29T14:58:17-07:00" level=error msg="failed to load C:\\Qt\\5.12.0\\mingw49_32\\bin\\qtenv2.bat" error="open C:\\Qt\\5.12.0\\mingw49_32\\bin\\qtenv2.bat: The system cannot find the path specified."
	time="2019-04-29T14:58:17-07:00" level=warning msg="failed to create qtenv.bat symlink in your PATH (C:\\Go\\bin\\qtenv.bat); please use C:\\Qt\\5.12.0\\mingw49_32\\bin\\qtenv2.bat instead"
	time="2019-04-29T14:58:17-07:00" level=info msg="running: 'qtsetup prep'"
	time="2019-04-29T14:58:17-07:00" level=info msg="successfully created qtrcc symlink in your PATH (C:\\Go\\bin\\qtrcc.exe)"
	time="2019-04-29T14:58:17-07:00" level=info msg="successfully created qtmoc symlink in your PATH (C:\\Go\\bin\\qtmoc.exe)"
	time="2019-04-29T14:58:17-07:00" level=info msg="successfully created qtminimal symlink in your PATH (C:\\Go\\bin\\qtminimal.exe)"
	time="2019-04-29T14:58:17-07:00" level=info msg="successfully created qtdeploy symlink in your PATH (C:\\Go\\bin\\qtdeploy.exe)"
	time="2019-04-29T14:58:17-07:00" level=info msg="running: 'qtsetup check windows' [docker=false] [vagrant=false]"
	time="2019-04-29T14:58:18-07:00" level=panic msg="failed to find g++, did you start the MinGW shell?" error="exec: \"g++\": executable file not found in %PATH%"
	panic: (*logrus.Entry) (0x935a20,0xc0002367e0)

	goroutine 1 [running]:
	github.com/sirupsen/logrus.Entry.log(0xc000048120, 0xc00034f860, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, ...)
	        C:/Users/tmcphill/go/src/github.com/sirupsen/logrus/entry.go:239 +0x2d5
	github.com/sirupsen/logrus.(*Entry).Log(0xc000236770, 0x0, 0xc000079a38, 0x1, 0x1)
	        C:/Users/tmcphill/go/src/github.com/sirupsen/logrus/entry.go:268 +0xf4
	github.com/sirupsen/logrus.(*Entry).Panic(0xc000236770, 0xc000079a38, 0x1, 0x1)
	        C:/Users/tmcphill/go/src/github.com/sirupsen/logrus/entry.go:306 +0x5c
	github.com/therecipe/qt/internal/cmd/setup.Check(0x940d3c, 0x7, 0x930000)
	        C:/Users/tmcphill/go/src/github.com/therecipe/qt/internal/cmd/setup/check.go:110 +0x2c2e
	main.main()
	        C:/Users/tmcphill/go/src/github.com/therecipe/qt/cmd/qtsetup/main.go:108 +0x3c8
	```

- Ran qtsetup.exe again from the MinGW prompt and still got what appear to be error messages:

    ```console
    C:\Qt\Qt5.12.3\5.12.3\mingw73_64>%GOPATH%\bin\qtsetup.exe
    time="2019-04-29T15:01:27-07:00" level=error msg="failed to load C:\\Qt\\5.12.0\\mingw49_32\\bin\\qtenv2.bat" error="open C:\\Qt\\5.12.0\\mingw49_32\\bin\\qtenv2.bat: The system cannot find the path specified."
    time="2019-04-29T15:01:27-07:00" level=warning msg="failed to create qtenv.bat symlink in your PATH (C:\\Go\\bin\\qtenv.bat); please use C:\\Qt\\5.12.0\\mingw49_32\\bin\\qtenv2.bat instead"
    time="2019-04-29T15:01:27-07:00" level=info msg="running: 'qtsetup prep'"
    time="2019-04-29T15:01:27-07:00" level=info msg="successfully created qtrcc symlink in your PATH (C:\\Go\\bin\\qtrcc.exe)"
    time="2019-04-29T15:01:27-07:00" level=info msg="successfully created qtmoc symlink in your PATH (C:\\Go\\bin\\qtmoc.exe)"
    time="2019-04-29T15:01:27-07:00" level=info msg="successfully created qtminimal symlink in your PATH (C:\\Go\\bin\\qtminimal.exe)"
    time="2019-04-29T15:01:27-07:00" level=info msg="successfully created qtdeploy symlink in your PATH (C:\\Go\\bin\\qtdeploy.exe)"
    time="2019-04-29T15:01:27-07:00" level=info msg="running: 'qtsetup check windows' [docker=false] [vagrant=false]"
    time="2019-04-29T15:01:28-07:00" level=info msg="GOOS:                        'windows'"
    time="2019-04-29T15:01:28-07:00" level=info msg="GOARCH:                      'amd64'"
    time="2019-04-29T15:01:28-07:00" level=info msg="GOVERSION:                   'go1.12.4'"
    time="2019-04-29T15:01:28-07:00" level=info msg="GOROOT:                   *  'C:\\Go\\'"
    time="2019-04-29T15:01:28-07:00" level=info msg="GOPATH:                   *  'C:\\Users\\tmcphill\\go'"
    time="2019-04-29T15:01:28-07:00" level=info msg="GOBIN:                       'C:\\Users\\tmcphill\\go\\bin'"
    time="2019-04-29T15:01:28-07:00" level=info msg="GOMOD:                       ''"
    time="2019-04-29T15:01:28-07:00" level=info msg="QT_HASH:                     '746779e4c012d961bc4fad16cfe6228b1b112215'"
    time="2019-04-29T15:01:28-07:00" level=info msg="QT_API:                      ''"
    time="2019-04-29T15:01:28-07:00" level=info msg="QT_VERSION:                  '5.12.0'"
    time="2019-04-29T15:01:28-07:00" level=info msg="QT_DIR:                      'C:\\Qt'"
    time="2019-04-29T15:01:28-07:00" level=info msg="QT_STUB:                     'false'"
    time="2019-04-29T15:01:28-07:00" level=info msg="QT_DEBUG:                    'false'"
    time="2019-04-29T15:01:28-07:00" level=info msg="QT_QMAKE_DIR:                ''"
    time="2019-04-29T15:01:28-07:00" level=info msg="QT_WEBKIT:                   'false'"
    time="2019-04-29T15:01:28-07:00" level=info msg="QT_DEBUG_CONSOLE:            'false'"
    time="2019-04-29T15:01:28-07:00" level=info msg="QT_MSYS2:                    'false'"
    time="2019-04-29T15:01:28-07:00" level=info msg="running: 'qtsetup generate windows' [docker=false] [vagrant=false]"
    time="2019-04-29T15:01:31-07:00" level=error msg="failed to run command" _func=RunCmd cmd="C:\\Qt\\5.12.0\\mingw49_32\\bin\\qmake -query QT_INSTALL_PREFIX" dir= env= error="exec: \"C:\\\\Qt\\\\5.12.0\\\\mingw49_32\\\\bin\\\\qmake\": file does not exist" name="query install prefix path for windows on windows" 
    ```

### Enabled MinGW executables in PowerShell

- Checked %PATH% variable in default CMD and PS:

	```console
	C:\Users\tmcphill>echo %PATH%
	C:\Program Files (x86)\Common Files\Oracle\Java\javapath;C:\Program Files\Docker\Docker\Resources\bin;C:\Program Files\Python36\Scripts\;C:\Program Files\Python36\;C:\Program Files\Microsoft MPI\Bin\;C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\WINDOWS\System32\OpenSSH\;C:\Program Files\dotnet\;C:\Program Files\Microsoft SQL Server\130\Tools\Binn\;C:\Program Files (x86)\Google\Google Apps Sync\;C:\Program Files (x86)\Google\Google Apps Migration\;C:\Program Files (x86)\NVIDIA Corporation\PhysX\Common;C:\Program Files\NVIDIA Corporation\NVIDIA NvDLISR;C:\ProgramData\chocolatey\bin;C:\HashiCorp\Vagrant\bin;C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\WINDOWS\System32\OpenSSH\;C:\Program Files\MiKTeX 2.9\miktex\bin\x64\;C:\Program Files\Microsoft SQL Server\Client SDK\ODBC\170\Tools\Binn\;C:\Program Files (x86)\IncrediBuild;C:\Program Files\PowerShell\6\;C:\Program Files\Git\cmd;C:\Program Files\Microsoft VS Code\bin;C:\Go\bin;C:\Users\tmcphill\AppData\Local\Microsoft\WindowsApps;C:\Program Files\Microsoft VS Code\bin;C:\Users\tmcphill\AppData\Local\GitHubDesktop\bin;C:\Users\tmcphill\AppData\Local\Microsoft\WindowsApps;;C:\Users\tmcphill\AppData\Local\Programs\Microsoft VS Code\bin;C:\Users\tmcphill\go\bin
	```
	
- Compared against %PATH% in MinGW terminal:

    ```console
    C:\Qt\Qt5.12.3\5.12.3\mingw73_64\bin>echo %PATH%
    C:\Qt\Qt5.12.3\5.12.3\mingw73_64\bin;C:/Qt/Qt5.12.3/Tools/mingw730_64\bin;C:\Program Files (x86)\Common Files\Oracle\Java\javapath;C:\Program Files\Docker\Docker\Resources\bin;C:\Program Files\Python36\Scripts\;C:\Program Files\Python36\;C:\Program Files\Microsoft MPI\Bin\;C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\WINDOWS\System32\OpenSSH\;C:\Program Files\dotnet\;C:\Program Files\Microsoft SQL Server\130\Tools\Binn\;C:\Program Files (x86)\Google\Google Apps Sync\;C:\Program Files (x86)\Google\Google Apps Migration\;C:\Program Files (x86)\NVIDIA Corporation\PhysX\Common;C:\Program Files\NVIDIA Corporation\NVIDIA NvDLISR;C:\ProgramData\chocolatey\bin;C:\HashiCorp\Vagrant\bin;C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\WINDOWS\System32\OpenSSH\;C:\Program Files\MiKTeX 2.9\miktex\bin\x64\;C:\Program Files\Microsoft SQL Server\Client SDK\ODBC\170\Tools\Binn\;C:\Program Files (x86)\IncrediBuild;C:\Program Files\PowerShell\6\;C:\Program Files\Git\cmd;C:\Program Files\Microsoft VS Code\bin;C:\Go\bin;C:\Users\tmcphill\AppData\Local\Microsoft\WindowsApps;C:\Program Files\Microsoft VS Code\bin;C:\Users\tmcphill\AppData\Local\GitHubDesktop\bin;C:\Users\tmcphill\AppData\Local\Microsoft\WindowsApps;;C:\Users\tmcphill\AppData\Local\Programs\Microsoft VS Code\bin;C:\Users\tmcphill\go\bin
    ```
- The difference is two new path elements:

    ```
    C:\Qt\Qt5.12.3\5.12.3\mingw73_64\bin
    C:/Qt/Qt5.12.3/Tools/mingw730_64\bin
    ```
- Added the two paths to the system-wide PATH variable and confirmed the executables in that directory are now accessible from a new instance of PowerShell:

    ```console
    PS C:\Users\tmcphill> qtpaths.exe --version
    qtpaths 1.0

	PS C:\Users\tmcphill> g++
	g++.exe: fatal error: no input files
	compilation terminated.
	```
	
### Successfully ran qtsetup to build and install Qt bindings to Go

- Noticed that error message running qtsetup above refers to an installation of Qt with a different version and location on disk than what I actually have.  
	- Error message refers to: `C:\Qt\5.12.0\`
	- Qt is installed at: `C:\Qt\Qt5.12.3\5.12.3`
- Found a suggested solution in a GitHub issue:
   https://github.com/therecipe/qt/issues/492#issuecomment-355209898
- Following the guidance in the comment to the GitHub issue above, set two system-wide environment variables in System Properties:
    ```
    QT_DIR = C:\Qt\Qt5.12.3\
    QT_VERSION = 5.12.3
    ```

- In new PowerShell instance ran qtsetup again, and this time the various qt components and bindings to Go were compiled and installed successfully:

    ```console
    PS C:\Users\tmcphill\go\bin> .\qtsetup.exe
    time="2019-04-29T18:37:37-07:00" level=info msg="successfully created qtenv.bat symlink in your PATH (C:\\Go\\bin\\qtenv.bat)"
    time="2019-04-29T18:37:37-07:00" level=info msg="running: 'qtsetup prep'"
    time="2019-04-29T18:37:37-07:00" level=info msg="successfully created qtrcc symlink in your PATH (C:\\Go\\bin\\qtrcc.exe)"
    time="2019-04-29T18:37:37-07:00" level=info msg="successfully created qtmoc symlink in your PATH (C:\\Go\\bin\\qtmoc.exe)"
    time="2019-04-29T18:37:37-07:00" level=info msg="successfully created qtminimal symlink in your PATH (C:\\Go\\bin\\qtminimal.exe)"
    time="2019-04-29T18:37:37-07:00" level=info msg="successfully created qtdeploy symlink in your PATH (C:\\Go\\bin\\qtdeploy.exe)"
    time="2019-04-29T18:37:37-07:00" level=info msg="running: 'qtsetup check windows' [docker=false] [vagrant=false]"
    time="2019-04-29T18:37:39-07:00" level=info msg="GOOS:                        'windows'"
    time="2019-04-29T18:37:39-07:00" level=info msg="GOARCH:                      'amd64'"
    time="2019-04-29T18:37:39-07:00" level=info msg="GOVERSION:                   'go1.12.4'"
    time="2019-04-29T18:37:39-07:00" level=info msg="GOROOT:                   *  'C:\\Go\\'"
    time="2019-04-29T18:37:39-07:00" level=info msg="GOPATH:                   *  'C:\\Users\\tmcphill\\go'"
    time="2019-04-29T18:37:39-07:00" level=info msg="GOBIN:                       'C:\\Users\\tmcphill\\go\\bin'"
    time="2019-04-29T18:37:39-07:00" level=info msg="GOMOD:                       ''"
    time="2019-04-29T18:37:39-07:00" level=info msg="QT_HASH:                     '746779e4c012d961bc4fad16cfe6228b1b112215'"
    time="2019-04-29T18:37:39-07:00" level=info msg="QT_API:                      ''"
    time="2019-04-29T18:37:39-07:00" level=info msg="QT_VERSION:               *  '5.12.3'"
    time="2019-04-29T18:37:39-07:00" level=info msg="QT_DIR:                   *  'C:\\Qt\\Qt5.12.3'"
    time="2019-04-29T18:37:39-07:00" level=info msg="QT_STUB:                     'false'"
    time="2019-04-29T18:37:39-07:00" level=info msg="QT_DEBUG:                    'false'"
    time="2019-04-29T18:37:39-07:00" level=info msg="QT_QMAKE_DIR:                ''"
    time="2019-04-29T18:37:39-07:00" level=info msg="QT_WEBKIT:                   'false'"
    time="2019-04-29T18:37:39-07:00" level=info msg="QT_DEBUG_CONSOLE:            'false'"
    time="2019-04-29T18:37:39-07:00" level=info msg="QT_MSYS2:                    'false'"
    time="2019-04-29T18:37:39-07:00" level=info msg="running: 'qtsetup generate windows' [docker=false] [vagrant=false]"
    time="2019-04-29T18:37:40-07:00" level=warning msg=parser.LoadModule error=EOF module=VirtualKeyboard
    time="2019-04-29T18:37:40-07:00" level=warning msg=parser.LoadModule error=EOF module=ScriptTools
    time="2019-04-29T18:37:40-07:00" level=warning msg=parser.LoadModule error=EOF module=Script
    time="2019-04-29T18:37:40-07:00" level=warning msg=parser.LoadModule error=EOF module=Purchasing
    time="2019-04-29T18:37:44-07:00" level=info msg="generating full qt/core"
    time="2019-04-29T18:37:58-07:00" level=info msg="generating full qt/androidextras"
    time="2019-04-29T18:37:59-07:00" level=info msg="generating full qt/gui"
    time="2019-04-29T18:38:19-07:00" level=info msg="generating full qt/network"
    time="2019-04-29T18:38:26-07:00" level=info msg="generating full qt/xml"
    time="2019-04-29T18:38:29-07:00" level=info msg="generating full qt/dbus"
    time="2019-04-29T18:38:32-07:00" level=info msg="generating full qt/nfc"
    time="2019-04-29T18:38:35-07:00" level=info msg="generating full qt/script"
    time="2019-04-29T18:38:36-07:00" level=info msg="generating full qt/sensors"
    time="2019-04-29T18:38:43-07:00" level=info msg="generating full qt/positioning"
    time="2019-04-29T18:38:46-07:00" level=info msg="generating full qt/widgets"
    time="2019-04-29T18:39:36-07:00" level=info msg="generating full qt/sql"
    time="2019-04-29T18:39:40-07:00" level=info msg="generating full qt/qml"
    time="2019-04-29T18:39:45-07:00" level=info msg="generating full qt/websockets"
    time="2019-04-29T18:39:47-07:00" level=info msg="generating full qt/xmlpatterns"
    time="2019-04-29T18:39:49-07:00" level=info msg="generating full qt/bluetooth"
    time="2019-04-29T18:39:54-07:00" level=info msg="generating full qt/webchannel"
    time="2019-04-29T18:39:56-07:00" level=info msg="generating full qt/svg"
    time="2019-04-29T18:39:58-07:00" level=info msg="generating full qt/multimedia"
    time="2019-04-29T18:40:16-07:00" level=info msg="generating full qt/quick"
    time="2019-04-29T18:40:23-07:00" level=info msg="generating full qt/help"
    time="2019-04-29T18:40:29-07:00" level=info msg="generating full qt/location"
    time="2019-04-29T18:40:32-07:00" level=info msg="generating full qt/scripttools"
    time="2019-04-29T18:40:34-07:00" level=info msg="generating full qt/uitools"
    time="2019-04-29T18:40:36-07:00" level=info msg="generating full qt/winextras"
    time="2019-04-29T18:40:39-07:00" level=info msg="generating full qt/testlib"
    time="2019-04-29T18:40:41-07:00" level=info msg="generating full qt/serialport"
    time="2019-04-29T18:40:52-07:00" level=info msg="generating full qt/serialbus"
    time="2019-04-29T18:40:55-07:00" level=info msg="generating full qt/printsupport"
    time="2019-04-29T18:41:00-07:00" level=info msg="generating full qt/designer"
    time="2019-04-29T18:41:07-07:00" level=info msg="generating full qt/scxml"
    time="2019-04-29T18:41:11-07:00" level=info msg="generating full qt/gamepad"
    time="2019-04-29T18:41:14-07:00" level=info msg="generating full qt/purchasing"
    time="2019-04-29T18:41:15-07:00" level=info msg="generating full qt/datavisualization    [GPLv3]"
    time="2019-04-29T18:41:17-07:00" level=info msg="generating full qt/charts               [GPLv3]"
    time="2019-04-29T18:41:18-07:00" level=info msg="generating full qt/virtualkeyboard      [GPLv3]"
    time="2019-04-29T18:41:20-07:00" level=info msg="generating full qt/speech"
    time="2019-04-29T18:41:22-07:00" level=info msg="generating full qt/quickcontrols2"
    time="2019-04-29T18:41:24-07:00" level=info msg="generating full qt/sailfish"
    time="2019-04-29T18:41:24-07:00" level=info msg="generating full qt/remoteobjects"
    time="2019-04-29T18:41:27-07:00" level=info msg="running: 'qtsetup install windows' [docker=false] [vagrant=false]"
    time="2019-04-29T18:41:28-07:00" level=info msg="installing full qt/core"
    time="2019-04-29T18:44:26-07:00" level=info msg="installing full qt/androidextras"
    time="2019-04-29T18:44:28-07:00" level=info msg="installing full qt/gui"
    time="2019-04-29T18:49:52-07:00" level=info msg="installing full qt/network"
    time="2019-04-29T18:50:45-07:00" level=info msg="installing full qt/xml"
    time="2019-04-29T18:51:03-07:00" level=info msg="installing full qt/dbus"
    time="2019-04-29T18:51:21-07:00" level=info msg="installing full qt/nfc"
    time="2019-04-29T18:51:38-07:00" level=info msg="installing full qt/script"
    time="2019-04-29T18:51:47-07:00" level=info msg="installing full qt/sensors"
    time="2019-04-29T18:52:16-07:00" level=info msg="installing full qt/positioning"
    time="2019-04-29T18:52:35-07:00" level=info msg="installing full qt/widgets"
    time="2019-04-29T19:02:44-07:00" level=info msg="installing full qt/sql"
    time="2019-04-29T19:03:12-07:00" level=info msg="installing full qt/qml"
    time="2019-04-29T19:03:35-07:00" level=info msg="installing full qt/websockets"
    time="2019-04-29T19:03:52-07:00" level=info msg="installing full qt/xmlpatterns"
    time="2019-04-29T19:04:09-07:00" level=info msg="installing full qt/bluetooth"
    time="2019-04-29T19:04:37-07:00" level=info msg="installing full qt/webchannel"
    time="2019-04-29T19:04:52-07:00" level=info msg="installing full qt/svg"
    time="2019-04-29T19:05:11-07:00" level=info msg="installing full qt/multimedia"
    time="2019-04-29T19:07:15-07:00" level=info msg="installing full qt/quick"
    time="2019-04-29T19:08:00-07:00" level=info msg="installing full qt/help"
    time="2019-04-29T19:08:38-07:00" level=info msg="installing full qt/location"
    time="2019-04-29T19:09:02-07:00" level=info msg="installing full qt/scripttools"
    time="2019-04-29T19:09:15-07:00" level=info msg="installing full qt/uitools"
    time="2019-04-29T19:09:34-07:00" level=info msg="installing full qt/winextras"
    time="2019-04-29T19:09:58-07:00" level=info msg="installing full qt/testlib"
    time="2019-04-29T19:10:17-07:00" level=info msg="installing full qt/serialport"
    time="2019-04-29T19:10:37-07:00" level=info msg="installing full qt/serialbus"
    time="2019-04-29T19:11:02-07:00" level=info msg="installing full qt/printsupport"
    time="2019-04-29T19:11:32-07:00" level=info msg="installing full qt/designer"
    time="2019-04-29T19:12:10-07:00" level=info msg="installing full qt/scxml"
    time="2019-04-29T19:12:34-07:00" level=info msg="installing full qt/gamepad"
    time="2019-04-29T19:12:57-07:00" level=info msg="installing full qt/purchasing"
    time="2019-04-29T19:13:12-07:00" level=info msg="installing full qt/datavisualization    [GPLv3]"
    time="2019-04-29T19:13:27-07:00" level=info msg="installing full qt/charts               [GPLv3]"
    time="2019-04-29T19:13:43-07:00" level=info msg="installing full qt/virtualkeyboard      [GPLv3]"
    time="2019-04-29T19:13:58-07:00" level=info msg="installing full qt/speech"
    time="2019-04-29T19:14:22-07:00" level=info msg="installing full qt/quickcontrols2"
    time="2019-04-29T19:14:40-07:00" level=info msg="installing full qt/sailfish"
    time="2019-04-29T19:14:44-07:00" level=info msg="installing full qt/remoteobjects"
    time="2019-04-29T19:15:11-07:00" level=info msg="running: 'qtsetup test windows' [docker=false] [vagrant=false]"
    time="2019-04-29T19:15:12-07:00" level=info msg="testing qml\\application"
    time="2019-04-29T19:17:26-07:00" level=info msg="testing qml\\drawer_nav_x"
    time="2019-04-29T19:20:08-07:00" level=info msg="testing qml\\gallery"
    time="2019-04-29T19:22:10-07:00" level=info msg="testing quick\\calc"
    time="2019-04-29T19:24:05-07:00" level=info msg="testing widgets\\line_edits"
    time="2019-04-29T19:26:07-07:00" level=info msg="testing widgets\\pixel_editor"
    time="2019-04-29T19:28:31-07:00" level=info msg="testing widgets\\textedit"
    ```
- All seven of the test application started at the end of the above build appear to be fully functional.

