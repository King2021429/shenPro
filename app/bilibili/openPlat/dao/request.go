package dao

import (
	"fmt"
	"github.com/monaco-io/request"
	"shenyue-gin/app/bilibili/openPlat/model"
	"strconv"
	"time"
)

// ApiRequest http request demo方法
func ApiRequest(reqJson, requestUrl string) (resp model.BaseResp, err error) {
	resp = model.BaseResp{}
	header := &model.CommonHeader{
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
	fmt.Println(requestUrl)
	fmt.Println(ToMap(header))

	cli := request.Client{
		Method: "POST",
		URL:    fmt.Sprintf("%s%s", model.MainOpenPlatformHttpHost, requestUrl),
		Header: ToMap(header),
		String: reqJson,
	}

	// 打印请求的cURL命令
	fmt.Println("cURL Command:")

	cliResp := cli.Send().Scan(&resp)
	if !cliResp.OK() {
		err = fmt.Errorf("[error] req:%+v resp:%+v err:%+v", reqJson, resp, cliResp.Error())
	}
	return
}
