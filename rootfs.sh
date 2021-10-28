#!/bin/sh

IMAGE_TAG=container-gameserver:latest

docker build -t ${IMAGE_TAG} .
docker export $(docker create ${IMAGE_TAG}) -o rootfs.tar
