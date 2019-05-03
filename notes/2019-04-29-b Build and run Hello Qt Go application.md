## 2019-04-29-b Build and run Hello Qt Go application

### Background
- The Qt bindings to Go appear to work on circe-win10.
- I would like to build and run a sample application in the book Hands-On GUI Application Development in Go.
- The Hello World application source code from the book is as follows:

    ```go
    package main
    
    import (
    	"os"
    
    	"github.com/therecipe/qt/widgets"
    )
    
    func main() {
    
    	app := widgets.NewQApplication(len(os.Args), os.Args)
    
    	window := widgets.NewQMainWindow(nil, 0)
    	window.SetWindowTitle("Hello World")
    
    	widget := widgets.NewQWidget(window, 0)
    	widget.SetLayout(widgets.NewQVBoxLayout())
    	window.SetCentralWidget(widget)
    
    	label := widgets.NewQLabel2("Hello World!", window, 0)
    	widget.Layout().AddWidget(label)
    
    	button := widgets.NewQPushButton2("Quit", window)
    	button.ConnectClicked(func(bool) {
    		app.QuitDefault()
    	})
    	widget.Layout().AddWidget(button)
    
    	window.Show()
    	widgets.QApplication_Exec()
    }
    ```
### Built and ran application
- Saved above source code in `go-demos\04_hello_qt\hello_qt.go`.
- Built the application to yield a 79 MB (!) executable.

    ```console
    PS C:\Users\tmcphill\Go\src\go-demos\04_hello_qt> go build .\hello_qt.go
    PS C:\Users\tmcphill\Go\src\go-demos\04_hello_qt> dir
     
        Directory: C:\Users\tmcphill\Go\src\go-demos\04_hello_qt
  
    Mode                LastWriteTime         Length Name
    ----                -------------         ------ ----
    -a----        4/29/2019  11:59 PM       78053263 hello_qt.exe
    -a----        4/29/2019  11:37 PM            623 hello_qt.go
    ```

- Ran the compiled application:

    ```console
    PS C:\Users\tmcphill\Go\src\go-demos\04_hello_qt> .\hello_qt.exe
    ```
- The program works presenting a tiny window titled "Hello World" and containing two widgets, a label with the text "Hello World!" and a button labeled "Quit".
- Clicking the Quit button causes the app to close but leaves a long stack trace in the PowerShell console:

    ```console
    Exception 0xc0000005 0x1 0x0 0x7ffcb2e9a9e6
    PC=0x7ffcb2e9a9e6
    
    runtime: unknown pc 0x7ffcb2e9a9e6
    stack: frame={sp:0x449e140, fp:0x0} stack=[0x0,0x449fdc0)
    000000000449e040:  00007ffcb2250000  00007ffcb35e9880
    000000000449e050:  00007ffcb2250000  000000000449e110
    000000000449e060:  0000000006b18290  00007ffcb2eaacb1
    000000000449e070:  0000000005330860  000000000449e380
    000000000449e080:  0000c8a8f920c987  0000000000000000
    000000000449e090:  0000000000000000  00007ffcb2eaabda
    000000000449e0a0:  0000000000000000  000000000449e380
    000000000449e0b0:  000000000449e380  0000000000000000
    000000000449e0c0:  00007ffcb2364b60  00007ffcb2e77dc9
    000000000449e0d0:  000000000449e1b0  00007ffcb22af099
    000000000449e0e0:  000074a473b0b434  00007ffcb2364ce0
    000000000449e0f0:  000000000449e1b0  00007ffcb2e89b29
    000000000449e100:  00007ffcb2250000  000000000449e3d0
    000000000449e110:  4f444e49575c3a43  65747379535c5357
    000000000449e120:  0000000000000000  0000006c6c642e46
    000000000449e130:  00007ffcb2364b60  00007ffcb2e9a970
    000000000449e140: <000000000449e1b0  000000000449e7c2
    000000000449e150:  0000000000aee600  00000000001bfd5c
    000000000449e160:  0000000000000000  00007ffcb22943db
    000000000449e170:  00007ffcb22937e0  000000000449e380
    000000000449e180:  00007ffcb2364b60  0000000000000400
    000000000449e190:  00007ffcb22937e0  00007ffcb2294415
    000000000449e1a0:  000000000449e200  0000000000000000
    000000000449e1b0:  0000000000000000  0000000000000004
    000000000449e1c0:  00007ffcb22937e0  00007ffcb229095a
    000000000449e1d0:  0000000000000000  000000000449e3d0
    000000000449e1e0:  0000000000490055 <_cgoexp_f9ac5476e5f3_callbackQAbstractAnimation_Pause+53>  00000000001bf130
    000000000449e1f0:  000000000449e400  00007ffcb22badec
    000000000449e200:  0000000000000001  000000000449e380
    000000000449e210:  00007ffcb2364b60  00007ffc9b1299e7
    000000000449e220:  000000000449e410  00007ffcb22aebe3
    000000000449e230:  0000000000000000  0000000000000000
    runtime: unknown pc 0x7ffcb2e9a9e6
    stack: frame={sp:0x449e140, fp:0x0} stack=[0x0,0x449fdc0)
    000000000449e040:  00007ffcb2250000  00007ffcb35e9880
    000000000449e050:  00007ffcb2250000  000000000449e110
    000000000449e060:  0000000006b18290  00007ffcb2eaacb1
    000000000449e070:  0000000005330860  000000000449e380
    000000000449e080:  0000c8a8f920c987  0000000000000000
    000000000449e090:  0000000000000000  00007ffcb2eaabda
    000000000449e0a0:  0000000000000000  000000000449e380
    000000000449e0b0:  000000000449e380  0000000000000000
    000000000449e0c0:  00007ffcb2364b60  00007ffcb2e77dc9
    000000000449e0d0:  000000000449e1b0  00007ffcb22af099
    000000000449e0e0:  000074a473b0b434  00007ffcb2364ce0
    000000000449e0f0:  000000000449e1b0  00007ffcb2e89b29
    000000000449e100:  00007ffcb2250000  000000000449e3d0
    000000000449e110:  4f444e49575c3a43  65747379535c5357
    000000000449e120:  0000000000000000  0000006c6c642e46
    000000000449e130:  00007ffcb2364b60  00007ffcb2e9a970
    000000000449e140: <000000000449e1b0  000000000449e7c2
    000000000449e150:  0000000000aee600  00000000001bfd5c
    000000000449e160:  0000000000000000  00007ffcb22943db
    000000000449e170:  00007ffcb22937e0  000000000449e380
    000000000449e180:  00007ffcb2364b60  0000000000000400
    000000000449e190:  00007ffcb22937e0  00007ffcb2294415
    000000000449e1a0:  000000000449e200  0000000000000000
    000000000449e1b0:  0000000000000000  0000000000000004
    000000000449e1c0:  00007ffcb22937e0  00007ffcb229095a
    000000000449e1d0:  0000000000000000  000000000449e3d0
    000000000449e1e0:  0000000000490055 <_cgoexp_f9ac5476e5f3_callbackQAbstractAnimation_Pause+53>  00000000001bf130
    000000000449e1f0:  000000000449e400  00007ffcb22badec
    000000000449e200:  0000000000000001  000000000449e380
    000000000449e210:  00007ffcb2364b60  00007ffc9b1299e7
    000000000449e220:  000000000449e410  00007ffcb22aebe3
    000000000449e230:  0000000000000000  0000000000000000
    
    goroutine 1 [running, locked to thread]:
    runtime.asmcgocall(0x456470, 0x1379708)
            C:/Go/src/runtime/asm_amd64.s:620 +0x47
    runtime.stdcall(0x9ede80, 0xc00008a018)
            C:/Go/src/runtime/os_windows.go:800 +0x93
    runtime.stdcall1(0x9ede80, 0x0, 0xc000000001)
            C:/Go/src/runtime/os_windows.go:820 +0x48
    runtime.exit(...)
            C:/Go/src/runtime/os_windows.go:498
    runtime.main()
            C:/Go/src/runtime/proc.go:222 +0x27d
    runtime.goexit()
            C:/Go/src/runtime/asm_amd64.s:1337 +0x1
    rax     0x0
    rbx     0x449e1b0
    rcx     0x74a477f96f840000
    rdi     0x7ffcb2364b60
    rsi     0xffffffffffffffff
    rbp     0x0
    rsp     0x449e140
    r8      0x0
    r9      0x449d750
    r10     0x0
    r11     0x449db80
    r12     0x7ffcb234b528
    r13     0xa3
    r14     0x449e1a8
    r15     0x7ffcb22937e0
    rip     0x7ffcb2e9a9e6
    rflags  0x10286
    cs      0x33
    fs      0x53
    gs      0x2b
    ```

