package fmonitor

import (
        "fmt"
        "os"
        "path/filepath"
        "runtime/debug"
        "time"
)

const (
        MAX_SLEEP_TIME = 100 /* ms */
)

type Op int

const (
        OP_ADD  Op      = 0x01
        OP_DEL  Op      = 0x02
        OP_MOD  Op      = 0x04
)

type FileOp struct {
        op      Op
        fname   string
}

type FileMonitor struct {
        fpath   string
        f       chan FileOp
}

func (fo FileOp) GetOpValue() Op {
        return fo.op
}

func (fo FileOp) GetNameValue() string {
        return fo.fname
}

func NewFileMonitor(name string) *FileMonitor {
        fm := new(FileMonitor)
        fm.fpath = name
        fm.f = make(chan FileOp)

        return fm
}

func (fm *FileMonitor) Release() {
        close(fm.f)
}

func (fm *FileMonitor) GetFileOp() FileOp {
        fo := <-fm.f
        return fo
}
func getDirInfo(fpath string, fin map[string]os.FileInfo) {
        for {
                if fi, err := os.Stat(fpath); err != nil {
                        panic(err)
                } else {
                        if fi.IsDir() {
                                if f, err := os.Open(fpath); err != nil {
                                        panic(err)
                                } else {
                                        v, err := f.Readdirnames(0)
                                        if err != nil {
                                                panic(err)
                                        }
                                        f.Close()
                                        for _, v0 := range v {
                                                v0 = filepath.Join(fpath, v0)
                                                getDirInfo(v0, fin)
                                        }
                                        break
                                }

                        } else {
                                fin[fpath] = fi
                                break
                        }
                }
        }
}

func (cmc *FileMonitor) cmpStat(f_old map[string]os.FileInfo, f_new map[string]os.FileInfo) {
        var fo FileOp
        for i, v := range f_old {
                if in, ok := f_new[i]; ok {
                        if !v.ModTime().Equal(in.ModTime()) {
                                // update
                                f_old[i] = in
                                delete(f_new, i)
                                fo.fname = i
                                fo.op = OP_MOD
                                cmc.f <- fo
                        }
                        delete(f_new, i)
                        // no changed
                } else {
                        delete(f_old, i)
                        fo.fname = i
                        fo.op = OP_DEL
                        cmc.f <- fo
                        // delete
                }
        }

        for i, v := range f_new {
                // add
                f_old[i] = v
                delete(f_new, i)
                fo.fname = i
                fo.op = OP_ADD
                cmc.f <- fo
        }
        debug.FreeOSMemory()
}

func (cmc *FileMonitor) Monitoring() {
        var ffi *os.FileInfo
        var fo FileOp
        if fi, err := os.Stat(cmc.fpath); err != nil {
                panic(err)
        } else {
                if fi.IsDir() {
                        e_new := make(map[string]os.FileInfo)
                        e_old := make(map[string]os.FileInfo)
                        // filepath.Walk(fpath, walkFunc)
                        getDirInfo(cmc.fpath, e_old)
                        for {
                                time.Sleep(time.Millisecond * time.Duration(MAX_SLEEP_TIME))
                                getDirInfo(cmc.fpath, e_new)
                                cmc.cmpStat(e_old, e_new)
                        }
                }
                ffi = &fi
        }

        for {
                time.Sleep(time.Millisecond * time.Duration(MAX_SLEEP_TIME))
                if fi, err := os.Stat(cmc.fpath); err != nil {
                        if os.IsNotExist(err) {
                                fo.fname = cmc.fpath
                                fo.op = OP_DEL
                                cmc.f <- fo
                                return
                        }
                } else {
                        if (*ffi).ModTime().Equal(fi.ModTime()) {
                                continue
                        }
                        fo.fname = cmc.fpath
                        fo.op = OP_MOD
                        cmc.f <- fo
                        ffi = &fi
                        debug.FreeOSMemory()
                }
        }
}

func print_stat(cmc *FileMonitor) {
        for {
                select {
                case s := <-cmc.f:
                        switch s.op {
                        case OP_ADD:
                                fmt.Printf("The file <%s> has been added.\n", s.fname)
                        case OP_DEL:
                                fmt.Printf("The file <%s> has been deleted.\n", s.fname)
                        case OP_MOD:
                                fmt.Printf("The file <%s> has been changed.\n", s.fname)
                        default:
                                fmt.Printf("Unknown operation on the file <%s>.\n", s.fname)
                        }

                }
        }
}
