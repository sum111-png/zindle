package handler

import (
	"backend/common/errorx"
	"backend/service/backend/cmd/api/internal/logic/systemRole"
	"backend/service/backend/cmd/api/internal/svc"
	"backend/service/backend/cmd/api/internal/types"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

func MenuParentListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SystemRoleListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v     %v", err, req), ""))
			return
		}

		l := logic.NewMenuParentListLogic(r.Context(), ctx)
		resp, err := l.MenuParentList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
