package main

import (
	"encoding/json"
	"fmt"
	"github.com/monaco-io/request"
)

const (
	AcceptHeader              = "Accept"
	ContentTypeHeader         = "Content-Type"
	AuthorizationHeader       = "Authorization"
	JsonType                  = "application/json"
	BiliVersion               = "1.0"
	HmacSha256                = "HMAC-SHA256"
	BiliTimestampHeader       = "x-bili-timestamp"
	BiliSignatureMethodHeader = "x-bili-signature-method"
	BiliSignatureNonceHeader  = "x-bili-signature-nonce"
	BiliAccessKeyIdHeader     = "x-bili-accesskeyid"
	BiliSignVersionHeader     = "x-bili-signature-version"
	BiliContentMD5Header      = "x-bili-content-md5"
	BilispyColor              = "x1-bilispy-color"
)

type CommonHeader struct {
	ContentType    string
	X1BilispyColor string
}

type BaseResp struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type WsStartData struct {
	ConnId        string `json:"conn_id"`
	WebsocketInfo struct {
		AuthBody string   `json:"auth_body"`
		WssLink  []string `json:"wss_link"`
	} `json:"websocket_info"`
}

// ApiRequest http request demo方法
func ApiRequest(reqJson, requestUrl string) (resp BaseResp, err error) {
	resp = BaseResp{}
	header := &CommonHeader{
		ContentType: JsonType,
	}

	cli := request.Client{
		Method: "POST",
		URL:    fmt.Sprintf("%s%s", OpenPlatformHttpHost, requestUrl),
		Header: header.ToMap(),
		String: reqJson,
	}
	cliResp := cli.Send().Scan(&resp)
	if !cliResp.OK() {
		err = fmt.Errorf("[error] req:%+v resp:%+v err:%+v", reqJson, resp, cliResp.Error())
	}
	return
}

// ToMap 所有字段转map<string, string>
func (h *CommonHeader) ToMap() map[string]string {
	return map[string]string{
		ContentTypeHeader: h.ContentType,
		BilispyColor:      h.X1BilispyColor,
	}
}
