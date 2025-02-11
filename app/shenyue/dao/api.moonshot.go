package dao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
		Model: "moonshot-v1-8k",

		//Model:       "deepseek-r1",
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
	url := d.c.Moonshot.Url
	//url := "http://localhost:11434/api/generate"

	//url := "https://api.moonshot.cn/v1/chat/completions" // 根据实际情况调整完整路径等
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", d.c.Moonshot.Authorization)

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

// AIChat 实现多轮对话的方法
func (d *Dao) AIChatDeep(query string, history *[]model.Message) string {
	// 将新的用户消息添加到历史记录中
	*history = append(*history, model.Message{
		Role:    "user",
		Content: query,
	})

	// 拼接历史消息和当前用户输入作为完整的提示
	prompt := ""
	for _, msg := range *history {
		if msg.Role == "user" {
			prompt += fmt.Sprintf("用户: %s\n", msg.Content)
		} else if msg.Role == "assistant" {
			prompt += fmt.Sprintf("助手: %s\n", msg.Content)
		}
	}

	// 构造请求体
	reqBody := model.OllamaRequest{
		Model:       "deepseek-r1",
		Prompt:      prompt,
		Temperature: 1.5,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("请求体序列化失败:", err)
		return ""
	}

	// 创建 POST 请求
	client := &http.Client{}
	url := "http://localhost:11434/api/generate"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return ""
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("请求失败，状态码: %d\n", resp.StatusCode)
		return ""
	}

	// 逐行读取响应并解析每个 JSON 对象
	var result string
	decoder := json.NewDecoder(resp.Body)
	for {
		var respData model.OllamaResponse
		err := decoder.Decode(&respData)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("解析响应体失败:", err)
			return ""
		}
		result += respData.Response
	}

	// 获取回复内容，并添加到历史记录中
	*history = append(*history, model.Message{
		Role:    "assistant",
		Content: result,
	})
	return result
}
