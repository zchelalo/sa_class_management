package bootstrap

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	config   Config
	onceConf sync.Once
)

type Config struct {
	Port                  int32  `mapstructure:"PORT"`
	DBHost                string `mapstructure:"DB_HOST"`
	DBUser                string `mapstructure:"DB_USER"`
	DBPass                string `mapstructure:"DB_PASS"`
	DBName                string `mapstructure:"DB_NAME"`
	DBPort                int32  `mapstructure:"DB_PORT"`
	PaginatorLimitDefault int32  `mapstructure:"PAGINATOR_LIMIT_DEFAULT"`
}

func LoadConfig(path string) (Config, error) {
	var err error
	onceConf.Do(func() {
		viper.AddConfigPath(path)
		viper.SetConfigName("app")
		viper.SetConfigType("env")
		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		err = viper.Unmarshal(&config)
	})
	return config, err
}

func GetConfig() Config {
	return config
}
