## 1、new
new是一个内置函数，它的函数签名如下：
```go
    func new(Type) *Type
```
解释：
1. Type表示类型，new函数只接受一个参数，这个参数是一个类型
2. *Type表示类型指针，new函数返回一个指向给类型内存地址的指针

new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。举个例子：
```go
func main() {
    a := new(int)
    b := new(bool)
    fmt.Println("%T\n", a) // *int
    fmt.Println("%T\n", b) // *b
    fmt.Println("%v\n", a) // 0
    fmt.Println("%v\n", b) // false
}
```
下方示例代码中```var a *int```只是声明了一个指针变量a，但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。在使用内置的new函数对a进行初始化之后就可以正常对其赋值了：
```go
func main() {
    var a *int
    a = new(int)
    *a = 10
    fmt.Println(*a)
}
```

## 2、make
make也是用于内存分配的，区别与new，它只用于slice，map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型。因为这三种类型就是引用类型，所以就没有返回它们的指针的必要了。make函数签名如下：
```go
func make(t Type, size ...IntegerType) Type
```
make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对他们进行操作。
下方示例代码中```var b map[string]int```只是声明变量b是一个map类型的变量，需要使用make函数进行初始化操作之后，才能泵对其进行赋值：
```go
func main() {
    var b map[string]int
    b = make(map[string]int, 10)    
    b[”测试“] = 100
    fmt.Println(b)
}
```
## 3、new和make的区别
1. 二者都是用来做内存分配的
2. make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身
3. new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针