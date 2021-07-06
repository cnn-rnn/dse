package text

import (
)


func Find(s string , x string, i int) int{
n := len(s)
m := len(x)
for j:= i; j<=n-m ;j++{
   b := true
   for k:=j; k<j+m; k++{
      if( s[k] != x[k-j]){
         b = false
         break
      }
   }
   if(b){
     return j
   }
}
return -1
}


func Rfind(s string , x string, i int) int{
n := len(s)
m := len(x)
i1 := i
if(i > n-m){
   i1 = n-m
}
for j:= i1; j>=0 ;j--{
   b := true
   for k:=j; k<j+m; k++{
      if( s[k] != x[k-j]){
         b = false
         break
      }
   }
   if(b){
     return j
   }
}
return -1
}



func min(a int,b int)int{
if(a<b){
  return a
}else{
  return b
}
}

func max(a int,b int)int{
if(a>b){
  return a
}else{
  return b
}
}


func Remove(s string, a string, b string)(string,[]string){
var x string
var w []string
n := len(s)
i := 0
for i<n{
  j := Find(s,a,i)
  if(j==-1){
     x += string([]byte(s[i:n]))
     break
  }
  k := Find(s,b,j)
  if( k== -1){
     x += string([]byte(s[i:n]))
     break
  }
  x += " "+string([]byte(s[i:j]))
  w = append(w,string([]byte(s[j:k+len(b)])))
  i = k+len(b)
}
return x,w
}
     



func Remove0(s string, a string, b string)string{
var x string
n := len(s)
i := 0
for i<n{
  j := Find(s,a,i)
  if(j==-1){
     x += string([]byte(s[i:n]))
     break
  }
  k := Find(s,b,j)
  if( k== -1){
     x += string([]byte(s[i:n]))
     break
  }
  x += " "+string([]byte(s[i:j]))
  i = k+len(b)
}
return x
}
     




func Btw(s string, a string, b string)(string,[]string){
var x string
var w []string
n := len(s)
i := 0
for i<n{
  j := Find(s,a,i)
  if(j==-1){
     x += string([]byte(s[i:n]))
     break
  }
  k := Find(s,b,j)
  if( k== -1){
     j1:=j-10
     if(j1<0){
       j1 = 0
     }
     j2 := j+30
     if(j2>len(s)){
        j2 = len(s)
     }
     x += string([]byte(s[i:n]))
     break
  }
  x += string([]byte(s[i:k+len(b)]))
  w = append(w,string([]byte(s[j:k+len(b)])))
  i = k+len(b)
}
return x,w
}



func Btw0(s string, a string, b string)([]string){
var w []string
n := len(s)
i := 0
for i<n{
  j := Find(s,a,i)
  if(j==-1){
     break
  }
  k := Find(s,b,j)
  if( k== -1){
     break
  }
  w = append(w,string([]byte(s[j:k+len(b)])))
  i = k+len(b)
}
return w
}



func Btw0m(s string, a string, b string)(map[string]bool){
var w = make(map[string]bool)
n := len(s)
i := 0
for i<n{
  j := Find(s,a,i)
  if(j==-1){
     break
  }
  k := Find(s,b,j)
  if( k== -1){
     break
  }
  w[string([]byte(s[j:k+len(b)]))] = true
  i = k+len(b)
}
return w
}






//----------------------


func Filt(s string)string{
var x string
for _,c := range s {
  if( !( ( 65<=c && c<=90) ||( 97 <=c && c <= 122  )  )){
    c = 32
  }
  x+=string(c)
}
return x
}


func Filt1(s string)string{
n := len(s)
x := make([]byte, n)
for i:=0;i<n;i++{
   c:=s[i]
   if( !( ( 65<=c && c<=90) ||( 97 <=c && c <= 122  )  )){
      x[i] =32
   }else{
      x[i]= c
   }
}   
return string(x)
}




func Replace(s string,a string, b string)string{
var x string
i:=0
n:=len(s)
for i<n{
  j:= Find(s,a,i)
  if(j == -1){
     y := string([]byte(s[i:n]))
     x+=y
     break
  }
  y := string([]byte(s[i:j]))
  x += y+b
  i = j + len(a)
}
return x
}






func Chop(s string)string{
x := s
i := Find(x,"#",0)
if(i!=-1){
   x = string([]byte(x[0:i]))
}
i = Find(x,"?",0)
if(i!=-1){
   x = string([]byte(x[0:i]))
}
return x
}




//       -------------------------------------




func Host(s string)string{
i := Find(s,"//",0)
if(i==-1){
   return ""
}
i+=2
n := len(s)
j:= Find(s,"/",i)
if(j!= -1){
   n = j
}
k := Rfind(s,".",n)
if(k==-1){
   return ""
}
k = Rfind(s,".",k-1)
if(k!=-1){
  i=k+1
}
if(i+3<len(s) && s[i:i+3]=="www"){
   i+=4
}
return string([]byte(s[i:n]))
}




func Host1(s string)string{
i := Find(s,"//",0)
if(i==-1){
   return ""
}
i+=2
n := len(s)
j:= Find(s,"/",i)
if(j!= -1){
   n = j
}
return string([]byte(s[0:n]))
}








func DetectNoquote(s string)int{
i := Find(s,"href",0)
if(i == -1){
  return 0
}
c :=0
for t :=0;t<i;t++{
  if(s[t]==34){
    c+=1
  }
}
 if(c%2 ==1){
   return 0
}


j := Find(s,"=",i)
if(j == -1){
  return 0
}
j++
n :=len(s)
for j<n{
 if(s[j] == ' '){
   j++
 }else{
   break
 }
}
if(j==n){
  return 0
}
if(int(s[j]) !=  34){
  return 0
}
return 1
}



func EndsWith(x string , s string)bool{
     n := len(x)
     var e int
     e = Find(x,s,0)
     if(e !=-1 && x[e:n] == s){
        return true
     }
     return false
}






func ParseHref(y string, bes string)string{

     j0 := Find(y,"href",0)
     if(j0==-1){
       return ""
     }
     nc := DetectNoquote(y)
     if(nc ==0){
        return ""
     }
     j := Find(y,"\"",j0)
     if( j==-1){
       return ""
     }
     k := Find(y,"\"",j+1)
     if( k==-1){
       return ""
     }
     x0 := string([]byte(y[j+1:k]))
     x := Chop(x0)
     if(Find(x,"javascript:",0)!= -1 || Find(x,"mailto:",0)!= -1 || Find(x,"tel:",0)!= -1){
        return ""
     }
     if(len(x)<4){
       return ""
     }
     if(x[0:4] != "http"){
       if(x[0:2] =="//"){
          x = "https:"+x
       }else{
          if(x[0]=='/' || bes[len(bes)-1]=='/'){
            x = "https://"+bes+x
          }else{
            x = "https://"+bes + "/" +x
          } 
       }
     }
     if( EndsWith(x,".gz") ||  EndsWith(x,".jpg") ||  EndsWith(x,".jpeg") ||  EndsWith(x,".png") ||  EndsWith(x,".pdf") ||  EndsWith(x,".tar")  ){
        return ""
     }
     return x
}



func GetHref(s string)string{
p:= Find(s,">",0)
if(p==-1){
   return ""
}
s = s[0:p]
i := Find(s,"href",0)
if(i ==-1){
   return ""
}
j := Find(s,"\"",i)
if(j ==-1){
   return ""
}
k := Find(s,"\"",j+1)
if(k ==-1){
   return ""
}
z:=string([]byte(s[j+1:k]))
return z
}



func HostUrl(s string)string{
i := Rfind(s,".",len(s))
if(i ==-1){
   return s
}
j := Rfind(s,".",i-1)
if(j ==-1){
   return s
}
return string([]byte(s[j+1:len(s)]))
}







