# Portal


![GitHub Workflow Status](https://img.shields.io/github/workflow/status/alexejk/portal/Build)
[![Go Report Card](https://goreportcard.com/badge/alexejk.io/portal)](https://goreportcard.com/report/alexejk.io/portal)

![GitHub](https://img.shields.io/github/license/alexejk/portal)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/alexejk/portal)

Simple tool to quickly connect to favorite servers, including creation of SSH tunnels.

## Status

This project is very much in it's initial,  **experimental** phase. Use at your own risk.

## Features

Portal provides following set of features that one could find useful:

* Bookmark-like feature for SSH connections - simply run `portal connect <name>`
* Support for complex connections such as multi-hop SSH tunnels
* Support for connections with AWS EC2 Instance Connect by temporarily pushing SSH key to the instance
* Provides (potentially )useful hints upon established connection, e.g what ports are being forwarded 

## Usage

### Configuration

Configuration file is looked for at `~/.portal/config.yaml`. 

```yaml
---
debug: false

default-public-key: ~/.ssh/id_rsa.pub
default-private-key: ~/.ssh/id_rsa

# List of all known destinations
portals:

  # Tunnel
  - name: server1.tunnel
    hint: Port 3000
    raw:
      command: ssh -L 3000:localhost:8080 me@server1.dot.ip

  # Direct connection
  - name: server.prod.bastion
    raw:
      command: ssh -A user@bastion.dot.ip
  
  # Double tunnel
  - name: server.prod.tunnel
    hint: Local 33333 -> 12345 @ backend.dot.ip via bastion.dot.ip
    raw:
      command: ssh -L 33333:localhost:33333 -A user@bastion.dot.ip ssh -L 33333:localhost:12345 -A user@backend.dot.ip
  
  # AWS EC2 Instance Connect
  - name: aws.host
    aws:
      instance-id: i-1234abcd
      region: us-west-1
      user: ec2-user
```

### Supported Commands

* Get all available connections:  
  `portal list`
* Connect to a server:  
  `portal connect <name>` where `<name>` is the favorite name

## TODOs

See separate [TODO document](TODO.md) for what is planned.  

## Building

Simply run `make all` or `make build-in-docker` if you have Docker installed 

## License

MIT license, as described in the [LICENSE](LICENSE) file.
