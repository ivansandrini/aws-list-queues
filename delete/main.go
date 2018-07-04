package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	svc := sqs.New(sess)

	result, err := svc.DeleteQueue(&sqs.DeleteQueueInput{
		QueueUrl: aws.String("https://sqs.us-east-1.amazonaws.com/265025916637/queue1"),
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Sucess", result)
}
