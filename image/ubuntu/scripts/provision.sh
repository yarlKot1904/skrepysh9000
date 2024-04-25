#!/bin/bash

mkdir -p /etc/skrepysh/
cat <<EOF >> /etc/skrepysh/config.yaml
log:
  level: DEBUG
EOF

cat <<EOF > $wd/etc/systemd/system/skrepysh-agent.service
[Unit]
Description=Skrepysh Agent
After=docker.service
Requires=docker.service

[Service]
Type=simple
ExecStart=/usr/local/bin/skrepysh-agent serve
Restart=never

[Install]
WantedBy=multi-user.target
EOF
