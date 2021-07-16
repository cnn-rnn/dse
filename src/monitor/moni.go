package moni


import(
"net"
//"fmt"
"time"
//"log"
//"syscall"
"os"
"os/exec"
"bufio"
"../protocol"
"../security"
"../logger"
)

var monitor_server = "174.138.39.149:7001"


//"localhost:7001"




type Mon struct{
  conn net.Conn
}

func Create()*Mon{
return &Mon{}
}

func (x * Mon)SendString(sec * security.Sec   , msg string){
if(x.conn != nil){
  ziga := sec.SignS([]byte(msg))
  y := msg  + ziga + protocol.SEP
  x.conn.Write([]byte(y+"\n"))
}
}


func (x * Mon)SendLinks(sec * security.Sec, name string, links map[string]bool){
if(x.conn != nil){
  y := protocol.Serialize(name,links)
  s := sec.SignS([]byte(y))
  y += s + protocol.SEP
  x.conn.Write([]byte(y+"\n"))
}
}


func (x * Mon)Maintain(PubK string){
for{
  if(x.conn == nil){
    var ec error
    x.conn ,  ec = net.Dial("tcp", monitor_server)
    if(ec != nil){
    }else{    
       x.conn.Write([]byte(PubK+"\n"))
    }      
  }
  time.Sleep(1*time.Second)
}
}


func (x * Mon)Listen(dir0 string){
for{
  if(x.conn == nil){
     time.Sleep(time.Second)
     continue
  }
  r := bufio.NewReader(x.conn)
  for{
    s,e := r.ReadString('\n')
    if( e != nil){
       x.conn.Close()
       x.conn = nil
       break
    }

    if( string(s) == "update\n"){
       cmd0 := exec.Command("systemd-run","sudo","systemctl","disable","dsed.service")
       err0 := cmd0.Run()
       if err0 != nil {
          logger.Loge("err0"+err0.Error())
       }
       cmd := exec.Command("systemd-run","sudo","dse","update5")
       cmd.Stderr = os.Stderr
       cmd.Stdin = os.Stdin
       cmd.Stdout = os.Stdout
       logger.Log("Running update scripts")
       err := cmd.Run()
       if err != nil {
          logger.Log(err.Error())
       }
       logger.Log("executed "+s+" closing conn")
       x.conn.Close()
       os.Exit(0)
       return
    }

    
    if( string(s) == "reload\n"){
       cmd := exec.Command("systemd-run","sudo","dse","reload")
       cmd.Stderr = os.Stderr
       cmd.Stdin = os.Stdin
       cmd.Stdout = os.Stdout
       logger.Log("Running reload scripts")
       err := cmd.Run()
       if err != nil {
          logger.Log(err.Error())
       }
       x.conn.Close()
       logger.Log("executed "+s+" closing conn")
       os.Exit(0)
       return
    }
    
  
    
    
  }
}
}















