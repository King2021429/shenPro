package dao

import "shenyue-gin/app/shenyue/model"

// CreateArticle 创建文章
func (d *Dao) CreateArticle(article model.Article) error {
	return d.db.Create(&article).Error
}

// GetArticle 获取文章
func (d *Dao) GetArticle(id uint) (model.Article, error) {
	var article model.Article
	err := d.db.First(&article, id).Error
	return article, err
}

// UpdateArticle 更新文章
func (d *Dao) UpdateArticle(article model.Article) error {
	return d.db.Save(&article).Error
}

// DeleteArticle 删除文章
func (d *Dao) DeleteArticle(id uint) error {
	return d.db.Delete(&model.Article{}, id).Error
}

// CreateArticleCollection 创建文章收藏
func (d *Dao) CreateArticleCollection(collection model.ArticleCollection) error {
	return d.db.Create(&collection).Error
}

// GetArticleCollection 获取文章收藏
func (d *Dao) GetArticleCollection(id uint) (model.ArticleCollection, error) {
	var collection model.ArticleCollection
	err := d.db.First(&collection, id).Error
	return collection, err
}

// UpdateArticleCollection 更新文章收藏
func (d *Dao) UpdateArticleCollection(collection model.ArticleCollection) error {
	return d.db.Save(&collection).Error
}

// DeleteArticleCollection 删除文章收藏
func (d *Dao) DeleteArticleCollection(id uint) error {
	return d.db.Delete(&model.ArticleCollection{}, id).Error
}

// CreateArticleLike 创建文章点赞
func (d *Dao) CreateArticleLike(like model.ArticleLike) error {
	return d.db.Create(&like).Error
}

// GetArticleLike 获取文章点赞
func (d *Dao) GetArticleLike(id uint) (model.ArticleLike, error) {
	var like model.ArticleLike
	err := d.db.First(&like, id).Error
	return like, err
}

// UpdateArticleLike  更新文章点赞
func (d *Dao) UpdateArticleLike(like model.ArticleLike) error {
	return d.db.Save(&like).Error
}

// DeleteArticleLike 删除文章点赞
func (d *Dao) DeleteArticleLike(id uint) error {
	return d.db.Delete(&model.ArticleLike{}, id).Error
}
