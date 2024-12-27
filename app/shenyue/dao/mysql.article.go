package dao

import (
	"context"
	"shenyue-gin/app/shenyue/model"
)

// CreateArticle 创建文章
func (d *Dao) CreateArticle(ctx context.Context, article model.Article) error {
	return d.db.Create(&article).Error
}

// GetArticle 获取文章
func (d *Dao) GetArticle(ctx context.Context, id int64) (model.Article, error) {
	var article model.Article
	err := d.db.First(&article, id).Error
	return article, err
}

// UpdateArticle 更新文章
func (d *Dao) UpdateArticle(ctx context.Context, article model.Article) error {
	return d.db.Save(&article).Error
}

// DeleteArticle 删除文章
func (d *Dao) DeleteArticle(ctx context.Context, id int64) error {
	return d.db.Delete(&model.Article{}, id).Error
}

// CreateArticleCollection 创建文章收藏
func (d *Dao) CreateArticleCollection(ctx context.Context, collection model.ArticleCollection) error {
	return d.db.Create(&collection).Error
}

// GetArticleCollection 获取文章收藏
func (d *Dao) GetArticleCollection(ctx context.Context, id int64) (model.ArticleCollection, error) {
	var collection model.ArticleCollection
	err := d.db.First(&collection, id).Error
	return collection, err
}

// UpdateArticleCollection 更新文章收藏
func (d *Dao) UpdateArticleCollection(ctx context.Context, collection model.ArticleCollection) error {
	return d.db.Save(&collection).Error
}

// DeleteArticleCollection 删除文章收藏
func (d *Dao) DeleteArticleCollection(ctx context.Context, id int64) error {
	return d.db.Delete(&model.ArticleCollection{}, id).Error
}

// CreateArticleLike 创建文章点赞
func (d *Dao) CreateArticleLike(ctx context.Context, like model.ArticleLike) error {
	return d.db.Create(&like).Error
}

// GetArticleLike 获取文章点赞
func (d *Dao) GetArticleLike(ctx context.Context, id int64) (model.ArticleLike, error) {
	var like model.ArticleLike
	err := d.db.First(&like, id).Error
	return like, err
}

// UpdateArticleLike  更新文章点赞
func (d *Dao) UpdateArticleLike(ctx context.Context, like model.ArticleLike) error {
	return d.db.Save(&like).Error
}

// DeleteArticleLike 删除文章点赞
func (d *Dao) DeleteArticleLike(ctx context.Context, id int64) error {
	return d.db.Delete(&model.ArticleLike{}, id).Error
}
