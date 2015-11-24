FROM debian:wheezy

ENV DEBIAN_FRONTEND noninteractive
ENV TERM dumb

RUN echo "deb http://ftp.us.debian.org/debian/ wheezy main contrib non-free" >> /etc/apt/sources.list
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

ENV SLACK_SERVICE_NAME "Slack Invite"
ENV SLACK_SERVICE_DESCRIPTION "Golang.hr Slack Invite is a small automation service written on top of Golang.hr Platform."
ENV SLACK_SERVICE_VERSION "0.1"
ENV SLACK_SERVICE_MANAGER_INTERRUPT_TIMEOUT "10"
ENV SLACK_SERVICE_GRPC_LISTEN_FOREVER "true"
ENV SLACK_SERVICE_GRPC_ADDR ":4772"
ENV SLACK_SERVICE_GRPC_TLS "true"
ENV SLACK_SERVICE_GRPC_TLS_DOMAIN "golang.hr"

ADD test_data/server.crt /certs/server.crt
ADD test_data/server.key /certs/server.key

ENV SLACK_SERVICE_GRPC_TLS_CERT "/certs/server.crt"
ENV SLACK_SERVICE_GRPC_TLS_KEY "/certs/server.key"

ENV SLACK_SERVICE_HTTP_ADDR ":8500"
ENV SLACK_SERVICE_HTTP_LISTEN_FOREVER "true"

EXPOSE 4772
EXPOSE 8500

ADD build/platform-slack-invite /usr/bin/platform-slack-invite
ENTRYPOINT ["platform-slack-invite"]
