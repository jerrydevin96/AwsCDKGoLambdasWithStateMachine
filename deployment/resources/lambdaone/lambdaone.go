package lambdaone

import (
	"aws-cdk-test/deployment/templates"
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func CreateStack(scope constructs.Construct, env *awscdk.Environment, stackName string, acctNo string) awslambda.Function {
	fmt.Println("account number is " + acctNo)
	stackData := templates.LambdaStackData{
		StackName:        stackName,
		StackDescription: "stack for lambda one created through aws cdk",
		Tags: &map[string]*string{
			"App":    jsii.String("cdk-test-app-stepfunctions"),
			"Author": jsii.String("Jerry"),
		},
		Env:             env,
		RoleName:        "cdk-testing-lambda-one-role",
		RoleDescription: "role for lambda one created through cdk",
		InlinePolicies: &map[string]awsiam.PolicyDocument{
			"lambdaBasicExecution": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
				Statements: &[]awsiam.PolicyStatement{
					awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
						Effect: awsiam.Effect_ALLOW,
						Actions: jsii.Strings(
							"logs:CreateLogGroup",
							"logs:CreateLogStream",
							"logs:PutLogEvents",
						),
						Resources: jsii.Strings(
							fmt.Sprintf("arn:aws:logs:eu-west-1:%s:*", acctNo),
						),
					}),
				},
			}),
		},
		FunctionName:        "aws-cdk-lambda-one",
		FunctionDescription: "lambda one created through cdk",
		LambdaCodePath:      "../build/lambdaone.zip",
		LambdaHandler:       "lambdaone/bootstrap",
	}

	return stackData.CreateLambdaWithAssociatedPolicy(scope)
}
