
## go概述

### go适合干什么
1. 区块链
2. 服务器端 / 游戏后台
3. 云计算 / 云服务


### go的历史
现有的语言系统没有充分利用多核，处理大并发的能力不够好

c/c++运行速度虽然很快，但是编译速度很慢，还存在内存泄漏等问题。


### go吉祥物
金花鼠(gordon)

### go特点
既能达到静态语言的安全和性能，又达到了动态语言的开发维护效率。Go=C+python

1. 引入包的概念，每个文件都要归属于一个package
2. 垃圾回收机制，自动回收内存。
3. 天然具有并发的能力
4. 管道机制，通过chan进行goroutine的通信
5. 切片，延迟执行(defer)
6. 函数可以返回多个值




## 安装

下载 -- 添加环境变量 -- cmd查看版本 -- 配置GOPATH环境变量 -- 建三个文件夹



查看go的版本`go version`

查看go安装的包`go list`

配置go国内的源

```shell
C:\Users\leishen>go env -w GO111MODULE=on
C:\Users\leishen>go env -w GOPROXY=https://goproxy.cn,direct
```








gin.SetMode(gin.ReleaseMode) 发行模式







安装gin`go get -u github.com/gin-gonic/gin`





在进行`Go`语言开发的时候，我们的代码总是会保存在`$GOPATH/src`目录下。在工程经过`go build`、`go install`或`go get`等指令后，会将下载的第三方包源代码文件放在`$GOPATH/src`目录下， 产生的二进制可执行文件放在 `$GOPATH/bin`目录下，生成的中间缓存文件会被保存在 `$GOPATH/pkg` 下。







## 特性

### 变量必须声明并且声明后必须使用





### 下划线的使用

go中不使用的变量会报错，因此需要有”虚拟的“变量去接收返回值，这时候就用下划线



补充：

```
    import "database/sql"
    import _ "github.com/go-sql-driver/mysql"
```

第二个import就是不直接使用mysql包，只是执行一下这个包的init函数，把mysql的驱动注册到sql包里，然后程序里就可以使用sql包来访问mysql数据库了。

即：import 加下划线 只执行了init，不能去调用这个包的其他功能



### iota常量计数器





### 1.1.7. 多行字符串

Go语言中要定义一个多行字符串时，就必须使用`反引号`字符：

```
    s1 := `第一行
    第二行
    第三行
    `
    fmt.Println(s1)
```

反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。



### 1.1.10. 修改字符串

要修改字符串，需要先将其转换成`[]rune或[]byte`，完成后再转换为`string`。无论哪种转换，都会重新分配内存，并复制字节数组。

```
    func changeString() {
        s1 := "hello"
        // 强制类型转换
        byteS1 := []byte(s1)
        byteS1[0] = 'H'
        fmt.Println(string(byteS1))

        s2 := "博客"
        runeS2 := []rune(s2)
        runeS2[0] = '狗'
        fmt.Println(string(runeS2))
    }
```





## 数组

一旦定义，长度不能变

长度是类型的一部分，var a[5] int和var a[10]int是不同的类型

指针数组 [n]*T，数组指针 *[n]T



### 切片

相当于数组的引用

第一次append的时候，长度没有超过容量，所以容量没有变。
第二次append的时候，**长度超过了容量，这时容量会扩展到原来的2倍**。(但是如果一次添加多个值，这个容量不是这样扩展的)

```
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4, 5, 8, 10, 234, 123, 18)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

```

同时，增加后的切片地址和原来不同，也就是说：
append操作可能会导致原本使用同一个底层数组的两个Slice变量变为使用不同的底层数组。





**你可能会问一个问题：如果切片是建立在数组之上的，而数组本身不能改变长度，那么切片是如何动态改变长度的呢？实际发生的情况是，当新元素通过调用 `append` 函数追加到切片末尾时，如果超出了容量，`append` 内部会创建一个新的数组。并将原有数组的元素被拷贝给这个新的数组，最后返回建立在这个新数组上的切片。这个新切片的容量是旧切片的二倍**

cap是倍增的，2的次方



**应及时将所需数据 copy 到较小的 slice，以便释放超大号底层数组内存。**





### 切片与数组的内存优化问题

# 内存优化

切片保留对底层数组的引用。只要切片存在于内存中，数组就不能被垃圾回收。这在内存管理方便可能是值得关注的。假设我们有一个非常大的数组，而我们只需要处理它的一小部分，为此我们创建这个数组的一个切片，并处理这个切片。这里要注意的事情是，数组仍然存在于内存中，因为切片正在引用它。

解决该问题的一个方法是使用 [copy](https://golang.org/pkg/builtin/#copy) 函数 `func copy(dst, src []T) int` 来创建该切片的一个拷贝。这样我们就可以使用这个新的切片，原来的数组可以被垃圾回收。

```
package main

import (  
    "fmt"
)

func countries() []string {  
    countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2]
    countriesCpy := make([]string, len(neededCountries))
    copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    return countriesCpy
}
func main() {  
    countriesNeeded := countries()
    fmt.Println(countriesNeeded)
}
```

在上面[程序](https://play.golang.org/p/35ayYBhcDE)中，第 9 行 `neededCountries := countries[:len(countries)-2]` 创建一个底层数组为 `countries` 并排除最后两个元素的切片。第 11 行将 `neededCountries` 拷贝到 `countriesCpy` 并在下一行返回 `countriesCpy`。现在数组`countries` 可以被垃圾回收，因为 `neededCountries` 不再被引用。





### go特性



go中没有while 单独的for就是while



defer



&i是取i的地址

*p是根据p取对应的值



---

map原理

slice原理







导入包的命名





### func

注意传递的是指针还是值，多说情况下传递的应该是指针。





### 结构体

> go中没有类，需要使用结构体进行类的操作

定义类方法就是把这个结构体作为函数定义的参数





## 接口

go中可以通过interface定义一种具有若干功能的接口。

interface是一种类型。通过type定义

```go
type Abser interface {
	Abs() float64
	// 定义了一个Abser接口，这个接口类型的变量需要实现Abs()功能
}
```







### import中的下划线

https://www.cnblogs.com/hezhixiong/p/4588964.html

总结：下划线只导入了这个包的init，无法通过包名来调用包中的其他函数。



### 导入的目录问题

在导入的时候，不用导入那个对应的文件，导入包即可，比如：在utils目录下有很多文件，其中一个wxpay.go中有一个func MakeMiniProgramUnifiedOrder

```
import (
	"baogo/utils"
)

// 这时导入utils即可，不用导入utils/wxpay
// 可以直接使用这个MakeMiniProgramUnifiedOrder
```





### go项目下生成的go.sum和go.mod

在`go mod init package`的时候生成这两个文件。

go.mod是记录了依赖包和版本

go.sum是将下载的依赖包和版本进行