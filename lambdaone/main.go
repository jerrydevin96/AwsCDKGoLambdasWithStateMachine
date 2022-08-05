package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type firstName struct {
	Name string `json:"firstName"`
}

type lastName struct {
	Name string `json:"lastName"`
}

func main() {
	log.Println("starting lambda one")

	lambda.Start(lambdaHandler)
}

func lambdaHandler(firstName firstName) (lastName, error) {
	log.Println("first name is " + firstName.Name)

	return lastName{
		Name: "Doe",
	}, nil
}
