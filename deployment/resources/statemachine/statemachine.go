package statemachine

import (
	"aws-cdk-test/deployment/templates"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func CreateStack(
	scope constructs.Construct,
	env *awscdk.Environment,
	stackName string,
	lambdaOne awslambda.Function,
	lambdaTwo awslambda.Function,
) {
	stackData := templates.StateMachineStackData{
		StackName:        stackName,
		StackDescription: "statemachine stack created with cdk",
		Tags: &map[string]*string{
			"App":    jsii.String("cdk-test-app-stepfunctions"),
			"Author": jsii.String("Jerry"),
		},
		Env:             env,
		RoleName:        "statemachine-role-cdk",
		RoleDescription: "statemachine role createad with cdk",
		InlinePolicies: &map[string]awsiam.PolicyDocument{
			"cloudwatch": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
				Statements: &[]awsiam.PolicyStatement{
					awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
						Effect: awsiam.Effect_ALLOW,
						Actions: jsii.Strings(
							"logs:CreateLogDelivery",
							"logs:GetLogDelivery",
							"logs:UpdateLogDelivery",
							"logs:DeleteLogDelivery",
							"logs:ListLogDeliveries",
							"logs:PutResourcePolicy",
							"logs:DescribeResourcePolicies",
							"logs:DescribeLogGroups",
						),
						Resources: jsii.Strings(
							"*",
						),
					}),
				},
			}),
			"lambda-exec": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
				Statements: &[]awsiam.PolicyStatement{
					awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
						Effect: awsiam.Effect_ALLOW,
						Actions: jsii.Strings(
							"lambda:InvokeFunction",
						),
						Resources: &[]*string{
							lambdaOne.FunctionArn(),
							lambdaTwo.FunctionArn(),
						},
					}),
				},
			}),
			"xray": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
				Statements: &[]awsiam.PolicyStatement{
					awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
						Effect: awsiam.Effect_ALLOW,
						Actions: jsii.Strings(
							"xray:PutTraceSegments",
							"xray:PutTelemetryRecords",
							"xray:GetSamplingRules",
							"xray:GetSamplingTargets",
						),
						Resources: jsii.Strings("*"),
					}),
				},
			}),
		},
		Defnition: createStateMachineDef(scope, lambdaOne, lambdaTwo),
	}

	stackData.CreateStateMachineWithAssociatedRole(scope)
}

func createStateMachineDef(
	scope constructs.Construct,
	lambdaOne awslambda.Function,
	lambdaTwo awslambda.Function,
) awsstepfunctions.IChainable {
	lambdaOneInvoke := awsstepfunctionstasks.NewLambdaInvoke(scope, jsii.String("invoke-lambda-one"), &awsstepfunctionstasks.LambdaInvokeProps{
		LambdaFunction: lambdaOne,
		Payload:        awsstepfunctions.TaskInput_FromJsonPathAt(jsii.String("$")),
		OutputPath:     jsii.String("$.Payload"),
	})

	lambdaTwoInvoke := awsstepfunctionstasks.NewLambdaInvoke(scope, jsii.String("invoke-lambda-two"), &awsstepfunctionstasks.LambdaInvokeProps{
		LambdaFunction: lambdaTwo,
		OutputPath:     jsii.String("$.Payload"),
		Payload:        awsstepfunctions.TaskInput_FromJsonPathAt(jsii.String("$")),
	})

	waitTask := awsstepfunctions.NewWait(scope, jsii.String("wait-task"), &awsstepfunctions.WaitProps{
		Time: awsstepfunctions.WaitTime_Duration(awscdk.Duration_Seconds(jsii.Number(30))),
	})

	def := lambdaOneInvoke.Next(waitTask).Next(lambdaTwoInvoke)

	return def
}
