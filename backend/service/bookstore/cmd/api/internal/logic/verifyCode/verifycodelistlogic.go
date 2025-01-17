package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
)

type VerifyCodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证码 list
func NewVerifyCodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyCodeListLogic {
	return VerifyCodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCodeListLogic) VerifyCodeList(req types.VerifyCodeListReq) (*types.VerifyCodeReply, error) {
	reqParam := utils.ListReq{}
	reqParam.Page = req.Page
	reqParam.PageSize = req.PageSize
	reqParam.Keyword = req.Keyword
	list, i, err := l.svcCtx.VerifyCodesModel.List(reqParam)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["page"] = req.Page
	data["pageSize"] = req.PageSize
	data["total"] = i
	data["list"] = list

	return nil, errorx.NewCodeError(200, "ok", data)
}
