package main

import (
	"github.com/alexejk/portal/cmd"
	"github.com/alexejk/portal/pkg/version"
	"github.com/sirupsen/logrus"
)

// AppVersion denotes version number of this application
var AppVersion = "x.y.z"

// AppName denotes name of this application
var AppName = "portal"

// AppCommit denotes git commit that this binary was built from
var AppCommit = "~unknown~"

// AppCommitDate denotes the date of the commit
var AppCommitDate = "~unknown~"

// AppCommitTime denotes the time of the commit
var AppCommitTime = "~unknown~"

func main() {

	v := version.NewVersionInfo(AppName, AppVersion, AppCommit, AppCommitDate, AppCommitTime)

	if err := cmd.RootCmd(v).Execute(); err != nil {
		logrus.Fatal(err)
	}
}
