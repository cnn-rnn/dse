package seed


import(
//"fmt"
//"os"
"strings"
"io/ioutil"
)



func check(e error){
if(e!=nil){
  panic(e)
}
}



func Seed(jobs chan string , dir0 string){
r,e1 := ioutil.ReadFile(dir0+"/seed.txt")
q := strings.Split(string(r),"\n")

if(e1 != nil){

    jobs<-"https://www.amazon.com"
    jobs<-"https://reuters.com"
    jobs<-"https://yahoo.com"
    jobs <-"https://pinterest.com"
    jobs<- "https://baidu.com"
    jobs <-"https://livejournal.com"
    jobs<-"https://cnn.com"
    jobs<-"https://wsj.com"
    jobs<-"https://youtube.com"
  return
}



for i:=range(q){
  if(len(q[i])>0){
    jobs <- q[i]
  }
}

}

