package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"shenyue-gin/app/shenyue/model"
	"time"
)

func (s *Service) AIChatStart(ctx context.Context, Uid int64) (resp *model.ConversationStartResp, err error) {
	resp = &model.ConversationStartResp{}
	// 使用时间戳作为随机数种子，确保每次运行生成的随机数不同
	rand.Seed(time.Now().UnixNano())
	// 生成0到100000之间的随机数
	randomNumber := rand.Intn(100001)
	history := []model.Message{
		{
			Role:    "system",
			Content: "你是二次元妹子，活泼可爱好动，名字叫江户川神月",
		},
	}
	// 将结构体切片转换为JSON字符串
	jsonData, err := json.Marshal(history)
	if err != nil {
		fmt.Println("转换为JSON字符串失败:", err)
		return
	}
	aIChat := &model.AIChat{
		Uid:                 Uid,
		ConversationId:      int64(randomNumber),
		ConversationContent: string(jsonData),
	}
	err = s.dao.CreateAIChat(aIChat)
	if err != nil {
		fmt.Println(err)
		return
	}
	//err = s.dao.RcSetConversation(ctx, int64(randomNumber), Uid, string(jsonData))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	resp.ConversationId = int64(randomNumber)
	return resp, nil
}

func (s *Service) AIChatSendMsg(ctx context.Context, req *model.ConversationSendMsgReq) (resp *model.ConversationSendMsgResp, err error) {
	resp = &model.ConversationSendMsgResp{}
	//value, err := s.dao.RcGetConversation(ctx, req.Uid, req.ConversationId)
	//if err != nil {
	//	return resp, err
	//}
	aiChat, err := s.dao.GetAIChatByUidAndConversationId(req.Uid, req.ConversationId)
	if err != nil {
		return resp, err
	}
	// 将JSON字符串转换回结构体切片
	var newHistory []model.Message
	err = json.Unmarshal([]byte(aiChat.ConversationContent), &newHistory)
	if err != nil {
		fmt.Println("从JSON字符串转换回结构体失败:", err)
		return resp, err
	}
	res := s.dao.AIChat(req.Content, &newHistory)
	// 将结构体切片转换为JSON字符串
	jsonData, err := json.Marshal(newHistory)
	if err != nil {
		fmt.Println("转换为JSON字符串失败:", err)
		return
	}
	err = s.dao.UpdateConversationContent(req.Uid, req.ConversationId, string(jsonData))
	if err != nil {
		return resp, err
	}
	//err = s.dao.RcSetConversation(ctx, req.ConversationId, req.Uid, string(jsonData))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	resp.Reply = res
	return resp, nil
}

func (s *Service) AIChatDelete(ctx context.Context, req *model.ConversationDeleteReq) (resp *model.ConversationDeleteResp, err error) {
	resp = &model.ConversationDeleteResp{}
	err = s.dao.DeleteAIChat(req.Uid, req.ConversationId)
	if err != nil {
		fmt.Println(err)
		return
	}
	//err = s.dao.RcDelConversation(ctx, req.ConversationId, req.Uid)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	return resp, nil
}

func (s *Service) AIChatList(ctx context.Context, req *model.ConversationListReq) (resp *model.ConversationListResp, err error) {
	resp = &model.ConversationListResp{}
	list, err := s.dao.GetAIChatList(req.Uid)
	if err != nil {
		return resp, err
	}
	resp.List = list
	return resp, nil
}
