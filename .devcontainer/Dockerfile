FROM mcr.microsoft.com/vscode/devcontainers/typescript-node:16-buster

ARG GO_VERSION=1.17.5
ARG GO_MIGRATE_VERSION=4.15.1

RUN DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get -y install sqlite3 \
    && DEBIAN_FRONTEND=dialog

RUN npm install -g @angular/cli

RUN wget https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz

RUN cd /usr/local/bin \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v$GO_MIGRATE_VERSION/migrate.linux-arm64.tar.gz | tar xvz \
    && rm README.md

ENV PATH="/usr/local/go/bin:${PATH}"

EXPOSE 4200
EXPOSE 8080