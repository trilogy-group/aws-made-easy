package main

import (
	"log"
	"os"
	"path/filepath"
	
	aws "github.com/aws/aws-sdk-go-v2/aws"
	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	awss3assets "github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	constructs "github.com/aws/constructs-go/constructs/v10"
)

type DebugStepFunctionsStackProps struct {
	awscdk.StackProps
}

func NewDebugStepFunctionsStack(scope constructs.Construct, id string, props *DebugStepFunctionsStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	placeBet := placeBet(stack)
	tossCoin := tossCoin(stack)
	playGame(stack, placeBet, tossCoin)

	return stack
}

func placeBet(stack awscdk.Stack) lambda.Function {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lambdaPath := filepath.Join(path, "../.dist/place-bet/main.zip")

	fn := lambda.NewFunction(stack, 
		aws.String("aws-made-easy-debug-step-functions-place-bet"), 
		&lambda.FunctionProps{
			Description:    aws.String("AWS Made Easy - Debug Step Functions - Place a bet"),
			FunctionName:   aws.String("aws-made-easy-debug-step-functions-place-bet"),
			MemorySize:     aws.Float64(128),
			Timeout:        awscdk.Duration_Seconds(aws.Float64(10)),
			Code: 			lambda.Code_FromAsset(&lambdaPath, &awss3assets.AssetOptions{}),
			Handler: 		aws.String("main"),
			Runtime:        lambda.Runtime_GO_1_X(),
		},
	)

	return fn
}

func tossCoin(stack awscdk.Stack) lambda.Function {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lambdaPath := filepath.Join(path, "../.dist/toss-coin/main.zip")

	fn := lambda.NewFunction(stack, 
		aws.String("aws-made-easy-debug-step-functions-toss-coin"), 
		&lambda.FunctionProps{
			Description:    aws.String("AWS Made Easy - Debug Step Functions - Toss a coin"),
			FunctionName:   aws.String("aws-made-easy-debug-step-functions-toss-coin"),
			MemorySize:     aws.Float64(128),
			Timeout:        awscdk.Duration_Seconds(aws.Float64(10)),
			Code: 			lambda.Code_FromAsset(&lambdaPath, &awss3assets.AssetOptions{}),
			Handler: 		aws.String("main"),
			Runtime:        lambda.Runtime_GO_1_X(),
		},
	)

	return fn
}

func playGame(stack awscdk.Stack, placeBet lambda.Function, tossCoin lambda.Function) {
	placeBetStep := tasks.NewLambdaInvoke(stack, 
		aws.String("Place bet"), 
		&tasks.LambdaInvokeProps{
			LambdaFunction: placeBet,
		},
	)
	
	tossCoinStep := tasks.NewLambdaInvoke(stack, 
		aws.String("Toss Coin"), 
		&tasks.LambdaInvokeProps{
			LambdaFunction: tossCoin,
		},
	)
	
	play := sfn.NewParallel(stack, aws.String("Let's play"), nil)
	play.Branch(placeBetStep)
	play.Branch(tossCoinStep)
	
	won := sfn.NewPass(stack, aws.String("You Won!"), nil)
	lose := sfn.NewFail(stack, aws.String("You Lost!"), nil)

	outcome := sfn.NewChoice(stack, aws.String("Outcome"), nil)
	outcome.When(
		sfn.Condition_Or(
			sfn.Condition_And(
				sfn.Condition_StringEquals(aws.String("$[0].Payload"), aws.String("tails")),
				sfn.Condition_StringEquals(aws.String("$[1].Payload"), aws.String("tails")),
			),
			sfn.Condition_And(
				sfn.Condition_StringEquals(aws.String("$[0].Payload"), aws.String("heads")),
				sfn.Condition_StringEquals(aws.String("$[1].Payload"), aws.String("heads")),
			), 
		), won).Otherwise(lose)

	definition := play.Next(outcome)

	sfn.NewStateMachine(stack, aws.String("StateMachine"), &sfn.StateMachineProps{
		StateMachineName: aws.String("AwsMadeEasy-DebugStepFunctions"),
		Definition: definition,
		Timeout: awscdk.Duration_Seconds(aws.Float64(30)),
	})

}

func main() {
	app := awscdk.NewApp(nil)

	NewDebugStepFunctionsStack(app, 
		"AwsMadeEasy-DebugStepFunctionsStack", 
		&DebugStepFunctionsStackProps{
			awscdk.StackProps{
				Env: env(),
			},
		},
	)

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: aws.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	  	Region:  aws.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
