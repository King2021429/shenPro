package model

type CreateArticleReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateArticleResp struct {
}
