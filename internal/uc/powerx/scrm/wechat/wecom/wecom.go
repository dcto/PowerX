package wecom

import (
	"PowerX/internal/config"
	organization2 "PowerX/internal/model/organization"
	"PowerX/internal/model/scene"
	"PowerX/internal/model/scrm/wechat/wecom/app"
	"PowerX/internal/model/scrm/wechat/wecom/customer"
	organization3 "PowerX/internal/model/scrm/wechat/wecom/organization"
	"PowerX/internal/model/scrm/wechat/wecom/resource"
	tag2 "PowerX/internal/model/scrm/wechat/wecom/tag"
	"PowerX/internal/repository"
	"PowerX/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"sync"
)

type WeComUseCase struct {
	//
	//  help
	//  @Description:
	//
	help
	//
	//  WeComUserRepository
	//  @Description:
	//
	WeComUserRepository *repository.BaseRepository[organization3.WeComUser] `inject:""`
	//
	//  WeComDepartmentRepository
	//  @Description:
	//
	WeComDepartmentRepository *repository.BaseRepository[organization3.WeComDepartment] `inject:""`
	//
	//  WeComTagRepository
	//  @Description:
	//
	WeComTagRepository *repository.BaseRepository[tag2.WeComTag] `inject:""`
	//
	//  WeComGroupTagRepository
	//  @Description:
	//
	WeComTagGroupRepository *repository.BaseRepository[tag2.WeComTagGroup] `inject:""`

	//
	//  db
	//  @Description:
	//
	db *gorm.DB
	//
	//  kv
	//  @Description:
	//
	kv *redis.Redis
	//
	//  wecom
	//  @Description:
	//
	Client *work.Work
	//
	//  ctx
	//  @Description:
	//
	ctx context.Context
	//
	//  gLock
	//  @Description:
	//
	gLock *sync.WaitGroup
	//
	//  modelOrganization
	//  @Description:
	//
	modelOrganization
	//
	//  modelWeComApp
	//  @Description:
	//
	modelWeComApp
	//
	//  modelWeComOrganization
	//  @Description:
	//
	modelWeComOrganization
	//
	//  modelWeComResource
	//  @Description:
	//
	modelWeComResource
	//
	//  modelWeComQRCode
	//  @Description:
	//
	modelWeComQRCode
	//
	//  modelWeComTag
	//  @Description:
	//
	modelWeComTag
	//
	//  modelWeComCustomer
	//  @Description:
	//
	modelWeComCustomer
}

type (
	help              struct{}
	hash              power.HashMap
	modelOrganization struct {
		user       organization2.User
		department organization2.Department
	}
	modelWeComApp struct {
		group app.WeComAppGroup
	}
	modelWeComOrganization struct {
		user       organization3.WeComUser
		department organization3.WeComDepartment
	}
	modelWeComResource struct {
		resource resource.WeComResource
	}
	modelWeComQRCode struct {
		qrcode scene.SceneQRCode
	}
	modelWeComTag struct {
		tag   tag2.WeComTag
		group tag2.WeComTagGroup
	}
	modelWeComCustomer struct {
		follow customer.WeComExternalContactFollow
	}
)

func NewWeComUseCase(db *gorm.DB, conf *config.Config) *WeComUseCase {
	// 初始化企业微信API SDK
	c, err := work.NewWork(&work.UserConfig{
		CorpID:  conf.WeCom.CropId,
		AgentID: conf.WeCom.AgentId,
		Secret:  conf.WeCom.Secret,
		OAuth: work.OAuth{
			Callback: "https://wecom.artisan-cloud.com/callback",
			Scopes:   nil,
		},
		Token:     conf.WeCom.Token,
		AESKey:    conf.WeCom.EncodingAESKey,
		HttpDebug: true,
	})

	if err != nil {
		panic(errors.Wrap(err, "wecom init failed"))
	}

	return &WeComUseCase{
		Client: c,
		db:     db,
		gLock:  new(sync.WaitGroup),
	}
}

var (
	HRedisSCRMGroupMessageKey = `scrm:app:group:%d`
)

type TimerTypeByte int

const (
	//app message
	AppMessageTimerTypeByte TimerTypeByte = iota + 1<<2
	//app group organization message
	AppGroupOrganizationMessageTimerTypeByte
	//app group customer message
	AppGroupCustomerMessageTimerTypeByte
)

// FindManyWeComDepartmentsOption
// @Description:
type FindManyWeComDepartmentsOption struct {
	WeComDepId []int
	Name       string
}

// FindManyWeComUsersOption
// @Description:
type FindManyWeComUsersOption struct {
	UserId                string `json:"wecom_user_id"` //员工唯一ID
	Ids                   []int64
	Names                 []string
	Alias                 []string
	Emails                []string
	Mobile                []string
	OpenUserId            []string
	WeComMainDepartmentId []int64
	Status                []int
	types.PageEmbedOption
}

type (

	// https://developer.work.weixin.qq.com/document/path/90248
	WechatAppRequestBase struct {
		ChatIds []string             `json:"chatIds"`
		MsgType string               `json:"msgtype"`
		Safe    int                  `json:"safe"`
		News    WechatAppRequestNews `json:"news"`
	}

	WechatAppRequestNews struct {
		Article []*WechatAppRequestNewsArticle `json:"articles"`
	}

	WechatAppRequestNewsArticle struct {
		Title       string `json:"title"`       // "领奖通知",
		Description string `json:"description"` // "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>",
		URL         string `json:"url"`         // "URL",
		PicURL      string `json:"picurl"`      // 多"
		AppID       string `json:"appid,omitempty"`
		PagePath    string `json:"pagepath,omitempty"`
	}
)

type (

	//
	//  FindManyWeComCustomerOption
	//  @Description:
	//
	FindManyWeComCustomerOption struct {
		UserId string `json:"user_id"`
		Name   string `gorm:"column:name" json:"name"`
		TagId  string `json:"tag_id"`
		types.PageEmbedOption
	}
)

// decode
//
//	@Description:
//	@receiver self
//	@param str
//	@param body
//	@return help
func (self help) decode(str string, body interface{}) help {

	_ = json.Unmarshal([]byte(str), &body)
	return self
}

// error
//
//	@Description:
//	@receiver self
//	@param ps
//	@param rsp
//	@return err
func (self help) error(ps string, rsp response.ResponseWork) (err error) {

	if rsp.ErrCode > 0 {
		marshal, _ := json.Marshal(rsp)
		err = fmt.Errorf(`%s.%s`, ps, string(marshal))
		logx.Error(err)
	}
	return err

}
