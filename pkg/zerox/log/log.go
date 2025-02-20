package log

import (
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type LokiConf struct {
	Enabled    bool              `json:",default=false"` // 是否启用 Loki
	URL        string            `json:",optional"`      // Loki 的 HTTP API 地址
	Labels     map[string]string `json:",optional"`      // Loki 的标签
	RetryCount int               `json:",default=3"`     // 重试次数（如果需要）
}

type LogConf struct {
	Logx logx.LogConf `json:",optional"`
	Loki LokiConf     `json:",optional"`
}

func MustSetupLog(conf *LogConf) {

	// 禁用统计信息（如果需要）
	if !conf.Logx.Stat {
		logx.DisableStat()
	}

	// 创建日志目录
	if conf.Logx.Path == "" {
		conf.Logx.Path = "logs"
	}
	if err := os.MkdirAll(conf.Logx.Path, os.ModePerm); err != nil {
		logx.Errorf("failed to create log directory: %v", err)
		panic(err)
	}

	// 初始化 logx
	logx.MustSetup(conf.Logx)

	// // 追加其他 writer
	// var writers []io.Writer

	// Loki 输出
	if conf.Loki.Enabled {

		lokiWriter := NewLokiWriter(conf.Loki)
		// writers = append(writers, lokiWriter)
		logx.SetWriter(lokiWriter)
	}

	// // 如果有额外 writer，则组合为 MultiWriter
	// if len(writers) > 0 {
	// 	multiWriter := io.MultiWriter(writers...)
	// 	logx.SetWriter(multiWriter)
	// }

}
