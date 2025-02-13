package dao

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"shenyue-gin/app/shenyue/model"
)

// 文章模块

// GetArticlesByIds 根据文章ID列表查询文章信息
func (d *Dao) GetArticlesByIds(ctx context.Context, articleIds []int64) ([]model.Article, error) {
	var articles []model.Article
	err := d.db.WithContext(ctx).Where("id IN ?", articleIds).Find(&articles).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get articles by ids: %w", err)
	}
	return articles, nil
}

// CreateArticle 创建文章
func (d *Dao) CreateArticle(ctx context.Context, article *model.Article) (uint, error) {
	result := d.db.Create(article)
	if result.Error != nil {
		return 0, result.Error
	}
	return article.ID, nil
}

// GetArticleById 获取文章
func (d *Dao) GetArticleById(ctx context.Context, id uint) (model.Article, error) {
	var article model.Article
	err := d.db.First(&article, id).Error
	return article, err
}

// UpdateArticle 更新文章
func (d *Dao) UpdateArticle(ctx context.Context, article *model.Article) error {
	return d.db.Save(&article).Error
}

func (d *Dao) GetArticleList(ctx context.Context, pageSize int, pageNumber int) ([]model.Article, error) {
	// 计算偏移量
	offset := (pageNumber - 1) * pageSize
	// 分页获取文章列表
	var articles []model.Article
	err := d.db.Limit(pageSize).Offset(offset).Find(&articles).Error
	if err != nil {
		fmt.Println("获取文章列表失败: ", err)
		return nil, err
	}
	return articles, nil
}

// DeleteArticle 删除文章
func (d *Dao) DeleteArticle(ctx context.Context, id uint) error {
	return d.db.Delete(&model.Article{}, id).Error
}

// 文章点赞模块

// CreateArticleLike 创建文章点赞
func (d *Dao) CreateArticleLike(ctx context.Context, like *model.ArticleLike) error {
	return d.db.Create(&like).Error
}

// GetArticleLikeByUserAndArticle 根据 Uid 和 ArticleID 查询记录
func (d *Dao) GetArticleLikeByUserAndArticle(ctx context.Context, uid, articleID int64) (*model.ArticleLike, error) {
	var articleLike model.ArticleLike
	result := d.db.WithContext(ctx).Where("uid = ? AND article_id = ?", uid, articleID).First(&articleLike)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // 未找到记录
		}
		return nil, result.Error // 其他错误
	}
	return &articleLike, nil
}

// UpdateArticleLikeStatus 根据 ID 更新 ArticleLike 的 Status
func (d *Dao) UpdateArticleLikeStatus(ctx context.Context, id, newStatus int64) error {
	result := d.db.WithContext(ctx).Model(&model.ArticleLike{}).Where("id = ?", id).Update("status", newStatus)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetArticleLikesByUser 根据用户ID查询点赞的文章列表
func (d *Dao) GetArticleLikesByUser(ctx context.Context, userId int64) ([]model.ArticleLike, error) {
	var articleLikes []model.ArticleLike
	err := d.db.WithContext(ctx).Where("uid = ?", userId).Find(&articleLikes).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get article likes by user: %w", err)
	}
	return articleLikes, nil
}

// 文章收藏模块

// GetArticleFavoriteByUserAndArticle 根据 Uid 和 ArticleID 查询记录
func (d *Dao) GetArticleFavoriteByUserAndArticle(ctx context.Context, uid, articleID int64) (*model.ArticleFavorite, error) {
	var articleFavorite model.ArticleFavorite
	result := d.db.WithContext(ctx).Where("uid = ? AND article_id = ?", uid, articleID).First(&articleFavorite)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // 未找到记录
		}
		return nil, result.Error // 其他错误
	}
	return &articleFavorite, nil
}

// CreateArticleFavorite 创建文章收藏
func (d *Dao) CreateArticleFavorite(ctx context.Context, favorite *model.ArticleFavorite) error {
	return d.db.Create(&favorite).Error
}

// UpdateArticleFavoriteStatus 根据 ID 更新 ArticleLike 的 Status
func (d *Dao) UpdateArticleFavoriteStatus(ctx context.Context, id, newStatus int64) error {
	result := d.db.WithContext(ctx).Model(&model.ArticleFavorite{}).Where("id = ?", id).Update("status", newStatus)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetArticleFavoriteByUser 根据用户ID查询收藏的文章列表
func (d *Dao) GetArticleFavoriteByUser(ctx context.Context, userId int64) ([]model.ArticleFavorite, error) {
	var articleFavoriteList []model.ArticleFavorite
	err := d.db.WithContext(ctx).Where("uid = ?", userId).Find(&articleFavoriteList).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get article likes by user: %w", err)
	}
	return articleFavoriteList, nil
}
