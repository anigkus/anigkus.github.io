<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >MongoDB 性能调优操作指南</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >MongoDB Performance Tuning Operation Guide</h1><br/>]:#

<center>
<img src="../assets/images/mongodb-performance-tuning-operation-guide/figure-1.jpeg" alt="Anigkus github article template title" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; 文章简要说明.<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; Some general notes on article.<br/>]:#
[> <br/>]:#

https://www.mongodb.com/docs/manual/reference/configuration-options/

# 入门调优



# 中级调优

# 高级调优

# 系统设置

## vm.zone_reclaim_mode
```bash
# 修改,重启后失效
sysctl -w vm.zone_reclaim_mode=0
# 永久修改
echo "vm.zone_reclaim_mode = 0" >> /etc/sysctl.conf 
sysctl -p
# 查询当前参数配置
sysctl -a |grep vm.zone_reclaim_mode
```


# 通用配置



# 单机版

# 主从复制

# 副本集

# 分片



<br>

### [back](./)
