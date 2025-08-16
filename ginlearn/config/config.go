package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {

	/**
	    mapstructure 是一个结构体字段标签，用于指定在将映射（map）数据解码到结构体时，该字段与映射中哪个键相对应
	**/
	ServerPort       string `mapstructure:"SERVER_PORT"` //解析的时候，将SERVER_PORT中的值映射到ServerPort中
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	JWTSecret        string `mapstructure:"JWT_SECRET"`
	LogLevel         string `mapstructure:"LOG_LEVEL"`
}

func LoadConfig(path string) (config Config, err error) {

	//转换为绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		return Config{}, fmt.Errorf("路径解析失败：%w", err)
	}

	//构建完整环境文件
	envPath := filepath.Join(absPath, ".env")

	//检查文件是否存在
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("配置文件不存在: %s", envPath)
	}

	// 设置Viper
	viper.SetConfigFile(envPath)

	viper.AutomaticEnv()

	// 读取配置文件
	if err = viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("读取配置文件失败: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("配置解析失败: %w", err)
	}

	return config, nil

}
