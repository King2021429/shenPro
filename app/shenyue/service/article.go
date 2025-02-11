package service

import (
	"context"
	"fmt"
	"shenyue-gin/app/shenyue/errorcode"
	"shenyue-gin/app/shenyue/model"
)

func (s *Service) CreateArticle(ctx context.Context, req *model.CreateArticleReq, uid int64) (resp *model.CreateArticleResp, err int64) {
	resp = &model.CreateArticleResp{}
	if req.Title == "" || req.Content == "" || uid == 0 {
		return nil, errorcode.ErrParam
	}
	newArticle := &model.Article{
		UID:     uid,
		Content: req.Content,
		Title:   req.Title,
	}
	id, errDb := s.dao.CreateArticle(ctx, newArticle)
	if errDb != nil {
		fmt.Println(errDb)
		return nil, errorcode.ERROR
	}
	resp.Id = int64(id)
	return resp, 0
}

func (s *Service) EditArticle(ctx context.Context, req *model.EditArticleReq, uid int64) (resp *model.EditArticleResp, err error) {
	resp = &model.EditArticleResp{}
	if req.Title == "" || req.Content == "" || uid == 0 {
		return nil, nil
	}
	article, err := s.dao.GetArticleById(ctx, uint(req.ArticleId))
	if err != nil {
		return nil, err
	}
	if article.UID != uid {
		return nil, nil
	}
	newArticle := &model.Article{
		UID:     uid,
		Content: req.Content,
		Title:   req.Title,
	}
	errDb := s.dao.UpdateArticle(ctx, newArticle)
	if errDb != nil {
		fmt.Println(errDb)
		return nil, nil
	}
	return resp, nil
}

func (s *Service) GetArticleList(ctx context.Context, req *model.GetArticleListReq) (resp *model.GetArticleListResp, err error) {
	resp = &model.GetArticleListResp{}
	if req.PageNum <= 0 || req.PageSize <= 0 || req.PageSize >= 10000 {
		return nil, nil
	}
	list, err := s.dao.GetArticleList(ctx, int(req.PageSize), int(req.PageNum))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp.ArticleList = list
	return resp, nil
}

func (s *Service) DeleteArticle(ctx context.Context, req *model.DeleteArticleReq, uid int64) (resp *model.GetArticleListResp, err error) {
	resp = &model.GetArticleListResp{}
	article, err := s.dao.GetArticleById(ctx, uint(req.ArticleId))
	if err != nil {
		return nil, err
	}
	if article.UID != uid {
		return nil, err
	}
	err = s.dao.DeleteArticle(ctx, uint(req.ArticleId))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}

// GetArticleById 根据id获取文章
func (s *Service) GetArticleById(ctx context.Context, req *model.GetArticleByIdReq, uid int64) (resp *model.GetArticleByIdResp, err error) {
	resp = &model.GetArticleByIdResp{}
	article, err := s.dao.GetArticleById(ctx, uint(req.ArticleId))
	if err != nil {
		return nil, err
	}
	if article.UID != uid {
		return nil, err
	}
	article, err = s.dao.GetArticleById(ctx, uint(req.ArticleId))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp.Article = article
	return resp, nil
}
