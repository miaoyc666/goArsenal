# goArsenal

### goArsenal是什么
目标是将goArsenal打造成Golang代码的武器库，要包罗万象。通用函数，代码小工具等等应有尽有

### 基础语法
[基础语法](https://github.com/miaoyc666/rd-manual/tree/main/Golang)

### 功能列表
1. [qps test](./qps/qps.go)，限制qps测试
2. [json](./json/)，json操作示例，包含多种json库（标准json库、fastjson和gjson）的解析与生成示例
3. [version diff](./versionDiff/versionDiff.go)，版本号比较    
4. [system](./system/system.go)，系统命令调用
5. [file](./file/file.go)，文件操作相关
6. [get mac addrs](./network/network.go)，获取mac地址列表，获取pci总线上真实的网卡顺序
7. [service register](./serviceRegister/serviceRegister.go), 接口对象注册注册，可用于业务逻辑抽象，逻辑层与io层分离
8. [udp](./udp/udpClient.go), udp程序示例
9. [flag](./flag/flag.go), 获取命令行参数
10. [aes](./aes/aes.go), aes ecb加解密示例，包含base64转换。python版本aes加解密[aes.py](https://github.com/miaoyc666/pyArsenal/blob/master/aes.py)
11. [panic](./panic/main.go), panic 
12. [chan](./chan/chan.go), chan 

### Go常见错误
[Go常见错误](https://github.com/miaoyc666/go-mistakes)

#### gotests
```bash
# 安装gotests
go get -u github.com/cweill/gotests/...
# env
export PATH=$PATH:$GOPATH/bin
# 生成所有测试用例
gotests -all {$filename}
```

#### go test
```bash
go test -v {$testfile} {$sourcefile}
go test -v {$testfile} {$sourcefile} -test.run {$test case name}
```
