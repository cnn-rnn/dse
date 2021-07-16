package proc



import (
"os"
"io/ioutil"
"strconv"
"strings"
"time"
)



const N = 100
const N1 = 1000



type Proc struct{
  cpu []float64
  cput []int64
  cpu_n int
  bw []float64
  bwt []int64
  bw_n int
  disk float64
  
  dir0 string
  cpu0 float64
  bw0 float64
  disk0 float64  
  
  tts * float64
  
  tstart int64
  
  t1 int64
  
}










func Create(dir0  string , cpu float64, bw float64, disk float64, tts * float64)*Proc{
return &Proc{make([]float64,N),make([]int64,N),0,make([]float64,N1),make([]int64,N1),0,0,dir0,cpu,bw,disk,tts, time.Now().UnixNano(), time.Now().UnixNano()}
}


func (x * Proc)Start(){
x.tstart = time.Now().UnixNano()
go x.KeepCpu()
go x.Run()
}



func sum0(z []float64)float64{
var y float64
for i:=0;i<len(z);i++{
  y += float64(z[i]);
}
return y/float64(len(z))
}



func sum(z []float64, t []int64 , m int, n int)float64{
var y float64
for i:=0;i<n;i++{
  y += float64(z[i]);
}
j2 := m - 1
if(j2 <0){
  j2 = n-1
}
j1 := j2 - 1
if(j1 <0){
  j1 = n-1
}
dt := float64( t[j2] - t[j1])*1e-9
return y/dt*1e-6   
}




func (x * Proc)Run(){
for{
  a := sum0(x.cpu)    
  if( time.Now().UnixNano() - x.tstart > 600*1e+9){
    *x.tts +=  0.001*(*x.tts)*(a-x.cpu0)
  }
/*  
  temp := make([]int,N)
  for i:=0;i < N;i++{
     temp[i] = int(x.cpu[i])
  }
  fmt.Println(temp,0.001*(*x.tts)*(a-x.cpu0))
*/  
  time.Sleep(5*time.Second)
}
}
  
  

 
 
func (x * Proc)AddToBw(t int64, y int){
if(x.bw_n ==0){
  x.t1 = time.Now().UnixNano()
}
x.bw[x.bw_n] =float64(y);
x.bwt[x.bw_n] =t - x.tstart;
x.bw_n +=1
x.bw_n %=N1
}


func (x * Proc)SumBw()float64{
var y float64
for i:=0;i<N;i++{
  y += float64(x.bw[i]);
}
j2 := x.bw_n - 1
if(j2 <0){
  j2 = N1-1
}
j1 := j2 - 1
if(j1 <0){
  j1 = N1-1
}
dt := float64( x.bwt[j2] - x.bwt[j1])*1e-9
return y/dt   // in Mb/s
}






func proc(pid int)[]float64{
pids := strconv.Itoa(pid)
r,_ := ioutil.ReadFile("/proc/"+pids+"/stat")
s := string(r)
i := 0
for {
  if(s[i] == ')'){
    break
  }
  i+=1
}
i+=2
for {
  if(s[i] == ' '){
    break
  }
  i+=1
}
q := strings.Split(s," ")
p := make([]float64,0)
for i := 4;i<len(q);i++{
   x := q[i]
   y,_ := strconv.ParseFloat(x,10)
   p = append(p,y)
}
return p
}





func usage(pid int)float64{
p := proc(pid)
y := p[10]+p[9]
return y
}



func (x * Proc) KeepCpu(){
pid := os.Getpid()
u1 := usage(pid)
for{
  u := usage(pid)  
  v := (u-u1)/3.0
  x.cpu[x.cpu_n] = v
  x.cput[x.cpu_n] = time.Now().UnixNano()
  x.cpu_n +=1
  x.cpu_n %= N
  time.Sleep(3*time.Second)
  u1 = u
}
}



  











