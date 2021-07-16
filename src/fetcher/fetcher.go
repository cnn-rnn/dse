package fetcher

import (
"net/http"
"net/url"
"strings"
"os"
"time"

"html"
"io"
"compress/gzip"
"../logger"
"../text"
"../murl"
)



var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    5 * time.Second,
	DisableCompression: true,
     }
var client = &http.Client{
         Transport: tr,
         Timeout: 10 * time.Second,
     }






func OnePageLinks(name string)(map[string]bool,string,int,int64,string){
  L := make(map[string]bool)
//start:
  req, err := http.NewRequest("GET", name, nil)
  if err != nil {
      logger.Logo("invalid url "+err.Error())
      return L,"",0,int64(0),"invalid url"+err.Error()
  }
  moz :=  "Mozilla/5.0 (X11; Linux i686; rv:88.0) Gecko/20100101 Firefox/88.0"

  req.Header.Set("User-Agent", moz)
  resp,err := client.Do(req)
  if(err != nil){
     logger.Logo("client.Do err="+err.Error())
     return L,"",0,int64(0),"client.Do err="+err.Error()
  }
  defer resp.Body.Close()
  
  var reader io.ReadCloser
  switch resp.Header.Get("Content-Encoding") {
    case "gzip":
      reader, err = gzip.NewReader(resp.Body)
      if(err != nil){
        logger.Logo("gunzip error ="+err.Error())
        return L,"",0,int64(0),"gunzip error ="+err.Error()
      }
      defer reader.Close()
    default:
      reader = resp.Body
      defer reader.Close()
  }  
  if(reader == nil){
      logger.Logo("reader is nil")
      return L,"",0,int64(0),"reader is nil"
  }
  typ := resp.Header.Get("Content-Type") 
  if(typ == "application/pdf"){
      logger.Logo("pdf crawled")
      return L,"",0,int64(0),"pdf crawled"
  }  
  if(typ == "image/jpeg"){
      logger.Logo("jpeg crawled")
      return L,"",0,int64(0),"jpeg crawled"
  }  
  if(strings.Index(typ,"text/html") == -1  && strings.Index(typ,"gzip") == -1){
     logger.Logo("non html non gzip non pdf non jpeg")
     return L,"",0,int64(0),"non html non gzip non pdf non jpeg"
  }
  buf := new(strings.Builder)
  io.Copy(buf, reader)
  s0 := buf.String()
 
  
//  s1 := text.Remove0(s0,"<script","</script>")
//  s2 := text.Remove0(s1,"<style","</style>")
//  s3 := text.Remove0(s2,"<!--","-->")

  href := text.Btw0m(s0,"<a ","</a>")
//  txt := text.Remove0(s3,"<",">")  
  
  host,e := url.Parse(name)
  if( e!= nil){
  }

  
  for i:= range href {
     z := i
     y := text.GetHref(z)
     y = strings.TrimSpace(y)
     
     y = html.UnescapeString(y)
     x,e1 := host.Parse(y)     
     if( e1!= nil){
       continue
     }
     u := x.String()
     u = strings.TrimSpace(u)
     
     if(text.Find(u,"\n",0)!=-1){
        os.Exit(0)
     }
     
     if j := text.Find(u,"#",0) ; j!=-1{
        u = u[0:j-1]
     }
     
     if(u ==""){
        continue
     }
     if(len(u)>=10 && u[0:10]=="javascript"){
         continue
     }
     if(len(u)>=6 && u[0:6]=="mailto"){
         continue
     }
     if(len(u)>=4 && u[0:4]=="tel:"){
         continue
     }
     if(len(u)>=4 && u[0:4]=="sms:"){
         continue
     }

     if(  text.Find(u,"//",0)==-1){
         continue
     }
     
     if(len(u)>7 && u[0:7] == "tencent"){
         continue
     }
     

     u = strings.ToLower(u)
     _,eu := murl.Host(u)
     if( eu != ""){
     }else{        
        L[u] = true
     }
  }
  return L,"",len(s0),time.Now().UnixNano(),""
}



