package main

import (
	"workPackage/rest"

	"github.com/sirupsen/logrus"
)

/**
 * Golang Echo Server
 */
func main() {
	logrus.Println("Starting Golang Echo Server")
	logrus.Fatal(rest.RunServer(":8080"))
}
