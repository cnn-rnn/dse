package murl


import(
"os"
"strings"
"../text"
)



func Host(s string)(string,string){
s = strings.ToLower(s)
i := text.Find(s,"//",0)
if(i != -1){
  s = string([]byte(s[i+2:len(s)]))
}

r:= text.Find(s,":",0)
if( r!=-1){
  s = string([]byte(s[0:r]))
}
p:= text.Find(s,"?",0)
if( p!=-1){
  s = string([]byte(s[0:p]))
}
q:= text.Find(s,"#",0)
if( q!=-1){
  s = string([]byte(s[0:q]))
}
j:= text.Find(s,"/",0)
if(j != -1){
  s = string([]byte(s[0:j]))
}
if(len(s)>4 && s[0:4] == "www."){
    s = string([]byte(s[4:len(s)]))
}

if( text.Find(s,".",0) ==-1){
  return "","no ."
}else{
  if( len(s)==0 || s[0] == '.'){
    return "","murl 2"
  }
  return s,""
}
os.Exit(0)
return "",""
}

    
  






func TheHost(s string)string{
i := text.Rfind(s,".",len(s))
if(i ==-1){
  os.Exit(0)
}
j := text.Rfind(s,".",i-1)
if(j ==-1){
  return s
}
return string([]byte(s[j+1:len(s)]))
}



func Scheme(s string)string{
i := text.Find(s,"://",0)
if( i==-1){
  return "https"
}
return s[0:i]
}


func Ho(s string)string{
h0,_ := Host(s)
return TheHost(h0)
}



  





