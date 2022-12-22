package main

import (
	"context"
	"fmt"
	"os"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//preparando para requisicao
	c := gogpt.NewClient("api-key")
	ctx := context.Background()

	//argumento cli
	wtf := os.Args[1]

	//montando requisicao
	req := gogpt.CompletionRequest{
		Model: "text-davinci-003",
		MaxTokens: 2048,
		Prompt:    wtf,
	}

	//requisicao para o openai
	resp, err := c.CreateCompletion(ctx, req)
	check(err)

	//print zuado
	fmt.Println("\n* Gerando arquivo do código... \n* Editar o arquivo file.temp antes de utilizar!\n ")

	//escrevendo no arquivo
	f1 := []byte(resp.Choices[0].Text)
	err = os.WriteFile("file.temp", f1, 0644)
	check(err)
}