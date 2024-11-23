package dao

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"shenyue-gin/app/bilibili/openPlat/model"
	"strconv"
	"time"
)

// DaoFormRequest http request demo方法
func DaoFormRequest(requestUrl string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("/Users/shenyue/Downloads/图标2.0.jpg")
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base("/Users/shenyue/Downloads/图标2.0.jpg"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	_ = writer.WriteField("staff_id", "1026006")
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", requestUrl, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	//s, _ := ioutil.ReadAll(req.Body)
	//bodyStr := string(s)
	bodyStrMds := Md5("")

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-bili-timestamp", timestamp)
	req.Header.Add("x-bili-signature-method", "HMAC-SHA256")
	req.Header.Add("x-bili-signature-version", model.BiliVersion)
	req.Header.Add("x-bili-signature-nonce", nonce)
	req.Header.Add("x-bili-accesskeyid", model.ClientIdProd)
	req.Header.Add("x-bili-content-md5", bodyStrMds)
	req.Header.Add("access-token", model.AccessTokenProd)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	header := &model.CommonHeader{
		ContentType:       writer.FormDataContentType(),
		ContentAcceptType: model.JsonType,
		Timestamp:         timestamp,
		SignatureMethod:   model.HmacSha256,
		SignatureVersion:  model.BiliVersion,
		Authorization:     "",
		Nonce:             nonce, //用于幂等,记得替换
		AccessKeyId:       model.ClientIdProd,
		ContentMD5:        bodyStrMds,
		AccessToken:       model.AccessTokenProd,
	}

	req.Header.Add("Authorization", CreateSignature(header, model.AppSecretProd))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	return
}
