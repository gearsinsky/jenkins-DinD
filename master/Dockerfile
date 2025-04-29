# 使用 jenkins 官方的 lts 版本作為 base image
FROM jenkins/jenkins:lts

# 進入 root 權限，方便安裝其他套件
USER root

# 安裝 Docker CLI（選擇合適的版本，這裡安裝最常見的 docker.io）
RUN apt-get update && apt-get install -y docker.io
RUN usermod -aG docker jenkins
# 也可以選擇只安裝 docker cli
RUN apt-get update && apt-get install -y docker-cli

# 如果有其他需要安裝的套件，像是工具或者插件，可以在這裡添加
# RUN apt-get install -y some_package

# 設定 Jenkins 使用的工作目錄
USER jenkins

