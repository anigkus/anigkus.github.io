<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >手动编译开源网关konga与kong集成</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Manually compile open source gateway konga and integrate with kong</h1><br/>]:#

<center>
<img src="../assets/images/manually-compile-open-source-gateway-konga-and-integrate-with-kong/figure-1.jpeg" alt="Manually compile open source gateway konga and integrate with kong" title="Github of Anigkus" >
</center>


> <br/>&nbsp;&nbsp;&nbsp;&nbsp; [`Kong`](https://docs.konghq.com/) 由Mashape公司开发的一款云原生、平台无关、可扩展的 API 网关,是基于OpenResty(Nginx + Lua）而编写的.而 [`konga`](https://github.com/pantsel/konga) 是社区贡献的,主要是为了更好和方便管理和操作kong而提供的一套Web系统. kong 和 konga 都无论是社区还是官方都提供容器部署版本,但是有时候我们的客户没法使用容器化方案,因此就需要使用二进制来安装.kong官方提供了rpm和多种系统的二进制安装教程,而konga就需要自己编译了.因此本文主要是简单说明下如何使用二进制的方式来部署来部署.<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; Some general notes on article.<br/>]:#
[> <br/>]:#


# 安装 PostgreSQL

&nbsp;&nbsp;&nbsp;&nbsp; 因为Kong只支持Cassandra或PostgreSQL,我这里就使用PG了.也可以看[官方安装文档](https://www.postgresql.org/download/linux/redhat/#yum).我这里使用二进制解压的方式来安装 Postgresql.


## 下载文件和解压

```
$ yum install gcc -y

$ yum install zlib-devel -y

$ yum install make -y

$ yum install readline eadline-devel -y

$ wget https://download.postgresql.org/pub/source/v11.2/postgresql-11.2.tar.gz

$ tar -vxf postgresql-11.2.tar.gz

$ cd postgresql-11.2

$ ./configure --prefix=/usr/local/postgresql 

$ make && make install
...
make[1]: Leaving directory '/root/postgresql-11.2/config'
PostgreSQL installation complete.
```

## 配置环境和用户

```
$ mkdir -p /data/postgresql-11.2/{data,logs,pg_archive} 

$ vi + /etc/profile
PGHOME=/usr/local/postgresql
export PGHOME
PGDATA=/data/postgresql-11.2/data
export PGDATA
PATH=$PATH:$HOME/.local/bin:$HOME/bin:$PGHOME/bin
export PATH

$ source /etc/profile

# useradd postgres

$ chown -R postgres:postgres /data/postgresql-11.2

$ chown -R postgres:postgres /usr/local/postgresql

$ vi /data/postgresql-11.2/data/postgresql.conf
listen_addresses = '*'

$ vi /data/postgresql-11.2/data/pg_hba.conf
 host  all  all 0.0.0.0/0 md5

```

## 首次初始数据库

```
# su - postgres

$ cd /usr/local/postgresql/bin

$ ./initdb -D /data/postgresql-11.2/data
The files belonging to this database system will be owned by user "postgres".
This user must also own the server process.

The database cluster will be initialized with locale "en_US.UTF-8".
The default database encoding has accordingly been set to "UTF8".
The default text search configuration will be set to "english".

Data page checksums are disabled.

fixing permissions on existing directory /data/postgresql-11.2/data ... ok
creating subdirectories ... ok
selecting default max_connections ... 100
selecting default shared_buffers ... 128MB
selecting dynamic shared memory implementation ... posix
creating configuration files ... ok
running bootstrap script ... ok
performing post-bootstrap initialization ... ok
syncing data to disk ... ok

WARNING: enabling "trust" authentication for local connections
You can change this by editing pg_hba.conf or using the option -A, or
--auth-local and --auth-host, the next time you run initdb.

Success. You can now start the database server using:

    ./pg_ctl -D /data/postgresql-11.2/data -l logfile start

$ psql -V
psql (PostgreSQL) 11.2

```

## 启动数据库服务

```
$ ./pg_ctl -D /data/postgresql-11.2/data -l logfile start

$ ./pg_ctl stop 
```

## 配置账号和密码

```
$ su postgres
$ psql

# Change postgres password of [postgres123456]
postgres-# \password postgres
Enter new password: 
Enter it again: 
postgres=# CREATE DATABASE kong;
CREATE DATABASE
postgres=# CREATE DATABASE konga;
CREATE DATABASE
postgres=# CREATE ROLE kong WITH SUPERUSER LOGIN PASSWORD 'kong123456';
CREATE ROLE
postgres=# CREATE ROLE konga WITH SUPERUSER LOGIN PASSWORD 'konga123456';
CREATE ROLE
postgres=# GRANT ALL PRIVILEGES ON DATABASE kong to kong;
GRANT
postgres=# GRANT ALL PRIVILEGES ON DATABASE konga to konga;
GRANT
# 至关重要,主从复制或者和 patroni 都用到了这2个玩意(如果已经同步了执行这个DDL语句,会自动复制到所有的slave节点上的),这里可以不需要忽略即可.
postgres=# create user replipuser replication login encrypted password 'replipuser123456';
postgres=# select * from pg_roles; #查询所有角色
postgres=# select * from pg_user; #查询所有用户

postgres=# exit
$ exit
$
```

## 做成启动服务项

```
$ su  # root

$ cd  /root/postgresql-11.2/contrib/start-scripts/ # 源码包目录

$ chmod a+x linux

$ cp linux /etc/init.d/postgresql


$ vi /etc/init.d/postgresql

# Old
prefix=/usr/local/pgsql
PGDATA="/usr/local/pgsql/data"

# New
prefix=/usr/local/postgresql
PGDATA="/data/postgresql-11.2/data"

$ chkconfig --add postgresql  

$  service postgresql start # systemctl start postgresql 

$ 

```

## 容器启动数据库
```
$ docker pull postgres:11.2
$ mkdir -p /soft/postgresql
$ vi /soft/postgresql/start_postgresql.sh

#!/bin/bash

# remove container
docker rm -f postgresql_m1 &> /dev/null

# run container
docker run -d --restart=always --restart=on-failure:3 --privileged=true \
        -p 5432:5432 --memory 2048m  --memory-reservation 1024m -w /data/postgresql \
        -v $PWD/m1/pgdata:/var/lib/postgresql/data \
        -v $PWD/m1/pg_archive/:/var/lib/postgresql/pg_archive \
        -v /etc/localtime:/etc/localtime:ro \
        -v /etc/timezone/timezone:/etc/timezone:ro \
        -e "POSTGRES_USER=kong" \
        -e "POSTGRES_DB=kong" \
        --name postgresql_m1 postgres:11.2

$ chmod +x /soft/start_postgresql.sh
sh start_postgresql.sh 
```

# 安装 Kong

```
$ wget --no-check-certificate  https://download.konghq.com/gateway-2.x-centos-8/Packages/k/kong-2.8.1.el8.amd64.rpm
$  yum install kong-2.8.1.el7.amd64.rpm -y 
```

## 配置 Kong
```
$ kong version
2.8.1

$ cp /etc/kong/kong.conf.default /etc/kong/kong.conf

$ vi /etc/kong/kong.conf

# x.x.x.x 是本机IP
$ grep "^\s*[^# \t].*$" /etc/kong/kong.conf
admin_listen = 127.0.0.1:8001,x.x.x.x:8001 reuseport backlog=16384, 127.0.0.1:8444,x.x.x.x:8444 http2 ssl reuseport backlog=16384
pg_host = localhost             # Host of the Postgres server.
pg_port = 5432                  # Port of the Postgres server.
pg_user = kong                  # Postgres user.
pg_password =kong123456                   # Postgres user's password.
pg_database = kong              # The database name to connect to.

```

## 初始化 kong
```
$ kong migrations bootstrap
2021/10/13 16:31:30 [warn] ulimit is currently set to "1024". For better performance set it to at least "4096" using "ulimit -n"

Bootstrapping database...
migrating core on database 'kong1'...
core migrated up to: 000_base (executed)
core migrated up to: 003_100_to_110 (executed)
core migrated up to: 004_110_to_120 (executed)
core migrated up to: 005_120_to_130 (executed)
core migrated up to: 006_130_to_140 (executed)
core migrated up to: 007_140_to_150 (executed)
core migrated up to: 008_150_to_200 (executed)
core migrated up to: 009_200_to_210 (executed)
core migrated up to: 010_210_to_211 (executed)
core migrated up to: 011_212_to_213 (executed)
core migrated up to: 012_213_to_220 (executed)
core migrated up to: 013_220_to_230 (executed)
migrating acl on database 'kong1'...
acl migrated up to: 000_base_acl (executed)
acl migrated up to: 002_130_to_140 (executed)
acl migrated up to: 003_200_to_210 (executed)
acl migrated up to: 004_212_to_213 (executed)
migrating acme on database 'kong1'...
acme migrated up to: 000_base_acme (executed)
migrating basic-auth on database 'kong1'...
basic-auth migrated up to: 000_base_basic_auth (executed)
basic-auth migrated up to: 002_130_to_140 (executed)
basic-auth migrated up to: 003_200_to_210 (executed)
migrating bot-detection on database 'kong1'...
bot-detection migrated up to: 001_200_to_210 (executed)
migrating hmac-auth on database 'kong1'...
hmac-auth migrated up to: 000_base_hmac_auth (executed)
hmac-auth migrated up to: 002_130_to_140 (executed)
hmac-auth migrated up to: 003_200_to_210 (executed)
migrating ip-restriction on database 'kong1'...
ip-restriction migrated up to: 001_200_to_210 (executed)
migrating jwt on database 'kong1'...
jwt migrated up to: 000_base_jwt (executed)
jwt migrated up to: 002_130_to_140 (executed)
jwt migrated up to: 003_200_to_210 (executed)
migrating key-auth on database 'kong1'...
key-auth migrated up to: 000_base_key_auth (executed)
key-auth migrated up to: 002_130_to_140 (executed)
key-auth migrated up to: 003_200_to_210 (executed)
migrating oauth2 on database 'kong1'...
oauth2 migrated up to: 000_base_oauth2 (executed)
oauth2 migrated up to: 003_130_to_140 (executed)
oauth2 migrated up to: 004_200_to_210 (executed)
oauth2 migrated up to: 005_210_to_211 (executed)
migrating rate-limiting on database 'kong1'...
rate-limiting migrated up to: 000_base_rate_limiting (executed)
rate-limiting migrated up to: 003_10_to_112 (executed)
rate-limiting migrated up to: 004_200_to_210 (executed)
migrating response-ratelimiting on database 'kong1'...
response-ratelimiting migrated up to: 000_base_response_rate_limiting (executed)
migrating session on database 'kong1'...
session migrated up to: 000_base_session (executed)
session migrated up to: 001_add_ttl_index (executed)
41 migrations processed
41 executed
Database is up-to-date

```
## 启动 kong
```
$ kong start 
2021/10/13 16:31:52 [warn] ulimit is currently set to "1024". For better performance set it to at least "4096" using "ulimit -n"
Kong started

$ ps -ef | grep kong
root      442302       1  0 14:12 ?        00:00:00 nginx: master process /usr/local/openresty/nginx/sbin/nginx -p /usr/local/kong -c nginx.conf
kong      442303  442302  0 14:12 ?        00:00:00 nginx: worker process
kong      442304  442302  0 14:12 ?        00:00:00 nginx: worker process
kong      442305  442302  0 14:12 ?        00:00:00 nginx: worker process
kong      442306  442302  0 14:12 ?        00:00:00 nginx: worker process
postgres  442313  439797  0 14:12 ?        00:00:00 postgres: kong kong 127.0.0.1(42400) idle
postgres  442315  439797  0 14:12 ?        00:00:00 postgres: kong kong 127.0.0.1(42402) idle
postgres  442316  439797  0 14:12 ?        00:00:00 postgres: kong kong 127.0.0.1(42404) idle
root      442320  438512  0 14:12 pts/0    00:00:00 grep --color=auto kong

```

# 安装 Konga

## 安装 Node
```
$ VERSION=v14.18.0

$ DISTRO=linux-x64

$ wget https://nodejs.org/download/release/v14.18.0/node-$VERSION-$DISTRO.tar.xz

$ mkdir -p  /usr/local/lib/nodejs

$ tar -xJvf node-$VERSION-$DISTRO.tar.xz -C /usr/local/lib/nodejs

$ vi ~/.node_profile
VERSION=v14.18.0
DISTRO=linux-x64
export PATH=/usr/local/lib/nodejs/node-$VERSION-$DISTRO/bin:$PATH

$ vi + ~/.bash_profile 

#node
source ~/.node_profile

$ source ~/.bash_profile

$ node -v
v14.18.0

$ npm -v
6.14.15

```
## 配置 Konga
```
$ yum install -g git

$ git clone --branch 0.14.9 https://github.com/pantsel/konga.git

$ cd konga

$ git branch -vv 
* (no branch) 57f40d4 Updated README.md # 0.14.9

#npm command have times not Unstable##
$ npm i

$ npm i pg -g


#OR Use yarn,not required, but required HangKong or foreign internet
$ npm i yarn -g

$ yarn -v
1.22.18

$ yarn
yarn install v1.22.15
warning package-lock.json found. Your project contains lock files generated by tools other than Yarn. It is advised not to mix package managers in order to avoid resolution inconsistencies caused by unsynchronized lock files. To clear this warning, remove package-lock.json.
[1/4] Resolving packages ...
...
...
...
success Already up-to-date.
$ bower --allow-root install
bower bootstrap-switch extra-resolution Unnecessary resolution: bootstrap-switch#~3.3.4
Done in 1.79s.

$ yarn add pg -g
```

&nbsp;&nbsp;&nbsp;&nbsp; 如果没有进行这一步,那么就是使用本地临时存储,在执行`npm run  production`的时候就会输出`No DB Adapter defined. Using localDB...`.

```
$ cp .env_example .env

$ vi .env
PPORT=1337
NODE_ENV=production
KONGA_HOOK_TIMEOUT=120000
DB_ADAPTER=postgres
DB_URI=postgresql://localhost:5432/konga
DB_USER=konga
DB_PASSWORD=konga123456
DB_DATABASE=konga
KONGA_LOG_LEVEL=warn
TOKEN_SECRET=some_secret_token

# $ node ./bin/konga.js prepare --adapter postgres --uri postgresql://konga:konga123456@localhost:5432/konga
$ node ./bin/konga.js prepare --adapter postgres --uri postgresql://localhost:5432/konga
debug: Preparing database...
Using postgres DB Adapter.

$ npm run  production
> kongadmin@0.14.9 production /root/konga
> node --harmony app.js --prod

(node:13866) Warning: Accessing non-existent property 'padLevels' of module exports inside circular dependency
(Use `node --trace-warnings ...` to show where the warning was created)
The default `sails-disk` adapter is not designed for use as a production database;
(it stores the entire contents of your database in memory)
Instead, please use another adapter; e.g. sails-postgresql or sails-mongo.
For more info, see: http://sailsjs.org/documentation/concepts/deployment
To hide this warning message, enable `sails.config.orm.skipProductionWarnings`.
```

## 错误

&nbsp;&nbsp;&nbsp;&nbsp; 这个错误,官方和网上都说增加 `hookTimeout`能解决,我怎么都解决不了,这个错误在 `node prepare ...` 或者 `npm run production` 都可能会出现,因为kong是使用sails开发的.但是这个错误出现了机率还是很大的.我一共安装过三次,只有一次是非常顺利的.其它三次都是要么卡在 `node prepare ...` 阶段,要么卡在 `npm run production` 阶段,因此运气成分很大啊.但似乎我已经提了[issue] (https://github.com/pantsel/konga/issues/775)了.

```
Error: The hook `orm` is taking too long to load.
Make sure it is triggering its `initialize()` callback, or else set `sails.config.orm._hookTimeout to a higher value (currently 60000)
    at Timeout.tooLong [as _onTimeout] (/root/konga/node_modules/sails/lib/app/private/loadHooks.js:85:21)
    at listOnTimeout (internal/timers.js:557:17)
    at processTimers (internal/timers.js:500:7)
```

## 访问 konga

打开浏览器, http://x.x.x.x:1337/register ,x.x.x.x 是本机IP

<center>
<img src="../assets/images/manually-compile-open-source-gateway-konga-and-integrate-with-kong/figure-2.png" alt="Manually compile open source gateway konga and integrate with kong" title="Github of Anigkus" >
</center>

<center>
<img src="../assets/images/manually-compile-open-source-gateway-konga-and-integrate-with-kong/figure-3.png" alt="Manually compile open source gateway konga and integrate with kong" title="Github of Anigkus" >
</center>

<center>
<img src="../assets/images/manually-compile-open-source-gateway-konga-and-integrate-with-kong/figure-4.png" alt="Manually compile open source gateway konga and integrate with kong" title="Github of Anigkus" >
</center>




<br>

### [back](./)
