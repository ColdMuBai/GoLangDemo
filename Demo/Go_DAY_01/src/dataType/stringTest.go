package main


import (
	"fmt"
	"strings"
	_"strconv"
)

const str = "This is a example of a string"

func main(){
	// HasPrefixTest()
	// ContainsTest()
	//FindIndex()
	// ReplaceStr()
	// CountStr()
	// RepeatStr()
	// StrToLowerOrUpper()
	// TrimStr()
	// SliceStr()
	// JoinStr()
	readerStr()
}



func HasPrefixTest(){

	
	fmt.Printf("字符串 %vs 是否包含前缀 %s",str,"Th")
	fmt.Printf("\t%t\n",strings.HasPrefix(str,"Th"))	
}

func ContainsTest(){
	var subStr = "example"
	var flag bool = strings.Contains(str,subStr)
	fmt.Printf("string is contain %s : %t",subStr,flag)
}

//查找并返回字符串s 在 str中的第一次出现的索引
func FindIndex(){
	var s string= "a"
	fmt.Printf("s = %v\t str = %v\n",s,str)
	//在 str中的第一次出现的索引 <s --> 查找的子串 str --> 被查找的完整字符串>
	firstIndex  := strings.Index(str,s) 
	fmt.Printf("第一次出现的索引为 : %v\t%[1]T\n",firstIndex)
	//在 str中最后出现的索引
	lastIndex := strings.LastIndex(str,s)
	fmt.Printf("第一次出现的索引为 : %v\t%[1]T\n",lastIndex)

	//查询非 ASCII 编码的字符在父字符串中的位置
	s = "这是一段非ASCll编码的字符串,用来测试查询非 ASCII 编码的字符在父字符串中的位置"
	runeIndex := strings.IndexRune(s,'符')
	fmt.Printf("第一次出现的索引为(Rune) : %v\t%[1]T\n",runeIndex)
}
//查找并替换字符串中对应索引范围的字符
func ReplaceStr(){
	//参数说明 : 被替换的完整字符串  需要替换的字符串片段  替换的字符片段  替换区域的截止索引(-1为替换全部)
	newstr := strings.Replace(str,"example","replace test",9)
	// newstr := strings.Replace(str,"s","replace test",-1)
	fmt.Printf("newStr = %v\t %[1]T",newstr)
}
//统计字符串出现次数
func CountStr(){
	count := strings.Count(str,"i")
	fmt.Printf("count TYpe : %T\t value : %[1]v",count)
}
//重复count次字符串s并返回新字符串
func RepeatStr(){
	// newStr := strings.Repeat(strings.Join([]string{str},"\t"),10000)
	newStr := strings.Repeat(str+"\n",10000)
	fmt.Printf("Repeat str : %v",newStr)
}
//字符串大小写转换
func StrToLowerOrUpper(){
	upper := strings.ToUpper(str)
	fmt.Printf("To Upper str : %v\n",upper)
	lower := strings.ToLower(str)
	fmt.Printf("To Lower str : %v",lower)
}
//字符串修剪
func TrimStr(){
	const s string = " this is a test strings this    "
	const trimStr string = " this"
	str1 := strings.TrimSpace(s)					//去除字符串前后的空格
	fmt.Printf("TrimSpace string : %v\n",str1)		
	str2 := strings.Trim(s,trimStr)					//去除字符串左右的指定字符,没有符合的则返回原字符
	fmt.Printf("TrimSpace string : %v\n",str2)
	str3 := strings.TrimLeft(s,trimStr)				//去除字符串左侧的指定字符,没有符合的则返回原字符
	fmt.Printf("TrimSpace string : %v\n",str3)
	str4 := strings.TrimRight(s,"  this")			//去除字符串右侧的指定字符,没有符合的则返回原字符
	fmt.Printf("TrimSpace string : %vxxxxx\n",str4)
}
//分割字符串
func SliceStr(){
	const tempStr = "这*是*一*段*用*逗*号*分*割*的*字*符*串"
	respStr := strings.Fields(tempStr) //默认以空格进行切片
	fmt.Printf("分割后的结果为: %v\t type : %[1]T",respStr)
	respStr = strings.Split(tempStr,"*")//以指定字符进行切片
	fmt.Printf("分割后的结果为: %v\t type : %[1]T",respStr)
	strJoin := strings.Join(respStr,"")
	fmt.Printf("slice joined : %v\t %[1]T",strJoin,strJoin)
}
//拼接字符串切片数组
func JoinStr(){
	var testStr = []string {"这","是","一","段","用","逗","号","分","割","的","字","符","串"}
	fmt.Printf("testStr type is %v\t %[1]T\n",testStr)
	strJoin := strings.Join(testStr,"")
	fmt.Printf("slice joined : %v\t %[1]T",strJoin,strJoin)
}
//从字符串中读取内容
func readerStr(){

	read := strings.NewReader(str)
	fmt.Println(read.Len())
	var buf []byte
	buf = make([]byte,5)
	readLen,err := read.Read(buf)
	fmt.Println("读取到的数据长度:",readLen)
	if err != nil {
		fmt.Println("错误:",err)
	}
	fmt.Println(buf)	//this buf中读取到的内容
	fmt.Printf("%c : %[1]T\n",buf) 
	fmt.Println(read.Len())	//读取后read的剩余长度
	fmt.Println(read.Size())	//read所加载的字符串的大小
}