# go-cloudwatchlogs/awslogs Go package

`Before using the 'awslogs' package function please export the environment variable AWS_SDK_LOAD_CONFIG=1 or go to get more details and
alternative configurations: https://github.com/aws/aws-sdk-go.
"The SDK has support for the shared configuration file (~/.aws/config). This support can be enabled by setting the environment variable, "AWS_SDK_LOAD_CONFIG=1", or enabling the feature in code when creating a Session via the Option's SharedConfigState parameter." `

The awslogs go package provides a function that you can use to post log messages (or any other string of interest) to AWS Cloudwatch.

The Go function LogToCloudwatch receives 2 arguments:
 - A string with the Cloudwatch log group name.
 - A string with the message to be logged into Cloudwatch.
 The function checks if the log group name/log stream name already exists and if not it creates the log group.
 
 Usage example:
     
     import (
        "github.com/bamaralf/go-cloudwatchlogs/awslogs"
     ) 
     ...
     LogToCloudwatch ("test-group", "test message")
     ...

 Usage example 2:
     
     import (
         "github.com/bamaralf/go-cloudwatchlogs/awslogs"
     )
     ...
     test := "test-group"
     LogToCloudwatch(test, "Cambio! This message will be logged on AWS Cloudwatch!")
     ...

 The log group name will be the value kept on the logGroupName variable with a "." added to the beginning of the string.
 
 ".test-group"

 The log stream inside the log group will be the hostname.

 # Dependencies:

   AWS SDK for the Go programming language. http://aws.amazon.com/sdk-for-go/

 
 brunoamaralf
