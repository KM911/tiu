FROM ubuntu:20.04

# Install dependencies

# 或许正确的做法是我们将go的环境放到这里来 然后将其配置好
# 拷贝文件
COPY ./ /app

# Set the working directory
WORKDIR /app

# Install dependencies



CMD ["echo hello world"]
