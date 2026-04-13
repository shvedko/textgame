package street

import (
	"github.com/pshvedko/textgame/component"
	"github.com/pshvedko/textgame/engine"
)

func init() {
	engine.Register(Street{})
}

type Street struct {
	component.Route
}

func (s Street) Append(string, ...string) bool {
	return false
}

func (s Street) New() engine.Location {
	return &s
}

func (s Street) Name() string {
	return "улица"
}

func (s Street) Around() string {
	return "переходи дорогу на зеленый свет"
}

func (s Street) Enter() string {
	return "на улице весна"
}

func (s Street) Pop(string) bool { return false }
