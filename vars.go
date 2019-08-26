// 定义一个名称为“variableName”，类型为"type"的变量
var variableName type

// 定义多个都是type类型变量
var vname1, vname2, vname3 type

// 定义变量并初始化值
var variableName type = value

// 同时初始化多个变量
var vname1, vname2, vname3 type = v1, v2, v3

// 忽略类型声明

/*
	定义三个变量，它们分别初始化为相应的值
	vname1为v1，vname2为v2，vname3为v3
	然后Go会根据其相应值的类型来帮你初始化它们
*/
var vname1, vname2, vname3 = v1, v2, v3

// 继续简化
/*
	定义三个变量，它们分别初始化为相应的值
	vname1为v1，vname2为v2，vname3为v3
	编译器会根据初始化的值自动推导出相应的类型
*/
// 只能使用在函数内部
// 在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量
vname1, vname2, vname3 := v1, v2, v3