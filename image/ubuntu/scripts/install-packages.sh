#!/bin/bash

apt install apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"
apt-cache policy docker-ce
apt install docker-ce
systemctl status docker
sudo usermod -aG docker ${USER}

apt install tar wget
wget -P /tmp https://github.com/prometheus/node_exporter/releases/download/v1.8.0/node_exporter-1.8.0.linux-amd64.tar.gz
tar -C /usr/local/bin -xvf /tmp/node_exporter-1.8.0.linux-amd64.tar.gz
mv /usr/local/bin/node_exporter-1.8.0.linux-amd64/node_exporter /usr/local/bin
rm -r /usr/local/bin/node_exporter-1.8.0.linux-amd64
rm /tmp/node_exporter-1.8.0.linux-amd64.tar.gz


