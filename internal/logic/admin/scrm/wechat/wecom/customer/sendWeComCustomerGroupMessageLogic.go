package customer

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendWeComCustomerGroupMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 客户群发信息
func NewSendWeComCustomerGroupMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendWeComCustomerGroupMessageLogic {
	return &SendWeComCustomerGroupMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendWeComCustomerGroupMessageLogic) SendWeComCustomerGroupMessage(req *types.WeComAddMsgTemplateRequest) (resp *types.WeComAddMsgTemplateResponse, err error) {
	template, err := l.svcCtx.PowerX.SCRM.WeCom.PushWoWorkCustomerTemplateRequest(l.OPT(req), req.SendTime)

	return &types.WeComAddMsgTemplateResponse{
		FailList: template.FailList,
		MsgId:    template.MsgID,
	}, err

}

// OPT
//
//	@Description:
//	@receiver message
//	@param opt
//	@return *request.RequestAddMsgTemplate
func (l *SendWeComCustomerGroupMessageLogic) OPT(opt *types.WeComAddMsgTemplateRequest) *request.RequestAddMsgTemplate {
	option := &request.RequestAddMsgTemplate{
		ChatType:       opt.ChatType,
		ExternalUserID: opt.ExternalUserId,
		Sender:         opt.Sender,
		Text:           l.text(opt.Text),
		Attachments:    l.attachments(opt.Attachments),
	}
	return option
}

// text
//
//	@Description:
//	@receiver message
//	@param msg
//	@return *request.TextOfMessage
func (l *SendWeComCustomerGroupMessageLogic) text(msg *types.WeComTextOfMessage) *request.TextOfMessage {

	if msg != nil {
		return &request.TextOfMessage{msg.Content}
	}
	return nil
}

// attachments
//
//	@Description:
//	@receiver message
//	@param contents
//	@return attachments
func (l *SendWeComCustomerGroupMessageLogic) attachments(contents []types.Content) (attachmentsMessageTemplateInterface []request.MessageTemplateInterface) {

	if len(contents) > 0 {
		attr := new(attachment)
		for _, content := range contents {
			attr.MsgType = content.Link.MsgType
			attr.Link = l.attachmentLink(&content.Link)
			attachmentsMessageTemplateInterface = append(attachmentsMessageTemplateInterface, attr)
		}
	}
	return attachmentsMessageTemplateInterface
}

// @Description:
// @receiver message
// @param image
// @return *request.Image
func (l *SendWeComCustomerGroupMessageLogic) attachmentImage(image *types.Image) *request.Image {

	if image != nil {
		return &request.Image{
			MediaID: image.MediaId,
			PicURL:  image.PicUrl,
		}
	}
	return nil

}

// attachmentLink
//
//	@Description:
//	@receiver message
//	@param link
//	@return *request.Link
func (l *SendWeComCustomerGroupMessageLogic) attachmentLink(link *types.Link) *request.Link {

	if link != nil {
		return &request.Link{
			Title:  link.Title,
			PicURL: link.PicURL,
			Desc:   link.Desc,
			URL:    link.URL,
		}
	}
	return nil

}

type attachment request.Attachment

func (attachment *attachment) GetMsgType() string {

	return ``
}
