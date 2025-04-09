package leader

import (
	"PowerX/internal/model/crm/customerDomain"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLeadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLeadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLeadLogic {
	return &CreateLeadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLeadLogic) CreateLead(req *types.CreateLeadRequest) (resp *types.CreateLeadReply, err error) {

	lead := &customerDomain.Lead{
		Name:        req.Name,
		Mobile:      req.Mobile,
		Email:       req.Email,
		InviterId:   req.InviterId,
		Source:      req.Source,
		Type:        req.Type,
		IsActivated: req.IsActivated,
	}

	err = l.svcCtx.PowerX.Lead.CreateLead(l.ctx, lead)

	return &types.CreateLeadReply{
		lead.Id,
	}, err

}

func TransformRequestToLead(leadRequest *types.Lead) (mdlLead *customerDomain.Lead) {

	mdlLead = &customerDomain.Lead{
		Name:        leadRequest.Name,
		Mobile:      leadRequest.Mobile,
		Email:       leadRequest.Email,
		InviterId:   leadRequest.InviterId,
		Source:      leadRequest.Source,
		Type:        leadRequest.Type,
		IsActivated: leadRequest.IsActivated,
		ExternalId: customerDomain.ExternalId{
			OpenIdInMiniProgram:           leadRequest.LeadExternalId.OpenIdInMiniProgram,
			OpenIdInWeChatOfficialAccount: leadRequest.LeadExternalId.OpenIdInWeChatOfficialAccount,
			OpenIdInWeCom:                 leadRequest.LeadExternalId.OpenIdInWeCom,
		},
	}

	return mdlLead

}
