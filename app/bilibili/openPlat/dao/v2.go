package dao

import (
	"encoding/json"
	"fmt"
	"github.com/monaco-io/request"
	"shenyue-gin/app/bilibili/openPlat/model"
	"strconv"
	"time"
)

// ApiRequestV2 http request demo方法
func ApiRequestV2(reqJson, requestUrl string) (resp model.BaseResp, err error) {
	resp = model.BaseResp{}
	header := &model.CommonHeader{
		AccessToken:       model.AccessTokenProd,
		ContentType:       model.JsonType,
		ContentAcceptType: model.JsonType,
		Timestamp:         strconv.FormatInt(time.Now().Unix(), 10),
		SignatureMethod:   model.HmacSha256,
		SignatureVersion:  model.BiliVersionV2,
		Authorization:     "",
		Nonce:             strconv.FormatInt(time.Now().UnixNano(), 10), //用于幂等,记得替换
		AccessKeyId:       model.ClientIdProd,
		ContentMD5:        Md5(reqJson),
		//X1BilispyColor:    model.Color,
	}
	header.Authorization = CreateSignature(header, model.AppSecretProd)
	cli := request.Client{
		Method: "POST",
		URL:    fmt.Sprintf("%s%s", model.MainOpenPlatformHttpHost, requestUrl),
		Header: ToMap(header),
		String: reqJson,
	}

	// 打印请求的cURL命令
	fmt.Println("cURL Command:")

	var respTest interface{}
	cliResp := cli.Send().Scan(&respTest)
	if !cliResp.OK() {
		err = fmt.Errorf("[error] req:%+v resp:%+v err:%+v", reqJson, resp, cliResp.Error())
	}
	// 使用json.MarshalIndent来格式化JSON数据
	jsonData, err := json.MarshalIndent(respTest, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// 打印格式化后的JSON字符串
	fmt.Println(string(jsonData))
	return
}
