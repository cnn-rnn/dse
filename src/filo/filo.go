package filo

	

import (
"strconv"
"strings"
"log"
"io"
"io/ioutil"
"os"
"time"
"crypto/sha256"
"math/rand"
"encoding/hex"
)


var SEEK_BEG = 0
var SEEK_SET = 0
var SEEK_END = 2
var SEEK_CUR = 1


var LENGTH5 = 10
var LENGTH1 = 4
var LENGTH0 = 64

var LENGTH = 4*LENGTH5 + LENGTH1  + LENGTH0



func GetBytes(f * os.File)[]byte{
r := make([]byte, LENGTH)
i,e := io.ReadAtLeast(f,r,LENGTH)
if( e != nil){
   log.Println("filo.GetBytes i,e =",i,e ,f)
   os.Exit(0)
}
if( i!= LENGTH){
   log.Println("filo.GetBytes i < Length i = ", i)
   os.Exit(0)
}
return r
}



func Convert(s string)string{
y  := sha256.Sum256([]byte(s))
id := hex.EncodeToString(y[:])
return id
}





func Create(fname string, s string){
if(len(s) != LENGTH0){
  log.Println("filo.Create :nonstandrd string",s)
  return
}
f, _ := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0644)
defer f.Close()
f.Truncate(0)
u := make([]byte,LENGTH5)
for i:=0;i<LENGTH5;i++{
  u[i]=' '
}
u[0] = '1'
w := make([]byte,LENGTH5)
for i:=0;i<LENGTH5;i++{
  w[i]=' '
}
w[0] = '-'
z := make([]byte,LENGTH5)
for i:=0;i<LENGTH5;i++{
  z[i]=' '
}
z[0] = '0'
v := make([]byte,LENGTH1)
for i:=0;i<LENGTH1;i++{
  v[i]=' '
}
v[0] = '-'
r := string(u) + string(w) + string(w) + string(z) + string(v) +s
//log.Println("create r=",r)
f.Write([]byte(r))
}



func CreateNewRecord(n int, parent_id string, p int, f * os.File, s string){
//log.Println("parent id = ",parent_id)
c := make([]byte,LENGTH5)
d := strconv.Itoa(n)
for i:=0;i<LENGTH5;i++{
      c[i] = ' '
}
for i:=0;i<len(d);i++{
      c[i] = d[i]
}
f.Seek(int64(p),SEEK_BEG)
f.Write(c)
u := make([]byte,LENGTH5)
for i:=0;i<LENGTH5;i++{
  u[i]=' '
}
u[0] = '-'
v := make([]byte,LENGTH1)
for i:=0;i<LENGTH1;i++{
  v[i]=' '
}
v[0] = '-'
t := parent_id + string(u) + string(u) + string(c) + string(v) + s
//log.Println("t =",t)
f.Seek(0,SEEK_END)
f.Write([]byte(t))
}





func AddString(fname string, s string)bool{
//log.Println("filo:   fname=",fname)
f,e := os.OpenFile(fname,os.O_RDWR,0644)
if(e != nil){
  if(os.IsNotExist(e) ){
     Create(fname, s)
     return true
  }else{
     log.Println("filo.AddString: e= ",e, os.IsNotExist(e) )
     os.Exit(0)
  }
}
defer f.Close()
return AddStringSolid(f,s)
}



func AddStringSolid(f *os.File, s string)bool{
if(len(s) != LENGTH0){
  log.Println("filo.AddStringSolid string nonstandard",s)
  return false
}
b := make([]byte,LENGTH5)
f.Seek(0,SEEK_SET)
f.Read(b)
rho := string(b)
rho = strings.Trim(rho," ")
n,err := strconv.Atoi(rho)
if(err != nil){
   log.Println("filo.AddStringSolid e = ",err," rho = ",rho)
   os.Exit(0)
}
f.Seek(0,SEEK_BEG)
w := Add_string(n,f,s)
if( w ){
   c := make([]byte,LENGTH5)
   for i:=0;i<LENGTH5;i++{
      c[i] = ' '
   }
   u := strconv.Itoa(n+1)
   if(len(u)>=LENGTH5){
      log.Println("LENGTH5 exceeded",u," s= ",s)
      os.Exit(0)
   }
   for i:=0;i<len(u);i++{
      c[i] = u[i]
   }
   f.Seek(0,SEEK_BEG)
   f.Write(c)
}
return w
}



func Add_string(n int, f * os.File, s string)bool{
x := GetBytes(f)
y := strings.Trim(string(x[4*LENGTH5+LENGTH1:len(x)])," ")
l := string(x[LENGTH5:2*LENGTH5])
r := string(x[2*LENGTH5:3*LENGTH5])
t := string(x[3*LENGTH5:4*LENGTH5])
//log.Println(s,y,s<y,s>y, len(s),len(y))
if( y == s){
  return false 
}
if( s < y){
  if( l[0] == '-'){
     p,e := strconv.Atoi(strings.Trim(t," "))
     if(e != nil){
        log.Println("filo.Add_string l : p, e =",p,e,"t=",string(t))
        os.Exit(0)
     }
     p*=LENGTH
     p += LENGTH5
     CreateNewRecord(n,t,p,f,s)
     return true
  }else{
     m,e := strconv.Atoi(strings.Trim(l," "))
     if(e != nil){
        log.Println("filo.Add_string left : m, e =",m,e,"x=",string(x))
        os.Exit(0)
     }
     f.Seek(int64(m*LENGTH) , SEEK_SET)
     return Add_string(n,f,s)
  }
}    
if( s > y){
  if( r[0] == '-'){
     p,e := strconv.Atoi(strings.Trim(t," "))
     if(e != nil){
        log.Println("filo.Add_string r : p, e =",p,e,"t=",string(t))
        os.Exit(0)
     }
     p*=LENGTH
     p += 2*LENGTH5
     CreateNewRecord(n,t,p,f,s)
     return true
  }else{
     m,e := strconv.Atoi(strings.Trim(r," "))
     if(e != nil){
        log.Println("filo.Add_string right : m, e =",m,e, "xxx=",string(x))
        os.Exit(0)
     }
     f.Seek(int64(m*LENGTH) , SEEK_SET)
     return Add_string(n,f,s)
  }
}
return false    
}





func Seek(fname string, s string)*os.File{
f, err := os.OpenFile(fname,os.O_RDWR,0655)
if( err != nil){
  log.Println(err)
  os.Exit(0)
}
x := GetBytes(f)
n0 := strings.Trim(string(x[0:LENGTH5])," ")
n,e := strconv.Atoi(n0)
if( e!= nil){
   log.Println("filo.Seek e=",e)
   os.Exit(0)
}
f.Seek(0,SEEK_BEG)
return seek(n,f,s)
}


func seek(n int , f * os.File, s string)*os.File{
//s = strings.Trim(s," ")
x := GetBytes(f)
y := strings.Trim(string(x[4*LENGTH5+LENGTH1:len(x)])," ")
l := strings.Trim(string(x[LENGTH5:2*LENGTH5])," ")
r := strings.Trim(string(x[2*LENGTH5:3*LENGTH5])," ")
if( y == s){
   f.Seek(-int64(LENGTH),SEEK_CUR)
   return f
}
if( s <y  ){
  if( l[0] == '-'){
     f.Close()
     return nil
  }
  l1,e := strconv.Atoi(l)
  if( e!= nil){
    log.Println("filo.seek l , e = ",e)
    os.Exit(0)
  }  
  f.Seek( int64(l1*LENGTH), SEEK_SET)
  return seek(n,f,s)
}
if( s >y  ){
  if(r[0]=='-'){
     f.Close()
     return nil
  }
  r1,e := strconv.Atoi(r)
  if( e!= nil){
    log.Println("filo.seek r , e= ",e)
    os.Exit(0)
  }  
  f.Seek( int64(r1*LENGTH), SEEK_SET)
  return seek(n,f,s)
}
log.Println("filo.seek return error")
os.Exit(0)
return nil
}





func GetRandom(fname string)string{
f,e := os.OpenFile(fname, os.O_RDONLY,0644)
if( e != nil){
   log.Println("filo.GetRandom e=",e)
   os.Exit(0)
}
defer f.Close()
w := GetBytes(f)
u := strings.Trim(string(w[0:LENGTH5])," ")
n,err := strconv.Atoi(u)
if(err != nil){
   log.Println("filo.Getrandom err = ",err)
   os.Exit(0)
}
rs := rand.NewSource(time.Now().UnixNano())
ra := rand.New(rs)
i := ra.Intn(n)
i0 := i
f.Seek(int64(i*LENGTH),SEEK_BEG)
r := GetBytes(f)
if( r[ 4*LENGTH5] == '-'){
   return string(r[4*LENGTH5+LENGTH1:LENGTH])
}
c := byte('+')
for c == '+' && i<n-1{
  i+=1
  f.Seek(int64(i*LENGTH),SEEK_BEG) 
  r = GetBytes(f)
  c = r[4*LENGTH5]
//  log.Println("c =",string(c))
}
if(c == '-'){
  return string(r[4*LENGTH5+LENGTH1:LENGTH])
}
i = i0
for c == '+' && i>0{
  i-=1
  f.Seek(int64(i*LENGTH),SEEK_BEG) 
  r = GetBytes(f)
  c = r[4*LENGTH5]
//  log.Println("c =",string(c))  
}
if(c == '-'){
  return string(r[4*LENGTH5+LENGTH1:LENGTH])
}
//log.Println("filo.GetRandom: nothin to be done in file ",fname)
return ""
}







func ChangeToDone(fname string, s string)bool{
f := Seek(fname,s)
if( f !=nil){
   f.Seek(int64(4*LENGTH5),SEEK_CUR)
   r,e :=f.Write([]byte("+"))
   if( r<1 || e != nil){
      log.Println("ChangeToDone")
      os.Exit(0)
   }
   f.Close()
   return true
}
return false
}






func GetRandomAndChange(fname string)string{
f,e := os.OpenFile(fname, os.O_RDWR,0644)
if( e != nil){
   log.Println("filo.GetRandom e=",e)
   os.Exit(0)
}
defer f.Close()
w := GetBytes(f)
u := strings.Trim(string(w[0:LENGTH5])," ")
n,err := strconv.Atoi(u)
if(err != nil){
   log.Println("filo.Getrandom err = ",err)
   os.Exit(0)
}
rs := rand.NewSource(time.Now().UnixNano())
ra := rand.New(rs)
i := ra.Intn(n)
i0 := i
f.Seek(int64(i*LENGTH),SEEK_BEG)
r := GetBytes(f)
if( r[ 4*LENGTH5] == '-'){
   f.Seek(-int64(LENGTH0+LENGTH1),SEEK_CUR)
   f.Write([]byte("+"))
   return string(r[4*LENGTH5+LENGTH1:LENGTH])
}
c := byte('+')
for c == '+' && i<n-1{
  i+=1
  f.Seek(int64(i*LENGTH),SEEK_BEG) 
  r = GetBytes(f)
  c = r[4*LENGTH5]
}
if(c == '-'){
  f.Seek(-int64(LENGTH0+LENGTH1),SEEK_CUR)
  f.Write([]byte("+"))
  return string(r[4*LENGTH5+LENGTH1:LENGTH])
}
i = i0
for c == '+' && i>0{
  i-=1
  f.Seek(int64(i*LENGTH),SEEK_BEG) 
  r = GetBytes(f)
  c = r[4*LENGTH5]
}
if(c == '-'){
  f.Seek(-int64(LENGTH+LENGTH1),SEEK_CUR)
  f.Write([]byte("+"))
  return string(r[4*LENGTH5+LENGTH1:LENGTH])
}
//log.Println("filo.GetRandom: nothin to be done in file ",fname)
return ""
}






func AsSlice(fname string)[]string{
y := make([]string,0)
f,e := ioutil.ReadFile(fname)
if(e != nil){
   log.Println("filo.AsSlice e ",e)
   return y
}
g := string(f)
n := len(g)/LENGTH
for i:=0;i<n;i++{
  s := g[i*LENGTH:(i+1)*LENGTH]
  x := s[4*LENGTH5+LENGTH1:LENGTH]
  y = append(y,x)
}
return y
}


func AsSlice0(fname string)([]string,[]string){
y := make([]string,0)
y1 := make([]string,0)
f,e := ioutil.ReadFile(fname)
if(e != nil){
   log.Println("filo.AsSlice e ",e)
   return y,y1
}
g := string(f)
n := len(g)/LENGTH
for i:=0;i<n;i++{
  s := g[i*LENGTH:(i+1)*LENGTH]
  x := s[4*LENGTH5+LENGTH1:LENGTH]
  z := s[4*LENGTH5:4*LENGTH5+LENGTH1]
  u := strings.Trim(string(z)," ")
  y = append(y,x)
  y1 = append(y1,u)  
}
return y,y1
}






