package version

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type Info struct {
	Name       *string
	Version    *string
	Commit     *string
	CommitTime *time.Time
}

func NewVersionInfo(name, version, commit, commitDate, commitTime string) *Info {

	t, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT%sZ", commitDate, commitTime))
	if err != nil {
		logrus.Errorf("unable to parse commit timestamp (%s %s): %s", commitDate, commitTime, err)
		t = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	i := Info{
		Name:       &name,
		Version:    &version,
		Commit:     &commit,
		CommitTime: &t,
	}

	return &i
}

func (i *Info) String() string {
	return fmt.Sprintf("%s (Commit: %s @ %s)",
		*i.Version,
		*i.Commit,
		i.CommitTime.Format(time.RFC3339))
}
