package openAIdialog

import (
	"SaltAIdDishes/pkg/loggers"
	"context"
	"github.com/sashabaranov/go-openai"
)

func Test(secret, name string) (string, error) {
	client := openai.NewClient(secret)
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "you are a chef",
			},
		},
	}
	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "Переведи на русский. Выйдай только перевод: " + name,
	})
	resp, err := client.CreateChatCompletion(
		context.Background(),
		req,
	)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
	//req.Messages = append(req.Messages, resp.Choices[0].Message)
	//req.Messages = append(req.Messages, openai.ChatCompletionMessage{
	//	Role:    openai.ChatMessageRoleUser,
	//	Content: "рецепт, с массой продуктов",
	//})
	//resp, err = client.CreateChatCompletion(
	//	context.Background(),
	//	req,
	//)
	//if err != nil {
	//	fmt.Printf("ChatCompletion error: %v\n", err)
	//	return
	//}
	//
	//fmt.Println(resp.Choices[0].Message.Content)

}
