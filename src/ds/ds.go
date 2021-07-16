package ds


import(
"math/rand"
"os"
"time"
"../logger"
"../murl"
"../filo"
"../urlstore"
)



type Ds struct{
  dir0 string
  sites_dir string
  ra *rand.Rand
}


func Create(s string)*Ds{
sites_dir := s+"/sites"
if _, err := os.Stat(sites_dir); os.IsNotExist(err) {
  em := os.Mkdir(sites_dir, 0755)
    if(em != nil){
       logger.Logee("ds em = "+em.Error())
    }
}
rs := rand.NewSource(time.Now().UnixNano())
ra := rand.New(rs)
return &Ds{s,sites_dir,ra}
}



func (x * Ds)Store(s string){
h := murl.Ho(s)
id_h := urlstore.Convert(h)        
filo.AddString(x.dir0+"/host_ids.txt",id_h)

urlstore.AddString(x.dir0+"/hosts.txt",h)
path := x.sites_dir+"/"+h
if _, err := os.Stat(path); os.IsNotExist(err) {
  em := os.Mkdir(path, 0755)
    if(em != nil){
      logger.Logee("dse em = "+em.Error())
    }
}
id := urlstore.Convert(s)
b := filo.AddString(x.sites_dir+"/"+h+"/pages.txt",id)
if(b){
     urlstore.AddString(x.dir0+"/urls.txt",s)
}        
}



func (x * Ds) Rand()string{
hf,ehf := os.OpenFile(x.dir0+"/host_ids.txt",os.O_RDONLY,0644)
if(ehf != nil){
   logger.Logee("ehf = "+ehf.Error())
}
fi,_ := hf.Stat()
M := int(fi.Size())/filo.LENGTH
m := x.ra.Intn(M)
hf.Seek(int64(m*filo.LENGTH),filo.SEEK_BEG)
z := filo.GetBytes(hf)
hf.Close()
id_h := z[4*filo.LENGTH5+filo.LENGTH1:len(z)]
h := urlstore.GetUrl(x.dir0+"/hosts.txt"  ,string(id_h))
id := filo.GetRandom(x.sites_dir+"/"+h+"/pages.txt")
if(id == ""){
  return ""
}
u := urlstore.GetUrl(x.dir0+"/urls.txt"  ,id)
return u
}









