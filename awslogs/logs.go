// This file keeps the method LogToCloudwatch. The method receives 2 arguments:
// - A string with the Cloudwatch log group name.
// - A string with the message to be logged into Cloudwatch.
// The method also checks if the log group name already exists and if not it creates the log group.
// Usage example:
//
//     LogToCloudwatch ("test-group", "test message")
// 
// Usage example 2:
//     
//     test := "test-group"
//     LogToCloudwatch(test, "Cambio! This message will be logged on AWS Cloudwatch!")
//
// The log group name will be the value kept on the logGroupName variable with a "." added to the beginning of the string.
// ".test-group"
//
// The log stream inside the log group will be the hostname.
//
// brunoamaralf
 
package awslogs

import (
	"fmt"
	"os"
	"time"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)
//(*cloudwatchlogs.CloudWatchLogs, string, error)
func LogToCloudwatchInit(logGroupName, logStreamName string) (*cloudwatchlogs.CloudWatchLogs, string, error)  {
	sess := session.Must(session.NewSession())
	client := cloudwatchlogs.New(sess)

	nextToken := ""

    //fetch group descriptors
	groupsOut, err := client.DescribeLogGroups(&cloudwatchlogs.DescribeLogGroupsInput{}) 
	groups := groupsOut.LogGroups
	groupExists := false
		if err == nil {
		//check if group exists
		for _, group := range groups {
			groupString := *group.LogGroupName 
			if groupString == logGroupName {
		 	   groupExists = true
		 	   break
		 	}
		 }
	}
		 
	// create group if it does not exist
 	if !groupExists { 
	  _,err := client.CreateLogGroup(&cloudwatchlogs.CreateLogGroupInput{LogGroupName: &logGroupName})  
		   if err != nil { 
			   fmt.Println("Error: The log group can not be created! %v", err)
			   return nil, "", err
		   }
 	  } 

 	// fetch group stream descriptors
 	streamsOut, err := client.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{LogGroupName: &logGroupName})
	streams := streamsOut.LogStreams
    streamExists := false
	if err == nil {
		// check if stream exists
	    for _, stream := range streams {
		    streamString := *stream.LogStreamName
		    if streamString == logStreamName {
				nextToken = *stream.UploadSequenceToken
			    streamExists = true
			    break
		}
	  }	
	 }
	 
    // create stream if it does not exist
 	if !streamExists { 
		_,err := client.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{LogGroupName: &logGroupName, LogStreamName: &logStreamName})  
			 if err != nil { 
				 fmt.Println("Error: The log stream can not be created! %v", err)
				 return nil, "", err
			 }
		 } 

		      return client, nextToken, nil
}

func LogToCloudwatch(logGroupName, messageLog string ) {
	logGroupName = "."+logGroupName // Avoiding to iterate over log groups pages. The created log group 
	                                // will be always placed at the beginning of the log groups list. 
 	logStreamName, err := os.Hostname()
 	if err != nil {
 		fmt.Printf("%s\n", err)
 		return
 	}

 	client, sequenceToken, err := LogToCloudwatchInit(logGroupName, logStreamName)
	if err != nil {
 		fmt.Printf("%s\n", err)
 		return
 	}

    // log to Cloudwatch

 	//create log events
	 events := make([]*cloudwatchlogs.InputLogEvent, 0)
	 msg := time.Now().Format("2006-01-02 15:04:05") + " " + messageLog
	 timestamp := time.Now().UnixNano() / 1000000 
     events = append(events, &cloudwatchlogs.InputLogEvent{Message: &msg, Timestamp: &timestamp})

 	// send log events
	if sequenceToken == "" {  
	reqOut,_ := client.PutLogEventsRequest(&cloudwatchlogs.PutLogEventsInput{
 		LogEvents:     events,
 		LogStreamName: &logStreamName,
 		LogGroupName:  &logGroupName})
		   err = reqOut.Send()
		   if err != nil {
			   fmt.Println(err)
		   }
    } else {
		reqOut,_ := client.PutLogEventsRequest(&cloudwatchlogs.PutLogEventsInput{
			LogEvents:     events,
			LogStreamName: &logStreamName,
			LogGroupName:  &logGroupName,
			SequenceToken: &sequenceToken})	
			err = reqOut.Send()
			if err != nil {
				fmt.Println(err)
			}
	}

}

