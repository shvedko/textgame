package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"

	"github.com/pshvedko/textgame/engine"

	_ "github.com/pshvedko/textgame/location/corridor"
	_ "github.com/pshvedko/textgame/location/kitchen"
	_ "github.com/pshvedko/textgame/location/room"
	_ "github.com/pshvedko/textgame/location/street"
)

func main() {
	/*
		в этой функции можно ничего не писать,
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
	initGame()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		command, err := r.ReadString('\n')
		if err != nil {
			break
		}
		answer := handleCommand(command)
		fmt.Println(answer)
	}
}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все локации
		если что-то было - оно корректно перезаписывается
	*/
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var cfg engine.Config
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	game, err = engine.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	return game.HandleCommand(command)
}

type Game interface {
	HandleCommand(string) string
}

var game Game
