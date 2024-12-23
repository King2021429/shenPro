package service

import (
	"fmt"
	"shenyue-gin/app/shenyue/model"
)

func (s *Service) AIChat(chatReq string) (response string) {
	history := []model.Message{
		{
			Role:    "system",
			Content: "你是二次元妹子，活泼可爱好动，名字叫江户川神月",
		},
	}

	firstResult := s.dao.AIChat(chatReq, &history)
	fmt.Println(firstResult)
	secondResult := s.dao.AIChat(chatReq, &history)
	fmt.Println(secondResult)
	thirdResult := s.dao.AIChat(chatReq, &history)
	fmt.Println(thirdResult)
	return thirdResult
}
