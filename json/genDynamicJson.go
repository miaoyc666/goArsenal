package json

/*
File name    : genDynamicJson.go
Author       : miaoyc
Create date  : 2022/2/7 12:03 下午
Description  : 生成动态json数据
*/

type ReqParam struct {
	Apikey     string   `form:"apikey" json:"apikey"`
	Param      string   `form:"param" json:"param"`
}

type ReqParams struct {
	Apikey     string   `form:"apikey" json:"apikey"`
	Params     []string `json:"params"`
}

func genDynamicJson() {

}
