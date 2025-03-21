package service

import (
	"context"
	"fmt"
	"shenyue-gin/app/shenyue/model"
	"shenyue-gin/app/shenyue/utils"
	"sync"
)

func (s *Service) SaveUser(ctx context.Context, req *model.User) (err error) {
	if req.Username == "" || req.Password == "" {
		return fmt.Errorf("param wrong")
	}
	uid := utils.GenerateUserId()
	user := &model.User{
		Uid:      uid,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	err = s.dao.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return
}

func (s *Service) LoginUser(ctx context.Context, req *model.User) (uid int64, err error) {
	user, err := s.dao.SelectByUsername(ctx, req.Username)
	if err != nil {
		return 0, fmt.Errorf("username not exist")

	}
	if user == nil || user.Password != req.Password {
		return 0, fmt.Errorf("password not exist")
	}
	return user.Uid, nil
}

func (s *Service) FindUserInfo(ctx context.Context, uid int64) (user model.User, err error) {
	user, err = s.dao.SelectByUid(ctx, uid)
	if err != nil {
		return user, fmt.Errorf("username not exist")
	}
	return user, nil
}

func (s *Service) SendUserEmail(ctx context.Context) (err error) {
	res, _ := s.dao.SelectAllEmail(ctx)
	var wg sync.WaitGroup
	// 设置要等待的协程数量为邮箱数量
	wg.Add(len(res))
	// 创建一个channel用于协调邮件发送顺序，确保一封发完再发下一封
	sendChan := make(chan struct{}, 1)
	go func() {
		for _, email := range res {
			// 每个协程执行发送邮件任务前先等待获取发送权限
			sendChan <- struct{}{}
			go func(email string) {
				defer wg.Done()
				defer func() {
					// 发送完邮件后释放发送权限
					<-sendChan
				}()
				if err := s.dao.SendEmail(ctx, email, "测试", "标题"); err != nil {
					// 如果发送邮件出现错误，可以在这里处理，比如记录错误等
					fmt.Printf("发送邮件到 %s 失败: %v\n", email, err)
				}
			}(email)
		}
		fmt.Println("这里")
		// 所有邮件发送任务都已启动，关闭channel
		close(sendChan)
	}()
	return
}
