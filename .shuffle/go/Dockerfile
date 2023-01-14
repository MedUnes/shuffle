ARG GO_VERSION

FROM golang:${GO_VERSION}-bullseye

ARG APP_PATH

ENV TZ="Europe/Berlin"
ENV LC_ALL="C.UTF-8"
ENV LANG="C.UTF-8"

RUN apt update && \
    apt install -y --no-install-recommends \
    procps \
    zip \
    unzip \
    git \
    jq \
    net-tools \
    dnsutils \
    bash-completion \
    nano \
    wget \
    netcat \
    sudo
RUN  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

RUN set -ex; \
    echo "PS1='\h:\w\$ '" >> /etc/bash.bashrc; \
    echo "alias ls='ls --color=auto'" >> /etc/bash.bashrc; \
    echo "alias grep='grep --color=auto'" >> /etc/bash.bashrc;

RUN useradd --create-home --shell /bin/bash docker \
    && passwd docker -d \
    && adduser docker sudo

RUN echo "docker ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

WORKDIR "${APP_PATH}"

EXPOSE 80
STOPSIGNAL SIGQUIT

CMD tail -f /dev/null
