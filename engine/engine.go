package engine

import (
	"fmt"
	"strings"
)

type Router interface {
	Path() []string
	Find(string) (Location, bool)
	Link(...Location)
	Same(string, string) bool
}

type Keeper interface {
	Pop(string) bool
	Append(at string, items ...string) bool
}

type Location interface {
	Name() string
	Around() string
	Enter() string
	Router
	Keeper
}

type Inventory []string

type Person struct {
	Location
	Inventory
}

func (i *Inventory) Put(item string) {
	*i = append(*i, item)
}

type Game struct {
	Person
}

func (g *Game) HandleCommand(command string) string {
	path := strings.Fields(command)
	switch {
	case len(path) == 1 && path[0] == "осмотреться":
		return fmt.Sprintf("%s. можно пройти - %s", g.Around(), strings.Join(g.Path(), ", "))
	case len(path) == 2 && path[0] == "идти":
		location, ok := g.Find(path[1])
		if ok {
			g.Location = location
			return fmt.Sprintf("%s. можно пройти - %s", g.Enter(), strings.Join(g.Path(), ", "))
		}
		return fmt.Sprintf("нет пути в %s", path[1])
	case len(path) == 2 && path[0] == "взять":
		if g.Pop(path[1]) {
			g.Put(path[1])
			return fmt.Sprintf("предмет добавлен в инвентарь: %s", path[1])
		}
		return "нет такого"
	default:
		return "неизвестная команда"
	}
}

type Factory interface {
	Name() string
	New() Location
}

var Registry = map[string]Factory{}

type ErrLocationNotExists string

func (e ErrLocationNotExists) Error() string {
	return fmt.Sprintf("location not exists: %q", string(e))
}

type ErrLocationWithNoItem string

func (e ErrLocationWithNoItem) Error() string {
	return fmt.Sprintf("location does not contain any items: %q", string(e))
}

func Register(factory Factory) {
	if _, exists := Registry[factory.Name()]; exists {
		panic(factory.Name())
	}
	Registry[factory.Name()] = factory
}

type Config struct {
	From      string `yaml:"from"`
	Locations []struct {
		Name    string              `yaml:"name"`
		Paths   []string            `yaml:"paths"`
		Items   map[string][]string `yaml:"items"`
		Aliases map[string]string   `yaml:"aliases"`
	} `yaml:"locations"`
}

func New(c Config) (*Game, error) {
	locations := map[string]Location{}
	for _, location := range c.Locations {
		factory, ok := Registry[location.Name]
		if !ok {
			return nil, fmt.Errorf("registry: %w", ErrLocationNotExists(location.Name))
		}
		_, ok = locations[location.Name]
		if !ok {
			locations[location.Name] = factory.New()
		}
		for at, items := range location.Items {
			if !locations[location.Name].Append(at, items...) {
				return nil, ErrLocationWithNoItem(location.Name)
			}
		}
		for _, path := range location.Paths {
			_, ok = locations[path]
			if !ok {
				factory, ok = Registry[path]
				if !ok {
					return nil, fmt.Errorf("registry: %w", ErrLocationNotExists(path))
				}
				locations[path] = factory.New()
			}
			locations[location.Name].Link(locations[path])
		}
		for path, to := range location.Aliases {
			if !locations[location.Name].Same(to, path) {
				return nil, fmt.Errorf("alias: %w", ErrLocationNotExists(to))
			}
		}
	}
	from, ok := locations[c.From]
	if !ok {
		return nil, fmt.Errorf("from: %w", ErrLocationNotExists(c.From))
	}
	return &Game{Person: Person{Location: from}}, nil
}
