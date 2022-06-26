<iframe src="../detail-header.html" title="Github of Anigkus" style="height:0px,widht:0px;display:none" id="kusifreamheader"></iframe>

<h1 style="color:#606c71;text-align:center;" id="h1" >理解MongoDB 这些命令的区别和原理</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Understand the differences and principles of these commands in MongoDB</h1><br/>]:#

<center>
<img src="../assets/images/understand-the-differences-and-principles-of-these-commands-in-mongodb/figure-1.jpeg" alt="Understand the differences and principles of these commands in MongoDB" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; 一般我们安装 `MongoDB` 组件之后,系统都会存在一些命令.这些命令就是 `MongoDB` 的核心命令. `MongoDB` 大致命令都是从 `mongo` 和 `mongo-tools` 项目构建出来的.我们在使用 `MongoDB`一些命令必须要明白所使用命令主要提供的功能是做什么的,它们的原理又是怎样以及大致的执行逻辑是什么.并且还需要知道某些功能可能在不同的命令中都提供相似效果,那么它们之间的区别又是什么呢.话说我也使用 `MongoDB` 有一段时间了,一直也没在意怎么去内部分析过.那么本文我将详细说明下关于 `mongo, mongosh, mongod , mongos` 这几个命令的功能以及执行逻辑,并且还会指出各个命令的源码出处,并且最后还会用一张图整体说明下各个命令架构和执行流程. <br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; Some general notes on article.<br/>]:#
[> <br/>]:#

# mongo


功能说明
功能使用

执行逻辑

源码位置

# 整体逻辑


mongo, 
    https://github.com/mongodb/mongo/blob/master/src/mongo/shell/mongo.cpp
    https://its401.com/article/pourtheworld/106175737
    https://www.cnblogs.com/daizhj/archive/2011/03/17/1987311.html

mongosh,
    https://github.com/mongodb-js/mongosh

mongod , 
    https://github.com/mongodb/mongo/blob/master/src/mongo/db/mongod_main.cpp
    https://github.com/mongodb/mongo/tree/master/src/mongo/db/commands
    https://mongoing.com/archives/78285

mongos
    https://github.com/mongodb/mongo/blob/master/src/mongo/s/mongos_main.cpp
    https://github.com/mongodb/mongo/tree/master/src/mongo/s/commands
    https://mongoing.com/archives/75581



<br>

### [back](./)
