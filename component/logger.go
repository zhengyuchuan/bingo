package component

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	logger *zap.Logger
)

func InitLog() {

	logFilePath := getLogFilePath() + getLogFileName()
	hook := &lumberjack.Logger{
		Filename:   logFilePath, // 日志文件路径
		MaxSize:    1024,        // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,           // 日志文件最多保存多少个备份
		MaxAge:     28,          // 文件最多保存多少天
		Compress:   true,        // 是否压缩
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),                       // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)), // 打印到控制台和文件
		zap.InfoLevel, // 日志级别 TODO：绑定viper配置
	)

	logger = zap.New(core)
}

func GetLogger() *zap.Logger {
	return logger
}

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", viper.GetString("log.runtime_root_path"), viper.GetString("log.log_save_path"))
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		viper.GetString("log.log_save_name"),
		time.Now().Format(viper.GetString("log.time_format")),
		viper.GetString("log.log_file_ext"),
	)
}
