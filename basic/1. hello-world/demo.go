package main

import "fmt"

func main() {
	serviceName := "service-demo"
	serviceVersion := "v0.1.0"
	serverPort := 8080
	serverEnv := "dev"
	fmt.Printf("serviceName: %s\nserviceVersion: %s\nserverPort: %d\nserverEnv: %s\n", 
				serviceName, serviceVersion, serverPort, serverEnv)
}