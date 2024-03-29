// 流程控制

// 流程控制在编程语言中是最伟大的发明了，因为有了它，你可以通过很简单的流程描述来表达很复杂的逻辑。Go中流程控制分三大类：条件判断，循环控制和无条件跳转。

// if  Go里面if条件判断语句中不需要括号，如下代码所示

if x > 10 {
	fmt.Println("x is greater than 10")
} else {
	fmt.Println("x is less than 10")
}

// Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了，如下所示

//// 计算获取值x,然后根据x返回的大小，判断是否大于10。
if x := computedValue(); x > 10 {
	fmt.Println("x is greater than 10")
} else {
	fmt.Println("x is less than 10")
}

//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
fmt.Println(x)

// 多个条件的时候如下所示：

if integer == 3 {
	fmt.Println("The integer is equal to 3")
} else if integer < 3 {
	fmt.Println("The integer is less than 3")
} else {
	fmt.Println("The integer is greater than 3")
}

// goto

// Go有goto语句——请明智地使用它。用goto跳转到必须在当前函数内定义的标签。例如假设这样一个循环：

func myFunc () {
	i := 0
Here:  //这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	goto Here   //跳转到Here去
}

标签名是大小写敏感的。

Go里面最强大的一个控制逻辑就是for，它既可以用来循环读取数据，又可以当作while来控制逻辑，还能迭代操作。它的语法如下：

for expression1; expression2; expression3 {
	//...
}

expression1、expression2和expression3都是表达式，其中expression1和expression3是变量声明或者函数调用返回值之类的，expression2是用来条件判断，expression1在循环开始之前调用，expression3在每轮循环结束之时调用。

一个例子比上面讲那么多更有用，那么我们看看下面的例子吧：

package main

import "fmt"

func main () {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to", sum)
}

输出：sum is equal to 45

有些时候需要进行多个赋值操作，由于Go里面没有,操作符，那么可以使用平行赋值i, j = i+1, j-1

有些时候如果我们忽略expression1和expression3：

sum := 1
for ; sum < 1000; {
	sum += sum
}

其中;也可以省略，那么就变成如下的代码了，是不是似曾相识？对，这就是while的功能。

sum := 1
for sum < 1000 {
	sum += sum
}

在循环里面有两个关键操作break和continue	,break操作是跳出当前循环，continue是跳过本次循环。当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置，详细参考如下例子：

for index := 10; index > 0; index-- {
	if index == 5{
		break // 或者continue
	}
	fmt.println(index)
}

// break打印出来10、9、8、7、6
// continue打印出来10、9、8、7、6、4、3、2、1

break和continue还可以跟着标号，用来跳到多重循环中的外层循环

for配合range可以用于读取slice和map的数据：

for k,v:=range map {
	fmt.Println("map's key:",k)
	fmt.Println("map's val:",v)
}

由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用_来丢弃不需要的返回值 例如

for _, v := range map{
	fmt.Println("map's val:", v)
}

// --------------------------------------------------------------------------------

switch

有些时候你需要写很多的if-else来实现一些逻辑处理，这个时候代码看上去就很丑很冗长，而且也不易于以后的维护，这个时候switch就能很好的解决这个问题。它的语法如下

switch sExpr {
case expr1:
	some instructions
case expr2:
	some other instructions
case expr3:
	some other instructions
default:
	other code
}

sExpr和expr1、expr2、expr3的类型必须一致。Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；而如果switch没有表达式，它会匹配true

i := 10
switch i {
case 1:
	fmt.Println("i is equal to 1")
case 2, 3, 4:
	fmt.Println("i is equal to 2, 3 or 4")
case 10:
	fmt.Println("i is equal to 10")
default:
	fmt.Println("All I know is that i is an integer")
}

在第5行中，我们把很多值聚合在了一个case里面，同时，Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。

integer := 6
switch integer {
case 4:
	fmt.Println("The integer was <= 4")
	fallthrough
case 5:
	fmt.Println("The integer was <= 5")
	fallthrough
case 6:
	fmt.Println("The integer was <= 6")
	fallthrough
case 7:
	fmt.Println("The integer was <= 7")
	fallthrough
case 8:
	fmt.Println("The integer was <= 8")
	fallthrough
default:
	fmt.Println("default case")
}

上面的程序将输出

The integer was <= 6
The integer was <= 7
The integer was <= 8
default case