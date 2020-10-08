package main

import "go-mid/mq"

func main() {
	mq.Publish()
	mq.Consume()
}
