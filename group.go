// 分组声明

// 在Go语言中，同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明。

// 例如下面的代码

import "fmt"
import "os"

const i = 100
const pi = 3.1415
const prefix = "Go_"

var i int
var pi float32
var prefix string

// 可以分组写成如下格式

import(
	"fmt"
	"os"
)

const(
	i = 100
	pi = 3.1415
	prefix  = "Go_"
)

var(
	i int
	pi float32
	prefix string
)

// iota枚举
// Go里面有一个关键字iota，这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1
