package component

import "github.com/pshvedko/textgame/engine"

type Route []engine.Location

func (r *Route) Find(to string) (engine.Location, bool) {
	for _, location := range *r {
		if to == location.Name() {
			return location, true
		}
	}
	return nil, false
}

func (r *Route) Path() (paths []string) {
	for _, location := range *r {
		paths = append(paths, location.Name())
	}
	return
}

func (r *Route) Link(locations ...engine.Location) {
	*r = append(*r, locations...)
}

type Alias struct {
	engine.Location
	string
}

func (a Alias) Name() string {
	return a.string
}

func (r *Route) Same(to, alias string) bool {
	for i, location := range *r {
		if to == location.Name() {
			(*r)[i] = Alias{Location: location, string: alias}
			return true
		}
	}
	return false
}
