package uc

import (
	"PowerX/internal/config"
	"time"

	"github.com/golang-module/carbon/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CustomUseCase struct {
	db *gorm.DB
}

func NewCustomUseCase(conf *config.Config, pxUseCase *PowerXUseCase) (uc *CustomUseCase, clean func()) {

	uc = &CustomUseCase{}

	// 需要打印当时系统的Timezone
	uc.CheckSystemTimeZone()
	return uc, func() {

	}
}

func (uc *CustomUseCase) CheckSystemTimeZone() {
	// 设置 Golang 的 time 包的默认时区
	cst := time.FixedZone("CST", 8*60*60)
	time.Local = cst

	// 设置 Carbon 库的默认时区
	strTimezone := "Asia/Shanghai"
	carbon.SetTimezone(strTimezone)

	// carbon 的timezone
	carbonTimezone := carbon.Now().Timezone()
	logx.Infof("check carbon datetime: timezone- %s\n", carbonTimezone)

	// 输出系统默认时区
	defaultTimezone := time.Now().Location()
	logx.Infof("check system datetime: timezone- %s\n", defaultTimezone.String())
}
