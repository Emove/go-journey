<!--
 * @author: Emove
 * @date: Do not edit
 -->
## Select
select语句类似于switch语句，但是select会随机执行一个可运行的case，如果没有case可以运行，它将阻塞，直到有case可以运行。
select是Go中的一个控制结构，每个case必须是一个通信操作，要么是发送要么是接收。一个默认的子句应该总是可运行的。
### 1、语法
```go
    select {
        case communication clause:
            statement(s);
        case communication clause:
            statement(s);
        default:
            statement(s);
    }
```
以下描述了select语句语法：
1. 在select语句中，Go会按照顺序从头到尾评估每一个发送和接收语句
2. 每个case都必须是一个通信
3. 所有channel表达式都会被求值
4. 所有被发送的表达式都会被求值
5. 如果任意某个通信可以进行，他就会执行；其他忽略
6. 如果有多个case可以执行，select会随机公平的选出一个执行；其他忽略
7. 如果有default子句，则执行该语句
8. 如果没有default语句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值
9. 每个case语句里必须是一个IO操作
