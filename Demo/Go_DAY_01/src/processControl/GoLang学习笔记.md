# GoLang学习笔记

## 基础命令

### 1.go 环境变量

1.在系统环境变量中新建 ==GoPath==

```go
变量值: 	D:\Environment\GO
```

2.在系统==path==末尾添加

```go
%GoPath%\bin
```

## go基本命令

### 1.go build命令

> go build 无参数编译

```go
// 编译，生成可执行文件
go build ***.go
// 运行可执行文件，windows上会加上exe，linux没有后缀
./可执行文件
// go build+文件列表，会生成多个可执行文件
go build ***.go ***.go ……
```

> go build 参数设置

| 参数  | 备注                                        |
| ----- | ------------------------------------------- |
| -o    | 编译的包名                                  |
| -v    | 编译时显示包名                              |
| -p n  | 开启并发编译，默认情况下该值为 CPU 逻辑核数 |
| -a    | 强制重新构建                                |
| -n    | 打印编译时会用到的所有命令，但不真正执行    |
| -x    | 打印编译时会用到的所有命令                  |
| -race | 开启竞态检测                                |

### 2.go clean 命令

> go clean 命令是用来移除当前源码包和关联源码包里面编译生成的文件。这些文件包括：

```go
_obj/ 旧的 object 目录，由 Makefiles 遗留
_test/ 旧的 test 目录，由 Makefiles 遗留
_testmain.go 旧的 gotest 文件，由M akefiles 遗留
test.out 旧的 test 记录，由 Makefiles 遗留
build.out 旧的 test 记录，由 Makefiles 遗留
*.[568ao] object 文件，由 Makefiles 遗留
DIR(.exe) 由 go build 产生
DIR.test(.exe) 由 go test -c 产生
MAINFILE(.exe) 由 go build MAINFILE.go 产生
*.so 由 SWIG 产生
```

> 参数说明: 

```go
//一般使用 go clean -i -n
-i 清除关联的安装的包和可运行文件，也就是通过 go install 安装的文件
-n 把需要执行的清除命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的
-r 循环的清除在 import 中引入的包
-x 打印出来执行的详细命令，其实就是 -n 打印的执行版本
```

### 3.go run 命令

> go run  编译源码，并且执行源码的 main() 函数，不会留下可执行文件。

go run 不能使用“go run+包”的方式进行编译，如需快速编译运行包，需要使用如下步骤来代替：

```go
//(编译并运行)
使用 go build 生成可执行文件。
运行可执行文件。
```

### 4.go fmt命令

> 格式化代码文件

- go fmt -w -l src，可以格式化整个项目。所以 go fmt 是 gofmt 的上层一个包装的命令。

> 参数介绍

- -l 显示那些需要格式化的文件
- -w 把改写后的内容直接写入到文件中，而不是作为结果打印到标准输出。
- -r 添加形如“a[b:len(a)] -> a[b:]”的重写规则，方便我们做批量替换
- -s 简化文件中的代码
- -d 显示格式化前后的 diff 而不是写入文件，默认是 false
- -e 打印所有的语法错误到标准输出。如果不使用此标记，则只会打印不同行的前 10 个错误。
- -cpuprofile 支持调试模式，写入相应的 cpufile 到指定的文件

### 5.go install命令

> (编译并安装) go install  将编译的中间文件放在 GOPATH 的 pkg 目录，将编译结果放在 GOPATH 的 bin 目录下

- 第1步是生成结果文件（可执行文件或者 .a 包），
- 第2步会把编译好的结果移到 $GOPATH/pkg 或者 $GOPATH/bin。

> go install 的编译过程有如下规律：

- go install 是建立在 GOPATH 上的，无法在独立的目录里使用 go install。
- GOPATH 下的 bin 目录放置的是使用 go install 生成的可执行文件，可执行文件的名称来自于编译时的包名。
- go install 输出目录始终为 GOPATH 下的 bin 目录，无法使用-o附加参数进行自定义。
- GOPATH 下的 pkg 目录放置的是编译期间的中间文件。

### 6. go get 命令

> 获取包，编译并安装

```go
go get -u + 远程包
```

> 参数介绍：

- -d 只下载不安装
- -f 只有在你包含了 -u 参数的时候才有效，不让 -u 去验证 import 中的每一个都已经获取了，这对
- 于本地 fork 的包特别有用
- -fix 在获取源码之后先运行fix，然后再去做其他的事情
- -t 同时也下载需要为运行测试所需要的包
- -u 强制使用网络去更新包和它的依赖包
- -v 显示执行的命令

### 7.go generate命令

go generate 命令是用于在编译前自动化生成某类代码。

```go
/**
*在同一目录下执行go generate就会自动运行命令command arg1 arg2。
*command可以是在PATH中的任何命令，应用非常广泛。
*/
//go:generate command arg1 arg2
```

> 注意事项：

- 该特殊注释必须在 .go 源码文件中。
- 每个源码文件可以包含多个 generate 特殊注释时。
- 显示运行 go generate 命令时，才会执行特殊注释后面的命令。
- 命令串行执行的，如果出错，就终止后面的执行。
- 特殊注释必须以"//go:generate"开头，双斜线后面没有空格。

### 8.go test命令

> 测试命令	go test 命令，会自动读取源码目录下的 *_test.go 文件，生成并运行测试用的可执行文件。

```go
主要提供“单元测试”和“基准测试”两种方案

单元测试——测试和验证代码的框架
– 1) 单元测试命令行
– 2) 运行指定单元测试用例
– 3) 标记单元测试结果
– 4) 单元测试日志

基准测试——获得代码内存占用和运行效率的性能数据
– 1) 基础测试基本使用
– 2) 基准测试原理
– 3) 自定义测试时间
– 4) 测试内存
– 5) 控制计时器
```

### 9.go pprof命令

> 性能分析命令	go pprof 可以帮助开发者快速分析及定位各种性能问题，如 CPU 消耗、内存分配及阻塞分析。

```go
性能分析首先需要使用 runtime.pprof 包嵌入到待分析程序的入口和结束处。runtime.pprof 包在运行时对程序进行每秒 100 次的采样，最少采样 1 秒。然后将生成的数据输出，让开发者写入文件或者其他媒介上进行分析。

go pprof 工具链配合 Graphviz 图形化工具可以将 runtime.pprof 包生成的数据转换为 PDF 格式，以图片的方式展示程序的性能分析结果。
```

```go
//安装第三方图形化显式分析数据工具（Graphviz）
yum install graphiviz
//安装第三方性能分析来分析代码包
go get github.com/pkg/profile
```

## GO 语言基础

### 1.格式化输出

- Println 	换行输出
- Printf       格式化输出
- Fprintf     格式化并输出
- Sprintf     格式化并返回一个字 符串而不带任何输出。

```go
reflect.TypeOf(v1)			//判断变量类型
fmt.Fprintf(os.Stderr, "格式化 %s\n", "error")
```

| 参数             | 格式                                                |
| ---------------- | --------------------------------------------------- |
| %d               | 十进制整数                                          |
| %x     %o     %b | 十六进制,八进制,二进制整数                          |
| %f      %g    %e | 浮点数:     3.141593 3.141592653589793 3.141593e+00 |
| %t               | 布尔：true或false                                   |
| %c               | 字符（rune） (Unicode码点)                          |
| %s               | 字符串                                              |
| %q               | 带双引号的字符串"abc"或带单引号的字符'c'            |
| %v               | 变量的自然形式（natural format）                    |
| %+v              | 在 %v的基础上额外输出字段名                         |
| %#v              | 在%+v的基础上额外输出类型名                         |
| %T               | 变量的类型                                          |
| %%               | 字面上的百分号标志（无操作数）                      |
| %U               | Unicode 格式 : U+1234; 与"U + %04x"相同             |
| %2.2f            | 输出时的间隔 2  格式化%f 精度2                      |
| %p               | 格式化指针的标识符                                  |

### 2.常量

> 常量使用关键字 `const` 定义，用于存储不会改变的数据。

```go
显式类型定义： const b string = "abc"
隐式类型定义： const b = "abc"
```

### 3.变量

- 变量声明

```Go

var 变量名字 类型 = 表达式
其中“类型”或“= 表达式”两个部分可以省略其中的一个。如果省略的是类型信息，那么将根据初始化表达式来推导变量的类型信息。如果初始化表达式被省略，那么将用零值初始化该变量。
//如果初始化表达式被省略，那么将用零值初始化该变量。
数值类型变量对应的零值是0，布尔类型变量对应的零值是false，字符串类型对应的零值是空字符串，接口或引用类型（包括slice、指针、map、chan和函数）变量对应的零值是nil。数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值。
var i, j, k int                 // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string
```

- 简短变量声明

  简短变量声明语句的形式可用于声明和初始化局部变量。	“名字 := 表达式”形式

  ```Go
  //简短变量声明
  anim := gif.GIF{LoopCount: nframes}
  freq := rand.Float64() * 3.0
  t := 0.0
  //简短变量声明一组变量
  i, j := 0, 1
  /**
  *第一个语句声明了in和err两个变量。
  *在第二个语句只声明了out一个变量，然后对已经声明的err进行了赋值操作。
  */
  
  in, err := os.Open(infile)
  out, err := os.Create(outfile)
  ```

  第一个语句声明了in和err两个变量。在第二个语句只声明了out一个变量，然后对已经声明的err进行了赋值操作。

  ```Go
  in, err := os.Open(infile)
  
  out, err := os.Create(outfile)
  ```

  简短变量声明语句中必须至少要声明一个新的变量，下面的代码将不能编译通过：

  ```Go
  f, err := os.Open(infile)
  //解决的方法是第二个简短变量声明语句改用普通的多重赋值语句。
  f, err := os.Create(outfile) // compile error: no new variables
  ```

### 4.指针

- "var x int"声明语句声明一个x变量

- &x表达式（取x变量的内存地址）将产生一个指向该整数变量的指针，

  指针对应的数据类型是`*int`，指针被称之为“指向int类型的指针”

- 如果指针名字为p，那么可以说“p指针指向变量x”，或者说“p指针保存了x变量的内存地址”。

   同时`*p`表达式对应p指针指向的变量的值

```go
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"
```

- 变量有时候被称为可寻址的值。即使变量由表达式临时生成，那么表达式也必须能接受`&`取地址操作。
- 任何类型的指针的零值都是nil。如果p指向某个有效变量，那么`p != nil`测试为真。指针之间也是可以进行相等测试的，只有当它们指向同一个变量或全部是nil时才相等。

```Go
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
```

### 5.strings常用方法

|                   方法名                    |                参数                 |                  作用                   |
| :-----------------------------------------: | :---------------------------------: | :-------------------------------------: |
|   strings.HasSuffix(s, suffix string)bool   |  s : 目标字符串  prefix : 前缀字符  |   判断字符串 `s` 是否以 `prefix` 开头   |
|   strings.HasPrefix(s, prefix string)bool   |  s : 目标字符串  prefix : 结尾字符  |   判断字符串 `s` 是否以 `prefix` 结尾   |
|   strings.Contains(s, substr string)bool    |  s : 目标字符串  substr : 子字符串  |      判断字符串 `s` 是否包含substr      |
|       strings.Index(s, str string)int       |   s : 目标字符串  str : 子字符串    |       `str` 在字符串 `s` 中的索引       |
|     strings.LastIndex(s, str string)int     |   s : 目标字符串  str : 子字符串    | `str` 在字符串 `s` 中最后出现位置的索引 |
|   strings.IndexRune(s string, r rune)int    |      s : 目标字符串  r : 字节       |  查询非 ASCII 编码的字符r在str中的位置  |
|   strings.Replace(str, old,new, n)string    | (string){str\|old\|new}\|(index){n} | 将str中的前 n个字符 `old` 替换为 `new`  |
|       strings.Count(s, str string)int       |         s:原字符串 str 子串         |   计算 `str` 在`s` 中出现的非重叠次数   |
|     strings.Repeat(s, count int)string      |      s:原字符串 count重复次数       |   返回重复 `count` 次字符串 `s`的结果   |
|          strings.ToLower(s)string           |              s原字符串              |            字符转为小写字符             |
|          strings.ToUpper(s)string           |              s原字符串              |            字符转为大写字符             |
|            strings.TrimSpace(s)             |              s原字符串              |       剔除字符串开头和结尾的空白        |
|             strings.TrimLeft(s)             |              s原字符串              |          剔除字符串开头的空白           |
|            strings.TrimRight(s)             |              s原字符串              |          剔除字符串结尾的空白           |
|              strings.Fields(s)              |              s原字符串              |    以(1,)个空白字符空分割目标字符串     |
|            strings.Split(s, sep)            |        s原字符串 sep分割字符        |           用sep分割目标字符串           |
| strings.Join(sl []string, sep string)string |      sl 切片数组 sep 分割字符       |       将sl以sepl连接成一个字符串        |
|           strings.NewReader(str)            |   read := strings.NewReader(str)    |      生成reader并读取字符串中内容       |
|              Package : strconv              |              类型转换               |            字符串转换工具类             |

### 6.时间和日期

>时间常用方法

| 方法名                                                | 参数               |             作用             |
| ----------------------------------------------------- | ------------------ | :--------------------------: |
| func Now() Time                                       | 无                 |         获取当前时间         |
| time.Now().Year()\|.Month()\|.Day()                   | time对象           |     获取时间对象的年月日     |
| func Parse(layout, value string) (Time, error)        | 字符串格式化成时间 |   layout必须时Go的诞生时间   |
| func (t Time) Format(layout string) string            | 时间转为字符串     |   根据参考时间转换成字符串   |
| time.Now().Unix                                       | 无                 |      获取当前时间时间戳      |
| func Unix(sec int64, nsec int64) Time                 | 秒数和纳秒数       |      将时间戳转换为时间      |
| func (t Time) Equal(u Time) bool 或 = =               | time t  time u     | 比较两个时间是否相同(含时区) |
| func (t Time) Before(u Time) bool                     | time t  time u     |      时间t 是否在u之前       |
| func (t Time) After(u Time) bool                      | time t  time u     |      时间t 是否在u之后       |
| func (t Time) Date() (year int, month Month, day int) | time               |     返回时间 t 的年月日      |
| func (t Time) Clock() (hour, min, sec int)            | time               |     返回时间 t 的时分秒      |

>时间运算

| 方法名                                                      | 参数  |          作用           |
| ----------------------------------------------------------- | ----- | :---------------------: |
| func (t Time) Add(d Duration) Time                          | time  |    返回t+d后的时间点    |
| func (t Time) AddDate(years int, months int, days int) Time | Y-M-D |   返回t+YMD后的时间点   |
| func (t Time) Sub(u Time) Duration                          | t  u  | 返回t-u的时间段(h-m-s)  |
| func (t Time) Round(d Duration) Time                        | t  d  |   返回距离t最近的时间   |
| func (t Time) Truncate(d Duration) Time                     | t  d  | 返回最接近但早于t的时间 |

>Timer定时器

| 方法名                                      | 参数 | 作用                                               |
| ------------------------------------------- | ---- | -------------------------------------------------- |
| func NewTimer(d Duration) *Timer            | d    | 在最少过去时间段d后到期，自己的C通道发送当时的时间 |
| func After(d Duration) <-chan Time          | d    | 在另一线程经过时间段d后向返回值发送当前时间        |
| func AfterFunc(d Duration, f func()) *Timer | d  f | 等待时间d然后调用函数f                             |
| func Sleep(d Duration)                      | d    | Sleep阻塞当前go程至少d代表的时间段                 |
| func Tick(d Duration) <-chan Time           | d    | 只提供对Ticker的通道的访问,无法关闭                |



### 7.控制结构

> if-else 结构

- if 和 else之后的 { 必须和关键字同行
- f 结构内有 break、continue、goto 或者 return 语句时,省略 else
- 使用简短方式 := 声明的变量的作用域只存在于 if 结构中 (如果在if之前被声明,那原有的值会被隐藏)

>Switch 结构

- { 必须和 switch 关键字在同一行 
- case 多个符合条件的值 case val1, val2, val3 :
- Go 执行完分支就会退出Switch代码块,所以不需要使用 ==break== 表示结束
- 进入分支后继续执行下方分支时使用 ==fallthrough== 继续执行后续分支



> for 结构

1. for 基本语法

```go
for 初始化语句; 条件语句; 修饰语句 {}
```

2. for 循环的三种形式

```go
1.完整形态
for (i = 0; i < 10; i++) { }
2.基于条件判断的迭代
for i >= 0 { i = i - 1 }
3.无限循环
for { }
```

3. for-range 结构

```go
for ix, val := range coll { }
```



> Break 与 continue

==Break的作用效果==

- break 的作用范围为该语句出现后的最内部的结构 
- 在 switch 或 select 语句中，break 作用结果是跳过整个代码块，执行后续的代码
- 嵌套的循环体, break 只会退出最内层的循环：

==Continue的作用效果==

- continue 忽略剩余的循环体而直接进入下一次循环的过程，但不是无条件执行下一次循环，执行之前依旧需要满足循环的判断条件

> 标签 与 goto

==标签 (Label)==

- for、switch 或 select 语句都可以配合标签（label）形式的标识符使用，即某一行第一个以冒号（`:`）结尾的单词
- 标签的名称是大小写敏感的，为了提升可读性，一般建议使用全部大写字母

==go to==

- 指向到对应标签继续运行代码 (不会干扰到已经存在的变量值)




### 8.函数 Function

#### 1.函数参数与返回值

> **niladic** 函数（niladic function）

```go
//函数定义时，它的形参一般是有名字的，也可以定义没有形参名的函数，只有相应的形参类型
func f(int, int, float64)
```

#### 2.按值传递（call by value） 按引用传递（call by reference）

==按值传递== : Go 默认使用按值传递来传递参数，也就是传递参数的副本

==按引用传递== : 如果需要函数直接修改参数的值,要将参数的地址（比如 &variable）传递给函数

- 在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递

- 如果一个函数需要返回多个值，可以传递一个==切片==给函数（如果返回值具有相同类型）或者是传递一个==结构体==（如果返回值具有不同的类型）

#### 3.命名的返回值 (named return variables)

> 命名的返回值作为结果形参,需要返回时 ==直接return== 即可

```go
//没有命名的返回值 函数
func namedFunc(input int) (int, int) {
    return 2 * input, 3 * input
}
//命名的返回值函数
func getX2AndX3_2(input int) (x2 int, x3 int) {
    x2 = 2 * input
    x3 = 3 * input
    // return x2, x3
    return
}
```

- return 或 return var 都是可以的。
- return var = expression（表达式） 会引发一个编译错误：syntax error: unexpected =, expecting semicolon or newline or }

#### 4.空白符（blank identifier）

> 空白符用来匹配一些不需要的值,然后丢弃

```go
func ThreeValues() (int, int, float32) {
    return 5, 6, 7.5
}
func main(){
    //函数有三个返回值,第二个使用空白符接收后即被自动丢弃
    i1, _, f1 = ThreeValues()
}
```

#### 5.改变外部变量（outside variable）

- 传递指针给函数可以节省内存,而且赋予函数直接修改外部变量的能力
- 修改变量不在需要使用return返回
- `reply` 是一个指向 `int` 变量的指针，通过这个指针，我们在函数内修改了这个 `int` 变量的数值

```go
func Multiply(a, b int, reply *int) {
    *reply = a * b
}
//在方法中使用指针直接修改外地传递的变量的原值
func main() {
    n := 0
    reply := &n
    Multiply(10, 5, reply)
    fmt.Println("Multiply:", *reply) // Multiply: 50
}
```

#### 6.传递可变长度参数

- 可变参数形式 ==...type== , 将切片拆解为多个参数 ==slice...==

```go
func F1(s ...string) { }
```

- 不同类型的可变参数(==存储所有可能的参数的结构体==)

```go
//包含所有可能的参数的结构体  
type Options struct {
     par1 type1,
     par2 type2,
     ...
 }
//使用方法时第三个参数为结构体
func F1(a, b, Options {}) { }
```

- 使用默认的空接口 (interface{ })

```go
// F1(a, b, Options {par1:val1, par2:val2})
func typecheck(..,..,values …interface{}){
    for _, value := range values {
        switch v := value.(type){
            caseint:…
            casefloat:…
            casestring:…
            casebool:…
            default:…
        }
    }
}
```



#### 7. defer 代码追踪

- 关键字 ==defer== 允许推迟到 ==函数返回之前== 或 任意位置执行 ==return语句== 之后
- 一般用于释放某些已分配的资源

> defer记录函数的参数与返回值

```go
package main
import (
    "io"
    "log"
)
func func1(s string) (n int, err error) {
    defer func() {
        log.Printf("func1(%q) = %d, %v", s, n, err)
    }()
    return 7, io.EOF
}
func main() {
    func1("Go")
}

输出结果 : Output: 2011/10/04 10:46:11 func1("Go") = 7, EOF
```



#### 8.内置函数

Go 语言拥有一些不需要导入就可以使用的内置函数

| 称                 | 说明                                                         |
| :----------------- | :----------------------------------------------------------- |
| close              | 用于管道通信                                                 |
| len、cap           | len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map） |
| new、make          | new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：`v := new(int)`。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作（详见第 7.2.3/4 节、第 8.1.1 节和第 14.2.1 节）**new() 是一个函数，不要忘记它的括号** |
| copy、append       | 用于复制和连接切片                                           |
| panic、recover     | 两者均用于错误处理机制                                       |
| print、println     | 底层打印函数 ，在部署环境中建议使用 fmt 包                   |
| complex、real imag | 用于创建和操作==复数==                                       |

#### 9.递归函数

> 在函数体内 调用自身

==栈溢出==  : 大量的递归调用导致的程序栈内存分配耗尽

#### 10. 将函数作为参数

```go
func fun01 (func(r rune) bool,s string) string{ }
func fun02 (r rune) bool {
    
}

func main(){
    str := fun01(fun02('沐'),"这是一段示例文本!")
    fmt.Println(str)
}
```

#### 11. 闭包函数

- 关键字 `defer` （详见第 6.4 节）经常配合匿名函数使用，它可以用于改变函数的命名返回值。
- 匿名函数还可以配合 `go` 关键字来作为 goroutine 使用

```go
//匿名函数赋值给变量
f := func (i int) { }
f()
//匿名函数自调用
func (i int) { }(10)
```

#### 12.应用闭包

- 应用闭包：将函数作为返回值
- 闭包中使用到的变量可以是在闭包函数体内声明的，也可以是在外部函数声明
- 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。

```go
//给未添加文件后缀的文件名添加后缀
func AutoAddSuffix(suffix string) func(string) string{
	return func(name string) string {
		if !strings.HasSuffix(name,suffix) {
			return name + suffix
		}
		return name
	}	

}
//斐波那契数列 闭包实现
func fibFunc() func () int{
	 num1,num2 := 0,1
	 return func () int {
	 	num1,num2 = num2,num1+num2
	 	return num2
	 }
}

func main (){
    addGo := AutoAddSuffix(".go")
	addHtml := AutoAddSuffix(".html")
	fmt.Printf("%s",addGo("goLang"))
	fmt.Printf("%s",addHtml("index"))
    
    f := fibFunc()
	arr := make([]int,0)
	for i := 0; i < 100; i++ {
		arr = append(arr,f())
	}
	fmt.Printf("%v\n",arr)
	fmt.Printf("切片长度为%d,容量为%d\n",len(arr),cap(arr))
}
```

#### 13.闭包调试

> 获取正在执行的文件中的函数 : ==runtime== 中的 ==Caller()== 方法  或  ==log==

- 使用where() 闭包函数和 runtime中的Caller()

```go
where := func() {
    _, file, line, _ := runtime.Caller(1)
    log.Printf("%s:%d", file, line)
}
where()
// some code
where()
// some more code
where()
```

- 使用log包中的 flag参数实现

```go
log.SetFlags(log.Llongfile)
log.Print("")

// 简短实现
var where = log.Print
func func1() {
where()
... some code
where()
... some code
where()
}
```

#### 14.计算函数执行时间

计算函数执行开始到完成所消耗的时间

```go
start := time.Now()
longCalculation()
end := time.Now()
delta := end.Sub(start)
fmt.Printf("longCalculation took this amount of time: %s\n", delta)
```

#### 15.通过内存缓存提升性能

- 进行大量的计算时 ==提升性能== 最有效的方式就是 避免重复计算. 

  通过==在内存中缓存==和==重复利用相同计算的结果==,称之为缓存. 

- 内存缓存在使用==计算成本相对昂贵的函数==时非常有用,譬如 大量进行相同参数的运算,

  还可用于输入想通透输出必定获得相同输出的纯函数中.



### 9.切片与数组

容器: 包含大量条目(item) 的数据结构,例如 ==数组 []==,==切片(slice)== 和 ==map== .

#### 1.数组的声明和初始化

- 数组是 ==值类型==
- 数组时具有==唯一类型== 的 一组已编号且长度固定的数据序列
- 类型可以是任意的原始类型: 整数,字符或自定义类型
- 数组长度必须是一个常量表达式或非负整数
- 数组长度也是数组类型的一部分,即不同长度的数组属于不同类型
- 数组声明时,其中的所有元素都会自动初始化为默认值
- 数组赋值可通过 `arr[i] = value`
- 如果数组索引越界时 编译会通过而运行时会 panic

==声明格式==

```go
var indetifier [len]type
例 : var arr [8]int

//初始化
var arr1 = [5]int{0,1,2,3,4}
var arr2 = [...]int{5,6,7,8,9}
var slice1 = []int{10,11,12,13,14}
//部分值初始化
var arr3 = [5]string{3: "arr3", 4 : "arr4"}
var slice2 = []string{3: "slice3", 4 : "slice4"}
```

#### 2.数组常量

- 如果数组值已知,则可以通过 ==数组常量== 的方法来初始化数组
- 数组常量赋值的三种形式

```go
//形式一 从左往右依次赋值
var arr = [5]int{1,2,3}
//形式二 不定长度数组
var arr = [...]int{1,3,5,7,9}
//形式三  key:value 形式
var arrValue = [5]string{3: "arr3",5: "arr5"}
```

#### 3.多为数组

- 数组通常是一维的,但是可以用来组装成多为数组, 例如: `[3][5]int`

#### 4.将数组传递给函数

- 使用数组的指针
- 使用数组的切片

#### 5.切片

- 切片（slice）是对数组一个==连续片段的引用==  （该数组我们称之为相关数组，通常是匿名的）
- 切片是  ==引用数据类型==
- 切片是可以索引的, 并且可以由 `len()` 函数获取长度
- 切片的长度可以在运行时修改,是一个 ==长度可变的数组==
- 切片提供了计算容量的函数 `cap(s)`  cap(s) = s[0] 到 s[len(s)-1]的长度
- 切片最长可达 = 切片长度 + 数组除切片之外的长度  所以 ==0<=len(s)<= cap(s)== 恒成立
- 多个切片表示同一个数组的片段时,可以==共享数据==;数组实际上是切片的构建块.
- 切片是引用,所以不需要使用额外的内存,比数组更有效率
- 切片扩展到大小上限 s = s[:cap(s)]
- 字符是不可变的字节数组,所以也可以被切分成切片
- 切片追加元素 ==append(slice, element)==

==切片声明==

```go
//切片初始化之前默认为nil
var identifier []type 	(无需指明长度)
//切片的初始化格式 (通过数组进行初始化)
var slice []type = arr[star:end]
//使用类似数组的方式初始化
s := [5]int{2, 3, 5, 7, 11}[:]
var s = []int{2, 3, 5, 7, 11}
```

 ==切片截取部分数组==

```go
//声明一个容量为end的数组
var arr =  [end]int{1,3,5,7,9,11,13,15}
//截取数组进行初始化的格式 [start,end)
var slice1 []type = arr[start:end]
//得到完整的数组
var slice2 []type = arr[:] 
//得到数组从start到末尾的部分
var slice3 []type = arr[start:]
//得到数组从开始到end的部分
var slice4 []type = arr[:end]
```

==切片传递给函数==

```go
func sum (a []int) int { }
func main() {
    var arr = [5]int{0,1,2,3,4}
    sum(arr[:])
}
```

==make()创建一个切片==

```go
var slice []typr = make([]type,len)
简写 : slice := make([]type,len)
//创建一个只占用数组 len个数项的切片
slice := make([]type,len,cap)
//make 使用的基本方式
func make([]T,len,cap){ }
```

==new() 和 make()== 的区别

- new(T) 为T分配一片内存,`初始化为0`,返回 `*T的内存地址` 

  返回值时一个指向类型为T 值为0的地址的指针,适用于 值类型 (数组和结构体)

- make(T) 返回一个类型为T的初始值,只适用于3中内建的引用类型 : `切片`| `map` | `channel`



==多维切片==

切片可以由一维组合成高维,通过分片的分片(切片的数组),长度可以任意动态变化

Go语言的多维切片可以任意切分,但内层的切片必须单独分配 (==通过make函数==)

#### 6. bytes

- 类型[]byte 的切片十分常见,Go语言有==bytes==包解决相关问题
- bytes包和字符串包类似, 而且包含 ==Buffer== 类型

长度可变的 bytes 的 buffer

```go
import"bytes"
type Bufferstruct{
...
}
```

读写 未知长度的 bytes 最好使用 buffer

```go
//Buffer定义
var buffer bytes.Buffer
//new() 获得Buffer指针
var r *bytes.Buffer = new(bytes.Buffer)
//通过函数创建Buffer对象 并且用buf初始化
func NewBuffer(buf []byte) *Buffer
```

通过 buffer 串联字符串 (比使用 `+=` 要更节省内存和 CPU)

```go
/**
* 通过 buffer.WriteString(s) 方法 将字符串s 追加到末尾
* 最后通过buffer.String() 方法将buffer转换为string
*/
var buffer bytes.Buffer
for{
    if s, ok := getNextString(); ok {//method getNextString() not shown here
         buffer.WriteString(s)
    }else{
   		 break
    }
}
fmt.Print(buffer.String(),"\n")
```

#### 7. 切片重组

```go
//初始化一个数组
slice := make([]type,start_length,capacity)
//切片的初始长度  
start_length
//切片的初始容量
capacity
//切片达到容量上限后进行扩容 (切片重组)
slice1 = slice[0:end] //end为新的末尾索引
sl = sl[0,len(sl)+1]
```

#### 8.切片的复制与追加

- slice... 切片解构
- 增加切片容量 : 创建新的大容量切片  复制原切片内容到新切片中
- `func append(s []T,x ...T) []T`  将n个相同类型的元素追加到切片后面,并返回新切片
- `func copy(dst,src []T) int`  将切片src拷贝到dst并覆盖dst的相关元素

#### 9.字符串生成切片

- c := []











