package dzgin

import (
	"github.com/Gouplook/dzgin/config"
	"github.com/Gouplook/dzgin/utils"
	"os"
	"path/filepath"
)

const (
	// DEV is for develop
	DEV = "dev"
	// PROD is for production
	PROD = "prod"
)

var (
	DzConfig  *Config
	AppConfig *dzginAppConfig
	AppPath   string
	WorkPath  string
)

type Config struct {
	AppName  string // Application name
	RunMode  string // Running Mode: dev | prod
	ConfPath string // Config app.ini path
	ConfName string // Config name

	WebConfig WebConfig
}

type WebConfig struct {
	TemplateLeft  string
	TemplateRight string
}

func init() {
	DzConfig = &Config{
		AppName:  "dzgin",
		RunMode:  "dev",
		ConfPath: "conf",
		ConfName: "config.prod",
		WebConfig: WebConfig{
			TemplateLeft:  "{{",
			TemplateRight: "}}",
		},
	}

	// 获取环境变量的中设置配置信息
	if os.Getenv("DZGIN_APPNAME") != "" {
		DzConfig.AppName = os.Getenv("DZGIN_APPNAME")
	}
	if os.Getenv("DZGIN_RUNMODE") != "" {
		DzConfig.RunMode = os.Getenv("DZGIN_RUNMODE")
	}
	if os.Getenv("DZGIN_CONFPATH") != "" {
		DzConfig.ConfPath = os.Getenv("DZGIN_CONFPATH")
	}
	if os.Getenv("DZGIN_CONFNAME") != "" {
		DzConfig.ConfName = os.Getenv("DZGIN_CONFNAME")
	}

	var err error
	WorkPath, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	AppPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(AppPath, DzConfig.ConfPath, DzConfig.ConfName)
	if !utils.FileExists(appConfigPath) {
		appConfigPath = filepath.Join(WorkPath, DzConfig.ConfPath, DzConfig.ConfName)
		if !utils.FileExists(appConfigPath) {
			AppConfig = &dzginAppConfig{innerConfig: config.NewFakeConfig()}
		}
	}
	conf, err := newAppConfig("ini", appConfigPath)
	if err != nil {
		panic(err)
	}
	if rpcname := conf.String("rpcname"); rpcname != "" {
		DzConfig.AppName = rpcname
	}

	if runmode := conf.String("runmode"); runmode != "" {
		DzConfig.RunMode = runmode
	}

	AppConfig = conf
}

type dzginAppConfig struct {
	// 内部
	innerConfig config.Configer
}

func newAppConfig(appConfigProvider, appConfigPath string) (*dzginAppConfig, error) {
	ac, err := config.NewConfig(appConfigProvider, appConfigPath)
	if err != nil {
		return nil, err
	}
	return &dzginAppConfig{ac}, nil
}

func (b *dzginAppConfig) Set(key, val string) error {
	if err := b.innerConfig.Set(DzConfig.RunMode+"::"+key, val); err != nil {
		return err
	}
	return b.innerConfig.Set(key, val)
}

func (b *dzginAppConfig) String(key string) string {
	if v := b.innerConfig.String(DzConfig.RunMode + "::" + key); v != "" {
		return v
	}
	return b.innerConfig.String(key)
}

func (b *dzginAppConfig) Strings(key string) []string {
	if v := b.innerConfig.Strings(DzConfig.RunMode + "::" + key); len(v) > 0 {
		return v
	}
	return b.innerConfig.Strings(key)
}

func (b *dzginAppConfig) Int(key string) (int, error) {
	if v, err := b.innerConfig.Int(DzConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Int(key)
}

func (b *dzginAppConfig) Int64(key string) (int64, error) {
	if v, err := b.innerConfig.Int64(DzConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Int64(key)
}

func (b *dzginAppConfig) Bool(key string) (bool, error) {
	if v, err := b.innerConfig.Bool(DzConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Bool(key)
}

func (b *dzginAppConfig) Float(key string) (float64, error) {
	if v, err := b.innerConfig.Float(DzConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Float(key)
}

func (b *dzginAppConfig) DefaultString(key string, defaultVal string) string {
	if v := b.String(key); v != "" {
		return v
	}
	return defaultVal
}

func (b *dzginAppConfig) DefaultStrings(key string, defaultVal []string) []string {
	if v := b.Strings(key); len(v) != 0 {
		return v
	}
	return defaultVal
}

func (b *dzginAppConfig) DefaultInt(key string, defaultVal int) int {
	if v, err := b.Int(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *dzginAppConfig) DefaultInt64(key string, defaultVal int64) int64 {
	if v, err := b.Int64(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *dzginAppConfig) DefaultBool(key string, defaultVal bool) bool {
	if v, err := b.Bool(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *dzginAppConfig) DefaultFloat(key string, defaultVal float64) float64 {
	if v, err := b.Float(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *dzginAppConfig) DIY(key string) (interface{}, error) {
	return b.innerConfig.DIY(key)
}

func (b *dzginAppConfig) GetSection(section string) (map[string]string, error) {
	return b.innerConfig.GetSection(section)
}

func (b *dzginAppConfig) SaveConfigFile(filename string) error {
	return b.innerConfig.SaveConfigFile(filename)
}
