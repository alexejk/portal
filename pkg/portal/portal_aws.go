package portal

import (
	"fmt"
	"io/ioutil"
	"strings"

	"alexejk.io/portal/pkg/config"
	"alexejk.io/portal/pkg/run"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	"github.com/mitchellh/go-homedir"
)

const (
	defaultPublicKey  = "~/.ssh/id_rsa.pub"
	defaultPrivateKey = "~/.ssh/id_rsa"
)

// AwsPortal is a type of Portal that supports looking up of connections via AWS APIs based on instanceId and region.
// This tool will support pushing of temporary ssh key to the instance prior to connection.
// Only direct connections are supported (no tunnels) at this point.
type AwsPortal struct {
	*portalShared

	instanceId *string
	region     *string
	user       *string

	instanceAz      *string
	instanceAddress *string
}

func newAwsPortal(name string, awsConfig *config.PortalAwsConfig) (*AwsPortal, error) {

	return &AwsPortal{
		portalShared: &portalShared{
			name: name,
		},

		instanceId: awsConfig.InstanceId,
		region:     awsConfig.Region,
		user:       awsConfig.User,
	}, nil
}

func (p *AwsPortal) Connect() error {

	sess, err := p.session()
	if err != nil {
		return err
	}

	if err := p.lookup(sess); err != nil {
		return fmt.Errorf("failed looking up AWS instance ip: %w", err)
	}

	if err := p.configure(sess); err != nil {
		return fmt.Errorf("failed configuring AWS instance: %w", err)
	}

	key, err := p.defaultPrivateKeyName()
	if err != nil {
		return err
	}

	args := []string{
		"-i", key,
		fmt.Sprintf("%s@%s", *p.user, *p.instanceAddress),
	}

	c := run.NewRunner("ssh", args...)
	return c.Run()
}

func (p *AwsPortal) session() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: p.region,
	})

	return sess, err
}

func (p *AwsPortal) configure(sess *session.Session) error {

	key, err := p.defaultPublicKeyBody()
	if err != nil {
		return err
	}

	ic := ec2instanceconnect.New(sess)

	req := &ec2instanceconnect.SendSSHPublicKeyInput{
		InstanceId:       p.instanceId,
		InstanceOSUser:   p.user,
		AvailabilityZone: p.instanceAz,
		SSHPublicKey:     key,
	}

	resp, err := ic.SendSSHPublicKey(req)
	if err != nil {
		return err
	}

	if !*resp.Success {
		return fmt.Errorf("could not set EC2 instance SSH-Key. %s", *resp.RequestId)
	}

	return nil
}

func (p *AwsPortal) defaultPublicKeyName() (string, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return "", err
	}

	key := defaultPublicKey
	if cfg.DefaultPublicKey != "" {
		key = cfg.DefaultPublicKey
	}

	return key, nil
}

func (p *AwsPortal) defaultPublicKeyBody() (*string, error) {

	key, err := p.defaultPublicKeyName()
	if err != nil {
		return nil, err
	}

	path, err := homedir.Expand(key)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	keyBody := string(b)
	if !strings.HasPrefix(keyBody, "ssh-") {
		return nil, fmt.Errorf("public key is invalid format")
	}

	return &keyBody, nil
}

func (p *AwsPortal) defaultPrivateKeyName() (string, error) {

	cfg, err := config.GetConfig()
	if err != nil {
		return "", err
	}

	key := defaultPrivateKey
	if cfg.DefaultPrivateKey != "" {
		key = cfg.DefaultPrivateKey
	}

	return key, nil
}

func (p *AwsPortal) lookup(sess *session.Session) error {

	ec2c := ec2.New(sess)
	r := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			p.instanceId,
		},
	}
	resp, err := ec2c.DescribeInstances(r)
	if err != nil {
		return err
	}

	if len(resp.Reservations) != 1 {
		return fmt.Errorf("cannot uniquely identify instances by id. Reservations found: %d", len(resp.Reservations))
	}
	instances := resp.Reservations[0].Instances

	if len(instances) != 1 {
		return fmt.Errorf("cannot uniquely identify instances by id. Instances found: %d", len(instances))
	}

	instance := instances[0]

	p.instanceAz = instance.Placement.AvailabilityZone
	p.instanceAddress = instance.PublicIpAddress

	return nil
}
