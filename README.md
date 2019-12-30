# Portal

Simple tool to quickly connect to favorite servers, including creation of SSH tunnels.

## Status

This project is very much in it's initial,  **experimental** phase. Use at your own risk.

## Usage

### Configuration

Configuration file is looked for at `~/.portal/config.yaml`.

Format of the file is:

```yaml
---
debug: false

# List of all known destinations
portals:

  # Tunnel
  - name: server1.tunnel
    hint: Port 3000
    command: ssh -L 3000:localhost:8080 me@server1.dot.ip

  # Direct connection
  - name: server.prod.bastion
    command: ssh -A user@bastion.dot.ip
  
  # Double tunnel
  - name: server.prod.tunnel
    hint: Local 33333 -> 12345 @ backend.dot.ip via bastion.dot.ip
    command: ssh -L 33333:localhost:33333 -A user@bastion.dot.ip ssh -L 33333:localhost:12345 -A user@backend.dot.ip
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
