package service

import (
	"context"
	"fmt"
	"shenyue-gin/app/shenyue/errorcode"
	"shenyue-gin/app/shenyue/model"
)

// CreateComment 创建评论
func (s *Service) CreateComment(ctx context.Context, req *model.CreateCommentReq, uid int64) (resp *model.CreateCommentResp, err int64) {
	resp = &model.CreateCommentResp{}
	if req.Content == "" || uid == 0 {
		return nil, errorcode.ErrParam
	}
	newComment := &model.Comment{
		ArticleID: req.ArticleID,
		Uid:       uid,
		Content:   req.Content,
		ParentID:  req.ParentID,
	}
	id, errDb := s.dao.CreateComment(ctx, *newComment)
	if errDb != nil {
		fmt.Println(errDb)
		return nil, errorcode.ERROR
	}
	resp.Id = int64(id)
	return resp, 0
}

// EditComment 编辑评论
func (s *Service) EditComment(ctx context.Context, req *model.EditCommentReq, uid int64) (resp *model.EditCommentResp, err error) {
	resp = &model.EditCommentResp{}
	if req.Content == "" || uid == 0 {
		return nil, nil
	}
	comment, err := s.dao.GetComment(ctx, uint(req.CommentId))
	if err != nil {
		return nil, err
	}
	if comment.Uid != uid {
		return nil, nil
	}
	newComment := &model.Comment{
		Uid:     uid,
		Content: req.Content,
	}
	errDb := s.dao.UpdateComment(ctx, *newComment)
	if errDb != nil {
		fmt.Println(errDb)
		return nil, nil
	}
	return resp, nil
}

// DeleteComment 删除评论
func (s *Service) DeleteComment(ctx context.Context, req *model.DeleteCommentReq, uid int64) (resp *model.GetCommentListResp, err error) {
	resp = &model.GetCommentListResp{}
	comment, err := s.dao.GetComment(ctx, uint(req.CommentId))
	if err != nil {
		return nil, err
	}
	if comment.Uid != uid {
		return nil, err
	}
	err = s.dao.DeleteComment(ctx, uint(req.CommentId))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}

// GetCommentsByArticleId 根据文章id获取评论列表
func (s *Service) GetCommentsByArticleId(ctx context.Context, req *model.GetCommentListReq) (resp *model.GetCommentListResp, err int64) {
	resp = &model.GetCommentListResp{}
	if req.ArticleID == 0 {
		return nil, errorcode.ErrParam
	}
	comments, errDb := s.dao.GetCommentsByArticleID(ctx, req.ArticleID)
	if errDb != nil {
		fmt.Println(errDb)
		return nil, errorcode.ERROR
	}

	resp.CommentList = s.buildCommentTree(comments)
	return resp, 0
}

// buildCommentTree 根据 parent_id 构建评论的树形结构
func (s *Service) buildCommentTree(comments []model.Comment) []*model.Comment {
	commentMap := make(map[int64]*model.Comment)
	var rootComments []*model.Comment

	// 先将所有评论存入 map 中
	for i := range comments {
		comment := &comments[i]
		commentMap[int64(comment.ID)] = comment
	}

	// 构建树形结构
	for _, comment := range comments {
		if comment.ParentID == 0 {
			rootComments = append(rootComments, commentMap[int64(comment.ID)])
		} else {
			parent, ok := commentMap[comment.ParentID]
			if ok {
				parent.Replies = append(parent.Replies, commentMap[int64(comment.ID)])
			}
		}
	}

	return rootComments
}
