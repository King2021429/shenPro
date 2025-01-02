package service

import (
	"context"
	cfg "shenyue-gin/app/shenyue/configs"
	"shenyue-gin/app/shenyue/dao"
)

type Service struct {
	dao *dao.Dao
	cfg *cfg.Config
}

func NewService(c *cfg.Config) (s *Service) {
	s = &Service{
		dao: dao.NewDao(c),
		cfg: c,
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
