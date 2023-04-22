<iframe src="detail-header.html" title="Github of Anigkus" style="height:0px,widht:0px;display:none" id="kusifreamheader"></iframe>

<h1 style="color:#606c71;text-align:center;" >Using cron tasks in Alpine Linux</h1><br/>

<center>
<img src="assets/images/using-cron-tasks-in-alpine-linux/figure-1.jpeg" alt="Manually compile open source gateway konga and integrate with kong" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; This article briefly explains how to enable scheduled tasks in Alpine Linux.<br/>
> <br/>

# To upgrade to a new stable release or edge

```
$ apk upgrade --available
```

# optimized toolbox

```
$ apk add openrc --no-cache
```

# manages the services

```
$ apk add busybox-initscripts
```

# Start service crond and add it to runlevel

```
$ rc-service crond start && rc-update add crond
## OR ##
# $ /etc/init.d/crond start
## OR ##
# $ service crond [start|stop|status|restart]
```

# See crond is running

```
$ rc-service -l | grep crond 
crond
```

# New `task1` your scripts file. Note: The script must not have a file suffix

```
$ vi /etc/periodic/15min/task1
#!/bin/sh
echo `data` >>/tmp/log.txt
```

# Authorized

```
$ chmod 777  /etc/periodic/15min/task1
```

# To check whether your scripts are likely to run

```
$ run-parts --test /etc/periodic/15min
```

# Append schedule tasks, at every minute

```
$ crontab -e
*/1     *       *       *       *       run-parts /etc/periodic/15min
```

# View all schedule tasks

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

# View root user  schedule tasks

```
$  cat /etc/crontabs/root
```

# View run log

```
$  watch cat /tmp/log.txt
```

<br>

### [back](./)
