# Welcome [AWS Made Easy](https://awsmadeeasy.com/) - Debugging Step Functions Demo



This is a demo project to help debug step functions.

Setup:
1. Install [go](https://go.dev/doc/install)
2. Install [Taskfile](https://taskfile.dev/installation/)
2. Install [AWS CDK](https://docs.aws.amazon.com/cdk/v2/guide/getting_started.html)
3. Export `CDK_DEFAULT_REGION` and `CDK_DEFAULT_ACCOUNT` environment variables
4. Setup [AWS credentials](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/setup-credentials.html)

Deploy:
1. goto [app/place-bet](app/place-bet) folder and run `task build`
2. goto [app/toss-coin](app/toss-coin) folder and run `task build`
3. goto [infra](infra) folder and run `task diff`, review the changes and then run `task deploy`

Debug:
1. Login to AWS console and navigate to [Step Functions](https://console.aws.amazon.com/states/home)
2. Click on the step function `AwsMadeEasy-DebugStepFunctions`
3. Click on `New execution` button and then click on `Start execution` button
4. Follow the steps in the article to debug the step function

Cleanup:
1. goto [infra](infra) folder and run `task destroy`
