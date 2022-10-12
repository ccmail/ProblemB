# 选择问题

该代码在constant中通过调整变量DataSetSelect的值来选择执行1/2问。

例如， 此时执行第一问

```go
const(
...
DataSetSelect = "A"
...
)
```

***
此时执行第二问

```go
const (
...
DataSetSelect = "B"
...
)
```

***

两问都会生成大量图片，存放在outPut文件夹下

环境为go 1.18.3

使用了第三方画图包gg，在新环境下会自动从github下载所需第三方包