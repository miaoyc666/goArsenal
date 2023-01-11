# goArsenal

### What is goArsenal?
目标是将goArsenal打造成Golang代码的武器库，要包罗万象。通用函数，代码小工具等等应有尽有

### 基础语法
[基础语法](https://github.com/miaoyc666/rd-manual/tree/main/Golang)

### 功能列表
1. [aes](./aes/aes.go), aes ecb加解密示例，包含base64转换，对应的python版本[aes.py](https://github.com/miaoyc666/pyArsenal/blob/master/aes.py)
2. [array](./array/array.go), 数组相关操作
3. [chan](./chan/chan.go), chan
4. [download](./download/download.go), download file
5. [file](./file/file.go)，文件操作相关
6. [flag](./flag/flag.go), 获取命令行参数
7. [json](./json/)，json操作示例，包含多种json库（标准json库、fastjson和gjson）的解析与生成示例
8. [mac addr](./network/network.go)，获取mac地址列表，获取pci总线上真实的网卡顺序
9. [mongo](./mongo/main.go), mongodb写入与读取
10. [orm](./orm/README.md), orm
11. [panic](./panic/main.go), panic 
12. [qps test](./qps/qps.go)，限制qps测试
13. [re](./re/re.go), regexp and regexp2
14. [service register](./serviceRegister/serviceRegister.go), 接口对象注册注册，可用于业务逻辑抽象，逻辑层与io层分离
15. [system](./system/system.go)，系统命令调用
16. [udp](./udp/udpClient.go), udp程序示例
17. [version diff](./versionDiff/versionDiff.go)，版本号比较

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
