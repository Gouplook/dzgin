package dzgin

import "github.com/Gouplook/dzgin/config"

const (
	// DEV is for develop
	DEV = "dev"
	// PROD is for production
	PROD = "prod"
)

var (
	KcConfig  *Config
	AppConfig *kcginAppConfig
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

type kcginAppConfig struct {
	// 内部
	innerConfig config.Configer
}

func newAppConfig(appConfigProvider, appConfigPath string) (*kcginAppConfig, error) {
	ac, err := config.NewConfig(appConfigProvider, appConfigPath)
	if err != nil {
		return nil, err
	}
	return &kcginAppConfig{ac}, nil
}

func (b *kcginAppConfig) Set(key, val string) error {
	if err := b.innerConfig.Set(KcConfig.RunMode+"::"+key, val); err != nil {
		return err
	}
	return b.innerConfig.Set(key, val)
}

func (b *kcginAppConfig) String(key string) string {
	if v := b.innerConfig.String(KcConfig.RunMode + "::" + key); v != "" {
		return v
	}
	return b.innerConfig.String(key)
}

func (b *kcginAppConfig) Strings(key string) []string {
	if v := b.innerConfig.Strings(KcConfig.RunMode + "::" + key); len(v) > 0 {
		return v
	}
	return b.innerConfig.Strings(key)
}

func (b *kcginAppConfig) Int(key string) (int, error) {
	if v, err := b.innerConfig.Int(KcConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Int(key)
}

func (b *kcginAppConfig) Int64(key string) (int64, error) {
	if v, err := b.innerConfig.Int64(KcConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Int64(key)
}

func (b *kcginAppConfig) Bool(key string) (bool, error) {
	if v, err := b.innerConfig.Bool(KcConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Bool(key)
}

func (b *kcginAppConfig) Float(key string) (float64, error) {
	if v, err := b.innerConfig.Float(KcConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Float(key)
}

func (b *kcginAppConfig) DefaultString(key string, defaultVal string) string {
	if v := b.String(key); v != "" {
		return v
	}
	return defaultVal
}

func (b *kcginAppConfig) DefaultStrings(key string, defaultVal []string) []string {
	if v := b.Strings(key); len(v) != 0 {
		return v
	}
	return defaultVal
}

func (b *kcginAppConfig) DefaultInt(key string, defaultVal int) int {
	if v, err := b.Int(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *kcginAppConfig) DefaultInt64(key string, defaultVal int64) int64 {
	if v, err := b.Int64(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *kcginAppConfig) DefaultBool(key string, defaultVal bool) bool {
	if v, err := b.Bool(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *kcginAppConfig) DefaultFloat(key string, defaultVal float64) float64 {
	if v, err := b.Float(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *kcginAppConfig) DIY(key string) (interface{}, error) {
	return b.innerConfig.DIY(key)
}

func (b *kcginAppConfig) GetSection(section string) (map[string]string, error) {
	return b.innerConfig.GetSection(section)
}

func (b *kcginAppConfig) SaveConfigFile(filename string) error {
	return b.innerConfig.SaveConfigFile(filename)
}
