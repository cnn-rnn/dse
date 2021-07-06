package urlstore

	

import (
"strconv"
"strings"
"log"
"io"
"os"
"crypto/sha256"
"encoding/hex"
)


var SEEK_BEG = 0
var SEEK_SET = 0
var SEEK_END = 2
var SEEK_CUR = 1


var LENGTH5 = 12
var LENGTH0 = 64

var LENGTH = 6*LENGTH5 + LENGTH0


func GetBytes(f * os.File)[]byte{
r := make([]byte, LENGTH)
i,e := io.ReadAtLeast(f,r,LENGTH)
if( e != nil){
   log.Println("urlstore.GetBytes i,e =",i,e )
   os.Exit(0)
}
if( i< LENGTH){
   log.Println("urlstore.GetBytes i < Length i = ", i)
   os.Exit(0)
}
beg0 := r[4*LENGTH5:5*LENGTH5]
end0 := r[5*LENGTH5:6*LENGTH5]
beg1 := string(beg0)
end1 := string(end0)
beg,e1 := strconv.Atoi(strings.Trim(beg1," "))
if(e1 != nil){
   log.Println("urlstore.GetBytes e1 = ",e1,beg1)
   os.Exit(0)
}
end,e1 := strconv.Atoi(strings.Trim(end1," "))
if(e1 != nil){
   log.Println("urlstore.GetBytes e1 = ",e1,end1)
   os.Exit(0)
}
//f.Seek(beg,SEEK_BEG)     ---  already at the right place
r2 := make([]byte,end-beg)
j,e2 := io.ReadAtLeast(f,r2,end-beg)
if( e2 != nil){
   log.Println("urlstore.GetBytes j,e2 =",j,e2 )
   os.Exit(0)
}
if( j< end-beg){
   log.Println("urlstore.GetBytes j < end-beg j = ", j)
   os.Exit(0)
}
r3 := make([]byte,LENGTH+end-beg)
for i:=0;i<LENGTH;i++{
  r3[i] = r[i]
}
for i:=0;i<end-beg;i++{
  r3[LENGTH+i] = r2[i]
}
return r3
}


func Convert(s string)string{
y  := sha256.Sum256([]byte(s))
id := hex.EncodeToString(y[:])
return id
}





func Create(fname string, s string){
if(len(s) > 10000){
  log.Println("urlstore.Create : string too long",s)
  return
}
n := len(s)
id := Convert(s)
f, _ := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0644)
defer f.Close()
f.Truncate(0)
e := ""
for i:=0;i<LENGTH5-1;i++{
  e+=" "
}
u := strconv.Itoa(6*LENGTH5+LENGTH0)
v := strconv.Itoa(6*LENGTH5+LENGTH0+n)

x := make([]byte,LENGTH5)
for i:=0;i<LENGTH5;i++{
  x[i] =' '
}
for i:=0;i<len(u);i++{
  x[i] = u[i]
}
y := make([]byte,LENGTH5)
for i:=0;i<LENGTH5;i++{
  y[i] =' '
}
for i:=0;i<len(u);i++{
  y[i] = v[i]
}
//   P     L     R     T   B E ID
r := "0"+e+"-"+e+"-"+e+"0"+e+string(x)+string(y)+id+s
//log.Println("create r=",r)
f.Write([]byte(r))
}



func int_to_bytes(i int64)[]byte{
d := strconv.FormatInt(i,10)
c := make([]byte,LENGTH5)
for i:=0;i<LENGTH5;i++{
      c[i] = ' '
}
for i:=0;i<len(d);i++{
      c[i] = d[i]
}
return c
}



func CreateNewRecord(m int64, parent_offset string, f * os.File, s string){
f.Seek(int64(m),SEEK_BEG)
fi, _ := f.Stat()
n := fi.Size()
this := int_to_bytes(n)
f.Write(this)
f.Seek(0,SEEK_END)
e :=""
for i:=0;i<LENGTH5-1;i++{
  e+=" "
}
beg := int_to_bytes(n+int64(LENGTH))
end := int_to_bytes(n+int64(LENGTH+len(s)))
t := parent_offset+  "-"+e  +  "-"+e  + string(this) + string(beg) + string(end) +  Convert(s) +  s
//log.Println("new record t =",t)
f.Write([]byte(t))
}







func AddString(fname string, s string){
f,e := os.OpenFile(fname,os.O_RDWR,0644)
if(e != nil){
  if(os.IsNotExist(e) ){
     Create(fname, s)
     return
  }else{
     log.Println("urlstore.AddString: e= ",e, os.IsNotExist(e) )
     os.Exit(0)
  }
}
defer f.Close()
AddStringSolid(f,s)
}



func AddStringSolid(f *os.File, s string){
b := make([]byte,LENGTH5)
f.Seek(0,SEEK_SET)
f.Read(b)
n0 := string(b)
n0 = strings.Trim(n0," ")
n,err := strconv.Atoi(n0)
if(err != nil){
   log.Println("urlstore.AddStringSolid e = ",err," n0 = ",n0)
   os.Exit(0)
   return
}
f.Seek(0,SEEK_BEG)
w := Add_string(0,f,s)
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
}



func Add_string(n int64, f * os.File, s string)bool{
id := Convert(s)
x := GetBytes(f)
y := strings.Trim(string(x[6*LENGTH5:6*LENGTH5+LENGTH0])," ")
l := string(x[LENGTH5:2*LENGTH5])
r := string(x[2*LENGTH5:3*LENGTH5])
t := string(x[3*LENGTH5:4*LENGTH5])
//log.Println(s,y,s<y,s>y, len(s),len(y))
if( y == id){
   return false 
}
if( id < y){
  if( l[0] == '-'){
     CreateNewRecord(n+int64(LENGTH5),t,f,s)
     return true
  }else{
     m,e := strconv.ParseInt(strings.Trim(l," "),10,64)
     if(e != nil){
       log.Println("urlstore.Add_string left : m, e =",m,e, "x=",string(x))
       os.Exit(0)
     }
     f.Seek(int64(m) , SEEK_SET)
     return Add_string(m,f,s)
  }
}    
if( id > y){
  if( r[0] == '-'){
     CreateNewRecord(n+2*int64(LENGTH5),t,f,s)
     return true
  }else{
     m,e := strconv.ParseInt(strings.Trim(r," "),10,64)
     if(e != nil){
        log.Println("urlstore.Add_string right : m, e =",m,e, "x=",string(x))
        os.Exit(0)
     }
     f.Seek(int64(m) , SEEK_SET)
     return Add_string(m,f,s)
  }
}
return false    
}





func Seek(fname string, id string)*os.File{
f, err := os.OpenFile(fname,os.O_RDWR,0655)
if( err != nil){
  log.Println(err)
}
return seek(f,id)
}


func seek( f * os.File, id string)*os.File{
x := GetBytes(f)
y := strings.Trim(string(x[6*LENGTH5:6*LENGTH5+LENGTH0])," ")
l := strings.Trim(string(x[LENGTH5:2*LENGTH5])," ")
r := strings.Trim(string(x[2*LENGTH5:3*LENGTH5])," ")
if( y == id){
   f.Seek(-int64(len(x)),SEEK_CUR)
   return f
}
if( id <y  ){
  if( l[0] == '-'){
     f.Close()
     return nil
  }
  l1,e := strconv.Atoi(l)
  if( e!= nil){
    log.Println("urlstore.seek l , e = ",e)
    os.Exit(0)
  }  
  f.Seek( int64(l1), SEEK_SET)
  return seek(f,id)
}
if( id >y  ){
  if(r[0]=='-'){
     f.Close()
     return nil
  }
  r1,e := strconv.Atoi(r)
  if( e!= nil){
    log.Println("urlstore.seek r , e= ",e)
    os.Exit(0)
  }  
  f.Seek( int64(r1), SEEK_SET)
  return seek(f,id)
}
log.Println("urlstore.seek return error")
os.Exit(0)
return nil
}


func GetUrl(fname string, id string)string{
f := Seek(fname,id)
if(f != nil){
   x := GetBytes(f)
   f.Close()
   n := len(x)
   n0 := 6*LENGTH5+LENGTH0
   return string(x[n0:n])
}
return ""
}














