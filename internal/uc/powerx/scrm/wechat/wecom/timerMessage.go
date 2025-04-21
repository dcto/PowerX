package wecom

import (
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	creq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/request"
	"strconv"
)

// pushTimerMessageToKV
//
//	@Description:
//	@receiver this
//	@param ttp
//	@param sendTime
//	@param message
func (uc *WeComUseCase) pushTimerMessageToKV(ttp TimerTypeByte, sendTime int64, message interface{}) {

	val := make(map[string]string)
	key := fmt.Sprintf(HRedisSCRMGroupMessageKey, ttp)
	msg, _ := json.Marshal(message)
	val[strconv.Itoa(int(sendTime))] = string(msg)

	err := uc.kv.HmsetCtx(uc.ctx, key, val)
	if err != nil {
		panic(err)
	}

}

// InvokeTimerMessageGrabUniteSend
//
//	@Description: todo
//	@receiver this
//	@param ttp
//	@param sendTime
//	@return error
func (uc *WeComUseCase) InvokeTimerMessageGrabUniteSend(ttp TimerTypeByte, sendTime int64) (err error) {

	key := fmt.Sprintf(HRedisSCRMGroupMessageKey, ttp)

	vals, _ := uc.kv.Hget(key, strconv.Itoa(int(sendTime)))
	if vals == `` {
		return nil
	}
	switch ttp {

	case AppGroupOrganizationMessageTimerTypeByte:
		err = uc.callAppGroupOrganizationMessage(key, sendTime, vals)

	case AppMessageTimerTypeByte:
		err = uc.callAppMessage(key, sendTime, vals)

	case AppGroupCustomerMessageTimerTypeByte:
		err = uc.callCustomerGroupMessage(key, sendTime, vals)
	}

	return err

}

// callAppGroupOrganizationMessage
//
//	@Description:
//	@receiver this
//	@param key
//	@param sendTime
//	@param vals
//	@return error
func (uc *WeComUseCase) callAppGroupOrganizationMessage(key string, sendTime int64, val string) error {

	message := &power.HashMap{}
	err := json.Unmarshal([]byte(val), &message)
	if err == nil {
		_, err = uc.PushAppWeComGroupMessageArticlesRequest(message, sendTime)
		_, err = uc.kv.Hdel(key, strconv.Itoa(int(sendTime)))
	}
	return err
}

// callAppMessage
//
//	@Description:
//	@receiver this
//	@param key
//	@param sendTime
//	@param vals
//	@return error
func (uc *WeComUseCase) callAppMessage(key string, sendTime int64, val string) error {

	message := &request.RequestMessageSendNews{}
	err := json.Unmarshal([]byte(val), &message)
	if err == nil {
		_, err = uc.PushAppWeComMessageArticlesRequest(message, sendTime)
		_, err = uc.kv.Hdel(key, strconv.Itoa(int(sendTime)))
	}

	return err

}

// callCustomerGroupMessage
//
//	@Description:
//	@receiver this
//	@param key
//	@param sendTime
//	@param val
//	@return error
func (uc *WeComUseCase) callCustomerGroupMessage(key string, sendTime int64, val string) error {

	message := &creq.RequestAddMsgTemplate{}
	err := json.Unmarshal([]byte(val), &message)
	if err == nil {
		_, err = uc.PushWoWorkCustomerTemplateRequest(message, sendTime)
		_, err = uc.kv.Hdel(key, strconv.Itoa(int(sendTime)))
	}

	return err

}
