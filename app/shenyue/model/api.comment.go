package model

type CreateCommentReq struct {
	ArticleID int64  `json:"article_id" validate:"required"`    // 关联的文章 ID，必填
	UserID    int64  `json:"user_id" validate:"required"`       // 评论用户的 ID，必填
	Content   string `json:"content" validate:"required,min=1"` // 评论内容，必填且至少 1 个字符
	ParentID  int64  `json:"parent_id,omitempty"`               // 父评论 ID，可选
}

type CreateCommentResp struct {
	// 可根据实际需求添加字段
}

type EditCommentReq struct {
	CommentId uint   `json:"comment_id"`
	Content   string `json:"content"`
}

type EditCommentResp struct {
	// 可根据实际需求添加字段
}

type DeleteCommentReq struct {
	CommentId uint `json:"comment_id"`
}

type GetCommentListReq struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}

type GetCommentListResp struct {
	CommentList []Comment `json:"comment_list"`
}

type GetCommentByIdReq struct {
	CommentId uint `json:"comment_id"`
}

type GetCommentByIdResp struct {
	Comment Comment `json:"comment"`
}
