package main

import (
	"bufio"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	CfgFieName = "cfg.ini"     //程序运行时的配置文件
	SECTION    = "test"        // 配置文件中的section项
	PoolNum    = "pool_number" // pool_number配置
	FileName   = "file_name"   // 待检测数据的文件
	PATH       = "configPath"  //指定配置文件地址
	QueryLimit = "qpsLimit"    // qps限制
	TotalCount = "totalCount"  // 压测总条目数

)

var globalCount uint64
var globalChan = make(chan string, 1000010)
var wg sync.WaitGroup
var p *ants.PoolWithFunc

/*
* 逐行读取文件里面的内容
* 获取要查询的原始数据
* 函数返回结果为string类型的map
 */
func GetData(fileName string) {
	f, _ := os.Open(fileName)
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.TrimSpace(line)
		globalChan <- line
	}
}

func Timer(qpsLimit int, totalCount int)  {
	d := time.Duration(time.Second * 1)
	t := time.NewTicker(d)
	defer t.Stop()

	currentProcess := 0
	for {
		<-t.C
		if currentProcess >= totalCount {
			break
		}
		count := 0
		bT := time.Now()
		for {
			if count >= qpsLimit{
				break
			}
			if len(globalChan) == 0 {
				currentProcess = totalCount
				break
			}
			v := <-globalChan
			wg.Add(1)
			p.Invoke(v)
			count += 1
			currentProcess += 1
		}
		eT := time.Since(bT)
		fmt.Printf("cost time is: %v\n", eT)
		fmt.Println("process:", globalCount , totalCount)
	}
}

/*
调用接口函数
 */
func GetTestFunc() string {
	return ""
}

/*
反序列化数据函数预留
*/
func Unmarshal(data unsafe.Pointer) int {
	return 0
}

func run() {
	// 读取配置文件
	conf, err := config.ReadDefault(CfgFieName)
	if err != nil {
		log.Fatalf("Read config file error: %s", err)
	}

	fileName, _ := conf.String(SECTION, FileName)
	poolNumbers, _ := conf.Int(SECTION, PoolNum)
	path, _ := conf.String(SECTION, PATH)
	qpsLimit, _ := conf.Int(SECTION, QueryLimit)
	totalCount, _ := conf.Int(SECTION, TotalCount)

	log.Printf("file_name is: %v", fileName)
	log.Printf("pool_numbers is: %v", poolNumbers)
	log.Printf("config path: %v", path)

	GetData(fileName)
	// 并发请求
	defer ants.Release()
	p, _ = ants.NewPoolWithFunc(poolNumbers, func(i interface{}) {
		data := GetTestFunc()
		Unmarshal(unsafe.Pointer(&data))
		atomic.AddUint64(&globalCount, 1)
		wg.Done()
	})
	defer p.Release()
	Timer(qpsLimit, totalCount)
	time.Sleep(time.Second * 2)
	wg.Wait()
	fmt.Println("process:", globalCount , totalCount)
	fmt.Println("end")
}

func main()  {
	run()
}