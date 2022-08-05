package main

import (
	"aws-cdk-test/deployment/resources/lambdaone"
	"aws-cdk-test/deployment/resources/lambdatwo"
	"aws-cdk-test/deployment/resources/statemachine"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	fmt.Println("starting cdk tests")
	env := os.Getenv("ENV")
	app := awscdk.NewApp(nil)

	if strings.EqualFold(env, "envone") {
		fmt.Println("processing env " + env)
		createStackForEnvOne(app, env)
	} else if strings.EqualFold(env, "envtwo") {
		fmt.Println("processing env " + env)
		createStackForEnvTwo(app, env)
	} else {
		fmt.Println("no valid env found")
	}

	app.Synth(nil)
}

func createStackForEnvOne(app constructs.Construct, env string) {
	acctNo := os.Getenv("ACCTNO")
	stackEnv := &awscdk.Environment{
		Account: jsii.String(acctNo),
		Region:  jsii.String("eu-west-1"),
	}
	funcOne := lambdaone.CreateStack(app, stackEnv, fmt.Sprintf("lambda-one-stack-cdk-%s", env), acctNo)
	funcTwo := lambdatwo.CreateStack(app, stackEnv, fmt.Sprintf("lambda-two-stack-cdk-%s", env), acctNo)
	statemachine.CreateStack(app, stackEnv, fmt.Sprintf("statemachine-stack-cdk-%s", env), funcOne, funcTwo)
}

func createStackForEnvTwo(app constructs.Construct, env string) {
	acctNo := os.Getenv("ACCTNO")
	stackEnv := &awscdk.Environment{
		Account: jsii.String(acctNo),
		Region:  jsii.String("eu-west-1"),
	}
	funcOne := lambdaone.CreateStack(app, stackEnv, fmt.Sprintf("lambda-one-stack-cdk-%s", env), acctNo)
	funcTwo := lambdatwo.CreateStack(app, stackEnv, fmt.Sprintf("lambda-two-stack-cdk-%s", env), acctNo)
	statemachine.CreateStack(app, stackEnv, fmt.Sprintf("statemachine-stack-cdk-%s", env), funcOne, funcTwo)
}
