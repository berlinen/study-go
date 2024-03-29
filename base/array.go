// array

// array就是数组，它的定义方式如下：

var arr [n]type

// 在[n]type中，n表示数组的长度，type表示存储元素的类型。对数组的操作和其它语言类似，都是通过[]来进行读取或赋值：

var arr [10]int  // 声明了一个int类型的数组
arr[0] = 42 //  // 数组下标是从0开始的
arr[1] = 14 // 赋值操作
fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回42
fmt.Printf("The last element is %d\n", arr[9]) //返回未赋值的最后一个元素，默认返回0

// 由于长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度。
// 数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，
// 而不是它的指针。如果要使用指针，那么就需要用到后面介绍的slice类型了。

// 数组可以使用另一种:=来声明

a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组

b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0

c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

// 也许你会说，我想数组里面的值还是数组，能实现吗？当然咯，Go支持嵌套数组，即多维数组。比如下面的代码就声明了一个二维数组：

// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

// 上面的声明可以简化，直接忽略内部的类型
easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

// --------------------------------------------------------------------------------------------------------------------------------------------

// slice

// 在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫slice

// slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。

// 和声明array一样，只是少了长度

var fslice []int

// 接下来我们可以声明一个slice，并初始化数据，如下所示：

slice := []byte {'a', 'b', 'c'}

// slice可以从一个数组或一个已经存在的slice中再次声明。slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。

// 声明一个含有10个元素元素类型为byte的数组
var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

// 声明两个含有byte的slice
var a, b []byte

// a指向数组的第3个元素开始，并到第五个元素结束，
a = ar[2:5]
// 现在a含有的元素: ar[2]、ar[3]和ar[4]

// b是数组ar的另一个slice
b = ar[3:5]
// b的元素是：ar[3]和ar[4

// 注意slice和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用...自动计算长度，而声明slice时，方括号内没有任何字符。

// slice有一些简便的操作

/*
* slice的默认开始位置是0，ar[:n]等价于ar[0:n]
* slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
* 如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]
*/

// 下面这个例子展示了更多关于slice的操作：

// 声明一个数组
var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
// 声明两个slice
var aslice, bslice []byte

// 演示一些简便操作
aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

// 从slice中获取slice
aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g

// slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值，例如上面的aSlice和bSlice，如果修改了aSlice中元素的值，那么bSlice相对应的值也会改变。

// 从概念上面来说slice像一个结构体，这个结构体包含了三个元素：

/*
* 一个指针，指向数组中slice指定的开始位置
* 长度，即slice的长度
* 最大长度，也就是slice开始位置到数组的最后位置的长度
*/

Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
Slice_a := Array_a[2:5]

// 对于slice有几个有用的内置函数：

/*
* len 获取slice的长度
* cap 获取slice的最大容量
* append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
* copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
*/

// 注：append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
// 但当slice中没有剩余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间。返回的slice数组指针将指向这个空间，
// 而原数组的内容将保持不变；其它引用此数组的slice则不受影响。

// 从Go1.2开始slice支持了三个参数的slice，之前我们一直采用这种方式在slice或者array基础上来获取一个slice

var array [10]int
slice := array[2:4]  // 这个例子里面slice的容量是8，新版本里面可以指定这个容量

slice = array[2:4:7] // 上面这个的容量就是7-2，即5。这样这个产生的新的slice就没办法访问最后的三个元素。
// 如果slice是这样的形式array[:i:j]，即第一个参数为空，默认值就是0。

// ---------------------------------------------------------------------------------------------------------------------------

// map

// map也就是Python中字典的概念，它的格式为map[keyType]valueType

// 我们看下面的代码，map的读取和设置也类似slice一样，通过key来操作，只是slice的index只能是｀int｀类型，而map多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型。

// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
var numbers map[string]int

// 另一种map的声明方式
numbers = make(map[string]int)
numbers["one"] = 1  //赋值
numbers["ten"] = 10 //赋值
numbers["three"] = 3

fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
// 打印出来如:第三个数字是: 3

// 这个map就像我们平常看到的表格一样，左边列是key，右边列是值

//使用map过程中需要注意的几点：

/*
	map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
	map的长度是不固定的，也就是和slice一样，也是一种引用类型
	内置的len函数同样适用于map，返回map拥有的key的数量
	map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
	map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
	map的初始化可以通过key:val的方式初始化值，同时map内置有判断是否存在key的方式
*/

// 通过delete删除map的元素：

// 初始化一个字典
rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
csharpRating, ok := rating["C#"]
if ok {
	fmt.Println("C# is in the map and its rating is ", csharpRating)
} else {
	fmt.Println("We have no rating associated with C# in the map")
}

delete(rating, "C")  // 删除key为C的元素

// 上面说过了，map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变：

m := make(map[string]string)
m["Hello"] = "Bonjour"
m1 := m
m1["Hello"] = "Salut"  // 现在m["hello"]的值已经是Salut了

// make、new操作

// make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。

// 内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要

// new返回指针

// 内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。

// make返回初始化后的（非零）值。

// -------------------------------------------------------------------------------------------------------------------------------------------------------------

// 零值

// 关于“零值”，所指并非是空值，而是一种“变量未填充前”的默认值，通常为0。 此处罗列 部分类型 的 “零值”

int     0
int8    0
int32   0
int64   0
uint    0x0
rune    0 //rune的实际类型是 int32
byte    0x0 // byte的实际类型是 uint8
float32 0 //长度为 4 byte
float64 0 //长度为 8 byte
bool    false
string  ""

