package room

import (
	"github.com/pshvedko/textgame/component"
	"github.com/pshvedko/textgame/engine"
)

func init() {
	engine.Register(Room{})
}

type Room struct {
	component.Route
	component.Items
}

func (r Room) New() engine.Location {
	return &r
}

func (r Room) Name() string {
	return "комната"
}

func (r Room) Around() string {
	if !r.Empty() {
		return r.Items.String()
	}
	return "пустая комната"
}

func (r Room) Enter() string {
	return "ты в своей комнате"
}
