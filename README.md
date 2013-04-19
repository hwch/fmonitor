<div><div>fmonitor</div><div>========</div><div><br></div><div>The file monitor</div><div><br></div><div>Cross platform, works on:</div><div><br></div><div>&nbsp; Windows</div><div>&nbsp; Linux</div><div>&nbsp; BSD</div><div>&nbsp; OSX</div></div><div><br></div><div><br></div><div>package main</div><div><br></div><div>import (</div><div>&nbsp; &nbsp; &nbsp; &nbsp; "flag"</div><div>&nbsp; &nbsp; &nbsp; &nbsp; "fmonitor"</div><div>&nbsp; &nbsp; &nbsp; &nbsp; "fmt"</div><div>&nbsp; &nbsp; &nbsp; &nbsp; "os"</div><div>&nbsp; &nbsp; &nbsp; &nbsp; "runtime"</div><div>)</div><div><br></div><div>func main() {</div><div>&nbsp; &nbsp; &nbsp; &nbsp; var fpath string</div><div>&nbsp; &nbsp; &nbsp; &nbsp; var help bool</div><div>&nbsp; &nbsp; &nbsp; &nbsp; runtime.GOMAXPROCS(runtime.NumCPU())</div><div>&nbsp; &nbsp; &nbsp; &nbsp; flag.StringVar(&amp;fpath, "path", "/tmp/test.txt", "--path The path to the file you want to monitor")</div><div>&nbsp; &nbsp; &nbsp; &nbsp; flag.BoolVar(&amp;help, "help", false, "-h &nbsp;--help Show the help information")</div><div>&nbsp; &nbsp; &nbsp; &nbsp; flag.BoolVar(&amp;help, "h", false, "")</div><div>&nbsp; &nbsp; &nbsp; &nbsp; flag.Usage = func() {</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; fmt.Fprintf(os.Stderr, "-h &nbsp;--help: Show the help information\n"+</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "--path: --path The path to the file you want to monitor\n")</div><div>&nbsp; &nbsp; &nbsp; &nbsp; }</div><div>&nbsp; &nbsp; &nbsp; &nbsp; flag.Parse()</div><div><br></div><div>&nbsp; &nbsp; &nbsp; &nbsp; if help {</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; flag.Usage()</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; return</div><div>&nbsp; &nbsp; &nbsp; &nbsp; }</div><div>&nbsp; &nbsp; &nbsp; &nbsp; if _, err := os.Stat(fpath); err != nil {</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; if os.IsNotExist(err) {</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; fmt.Printf("The file '%s' is not exist !\n", fpath)</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; return</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; }</div><div>&nbsp; &nbsp; &nbsp; &nbsp; }</div><div>&nbsp; &nbsp; &nbsp; &nbsp; cmc := fmonitor.NewFileMonitor(fpath)</div><div>&nbsp; &nbsp; &nbsp; &nbsp; defer cmc.Release()</div><div>&nbsp; &nbsp; &nbsp; &nbsp; go cmc.Monitoring()</div><div>&nbsp; &nbsp; &nbsp; &nbsp; for {</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; fo := cmc.GetFileOp()</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; switch fo.GetOpValue() {</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; case fmonitor.OP_ADD:</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; fmt.Printf("The file &lt;%s&gt; has been added.\n", fo.GetNameValue())</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; case fmonitor.OP_DEL:</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; fmt.Printf("The file &lt;%s&gt; has been deleted.\n", fo.GetNameValue())</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; case fmonitor.OP_MOD:</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; fmt.Printf("The file &lt;%s&gt; has been changed.\n", fo.GetNameValue())</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; default:</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; fmt.Printf("Unknown operation on the file &lt;%s&gt;.\n", fo.GetNameValue())</div><div>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; }</div><div>&nbsp; &nbsp; &nbsp; &nbsp; }</div><div><br></div><div>}</div>
