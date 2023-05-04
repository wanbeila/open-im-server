package push

import (
	"context"
	"sync"

	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/db/cache"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/db/controller"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/db/localcache"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/discoveryregistry"
	pbPush "github.com/OpenIMSDK/Open-IM-Server/pkg/proto/push"
	"google.golang.org/grpc"
)

type pushServer struct {
	pusher *Pusher
}

func Start(client discoveryregistry.SvcDiscoveryRegistry, server *grpc.Server) error {
	rdb, err := cache.NewRedis()
	if err != nil {
		return err
	}
	cacheModel := cache.NewCacheModel(rdb)
	offlinePusher := NewOfflinePusher(cacheModel)
	database := controller.NewPushDatabase(cacheModel)
	pusher := NewPusher(client, offlinePusher, database, localcache.NewGroupLocalCache(client), localcache.NewConversationLocalCache(client))
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		pbPush.RegisterPushMsgServiceServer(server, &pushServer{
			pusher: pusher,
		})
	}()
	go func() {
		defer wg.Done()
		consumer := NewConsumer(pusher)
		consumer.initPrometheus()
		consumer.Start()
	}()
	wg.Wait()
	return nil
}

func (r *pushServer) PushMsg(ctx context.Context, pbData *pbPush.PushMsgReq) (resp *pbPush.PushMsgResp, err error) {
	switch pbData.MsgData.SessionType {
	case constant.SuperGroupChatType:
		err = r.pusher.MsgToSuperGroupUser(ctx, pbData.ConversationID, pbData.MsgData)
	default:
		err = r.pusher.MsgToUser(ctx, pbData.ConversationID, pbData.MsgData)
	}
	if err != nil {
		return nil, err
	}
	return &pbPush.PushMsgResp{}, nil
}

func (r *pushServer) DelUserPushToken(ctx context.Context, req *pbPush.DelUserPushTokenReq) (resp *pbPush.DelUserPushTokenResp, err error) {
	if err = r.pusher.database.DelFcmToken(ctx, req.UserID, int(req.PlatformID)); err != nil {
		return nil, err
	}
	return &pbPush.DelUserPushTokenResp{}, nil
}
