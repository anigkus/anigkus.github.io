<iframe src="../detail-header.html" title="Github of Anigkus" style="height:0px,widht:0px;display:none" id="kusifreamheader"></iframe>

<h1 style="color:#606c71;text-align:center;"  id="h1">记录一个关于Git中的 ‘Permission denied’ 的解决思路</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Record a solution about 'Permission denied' in Git</h1><br/>]:#

<center>
<img src="../assets/images/record-a-solution-about-permission-denied-in-git/figure-1.jpeg" alt="Record a solution about 'Permission denied' in Git" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; 平常我们使用`git clone`代码的时候,都是安装好git工具,然后生成ssh 私钥对,并把公钥添加到github或者gitlib的SSH keys即可.但是我有时就碰到这种情况,提示`Permission denied`.那么我这里就记录下,我是如何解决这个问题的,也多给大伙提供一个解决这个问题的思路吧.<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; Usually when we use the `git clone` code, we install the git tool, then generate the ssh private key pair, and add the public key to the SSH keys of github or gitlib. But I sometimes encounter this situation, Prompt `Permission denied`. Then I will record here, how I solved this problem, and provide you with more ideas to solve this problem.<br/>]:#
[> <br/>]:#


# 问题现象

[# Problem ]:#

&nbsp;&nbsp;&nbsp;&nbsp; 这个不就是我们平常经常使用的命令吗?怎么一下子就不行了呢.从这个现象来说,应是公钥不存在.

[&nbsp;&nbsp;&nbsp;&nbsp; Isn't this a command we usually use? Why does it fail all of a sudden. From this phenomenon, it should be that the public key does not exist.]:#

```
$ git clone git@github.com:graalvm/mx.git
Cloning into 'mx'...
git@github.com: Permission denied (publickey).
fatal: Could not read from remote repository.
```

# 解决思路

[# Solutions]:#

## 常规思路

[## General idea]:#

### 生成私钥对

[### Generate private key pair]:#

```
$ ssh-keygen -o -t rsa -b 4096 -C "xxx@xxx.com"
Generating public/private rsa key pair.
Enter file in which to save the key (/Users/xxx/.ssh/id_rsa): /Users/xxx/.ssh/id_rsa_xxx #Support custom file names
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /Users/xxx/.ssh/id_rsa_xxx.
Your public key has been saved in /Users/xxx/.ssh/id_rsa_xxx.pub.
The key fingerprint is:
SHA256:QBo6c0OT2g+SmTIElhDCvuQVxwqK4Wp7JOnmwm5U1/A xxx@xxx.com
The key's randomart image is:
+---[RSA 4096]----+
|Oo. =..          |
|+= +.O           |
|*.=*B =          |
|=+*Boo E         |
|+o=..o  S        |
|.B .  .          |
|= +              |
|.* .             |
|*o.              |
+----[SHA256]-----+
```

### 复制公钥

[### Copy public key]:#

&nbsp;&nbsp;&nbsp;&nbsp; 复制公钥到gitlab重的`SSH keys`页面中去即可.

[&nbsp;&nbsp;&nbsp;&nbsp; Copy the public key to the `SSH keys` page of gitlab.]:#

```
$ cat  ~/.ssh/id_rsa_xxx.pub 
```

### 添加私钥

[### Add private key]:#

```
$ ssh-add -K ~/.ssh/id_rsa_xxx
```

## 特殊情况

[## Special cases]:#

### 查看密钥

[### View key]:#

&nbsp;&nbsp;&nbsp;&nbsp; 先用 `ssh-add -l` 查看下私钥.正常情况下应该出现RSA这行记录,如果没有就添加,但是Agent不存在.

[&nbsp;&nbsp;&nbsp;&nbsp; First use `ssh-add -l` to view the private key. Under normal circumstances, the RSA line should appear, if not, add it, but the Agent does not exist.]:#

```
$ ssh-add -l
Could not open a connection to your authentication agent.
```

### 查看代理

&nbsp;&nbsp;&nbsp;&nbsp; 那么ssh-agent是干嘛的呢,我这里只简单说下这个ssh-agent的作用,其实作用就两个.

* 当使用不同的密钥连接到不同的主机时,ssh代理可以帮助我们选择对应的密钥进行认证.

* ssh-agnet能够帮助我们免去重复的输入密码的操作.

[### View proxy]:#

[&nbsp;&nbsp;&nbsp;&nbsp; So what does ssh-agent do, I will only briefly talk about the role of this ssh-agent, but there are actually two functions.]:#

[* When using different keys to connect to different hosts, the ssh agent can help us select the corresponding key for authentication.]:#

[* ssh-agnet can help us avoid the repeated operation of entering passwords.]:#

```
# Does not exist
$ ps -ef | grep ssh-agent 
root      499761  498856  0 03:56 pts/6    00:00:00 grep --color=auto ssh-agent 

# Start Agent
$ eval `ssh-agent -s`
Agent pid 497890

# View Agent
$ ssh-agent
SSH_AUTH_SOCK=/var/folders/7g/j8kn3xy50mgdb1c7jz_vp6_c0000gn/T//ssh-wTB99grIdHS9/agent.49789; export SSH_AUTH_SOCK;
SSH_AGENT_PID=497890; export SSH_AGENT_PID;
echo Agent pid 497890;
```

### 添加密钥

[### Add private key ]:#

```
#Add default private key
$ ssh-add 
Identity added: /Users/xxx/.ssh/id_rsa (xxx@xxx.local)

# Delete default private key
$ ssh-add -d                     
Identity removed: /Users/xxx/.ssh/id_rsa RSA (xxx@xxx.local)
```

&nbsp;&nbsp;&nbsp;&nbsp; 添加指定私钥,有个注意的地方.就是如果是`Mac`系统,就使用大写 ` K `.如果是`Linux`系统,就使用小写 ` k `.

[&nbsp;&nbsp;&nbsp;&nbsp; Add the specified private key, there is a place to pay attention. That is, if it is a `Mac` system, use uppercase `K `. If it is a `Linux` system, use lowercase `k`.]:#

```
# Add the specified private key
$ ssh-add -K ~/.ssh/id_rsa_xxx      
Identity added: /Users/xxx/.ssh/id_rsa_xxx (xxx@xxx.com)

# Delete the specified private key
➜  ssh-add -d  ~/.ssh/id_rsa_xxx                   
Identity removed: /Users/xxx/.ssh/id_rsa_github RSA (xxx@xxx.com)
```

### 再次查看私钥

[### Check the private key again]:#

&nbsp;&nbsp;&nbsp;&nbsp; 如果有就表示正常,否则就表示没有添加成功或者已经过期了(过期的时间没弄明白).

[&nbsp;&nbsp;&nbsp;&nbsp; If there is, it means normal, otherwise it means that it has not been added successfully or has expired (the expiration time is not clear).]:#

```
$ ssh-add -l
4096 SHA256:bTbs6kx7yG96kNd1fGO1lgaCNIZAcZzS8OtQzKDcai8 xx@xx.com (RSA)
```

# 验证情况

&nbsp;&nbsp;&nbsp;&nbsp; 如果返回以下内容就可以了.

[# Verification status]:#

[&nbsp;&nbsp;&nbsp;&nbsp; It's ok if it returns the following.]:#

```
# Gitlab
ssh -T git@gitlab.company.com
Welcome to GitLab, @xxx!

# Github
ssh -T  git@github.com
Hi xxx! You've successfully authenticated, but GitHub does not provide shell access.
```


<br>

### [back](./)
