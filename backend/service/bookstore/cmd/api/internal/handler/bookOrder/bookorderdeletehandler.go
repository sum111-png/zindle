package handler

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/logic/bookOrder"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

// 图书订单 delete
func BookOrderDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BookOrderDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewBookOrderDeleteLogic(r.Context(), ctx)
		resp, err := l.BookOrderDelete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
