<iframe src="../detail-header.html" title="Github of Anigkus" style="height:0px,widht:0px;display:none" id="kusifreamheader"></iframe>

<h1 style="color:#606c71;text-align:center;" id="h1" >理解MongoDB 这些命令的区别和原理,理解MongoDB 这几个命令的原理和区别</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Understand the differences and principles of these commands in MongoDB</h1><br/>]:#

<center>
<img src="../assets/images/understand-the-differences-and-principles-of-these-commands-in-mongodb/figure-1.jpeg" alt="Understand the differences and principles of these commands in MongoDB" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; 一般我们安装 `MongoDB` 组件之后,系统都会存在一些命令.这些命令就是 `MongoDB` 的核心命令. `MongoDB` 大致命令都是从 `mongo` 和 `mongo-tools` 项目构建出来的.我们在使用 `MongoDB`一些命令必须要明白所使用命令主要提供的功能是做什么的,它们的原理又是怎样以及大致的执行逻辑是什么.并且还需要知道某些功能可能在不同的命令中都提供相似效果,那么它们之间的区别又是什么呢.话说我也使用 `MongoDB` 有一段时间了,一直也没在意怎么去内部分析过.那么本文我将详细说明下关于 `mongo, mongosh, mongod , mongos` 这几个命令的功能以及执行逻辑,并且还会指出各个命令的源码出处,并且最后还会用一张图整体说明下各个命令架构和执行流程. <br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; Some general notes on article.<br/>]:#
[> <br/>]:#

# mongo

&nbsp;&nbsp;&nbsp;&nbsp;它是一个命令客户端,使用 Javascript编写而封装的一个组件,通过底层网络接口与 MongoDB 服务器 (mongod) 进行交互.可以使用这个命令行工具查询、更新数据和执行一些管理操作等.

https://its401.com/article/pourtheworld/106175737
https://www.cnblogs.com/daizhj/archive/2011/03/17/1987311.html

&nbsp;&nbsp;&nbsp;&nbsp; 功能说明.

## 功能用法

&nbsp;&nbsp;&nbsp;&nbsp; 常用的一些命令.

[http://blog.codekissyoung.com/MongoDB数据库/mongo客户端命令]:#

```bash
# mongo # Connetion localhost,default test db.

# mongo 192.168.1.101:27017 # Connetion 192.168.1.101,default test db.

# mongo --nodb   # start mongo shell without connecting to any server 
# 连接操作
>  mongodb://[username:password@]host1[:port1],...[,hostN[:portN]]][/[database][?options]] 
> mongodb://admin:123456@localhost/  # 使用用户 admin 使用密码 123456 连接到本地的 MongoDB 服务
> mongodb://admin:123456@localhost/test # 顺便指定了数据库

mongodb:// 这是固定的格式，必须要指定

username:password@ 可选项，如果设置，在连接数据库服务器之后，驱动都会尝试登陆这个数据库

host1 必须的指定至少一个host, host1 是这个URI唯一要填写的它指定了要连接服务器的地址如果要连接复制集，请指定多个主机地址

portX 可选的指定端口，如果不填，默认为27017

/database 如果指定username:password@，连接并验证登陆指定数据库若不指定，默认打开 test 数据库

?options 是连接选项如果不使用/database，则前面需要加上/所有连接选项都是键值对name=value，键值对之间通过 & 或 ; 隔开

replicaSet=name 验证replica set的名称 Impliesconnect=replicaSet

slaveOk=true|false

true: 在connect=direct模式下，驱动会连接第一台机器，即使这台服务器不是主在connect=replicaSet模式下，驱动会发送所有的写请求到主并且把读取操作分布在其他从服务器

false: 在connect=direct模式下，驱动会自动找寻主服务器. 在connect=replicaSet模式下，驱动仅仅连接主服务器，并且所有的读写命令都连接到主服务器

safe=true|false

true: 在执行更新操作之后，驱动都会发送getLastError命令来确保更新成功(还要参考 wtimeoutMS).

false: 在每次更新之后，驱动不会发送getLastError来确保更新成功

w=n 驱动添加 { w : n } 到getLastError命令. 应用于safe=true

wtimeoutMS=ms驱动添加 { wtimeout : ms } 到 getlasterror 命令. 应用于 safe=true.

fsync=true|false

true: 驱动添加 { fsync : true } 到 getlasterror 命令.应用于 safe=true.

false: 驱动不会添加到getLastError命令中

journal=true|false如果设置为 true, 同步到 journal (在提交到数据库前写入到实体中). 应用于 safe=true

connectTimeoutMS=ms可以打开连接的时间

socketTimeoutMS=ms发送和接受sockets的时间


# 数据库操作
> show dbs;
admin   0.078GB
config  0.078GB
local   0.078GB
> use admin; # 选择切换数据库 / 创建数据库
> use runoob;
switched to db runoob
> db.runoob.insert({"name":"菜鸟教程"})
WriteResult({ "nInserted" : 1 })
> show dbs;
admin   0.078GB
config  0.078GB
local   0.078GB
runoob  0.078GB
> show tables;
logs
product
system.indexes

# 文档操作


# 索引操作
```

&nbsp;&nbsp;&nbsp;&nbsp; 更多参数详见`mongo --help`或者[官方文档](https://www.mongodb.com/docs/mongodb-shell/).

<details>
<summary>mongo --help</summary>

```bash
# mongo --help
MongoDB shell version v5.0.9
usage: mongo [options] [db address] [file names (ending in .js)]
db address can be:
  foo                   foo database on local machine
  192.168.0.5/foo       foo database on 192.168.0.5 machine
  192.168.0.5:9999/foo  foo database on 192.168.0.5 machine on port 9999
  mongodb://192.168.0.5:9999/foo  connection string URI can also be used
Options:
  --ipv6                               enable IPv6 support (disabled by 
                                       default)
  --host arg                           server to connect to
  --port arg                           port to connect to
  -h [ --help ]                        show this usage information
  --version                            show version information
  --verbose                            increase verbosity
  --shell                              run the shell after executing files
  --nodb                               don't connect to mongod on startup - no 
                                       'db address' arg expected
  --norc                               will not run the ".mongorc.js" file on 
                                       start up
  --quiet                              be less chatty
  --eval arg                           evaluate javascript
  --apiVersion arg                     set the MongoDB API version
  --apiStrict                          disable all features not included in the
                                       MongoDB Versioned API
  --apiDeprecationErrors               disable all features deprecated in the 
                                       MongoDB Versioned API
  --disableJavaScriptJIT               disable the Javascript Just In Time 
                                       compiler
  --enableJavaScriptJIT                enable the Javascript Just In Time 
                                       compiler
  --disableJavaScriptProtection        allow automatic JavaScript function 
                                       marshalling
  --retryWrites                        automatically retry write operations 
                                       upon transient network errors
  --disableImplicitSessions            do not automatically create and use 
                                       implicit sessions
  --jsHeapLimitMB arg                  set the js scope's heap size limit
  --idleSessionTimeout arg (=0)        Terminate the Shell session if it's been
                                       idle for this many seconds

Authentication Options:
  -u [ --username ] arg                username for authentication
  -p [ --password ] arg                password for authentication
  --authenticationDatabase arg         user source (defaults to dbname)
  --authenticationMechanism arg        authentication mechanism
  --gssapiServiceName arg (=mongodb)   Service name to use when authenticating 
                                       using GSSAPI/Kerberos
  --gssapiHostName arg                 Remote host name to use for purpose of 
                                       GSSAPI/Kerberos authentication

TLS Options:
  --tls                                use TLS for all connections
  --tlsCertificateKeyFile arg          PEM certificate/key file for TLS
  --tlsCertificateKeyFilePassword arg  Password for key in PEM file for TLS
  --tlsCAFile arg                      Certificate Authority file for TLS
  --tlsCRLFile arg                     Certificate Revocation List file for TLS
  --tlsAllowInvalidHostnames           Allow connections to servers with 
                                       non-matching hostnames
  --tlsAllowInvalidCertificates        Allow connections to servers with 
                                       invalid certificates
  --tlsFIPSMode                        Activate FIPS 140-2 mode at startup
  --tlsDisabledProtocols arg           Comma separated list of TLS protocols to
                                       disable [TLS1_0,TLS1_1,TLS1_2,TLS1_3]

AWS IAM Options:
  --awsIamSessionToken arg             AWS Session Token for temporary 
                                       credentials

FLE AWS Options:
  --awsAccessKeyId arg                 AWS Access Key for FLE Amazon KMS
  --awsSecretAccessKey arg             AWS Secret Key for FLE Amazon KMS
  --awsSessionToken arg                Optional AWS Session Token ID
  --keyVaultNamespace arg              database.collection to store encrypted 
                                       FLE parameters
  --kmsURL arg                         Test parameter to override the URL for 
                                       KMS

file names: a list of files to run. files have to end in .js and will exit after unless --shell is specified
```
</details>

## 执行逻辑


## 源码位置

&nbsp;&nbsp;&nbsp;&nbsp; [mongo.cpp](https://github.com/mongodb/mongo/blob/master/src/mongo/shell/mongo.cpp),以下是入口`mongo`入口函数代码片段.

```cpp
...

#ifdef _WIN32
int wmain(int argc, wchar_t* argvW[]) {
    mongo::quickExit(mongo::mongo_main(argc, mongo::WindowsCommandLine(argc, argvW).argv()));
}
#else   // #ifdef _WIN32
int main(int argc, char* argv[]) {
    mongo::quickExit(mongo::mongo_main(argc, argv));
}
#endif  // #ifdef _WIN32
```

# mongosh

&nbsp;&nbsp;&nbsp;&nbsp; 从 MongoDB v5.0 开始,  `mongo shell` 已经被官方宣布弃用.就是后续如果 `mongo shell` 对有新的功能添加或者维护都已经转向 `mongosh`. `mongosh` 是一个使用`TypeScript`重写的一个用于操作 `MongoDB` 的命令.如果你需要交互或更改数据时运行是,就可以使用 `mongosh` 命令.相对于使用 `mongo` 命令操作, `mongosh` 命令改进的语法突出显、改进的命令历史记录、改进的日志记录等.

&nbsp;&nbsp;&nbsp;&nbsp; 参考资料

[mongosh vs mongo](https://docs.mongodb.com/manual/reference/mongo/#std-label-compare-mongosh-mongo).

## 功能使用
&nbsp;&nbsp;&nbsp;&nbsp; 常用功能.

&nbsp;&nbsp;&nbsp;&nbsp; 更多使用方式,详见以下帮助文档或者[Mongosh GitHub](https://github.com/mongodb-js/mongosh)

```bash
# mongosh --help
```
## 执行逻辑

## 源码位置
[Mongosh GitHub](https://github.com/mongodb-js/mongosh).



# mongod

https://github.com/mongodb/mongo/blob/master/src/mongo/db/mongod_main.cpp
https://github.com/mongodb/mongo/tree/master/src/mongo/db/commands
https://mongoing.com/archives/78285

&nbsp;&nbsp;&nbsp;&nbsp; 功能说明.

## 功能使用

```bash
# mongod --help
Options:
  --networkMessageCompressors arg (=snappy,zstd,zlib)
                                        Comma-separated list of compressors to 
                                        use for network messages

General options:
  -h [ --help ]                         Show this usage information
  --version                             Show version information
  -f [ --config ] arg                   Configuration file specifying 
                                        additional options
  --configExpand arg                    Process expansion directives in config 
                                        file (none, exec, rest)
  --port arg                            Specify port number - 27017 by default
  --ipv6                                Enable IPv6 support (disabled by 
                                        default)
  --listenBacklog arg (=128)            Set socket listen backlog size
  --maxConns arg (=1000000)             Max number of simultaneous connections
  --pidfilepath arg                     Full path to pidfile (if not set, no 
                                        pidfile is created)
  --timeZoneInfo arg                    Full path to time zone info directory, 
                                        e.g. /usr/share/zoneinfo
  --nounixsocket                        Disable listening on unix sockets
  --unixSocketPrefix arg                Alternative directory for UNIX domain 
                                        sockets (defaults to /tmp)
  --filePermissions arg                 Permissions to set on UNIX domain 
                                        socket file - 0700 by default
  --fork                                Fork server process
  -v [ --verbose ] [=arg(=v)]           Be more verbose (include multiple times
                                        for more verbosity e.g. -vvvvv)
  --quiet                               Quieter output
  --logpath arg                         Log file to send write to instead of 
                                        stdout - has to be a file, not 
                                        directory
  --syslog                              Log to system's syslog facility instead
                                        of file or stdout
  --syslogFacility arg                  syslog facility used for mongodb syslog
                                        message
  --logappend                           Append to logpath instead of 
                                        over-writing
  --logRotate arg                       Set the log rotation behavior 
                                        (rename|reopen)
  --timeStampFormat arg                 Desired format for timestamps in log 
                                        messages. One of iso8601-utc or 
                                        iso8601-local
  --setParameter arg                    Set a configurable parameter
  --bind_ip arg                         Comma separated list of ip addresses to
                                        listen on - localhost by default
  --bind_ip_all                         Bind to all ip addresses
  --noauth                              Run without security
  --transitionToAuth                    For rolling access control upgrade. 
                                        Attempt to authenticate over outgoing 
                                        connections and proceed regardless of 
                                        success. Accept incoming connections 
                                        with or without authentication.
  --slowms arg (=100)                   Value of slow for profile and console 
                                        log
  --slowOpSampleRate arg (=1)           Fraction of slow ops to include in the 
                                        profile and console log
  --profileFilter arg                   Query predicate to control which 
                                        operations are logged and profiled
  --auth                                Run with security
  --clusterIpSourceAllowlist arg        Network CIDR specification of permitted
                                        origin for `__system` access
  --profile arg                         0=off 1=slow, 2=all
  --cpu                                 Periodically show cpu and iowait 
                                        utilization
  --sysinfo                             Print some diagnostic system 
                                        information
  --noscripting                         Disable scripting engine
  --notablescan                         Do not allow table scans
  --shutdown                            Kill a running server (for init 
                                        scripts)
  --keyFile arg                         Private key for cluster authentication
  --clusterAuthMode arg                 Authentication mode used for cluster 
                                        authentication. Alternatives are 
                                        (keyFile|sendKeyFile|sendX509|x509)

Replication options:
  --oplogSize arg                       Size to use (in MB) for replication op 
                                        log. default is 5% of disk space (i.e. 
                                        large is good)

Replica set options:
  --replSet arg                         arg is <setname>[/<optionalseedhostlist
                                        >]
  --enableMajorityReadConcern [=arg(=1)] (=1)
                                        Enables majority readConcern. 
                                        enableMajorityReadConcern=false is no 
                                        longer supported

Sharding options:
  --configsvr                           Declare this is a config db of a 
                                        cluster; default port 27019; default 
                                        dir /data/configdb
  --shardsvr                            Declare this is a shard db of a 
                                        cluster; default port 27018

Storage options:
  --storageEngine arg                   What storage engine to use - defaults 
                                        to wiredTiger if no data files present
  --dbpath arg                          Directory for datafiles - defaults to 
                                        /data/db
  --directoryperdb                      Each database will be stored in a 
                                        separate directory
  --syncdelay arg (=60)                 Seconds between disk syncs
  --journalCommitInterval arg (=100)    how often to group/batch commit (ms)
  --upgrade                             Upgrade db if needed
  --repair                              Run repair on all dbs
  --journal                             Enable journaling
  --nojournal                           Disable journaling (journaling is on by
                                        default for 64 bit)
  --oplogMinRetentionHours arg (=0)     Minimum number of hours to preserve in 
                                        the oplog. Default is 0 (turned off). 
                                        Fractions are allowed (e.g. 1.5 hours)

TLS Options:
  --tlsOnNormalPorts                    Use TLS on configured ports
  --tlsMode arg                         Set the TLS operation mode 
                                        (disabled|allowTLS|preferTLS|requireTLS
                                        )
  --tlsCertificateKeyFile arg           Certificate and key file for TLS
  --tlsCertificateKeyFilePassword arg   Password to unlock key in the TLS 
                                        certificate key file
  --tlsClusterFile arg                  Key file for internal TLS 
                                        authentication
  --tlsClusterPassword arg              Internal authentication key file 
                                        password
  --tlsCAFile arg                       Certificate Authority file for TLS
  --tlsClusterCAFile arg                CA used for verifying remotes during 
                                        inbound connections
  --tlsCRLFile arg                      Certificate Revocation List file for 
                                        TLS
  --tlsDisabledProtocols arg            Comma separated list of TLS protocols 
                                        to disable [TLS1_0,TLS1_1,TLS1_2,TLS1_3
                                        ]
  --tlsAllowConnectionsWithoutCertificates 
                                        Allow client to connect without 
                                        presenting a certificate
  --tlsAllowInvalidHostnames            Allow server certificates to provide 
                                        non-matching hostnames
  --tlsAllowInvalidCertificates         Allow connections to servers with 
                                        invalid certificates
  --tlsFIPSMode                         Activate FIPS 140-2 mode at startup
  --tlsLogVersions arg                  Comma separated list of TLS protocols 
                                        to log on connect [TLS1_0,TLS1_1,TLS1_2
                                        ,TLS1_3]

AWS IAM Options:
  --awsIamSessionToken arg              AWS Session Token for temporary 
                                        credentials

WiredTiger options:
  --wiredTigerCacheSizeGB arg           Maximum amount of memory to allocate 
                                        for cache; Defaults to 1/2 of physical 
                                        RAM
  --zstdDefaultCompressionLevel arg (=6)
                                        Default compression level for zstandard
                                        compressor
  --wiredTigerJournalCompressor arg (=snappy)
                                        Use a compressor for log records 
                                        [none|snappy|zlib|zstd]
  --wiredTigerDirectoryForIndexes       Put indexes and data in different 
                                        directories
  --wiredTigerCollectionBlockCompressor arg (=snappy)
                                        Block compression algorithm for 
                                        collection data [none|snappy|zlib|zstd]
  --wiredTigerIndexPrefixCompression arg (=1)
                                        Use prefix compression on row-store 
                                        leaf pages

Free Monitoring Options:
  --enableFreeMonitoring arg            Enable Cloud Free Monitoring 
                                        (on|runtime|off)
  --freeMonitoringTag arg               Cloud Free Monitoring Tags

```

## 执行逻辑

## 源码位置

# mongos

https://github.com/mongodb/mongo/blob/master/src/mongo/s/mongos_main.cpp
https://github.com/mongodb/mongo/tree/master/src/mongo/s/commands
https://mongoing.com/archives/75581

&nbsp;&nbsp;&nbsp;&nbsp; 功能说明.

## 功能使用

## 执行逻辑

## 源码位置

# 整体逻辑


<br>

### [back](./)
