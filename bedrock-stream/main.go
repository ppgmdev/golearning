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

	client := bedrockruntime.NewFromConfig(sdkConfig)

	modelId := "amazon.nova-micro-v1:0"

	messages := []bedrockTypes.Message{}

	for {

		userMessage := bedrockTypes.Message{
			Content: []bedrockTypes.ContentBlock{
				&bedrockTypes.ContentBlockMemberText{
					Value: getProompt("\nUser: "),
				},
			},
			Role: bedrockTypes.ConversationRoleUser,
		}

		messages = append(messages, userMessage)

		output, err := client.ConverseStream(ctx, &bedrockruntime.ConverseStreamInput{
			ModelId:  &modelId,
			Messages: messages,
		})

		if err != nil {
			fmt.Println("error")
			fmt.Println(err)
			return
		}

		fmt.Println("Bedrockazo: ")

		assitantMessage, err := processStreamingOutput(output, func(ctx context.Context, part string) error {
			fmt.Print(part)
			return nil
		})

		if err != nil {
			fmt.Println("error")
			fmt.Println(err)
			return
		}

		messages = append(messages, assitantMessage)
	}

}

type StreamingOutputHandler func(ctx context.Context, part string) error

func processStreamingOutput(output *bedrockruntime.ConverseStreamOutput, handler StreamingOutputHandler) (bedrockTypes.Message, error) {
	var combinedResult string

	msg := bedrockTypes.Message{}

	for event := range output.GetStream().Events() {
		switch v := event.(type) {
		case *bedrockTypes.ConverseStreamOutputMemberMessageStart:
			msg.Role = v.Value.Role
		case *bedrockTypes.ConverseStreamOutputMemberContentBlockDelta:
			textResponse := v.Value.Delta.(*bedrockTypes.ContentBlockDeltaMemberText)
			handler(context.Background(), textResponse.Value)
			combinedResult = combinedResult + textResponse.Value
		case *bedrockTypes.UnknownUnionMember:
			fmt.Println("unknown union tag: ", v.Tag)
		}
	}

	msg.Content = append(msg.Content,
		&bedrockTypes.ContentBlockMemberText{
			Value: combinedResult,
		},
	)

	return msg, nil
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
