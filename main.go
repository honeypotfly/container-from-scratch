package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

// docker               run image <cmd> <params>
// go run main.go       run       <cmd> <params>


func main() {
    switch os.Args[1] {
    case "run":
        run()
    case "child":
        child()
    default:
        panic ("bad command")
    }
}


// This one is going to create the system, it is going to run *itself*. We need this in order to set the namespace.
func run() {
    fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

    cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
    // Need the below routes so that we can see things happen
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr= os.Stderr

    // Now, we finally set the namespace
    cmd.SysProcAttr = &syscall.SysProcAttr {
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS
        // NEW_UTS is a new unix timesharing system. It just gives a hostname to the container.
        // NEW_PID gives us a namespace for the processes
        // NEWNS gives us a mount namespace so that we dont stay on the root filesystem.
    }

//    cmd.Run()
    if err := cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }

}

// This one is going to create the system, it is going to run *itself*. We need this in order to set the things we define in the namespace.
func child() {
    fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

    // Now, we can set the hostname
    syscall.Sethostname([]byte("container"))

    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    // Need the below routes so that we can see things happen
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr= os.Stderr

    cmd.Run()

}


func must (err error) {
    if err != nil {
        panic(err)
    }
}
