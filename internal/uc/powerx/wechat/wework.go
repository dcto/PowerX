package wechat

import (
	"PowerX/internal/config"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type WeComUseCase struct {
	API *work.Work
	db  *gorm.DB
}

func NewWeComUseCase(db *gorm.DB, conf *config.Config) *WeComUseCase {
	// 初始化企业微信API SDK
	api, err := work.NewWork(&work.UserConfig{
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
		API: api,
		db:  db,
	}
}
