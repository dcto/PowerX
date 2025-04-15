package wecom

import (
	"PowerX/internal/model/scrm/wechat/wecom/app"
	"encoding/json"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	kresp "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/appChat/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/appChat/response"
	"time"
)

// CreateWeComAppGroupRequest
//
//	@Description:
//	@receiver this
//	@param option
//	@return reply
//	@return err
func (uc *WeComUseCase) CreateWeComAppGroupRequest(option *request.RequestAppChatCreate) (reply *response.ResponseAppChatCreate, err error) {

	reply, err = uc.Client.MessageAppChat.Create(uc.ctx, option)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.wechat.app.group.create.error`, reply.ResponseWork)
	}
	if err == nil {
		users, _ := json.Marshal(option.UserList)
		uc.modelWeComApp.group.Action(uc.db, []*app.WeComAppGroup{
			{
				Name:     option.Name,
				Owner:    option.Owner,
				UserList: string(users),
				ChatID:   reply.ChatID,
			},
		})
	}
	return reply, err

}

// UpdateWeComAppGroupRequest
//
//	@Description:
//	@receiver this
//	@param option
//	@return error
func (uc *WeComUseCase) UpdateWeComAppGroupRequest(option *request.RequestAppChatUpdate) error {

	reply, err := uc.Client.MessageAppChat.Update(uc.ctx, option)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.wecom.app.group.update.error`, *reply)
	}
	return err

}

// PullListWeComAppGroupRequest
//
//	@Description:
//	@receiver this
//	@param chatIDs
//	@return replys
//	@return err
func (uc *WeComUseCase) PullListWeComAppGroupRequest(chatIDs ...string) (replys []*power.HashMap, err error) {

	if chatIDs == nil {
		groups := uc.modelWeComApp.group.Query(uc.db)
		chatIDs = app.AdapterGroupSliceChatIDs(func(groups []*app.WeComAppGroup) (ids []string) {
			for _, group := range groups {
				ids = append(ids, group.ChatID)
			}
			return ids
		})(groups)
	}
	for _, id := range chatIDs {

		reply, err := uc.Client.MessageAppChat.Get(uc.ctx, id)
		if err != nil {
			panic(err)
		} else {
			err = uc.help.error(`scrm.pull.wecom.app.group.detail.error`, reply.ResponseWork)
		}
		// update local
		if err == nil {
			group := hash(*reply.ChatInfo).fromHashMapToAppGroup()
			uc.modelWeComApp.group.Action(uc.db, []*app.WeComAppGroup{group})
			replys = append(replys, reply.ChatInfo)
		}
	}

	return replys, err

}

// PushAppWeComGroupMessageArticlesRequest
//
//	@Description:
//	@receiver this
//	@param messages
//	@return *response.ResponseWork
//	@return error
func (uc *WeComUseCase) PushAppWeComGroupMessageArticlesRequest(messages *power.HashMap, sendTime int64) (reply *kresp.ResponseWork, err error) {

	if sendTime > time.Now().Unix() {

		uc.pushTimerMessageToKV(AppGroupOrganizationMessageTimerTypeByte, sendTime, messages)

	} else {
		msg := *messages
		chatIds := msg[`chatIds`].([]interface{})
		for _, id := range chatIds {
			msg[`chatid`] = id
			reply, err = uc.Client.MessageAppChat.Send(uc.ctx, &msg)
			if err != nil {
				panic(err)
			} else {
				err = uc.help.error(`scrm.push.wecom.app.group.message.articles.error`, *reply)
			}
		}

	}

	return reply, err

}

// fromHashMapToAppGroup
//
//	@Description:
//	@receiver this
//	@param obj
func (hash hash) fromHashMapToAppGroup() *app.WeComAppGroup {

	users, _ := json.Marshal(hash[`userlist`])
	return &app.WeComAppGroup{
		Name:     hash[`name`].(string),
		Owner:    hash[`owner`].(string),
		UserList: string(users),
		ChatID:   hash[`chatid`].(string),
	}

}
