package main

import (
	"fmt"
	"log"
	"omni/openai"
	"os"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	// db, err := dbase.InitDatabase()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// server.Serv(db)

	client := openai.InitClient()

	r := client.OpenAiCall.CreateAssistant()
	f := client.OpenAiCall.ChatCompletion("Hello, I am Rylan. How are you.")
	fmt.Println(r)
	fmt.Println(f)
}
