# goArsenal

### What is goArsenal?
目标是将goArsenal打造成Golang代码的武器库，要包罗万象。通用函数，代码小工具等等应有尽有

### 基础语法
[基础语法](https://github.com/miaoyc666/rd-manual/tree/main/Golang)

### 使用方法
在go.mog中引用github.com/miaoyc666/goArsenal即可使用函数库中的函数

### 功能列表
1. [aes](crypto/aes/aes.go), aes ecb加解密示例，包含base64转换，对应的python版本[aes.py](https://github.com/miaoyc666/pyArsenal/blob/master/aes.py)
2. [array](array/array.go), 数组相关操作
3. [chan](chan/chan.go), chan
4. [crypto](crypto/md5.go), 字符串加密
5. [download](download/download.go), download file
6. [file](file/file.go)，文件操作相关
7. [flag](flag/flag.go), 获取命令行参数
8. [http](http/http.go): http客户端
9. [json](json), json操作示例，包含多种json库（标准json库、fastjson和gjson）的解析与生成示例
10. [mac addr](network/network.go)，获取mac地址列表，获取pci总线上真实的网卡顺序
11. [mongo](mongo/main.go), mongodb写入与读取
12. [orm](orm/README.md), orm
13. [panic](panic/main.go), panic 
14. [qps test](qps/qps.go)，限制qps测试
15. [re](re/re.go), regexp and regexp2
16. [service register](serviceRegister/serviceRegister.go), 接口对象注册注册，可用于业务逻辑抽象，逻辑层与io层分离
17. [system](system/system.go)，系统命令调用
18. [udp](udp/udpClient.go), udp程序示例
19. [version diff](versionDiff/versionDiff.go)，版本号比较

### 函数列表
#### crypt函数列表
- Md5 获取字符串md5
- MD5sumFromFile 计算文件md5

#### logger函数列表
- Setup 设置日志模块配置
- Debug/Info/Warn/Error...

#### http
- NewHttpClient 新建http客户端，可指定代理和tls参数

### Go常见错误
[go-mistakes](https://github.com/miaoyc666/go-mistakes)

### gotests
```bash
# install gotests
go get -u github.com/cweill/gotests/...
# export env
export PATH=$PATH:$GOPATH/bin
# generate all test cases
gotests -all {$filename}
```

### Unit Testing, use go test
- test file  
`go test -v {$testfile} {$sourcefile}`
- test single function  
`go test -v {$testfile} {$sourcefile} -test.run {$test case name}`
- run benchmark  
`go test -bench=. {$testfile} {$sourcefile}`
