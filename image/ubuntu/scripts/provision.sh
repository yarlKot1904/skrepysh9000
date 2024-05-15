#!/bin/bash

mkdir -p /etc/skrepysh/
cat <<EOF >> /etc/skrepysh/config.yaml
log:
  level: DEBUG

server-port: 48934

node-exporter:
  host: localhost
  port: 9100

EOF

cat <<EOF > $wd/etc/systemd/system/skrepysh-agent.service
[Unit]
Description=Skrepysh Agent
After=docker.service

[Service]
Type=simple
ExecStart=/usr/local/bin/skrepysh-agent serve
Restart=always

[Install]
WantedBy=multi-user.target
EOF

cat <<EOF > $wd/etc/systemd/system/node-exporter.service
[Unit]
Description=Node Exporter
After=docker.service
Requires=skrepysh-agent.service

[Service]
Type=simple
ExecStart=/usr/local/bin/node_exporter
Restart=always

[Install]
WantedBy=multi-user.target
EOF
