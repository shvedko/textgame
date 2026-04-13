package corridor

import (
	"github.com/pshvedko/textgame/component"
	"github.com/pshvedko/textgame/engine"
)

func init() {
	engine.Register(Corridor{})
}

type Corridor struct {
	component.Route
}

func (c Corridor) Append(string, ...string) bool {
	return false
}

func (c Corridor) New() engine.Location {
	return &c
}

func (c Corridor) Name() string {
	return "коридор"
}

func (c Corridor) Around() string {
	return "не стой в коридоре"
}

func (c Corridor) Enter() string {
	return "ничего интересного"
}

func (c Corridor) Pop(string) bool { return false }
