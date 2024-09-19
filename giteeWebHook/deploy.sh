#!/bin/bash

# 定义路径
path="/opt/webhook/blue-whale/"
deployPath="/opt/www/blueWhale/"
gitPath="git@gitee.com:jzwPro/blue-whale.git";

# 检查文件夹是否存在
if [ ! -d "${path}" ]; then
    echo "Directory does not exist. Cloning the repository..."
    # 从 Git 仓库克隆项目
    git clone ${gitPath} "${path}"
else
    echo "Directory exists. Pulling the latest changes..."
fi

cd "${path}" || exit
git pull

# 生成 Hugo 静态文件
/opt/hugo/hugo

# 删除目标目录及其内容（如果存在）
rm -rf "${deployPath}"
mkdir -p "${deployPath}"

# 复制生成的静态文件到目标目录
cp -r ./public/* "${deployPath}"

echo "Deployment completed successfully."
