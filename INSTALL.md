# [Golang.hr] Slack Invite Installation Guide

**TBD**

#### Etcd Required

Yes, yes! We're limiting stuff. I know, that's not good but as Golang.hr wants one day to
have it's own services, it's logical that we have configuration management on one place that's
resiliant and just works. Due to that, you'll have to bring up etcd instance too.

Luckly, that should be easy :D Here you can find how to do it...

```sh
pushd /path/to/slack-invite
  go get -u github.com/coreos/etcd
  go get -u github.com/mattn/goreman
  wget https://raw.githubusercontent.com/coreos/etcd/master/Procfile
  goreman start
popd
```
That's about it! You now have as etcd README says

```
This will bring up 3 etcd members infra1, infra2 and infra3 and etcd proxy proxy, which runs locally and composes a cluster.
You can write a key to the cluster and retrieve the value back from any member or proxy.
```

[Golang.hr]: <https://github.com/golanghr>
[Golang.hr Slack Invite]: <https://github.com/golanghr/slack-invite>
[Golang.hr Platform]: <https://github.com/golanghr/platform>
