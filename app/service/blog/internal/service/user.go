package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"shenyue-gin/app/service/blog/model/api"
	"shenyue-gin/app/service/blog/model/db"
)

func (s *Service) SaveUser(ctx context.Context, req *api.User) (err error) {
	user := &db.User{
		Nick:     req.Nick,
		Password: req.Password,
		Email:    req.Email,
		QQ:       req.QQ,
		Wechat:   req.Wechat,
	}
	err = s.dao.Save(ctx, user)
	if err != nil {
		return err
	}
	return
}

func (s *Service) GetUserById(c *gin.Context) {
	return
}
