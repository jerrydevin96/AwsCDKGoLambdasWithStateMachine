package templates

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaStackData struct {
	StackName           string
	StackDescription    string
	Tags                *map[string]*string
	Env                 *awscdk.Environment
	RoleName            string
	RoleDescription     string
	InlinePolicies      *map[string]awsiam.PolicyDocument
	FunctionName        string
	FunctionDescription string
	LambdaCodePath      string
	LambdaHandler       string
}

func (sd LambdaStackData) CreateLambdaWithAssociatedPolicy(scope constructs.Construct) awslambda.Function {

	stack := awscdk.NewStack(scope, jsii.String(sd.StackName), &awscdk.StackProps{
		StackName:   jsii.String(sd.StackName),
		Description: jsii.String(sd.StackDescription),
		Tags:        sd.Tags,
		Env:         sd.Env,
	})

	role := awsiam.NewRole(stack, jsii.String(sd.RoleName), &awsiam.RoleProps{
		RoleName:       jsii.String(sd.RoleName),
		Description:    jsii.String(sd.RoleDescription),
		InlinePolicies: sd.InlinePolicies,
		AssumedBy:      awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
	})

	function := awslambda.NewFunction(stack, jsii.String(sd.FunctionName), &awslambda.FunctionProps{
		FunctionName: jsii.String(sd.FunctionName),
		Description:  jsii.String(sd.FunctionDescription),
		Role:         role,
		Code: awslambda.AssetCode_FromAsset(
			jsii.String(sd.LambdaCodePath),
			nil,
		),
		Architecture: awslambda.Architecture_X86_64(),
		Handler:      jsii.String(sd.LambdaHandler),
		Runtime: awslambda.NewRuntime(
			jsii.String("go1.x"),
			awslambda.RuntimeFamily_GO,
			nil,
		),
	})

	return function

}
