package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mozamimy/codepipeline-notify/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
