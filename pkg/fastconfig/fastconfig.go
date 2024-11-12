package fastconfig

import (
	"github.com/spf13/viper"
)

type FastConfig struct {
	file string
}

func NewFastConfig[T interface{}](conf T, options ...Option) T {
	c := &FastConfig{
		file: "./config.yaml",
	}
	for _, o := range options {
		o.Apply(c)
	}
	v := viper.New()
	v.SetConfigFile(c.file)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}
	return conf
}

type Option interface {
	Apply(*FastConfig)
}

type OptionFunc func(*FastConfig)

func (f OptionFunc) Apply(mutex *FastConfig) {
	f(mutex)
}

func WithFile(file string) Option {
	return OptionFunc(func(c *FastConfig) {
		c.file = file
	})
}
