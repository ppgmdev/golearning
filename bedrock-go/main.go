package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	bedrockTypes "github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

func main() {

	region := flag.String("region", "us-east-1", "The AWS region")
	flag.Parse()

	fmt.Printf("Using AWS region: %s\n", *region)

	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion(*region))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}

	proompt := getProompt("Write your proompt:")

	userMessage := bedrockTypes.Message{
		Content: []bedrockTypes.ContentBlock{
			&bedrockTypes.ContentBlockMemberText{
				Value: proompt,
			},
		},
		Role: bedrockTypes.ConversationRoleUser,
	}

	messages := []bedrockTypes.Message{userMessage}

	client := bedrockruntime.NewFromConfig(sdkConfig)

	modelId := "amazon.nova-micro-v1:0"

	output, err := client.Converse(ctx, &bedrockruntime.ConverseInput{
		ModelId:  &modelId,
		Messages: messages,
	})

	response, ok := output.Output.(*bedrockTypes.ConverseOutputMemberMessage)

	if !ok {
		fmt.Println(err)
		fmt.Println("An error has occurred")
		panic("exit")
	}

	responseContentBlock := response.Value.Content[0]

	text, ok := responseContentBlock.(*bedrockTypes.ContentBlockMemberText)

	if !ok {
		fmt.Println(err)
		fmt.Println("An error has occurred")
		panic("exit")
	}

	fmt.Println("Bedrockazo:")
	fmt.Println(text.Value)

	if err != nil {
		fmt.Println(err)
		fmt.Println("An error has occurred")
		panic("exit")
	}
}

func getProompt(cmdText string) string {
	fmt.Println(cmdText)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
