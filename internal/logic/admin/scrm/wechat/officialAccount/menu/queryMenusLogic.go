package menu

import (
	"PowerX/internal/types/errorx"
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询菜单列表
func NewQueryMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenusLogic {
	return &QueryMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMenusLogic) QueryMenus() (resp *types.QueryMenusReply, err error) {

	res, err := l.svcCtx.PowerX.WechatOA.App.Menu.Get(l.ctx)
	if err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errorx.WithCause(errorx.ErrNotFoundObject, res.ErrMsg)
	}

	return &types.QueryMenusReply{
		Button:    res.Menus.Buttons,
		MatchRule: object.HashMap{},
	}, nil
}
