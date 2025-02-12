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

// OllamaRequest 定义 Ollama 请求结构体
type OllamaRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature"`
}

// OllamaResponse 定义 Ollama 响应结构体
type OllamaResponse struct {
	Response string `json:"response"`
}

// deepseek
//代码生成/数学	0.0
//数据抽取/分析	1.0
//通用对话	1.3
//翻译	1.3
//创意类写作/诗歌创作	1.5

const (
	TaskCodeGenerationMath     = 0.0
	TaskCommon                 = 0.3
	TaskDataExtractionAnalysis = 1.0
	TaskGeneralConversation    = 1.3
	TaskTranslation            = 1.3
	TaskCreativeWritingPoetry  = 1.5
)
const (
	DeepSeekModel = "deepseek-r1"
	MoonshotModel = "moonshot-v1-8k"
)
