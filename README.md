fmonitor
========

The file monitor

Cross platform, works on:

Windows
Linux
BSD
OSX

Demos:

package main

import (
        "flag"
        "fmonitor"
        "fmt"
        "os"
        "runtime"
)

func main() {
        var fpath string
        var help bool
        runtime.GOMAXPROCS(runtime.NumCPU())
        flag.StringVar(&fpath, "path", "/tmp/test.txt", "--path The path to the file you want to monitor")
        flag.BoolVar(&help, "help", false, "-h  --help Show the help information")
        flag.BoolVar(&help, "h", false, "")
        flag.Usage = func() {
                fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
                fmt.Fprintf(os.Stderr, "-h  --help: Show the help information\n"+
                        "--path: --path The path to the file you want to monitor\n")
        }
        flag.Parse()

        if help {
                flag.Usage()
                return
        }
        if _, err := os.Stat(fpath); err != nil {
                if os.IsNotExist(err) {
                        fmt.Printf("The file '%s' is not exist !\n", fpath)
                        return
                }
        }
        cmc := fmonitor.NewFileMonitor(fpath)
        defer cmc.Release()
        go cmc.Monitoring()
        for {
                fo := cmc.GetFileOp()
                switch fo.GetOpValue() {
                case fmonitor.OP_ADD:
                        fmt.Printf("The file <%s> has been added.\n", fo.GetNameValue())
                case fmonitor.OP_DEL:
                        fmt.Printf("The file <%s> has been deleted.\n", fo.GetNameValue())
                case fmonitor.OP_MOD:
                        fmt.Printf("The file <%s> has been changed.\n", fo.GetNameValue())
                default:
                        fmt.Printf("Unknown operation on the file <%s>.\n", fo.GetNameValue())
                }
        }

}
