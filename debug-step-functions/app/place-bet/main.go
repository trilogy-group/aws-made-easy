package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
    return bet(), nil
}

func bet() (string){
	rand.Seed(time.Now().UnixNano())

	bet := ""
	switch rand.Intn(100000)%2 {
	case 0: bet = "tails"
	case 1: bet = "heads"
	}
	return bet
}

func main() {
	lambda.Start(HandleRequest)
}