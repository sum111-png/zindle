package handler

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/logic/reader"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)
// 读者 list
func ReaderListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReaderListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewReaderListLogic(r.Context(), ctx)
		resp, err := l.ReaderList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
