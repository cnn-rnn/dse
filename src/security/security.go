package security

import (
"os"
"io/ioutil"
"crypto"
"crypto/rand"
"crypto/rsa"
"crypto/x509"
"crypto/sha256"
"encoding/base64"
"../logger"
)


type Sec struct{
  PK *rsa.PrivateKey
}

func check(err error){
if err != nil {
panic(err)
}
}


func Create(dir0 string)*Sec{
exi := Load(dir0+"/security.txt")
if(exi != nil){
   return exi
}   
PK, err := rsa.GenerateKey(rand.Reader, 2048)
check(err)
x := &Sec{PK}
x.Save(dir0+"/security.txt")
return x
}



func (x *Sec)GetPK()*rsa.PrivateKey {
return x.PK
}


func (x * Sec) GetPubK()*rsa.PublicKey{
return &x.PK.PublicKey
}

func (x * Sec)GetPubKAsString()string{
return PubKToString(&x.PK.PublicKey)
}


func ToString(PK * rsa.PrivateKey)string{
m := x509.MarshalPKCS1PrivateKey(PK)
sm := base64.StdEncoding.EncodeToString(m)
return string(sm)
}

func PubKToString(p * rsa.PublicKey)string{
m := x509.MarshalPKCS1PublicKey(p)
sm := base64.StdEncoding.EncodeToString(m)
return string(sm)
}



func PrivKToString(PK * rsa.PrivateKey)string{
p := PK.PublicKey
m := x509.MarshalPKCS1PublicKey(&p)
sm := base64.StdEncoding.EncodeToString(m)
return string(sm)
}


func Base64EncodePrivate(p * rsa.PrivateKey)string{
m := x509.MarshalPKCS1PrivateKey(p)
sm := base64.StdEncoding.EncodeToString(m)
return sm
}

func Base64EncodePublic(p * rsa.PublicKey)string{
m := x509.MarshalPKCS1PublicKey(p)
sm := base64.StdEncoding.EncodeToString(m)
return sm
}



func (x * Sec)Save(fname string){
m := x509.MarshalPKCS1PrivateKey(x.PK)
sm := base64.StdEncoding.EncodeToString(m)
f,e := os.OpenFile(fname,os.O_CREATE|os.O_WRONLY,0644)
check(e)
f.Truncate(0)
f.Write([]byte(sm))
f.Close()
}


func Base64DecodePrivate(sm []byte)*rsa.PrivateKey{
n := base64.StdEncoding.DecodedLen(len(sm))
var buff = make( []byte,n)
l, _ := base64.StdEncoding.Decode(buff, []byte(sm))
p , e1 := x509.ParsePKCS1PrivateKey(buff[:l])
if e1 != nil {
  logger.Loge("Base64DecodePrivate ")
  return nil
}
return p
}

func Base64DecodePublic(sm []byte)*rsa.PublicKey{
n := base64.StdEncoding.DecodedLen(len(sm))
var buff = make( []byte,n)
l, _ := base64.StdEncoding.Decode(buff, []byte(sm))
p , e1 := x509.ParsePKCS1PublicKey(buff[:l])
if e1 != nil {
  logger.Loge("Base64DecodePublic "+e1.Error())
  return nil
}
return p
}


func Load(fname string)*Sec{
sm,e := ioutil.ReadFile(fname)
if(e!=nil){
  logger.Loge("security.Load e="+e.Error())
  return nil
}
n := base64.StdEncoding.DecodedLen(len(sm))
var buff = make( []byte,n)
l, _ := base64.StdEncoding.Decode(buff, []byte(sm))
var e1 error
PK , e1 := x509.ParsePKCS1PrivateKey(buff[:l])
check(e1)
return &Sec{PK}
}



func Digest(msg []byte)string{
msgHash := sha256.New()
var err error
_, err = msgHash.Write(msg)
if err != nil {
  logger.Logee("security.Digest"+err.Error())
}
msgHashSum := msgHash.Sum(nil)
return string(msgHashSum)
}


func (x * Sec)Sign(msg []byte)[]byte{
msgHash := sha256.New()
var err error
_, err = msgHash.Write(msg)
if err != nil {
  logger.Logee("security.Sign err="+err.Error())
}
msgHashSum := msgHash.Sum(nil)
signature, e1 := rsa.SignPSS(rand.Reader, x.PK, crypto.SHA256, msgHashSum, nil)
if e1 != nil {
  logger.Logee("security.Sign e1="+e1.Error())
}
return signature
}




func (x * Sec)SignS(msg []byte)string{
return base64.StdEncoding.EncodeToString(x.Sign(msg))
}



func Verify(msgHashSum []uint8,signature []uint8 ,PubK * rsa.PublicKey )bool{
err := rsa.VerifyPSS(PubK, crypto.SHA256, msgHashSum, signature, nil)
if err != nil {
  return false
}
return true
}


func VerifyS(msg []byte,signature []uint8 ,PubK * rsa.PublicKey )bool{
msgHashSum := Digest(msg)
err := rsa.VerifyPSS(PubK, crypto.SHA256, []byte(msgHashSum), signature, nil)
if err != nil {
  return false
}
return true
}


func VerifySS(msg []byte,signature string ,PubK * rsa.PublicKey )bool{
s,e := base64.StdEncoding.DecodeString(signature)
if( e != nil){
  logger.Loge("security.VerifySS"+e.Error())
  return false
}
return VerifyS(msg, s, PubK)
}




