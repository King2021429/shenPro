package dao

import (
	"context"
	"fmt"
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

// UpdateArticleLike  更新文章点赞
func (d *Dao) UpdateArticleLike(ctx context.Context, like *model.ArticleLike) error {
	return d.db.Save(&like).Error
}
