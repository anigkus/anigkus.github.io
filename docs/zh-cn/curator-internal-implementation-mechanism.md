<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >Curator内部实现机制</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Curator internal implementation mechanism</h1><br/>]:#

<center>
<img src="../assets/images/curator-internal-implementation-mechanism/figure-1.jpeg" alt="Curator internal implementation mechanism" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; 文章简要说明.<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; Some general notes on article.<br/>]:#
[> <br/>]:#

# 如何正确的使用开源

[# What's the right posture to use open source?]:#

<center>
<img src="../assets/images/curator-internal-implementation-mechanism/figure-2.png" alt="Curator internal implementation mechanism" title="Github of Anigkus" >
</center>

* All-takenism: 称之“拿来主义”阶段,不管它是怎么样的,先用着再说.

* Know the roots: 到后来用着用着时间九了,也就越来约熟悉了,对些原理和问题也了解的差不多了,称之“知根知底”阶段.

* Patch up: 再到后来用的过程中,总会发现些问题或者某些功能不如意.由于我们知道了它怎么弄的,就会尝试着去更改它的问题或者增加功能,这个阶段称之为“修修补补”.

* Make a new start: 等到完全吃透了,发现简单的修改或者增加功能根本没法改变底层越来不足的地方.这个时候发现投入的补丁时间比越来越多,而换来的结果又有点不理想.想着想着还不如重新’造一个‘,反正已经完全弄懂了它的原理,并且还可以为未来增加新功能或者维护带来诸多方便,那就造一个呗,这个阶段称之为“推到重来”.

[* All-takenism: Call it the stage of "All-takenism", no matter what it is, use it first and then talk about it.]:#

[* Know the roots: After using it for nine years, I became more and more familiar with it, and I have a good understanding of some principles and problems, which is called the "knowing the roots" stage.]:#

[* Patch up: In the process of later use, there will always be some problems or some unsatisfactory functions. Since we know how it is done, we will try to change its problems or add functions. This stage is called "Patch up".]:#

[* Make a new start: Wait until it is completely understood, and find that simple modifications or adding functions can't change the underlying shortcomings. At this time, it is found that the patch time invested is more and more, and the results are a little different. Ideal. If you think about it, it might be better to 'build one' again. Anyway, you have fully understood its principle, and it can also bring a lot of convenience for adding new functions or maintenance in the future. Then build one. This stage is called "Make a new start".]:#

# Curator 简要指南

[# Introduction to Curator]:#

&nbsp;&nbsp;&nbsp;&nbsp; Apache Curator 是一个用于 Apache ZooKeeper 的 Java/JVM 客户端库,它是一种分布式协调服务.它包括一个高级 API 框架和实用程序,使 Apache ZooKeeper 的使用更加轻松和可靠.它还包括常见用例和扩展的秘诀,例如服务发现和 Java 8 异步 DSL.

&nbsp;&nbsp;&nbsp;&nbsp; 以至于 Zookeeper 的创建者 Patrick Hunt 对 Curator 框架的评价是: `Guava is to Java What Curator is to Zookeeper`. 由此可见 Curator 的优秀程度.


[&nbsp;&nbsp;&nbsp;&nbsp; Apache Curator is a Java/JVM client library for Apache ZooKeeper, a distributed coordination service. It includes a highlevel API framework and utilities to make using Apache ZooKeeper much easier and more reliable. It also includes recipes for common use cases and extensions such as service discovery and a Java 8 asynchronous DSL.]:#

[&nbsp;&nbsp;&nbsp;&nbsp; So much so that Patrick Hunt, the creator of Zookeeper, commented on the Curator framework: `Guava is to Java What Curator is to Zookeeper`. This shows how excellent Curator is.]:#

## Curator简介

[## Introduction to Curator]:#

* 封装并简化了Zookeeper原生API

* 自动化连接管理

* Watcher反复注册

* 流式编程方式

* 失败重试机制 

* 强大的分布式环境工具菜谱

* 易测试,快速启动嵌入集群

[* Wraps and simplifies Zookeeper native API]:#

[* Automated connection management]:#

[*  Watcher repeated registration]:#

[* Stream programming]:#

[* Failure retry mechanism]:# 

[* Powerful Distributed Environment Tools Recipe]:#

[* Easy to test, quickly start embedded cluster]:#

<center>
<img src="../assets/images/curator-internal-implementation-mechanism/figure-3.png" alt="Curator internal implementation mechanism" title="Github of Anigkus" >
</center>

* curator-framework：对zookeeper的底层api的一些封装.

* curator-client：提供一些客户端的操作,例如重试策略等.

* curator-recipes：封装了一些高级特性,如：Cache事件监听、选举、分布式锁、分布式计数器、分布式Barrier等.

* curator-async：Curator异步操作封装类,就像Java 8's lambdas表达式那样.

[* curator-framework: some encapsulation of zookeeper's underlying api.]:#

[* curator-client: Provides some client-side operations, such as retry strategies, etc.]:#

[* curator-recipes: encapsulates some advanced features, such as: Cache event monitoring, election, distributed lock, distributed counter, distributed Barrier, etc.]:#

[* curator-async: Curator wrapper class for asynchronous operations, just like Java 8's lambdas.]:#

## Curator 版本兼容

[## Curator Compatibility]:# 

&nbsp;&nbsp;&nbsp;&nbsp; Curator 现在市面上用的比较多的版本有 `5.X`、`4.X`、`2.X`这些.所有典型应用场景需要依赖client和framework.

* 5.x支持 ZooKeeper 3.5和3.6(优先3.6),现在zookeeper最后一个Release 是3.6.2.

* 4.x支持 ZooKeeper 3.5和3.6,现在zookeeper最后一个Release 是3.5.8.

* 2.x支持 ZooKeeper 3.x及以下,现在zookeeper最后一个Release 是3.4.14

[&nbsp;&nbsp;&nbsp;&nbsp; There are many versions of Curator currently on the market, such as `5.X`, `4.X`, `2.X`. All typical application scenarios need to rely on client and framework.]:#

[* 5.x supports ZooKeeper 3.5 and 3.6 (3.6 is preferred), now the last release of zookeeper is 3.6.2.]:#

[* 4.x supports ZooKeeper 3.5 and 3.6, now the last release of zookeeper is 3.5.8.]:#

[* 2.x supports ZooKeeper 3.x and below, now the last release of zookeeper is 3.4.14.]:#

```
<dependency>
  <groupId>org.apache.curator</groupId>
    <artifactId>curator-recipes</artifactId>
    <version>x.x</version>
</dependency>
```

## Curator快速指南（一）：创建会话

[## Curator Quick Guide (1): Creating a Session]:#

&nbsp;&nbsp;&nbsp;&nbsp; 两种方式使用,一种是通过工厂方法,一种是通过流式的builder方法.

[&nbsp;&nbsp;&nbsp;&nbsp; It is used in two ways, one is through the factory method, the other is through the flow builder method]:#

### 1.Factory

```
CuratorFramework  curatorClient = CuratorFrameworkFactory .newClient(connectString,  sessionTimeoutMs,  connectionTimeoutMs, retryPolicy)
```

### 2.Stream

```
CuratorFramework curatorClient = CuratorFrameworkFactory.builder().
connectString(connectString).
sessionTimeoutMs(sessionTimeoutMs)
.connectionTimeoutMs(connectionTimeoutMs).
retryPolicy(retryPolicy).
build()
```

参数解释:

connectString: zk集群连接字符串,格式host1:port1,host2:port2,host3:port3

connectionTimeoutMs: 连接超时时间

sessionTimeoutMs: 会话超时时间

retryPolicy: 失败重试策略


[Parameter explanation:]:#

[connectString: zk cluster connection string, format host1:port1,host2:port2,host3:port3]:#

[connectionTimeoutMs: connection timeout time]:#

[sessionTimeoutMs: session timeout]:#

[retryPolicy: Failed retry policy]:#

## Curator快速指南（二）：启动

[## Curator Quick Guide (2): Getting Started ]:#

Curator会话创建好之后,就直接启动,启动成功之后就可以进行各种操作.

[After the Curator session is created, it will be started directly. Various operations can be performed after the startup is successful.]:#

```
curatorClient.start();

curatorClient.blockUntilConnected();
```

## Curator快速指南（三）：结点操作

[## Curator Quick Guide (3): Node Operations]:#

```
curatorClient.create().withMode(CreateMode.EPHEMERAL).forPath(“/temp”);  //Create a temporary node

curatorClient.delete().forPath("/todelete");  //Delete node

curatorClient.setData().forPath("/update", "value".getBytes());  //Update Data

curatorClient.getChildren().forPath("/parentNodes");  // Get Nodes
```

## Curator快速指南（四）：Retry策略

[## Curator Quick Guide (4): Retry Strategy]:#

&nbsp;&nbsp;&nbsp;&nbsp; Curator所有的操作,包括setData create delete等等.如果失败,都允许重试系统内部了几种重试策略.开发人员可以根据需要,实现RetryPolicy接口定制自己的策略.

[&nbsp;&nbsp;&nbsp;&nbsp; All operations of Curator, including setData create delete, etc. If it fails, it is allowed to retry several retry strategies within the system. Developers can customize their own strategies by implementing the RetryPolicy interface as needed.]:#

默认重试策略:

* RetryNTimes(int n, int sleepMsBetweenRetries)
n: 最大重试次数.
sleepMsBetweenRetry: 重试间隔的时间.

* RetryOneTime(int sleepMsBetweenRetry)
sleepMsBetweenRetry: 重试间隔的时间

* RetryUntilElapsed(int maxElapsedTimeMs, int sleepMsBetweenRetries)
重试的时间超过最大时间后,就不再重试.
maxElapsedTimeMs: 最大重试时间.
sleepMsBetweenRetriees: 每次重试的间隔时间.

* ExponentialBackoffRetry(int baseSleepTimeMs, int maxRetries, int maxSleepMs)
baseSleepTimeMs: 初始sleep时间.
maxRetries: 最大重试次数.
maxSleepMs: 最大重试时间.

当前应该 sleep 的时间: baseSleepTimeMs*Math.max(1,random.nextInt(1 << retryCount+1)),随着重试次数,增加重试时间间隔,指数增加.

[Default retry policy:]:#

[* RetryNTimes(int n, int sleepMsBetweenRetries)]:#
[n: maximum number of retries.]:#
[sleepMsBetweenRetry: time between retry.]:#

[* RetryOneTime(int sleepMsBetweenRetry)]:#
[sleepMsBetweenRetry: time between retry]:#

[* RetryUntilElapsed(int maxElapsedTimeMs, int sleepMsBetweenRetries)]:#
[After the retry time exceeds the maximum time, no more retries.]:#
[maxElapsedTimeMs: Maximum retry time.]:#
[sleepMsBetweenRetriees: The interval between each retry.]:#

[* ExponentialBackoffRetry(int baseSleepTimeMs, int maxRetries, int maxSleepMs)]:#
[baseSleepTimeMs: initial sleep time.]:#
[maxRetries: Maximum number of retries.]:#
[maxSleepMs: Maximum retry time.]:#

[The current time for sleep: baseSleepTimeMs*Math.max(1, random.nextInt(1 << retryCount+1)), with the number of retries, increase the retry interval, and the index increases.]:#

## Zk结点操作核心API

[## Core API for Zk Node Operations]:#

&nbsp;&nbsp;&nbsp;&nbsp; Curator所有的zk结点操作都是通过Builder完成.每一种操作后台对应一个Builder.负责zk结点的操作.
由于所有的操作都雷同,下面以`CreateBuilderImpl`为例,说明其内部的实现.

[&nbsp;&nbsp;&nbsp;&nbsp; All zk node operations of Curator are done through Builder. Each operation background corresponds to a Builder, which is responsible for the operation of the zk node.]:#
[Since all operations are the same, the following takes `CreateBuilderImpl` as an example to illustrate its internal implementation.]:#

| 核心API | 内部实现机制 | 描述 |
| :--- | :---  | :---  |
| create() | CreateBuilderImpl | 开始创建操作, 在最后调用`forPath()`指定要操作的ZNode |
| delete() | DeleteBuilderImpl | 开始删除操作,在最后调用`forPath()`指定要操作的ZNode |
| checkExists() | ExistsBuilderImpl | 开始检查ZNode是否存在的操作, 在最后调用`forPath()`指定要操作ZNode |
| getData() | GetDataBuilderImpl | 开始获得ZNode节点数据的操作,在最后调用`forPath()`指定要操作ZNode |
| setData() | SetDataBuilderImpl | 开始设置ZNode节点数据的操作,在最后调用`forPath()`指定要操作ZNode |
| getChildren() | GetChildrenBuilderImpl | 开始获得ZNode的子节点列表,在最后调用`forPath()`指定要操作ZNode |

[| Core API | Internal implementation | Describe |]:#
[| :--- | :---  | :---  |]:#
[| create() | CreateBuilderImpl | Start creating the operation, and call f`orPath()` at the end to specify the ZNode to operate |]:#
[| delete() | DeleteBuilderImpl | Start the delete operation, Call `forPath()` at the end to specify the ZNode to be operated on |]:#
[| checkExists() | ExistsBuilderImpl | Begin the operation of checking whether the ZNode exists,At the end, call `forPath()` to specify that the ZNode is to be operated]:#
[| getData() | GetDataBuilderImpl | Start the operation to get the data of the ZNode node. At the end, call `forPath()` to specify that the ZNode is to be operated |]:#
[| setData() | SetDataBuilderImpl | Start the operation of setting the ZNode node data. At the end, call `forPath()` to specify the ZNode to be operated |]:#
[| getChildren() | GetChildrenBuilderImpl | Start getting the list of child nodes of ZNode, Call `forPath()` at the end to specify the ZNode to be manipulated |]:#

# Curator 应用场合

[# Curator Application Scenario]:#



# Curator 内部剖析

# Curator 最佳实践

# 总结

<br>

### [back](./)
