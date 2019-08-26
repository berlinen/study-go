// 常量

// 所谓常量，也就是在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。在Go程序中，常量可定义为数值、布尔值或字符串等类型。

// 语法如下

const constantName = value
// 如果需要， 也可以明确常量的类型
const Pi float32 = 3.1415926

// 下面是常量的一些例子

const Pi = 3.141502
const i = 10000
const MaxThread = 10
const prefix = "astaxie_"

// Go 常量和一般程序语言不同的是，可以指定相当多的小数位数(例如200位)， 若指定給float32自动缩短为32bit，指定给float64自动缩短为64bit
