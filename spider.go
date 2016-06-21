package main
import(
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
	"time"
	"strings"
	"runtime"
	"regexp"
)

type superchao struct{
	urls chan string // 带爬取
	re *regexp.Regexp // 正则表达式
}

func main(){
		runtime.GOMAXPROCS(5)
		c:=superchao{
				make(chan string,3),
				regexp.MustCompile("<input type=\"submit\".*?value=\"(.*?)\""),
			}
		c.initialUrl()
		c.start()
		time.Sleep(time.Second*5)
		c.findLine("www.baidu.com")
}
func (c *superchao)initialUrl(){
	c.urls<-"http://www.baidu.com"
	close(c.urls)
}

func (c *superchao)findLine(path string){
	content,_:=ioutil.ReadFile(path)
	result:=c.re.FindStringSubmatch(string(content))
	fmt.Printf("%v",result[1])
}

func (c *superchao)start(){
	
	for x:=range c.urls{
		go func(){
			fmt.Println(x)
		   	 spider(x,strings.Split(x,"//")[1])			
		}()
		time.Sleep(time.Second*1)
	}
	

}



func  spider(url,path string){
	file,_:=os.OpenFile(path,os.O_CREATE|os.O_RDWR|os.O_APPEND,0666)
	resp,_:=http.Get(url)
	body,_:=ioutil.ReadAll(resp.Body)	
	file.WriteString(string(body))

}
