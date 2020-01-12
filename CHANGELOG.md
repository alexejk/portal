## WIP

_TBD_

## 0.0.2

### Enhancements

#### AWS Portals

Added support for AWS connections through AWS EC2 Instance Connect.
AWS Portals can be defined with `aws:` key in portal definition.

```yaml
portals:
  - name: aws.host
    aws:
      instace-id: i-abc1234
      region: us-west-2
      user: ec2-user
```

#### Default keys

To support AWS Portals, new top-level configuration properties `default-public-key` and `default-private-key` are added.
Default values are `~/.ssh/id_rsa.pub` and `~/.ssh/id_rsa` respectively, but can be changed if you use different format.

This feature will come extra handy with future support for:

* declarative connections
* per-connection keys including temporary keys

### Breaking change

Config format for basic command connections has changed: `portals.[].command` -> `portals.[].raw.command`. 
This is done to support declarative connections in the future.

Instead of 

```yaml
  - name: server.prod.bastion
    command: ssh -A user@bastion.dot.ip
```

Use

```yaml
  - name: server.prod.bastion
    raw:
      command: ssh -A user@bastion.dot.ip
```


## 0.0.1

Initial release with support for basic connections and tunneling operations.

