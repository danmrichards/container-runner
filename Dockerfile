FROM ubuntu:18.04

RUN useradd -ms /bin/sh mpukgame

USER mpukgame
WORKDIR /home/mpukgame
