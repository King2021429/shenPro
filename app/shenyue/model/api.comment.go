package model

type CreateCommentReq struct {
	Content string `json:"content"`
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
