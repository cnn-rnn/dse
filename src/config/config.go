package config


import(
"strconv"
"strings"
"io/ioutil"
"../logger"
)


type Conf struct{
  dir0 string
  cpu float64
  bw float64
  disk float64
  mon string
}


func check(e error){
if(e!=nil){
  logger.Logee("error "+e.Error())
}
}



func Configure()*Conf{
cname := "/etc/dse/dsed.conf"
x := &Conf{}
r,e1 := ioutil.ReadFile(cname)
check(e1)
q := strings.Split(string(r),"\n")
var y float64
var e error
x.dir0 = q[1]
y,e = strconv.ParseFloat(q[3],64)
check(e)
x.cpu = y
y,e = strconv.ParseFloat(q[5],64)
check(e)
x.bw = y
y,e = strconv.ParseFloat(q[7],64)
check(e)
x.disk = y
x.mon = q[8]
return x
}


func (x * Conf)Dir0()string{
return x.dir0
}
func (x * Conf)Cpu()float64{
return x.cpu
}
func (x * Conf)Bw()float64{
return x.bw
}
func (x * Conf)Disk()float64{
return x.disk
}
func (x * Conf)MonAddr()string{
return x.mon
}






