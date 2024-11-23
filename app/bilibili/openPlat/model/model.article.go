package model

type ArticleAddReq struct {
	Title      string `json:"title"`
	Category   int64  `json:"category"`
	TemplateId int64  `json:"template_id"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
}
