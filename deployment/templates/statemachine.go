package templates

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type StateMachineStackData struct {
	StackName        string
	StackDescription string
	Tags             *map[string]*string
	Env              *awscdk.Environment
	RoleName         string
	RoleDescription  string
	InlinePolicies   *map[string]awsiam.PolicyDocument
	Defnition        awsstepfunctions.IChainable
}

func (sd StateMachineStackData) CreateStateMachineWithAssociatedRole(scope constructs.Construct) {

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
		AssumedBy:      awsiam.NewServicePrincipal(jsii.String("states.amazonaws.com"), nil),
	})

	awsstepfunctions.NewStateMachine(stack, jsii.String("statemachine-cdk"), &awsstepfunctions.StateMachineProps{
		StateMachineName: jsii.String("statemachine-lambdas-cdk"),
		Role:             role,
		StateMachineType: awsstepfunctions.StateMachineType_STANDARD,
		Definition:       sd.Defnition,
	})

}
