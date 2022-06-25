---
layout: default
title: 解读 mongo, mongosh, mongod , mongos 是什么
description: Interpretation of mongo, mongosh, mongod , mongos What is
---

<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >解读 mongo, mongosh, mongod , mongos 是什么</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Interpretation of mongo, mongosh, mongod , mongos What is</h1><br/>]:#

<center>
<img src="../assets/images/figure-1.jpeg" alt="Anigkus github article template title" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; 文章简要说明.<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; Some general notes on article.<br/>]:#
[> <br/>]:#


背景

功能说明
功能使用

执行逻辑

源码位置


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
