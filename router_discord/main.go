package lab

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/RainrainWu/quizdeck"
)

func HandleRequest(ctx context.Context, event interface{}) (string, error) {
	fmt.Printf("%v\n", event)
	return fmt.Sprintf("current debug mode: %v", quizdeck.Config.GetDebugMode()), nil
}

func main() {
	lambda.Start(HandleRequest)
}
