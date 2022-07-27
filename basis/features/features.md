## 主要特性

1. 自动立即回收
2. 丰富的内置类型
3. 函数多返回值
4. 错误处理
5. 匿名函数和闭包
6. 类型和接口
7. 并发编程
8. 反射
9. 语言交互性

## 命名规范

1. 首字母可以是任意的Unicode字符或者下划线
2. 剩余字符可以是Unicode字符、下划线、数字
3. 字符长度不限
4. 常用命名规范为驼峰命名法

## 25个关键字

```
break 		default		func	interface		select
case		defer		go		map				struct
chan		else		goto	package			switch
const		fallthrough	range	if				type
continue	for			import	return			var
```

## 37个保留字

```
Constants：	true	false	iota	nil
Types:	int  int8  int16  int32  int64
		uint uint8  uint16  uint32  uint64
		float32  float64  complex64 complex128
		bool	byte	rune	string	error
Functions:	make	len		cap		new		append
			complex		real	imag	panic	
			recove
```

## 可见性

1. 声明在函数内部，是函数的本地值，类似于private
2. 声明在函数外部，是对当前包的可见（包内所有.go文件都可见）的全局值，类似于default
3. 声明在函数外部且首字母大写是所有包可见的全局值，类似于public

## 声明方式

```go
var(生命变量)
const(生命常量)
type(生命类型)
func(生命函数)
```

