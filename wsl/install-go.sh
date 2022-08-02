#!/bin/bash

sudo rm -rf /usr/local/go* && sudo rm -rf /usr/local/go

VERSION=1.18
OS=linux
ARCH=amd64

cd $HOME
wget https://storage.googleapis.com/golang/go$VERSION.$OS-$ARCH.tar.gz
tar -xvf go$VERSION.$OS-$ARCH.tar.gz
sudo mv go /usr/local