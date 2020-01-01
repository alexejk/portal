package main

import (
	"alexejk.io/portal/cmd"
	"alexejk.io/portal/pkg/version"
	"github.com/sirupsen/logrus"
)

// appVersion denotes version number of this application
var appVersion = "x.y.z"

// AppName denotes name of this application
var appName = "portal"

// AppCommit denotes git commit that this binary was built from
var appCommit = "~unknown~"

// AppCommitDate denotes the date of the commit
var appCommitDate = "~unknown~"

// AppCommitTime denotes the time of the commit
var appCommitTime = "~unknown~"

func main() {

	v := version.NewVersionInfo(appName, appVersion, appCommit, appCommitDate, appCommitTime)

	if err := cmd.RootCmd(v).Execute(); err != nil {
		logrus.Fatal(err)
	}
}
