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

