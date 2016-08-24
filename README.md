# SQS Piper

A utility to receive an AWS SQS message from one queue and then send it on
another queue

`AWS_ACCOUNT_ID=121749107756 AWS_SDK_LOAD_CONFIG=true go run main.go crespi-queue-dev crespi-queue-simon`

A potential use case is to pull a message from a remote queue and pipe it into
a queue you listen on locally for debugging purposes.
