package main

import (
	"github.com/alexejk/portal/cmd"
	"github.com/sirupsen/logrus"
)

func main() {

	if err := cmd.RootCmd().Execute(); err != nil {
		logrus.Fatal(err)
	}
}
