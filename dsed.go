package main


//cc

import (
//"fmt"
"log"
//"time"
//"math/rand"
//"os"

"net/http"
_ "net/http/pprof"


"./src/fetcher"
"./src/protocol"
"./src/config"
"./src/ds"
"./src/monitor"
"./src/process"
"./src/security"
"./src/seed"
"./src/logger"
)



var nJobs = 2000
var nResults = 10000
var nTodo = 100
var nWorkers = 100
var ttosleep = 0.0

var tstart int64




var dstore * ds.Ds
var mon * moni.Mon
var cnf * config.Conf
var sec * security.Sec
var pro * proc.Proc


var S = make(map[int]int)


type Resp struct{
  name string
  links map[string]bool
  txt string
  siz int
  t int64
  err string
}



func worker(id int, jobs <-chan string, results chan<- Resp) {
  for j := range jobs {
    links,txt,siz,t, err := fetcher.OnePageLinks(j)
    if( err == ""){
       results <- Resp{j,links,txt,siz,t,err}
        
    }
  }
}



func Process_links(name string, links map[string]bool){
for i:= range links{
  dstore.Store(i)
}
msg := protocol.Serialize(name,links)
mon.SendString(sec,  msg)
}





func main(){
  logger.Log("dsed started")  

  cnf = config.Configure()  
  dstore = ds.Create(cnf.Dir0() )
  sec = security.Create(cnf.Dir0())  
  pro = proc.Create(cnf.Dir0(), cnf.Cpu() , cnf.Bw(), cnf.Disk(),&ttosleep)
  pro.Start()


  mon = moni.Create()
  go mon.Maintain(sec.GetPubKAsString())
  go mon.Listen(cnf.Dir0())





  jobs := make(chan string, nJobs)
  results := make(chan Resp, nResults)

  seed.Seed(jobs, cnf.Dir0())

  for w := 1; w <= nWorkers; w++ {
    go worker(w, jobs, results)
  }


go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}()

  

  for {

    R := <-results
    name := R.name
    lin := R.links
//    erro := R.err
    pro.AddToBw(R.t , R.siz)
    
    
    Process_links( name,lin)    

//    log.Println("N=",N,"results= ",len(results),"jobs=",len(jobs),"\n\n\n")
    
    count := 0
    for len(jobs) < nJobs && count <10{
      count+=1
      u := dstore.Rand()
      if(len(u)>0){
        jobs <- u
      }
    }
    
  }  
}  

  
  
  

  
  




