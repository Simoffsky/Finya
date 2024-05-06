package main

import (
	"finance-bot/internal/config"
	"finance-bot/internal/log"
)

type Default interface {
	Test()
	Print()
}

type DefaultImpl struct {
	a int
}

func (d *DefaultImpl) Test() {
	d.a = d.a + 1
}

func (d DefaultImpl) Print() {
	println(d.a)
}

func getDefault() Default {
	return &DefaultImpl{a: 1}
}

func main() {
	config := config.NewEnvConfig()

	logger := log.NewDefaultLogger(
		log.LevelFromString(config.LogLevel),
	).WithTimePrefix()

}
