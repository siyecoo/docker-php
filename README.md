DOCKER-PHP（Docker + Nginx + MySQL + PHP7（多个版本 7.3/7.2/7.1） + Redis）是一款全功能的**LNMP一键安装程序**。

DOCKER-PHP项目特点：

1. 支持**多版本PHP**共存，可任意切换（PHP7.1、PHP7.2、PHP7.3) （**3.1 切换Nginx使用的PHP版本**）
2. 可一键选配常用服务：
    - 多PHP版本：PHP7.1、PHP7.2、PHP7.3
    - Web服务：Nginx
    - 数据库：MySQL5、Redis、


    
## 1.目录结构

```
/
├── data                        数据库数据目录
│   └── mysql                   MySQL5 数据目录
├── services                    服务构建文件和配置文件目录
│   ├── mysql                   MySQL 配置文件目录
│   ├── nginx                   Nginx 配置文件目录
│   ├── php73                   PHP73 配置目录 
│   ├── php72                   PHP72 配置目录
│   ├── php71                   PHP71 配置目录
│   └── redis                   Redis 配置目录
├── logs                        日志目录
├── docker-compose.sample.yml   Docker 服务配置示例文件
├── env.smaple                  环境配置示例文件
└── www                         PHP 代码目录
```    


## 2.快速使用

1.  首先clone项目 https://github.com/siyecoo/docker-php.git
2.  如果当前不是root用户,需要先将用户加入**docker**用户组  sudo gpasswd -a ${USER} docker
3.  启动项目操作流程:
    ````
        1.cd docker-php 目录
        2.cp env.sample .env  docker-compose-sample.yml docker-compose.yml  (需要配置一下环境变量,docker-compose.yml可以开启一些注释的PHP多版本,如果只想单纯版本直接复制一份即可)
        3.docker-compose up  -d    启动 (-d是后台执行,docker-compose up -d nginx redis php mysql 也可以这样执行,看你需要启动的镜像容器需要哪些 也可以单个)

4. 默认已经帮配置了localhost演示环境  `http://localhost`就能看到效果,代码在www目录下的localhost


然后重新build PHP镜像。
```bash
docker-compose build php71     #php72 php73 3个版本都可以
```
可用的扩展请看同文件的`env.sample`注释块说明。


### 3.3 使用composer

在容器内使用composer命令 

```bash
docker exec -it php /bin/sh
cd /www/localhost
composer update
```
也可以在宿主机上使用composer 需要自行设置环境变量(这里需要关联 例如 alias 可自行百度)

## 4.管理命令
### 4.1 服务器启动和构建命令
如需管理服务，请在命令后面加上服务器名称，例如：
```bash
$ docker-compose up                         # 创建并且启动所有容器
$ docker-compose up -d                      # 创建并且后台运行方式启动所有容器
$ docker-compose up nginx php mysql         # 创建并且启动nginx、php、mysql的多个容器
$ docker-compose up -d nginx php  mysql     # 创建并且已后台运行的方式启动nginx、php、mysql容器


$ docker-compose start php                  # 启动服务
$ docker-compose stop php                   # 停止服务
$ docker-compose restart php                # 重启服务
$ docker-compose build php                  # 构建或者重新构建服务

$ docker-compose rm php                     # 删除并且停止php容器
$ docker-compose down                       # 停止并删除容器，网络，图像和挂载卷
```

首先，在主机中查看可用的容器：
```bash
$ docker ps           # 查看所有运行中的容器
$ docker ps -a        # 所有容器
```
输出的`NAMES`那一列就是容器的名称，如果使用默认配置，那么名称就是`nginx`、`php71`、`php72`、`php73`、`mysql`等。










## License
MIT