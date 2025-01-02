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

type DeleteArticleReq struct {
	ArticleId int64 `json:"article_id"`
}

type DeleteArticleResp struct {
}
