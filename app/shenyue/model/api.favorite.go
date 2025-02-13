package model

// FavoriteArticleReq 收藏文章请求结构体
type FavoriteArticleReq struct {
	ArticleId int64 `json:"article_id"`
	Status    int64 `json:"status"` // 1 收藏 2 取消收藏
}

// FavoriteArticleResp 收藏文章响应结构体
type FavoriteArticleResp struct {
}

// FavoriteArticleListReq 收藏文章列表请求结构体
type FavoriteArticleListReq struct {
	UserId int64 `json:"user_id"`
}

// FavoriteArticleListResp 收藏文章列表响应结构体
type FavoriteArticleListResp struct {
	ArticleList []Article `json:"article_list"`
}
