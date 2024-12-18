package model

// Message Message结构体用于表示交互消息中的元素，包含角色和内容
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest ChatRequest结构体用于构造向AI服务发送的请求体中的内容
type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}
