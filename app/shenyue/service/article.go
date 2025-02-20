package service

import (
	"context"
	"fmt"
	"shenyue-gin/app/shenyue/errorcode"
	"shenyue-gin/app/shenyue/model"
)

func (s *Service) CreateArticle(ctx context.Context, req *model.CreateArticleReq, Uid int64) (resp *model.CreateArticleResp, err int64) {
	resp = &model.CreateArticleResp{}
	if req.Title == "" || req.Content == "" || Uid == 0 {
		return nil, errorcode.ErrParam
	}
	newArticle := &model.Article{
		Uid:     Uid,
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

func (s *Service) EditArticle(ctx context.Context, req *model.EditArticleReq, Uid int64) (resp *model.EditArticleResp, err int64) {
	resp = &model.EditArticleResp{}
	if req.Title == "" || req.Content == "" || Uid == 0 {
		return nil, errorcode.ErrParam
	}
	article, errDb := s.dao.GetArticleById(ctx, uint(req.ArticleId))
	if errDb != nil {
		return nil, errorcode.ERROR
	}
	if article.Uid != Uid {
		return nil, errorcode.ERROR
	}
	newArticle := &model.Article{
		Uid:     Uid,
		Content: req.Content,
		Title:   req.Title,
	}
	errDb = s.dao.UpdateArticle(ctx, newArticle)
	if errDb != nil {
		fmt.Println(errDb)
		return nil, errorcode.ERROR
	}
	return resp, 0
}

func (s *Service) GetArticleList(ctx context.Context, req *model.GetArticleListReq) (resp *model.GetArticleListResp, err int64) {
	resp = &model.GetArticleListResp{}
	if req.PageNum <= 0 || req.PageSize <= 0 || req.PageSize >= 10000 {
		return nil, errorcode.ERROR
	}
	list, errDb := s.dao.GetArticleList(ctx, int(req.PageSize), int(req.PageNum))
	if errDb != nil {
		fmt.Println(errDb)
		return nil, errorcode.ERROR
	}
	resp.ArticleList = list
	return resp, 0
}

func (s *Service) DeleteArticle(ctx context.Context, req *model.DeleteArticleReq, Uid int64) (resp *model.GetArticleListResp, err int64) {
	resp = &model.GetArticleListResp{}
	article, errDb := s.dao.GetArticleById(ctx, uint(req.ArticleId))
	if errDb != nil {
		return nil, errorcode.ERROR
	}
	if article.Uid != Uid {
		return nil, errorcode.ERROR
	}
	errDb = s.dao.DeleteArticle(ctx, uint(req.ArticleId))
	if errDb != nil {
		fmt.Println(errDb)
		return nil, errorcode.ERROR
	}
	return resp, 0
}

// GetArticleById 根据id获取文章
func (s *Service) GetArticleById(ctx context.Context, req *model.GetArticleByIdReq, Uid int64) (resp *model.GetArticleByIdResp, err int64) {
	resp = &model.GetArticleByIdResp{
		ArticleToUser: model.ArticleToUser{},
	}
	article, errDb := s.dao.GetArticleById(ctx, uint(req.ArticleId))
	if errDb != nil {
		return nil, errorcode.ERROR
	}
	if article.Uid != Uid {
		return nil, errorcode.ERROR
	}
	article, errDb = s.dao.GetArticleById(ctx, uint(req.ArticleId))
	if errDb != nil {
		fmt.Println(errDb)
		return nil, errorcode.ERROR
	}
	if article.Uid != Uid {
		return nil, errorcode.ERROR
	}
	resp.ArticleToUser.Uid = article.Uid
	resp.ArticleToUser.Title = article.Title
	resp.ArticleToUser.Content = article.Content
	resp.ArticleToUser.Cover = article.Cover

	user, errDb := s.dao.SelectByUid(ctx, article.Uid)
	if errDb != nil {
		fmt.Println(errDb)
		return nil, errorcode.ERROR
	}
	resp.ArticleToUser.Username = user.Username
	resp.ArticleToUser.Avatar = user.Avatar

	// 是否点赞
	like, errDb := s.dao.GetArticleLikeByUserAndArticle(ctx, Uid, req.ArticleId)
	if errDb != nil {
		fmt.Println(errDb)
		return nil, errorcode.ERROR
	}
	resp.ArticleToUser.LikeStatus = like.Status
	// 是否收藏
	favorite, errDb := s.dao.GetArticleFavoriteByUserAndArticle(ctx, Uid, req.ArticleId)
	if errDb != nil {
		fmt.Println(errDb)
		return nil, errorcode.ERROR
	}
	resp.ArticleToUser.FavoriteStatus = favorite.Status
	return resp, 0
}
