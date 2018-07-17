package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/yanzay/tbot"
)

func main() {
	bot, err := tbot.NewServer(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	bot.HandleFunc("/griller", GrillerHandler)
	bot.HandleFunc("/dessert", DessertHandler)
	bot.HandleFunc("/pay",PayHandler)
	bot.HandleFunc("/time",TimeHandler)
	//bot.HandleFunc("/menu", MenuHandler)
	//bot.HandleFunc("/where",WhereHandler)
	err = bot.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func GrillerHandler(m *tbot.Message){
	message := `GRIGLIATORI\n (anche alternandosi, a necessità, sapete a chi chiedere il cambio)\n
	* Silvano\n
	* Alberto d'este\n
	* Marco Falcier\n
	* Julian (Vegetarian grill master)\n
	* Daniele Scarpa\n
	* Daniele Crosera\n
	* Simone Pavlovich\n`
	m.Reply(message)
}

func DessertHandler(m *tbot.Message){
	message := `CHI PORTERA' QUALCOSA DI DOLCE:\n
				* Marco Favin\n
				* Massimo F\n
				* Simone\n
				* Julian\n
				* Alice\n`
	m.Reply(message)
}

func TimeHandler(m *tbot.Message){
	m.Reply("Chi aiuta a preparare dalle 10.00\nchi mangia entro le 12.00")
}

func PayHandler(m *tbot.Message){
	message := `
Menu carne + alcool:\t16.00
Menu carne NO Alcool:\t14.00
Menu vegetariano + alcool:\t14:00
Menu vegetariano NO alcool:\t12:00`
m.Reply(message)
}
