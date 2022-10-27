package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type SenderConfiger interface {
	SenderConfig() *SenderConfig
}

type SenderConfig struct {
	Endpoint string `fig:"endpoint"`
}

func NewSenderConfiger(getter kv.Getter) SenderConfiger {
	return &senderConfig{
		getter: getter,
	}
}

type senderConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *senderConfig) SenderConfig() *SenderConfig {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "eth_ws")
		config := AuthConfig{}
		err := figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &config
	}).(*SenderConfig)
}
