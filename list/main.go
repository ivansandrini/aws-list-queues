package list

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

	result, err := svc.ListQueues(nil)
	if err != nil {
		fmt.Println("Error ", err)
	}

	fmt.Println("Success")
	for i, urls := range result.QueueUrls {
		if urls == nil {
			continue
		}
		fmt.Printf("%d: %s\n", i, *urls)
	}
}
