#!/bin/bash

wget https://golang.org/dl/go1.17.1.linux-amd64.tar.gz

tar -xvf go1.*

cp -rv go /usr/bin/

echo "export PATH=$PATH:/usr/bin/go/bin" >> ~/.bashrc

source ~/.bashrc

go version
