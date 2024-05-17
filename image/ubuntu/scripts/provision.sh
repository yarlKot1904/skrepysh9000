#!/bin/bash

mkdir -p /etc/skrepysh/
cat <<EOF >> /etc/skrepysh/config.yaml
log:
  level: DEBUG

server-port: 48934

node-exporter:
  host: localhost
  port: 9100

skrepysh-backend:
  host: 192.168.64.1
  port: 8080
EOF

cat <<EOF > $wd/etc/systemd/system/skrepysh-agent.service
[Unit]
Description=Skrepysh Agent
After=docker.service
Requires=node-exporter.service

[Service]
Type=simple
ExecStart=/usr/local/bin/skrepysh-agent serve
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
EOF
systemctl enable skrepysh-agent.service

cat <<EOF > $wd/etc/systemd/system/node-exporter.service
[Unit]
Description=Node Exporter
After=docker.service

[Service]
Type=simple
ExecStart=/usr/local/bin/node_exporter
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
EOF
systemctl enable node-exporter.service