# [Golang.hr] Slack Invite Installation Guide

**STILL WORK IN PROGRESS**

## Requirements
----

Before we even start with anything related to installation, here's what you need to have.

#### [Go] Required

You will need to have [Go] 1.4+ installed and ready for use.

If you're unsure how to do it or need any sort of help, please navigate yourself to
[Go Getting Started]. Go team explained very nicely what you need to do depending on your
environment.

#### [Etcd] Required

Yes, yes! We're limiting stuff. I know, that's not good but as [Golang.hr] wants one day to
have it's own services, it's logical that we have configuration management on one place that's
resilient and just works. Due to that, you'll have to bring up [Etcd] instance in order to use this service.

In order to install [Etcd] please go to [Etcd Git] and read through their README


## Installation
----

I'm sorry but in this moment I have no idea how to set it up on Windows. In fact, I've never used
[Go] on Windows so please, if you did and have will to help us out, add it and make pull request :)

Installation for rest of the world should be similar if not the same.

#### From Source

```sh
# Write to the .profile
export SERVICE_SLACK_NAME="Slack Invite"
export SERVICE_SLACK_DESCRIPTION="Golang.hr Slack Invite is a small automation service written on top of Golang.hr Platform."
export SERVICE_SLACK_VERSION="0.0.1a"
export SERVICE_SLACK_SERVER_ADDR=":4010"

go get -u github.com/golanghr/slack-invite

# This is just dirty way of how to run it. For something more sophisticated you'll need to
# user or systemd or supervisor or upstart or something else.
./slack-invite
```

#### Docker Image

```sh

```



[Golang.hr]: <https://github.com/golanghr>
[Golang.hr Slack Invite]: <https://github.com/golanghr/slack-invite>
[Golang.hr Platform]: <https://github.com/golanghr/platform>

[Etcd]: <https://coreos.com/etcd/>
[Etcd Git]: <https://github.com/coreos/etcd>

[Go]: <http://golang.org/>
[Go Getting Started]: <https://golang.org/doc/install>
