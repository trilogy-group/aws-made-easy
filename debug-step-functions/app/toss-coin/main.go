package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)


func HandleRequest(ctx context.Context) (string, error) {
    return toss(), nil
}

func toss() (string){
	rand.Seed(time.Now().UnixNano())

	face := ""
	switch(rand.Intn(100000)%2) {
	case 0: face = "tails"
	case 1: face = "heads"
	}
	return face
}

func main() {
	lambda.Start(HandleRequest)
}