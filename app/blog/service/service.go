package service

import (
	"context"
	"shenyue-gin/app/blog/dao"
)

type Service struct {
	dao *dao.Dao
}

func NewService() (s *Service) {
	s = &Service{
		dao: dao.NewDao(),
	}
	s.NewCorn()
	return s
}

// Ping ping success.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
