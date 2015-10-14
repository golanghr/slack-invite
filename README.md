[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/golanghr/slack-invite/tree/master/LICENSE)
[![Build Status](https://travis-ci.org/golanghr/slack-invite.svg)](https://travis-ci.org/golanghr/slack-invite)
[![Go 1.3 Ready](https://img.shields.io/badge/Go%201.3-Ready-green.svg?style=flat)]()
[![Go 1.4 Ready](https://img.shields.io/badge/Go%201.4-Ready-green.svg?style=flat)]()
[![Go 1.5 Ready](https://img.shields.io/badge/Go%201.5-Ready-green.svg?style=flat)]()

# [Golang.hr] Slack Invite

[Golang.hr Slack Invite] is a small automation service written on top of [Golang.hr Platform].

Point of the service is for you to be able quickly setup and run [slack] invitation system
that will send out invitation email to the customer if he's not already registered under
your team.


### Installing Service

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

### License

```
The MIT License (MIT)

Copyright (c) 2015 Golang Hrvatska (Croatia)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

[Golang.hr]: <https://github.com/golanghr>
[Golang.hr Slack Invite]: <https://github.com/golanghr/slack-invite>
[Golang.hr Platform]: <https://github.com/golanghr/platform>
[filing an issue]: <https://github.com/golanghr/slack-invite/issues/new>

[Golang.hr Slack]: <http://slack.golang.hr>
[Golang.hr Facebook]: <https://www.facebook.com/groups/golanghr/>

[slack]: <https://slack.com/>
