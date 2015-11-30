[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/golanghr/slack-invite/tree/master/LICENSE)
[![Build Status](https://travis-ci.org/golanghr/slack-invite.svg)](https://travis-ci.org/golanghr/slack-invite)
[![Go 1.4 Ready](https://img.shields.io/badge/Go%201.4-Ready-green.svg?style=flat)]()
[![Go 1.5 Ready](https://img.shields.io/badge/Go%201.5-Ready-green.svg?style=flat)]()

# [Golang.hr] Slack Invite

[Golang.hr Slack Invite] is [slack] automated invitation service written on top of [Golang.hr Platform].

**Under Heavy Development. Still not even close to consider it for stable/production environments** 

Point of the service is for you to be able quickly setup and run [slack] invitation system
that will send out invitation email to the customer if he's not already registered under
your team.

Service supports REST API and gRPC endpoint. Additionally, it does have nice UI that can be used by WWW.
REST and gRPC are here mainly to help out with frontend but never the less, they are available.
Read me on how to access gRPC as REST will be provided soon.

Additional note is that we're using [AngularJS] and [Bootstrap] for frontend.

### Demo?

Can be seen at [slack.golang.hr]

### Installation Guide

Installation guide can bee seen at [Golang.hr Slack Invite Installation Guide]

### Contributions

Please make sure to read [Golang.hr Slack Invite Contribution Guide] if you are interested into
doing some code contributions :)

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
FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

[Golang.hr]: <https://github.com/golanghr>
[Golang.hr Slack Invite]: <https://github.com/golanghr/slack-invite>
[Golang.hr Slack Invite Installation Guide]: <https://github.com/golanghr/slack-invite/blob/master/INSTALL.md>
[Golang.hr Slack Invite Contribution Guide]: <https://github.com/golanghr/slack-invite/blob/master/CONTRIBUTING.md>
[Golang.hr Platform]: <https://github.com/golanghr/platform>
[filing an issue]: <https://github.com/golanghr/slack-invite/issues/new>

[Golang.hr Slack]: <http://slack.golang.hr>
[Golang.hr Facebook]: <https://www.facebook.com/groups/golanghr/>

[slack]: <https://slack.com/>
[slack.golang.hr]: <http://slack.golang.hr>

[AngularJS]: <https://angularjs.org/>
[Bootstrap]: <http://getbootstrap.com/>
