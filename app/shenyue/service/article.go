package service

import (
	"context"
	"fmt"
	"shenyue-gin/app/shenyue/model"
)

func (s *Service) CreateArticle(ctx context.Context, req *model.CreateArticleReq, uid int64) (resp *model.CreateArticleResp, err error) {
	resp = &model.CreateArticleResp{}
	if req.Title == "" || req.Content == "" || uid == 0 {
		return nil, err
	}
	newArticle := &model.Article{
		UID:     uid,
		Content: req.Content,
		Title:   req.Title,
	}
	err = s.dao.CreateArticle(ctx, newArticle)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}

func (s *Service) GetArticleList(ctx context.Context, req *model.GetArticleListReq, uid int64) (resp *model.GetArticleListResp, err error) {
	resp = &model.GetArticleListResp{}
	_, err = s.dao.GetArticleList(ctx, 1, 1)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}
