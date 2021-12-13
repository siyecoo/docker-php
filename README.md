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
│   ├── php                     PHP   配置目录 (默认是PHP7.3版本)
│   ├── php72                   PHP72 配置目录(这个是需要自己创建的,多版本的时候使用)
│   ├── php71                   PHP71 配置目录(这个是需要自己创建的,多版本的时候使用)
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

5. 如果想要创建多个php版本：
    ```
    在docker-compose.yml 把PHP7.2 注释打开
    services文件夹复制一份php文件夹,修改为php72              #注意事项:Dockerfile 需要修改相应的配置文件名称已免重复版本
    docker-compose up -d  修改为php72                    #即可启动
    ```

 
## 3.PHP和扩展
### 3.1 切换Nginx使用的PHP版本

NGINX中配置如果多个项目想使用不同的PHP版本的话
1.在services/nginx/conf.d 配置 多个conf.conf的时候  

      location ~ \.php$ {
            fastcgi_pass   php:9000;             这里的fastcgi_pass  php:9000  如想切换到 php.72 php72:9000  这里为什么写php:9000 这指的是我们容器的名称  跟127.0.0.1:9000 是一个道理 为什么使用容器 可以百度了解一下 容器之间如何关联起来
            include        fastcgi-php.conf;
            include        fastcgi_params;
        }

然后配置好之后 需要 **重启 Nginx** 生效。 
在宿主机上 执行 命令 :docker-exec -it nginx nginx -t (先查询下是否配置没问题) docker-exec -it nginx nginx -s reload

这里两个`nginx`，第一个是容器名，第二个是容器中的`nginx`程序。


### 3.2 安装PHP扩展
PHP的很多功能都是通过扩展实现，而安装扩展是一个略费时间的过程，
所以，除PHP内置扩展外，在`env.sample`文件中我们仅默认安装少量扩展，
如果要安装更多扩展，请打开你的`.env`文件修改如下的PHP配置，
增加需要的PHP扩展：
```bash
PHP_EXTENSIONS=pdo_mysql,opcache,redis       # PHP 要安装的扩展列表，英文逗号隔开
PHP54_EXTENSIONS=opcache,redis                 # PHP 5.4要安装的扩展列表，英文逗号隔开
```
然后重新build PHP镜像。
```bash
docker-compose build php
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
输出的`NAMES`那一列就是容器的名称，如果使用默认配置，那么名称就是`nginx`、`php`、`php56`、`mysql`等。


用于填写`extra_hosts`容器访问宿主机的`hosts`地址



## 4 经常遇到的问题


这要分两种情况，

第一种情况，在**PHP代码中**。
```php
// 连接MySQL
$dbh = new PDO('mysql:host=mysql;dbname=mysql', 'root', '123456');

// 连接Redis
$redis = new Redis();
$redis->connect('redis', 6379);
```

第二种情况，**在主机中**通过**命令行**或者**Navicat**等工具连接。主机要连接mysql和redis的话，要求容器必须经过`ports`把端口映射到主机了。以 mysql 为例，`docker-compose.yml`文件中有这样的`ports`配置：`3306:3306`，就是主机的3306和容器的3306端口形成了映射，所以我们可以这样连接：
```bash
$ mysql -h127.0.0.1 -uroot -p123456 -P3306
$ redis-cli -h127.0.0.1
```
这里`host`参数不能用localhost是因为它默认是通过sock文件与mysql通信，而容器与主机文件系统已经隔离，所以需要通过TCP方式连接，所以需要指定IP。



## 5.使用Log


### 5.1 Nginx日志
Nginx日志是我们用得最多的日志，所以我们单独放在根目录`log`下。

`log`会目录映射Nginx容器的`/var/log/nginx`目录，所以在Nginx配置文件中，需要输出log的位置，我们需要配置到`/var/log/nginx`目录，如：
```
error_log  /var/log/nginx/nginx.localhost.error.log  warn;
```


### 5.2 PHP-FPM日志
大部分情况下，PHP-FPM的日志都会输出到Nginx的日志中，所以不需要额外配置。

另外，建议直接在PHP中打开错误日志：
```php
error_reporting(E_ALL);
ini_set('error_reporting', 'on');
ini_set('display_errors', 'on');
```

如果确实需要，可按一下步骤开启（在容器中）。

1. 进入容器，创建日志文件并修改权限：
    ```bash
    $ docker exec -it php /bin/sh
    $ mkdir /var/log/php
    $ cd /var/log/php
    $ touch php-fpm.error.log
    $ chmod a+w php-fpm.error.log
    ```
2. 主机上打开并修改PHP-FPM的配置文件`conf/php-fpm.conf`，找到如下一行，删除注释，并改值为：
    ```
    php_admin_value[error_log] = /var/log/php/php-fpm.error.log
    ```
3. 重启PHP-FPM容器。

### 5.3 MySQL日志
因为MySQL容器中的MySQL使用的是`mysql`用户启动，它无法自行在`/var/log`下的增加日志文件。所以，我们把MySQL的日志放在与data一样的目录，即项目的`mysql`目录下，对应容器中的`/var/lib/mysql/`目录。
```bash
slow-query-log-file     = /var/lib/mysql/mysql.slow.log
log-error               = /var/lib/mysql/mysql.error.log
```
以上是mysql.conf中的日志文件的配置。

## License
MIT