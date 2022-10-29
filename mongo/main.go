package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var CollectionTest *mongo.Collection
var myTestModel TestModel

const MaxExecTime = 300
const TB = "test"

type TestModel struct {
	Action int32
	Value1 string
	Value2 int32
	Tag    []string
	OldId  int64 `bson:"old_id,omitempty" json:"old_id"`
}

func GetMongoClient(addr string) (*mongo.Client, error) {
	//获取mongo客户端
	clientOptions := options.Client().ApplyURI(addr)
	clientOptions.SetMaxPoolSize(1000)

	//初始化连接设置超时时间
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查db是否连接成功
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}

func InitLocalConnection() {
	// 连接本地db
	mongoAddr := "mongodb://root:xxx@xxxx:27017"
	client, _ := GetMongoClient(mongoAddr)

	// 利用client连接到具体的数据库,创建handle
	CollectionTest = client.Database("miaoyc").Collection(TB)
}

// ReadFileContent2db
/*
 *逐行解析文件里面的内容
 *将文件里面的内容批量添加到db
 */
func ReadFileContent2db(fileName, tableName string) error {
	fi, err := os.Open(fileName)
	defer fi.Close()
	buf := bufio.NewReader(fi)
	count := 0
	var docs []interface{}

	for {
		a, c := buf.ReadString('\n')
		if c == io.EOF {
			break
		}
		count++
		json.Unmarshal([]byte(a), &myTestModel)
		s, _ := bson.Marshal(myTestModel)
		docs = append(docs, s)

		if count == 10 {
			err = BatchCreateDocument(docs, CollectionTest)
			fmt.Println("BatchCreate is", err)
			count = 0
			docs = docs[:0]
		}
	}
	err = BatchCreateDocument(docs, CollectionTest)
	CreateIndex(CollectionTest, "old_id", TB)
	return err
}

// BatchCreateDocument 在db中批量创建文档
func BatchCreateDocument(docs []interface{}, tableCollection *mongo.Collection) error {
	opts := options.InsertMany().SetOrdered(false)
	_, err := tableCollection.InsertMany(context.TODO(), docs, opts)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndex 创建索引
func CreateIndex(tableCollection *mongo.Collection, filedName, tableName string) error {
	indexName := tableName + "_Index"
	model := mongo.IndexModel{
		Keys:    bson.D{{filedName, 1}},
		Options: options.Index().SetName(indexName),
	}
	opts := options.CreateIndexes().SetMaxTime(MaxExecTime * time.Second)
	_, err := tableCollection.Indexes().CreateOne(context.TODO(), model, opts)
	if err != nil {
		return err
	}
	return nil
}

func (m *TestModel) GetByOid(oldId string) TestModel {
	var info TestModel
	n, _ := strconv.ParseInt(oldId, 10, 64)
	manualCondition := bson.M{"old_id": n}
	CollectionTest.FindOne(context.TODO(), manualCondition).Decode(&info)
	return info
}

func main() {
	InitLocalConnection()
	fileName := "case.json"
	ReadFileContent2db(fileName, TB)
	var info TestModel
	r := info.GetByOid("8088464930757410816")
	fmt.Println(r)
}
