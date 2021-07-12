package logger




import(
"time"
"os"
)



func Log(s string){
t := time.Now().String()
s = t+ "  "+s + "\n"
f,_ := os.OpenFile("/var/log/dse/dsed.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
f.Write([]byte(s))
f.Close()
}
