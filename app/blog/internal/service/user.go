package service

import (
	"context"
	"shenyue-gin/app/blog/model/api"
	"shenyue-gin/app/blog/model/db"
)

func (s *Service) SaveUser(ctx context.Context, req *api.UserReq) (err error) {
	user := &db.User{
		Nick:     req.Nick,
		Password: req.Password,
		Email:    req.Email,
		QQ:       req.QQ,
	}
	err = s.dao.Save(ctx, user)
	if err != nil {
		return err
	}
	return
}

func (s *Service) GetUserById(ctx context.Context) (err error) {
	return
}

func (s *Service) GetUserByMobile(ctx context.Context) (err error) {
	return
}
