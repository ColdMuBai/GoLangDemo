package main

import(
	"fmt"
	"log"
	"net/http"
	"sync"
	"css"
	"strconv"
	"os"
)


var mu sync.Mutex
var count int
var cycles int=0;

func main(){
	http.HandleFunc("/",httpServer)
	http.HandleFunc("/count",counter)
	http.HandleFunc("/gif",func(w http.ResponseWriter,r *http.Request){
		css.Lissajous(w)
	})
	http.HandleFunc("/number",ChangeQuery)
	http.HandleFunc("/golanghtml",OutPutHtml)

	log.Fatal(http.ListenAndServe("127.0.0.1:9600",nil))
}

func OutPutHtml(w http.ResponseWriter,r *http.Request){
	f,err := os.Open("./GoLang学习笔记.html")
	//defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
	defer f.Close() 
	if err !=nil {
		//主动触发宕机
		panic(err)
	} else {
		buffer := make([]byte,20)
		//func (f *File) Read(b []byte) (m int,err errpr)
		//使用read方法,将文件内容读入到buffer切片中
		for {
			// length,err := f.Read(buffer)
			_,err := f.Read(buffer)
			if err != nil {
				panic(err)
			} else {
				fmt.Fprintf(w,"%v",string(buffer))
			}
		}
	}
}

func ChangeQuery(w http.ResponseWriter,r *http.Request){
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		for k,v := range r.Form {
			if(k == "cycles"){
				var str string= ""
				for _,i := range v{
					str += i
				}
				resp,_ := strconv.Atoi(str);
				cycles = resp
				fmt.Println(cycles)
			}
		}
		fmt.Fprintf(w,"cycles = %d",cycles)
		fmt.Println(cycles)

	}

func printHeader(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"%s %s %s\n",r.Method,r.URL,r.Proto)
	for k,v := range r.Header {
		fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
	}

	fmt.Fprintf(w,"Host = %q\n",r.Host)
	fmt.Fprintf(w,"RemoteAddr = %q\n",r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k,v := range r.Form {
		fmt.Fprintf(w,"Form[%q] = %q\n",k,v)
	}
}



func httpServer(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w,"HTTP Request Url : =%q\n",r.URL.Path)
	printHeader(w,r);
}

func counter(w http.ResponseWriter,r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w,"服务器被请求了: %d次\n",count)
	mu.Unlock()
}
