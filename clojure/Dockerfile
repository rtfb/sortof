FROM ubuntu:22.04

# Conjured this Dockerfile following this installation tutorial:
# https://ericnormand.me/guide/how-to-install-clojure

RUN apt-get update && apt-get install -y \
    curl \
    default-jre \
    default-jdk \
    git \
    leiningen \
    rlwrap \
    sudo \
    vim \
    wget

RUN curl https://download.clojure.org/install/linux-install-1.11.1.1200.sh \
        -o /tmp/linux-install-1.11.1.1200.sh && \
    chmod +x /tmp/linux-install-1.11.1.1200.sh && \
    /tmp/linux-install-1.11.1.1200.sh

RUN curl https://raw.githubusercontent.com/technomancy/leiningen/stable/bin/lein \
        > /usr/local/bin/lein && \
    chmod a+x /usr/local/bin/lein && \
    lein version

RUN adduser --disabled-password --gecos '' rtfb
RUN adduser rtfb sudo
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER rtfb

WORKDIR /home/rtfb

# ENTRYPOINT ["/usr/local/bin/lein", "repl"]
ENTRYPOINT ["/bin/bash"]
