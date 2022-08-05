package lambdatwo

import (
	"aws-cdk-test/deployment/templates"
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func CreateStack(scope constructs.Construct, env *awscdk.Environment, stackName string, accountNo string) awslambda.Function {
	fmt.Println("account number is " + accountNo)
	stackData := templates.LambdaStackData{
		StackName:        stackName,
		StackDescription: "stack for lambda two created through aws cdk",
		Tags: &map[string]*string{
			"App":    jsii.String("cdk-test-app-stepfunctions"),
			"Author": jsii.String("Jerry"),
		},
		Env:             env,
		RoleName:        "cdk-testing-lambda-two-role",
		RoleDescription: "role for lambda two created through cdk",
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
							fmt.Sprintf("arn:aws:logs:eu-west-1:%s:*", accountNo),
						),
					}),
				},
			}),
		},
		FunctionName:        "aws-cdk-lambda-two",
		FunctionDescription: "lambda two created through cdk",
		LambdaCodePath:      "../build/lambdatwo.zip",
		LambdaHandler:       "lambdatwo/bootstrap",
	}

	return stackData.CreateLambdaWithAssociatedPolicy(scope)
}
