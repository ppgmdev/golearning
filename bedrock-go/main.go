package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	bedrockTypes "github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

// Each model provider defines their own individual request and response formats.
// For the format, ranges, and default values for the different models, refer to:
// https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html

func main() {

	region := flag.String("region", "us-west-2", "The AWS region")
	flag.Parse()

	fmt.Printf("Using AWS region: %s\n", *region)

	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion(*region))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}

	userMessage := bedrockTypes.Message{
		Content: []bedrockTypes.ContentBlock{
			&bedrockTypes.ContentBlockMemberText{
				Value: "Hello. How can i deploy a lambda in go in AWS?",
			},
		},
		Role: bedrockTypes.ConversationRoleUser,
	}

	messages := []bedrockTypes.Message{userMessage}

	client := bedrockruntime.NewFromConfig(sdkConfig)

	modelId := "anthropic.claude-3-5-haiku-20241022-v1:0"

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
