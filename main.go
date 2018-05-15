// This is an example of using the awslogs package from the go-cloudwatchlogs repository.


package main

 import (
	 "github.com/bamaralf/go-cloudwatchlogs/awslogs"
 )
	func main () {
		awslogs.LogToCloudwatch("test-group","test")
	}