package conv


import(
"log"
"os"
"strconv"
"strings"
"crypto/sha256"
"encoding/hex"

)



func S_i64(s string)int64{
z,e := strconv.ParseInt(strings.Trim(s," "), 10, 64)
if(e != nil){
   log.Println("conv.S_i64 e = ",e,"s= ",s)
   os.Exit(0)
}
return int64(z)
}



func I64_s(u int64, n int)string{
d := strconv.FormatInt(u,10)
c := make([]byte,n)
for i:=0;i<n;i++{
      c[i] = ' '
}
for i:=0;i<len(d);i++{
      c[i] = d[i]
}
return string(c)
}




func I_s(x int, n int)string{
d := strconv.Itoa(x)
c := make([]byte,n)
for i:=0;i<n;i++{
      c[i] = ' '
}
for i:=0;i<len(d);i++{
      c[i] = d[i]
}
return string(c)
}


func Pad(s string, i int)string{
if(len(s)>=int(i)){
  return s[0:i]
}
b:= make([]byte,i)
for j:= range(s){
  b[j] = s[j]
}
for j:= len(s);j<i;j++ {
  b[j] = ' '
}
return string(b)
}




func Hash(s string)string{
y  := sha256.Sum256([]byte(s))
id := hex.EncodeToString(y[:])
return id
}



func B_i(x []byte) int{
y := string(x)
z,e := strconv.Atoi(strings.Trim(y," "))
if(e != nil){
   log.Println("conv.B_i e = ",e,"x= ",string(x))
   os.Exit(0)
}
return z
}


func S_i(x string) int{
return B_i([]byte(x))
}



