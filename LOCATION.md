### 🏠 Как добавить новую локацию (на примере «Чердака»)
Чтобы добавить новую локацию, достаточно выполнить 3 шага:

#### 1. Создайте пакет локации

Создайте файл `locations/attic/attic.go`:

```go
package attic

import (
	"game/component"
	"game/engine"
)

type Attic struct {
	component.Route
	component.Items // Если на чердаке будут предметы
}

func (a Attic) New() engine.Location {	return &a }
func (a Attic) Name() string  { return "чердак" }
func (a Attic) Around() string { return "тут пыльно и темно" }
func (a Attic) Enter() string  { return "ты поднялся на чердак" }

func init() {
    // Само-регистрация в реестре движка
    engine.Register(Attic{})
}
```

#### 2. Подключите пакет в main.go

Добавьте одну строку в блок import (через пустой импорт _), чтобы Go выполнил init() нового пакета:
```go
package main

import (
   _ "game/locations/attic" // <-- Новая локация подключена
   _ "game/locations/kitchen"
   // ...
)
```

#### 3. Опишите связи в config.yaml

Добавьте чердак в конфиг и пропишите пути к нему из других комнат:
```yaml
locations:
  - name: "чердак"
    items:
      "в углу": ["сундук", "паутина"]
    paths: ["коридор"]

  - name: "коридор"
    paths: ["кухня", "комната", "улица", "чердак"] # Добавили путь на чердак
```
