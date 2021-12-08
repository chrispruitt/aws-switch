# aws-switch

Use tags to identify aws resources to halt or resume expensive services.

The basic idea:
1. identify resources based on a given set of tags
1. halt   - scale resources down to zero and store previous values in a state
1. resume - scale resources back to previous value based on stored state

Currenlty supports
- [x] ecs:services (does not include underlying compute if running on ec2)
- [x] rds:clusters (provissioned only as serverless auto scales on its own)
- [x] rds:db (all database instances that are not a part of a cluster)

### Homebrew
```bash
brew install chrispruitt/tap/aws-switch
```

### Install from source

Install binary from [source](https://github.com/chrispruitt/aws-switch/releases).

## Usage

```text
    $> aws-switch --help

    aws-switch is used to halt and resume all aws services by tag.

    Usage:
    aws-switch [command]

    Available Commands:
    configure   Creates an s3 bucket for the aws-switch state to reside
    halt        halt an aws service
    help        Help about any command
    resume      Resume a halted aws service

    Flags:
    -h, --help      help for aws-switch
        --version   version for aws-switch
```

## Roadmap
- [ ] --exclude-tag flag to optionally exclude resources with specific tags
- [ ] different backends - local, dynamodb
- [ ] read configuration from yaml file or env vars
- [x] show changes and ask for confirmation before scaling up or down 
- [ ] flag to select specific services - rds, ecs - default to all
- [ ] show state
- [ ] resume specific items in state - work around using a good tagging stradegty and just identify resources by tags