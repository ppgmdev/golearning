package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type PipelineStackProps struct {
	awscdk.StackProps
}

func NewPipelineStack(scope constructs.Construct, id string, props *PipelineStackProps) awscdk.Stack {
	var sprops awscdk.StackProps

	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)

	repo := awscodecommit.NewRepository(stack, jsii.String("WorkshopRepo"), &awscodecommit.RepositoryProps{
		RepositoryName: jsii.String("CDKgoWorkshopRepo"),
	})

	pipeline := pipelines.NewCodePipeline(stack, jsii.String("WorkshopPipeline"), &pipelines.CodePipelineProps{
		PipelineName: jsii.String("WorkshopPipeline"),
		Synth: pipelines.NewCodeBuildStep(jsii.String("SynthStep"), &pipelines.CodeBuildStepProps{
			Input: pipelines.CodePipelineSource_CodeCommit(repo, jsii.String("master"), nil),
			Commands: jsii.Strings(
				"npm install -g aws-cdk",
				"goenv install 1.18.3",
				"goenv local 1.18.3",
				"npx cdk synth",
			),
		}),
	})

	deploy := NewWorkshopPipelineStage(stack, "Deploy", nil)
	pipeline.AddStage(deploy, nil)

	return stack
}
