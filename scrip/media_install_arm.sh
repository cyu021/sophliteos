#!/bin/bash

local_ip="127.0.0.1"

echo "修改/etc/hosts....."
sed -i '/clServer/d' /etc/hosts
sed -i '/mediagateway/d' /etc/hosts
echo $local_ip clServer >> /etc/hosts
echo $local_ip mediagateway >> /etc/hosts
echo "修改完成"

echo "导入镜像....."
sudo docker load -i license_service.tar.gz
sudo docker load -i media_gateway.tar.gz
sudo docker load -i device_detector.tar.gz
echo "导入镜像完成"


echo "启动流媒体容器....."

sudo docker run -itd --restart always   --name device_detector --net host  device_detector

sudo docker run -itd  --restart always --name license_service  -p 26085:10085 license_service

echo "流媒体容器启动成功"

