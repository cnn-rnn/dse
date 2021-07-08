package  protocol

import (
"strings"
"crypto/sha256"
"encoding/hex"

)


var SEP = "<!#####!>"



func Serialize( name string, links map[string]bool)string{
  x:=name+SEP
  for i:= range links{
    x += string([]byte(i))+SEP
  }
  return x
}


func Deserialize(s string)(string,[]string){
  q := strings.Split(s,SEP)
  return q[0],q[1:]
}


func Digest(s string)string{
y  := sha256.Sum256([]byte(s))
dg := hex.EncodeToString(y[:])
return dg
}





