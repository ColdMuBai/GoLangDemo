package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)



func main(){


	// reuseIoRead()

	reuseInFile()

	// reuse()
}


func reuseIoRead() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//一次读取全部文件
		data,err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr,"error msg: %v\n",err)
			continue
		}

		for _,line := range strings.Split(string(data),"\n") {
			counts[line]++
		}

	}

	for line,n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}



}



//从具名文件中读取数据找出重复的行数据,如果没有文件输入则读取cmd输入的数据
func reuseInFile(){
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin,counts)
	} else {
		for _,arg := range files {
			fmt.Println(arg)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr,"error msg: %v\n",err)
				continue
			}
			// countLines(f,counts)
			countLinesRE(arg,f,counts)
			f.Close();
		}
	}

		
	// for line,n := range counts {

	// 	sort.Strings(counts)
	// 	if n > 1{
	// 		// fmt.Printf("%s\n",files)
	// 		fmt.Printf("%d\t%s\n",n,line)
	// 	}
	// }

	var keys []string
	for k,v :=range counts {
		if v > 1 {
			keys = append(keys,k)
		}
	}
	sort.Strings(keys)
	for _,k := range keys {
		fmt.Printf("%d\t%s\n",counts[k],k)
	}

}


func countLinesRE(filename string,f *os.File,counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if strings.Compare(input.Text(),"over") == 0 {
			break
		}
		counts[filename+": "+input.Text()]++
	}

}


func countLines(f *os.File,counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if strings.Compare(input.Text(),"over") == 0 {
			break
		}
		counts[input.Text()]++
	}

}



//测试输入数据中重复的行
func reuse(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++
		if input.Text() == "over"{
			break
		}
	}

	for k,v := range counts {
		if v >1 {
			fmt.Printf("%d\t%s\n",v,k)
			fmt.Printf("%v\t%v\n",v,k)
			// fmt.Printb("%v",v,k)
		}
	}




}