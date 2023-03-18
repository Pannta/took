package logic

import (
	"context"

	"took/chat/rpc/types/chat"
	"took/user/model"
	"took/user/rpc/internal/svc"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *user.FriendListReq) (*user.FriendListResp, error) {
	friendList := user.NewFriendList(l.svcCtx.FollowModel.GetFriendById(l.ctx, in.ToUserId))
	
	for i := range friendList {
		isFollow, _ := l.svcCtx.FollowModel.Exist(l.ctx, &model.Follow{
			UserId: friendList[i].Id,
			FanId: in.UserId,
		})
		friendList[i].IsFollow = isFollow
	}

	for i := range friendList {
		var message chat.Message
		has, err := l.svcCtx.CommonModel.Where("(from_user_id=? AND to_user_id=?) OR (from_user_id=? AND to_user_id=?)", 
			in.ToUserId, friendList[i].Id, friendList[i].Id, in.ToUserId).Desc("id").Get(&message)
		if err != nil {
			return nil, err
		}
		if has {
			friendList[i].Message = message.Content
			if message.ToUserId == in.ToUserId {
				friendList[i].MsgType = 0
			} else {
				friendList[i].MsgType = 1
			}
		}
	}

	return &user.FriendListResp{
		StatusCode: 0,
		StatusMsg: "success",
		UserList: friendList,
	}, nil
}
