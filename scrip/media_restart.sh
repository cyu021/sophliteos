#!/bin/bash

# 全局变量用于存储IP地址
declare -g local_ip=""

function get_ip() {
    local active_interface
    active_interface=$(ip route get 8.8.8.8 | awk 'NR==1 {print $5}')
    
    if [[ -z $active_interface ]]; then
        echo "无法自动获取活动网络接口，请手动输入本机IP地址："
        read user_input_ip
        check_ip "$user_input_ip"
    else
        local ip_address
        ip_address=$(ip -4 addr show dev "$active_interface" | grep -oP '(?<=inet\s)\d+(\.\d+){3}')
        echo "获取到本机IP地址： $ip_address"
        confirm_ip "$ip_address"
    fi
}

function confirm_ip() {
    local ip="$1"
    while true; do
        echo "您的IP地址是 $ip。是否正确？(y/n)"
        read -r answer
        case $answer in
            [Yy]*) local_ip=$ip; break ;;
            [Nn]*) echo "请手动输入IP地址："; read -r new_ip; check_ip "$new_ip"; break ;;
            *) echo "请输入 y 或 n." ;;
        esac
    done
}

# 函数用于检查IP地址是否合法
function check_ip() {
    local ip="$1"
    local regex="^([0-9]{1,3}\.){3}[0-9]{1,3}$"

    if [[ $ip =~ $regex ]]; then
        IFS='.' read -ra ip_parts <<< "$ip"
        valid=true

        for part in "${ip_parts[@]}"; do
            if ((part < 0 || part > 255)); then
                valid=false
                break
            fi
        done

        if [ "$valid" = true ]; then
            local_ip=$ip
            echo "IP地址已设置为: $local_ip"
            return 0  # 返回0表示真
        else
            echo "IP地址 $ip 不合法，每个段的取值范围应在0-255之间。"
            exit 1
        fi
    else
        echo "IP地址 $ip 不合法，格式不正确。"
        exit 1
    fi
}

get_ip
echo "IP地址：$local_ip"


echo "修改/etc/hosts....."
sed -i '/clServer/d' /etc/hosts
sed -i '/mediagateway/d' /etc/hosts
echo $local_ip clServer >> /etc/hosts
echo $local_ip mediagateway >> /etc/hosts
echo "修改完成"


echo "删除流媒体容器....."
sudo docker run -d -p 4406:3306 --restart always  --name media-mysql -e MYSQL_ROOT_PASSWORD='Bitmain_(root123)'  ubuntu/mysql:latest

sudo docker run -itd --restart always  --name media_gateway --net host media_gateway -bind_ip 127.0.0.1 -db_ip 127.0.0.1 -db_port 4406 -db_user root -db_pass 'Bitmain_(root123)'

sudo docker run -itd --restart always   --name device_detector --net host --restart always device_detector

sudo docker run -itd  --restart always --name license_service --restart always -p 26085:10085 license_service
echo "流媒体容器启动成功"