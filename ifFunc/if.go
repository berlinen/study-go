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