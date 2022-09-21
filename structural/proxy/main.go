package main

import "fmt"

func main() {
	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

}
