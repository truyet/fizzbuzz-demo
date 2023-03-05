package config

import "github.com/zeromicro/go-zero/rest"

type Message struct {
	MultipleOf int32
	Message    string
}
type Config struct {
	rest.RestConf
	Messages []Message
}
