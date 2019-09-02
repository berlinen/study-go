函数

函数是Go里面的核心设计，它通过关键字func来声明，它的格式如下：

func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	//这里是处理逻辑代码
	//返回多个值
	return value1, value2
}

上面的代码我们看出

关键字func用来声明一个函数funcName
函数可以有一个或者多个参数，每个参数后面带有类型，通过,分隔
函数可以返回多个值
上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型
如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
如果没有返回值，那么就直接省略最后的返回信息
如果有返回值， 那么必须在函数的外层添加return语句

下面我们来看一个实际应用函数的例子（用来计算Max值）

package main

import "fmt"

// 返回a、b中最大值.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x := 3
	y := 4
	z := 5

	max_xy := max(x, y) //调用函数max(x, y)
	max_xz := max(x, z) //调用函数max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y,z)) // 也可在这直接调用它
}

上面这个里面我们可以看到max函数有两个参数，它们的类型都是int，那么第一个变量的类型可以省略（即 a,b int,而非 a int, b int)，默认为离它最近的类型，同理多于2个同类型的变量或者返回值。同时我们注意到它的返回值就是一个类型，这个就是省略写法。


多个返回值

Go语言比C更先进的特性，其中一点就是函数能够返回多个值。

我们直接上代码看例子

package main

import "fmt"

//返回 A+B 和 A*B
func SumAndProduct(A, B int) (int, int) {
	return A+B, A*B
}

func main() {
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}
