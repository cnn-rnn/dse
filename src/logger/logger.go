package logger




import(
"time"
"io/ioutil"
"strings"
"os"
)



func Log(s string){
t := time.Now().String()
s = t+ "  "+s + "\n"
f,_ := os.OpenFile("/var/log/dse/dsed.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
f.Write([]byte(s))
f.Close()
}



func get_dir0()string{
r,_ := ioutil.ReadFile("/etc/dse/dsed.conf")
s := string(r)
q := strings.Split(s,"\n")
return q[1]
}


func Log0(s string){
dir0 := get_dir0()
t := time.Now().String()
s = t+ "  "+s + "\n"
f,_ := os.OpenFile(dir0+"/dsed.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
f.Write([]byte(s))
f.Close()
}


func Loge(s string){
dir0 := get_dir0()
s += "\n"
f,_ := os.OpenFile(dir0+"/err.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
f.Write([]byte(s))
f.Close()
}




