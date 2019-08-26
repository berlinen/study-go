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

package main

import (
	"fmt"
)

const (
	x = iota // x== 0
	y = iota // y == 1
	z = iota // z == 2
	w  // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

func main () {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
}

// Go程序设计的一些规则

// Go之所以会那么简洁，是因为它有一些默认的行为：

/*
*  大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
*  大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。
*/