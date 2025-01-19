package auth

import (
	"PowerX/internal/logic/admin/crm/customerDomain/customer"
	customerDomain2 "PowerX/internal/model/crm/customerDomain"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"PowerX/internal/uc/powerx/crm/customerDomain"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoReplyForWeb, err error) {

	vAuthCustomer := l.ctx.Value(customerDomain.AuthCustomerKey)
	authCustomer := vAuthCustomer.(*customerDomain2.Customer)

	customer := customer.TransformCustomerToReply(l.svcCtx, authCustomer)
	customer.AccountId = fmt.Sprintf("%d", customer.Id)
	return &types.GetUserInfoReplyForWeb{
		Customer: customer,
	}, nil
}
