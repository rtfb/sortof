FROM ubuntu:22.04

RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y \
    build-essential \
    curl \
    git \
    make \
    python3 \
    python3-virtualenv \
    wget

RUN adduser --disabled-password --gecos '' rtfb
RUN adduser rtfb sudo
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

RUN mkdir -p /app
RUN chown -R rtfb:rtfb /app
RUN mkdir -p /home/rtfb/dev /home/rtfb/.ssh
RUN chown -R rtfb:rtfb /home/rtfb/dev /home/rtfb/.ssh

USER rtfb
WORKDIR /app

ADD run.sh /app/run.sh
RUN pip3 install --no-warn-script-location huggingface-hub
RUN git clone https://github.com/ggerganov/llama.cpp
RUN cd llama.cpp && make -j -k

RUN echo 'PATH=$PATH:/home/rtfb/.local/bin' >> /home/rtfb/.bashrc

RUN /home/rtfb/.local/bin/huggingface-cli download \
    TheBloke/Nous-Capybara-34B-GGUF nous-capybara-34b.Q4_K_M.gguf \
    --local-dir . \
    --local-dir-use-symlinks False

ENTRYPOINT ["/bin/bash"]
