package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage:", os.Args[0], "<source_queue_url> <sink_queue_url>")
	}

	// awsAccountId := os.Getenv("AWS_ACCOUNT_ID")

	sourceUrl := os.Args[1]
	sinkUrl := os.Args[2]

	log.Println("Source url:", sourceUrl)
	log.Println("Sink url:", sinkUrl)

	sess, err := session.NewSessionWithOptions(session.Options{
		// Config:  aws.Config{Region: aws.String("us-east-1")},
		Profile: "development",
	})

	svc := sqs.New(sess)
	//sourceUrlRes, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
	//	QueueName:              aws.String(os.Args[1]),
	//	QueueOwnerAWSAccountId: aws.String(awsAccountId),
	//})
	//if err != nil {
	//	log.Fatalln("Source GetQueueURL failed:", err)
	//}

	//sinkUrlRes, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
	//	QueueName:              aws.String(os.Args[2]),
	//	QueueOwnerAWSAccountId: aws.String(awsAccountId),
	//})
	//if err != nil {
	//	log.Fatalln("Sink GetQueueURL failed:", err)
	//}

	// TODO: increase wait time
	receiveResponse, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: &sourceUrl,
	})
	if err != nil {
		log.Fatalln("ReceiveMessage failed:", err)
	}

	message := receiveResponse.Messages[0]

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: message.Body,
		QueueUrl:    &sinkUrl,
	})
	if err != nil {
		log.Fatalln("SendMessage failed:", err)
	}
}
