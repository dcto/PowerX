package wecom

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/groupRobot/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/groupRobot/response"
	"github.com/zeromicro/go-zero/core/logx"
)

// PushWeComBotTextRequest
//
//	@Description:
//	@receiver this
//	@param key
//	@param text
//	@return error
func (uc *WeComUseCase) PushWeComBotTextRequest(key string, text *request.GroupRobotMsgText) error {

	reply, err := uc.Client.GroupRobot.SendText(context.TODO(), key, text)
	logx.Debug(reply, err)
	return err

}

// PushWeComBotFileRequest
//
//	@Description:
//	@receiver this
//	@param key
//	@param file
//	@return error
func (uc *WeComUseCase) PushWeComBotFileRequest(key string, file *request.GroupRobotMsgFile) error {

	reply, err := uc.Client.GroupRobotMessenger.SendFile(context.TODO(), key, file)
	logx.Debug(reply, err)
	return err

}

// PushWeComBotImageRequest
//
//	@Description:
//	@receiver this
//	@param key
//	@param file
//	@return error
func (uc *WeComUseCase) PushWeComBotImageRequest(key string, image *request.GroupRobotMsgImage) error {

	reply, err := uc.Client.GroupRobotMessenger.SendImage(context.TODO(), key, image)
	logx.Debug(reply, err)
	return err

}

// PushWeComBotMarkdownRequest
//
//	@Description:
//	@receiver this
//	@param key
//	@param markdown
//	@return error
func (uc *WeComUseCase) PushWeComBotMarkdownRequest(key string, markdown *request.GroupRobotMsgMarkdown) error {

	reply, err := uc.Client.GroupRobotMessenger.SendMarkdown(context.TODO(), key, markdown)
	logx.Debug(reply, err)
	return err

}

// PushWeComBotArticlesRequest
//
//	@Description:
//	@receiver this
//	@param key
//	@param articles
//	@return *response.ResponseGroupRobotSend
//	@return error
func (uc *WeComUseCase) PushWeComBotArticlesRequest(key string, articles []*request.GroupRobotMsgNewsArticles) (resp *response.ResponseGroupRobotSend, err error) {

	reply, err := uc.Client.GroupRobotMessenger.SendNewsArticles(uc.ctx, key, articles)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.push.wecom.bot.articles.error`, reply.ResponseWork)
	}
	return reply, err

}

// PushWeComBotTemplateRequest
//
//	@Description:
//	@receiver this
//	@param key
//	@param template
//	@return error
func (uc *WeComUseCase) PushWeComBotTemplateRequest(key string, template *request.GroupRobotMsgTemplateCard) error {

	reply, err := uc.Client.GroupRobotMessenger.SendTemplateCard(context.TODO(), key, template)
	logx.Debug(reply, err)
	return err

}
