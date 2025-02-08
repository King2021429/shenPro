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
	}
	errDb := s.dao.CreateComment(ctx, *newComment)
	if errDb != nil {
		fmt.Println(errDb)
		return nil, errorcode.ERROR
	}
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
