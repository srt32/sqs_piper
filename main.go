package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage:", os.Args[0], "<source_queue> <sink_queue>")
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String("us-east-1")},
	})

	svc := sqs.New(sess)
	sourceUrlRes, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(os.Args[1]),
	})
	if err != nil {
		log.Fatalln("GetQueueURL failed:", err)
	}

	sinkUrlRes, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(os.Args[2]),
	})
	if err != nil {
		log.Fatalln("GetQueueURL failed:", err)
	}

	receiveResponse, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: sourceUrlRes.QueueUrl,
	})
	if err != nil {
		log.Fatalln("ReceiveMessage failed:", err)
	}

	message := receiveResponse.Messages[0]

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: message.Body,
		QueueUrl:    sinkUrlRes.QueueUrl,
	})
	if err != nil {
		log.Fatalln("SendMessage failed:", err)
	}
}
