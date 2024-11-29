package hitcounter

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type HitCounterProps struct {
	//Downstream is the function for which we want to count this
	Downstream   awslambda.IFunction
	ReadCapacity float64
}

type hitCounter struct {
	constructs.Construct
	handler awslambda.IFunction
}

type HitCounter interface {
	constructs.Construct
	Handler() awslambda.IFunction
}

func (h *hitCounter) Handler() awslambda.IFunction {
	return h.handler
}

func NewHitCounter(scope constructs.Construct, id string, props *HitCounterProps) HitCounter {
	if props.ReadCapacity < 5 || props.ReadCapacity > 20 {
		panic("ReadCapacity must be between 5 and 20")
	}
	this := constructs.NewConstruct(scope, &id)

	table := awsdynamodb.NewTable(this, jsii.String("Hits"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("path"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		Encryption:    awsdynamodb.TableEncryption_AWS_MANAGED,
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		ReadCapacity:  &props.ReadCapacity,
	})

	handler := awslambda.NewFunction(this, jsii.String("HitCounterHandler"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_NODEJS_16_X(),
		Handler: jsii.String("hitcounter.handler"),
		Code:    awslambda.Code_FromAsset(jsii.String("lambda"), nil),
		Environment: &map[string]*string{
			"DOWNSTREAM_FUNCTION_NAME": props.Downstream.FunctionName(),
			"HITS_TABLE_NAME":          table.TableName(),
		},
	})

	table.GrantReadWriteData(handler)

	props.Downstream.GrantInvoke(handler)

	return &hitCounter{
		this,
		handler,
	}
}
