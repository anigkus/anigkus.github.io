<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >Record a solution about 'Permission denied' in Git</h1><br/>

<center>
<img src="assets/images/record-a-solution-about-permission-denied-in-git/figure-1.jpeg" alt="Record a solution about 'Permission denied' in Git" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; Usually when we use the `git clone` code, we install the git tool, then generate the ssh private key pair, and add the public key to the SSH keys of github or gitlib. But I sometimes encounter this situation, Prompt `Permission denied`. Then I will record here, how I solved this problem, and provide you with more ideas to solve this problem.<br/>
> <br/>

# Problem 

&nbsp;&nbsp;&nbsp;&nbsp; Isn't this a command we usually use? Why does it fail all of a sudden. From this phenomenon, it should be that the public key does not exist.

```
$ git clone git@github.com:graalvm/mx.git
Cloning into 'mx'...
git@github.com: Permission denied (publickey).
fatal: Could not read from remote repository.
```

# Solutions

## General idea

### Generate private key pair

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

### Copy public key

&nbsp;&nbsp;&nbsp;&nbsp; Copy the public key to the `SSH keys` page of gitlab.

```
$ cat  ~/.ssh/id_rsa_xxx.pub 
```

### Add private key

```
$ ssh-add -K ~/.ssh/id_rsa_xxx
```

## Special cases

### View key

&nbsp;&nbsp;&nbsp;&nbsp; First use `ssh-add -l` to view the private key. Under normal circumstances, the RSA line should appear, if not, add it, but the Agent does not exist.

```
$ ssh-add -l
Could not open a connection to your authentication agent.
```

### View proxy

&nbsp;&nbsp;&nbsp;&nbsp; So what does ssh-agent do, I will only briefly talk about the role of this ssh-agent, but there are actually two functions.

* When using different keys to connect to different hosts, the ssh agent can help us select the corresponding key for authentication.

* ssh-agnet can help us avoid the repeated operation of entering passwords.

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

### Add private key 

```
#Add default private key
$ ssh-add 
Identity added: /Users/xxx/.ssh/id_rsa (xxx@xxx.local)

# Delete default private key
$ ssh-add -d                     
Identity removed: /Users/xxx/.ssh/id_rsa RSA (xxx@xxx.local)
```

&nbsp;&nbsp;&nbsp;&nbsp; Add the specified private key, there is a place to pay attention. That is, if it is a `Mac` system, use uppercase `K `. If it is a `Linux` system, use lowercase `k`.

```
# Add the specified private key
$ ssh-add -K ~/.ssh/id_rsa_xxx      
Identity added: /Users/xxx/.ssh/id_rsa_xxx (xxx@xxx.com)

# Delete the specified private key
➜  ssh-add -d  ~/.ssh/id_rsa_xxx                   
Identity removed: /Users/xxx/.ssh/id_rsa_github RSA (xxx@xxx.com)
```

### Check the private key again

&nbsp;&nbsp;&nbsp;&nbsp; If there is, it means normal, otherwise it means that it has not been added successfully or has expired (the expiration time is not clear).

```
$ ssh-add -l
4096 SHA256:bTbs6kx7yG96kNd1fGO1lgaCNIZAcZzS8OtQzKDcai8 xx@xx.com (RSA)
```

# Verification status

&nbsp;&nbsp;&nbsp;&nbsp; It's ok if it returns the following.

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
