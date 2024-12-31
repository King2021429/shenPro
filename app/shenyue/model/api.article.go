package model

type CreateArticleReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateArticleResp struct {
}

type GetArticleListReq struct {
}

type GetArticleListResp struct {
}
