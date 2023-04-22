<iframe src="../detail-header.html" title="Github of Anigkus" style="height:0px,widht:0px;display:none" id="kusifreamheader"></iframe>

<h1 style="color:#606c71;text-align:center;" id="h1" >在 Alpine Linux中使用定时任务</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Using cron tasks in Alpine Linux</h1><br/>]:#

<center>
<img src="../assets/images/using-cron-tasks-in-alpine-linux/figure-1.jpeg" alt="Manually compile open source gateway konga and integrate with kong" title="Github of Anigkus" >
</center>


> <br/>&nbsp;&nbsp;&nbsp;&nbsp; 本文简单说下在Alpine Linux中如何启用定时任务.<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; This article briefly explains how to enable scheduled tasks in Alpine Linux.<br/>]:#
[> <br/>]:#

# 首先升级成最新系统
[# To upgrade to a new stable release or edge]:#

```
$ apk upgrade --available
```

# 安装openrc工具包
[# optimized toolbox]:#

```
$ apk add openrc --no-cache
```

# 安装busybox工具包
[# manages the services]:#

```
$ apk add busybox-initscripts
```

# 启动调度服务并增加系统启动项
[# Start service crond and add it to runlevel]:#

```
$ rc-service crond start && rc-update add crond
## OR ##
# $ /etc/init.d/crond start
## OR ##
# $ service crond [start|stop|status|restart]
```

# 查看调度进程
[# See crond is running]:#

```
$ rc-service -l | grep crond 
crond
```

# 定义自己的脚本任务.注意:脚本一定不能有文件后缀
[# New `task1` your scripts file. Note: The script must not have a file suffix]:#

```
$ vi /etc/periodic/15min/task1
#!/bin/sh
echo `data` >>/tmp/log.txt
```

# 授权
[# Authorized]:#

```
$ chmod 777  /etc/periodic/15min/task1
```

# 检查下脚本有无问题
[# To check whether your scripts are likely to run]:#

```
$ run-parts --test /etc/periodic/15min
```

# 设置脚本调度周期,每一分钟
[# Append schedule tasks, at every minute]:#

```
$ crontab -e
*/1     *       *       *       *       run-parts /etc/periodic/15min
```

# 查看当前用户所有定时任务
[# View all schedule tasks]:#

```
$  crontab -l
# do daily/weekly/monthly maintenance
# min   hour    day     month   weekday command
*/15    *       *       *       *       run-parts /etc/periodic/15min
0       *       *       *       *       run-parts /etc/periodic/hourly
0       2       *       *       *       run-parts /etc/periodic/daily
0       3       *       *       6       run-parts /etc/periodic/weekly
0       5       1       *       *       run-parts /etc/periodic/monthly
*/1     *       *       *       *       run-parts /etc/periodic/15min
```

# 和上面一样，只是这个是文件
[# View root user  schedule tasks]:#

```
$  cat /etc/crontabs/root
```

# 验证调度日志情况
[# View run log]:#

```
$  watch cat /tmp/log.txt
```

<br>

### [back](./)
