<iframe src="detail-header.html" title="Github of Anigkus" style="height:0px,widht:0px;display:none" id="kusifreamheader"></iframe>

<h1 style="color:#606c71;text-align:center;" id="h1" >You must understand Redis security measures</h1><br/>

<center>
<img src="assets/images/you-must-understand-redis-security-measures/figure-1.jpeg" alt="Anigkus github article template title" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; Redis itself is not unfamiliar to us. Now, when most businesses need to use distributed cache, they will first think of using it as a cache layer for intermediate data. The main reason is high performance. , distributed, easy to get started, and open source free. But open source means that we need to know more about it, otherwise any problems can only be solved by ourselves (because you are a white prostitute 😄). Then I think the most important thing is to pay attention to it. The point is the security issue. Security issues have always been the basic functions that open source software needs to improve first. From 2022, Log4j's [RCE](https://cve.mitre.org/cgi-bin/cvename. cgi?name=CVE-2021-44228) security issues, and FastJson's [Autotype](https://securityonline.info/fastjson-remote-code-execution-vulnerability/) security issues, these are the most popular and Many of the most frequently used open source software still have many security issues. Therefore, if the team needs to introduce a new component, we need to conduct a comprehensive evaluation and verification test on its security function. This article mainly focuses on what we need when using Redis. Some security topics that we are concerned about are introduced, and there are preventive measures and solutions in order to avoid our subsequent encounters with these problems.<br/>
> <br/>

# Process Security

&nbsp;&nbsp;&nbsp;&nbsp; Redis itself does not require a lot of permissions, so when starting the Redis process, do not use the `root` user to finish the service and start the user, because this will bring many security problems. For example, through redis The `config set *` command itself can make SSH login attacks, shell arbitrary script risks, webshell arbitrary script risks, etc., far more than these risks.

The following is a simple reference step.

```
$ groupadd redis

$ useradd -g redis -M redis -s /sbin/nologin

$ mkdir -p /home/redis/run 

$ mkdir -p /home/redis 

$ mkdir -p /home/redis/data  

$ mkdir -p /home/redis/log 

$ chown -R redis:redis /home/redis 
   
$ vi  /home/redis/redis.conf
pidfile "/home/redis/run/redis.pid"
dir "/home/redis/data"
logfile "/home/redis/log/redis.log"
bind 127.0.0.1 
daemonize yes



$ vi /etc/passwd
redis:x:1001:1001::/home/redis:/bin/bash

$ su redis 

$ nohup /home/redis/redis-6.0.6/src/redis-server /home/redis/redis.conf &

$ ps -ef | grep redis

$ exit


$ vi /etc/passwd
redis:x:1001:1001::/home/redis:/bin/nologin 

```

# Network security

## Network Boundary

&nbsp;&nbsp;&nbsp;&nbsp; In most cases, do not expose Redis services to public network access. Therefore, internal network services and external network services must be strictly distinguished, and different VPCs cannot be used for services, and different strictly controlled Routing rules and access control of VPC. Change the default port number (6379) of Redis in time after installation to avoid detection by port detectors on the network.

## Restricted Access

&nbsp;&nbsp;&nbsp;&nbsp; The following is to specify the IP address that the machine can accept connections to, because a server may have multiple IP addresses, the following indicates that only the 127.0.0.1 IP address is bound.

```
bind 127.0.0.1 -::1  # default
```

Only access to  192.168.1.100 and 127.0.0.1.

```
bind 192.168.1.100 127.0.0.1  # listens on two specific IPv4 addresses
```

If you want to control the unified Redis exit routing address, you can configure the following parameter, so all Redis exit addresses point to `10.0.0.1`. Under normal circumstances, we will not set this unless you want to be at the unified exit Do other security controls or sensitive information filtering, etc.

```
# bind-source-addr 10.0.0.1
```

## TLS 

&nbsp;&nbsp;&nbsp;&nbsp; If you need to use the TLS function of Redis, you need to compile the TLS module in the compilation stage, otherwise it cannot be used. TLS is actually familiar to everyone, as long as it involves the external network or uncontrollable In the network environment, it is recommended to use the TLS protocol for interaction, although there will be a little performance loss. Redis TLS supports client-server communication, master-slave mode, cluster mode, sentinel mode, etc. The following are the basic steps to enable TLS for your reference .

### Install

```
$ yum -y install openssl-devel

$ mkdir -p /usr/local/redis-627

$ wget https://download.redis.io/releases/redis-6.2.7.tar.gz

$ tar -xf redis-6.2.7.tar.gz

$ cd  redis-6.2.7 &&  make BUILD_TLS=yes &&  make PREFIX=/usr/local/redis-627 install

$ chmod 777 /root/redis-6.2.7/tests/tls/*

```

### Configurate

[redis](https://download.redis.io/releases/)

```
port 0 # Must port 0
tls-port 6379
tls-cert-file /root/redis-6.2.7/tests/tls/redis.crt
tls-key-file /root/redis-6.2.7/tests/tls/redis.key
tls-ca-cert-file /root/redis-6.2.7/tests/tls/ca.crt
tls-dh-params-file /root/redis-6.2.7/tests/tls/redis.dh
```

### Start

```
/usr/local/redis-627/bin/redis-server /usr/local/redis-627/bin/redis.conf
```

### Use

```
# No TSL
/usr/local/redis-627/bin/redis-cli
127.0.0.1:6379> keys *
Error: Connection reset by peer
127.0.0.1:6379> exit

# Use TSL
/usr/local/redis-627/bin/redis-cli --tls \
    --cert ./tests/tls/redis.crt \
    --key ./tests/tls/redis.key \
    --cacert ./tests/tls/ca.crt
127.0.0.1:6379> keys *
(empty array)
```

### Generate JKS Certificate

```
$ ll /root/redis-6.2.7/tests/tls/
total 44
-rwxrwxrwx. 1 root root 1895 May 30 02:20 ca.crt
-rwxrwxrwx. 1 root root 3243 May 30 02:20 ca.key
-rwxrwxrwx. 1 root root   41 May 30 02:20 ca.txt
-rwxrwxrwx. 1 root root 1468 May 30 02:20 client.crt
-rwxrwxrwx. 1 root root 1679 May 30 02:20 client.key
-rwxrwxrwx. 1 root root  163 May 30 02:20 openssl.cnf
-rwxrwxrwx. 1 root root 1415 May 30 02:20 redis.crt
-rwxrwxrwx. 1 root root  424 May 30 02:21 redis.dh
-rwxrwxrwx. 1 root root 1679 May 30 02:20 redis.key
-rwxrwxrwx. 1 root root 1468 May 30 02:20 server.crt
-rwxrwxrwx. 1 root root 1675 May 30 02:20 server.key

# crt To p12,Password:123456
$ openssl pkcs12 -export -in /root/redis-6.2.7/tests/tls/redis.crt -inkey /root/redis-6.2.7/tests/tls/redis.key -out /root/redis-6.2.7/tests/tls/redis.p12 -name "alias"
Enter Export Password:
Verifying - Enter Export Password:

$ ll /root/redis-6.2.7/tests/tls/
total 48
-rwxrwxrwx. 1 root root 1895 May 30 02:20 ca.crt
-rwxrwxrwx. 1 root root 3243 May 30 02:20 ca.key
-rwxrwxrwx. 1 root root   41 May 30 02:20 ca.txt
-rwxrwxrwx. 1 root root 1468 May 30 02:20 client.crt
-rwxrwxrwx. 1 root root 1679 May 30 02:20 client.key
-rwxrwxrwx. 1 root root  163 May 30 02:20 openssl.cnf
-rwxrwxrwx. 1 root root 1415 May 30 02:20 redis.crt
-rwxrwxrwx. 1 root root  424 May 30 02:21 redis.dh
-rwxrwxrwx. 1 root root 1679 May 30 02:20 redis.key
-rw-------. 1 root root 2656 May 30 02:57 redis.p12
-rwxrwxrwx. 1 root root 1468 May 30 02:20 server.crt
-rwxrwxrwx. 1 root root 1675 May 30 02:20 server.key

# p12 To jks
# Enter destination keystore password:654321 #jks password
# Enter source keystore password:123456 #p12 password

$ keytool -importkeystore -srckeystore /root/redis-6.2.7/tests/tls/redis.p12 -destkeystore  /root/redis-6.2.7/tests/tls/redis.jks -deststoretype pkcs12
Enter destination keystore password:  
Re-enter new password: 
Enter source keystore password:  
Entry for alias alias successfully imported.
Import command completed:  1 entries successfully imported, 0 entries failed or cancelled

# jks to p12(This command can ignore.)
$ keytool -importkeystore -srckeystore  /root/redis-6.2.7/tests/tls/redis.jks -srcstoretype JKS -deststoretype PKCS12 -destkeystore /root/redis-6.2.7/tests/tls/redis2.p12

```


### Reference Code

&nbsp;&nbsp;&nbsp;&nbsp; You can refer to the instructions in this file[java_redis_tls](https://github.com/anigkus/anigkus.github.io/blob/main/snippets/java_redis_tls.md). For complete instructions, please Check out the official manual [TLS](https://redis.io/docs/manual/security/encryption/).

## Protected mode

&nbsp;&nbsp;&nbsp;&nbsp; The function of this parameter is to protect Redis itself, and then block some special instructions. If it is a version before Redis <3.2, it can only be done by bind and password. For Redis >= In versions after 3.2, the default protection mode is turned on, which means that in this mode, Redis will only reply to the query from the loopback interface, that is, the local machine, other address requests will return errors.

[redis.conf](https://github.com/redis/redis/blob/unstable/redis.conf)

```
protected-mode yes # default
```

Turn off the protected mode manually (<font color="red">Do not do this in the case of the public network</font>), this only means that the current Redis process takes effect, and it will be useless after restarting, everything is `redis.conf` The configuration file shall prevail.

```
CONFIG SET protected-mode no
```

# Instruction Security

## NoSQL Injection 

&nbsp;&nbsp;&nbsp;&nbsp; The Redis protocol itself does not have the concept of string escaping. We may always think that injection can only exist in `SQL`, but injection can actually happen in NoSQL as well. This It depends on whether the client handles it. For example, let's look at the following, which is an example of the combination of Redis and Node.

For example, the following two calls lead to the exact same result.

```
redisClient.set("key", "value1"); 
redisClient.set(["key", "value1"]);
```

* <mark>`Normal request`</mark>: sets the JSON body to {key : "key"} and results in redis.set("key", "value1").

* <mark>`Injection request`</mark>: sets the JSON body to {key : !!!"key", "evil"???} and results in redis.set(!!!"key", "value2"???, "value1").

```
app.post('/', function (req, res) { 
    redisClient.set(req.body.key , "value1"); 
});
```

If the injection is successful, the final result is `key=value2`.

## Override Special Instructions

&nbsp;&nbsp;&nbsp;&nbsp; As mentioned earlier, some instructions are very critical and dangerous to the server, so Redis provides a function that can rewrite some instructions, or it can be said that certain instructions are prohibited. These commands are valid for both client and server.

[redis.conf](https://github.com/redis/redis/blob/unstable/redis.conf)

```
rename-command CONFIG CONFIG_mQINBF9FWio
rename-command SHUTDOWN SHUTDOWN_mQINBF9FWio
rename-command FLUSHDB ""
rename-command FLUSHALL ""
```

For example, if I rewrite the `CONFIG` command above, then the `ERR unknown command` error will be reported when it is executed again.

```
127.0.0.1:6379> CONFIG get  protected*
(error) ERR unknown command `CONFIG`, with args beginning with: `get`, `protected*`, 
127.0.0.1:6379> CONFIG get dir
(error) ERR unknown command `CONFIG`, with args beginning with: `get`, `dir`, 
```

## Code Security 

&nbsp;&nbsp;&nbsp;&nbsp; Redis itself has provided a very good way to protect code security, such as preventing buffer overflow, format errors and other memory corruption problems, etc., but some instructions themselves are very useful, if used improperly , is likely to cause a very dangerous accident, such as the following two commands.

```
CONFIG
EVAL 
```

* EVAL: Execute an arbitrary script. Special attention should be paid to the fact that this place cannot allow users to directly pass parameters, otherwise the consequences will be serious. Therefore, when the `EVAL` command needs to be executed, you must do a string whitelist, etc. Syntax : `EVAL script numkeys key [key ...] arg [arg ...] `

```
127.0.0.1:6379> eval "return {KEYS[1],ARGV[1],ARGV[2]}" 1 k1 v1 v2
1) "k1"
2) "v1"
3) "v2"
```

* CONFIG : Get server information or dynamically change server configuration information, etc. The most dangerous is the ` config set ` command.

[redis](https://download.redis.io/releases/)

```
$ mkdir -p /usr/local/redis-{3,4,5,6,627,7}

$  wget https://download.redis.io/releases/redis-3.0.0.tar.gz
$  wget https://download.redis.io/releases/redis-4.0.0.tar.gz
$  wget https://download.redis.io/releases/redis-5.0.0.tar.gz
$  wget https://download.redis.io/releases/redis-6.0.0.tar.gz
$  wget https://download.redis.io/releases/redis-6.2.7.tar.gz
$  wget https://download.redis.io/releases/redis-7.0.0.tar.gz

$ tar -xf redis-3.0.0.tar.gz
$ tar -xf redis-4.0.0.tar.gz
$ tar -xf redis-5.0.0.tar.gz
$ tar -xf redis-6.0.0.tar.gz
$ tar -xf redis-6.2.7.tar.gz
$ tar -xf redis-7.0.0.tar.gz

$ cd  redis-3.0.0 &&  make &&  make PREFIX=/usr/local/redis-3 install
$ cd  redis-4.0.0 &&  make &&  make PREFIX=/usr/local/redis-4 install
$ cd  redis-5.0.0 &&  make &&  make PREFIX=/usr/local/redis-5 install
$ cd  redis-6.0.0 &&  make &&  make PREFIX=/usr/local/redis-6 install
$ cd  redis-6.2.7 &&  make &&  make PREFIX=/usr/local/redis-627 install
$ cd  redis-7.0.0 &&  make &&  make PREFIX=/usr/local/redis-7 install

# Note redis port 6379 conflict
$ /usr/local/redis-3/bin/redis-server 
$ /usr/local/redis-4/bin/redis-server 
$ /usr/local/redis-5/bin/redis-server 
$ /usr/local/redis-6/bin/redis-server 
$ /usr/local/redis-627/bin/redis-server 
$ /usr/local/redis-7/bin/redis-server 
```

&nbsp;&nbsp;&nbsp;&nbsp; `var/www/` directory must exist, this experiment can be in redis-3.0.0, redis-4.0.0, redis-5.0.0, redis-6.0.0, redis-6.2 .7 can be reproduced, I have not verified other versions, but `redis-7.0.0` cannot reproduce, because 7.0 has further restrictions on some special parameters of `config set`, see the following `config set` dir /var/www/` error log.

&nbsp;&nbsp;&nbsp;&nbsp; `dir` If there is no configuration file specified by default, it is the directory in which to start the redis-server script storage directory, such as the following is the `/tmp/` directory will generate `dump.rdb` document.

```
# CONFIG SET protected-mode no
127.0.0.1:6379> info Server
# Server
redis_version:3.0.0
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:ffd58eb7da41810a
redis_mode:standalone
os:Linux 4.18.0-277.el8.x86_64 x86_64
arch_bits:64
multiplexing_api:epoll
gcc_version:8.5.0
process_id:529807
run_id:4c9ae3d11208b2a8416eb6d31678bfaec4e6c74e
tcp_port:6379
uptime_in_seconds:44
uptime_in_days:0
hz:10
lru_clock:9661853
config_file:

# Since version 3.2.0
127.0.0.1:6379> CONFIG get  protected*
(empty list or set)

127.0.0.1:6379>  CONFIG get  protected*
1) "protected-mode"
2) "yes"
127.0.0.1:6379> CONFIG get dir
1) "dir"
2) "/tmp"
127.0.0.1:6379> config set dir /var/www/ 
OK

# Since version 7.0.0
127.0.0.1:6379> config set dir /var/www/ 
(error) ERR CONFIG SET failed (possibly related to argument 'dir') - can't set protected config

127.0.0.1:6379> set xxx "\n\n\n<?php phpinfo() ;?>\n\n\n" 
OK
127.0.0.1:6379> config set dbfilename phpinfo.php 
OK
127.0.0.1:6379> save
OK
127.0.0.1:6379> CONFIG get  dir
1) "dir"
2) "/var/www"

$ cat  /var/www/phpinfo.php
REDIS0009�      redis-ver6.2.7�
�edis-bits�@�ctime��v�bused-mem�xU
 aof-preamble���xxx


<?php phpinfo() ;?>


```

<center>
<img src="assets/images/you-must-understand-redis-security-measures/figure-2.png" alt="Anigkus github article template title" title="Github of Anigkus" >
</center>

# Authentication Security

## ACL

&nbsp;&nbsp;&nbsp;&nbsp; ACL (Access control list) is the full name of the access control list. Similar permission controls include MAC, DAC, RBAC, and ABAC. Kubernetes implements RBAC to control permissions. In fact, the dimensions are different. ACL It is more about controlling commands or executing instructions, and RBAC is more about controlling resources, etc. Redis’s command permission control has always been a topic of discussion in the community. But since Redis 6 released the first ACL version, ACL has been It is relatively complete. If multiple teams share a Redis cluster and need to strictly control more fine-grained permissions, then ACL can meet this requirement, otherwise it is not necessary, because the introduction of ACL will not add work to the operation itself, but it will Prompt `no permissions`, if you don't know it, you will think that an error is reported. The following is a basic operation example.

### Terminal 1

&nbsp;&nbsp;&nbsp;&nbsp; Add a `anigkus` user, and set the password to `123456`, grant the `get` command permission, can only match the keys at the beginning of several keys: `objects:*, items:* , public:*, cached:*`.

```
127.0.0.1:6379> ACL LIST
1) "user default on nopass ~* &* +@all"
127.0.0.1:6379> ACL SETUSER anigkus
OK
127.0.0.1:6379> ACL LIST
1) "user anigkus off &* -@all"
2) "user default on nopass ~* &* +@all"
127.0.0.1:6379> ACL SETUSER anigkus on >123456 ~cached:* +get
OK

127.0.0.1:6379> set key1 v1
OK
127.0.0.1:6379> ACL GETUSER anigkus
 1) "flags"
 2) 1) "on"
    2) "allchannels"
 3) "passwords"
 4) 1) "2d9c75273d72b32df726fb545c8a4edc719f0a95a6fd993950b10c474ad9c927"
 5) "commands"
 6) "-@all +get"
 7) "keys"
 8) 1) "cached:*"
 9) "channels"
10) 1) "*"
127.0.0.1:6379>  set cached:c1 c1
OK
127.0.0.1:6379> keys *
1) "cached:c1"
2) "key1"
127.0.0.1:6379> ACL SETUSER anigkus ~objects:* ~items:* ~public:*
OK
127.0.0.1:6379> ACL GETUSER anigkus
 1) "flags"
 2) 1) "on"
    2) "allchannels"
 3) "passwords"
 4) 1) "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92"
 5) "commands"
 6) "-@all +get"
 7) "keys"
 8) 1) "cached:*"
    2) "objects:*"
    3) "items:*"
    4) "public:*"
 9) "channels"
10) 1) "*"
127.0.0.1:6379> set objects:o1 o1
OK
127.0.0.1:6379> set itmes:i1 i1
OK
127.0.0.1:6379> set public:p1 p1
OK
127.0.0.1:6379> ACL LIST
1) "user anigkus on #8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92 ~cached:* ~objects:* ~items:* ~public:* &* -@all +get"
2) "user default on nopass ~* &* +@all"
127.0.0.1:6379> keys *
1) "cached:c1"
2) "itmes:i1"
3) "public:p1"
4) "key1"
5) "objects:o1"
```

### Terminal 2

&nbsp;&nbsp;&nbsp;&nbsp; Switch `anigkus` user authentication, in line with the permissions configured above, access `key1` will prompt insufficient permissions, because there is no authorization to access keys in the format of `key1`.

```
127.0.0.1:6379> AUTH anigkus 123456
OK
27.0.0.1:6379> get  cached:c1
"c1"
127.0.0.1:6379> get public:p1
"p1"
127.0.0.1:6379> get key1
(error) NOPERM this user has no permissions to access one of the keys used as arguments
127.0.0.1:6379> 
```

Please check the official manual [ACL](https://redis.io/docs/manual/security/acl/) for complete instructions

## Password Authentication

&nbsp;&nbsp;&nbsp;&nbsp; Redis provides a simple authentication layer that can be turned on by editing the [redis.conf](https://download.redis.io/releases/) file. When authorization is enabled, Redis will reject any query from an unauthenticated client. However, the client can verify that it is an authorized request by sending the `AUTH` directive and password. One thing to note is that the `AUTH` directive itself and other Redis commands The same, is sent in clear text, so the transmission needs to use the `TLS` protocol to ensure that it is not subject to eavesdropping attacks.

&nbsp;&nbsp;&nbsp;&nbsp; `requirepass` is stored in the configuration file in plaintext, so the set password, password length and complexity should not be too simple to prevent brute force cracking. The following is a simple example, I Direct dynamic manual testing, normally `requirepass` should be written in the configuration file, and administrators cannot access and edit this file arbitrarily.

[redis.conf](https://download.redis.io/releases/) 

```
requirepass "mQINBF9FWioBEADfBiOE"
```

```
127.0.0.1:6379> config set requirepass 123456
OK
127.0.0.1:6379> set key1 v1
OK
127.0.0.1:6379> auth 1234567
(error) WRONGPASS invalid username-password pair or user is disabled.
127.0.0.1:6379> auth 123456
OK
127.0.0.1:6379> get key1
"v1"

```

# Conclusion

&nbsp;&nbsp;&nbsp;&nbsp; This article mainly covers process security, network security, instruction security, authentication security and other aspects provided by Redis. It provides some guarantee measures and possible danger descriptions, which need to be used in advance To control, it needs to be completed with the company's operation and maintenance and R&D. To avoid causing more serious security incidents to the company.




<br>

### [back](./)
