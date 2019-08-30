// 内置基础类型

// Boolean

var isActive bool // 全局变量声明
var enable, disabled = true, false // 忽略类型声明
func test () {
	var available bool // 一般声明
	valid := false // 简短声明
  available = true // 赋值操作
}

// 数值类型  Number

// Go里面也有直接定义好位数的类型：rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64。其中rune是int32的别称，byte是uint8的别称。

// 需要注意的一点是，这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错。

// 如下的代码会产生错误：invalid operation: a + b (mismatched types int8 and int32)

var a int8
var b int32
c := a + b

// 浮点数的类型有float32和float64两种（没有float类型），默认是float64

// 复数

var c complex64 = 5+5i
//output: (5+5i)
fmt.Printf("Value is: %v", c)


// 字符串

// Go中的字符串都是采用UTF-8字符集编码。字符串是用一对双引号（""）或反引号（` `）括起来定义，它的类型是string。

var frenchHello string
var emptyString string = ""
func test () {
	no, yes, maybe := "no", "yes", "maybe"
	jappaneseHello := "Konichiwa"
	frenchHello = "Bonjour"
}

// 在Go中字符串是不可变的，例如下面的代码编译时会报错：cannot assign to s[0]

var s string = "hello"
s[0] = 'c'

// 但如果真的想要修改怎么办呢？下面的代码可以实现：

s := 'hello'
c := []byte(s)  // 将字符串 s 转换为 []byte 类型
c[0] = 'c'
s2 := string(c)  // 再转换回 string 类型
fmt.Printf("%s\n", s2)

// Go中可以使用+操作符来连接两个字符串：

s := "hello"
m := "world"
a := s + m
fmt.Printf("%s\n", a)

// 修改字符串也可写为：

s := "hello"
s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
fmt.Printf("%s\n", s)

// 如果要声明一个多行的字符串怎么办？可以通过`来声明：

m := `hello
	world`

// error 类型

// Go内置有一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包errors来处理错误：

err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
	fmt.print(err)
}






