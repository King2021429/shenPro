package service

import (
	"context"
	"fmt"
	"shenyue-gin/app/shenyue/errorcode"
	"shenyue-gin/app/shenyue/model"
)

// 其他已有函数...

// LikeArticle 实现文章点赞功能
func (s *Service) LikeArticle(ctx context.Context, req *model.LikeArticleReq, uid int64) (resp *model.LikeArticleResp, errCode int64) {
	resp = &model.LikeArticleResp{}

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
	fmt.Println(1)

	// 获取文章点赞信息
	articleLike, err := s.dao.GetArticleLikeByUserAndArticle(ctx, uid, req.ArticleId)
	fmt.Println(2)

	if err != nil {
		// 如果没有点赞记录，创建一个新的点赞记录
		fmt.Println(3)

		newArticleLike := &model.ArticleLike{
			UserID:    uid,
			ArticleID: req.ArticleId,
			Status:    req.Status,
		}
		err = s.dao.CreateArticleLike(ctx, newArticleLike)
		if err != nil {
			fmt.Println(err)
			return nil, errorcode.ErrParam
		}
	} else {
		fmt.Println(4)
		// 如果已经有点赞记录，更新点赞状态
		err = s.dao.UpdateArticleLikeStatus(ctx, int64(articleLike.ID), req.Status)
		if err != nil {
			fmt.Println(err)
			return nil, errorcode.ErrParam
		}
	}

	return resp, 0
}

// GetLikeList 获取用户点赞列表
func (s *Service) GetLikeList(ctx context.Context, req *model.LikeArticleListReq) (resp *model.LikeArticleListResp, errCode int64) {
	resp = &model.LikeArticleListResp{}

	// 检查请求参数
	if req.UserId == 0 {
		return nil, errorcode.ErrParam
	}

	// 根据用户ID查询点赞的文章列表
	articleLikes, err := s.dao.GetArticleLikesByUser(ctx, req.UserId)
	if err != nil {
		fmt.Println(err)
		return nil, errorcode.ErrParam
	}

	// 提取文章ID列表
	articleIds := make([]int64, 0, len(articleLikes))
	for _, like := range articleLikes {
		articleIds = append(articleIds, int64(like.ArticleID))
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
