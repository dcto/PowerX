package wecom

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/response"
	"time"
)

// PushAppWeComMessageArticlesRequest
//
//	@Description:
//	@receiver this
//	@param opt
//	@return *response.ResponseMessageSend
//	@return error
func (uc *WeComUseCase) PushAppWeComMessageArticlesRequest(opt *request.RequestMessageSendNews, sendTime int64) (reply *response.ResponseMessageSend, err error) {

	if sendTime > time.Now().Unix() {

		uc.pushTimerMessageToKV(AppMessageTimerTypeByte, sendTime, opt)

	} else {

		reply, err = uc.Client.Message.SendNews(uc.ctx, opt)
		if err != nil {
			panic(err)

		} else {
			err = uc.help.error(`scrm.push.wecom.app.message.articles.error`, reply.ResponseWork)

		}
	}
	return reply, err

}
