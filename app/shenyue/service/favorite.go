package service

import (
	"context"
	"fmt"
	"shenyue-gin/app/shenyue/errorcode"
	"shenyue-gin/app/shenyue/model"
)

// FavoriteArticle 实现文章收藏功能
func (s *Service) FavoriteArticle(ctx context.Context, req *model.FavoriteArticleReq, uid int64) (resp *model.FavoriteArticleResp, errCode int64) {
	resp = &model.FavoriteArticleResp{}

	// 检查请求参数
	if req.ArticleId == 0 || uid == 0 || (req.Status != 1 && req.Status != 2) {
		return nil, errorcode.ErrParam
	}

	// 先尝试获取文章信息，检查文章是否存在
	article, err := s.dao.GetArticleById(ctx, uint(req.ArticleId))
	if err != nil || article.ID == 0 {
		fmt.Println(err)
		return nil, errorcode.ErrParam
	}
	// 获取文章点赞信息
	articleLike, err := s.dao.GetArticleFavoriteByUserAndArticle(ctx, uid, req.ArticleId)
	if err != nil {
		// 如果没有点赞记录，创建一个新的点赞记录
		newArticleLike := &model.ArticleFavorite{
			Uid:       uid,
			ArticleID: req.ArticleId,
			Status:    req.Status,
		}
		err = s.dao.CreateArticleFavorite(ctx, newArticleLike)
		if err != nil {
			fmt.Println(err)
			return nil, errorcode.ErrParam
		}
	} else {
		// 如果已经有点赞记录，更新点赞状态
		err = s.dao.UpdateArticleFavoriteStatus(ctx, int64(articleLike.ID), req.Status)
		if err != nil {
			fmt.Println(err)
			return nil, errorcode.ErrParam
		}
	}

	return resp, 0
}

// GetFavoriteList 获取用户收藏列表
func (s *Service) GetFavoriteList(ctx context.Context, req *model.FavoriteArticleListReq) (resp *model.FavoriteArticleListResp, errCode int64) {
	resp = &model.FavoriteArticleListResp{}

	// 检查请求参数
	if req.Uid == 0 {
		return nil, errorcode.ErrParam
	}

	// 根据用户ID查询点赞的文章列表
	articleFavoriteList, err := s.dao.GetArticleFavoriteByUser(ctx, req.Uid)
	if err != nil {
		fmt.Println(err)
		return nil, errorcode.ErrParam
	}

	// 提取文章ID列表
	articleIds := make([]int64, 0, len(articleFavoriteList))
	for _, like := range articleFavoriteList {
		articleIds = append(articleIds, like.ArticleID)
	}

	// 根据文章ID列表查询文章信息
	articles, err := s.dao.GetArticlesByIds(ctx, articleIds)
	if err != nil {
		fmt.Println(err)
		return nil, errorcode.ErrParam
	}

	// 将文章信息填充到响应结构体中
	resp.ArticleList = articles
	return resp, 0
}
