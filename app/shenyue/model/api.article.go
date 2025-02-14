package model

type CreateArticleReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateArticleResp struct {
	Id int64 `json:"id"`
}

type GetArticleListReq struct {
	PageNum  int64 `json:"page_num"`
	PageSize int64 `json:"page_size"`
}

type GetArticleListResp struct {
	Total       int64     `json:"total"`
	ArticleList []Article `json:"article_list"`
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
	Article Article `json:"article_api"`
}

type ArticleToUser struct {
	// 文章信息
	Uid     int64  `json:"uid"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Cover   string `json:"cover"`
	// 文章作者信息
}
