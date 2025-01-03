package model

type CreateArticleReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateArticleResp struct {
}

type GetArticleListReq struct {
	PageNum  int64 `json:"page_num"`
	PageSize int64 `json:"page_size"`
}

type GetArticleListResp struct {
}

type DeleteArticleReq struct {
	ArticleId int64 `json:"article_id"`
}

type DeleteArticleResp struct {
}

type EditArticleReq struct {
	ArticleId int64  `json:"article_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type EditArticleResp struct {
}

type GetArticleByIdReq struct {
	ArticleId int64 `json:"article_id"`
}

type GetArticleByIdResp struct {
}
