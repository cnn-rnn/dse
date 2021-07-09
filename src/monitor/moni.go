package moni


import(
"net"
//"fmt"
"time"
//"log"
"syscall"
"os"
"bufio"
"../protocol"
"../security"
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
//      log.Println("dse : ec = ",ec)
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
       binary := "/usr/bin/dse"

       args := []string{"dse","update"}

       env := os.Environ()

       syscall.Exec(binary, args, env)
           
       os.Exit(0)
    }
  }
}
}















