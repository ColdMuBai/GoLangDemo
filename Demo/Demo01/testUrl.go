package main

import (
	"fmt"
	"io/ioutil"
	"io"
	"net/http"
	"os"
	"bufio"
	"strings"
)


func main(){

	

	httpGet()
	// httpIoCopy()


}



func httpIoCopy(){
	for _,url := range os.Args[1:] {
		resp,err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		out,err := os.Create("./test.zip")
		write := bufio.NewWriter(out) 
		defer out.Close()
		// b,err := ioutil.ReadAll(resp.Body)
		b,err := io.Copy(write,resp.Body)
		fmt.Println("write",b)
		if err != nil {
			fmt.Fprintf(os.Stderr,"Respone reading errMsg %s: %v\n",url,err)
		}
		write.Flush()
		// ioutil.WriteFile("./test.html",[]byte(b),0666);
		// fmt.Printf("%s",b)
	}
}



func httpGet(){
	for _,url := range os.Args[1:] {

		prefix := "https://"
		if !strings.HasPrefix(url,prefix) {
			url = prefix+url
		} 
		fmt.Println(url)
		resp,err := http.Get(url)
		if err != nil {
			continue
		}

		b,err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,"Respone reading %s: %v\n",url,err)
		}
		//输出HTTP响应的状态码
		fmt.Println(resp.Status)
		ioutil.WriteFile("./test.html",[]byte(b),0666);
		// fmt.Printf("%s",b)
	}


}