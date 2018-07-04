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

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("queue1"),
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", *result.QueueUrl)
}
