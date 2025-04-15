package scrm

import (
	"PowerX/internal/config"
	"PowerX/internal/uc/powerx/scrm/wechat/wecom"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type SCRMUseCase struct {
	db   *gorm.DB
	kv   *redis.Redis
	Cron *cron.Cron
	// Wechat
	WeCom *wecom.WeComUseCase

	//DTalk
}

func NewSCRMUseCase(db *gorm.DB, conf *config.Config, c *cron.Cron, kv *redis.Redis) *SCRMUseCase {
	work := wecom.NewWeComUseCase(db, conf)

	return &SCRMUseCase{
		db:    db,
		Cron:  c,
		WeCom: work,
	}
}

// Schedule
//
//	@Description:
//	@receiver this
func (uc *SCRMUseCase) Schedule() {

	//sl := []wecom.TimerTypeByte{
	//	// customer message
	//	wecom.AppGroupCustomerMessageTimerTypeByte,
	//	// app group organization message
	//	wecom.AppGroupOrganizationMessageTimerTypeByte,
	//	// app message
	//	wecom.AppMessageTimerTypeByte,
	//}
	//
	//_, _ = uc.Cron.AddFunc(`*/1 * * * *`, func() {
	//	var err error
	//	//  timer 1 minute
	//	unix := time.Now()
	//
	//	for _, val := range sl {
	//		err = uc.WeCom.InvokeTimerMessageGrabUniteSend(val, unix.Unix())
	//		if err != nil {
	//			logx.Info(fmt.Sprintf(`--- [%s] cron.schedule.call.message.%d.error, %v.`, unix.String(), val, err))
	//		}
	//	}
	//
	//})
	//
	//go uc.Cron.Start()

}
