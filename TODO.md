# [Golang.hr] Slack Invite TODO

While building [Golang.hr Platform] and [Golang.hr Slack Invite] I've encounter with
problematic design concepts that **must** be taken care off before we can treat this
service as stable and operational.

- Make sure GRPC/REST handlers are properly invoked. Get rid of [grpc-gateway] or at least minimize
  use of it as much as possible. It's adding layers of complexity that should be avoided.
- WebSocket support instead of plain REST API call. We want to push changes, not pull for changes.




[Golang.hr]: <https://github.com/golanghr>
[Golang.hr Slack Invite]: <https://github.com/golanghr/slack-invite>
[Golang.hr Platform]: <https://github.com/golanghr/platform>

[Etcd]: <https://coreos.com/etcd/>
[Etcd Git]: <https://github.com/coreos/etcd>

[Go]: <http://golang.org/>
[Go Getting Started]: <https://golang.org/doc/install>

[grpc-gateway]: <https://github.com/gengo/grpc-gateway/>
