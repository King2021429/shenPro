package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// MyMessage 表示预期接收的消息结构
type MyMessage struct {
	Event string `json:"event"`
	// 可以添加更多字段
}

// Response 结构用于定义响应格式
type Response struct {
	ActionStatus int64 `json:"action_status"`
}

type DlcCallbackRespData struct {
	// 执行状态
	ActionStatus int64 `json:"action_status"`
}

func main() {
	http.HandleFunc("/dlc", handler)

	// 在80端口上运行服务器
	fmt.Println("Server is running on port 80...")
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析JSON请求体
	var req MyMessage
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// 打印请求信息
	fmt.Printf("request: %+v\n", req)
	fmt.Printf("header: %+v\n", r.Header)

	// 正确的响应结果
	correctResult := DlcCallbackRespData{ActionStatus: 0}

	// 错误的响应结果
	errorResult := DlcCallbackRespData{ActionStatus: -1}
	//
	//// 验签逻辑
	//signature := r.Header.Get("Authorization")
	//timestamp := r.Header.Get("X-Bili-Timestamp")
	//nonce := r.Header.Get("X-Bili-Signature-Nonce")
	//accessKeyID := r.Header.Get("X-Bili-Accesskeyid")
	//
	//// 构建签名字符串
	//signStr := fmt.Sprintf("x-bili-timestamp:%s\nx-bili-signature-nonce:%s\n", timestamp, nonce)
	//
	//// 计算自己的签名
	//mySign := hmac.New(sha256.New, []byte("1bz5yRkaHlYbbyIAH0unyAvN"))
	//mySign.Write([]byte(signStr))
	//calculatedSign := mySign.Sum(nil)
	//
	//// Base64编码
	//calculatedSignB64 := base64.StdEncoding.EncodeToString(calculatedSign)
	//
	//// 比较签名
	//if signature != calculatedSignB64 {
	//	fmt.Println("sign mismatch!")
	//	json.NewEncoder(w).Encode(errorResult)
	//	return
	//}
	w.Header().Set("Content-Type", "application/json")

	// 业务逻辑处理
	if req.Event == "verify_webhooks" || req.Event == "add_audit_user" || req.Event == "add_user" {
		json.NewEncoder(w).Encode(correctResult)
	} else {
		fmt.Println("not expected event!")
		json.NewEncoder(w).Encode(errorResult)
	}

}
