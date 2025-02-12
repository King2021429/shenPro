package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"shenyue-gin/app/shenyue/model"
)

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

// CreateArticleCollection 创建文章收藏
func (d *Dao) CreateArticleCollection(ctx context.Context, collection model.ArticleCollection) error {
	return d.db.Create(&collection).Error
}

// GetArticleCollection 获取文章收藏
func (d *Dao) GetArticleCollection(ctx context.Context, id uint) (model.ArticleCollection, error) {
	var collection model.ArticleCollection
	err := d.db.First(&collection, id).Error
	return collection, err
}

// UpdateArticleCollection 更新文章收藏
func (d *Dao) UpdateArticleCollection(ctx context.Context, collection model.ArticleCollection) error {
	return d.db.Save(&collection).Error
}

// DeleteArticleCollection 删除文章收藏
func (d *Dao) DeleteArticleCollection(ctx context.Context, id uint) error {
	return d.db.Delete(&model.ArticleCollection{}, id).Error
}

// CreateArticleLike 创建文章点赞
func (d *Dao) CreateArticleLike(ctx context.Context, like *model.ArticleLike) error {
	return d.db.Create(&like).Error
}

// GetArticleLike 获取文章点赞
func (d *Dao) GetArticleLike(ctx context.Context, id uint) (model.ArticleLike, error) {
	var like model.ArticleLike
	err := d.db.First(&like, id).Error
	return like, err
}

// GetArticleLikeByUserAndArticle 根据 UserID 和 ArticleID 查询记录
func (d *Dao) GetArticleLikeByUserAndArticle(ctx context.Context, userID, articleID int64) (*model.ArticleLike, error) {
	var articleLike model.ArticleLike
	result := d.db.WithContext(ctx).Where("user_id = ? AND article_id = ?", userID, articleID).First(&articleLike)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
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
	err := d.db.WithContext(ctx).Where("user_id = ?", userId).Find(&articleLikes).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get article likes by user: %w", err)
	}
	return articleLikes, nil
}

// GetArticlesByIds 根据文章ID列表查询文章信息
func (d *Dao) GetArticlesByIds(ctx context.Context, articleIds []int64) ([]model.Article, error) {
	var articles []model.Article
	err := d.db.WithContext(ctx).Where("id IN ?", articleIds).Find(&articles).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get articles by ids: %w", err)
	}
	return articles, nil
}
