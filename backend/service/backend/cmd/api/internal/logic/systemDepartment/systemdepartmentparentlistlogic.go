package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	"backend/service/backend/cmd/api/internal/svc"
	"backend/service/backend/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
)

type SystemDepartmentParentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemDepartmentParentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemDepartmentParentListLogic {
	return SystemDepartmentParentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemDepartmentParentListLogic) SystemDepartmentParentList(req types.SystemDepartmentListReq) (*types.SystemDepartmentReply, error) {
	reqParam := utils.ListReq{}
	reqParam.Page = req.Page
	reqParam.PageSize = req.PageSize
	reqParam.Keyword = req.Keyword
	list, i, err := l.svcCtx.SystemDepartmentsModel.ListParent(reqParam)
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
