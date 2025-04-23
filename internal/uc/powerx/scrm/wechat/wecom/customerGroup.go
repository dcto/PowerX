package wecom

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat/response"
	creq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"
	crsp "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/response"
	"time"
)

// PullListWeComCustomerGroupRequest
//
//	@Description:
//	@receiver this
//	@param opt
//	@return list
//	@return error
func (uc *WeComUseCase) PullListWeComCustomerGroupRequest(opt *request.RequestGroupChatList) (list []*response.ResponseGroupChatGet, err error) {

	reply, err := uc.Client.ExternalContactGroupChat.List(uc.ctx, opt)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.wecom.list.customer.group.error`, reply.ResponseWork)
	}

	if reply != nil {
		uc.gLock.Add(len(reply.GroupChatList))
		for _, chat := range reply.GroupChatList {
			go func(chatID string) {
				get, _ := uc.Client.ExternalContactGroupChat.Get(uc.ctx, chatID, 1)
				if get.ErrCode == 0 {
					list = append(list, get)
				}
				uc.gLock.Done()
			}(chat.ChatID)
		}
		uc.gLock.Wait()
	}
	return list, err

}

// PushWoWorkCustomerTemplateRequest
//
//	@Description:
//	@receiver this
//	@param opt
//	@return *crsp.ResponseAddMessageTemplate
//	@return error
func (uc *WeComUseCase) PushWoWorkCustomerTemplateRequest(opt *creq.RequestAddMsgTemplate, sendTime int64) (*crsp.ResponseAddMessageTemplate, error) {

	if sendTime > time.Now().Unix() {

		uc.pushTimerMessageToKV(AppGroupCustomerMessageTimerTypeByte, sendTime, opt)

	}
	reply, err := uc.Client.ExternalContactMessageTemplate.AddMsgTemplate(uc.ctx, opt)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.push.wecom.customer.message.error.`, reply.ResponseWork)
	}
	return reply, err

}
