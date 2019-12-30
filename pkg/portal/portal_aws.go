package portal

import "errors"

// AwsPortal is a type of Portal that supports looking up of connections via AWS APIs based on instanceId and region.
// This tool will support pushing of temporary ssh key to the instance prior to connection.
// Only direct connections are supported (no tunnels) at this point.
type AwsPortal struct {
	*portalShared

	instanceId string
	region     string
}

func newAwsPortal(name string, instanceId string, region string) (*AwsPortal, error) {

	return &AwsPortal{
		portalShared: &portalShared{
			name: name,
		},

		instanceId: instanceId,
		region:     region,
	}, nil
}

func (p *AwsPortal) Connect() error {
	return errors.New("connect is not yet implemented for AWS portals")
}
