package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type lastName struct {
	Name string `json:"lastName"`
}

func main() {
	log.Println("starting lambda two")

	lambda.Start(lambdaHandler)
}

func lambdaHandler(lastName lastName) {
	log.Println("last name is " + lastName.Name)
}
