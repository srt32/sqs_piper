# SQS Piper

A utility to receieve an SQS message from one queue and then send it on another
queue

`go run main.go <source queue name> <sink source name>`

A potential use case is to pull a message from a remote queue and pipe it into
a queue you listen on locally for debugging purposes.
