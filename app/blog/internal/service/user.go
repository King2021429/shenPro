package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"shenyue-gin/app/blog/model/api"
	"shenyue-gin/app/blog/model/db"
)

func (s *Service) SaveUser(ctx context.Context, req *api.UserReq) (err error) {
	user := &db.User{
		Nick:     req.Nick,
		Password: req.Password,
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
