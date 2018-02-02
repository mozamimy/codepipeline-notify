# codepipeline-notify

A serverless application to notify whether it succeeded or not.

## Development

You need to set following environment variable to run this application in your local machine. You need to install [aws-sam-local](https://github.com/awslabs/aws-sam-local) before execute following instruction.

```sh
$ GOOS=linux make main.zip
$ aws-sam-local local invoke CodePipelineNotify -e sample-event.json --template=deploy/template/staging.yml
```

Also you can build and upload to run it simply on AWS Lambda.

```sh
$ GOOS=linux go build -o main
$ zip main.zip main
```
