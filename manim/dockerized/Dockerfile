FROM ubuntu:22.04

RUN apt-get update
RUN apt-get install -y \
    ffmpeg \
    libcairo2-dev \
    libpango1.0-dev \
    python3 \
    python3-cairo \
    python3-virtualenv \
    sudo \
    vim

RUN mkdir -p /app
ADD requirements.txt /app/requirements.txt

WORKDIR /app

RUN adduser --disabled-password --gecos '' rtfb
RUN adduser rtfb sudo
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

RUN chown -R rtfb:rtfb /app
RUN mkdir -p /home/rtfb/dev
RUN chown -R rtfb:rtfb /home/rtfb/dev

USER rtfb

RUN echo 'PATH=$PATH:/home/rtfb/.local/bin' >> /home/rtfb/.bashrc
RUN pip3 install --no-warn-script-location -r requirements.txt

WORKDIR /home/rtfb/dev

ENTRYPOINT ["/bin/bash"]
