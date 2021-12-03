package config

/**
本地文件配置
参考 beego
*/

type VAppConfig struct {
	innerConfig Configer
	runMode     string
}




// NewConfig adapterName is ini/json/xml/yaml.
// filename is the config file path.
func NewAppConfig(adapterName, filename string) (*VAppConfig, error) {
	ac, err := NewConfig(adapterName, filename)
	if err != nil {
		return nil, err
	}
	return &VAppConfig{ac, ""}, nil
}

func (b *VAppConfig) SetRunMode(runMode string) {
	b.runMode = runMode
}

func (b *VAppConfig) Set(key, val string) error {
	if err := b.innerConfig.Set(b.runMode+"::"+key, val); err != nil {
		return err
	}
	return b.innerConfig.Set(key, val)
}

func (b *VAppConfig) String(key string) string {
	if v := b.innerConfig.String(b.runMode + "::" + key); v != "" {
		return v
	}
	return b.innerConfig.String(key)
}

func (b *VAppConfig) Strings(key string) []string {
	if v := b.innerConfig.Strings(b.runMode + "::" + key); len(v) > 0 {
		return v
	}
	return b.innerConfig.Strings(key)
}

func (b *VAppConfig) Int(key string) (int, error) {
	if v, err := b.innerConfig.Int(b.runMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Int(key)
}

func (b *VAppConfig) Int64(key string) (int64, error) {
	if v, err := b.innerConfig.Int64(b.runMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Int64(key)
}

func (b *VAppConfig) Bool(key string) (bool, error) {
	if v, err := b.innerConfig.Bool(b.runMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Bool(key)
}

func (b *VAppConfig) Float(key string) (float64, error) {
	if v, err := b.innerConfig.Float(b.runMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Float(key)
}

func (b *VAppConfig) DefaultString(key string, defaultVal string) string {
	if v := b.String(key); v != "" {
		return v
	}
	return defaultVal
}

func (b *VAppConfig) DefaultStrings(key string, defaultVal []string) []string {
	if v := b.Strings(key); len(v) != 0 {
		return v
	}
	return defaultVal
}

func (b *VAppConfig) DefaultInt(key string, defaultVal int) int {
	if v, err := b.Int(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *VAppConfig) DefaultInt64(key string, defaultVal int64) int64 {
	if v, err := b.Int64(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *VAppConfig) DefaultBool(key string, defaultVal bool) bool {
	if v, err := b.Bool(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *VAppConfig) DefaultFloat(key string, defaultVal float64) float64 {
	if v, err := b.Float(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *VAppConfig) DIY(key string) (interface{}, error) {
	return b.innerConfig.DIY(key)
}

func (b *VAppConfig) GetSection(section string) (map[string]string, error) {
	return b.innerConfig.GetSection(section)
}

func (b *VAppConfig) SaveConfigFile(filename string) error {
	return b.innerConfig.SaveConfigFile(filename)
}
