<!--
 * @author: Emove
 * @date: Do not edit
 -->
## 内置类型
### 1、值类型
```go
bool
int(32 or 64),int8,int16,int32,int64
uint(32 or 64),uint8(byte),uint16,uint32,uint64
float32,float64
string
complex64,complex128
array -- 固定长度的数组
```
### 2、引用类型（指针类型）
``` go
slice -- 序列数组（最常用）
map -- 映射
chan -- 管道
```
### 3、内置函数
```go
append  -- 用来追加元素到数组，slice中，返回修改后的数组、slice
close   -- 主要用来关闭channel
delete  -- 从map中删除key对应的value
panic   -- 停止常规的goroutime（panic和recover：用来做错误处理）
recover -- 允许程序定义goroutine的panic动作
real    -- 返回complex的实部 （complex、real、imag：用来创建和操作复数）
make    -- 用来分配内存，主要用来分配值类型，比如int、struct、channel
new     -- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针
cap     -- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和map）
copy    -- 用于复制和连接slice，返回复制的数目
len     -- 用来求长度，比如string、array、slice、map、channel
print、println  --底层打印函数，再部署环境在建议使用fmt包
```
### 4、内置接口error
```go
type error interface {
    // 只要实现了Error()函数，返回值为String的都实现了error接口
    Error() String
}
```