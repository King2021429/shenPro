package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"shenyue-gin/app/service/vrgo/internal/model"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	pb "shenyue-gin/app/service/vrgo/api"
	"shenyue-gin/app/service/vrgo/internal/dao"
)

// Service service.
type Service struct {
	dao *dao.Dao
}

// New new a service and return.
func New() *Service {
	d := dao.New()
	s := &Service{
		dao: d,
	}
	return s
}
func (s *Service) Format(numStr string) (num int) {
	if len(numStr) == 0 {
		return 1
	}
	num, err := strconv.Atoi(numStr)
	if err != nil {
		// 如果转换失败，打印错误信息
		fmt.Println("转换错误:", err)
		num = 1
	} else {
		// 如果转换成功，打印结果
		fmt.Println("转换结果:", num)
	}
	return
}

// SayHelloURL bm demo func.
func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (reply *pb.InfoResp, err error) {
	reply = &pb.InfoResp{}

	key := fmt.Sprintf("second_%d.jpg", req.Time)
	// 解析JSON字符串到ImageMap结构体
	var imageMap model.ImageMap
	err = json.Unmarshal([]byte(model.JsonBlue), &imageMap)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 解析JSON字符串到ImageMap结构体
	var imageMap2 model.ImageMap
	err = json.Unmarshal([]byte(model.JsonRed), &imageMap2)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	imageMapNow, ok1 := imageMap[key]
	imageMap2Now, ok2 := imageMap2[key]

	reply.TeamNow = &pb.TeamNow{
		BlueName: "TES",
		RedName:  "BLG",
	}
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 初始化红蓝两方队伍
	redTeam := model.Team{Name: "BLG", Heroes: [5]model.Hero{{"leona", 0}, {"tristana", 0}, {"ezreal", 0}, {"kennen", 0}, {"maokai", 0}}}
	blueTeam := model.Team{Name: "TES", Heroes: [5]model.Hero{{"senna", 0}, {"azir", 0}, {"leesin", 0}, {"jayce", 0}, {"ornn", 0}}}
	// 模拟经济变化
	for i := 0; i < 5; i++ {
		redTeam.Heroes[i].Wealth = 500 + (100+rand.Intn(10))*int(req.Time) // 假设每个角色每分钟经济增长100-150
		blueTeam.Heroes[i].Wealth = 500 + (100+rand.Intn(10))*int(req.Time)
	}

	for _, hero := range redTeam.Heroes {
		x := 1
		y := 1
		if ok1 {
			if pos, ok := imageMapNow[hero.Name]; ok {
				x = pos[0]
				y = pos[1]
			}
		}

		item := pb.PlayerNow{
			Coler:    "red",
			Name:     hero.Name,
			Economic: int64(hero.Wealth),
			X:        int64(x),
			Y:        int64(y),
		}
		reply.PlayerNow = append(reply.PlayerNow, &item)
	}

	// 输出蓝方队伍的经济情况
	for _, hero := range blueTeam.Heroes {
		x := 1
		y := 1
		if ok2 {
			if pos, ok := imageMap2Now[hero.Name]; ok {
				x = pos[0]
				y = pos[1]
			}
		}

		item := pb.PlayerNow{
			Coler:    "blue",
			Name:     hero.Name,
			Economic: int64(hero.Wealth),
			X:        int64(x),
			Y:        int64(y),
		}
		reply.PlayerNow = append(reply.PlayerNow, &item)
	}

	fmt.Println(reply)
	reply.Time = req.Time

	if req.Time == 20 || req.Time == 50 {
		reply.Status = pb.Status_Status2
	} else {
		reply.Status = pb.Status_Status1
	}
	return reply, nil
}

func (s *Service) MatchList(ctx context.Context, req *pb.MatchListReq) (reply *pb.MatchListResp, err error) {
	reply = &pb.MatchListResp{}
	err = json.Unmarshal([]byte(model.MatchInfo), &reply)
	if err != nil {
		log.Println(err)
	}
	return
}

func (s *Service) GradeInfo(ctx context.Context, req *pb.GradeInfoReq) (reply *pb.GradeInfoResp, err error) {
	reply = &pb.GradeInfoResp{}
	err = json.Unmarshal([]byte(model.GradeInfo1), &reply)
	if err != nil {
		log.Println(err)
	}
	return
}

func (s *Service) Economic(ctx context.Context, req *pb.EconomicReq) (reply *pb.EconomicResp, err error) {
	reply = &pb.EconomicResp{}
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	// 初始化红蓝两方队伍
	redTeam := model.Team{Name: "Red Team", Heroes: [5]model.Hero{{"Red1", 0}, {"Red2", 0}, {"Red3", 0}, {"Red4", 0}, {"Red5", 0}}}
	blueTeam := model.Team{Name: "Blue Team", Heroes: [5]model.Hero{{"Blue1", 0}, {"Blue2", 0}, {"Blue3", 0}, {"Blue4", 0}, {"Blue5", 0}}}
	// 模拟经济变化
	for i := 0; i < 5; i++ {
		redTeam.Heroes[i].Wealth = 500 + (100+rand.Intn(50))*int(req.Time) // 假设每个角色每分钟经济增长100-150
		blueTeam.Heroes[i].Wealth = 500 + (100+rand.Intn(50))*int(req.Time)
	}

	for _, hero := range redTeam.Heroes {
		item := pb.EconomicItem{
			Coler: "red",
			Name:  hero.Name,
			Value: int64(hero.Wealth),
		}
		reply.Economic = append(reply.Economic, &item)
	}

	// 输出蓝方队伍的经济情况
	for _, hero := range blueTeam.Heroes {
		item := pb.EconomicItem{
			Coler: "blue",
			Name:  hero.Name,
			Value: int64(hero.Wealth),
		}
		reply.Economic = append(reply.Economic, &item)
	}
	return
}

func (s *Service) Webhook(ctx context.Context, req *pb.WebhookReq) (reply *pb.WebhookResp, err error) {
	reply = &pb.WebhookResp{}
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
