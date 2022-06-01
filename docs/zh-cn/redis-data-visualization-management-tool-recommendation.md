<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >Redis 数据可视化管理工具推荐</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Redis data visualization management tool recommendation</h1><br/>]:#

<center>
<img src="../assets/images/redis-data-visualization-management-tool-recommendation/figure-1.jpeg" alt="Redis data visualization management tool recommendation" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; 作为程序员来说, Redis 对我们来说并不会陌生,几乎都会与它打过交道. Redis 是一款 NoSQL 类型的键值对内存型数据库,在大数据量高并发下的场景下尤为重要,但是大多数的时候还是作为缓存使用,是为了减少后端关系型数据库IO压力. Redis 本身查询数据展现方式还是不够直观,安装包中只提供了一个 `redis-cli`命令行的工具.其实社区有不少图形化的管理工具,某些时候使用图形化确实比较惬意(只想鼠标点点的时候,😄).工具基本也分是管理工具还是监控工具,我在下面也都会说到.那么我这里推荐几款开源的图形化工具,非常实用,大家有空可以去试试.<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; Some general notes on article.<br/>]:#
[> <br/>]:#

[![Release](https://img.shields.io/github/v/release/RedisInsight/RedisInsight.svg?sort=semver)](https://github.com/RedisInsight/RedisInsight/releases)
[![CircleCI](https://circleci.com/gh/RedisInsight/RedisInsight/tree/main.svg?style=svg)](https://circleci.com/gh/RedisInsight/RedisInsight/tree/main)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/RedisInsight/RedisInsight.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/RedisInsight/RedisInsight/alerts/)




RedisInsight

Redis Desktop Manager

Redis Commander

p3x-redis-ui

Redis Browser

QRedis

Redis Explorer

RedisStudio

AnotherRedisDesktopManager

QuickRedis


# 测试数据
 

```
127.0.0.1:6379> keys *
(empty array)
127.0.0.1:6379> info server
# Server
redis_version:6.2.7
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:f300021a674a92d6
redis_mode:standalone
os:Linux 4.18.0-277.el8.x86_64 x86_64
arch_bits:64
monotonic_clock:POSIX clock_gettime
multiplexing_api:epoll
atomicvar_api:c11-builtin
gcc_version:8.5.0
process_id:568483
process_supervised:no
run_id:920f24df4f0321d047e6d1aa2ca7900bb16109f5
tcp_port:6379
server_time_usec:1654074930873401
uptime_in_seconds:181028
uptime_in_days:2
hz:10
configured_hz:10
lru_clock:9907762
executable:/usr/local/redis-627/bin/redis-server
config_file:
io_threads_active:0

# Sting
127.0.0.1:6379> SET string:user:info '{"name":"Andi","gender":0}'
OK
127.0.0.1:6379> GET  string:user:info
"{\"name\":\"Andi\",\"gender\":0}"

# List
127.0.0.1:6379> LPUSH list:user:companys "facebook"
(integer) 1
127.0.0.1:6379> LPUSH list:user:companys "twitter"
(integer) 2
127.0.0.1:6379> LPUSH list:user:companys "microsoft"
(integer) 3
127.0.0.1:6379> LPUSH list:user:companys "facebook"
(integer) 4
127.0.0.1:6379> 
127.0.0.1:6379> LRANGE list:user:companys 0 -1
1) "facebook"
2) "microsoft"
3) "twitter"
4) "facebook"

# Hash
127.0.0.1:6379> HMSET hash:user:info username Andi company Facebook gender 0
OK
127.0.0.1:6379> HGETALL hash:user:info
1) "username"
2) "Andi"
3) "company"
4) "Facebook"
5) "gender"
6) "0"

# Set
127.0.0.1:6379> SADD set:user:companys "facebook"
(integer) 1
127.0.0.1:6379> SADD set:user:companys "twitter"
(integer) 1
127.0.0.1:6379> SADD set:user:companys "microsoft"
(integer) 1
127.0.0.1:6379> SADD set:user:companys "facebook"
(integer) 0
127.0.0.1:6379> SMEMBERS set:user:companys
1) "microsoft"
2) "facebook"
3) "twitter"

# Sorted Set
127.0.0.1:6379> ZADD  sortedset:user:companys 10 facebook
(integer) 1
127.0.0.1:6379> ZADD  sortedset:user:companys 60 twitter
(integer) 1
127.0.0.1:6379> ZADD  sortedset:user:companys 40 microsoft
(integer) 1
127.0.0.1:6379> ZRANGE sortedset:user:companys 0 -1
1) "facebook"
2) "microsoft"
3) "twitter"

# Bitmaps
127.0.0.1:6379> setbit bitmaps:user:info 1001 1
(integer) 0
127.0.0.1:6379> setbit bitmaps:user:info 1002 0
(integer) 0
127.0.0.1:6379> setbit bitmaps:user:info 1003 1
(integer) 0
127.0.0.1:6379> getbit  bitmaps:user:info 1001
(integer) 1
127.0.0.1:6379> getbit  bitmaps:user:info 1002
(integer) 0
127.0.0.1:6379> getbit  bitmaps:user:info 1003
(integer) 1
127.0.0.1:6379> bitcount bitmaps:user:info
(integer) 2

```


# RedisInsight

## 软件说明

## 安装方法

## 软件截图

## 评测心得


<br/>

<mark>推荐指数</mark> : 🌟🌟🌟🌟🌟</br>

<mark>是否开源</mark> : 是/否</br>

<mark>仓库地址</mark> : http://github.com/anigkus</br>

<mark>下载地址</mark> : http://github.com/anigkus<</br>

<mark>是否收费</mark> : 是/否</br>

<mark>开发语言</mark> : Java+Golang</br>

<mark>用户群体</mark> : 美国、中国</br>

<mark>支持平台</mark> : Mac、Windows、Linux</br>








<br>

### [back](./)
