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
	ArticleToUser ArticleToUser `json:"article_to_user"`
}

type ArticleToUser struct {
	// 文章信息
	Uid     int64  `json:"uid"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Cover   string `json:"cover"`
	// 文章作者信息
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	// 是否点赞
	IsLike bool `json:"is_like"`
	// 是否收藏
	IsFavorite bool `json:"is_favorite"`
}
