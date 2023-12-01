package main

import (
	"fmt"
	"log"
	"omni/openai"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// db, err := dbase.InitDatabase()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// server.Serv(db)

	client := openai.New()
	// Thread := client.OpenAiCall.CreateThread()

	// ai := openai.NewAssistant(client, Thread)
	// r := ai.MessageAssistant("Can you please get the news for me?").(*openai.ListRunStep)
	// fmt.Println(r)

	// r := client.OpenAiCall.VisionCompletion("What is in this photo?")
	// fmt.Println(r)

	c := client.OpenAiCall.ChatCompletion("Hello")
	fmt.Println(c)
}
