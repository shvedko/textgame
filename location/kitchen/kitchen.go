package kitchen

import (
	"github.com/pshvedko/textgame/component"
	"github.com/pshvedko/textgame/engine"
)

func init() {
	engine.Register(Kitchen{})
}

type Kitchen struct {
	component.Route
}

func (k Kitchen) Append(string, ...string) bool {
	return false
}

func (k Kitchen) New() engine.Location {
	return &k
}

func (k Kitchen) Name() string {
	return "кухня"
}

func (k Kitchen) Around() string {
	return "ты находишься на кухне, надо собрать рюкзак и идти в универ"
}

func (k Kitchen) Enter() string {
	return "кухня, ничего интересного"
}

func (k Kitchen) Pop(string) bool { return false }
