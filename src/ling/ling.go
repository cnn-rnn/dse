package ling

import (
"strings"
//"os"
"html"
"../text"
)



func Words( s string)[]string{

s2 := Filter_9_10(s)

v := strings.ToLower(s2)
u := text.Replace(v,"/"," ")
p := strings.Split(u," ")
q := make([]string,0)
for i:=0;i<len(p);i++{
  if(p[i] != "" && p[i] != "\n"){
     q = append(q,p[i])
  }
}
return q
}


func Filter_9_10(x string)string{
s := []byte(x)
for i:=0; i<len(s);i++{
  if(s[i] == 9 || s[i] == 10 || s[i] == 13){
     s[i] = ' '
  }
}
return string(s)
}






func Filter_1(x string)string{
s := []byte(x)
for i:=0; i<len(s);i++{
  if(s[i] == 9 || s[i] == 10 || s[i] == 13 || s[i] == 34  || s[i] == ')'  || s[i] == '('    || s[i] == '-'  || s[i] == '$'  || s[i] == ':' || s[i] == ';' || s[i] == '*' || s[i] == '&' || s[i] == '^' || s[i] == '@' || s[i] == '_' || s[i] == '#' || s[i] == '!' || s[i] == '?' || s[i] == '}' || s[i] == '{' || s[i] == '[' || s[i] == ']' || s[i] == '%' || s[i] == '<' || s[i] == '>' || s[i] == '=' || s[i] == '|' || s[i] == '»'  || s[i] == '\t'  || s[i] == '\r'  || s[i] == '\n' ){
     s[i] = ' '
  }
}
return string(s)
}




func Filter512( s string)string{
x := html.UnescapeString(s)
y := ""
for _,c := range x{
  if c<512{
    y += string(c)
  }else{
    y += " "
  }
}     
return y
}



func Filter0(s string)string{
return Filter_1(Filter512(s))
}






