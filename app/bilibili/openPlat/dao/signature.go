package dao

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"shenyue-gin/app/bilibili/openPlat/model"
	"sort"
	"strings"
)

// CreateSignature 生成Authorization加密串
func CreateSignature(header *model.CommonHeader, accessKeySecret string) string {
	sStr := ToSortedString(header)
	return HmacSHA256(accessKeySecret, sStr)
}

// Md5 md5加密
func Md5(str string) (md5str string) {
	data := []byte(str)
	has := md5.Sum(data)
	md5str = fmt.Sprintf("%x", has)
	return md5str
}

// HmacSHA256 HMAC-SHA256算法
func HmacSHA256(key string, data string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

// ToMap 所有字段转map<string, string>
func ToMap(h *model.CommonHeader) map[string]string {
	return map[string]string{
		model.BiliTimestampHeader:       h.Timestamp,
		model.BiliSignatureMethodHeader: h.SignatureMethod,
		model.BiliSignatureNonceHeader:  h.Nonce,
		model.BiliAccessKeyIdHeader:     h.AccessKeyId,
		model.BiliSignVersionHeader:     h.SignatureVersion,
		model.BiliContentMD5Header:      h.ContentMD5,
		model.AuthorizationHeader:       h.Authorization,
		model.ContentTypeHeader:         h.ContentType,
		model.AcceptHeader:              h.ContentAcceptType,
		//model.BilispyColor:              h.X1BilispyColor,
		model.AccessToken: h.AccessToken,
	}
}

// ToSortMap 参与加密的字段转map<string, string>
func ToSortMap(h *model.CommonHeader) map[string]string {
	return map[string]string{
		model.BiliTimestampHeader:       h.Timestamp,
		model.BiliSignatureMethodHeader: h.SignatureMethod,
		model.BiliSignatureNonceHeader:  h.Nonce,
		model.BiliAccessKeyIdHeader:     h.AccessKeyId,
		model.BiliSignVersionHeader:     h.SignatureVersion,
		model.BiliContentMD5Header:      h.ContentMD5,
	}
}

// ToSortedString 生成需要加密的文本
func ToSortedString(h *model.CommonHeader) (sign string) {
	hMap := ToSortMap(h)
	var hSil []string
	for k := range hMap {
		hSil = append(hSil, k)
	}
	sort.Strings(hSil)
	for _, v := range hSil {
		sign += v + ":" + hMap[v] + "\n"
	}
	sign = strings.TrimRight(sign, "\n")
	return
}
