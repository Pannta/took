// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package videoservice

import (
	"context"

	"took/video/video/rpc/types/video"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment                = video.Comment
	CommentActionRequest   = video.CommentActionRequest
	CommentActionResponse  = video.CommentActionResponse
	CommentListRequest     = video.CommentListRequest
	CommentListResponse    = video.CommentListResponse
	FavoriteActionRequest  = video.FavoriteActionRequest
	FavoriteActionResponse = video.FavoriteActionResponse
	FavoriteListRequest    = video.FavoriteListRequest
	FavoriteListResponse   = video.FavoriteListResponse
	FeedRequest            = video.FeedRequest
	FeedResponse           = video.FeedResponse
	PublishListRequest     = video.PublishListRequest
	PublishListResponse    = video.PublishListResponse
	Response               = video.Response
	User                   = video.User
	Video                  = video.Video

	VideoService interface {
		GetVideo(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error)
		PublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListResponse, error)
		GetCommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error)
		CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error)
		FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error)
		FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionResponse, error)
	}

	defaultVideoService struct {
		cli zrpc.Client
	}
)

func NewVideoService(cli zrpc.Client) VideoService {
	return &defaultVideoService{
		cli: cli,
	}
}

func (m *defaultVideoService) GetVideo(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.GetVideo(ctx, in, opts...)
}

func (m *defaultVideoService) PublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.PublishList(ctx, in, opts...)
}

func (m *defaultVideoService) GetCommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.GetCommentList(ctx, in, opts...)
}

func (m *defaultVideoService) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.CommentAction(ctx, in, opts...)
}

func (m *defaultVideoService) FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.FavoriteList(ctx, in, opts...)
}

func (m *defaultVideoService) FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.FavoriteAction(ctx, in, opts...)
}
