package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"model_video/core/internal/logic"
	"model_video/core/internal/svc"
	"model_video/core/internal/types"
)

func GetVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FeedRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetVideoListLogic(r.Context(), svcCtx)
		resp, err := l.GetVideoList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
