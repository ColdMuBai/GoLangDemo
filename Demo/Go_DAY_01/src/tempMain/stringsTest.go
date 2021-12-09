package main

import(
	"fmt"
	// "strings"
)


func main(){

	str1,str2 :=ActionScope("this is a lowercase string","THIS IS AN UPPERCASE STRING")
	fmt.Printf("\nlowerCase : "+str1+"\nupperCase :"+str2);
}


func  ActionScope(sm string,lg string) (string,string){
	var lgStr,smStr string
	for _,i := range sm {
		if i != '!'{
			if i!=' '{
				fmt.Printf("%c",i)
			}else {
				fmt.Print(" ")
			}
			i := i+'A'-'a'
			lgStr += string(i)
		}
	}

	fmt.Println("\n----------------------------------")
	
	for i := 0; i < len(lg);i++{
		i := lg[i]
		if i != '$'{
			if i!=' ' {
				fmt.Printf("%c",i)
			} else {
				fmt.Print(" ")
			}
			i := i+'a'-'A'
			smStr += string(i)
		}
	}
	fmt.Println("\n----------------------------------")
	return lgStr, smStr
}