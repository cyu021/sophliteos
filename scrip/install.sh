#!/bin/bash

set -e

if [[ -f /etc/systemd/system/algoliteos.service ]]; then
  systemctl stop algoliteos.service
  systemctl disable algoliteos.service
fi
rm -rf /data/pictures
mkdir -p /etc/sophliteos/config /var/log/sophliteos /data/pictures

cp algoliteos /bin
cp config/algoliteos.yaml /etc/sophliteos/config
cp database/algoliteos.db /var/lib/sophliteos/db
cp algoliteos.service /etc/systemd/system/

systemctl daemon-reload
systemctl enable algoliteos.service
systemctl start algoliteos.service