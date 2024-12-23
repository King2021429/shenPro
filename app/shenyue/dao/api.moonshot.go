package dao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shenyue-gin/app/shenyue/model"
)

func (d *Dao) AIChat(query string, history *[]model.Message) string {
	// 将新的用户消息添加到历史记录中
	*history = append(*history, model.Message{
		Role:    "user",
		Content: query,
	})

	// 构造请求体
	reqBody := model.ChatRequest{
		Model:       "moonshot-v1-8k",
		Messages:    *history,
		Temperature: 0.3,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("请求体序列化失败:", err)
		return ""
	}

	// 创建POST请求
	client := &http.Client{}
	url := "https://api.moonshot.cn/v1/chat/completions" // 根据实际情况调整完整路径等
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	// 这里假设需要类似的认证头，实际按照其API要求来准确设置
	req.Header.Set("Authorization", "Bearer sk-D3lXzQcv8Uflx0AjdyolQdjGlO7yPa7be0oligkSt2dfX7Ab")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return ""
	}
	defer resp.Body.Close()

	// 读取响应体内容
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return ""
	}

	// 这里简单假设响应体里有个类似的结构体包含回复内容，实际要根据其API文档解析正确的数据结构
	var respData struct {
		Choices []struct {
			Message model.Message `json:"message"`
		} `json:"choices"`
	}
	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		fmt.Println("解析响应体失败:", err)
		return ""
	}

	// 获取回复内容，并添加到历史记录中
	result := respData.Choices[0].Message.Content
	*history = append(*history, model.Message{
		Role:    "assistant",
		Content: result,
	})
	return result
}
