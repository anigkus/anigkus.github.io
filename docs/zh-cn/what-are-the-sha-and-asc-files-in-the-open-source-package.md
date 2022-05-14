<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >开源包中SHA和ASC两个文件到底是什么?</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >What are the SHA and ASC files in the open source package?</h1><br/>]:#

<!--文章模版标题 -->
<center>
<img src="../assets/images/what-are-the-sha-and-asc-files-in-the-open-source-package/figure-1.jpeg" alt="What are the SHA and ASC files in the open source package?" title="Github of Anigkus" >
</center>

[<center>]:#
[<img src="../assets/images/what-are-the-sha-and-asc-files-in-the-open-source-package/figure-1.jpeg" alt="What are the SHA and ASC files in the open source package?" title="Github of Anigkus" >]:#
[</center>]:#

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; 我们都知道社区提供release包时还都会提供一些验证文件.现在大多数都会有SHA文件和ASC文件.这些文件的作用就是验证软件最终发布时的内容完整性而提供的.为了防止网络中传输丢包或者有人串改了里面的内容而提供的.那你如何知道怎样去检验文件的正确性和完整性以及生成自己的SHA和ASC文件?那么如下这篇文章我将简单分享下原理以及文件制作过程.<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; We all know that the community will also provide some verification files when they provide release packages. Now most of them will have SHA files and ASC files. The purpose of these files is to verify the content integrity of the software when it is finally released. In order to prevent transmission loss in the network Packages or provided by someone who has modified the content inside. Then how do you know how to verify the correctness and integrity of the files and generate your own SHA and ASC files? Then I will briefly share the principles and file production in the following article process. <br/>]:#
[> <br/>]:#

# 关于这两个函数的背景

[# Background on both functions]:#

&nbsp;&nbsp;&nbsp;&nbsp; 先说SHA函数.行业选择SHA主要是作为数字签名.SHA分为SHA-1和SHA-2,SHA-2（Secure Hash Algorithm 2）的缩写.由美国国家标准与技术研究院（NIST）在2001年发布,属于SHA算法之一,是SHA-1的后继者.其下又可再分为六个不同的算法标准,包括了：SHA-224、SHA-256、SHA-384、SHA-512、SHA-512/224、SHA-512/256.现在软件包中用的最多的还是SHA-256、SHA-512.其中SHA1输出160bit、SHA256输出256bit、SHA512输出512bit、MD5其实是输出128bit.因为只要使用一个源文件经过SHA后会生成一串hash码,而这个hash码是和这个文件的内容是一对一的关系.只要这个文件当中修改了任何一个字符,那么重新进行hash时就会生成另一个新的hash码.因此2个hash码只要一样,"基本"就可以认为两个文件是一样的.我说的基本,因为从理论上来说,就算2个文件内部不一样,还是有可能hash出来的hash码一样.这个就是通常所说的hash冲突.解决哈希冲突的方法一般有：开放定址法(再散列法)、链地址法（HashMap的哈希冲突解决方法）、再哈希法(哈希多次)、建立公共溢出区(弄个公共区域存储冲突数据)等方法,详细说明这几种处理哈希冲突方法的原理就不在此说明.

[&nbsp;&nbsp;&nbsp;&nbsp; Let's talk about the SHA function first. The industry chooses SHA mainly as a digital signature. SHA is divided into SHA-1 and SHA-2, the abbreviation of SHA-2 （Secure Hash Algorithm 2）. Released by the National Institute of Standards and Technology (NIST) in 2001, it belongs to one of the SHA algorithms and is the successor of SHA-1. It can be further divided into six different algorithm standards, including: SHA-224, SHA-256, SHA-384, SHA-512, SHA-512/224, SHA-512/256. Now the software package uses The most common ones are SHA-256 and SHA-512. Among them, SHA1 outputs 160bit, SHA256 outputs 256bit, SHA512 outputs 512bit, and MD5 actually outputs 128bit. Because as long as a source file is used after SHA, a string of hash codes will be generated, and this hash The code is in a one-to-one relationship with the content of this file. As long as any character is modified in this file, another new hash code will be generated when re-hash is performed. Therefore, as long as the two hash codes are the same, "basic" It can be considered that the two files are the same. I said the basics, because theoretically, even if the two files are different internally, it is still possible that the hash codes from the hash are the same. This is the so-called hash conflict. The methods of Greek conflict generally include: open addressing method (re-hash method), chain address method (HashMap's hash conflict resolution method), re-hash method (hash multiple times), establishing a public overflow area (getting a public area to store conflicting data) and other methods, and detailing the principles of these methods for dealing with hash collisions will not be explained here. ]:#

&nbsp;&nbsp;&nbsp;&nbsp; 再说ASC,一般ASC文件都是使用GPG命令生成的,是另一种签名文件.要了解什么是GPG,就要先了解PGP.可以这么简单认为就是PGP是商业软件,GPG是开源社区开发用来替代PGP的.因为GPG是Gnu自由基金会社区开发出来的,因此就取名为GnuPG,这就是GPG的由来.GPG有许多用途,本文主要介绍文件信息的加密和解密.这要用到GnuPG软件,它是目前最流行、最好用的加密工具之一,其他类似加密解密的工具有RAS、OpenSSL等.

[&nbsp;&nbsp;&nbsp;&nbsp; Besides ASC, ASC files are generally generated using the GPG command, which is another type of signature file. To understand what GPG is, you must first understand PGP. It can be simply thought that PGP is commercial software, and GPG is developed by the open source community. Instead of PGP, because GPG was developed by the Gnu Freedom Foundation community, it was named GnuPG, which is the origin of GPG. GPG has many uses, This article mainly introduces the encryption and decryption of file information. This uses GnuPG software, it is one of the most popular and best encryption tools at present, other similar encryption and decryption tools are RAS, OpenSSL, etc. ]:#

用两张简单的图来概括下SHA和ASC文件的原文和密文:

[Use two simple diagrams to summarize the original text and ciphertext of SHA and ASC files: ]:#

<center>
<img src="../assets/images/what-are-the-sha-and-asc-files-in-the-open-source-package/figure-2.png" alt="What are the SHA and ASC files in the open source package?" title="Github of Anigkus" >
</center>

<center>
<img src="../assets/images/what-are-the-sha-and-asc-files-in-the-open-source-package/figure-3.png" alt="What are the SHA and ASC files in the open source package?" title="Github of Anigkus" >
</center>

<mark>Note:</mark>

&nbsp;&nbsp;&nbsp;&nbsp; ASC 的签名文件内容是根据每个人的`passphrase`密码不同以及个人信息不同,最终签名出来的内容也会是不一样的.

[&nbsp;&nbsp;&nbsp;&nbsp; The content of ASC's signature file is based on each person's `passphrase` password and personal information, and the final signed content will also be different.]:#

# 使用已有文件检验举例

[# Example of inspection with existing documents]:#

&nbsp;&nbsp;&nbsp;&nbsp; 就拿apache pulsar(原生云、分布式消息中间件)官网提供的例子说明下,如图所示有两个加密文件(asc、sha512).就以 `apache-pulsar-2.8.1-bin.tar.gz` 这个文件说明.

[&nbsp;&nbsp;&nbsp;&nbsp; Take the example provided by the official website of apache pulsar (native cloud, distributed message middleware), as shown in the figure, there are two encrypted files (asc, sha512). Just use `apache-pulsar-2.8.1-bin.tar.gz` This file description. ]:#

<center>
<img src="../assets/images/what-are-the-sha-and-asc-files-in-the-open-source-package/figure-4.png" alt="What are the SHA and ASC files in the open source package?" title="Github of Anigkus" >
</center>

我们分别看下asc和sha512文件内容是什么.

[Let's look at the contents of the asc and sha512 files respectively. ]:#

[apache-pulsar-2.8.1-bin.tar.gz.sha512](https://downloads.apache.org/pulsar/pulsar-2.8.1/apache-pulsar-2.8.1-bin.tar.gz.sha512) #HASHES

```
a981925e4c801bb5c0a4a3b683445ab36fd04ca10cf023268d9a60801a0e19a4176d1551d3a9b580f53cfcb2fe43804eabc8fab063643ae4edcf618bfacec992  ./apache-pulsar-2.8.1-bin.tar.gz
```

[apache-pulsar-2.8.1-bin.tar.gz.asc](https://downloads.apache.org/pulsar/pulsar-2.8.1/apache-pulsar-2.8.1-bin.tar.gz.asc) #SIGNATURES

```
-----BEGIN PGP SIGNATURE-----


iQIzBAABCAAdFiEEwhelw+/rwQvkHOHUpkE3jJeUTTQFAmEzBJ0ACgkQpkE3jJeU
TTSm+RAAiI6gESg1xI11ncYOjpsRVdTgO0s0k9kN3o9qfFjqMXRlOYFJLYQSkSA4
FOh7CX9OhrksWu90ez30qRLYnn2reFKgQBR433P3XDKcqt/7muxLgHYRhXz3egp5
09rRn91T0vZxkw+oY+9kN9uKwbuvDCKi++4ObGTmoFSSgVATe0HbyQw1rOrJ3TpW
e+SlzGz9I/Au/V6EvUNMHC7bftsuN3NTOUGM9pm1KzwVYZdNsYhPlvklsS2XBnjH
Kli9Nr0mwxeQN61FZsVIZ/Qwvf8KGDhwtukU1L56J42pp9jQvKIKM8aFPCypEJq0
CzN3Kw/LrbTKpdMQogj4IatajbEh/S4Ox3dhIzAzluxoN3wHwih+Rox2Sm82HsQX
jLnANW50SZmVoW9yhTYiRKjsuxGjBUmtuJUj19pEBed2IllRZz/G/j0w8BevyTPG
58TIsBrdFKjbX4LJ/M5S5QlTS5+zt+s8u+5p/PCwgqPYe3cvcZ9Dl9TOSu702bBT
AaeOosSOmMvl816Qc21sCAVg2gVJqRBWf+InnqfyWr4ocbAoV5uiGFSN6XQ/F6kB
bugxUErlLAsaWeIcQtduVfugdpVWC4D0kxvIySLoI4yJZV1eWcjYjki4MgaIYhJE
w4BgnccOAuPrE10ZDUIzxclvJulbe4E83PETTqh4dthn6Fy32Tg=
=D4N2
-----END PGP SIGNATURE-----
```

那我们怎么去检验`apache-pulsar-2.8.1-bin.tar.gz`中的 (`HASHES & SIGNATURES`) 是正确性和完整性呢?那么我们就按如下步骤开始说明:

[So how do we verify that (`HASHES & SIGNATURES`) in `apache-pulsar-2.8.1-bin.tar.gz` is correctness and integrity? Then we start the explanation as follows: ]:#

<mark>Note:</mark>

系统必须有`sha512sum`、`gpg`这两个命令,基本还不存在现有主流的系统上没有这两个命令的(😄).

[The system must have `sha512sum`,`gpg` these two commands, basically there is no existing mainstream system without these two commands(😄). ]:#

```
➜  yum install gnupg -y
➜  yum install coreutils -y
```

```
➜  sha512sum --version
sha512sum (GNU coreutils) 8.22
Copyright (C) 2013 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.


Written by Ulrich Drepper, Scott Miller, and David Madore.
➜  gpg --version
gpg (GnuPG) 2.0.22
libgcrypt 1.5.3
Copyright (C) 2013 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.


Home: ~/.gnupg
Supported algorithms:
Pubkey: RSA, ?, ?, ELG, DSA
Cipher: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
        CAMELLIA128, CAMELLIA192, CAMELLIA256
Hash: MD5, SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
Compression: Uncompressed, ZIP, ZLIB, BZIP2
[root@CentOS_Test pulsar]#
```

## 检验 HASHES 文件

[## Verify HASHES file]:#

&nbsp;&nbsp;&nbsp;&nbsp; 要检查 `HASHES` 您必须计算刚刚下载的文件的正确校验值.然后将其与已发布的原始校验值进行比较,文件哈希值只是用来检查文件没有被修改,不保证文件里面的内容真实可靠.如果由于网络下载而导致不完整或错误,那么 `HASHES` 也会对不上.

[To Verify the `HASHES` you have to calculate the correct checksum of the file you just downloaded. Then compare it with the original published checksum, the file hash is only used to check that the file has not been modified, it does not guarantee that the The content is authentic and reliable. If it is incomplete or wrong due to network download, then `HASHES` will not match. ]:#

```
# 为了环境干净,可以选择删除当前根目录下的gnupg文件
➜  rm -rf /root/.gnupg
➜  mkdir pulsar

➜  wget https://downloads.apache.org/pulsar/pulsar-2.8.1/apache-pulsar-2.8.1-bin.tar.gz
➜  wget https://downloads.apache.org/pulsar/pulsar-2.8.1/apache-pulsar-2.8.1-bin.tar.gz.asc
➜  wget https://downloads.apache.org/pulsar/pulsar-2.8.1/apache-pulsar-2.8.1-bin.tar.gz.sha512

➜   ll
total 325220
-rw-r--r--. 1 root root 333016170 Sep  4 02:06 apache-pulsar-2.8.1-bin.tar.gz
-rw-r--r--. 1 root root       833 Sep  4 02:06 apache-pulsar-2.8.1-bin.tar.gz.asc
-rw-r--r--. 1 root root       163 Sep  4 02:06 apache-pulsar-2.8.1-bin.tar.gz.sha512

# sha512sum输出的和官网提供的sha512文件一样一样,说明文件的完整性没有问题
➜  sha512sum ./apache-pulsar-2.8.1-bin.tar.gz 
a981925e4c801bb5c0a4a3b683445ab36fd04ca10cf023268d9a60801a0e19a4176d1551d3a9b580f53cfcb2fe43804eabc8fab063643ae4edcf618bfacec992  ./apache-pulsar-2.8.1-bin.tar.gz
➜  cat apache-pulsar-2.8.1-bin.tar.gz.sha512 
a981925e4c801bb5c0a4a3b683445ab36fd04ca10cf023268d9a60801a0e19a4176d1551d3a9b580f53cfcb2fe43804eabc8fab063643ae4edcf618bfacec992  ./apache-pulsar-2.8.1-bin.tar.gz
[root@CentOS_Test pulsar]# 

# 用一条命令查看也行,如果没有任何输出就可以证明和官网一样
➜  sha512sum ./apache-pulsar-2.8.1-bin.tar.gz > ./compared.sha512 && diff ./compared.sha512 ./apache-pulsar-2.8.1-bin.tar.gz.sha512
```

[# In order to clean the environment, you can choose to delete the gnupg file in the current root directory]:#
[# The output of sha512sum is the same as the sha512 file provided by the official website, indicating that there is no problem with the integrity of the file]:#
[# You can also view it with one command, if there is no output, it can be proved that it is the same as the official website]:#

## 检验 SIGNATURES 文件

[## Verify SIGNATURES file]:#

&nbsp;&nbsp;&nbsp;&nbsp; 要检查 `SIGNATURES` ,我们需要发布者的公钥.公钥一般都是发布者提供的.基本发布者都会把公钥上传去全球公开的公钥平台.可以根据 (`发信者ID`) 去下载.但是由于公钥密码学的性质,您需要另外验证密钥是否由真正的公钥官方发布者签的,还需要去比较公钥里面的发布者是否和密钥指纹的一样.签名的发布者可以有多个,大多数都在对应软件的`KEYS`文件列表中,但是有时也会不存在里面,因为和的`KEYS`文件更新机制也有关系.

[&nbsp;&nbsp;&nbsp;&nbsp; To check `SIGNATURES`, we need the public key of the publisher. The public key is usually provided by the publisher. Basically, The publisher will upload the public key to the public key platform that is open globally. (`Sender ID`) to download. However, due to the nature of public key cryptography, you need to additionally verify whether the key is signed by the real public key official issuer,and you also need to compare whether the issuer in the public key is the same as the password. The key fingerprint is the same. There can be multiple issuers of the signature, most of them are in the `KEYS` file list of the corresponding software, but sometimes they do not exist, because it is also related to the `KEYS` file update mechanism. ]:#


```
# 开始验签
➜  gpg --verify apache-pulsar-2.8.1-bin.tar.gz.asc apache-pulsar-2.8.1-bin.tar.gz
gpg: directory `/root/.gnupg' created
https://pgpkeys.mit.edu/gpg: new configuration file `/root/.gnupg/gpg.conf' created
gpg: WARNING: options in `/root/.gnupg/gpg.conf' are not yet active during this run
gpg: keyring `/root/.gnupg/pubring.gpg' created
gpg: Signature made Sat 04 Sep 2021 01:31:09 AM EDT using RSA key ID 97944D34 #没有找到
gpg: Can't check signature: No public key

# 导入认证(未找到)
➜  gpg --keyserver keyring.debian.org --recv-key 97944D34
gpg: requesting key 97944D34 from hkp server keyring.debian.org
gpgkeys: key 97944D34 can't be retrieved
gpg: no valid OpenPGP data found.


# 或者用这个(未找到)
➜  gpg --keyserver pgpkeys.mit.edu --recv-key 97944D34
gpg: requesting key 97944D34 from hkp server kpgpkeys.mit.edu
gpgkeys: key 97944D34 can't be retrieved
gpg: no valid OpenPGP data found.


# 或者用这个(找到了)
➜  gpg --keyserver keyserver.ubuntu.com --recv-key 97944D34
gpg: keyring `/root/.gnupg/secring.gpg' created
gpg: requesting key 97944D34 from hkp server keyserver.ubuntu.com


gpg: /root/.gnupg/trustdb.gpg: trustdb created
gpg: key 97944D34: public key "chenhang (CODE SIGNING KEY) <chenhang@apache.org>" imported
gpg: no ultimately trusted keys found
gpg: Total number processed: 1
gpg:               imported: 1  (RSA: 1)

# 私钥和公钥的默认位置,这三个文件只有第一次导入或者生成时才会有(gpg.conf、pubring.gpg、pubring.gpg~)
[root@CentOS_Test ~]# ll /root/.gnupg/
total 84
-rw-------. 1 root root  7680 Oct  5 05:58 gpg.conf
-rw-------. 1 root root 36708 Oct  5 06:08 pubring.gpg
-rw-------. 1 root root 36708 Oct  5 06:08 pubring.gpg~
-rw-------. 1 root root     0 Oct  5 05:58 secring.gpg
-rw-------. 1 root root  1200 Oct  5 06:08 trustdb.gpg

# 再次验签
➜  gpg --verify apache-pulsar-2.8.1-bin.tar.gz.asc apache-pulsar-2.8.1-bin.tar.gz
gpg: Signature made Sat 04 Sep 2021 01:31:09 AM EDT using RSA key ID 97944D34
gpg: Good signature from "chenhang (CODE SIGNING KEY) <chenhang@apache.org>"
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: C217 A5C3 EFEB C10B E41C  E1D4 A641 378C 9794 4D34

# 前面说过,就算签名是好的,我们也不能就信任这个密钥,我们还需要检验97944D34是不是真由chenhang@apache.org发布的,基本可以确认是由chenhang签的.
➜  gpg --fingerprint 97944D34
pub   4096R/97944D34 2021-08-18
      Key fingerprint = C217 A5C3 EFEB C10B E41C  E1D4 A641 378C 9794 4D34
uid                  chenhang (CODE SIGNING KEY) <chenhang@apache.org>
sub   4096R/AD3291ED 2021-08-18

# 我们在顺便检验下chenhang@apache.org在不在对应的KEYS文件里面呢(不一定)
➜  wget https://downloads.apache.org/pulsar/KEYS
➜  cat KEYS  | grep chenhang # 如果没找到就表示不在里面(这个也没关系)

# 如果需要在里面通过KEYS文件里面说明,把自己追加到KEYS即可,<your name>就是我们在生成自己公私钥对的时候的Real name,下面会说到
➜  head -24 KEYS 
This file contains the PGP keys of various developers.


Users: pgp < KEYS
or
       gpg --import KEYS




Developers:
    pgp -kxa <your name> and append it to this file.
or
    (pgpk -ll <your name> && pgpk -xa <your name>) >> this file.
or
    (gpg --list-sigs <your name>
    && gpg --armor --export <your name>) >> this file.


pub   rsa2048 2017-07-23 [SC] [expires: 2019-07-23]
      8C75C738C33372AE198FD10CC238A8CAAC055FD2
uid           [ultimate] Matteo Merli <mmerli@apache.org>
sig 3        C238A8CAAC055FD2 2017-07-23  Matteo Merli <mmerli@apache.org>
sub   rsa2048 2017-07-23 [E] [expires: 2019-07-23]
sig          C238A8CAAC055FD2 2017-07-23  Matteo Merli <mmerli@apache.org>


-----BEGIN PGP PUBLIC KEY BLOCK-----
```

[# start verification]:#
[# import certificate (not found)]:#
[# or use this (not found)]:#
[# or use this (found)]:#
[# The default location of the private key and the public key, these three files will only exist when they are imported or generated for the first time (gpg.conf, pubring.gpg, pubring.gpg~)]:#
[# Check again]:#
[# As mentioned earlier, even if the signature is good, we can't trust this key. We also need to check whether 97944D34 is really released by chenhang@apache.org, and it can basically be confirmed that it was signed by chenhang.]:#
[# Let's check by the way whether chenhang@apache.org is in the corresponding KEYS file (not necessarily)]:#
[# If it is not found, it means it is not there (this is okay)]:#
[# If you need to describe it in the KEYS file, you can append yourself to KEYS. <your name> is the Real name when we generate our own public and private key pair, which will be discussed below]:#

[# start verification]:# 
[# import certificate(not found)]:# 
[# or use this(not found)]:# 
[# or use this (found)]:# 
[# The default location of the private key and the public key, these three files will only exist when they are imported or generated for the first time (gpg.conf, pubring.gpg, pubring.gpg~)]:# 
[# Check agin]:# 
[# As mentioned earlier, even if the signature is good, we can't trust this key. we also need to check whether 97944D34 is really released by chenhang@apache.org, and it can basically be confirmed that it was signed by chenhang. ]:# 
[# Let's check by the way whether chenhang@apache.org is in the corresponding KEYS file (not necessarily) ]:# 
[# If it is not found, it means it is not there (this is okay)]:#
[# If you need to describe it in the KEYS file, you can append yourself to KEYS. <your name> is the Real name when we generate our own public and private key pair, which will be discussed below]:#


&nbsp;&nbsp;&nbsp;&nbsp; 如果没有找到 <font color="red">97944D34</font> 公钥文件,那么可以选择手工 `gpg —import (KEYS|xx.asc)`,也可以根据`97944D34`去全球提供的公钥库下载,因为发布者都会把自己的公钥上传到一个或者多个公钥平台,然后公钥平台会私下分发到其他公钥平台,常用的公钥平台有如下几个:

[&nbsp;&nbsp;&nbsp;&nbsp; if the <font color="red">97944D34</font> public key file is not found, you can choose manual `gpg —import (KEYS|xx.asc)`, or according to `97944D34` go to the public key library provided globally to download, because the publisher will upload his public key to one or more public key platforms, and then the public key platform will be distributed to other public key platforms privately. The commonly used public key platforms are as follows indivual: ]:#

[pgpkeys.mit.edu](http://pgpkeys.mit.edu/)

[keyring.debian.org](http://keyring.debian.org/)

[keyserver.ubuntu.com](http://keyserver.ubuntu.com/)

# 制作函数以及过程剖析

[# Make function and process analysis]:#

```
# 为了环境干净,可以选择删除当前根目录下的gnupg文件
➜  rm -rf /root/.gnupg

➜  echo "123456">input.txt
➜  ll
total 4
-rw-r--r--. 1 root root 7 Oct 30 16:57 input.txt

# 非交互方式--batch, 密码:123456
➜  gpg --batch --passphrase 123456 -c input.txt 
gpg: directory `/root/.gnupg' created
gpg: new configuration file `/root/.gnupg/gpg.conf' created
gpg: WARNING: options in `/root/.gnupg/gpg.conf' are not yet active during this run
gpg: keyring `/root/.gnupg/pubring.gpg' created

➜  ll /root/.gnupg/
total 12
-rw-------. 1 root root 7680 Oct 30 16:57 gpg.conf
drwx------. 2 root root    6 Oct 30 16:57 private-keys-v1.d
-rw-------. 1 root root    0 Oct 30 16:57 pubring.gpg
-rw-------. 1 root root  600 Oct 30 16:57 random_seed
srwxr-xr-x. 1 root root    0 Oct 30 16:57 S.gpg-agent

# 加密文件(input.txt.gpg)和源文件(input.txt)
➜  ll input.txt*
-rw-r--r--. 1 root root  7 Oct 30 16:57 input.txt
-rw-r--r--. 1 root root 55 Oct 30 16:57 input.txt.gpg

# 没有密码,查看是乱码
➜  cat input.txt.gpg 
)dЊ�Eu�?`���_�)

# 非交互方式解密输出文件内容到output.txt
➜  gpg  --batch --passphrase 123456 -o output.txt -d  input.txt.gpg 
gpg: CAST5 encrypted data
gpg: encrypted with 1 passphrase
gpg: WARNING: message was not integrity protected
➜  ll
total 12
-rw-r--r--. 1 root root  7 Oct 30 16:57 input.txt
-rw-r--r--. 1 root root 55 Oct 30 16:57 input.txt.gpg
-rw-r--r--. 1 root root  7 Oct 30 17:01 output.txt

# 查看解密后的文件内容
➜  cat output.txt
123456
```

[# In order to clean the environment, you can choose to delete the gnupg file in the current root directory]:#

[# Non-interactive --batch, pass:123456]:#
[# Encrypted file (input.txt.gpg) and source file(input.txt)]:#
[# No pass key,view is garbled]:#
[# Decrypt the output file content noo-interactively to output.txt]:#
[# View the contents of the decrypted file]:#


前面我们拿 apache 官方的 pulsar 举例,那是别人已经提供好的asc和sha512文件.那我们能否自己也提供公钥签名给别人?现在让我们来生成自己 (`publisher`) 的公钥签名并上传到公钥平台,然后在另一台机器 (`developer`) 去检验.

[Earlier, we took the official pulsar of apache as an example, which is the asc and sha512 files that others have provided. Can we also provide public key signatures to others? Now let's generate our own (`publisher`) public key signature And upload it to the public key platform, and then check it on another machine (`developer`). ]:#

* centos-linux-k8s-sealos-5 #publisher
* centos-linux-k8s-sealos-6 #developer

<mark>Note</mark>

如果出现`...generator a better chance to gain enough entropy`的提示时会一直卡住的可以用这个解决方案.

[If you get stuck at the prompt `...generator a better chance to gain enough entropy` , you can use this solution.]:#

```
➜   yum install rng-tools -y
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
epel/x86_64/metalink                                                                             | 3.8 kB  00:00:00     
 * base: mirrors.huaweicloud.com
 * epel: ftp.iij.ad.jp
 * extras: mirrors.163.com
 * updates: mirrors.163.com
base                                                                                             | 3.6 kB  00:00:00     
epel                                                                                             | 4.7 kB  00:00:00     
extras                                                                                           | 2.9 kB  00:00:00     
updates                                                                                          | 2.9 kB  00:00:00     
(1/2): epel/x86_64/updateinfo                                                                    | 1.0 MB  00:00:52     
(2/2): epel/x86_64/primary_db                                                                    | 7.0 MB  00:00:58     
Package rng-tools-6.3.1-5.el7.x86_64 already installed and latest version
Nothing to do

# 原理就是读写硬盘会影响到gpg生成,这也是很关键的,目的就是为了让发生器有更好的机会获得足够的熵数.
# 只能新建一个终端,然后执行如下命令,那么另一个卡住的节目就会马上结束.如果没有卡住,就可以不管
➜  rngd -r /dev/urandom




Initalizing available sources




Initalizing entropy source Hardware RNG Device




Enabling RDSEED rng support




Initalizing entropy source Intel RDRAND Instruction RNG




Enabling JITTER rng support




Initalizing entropy source JITTER Entropy generator
```

[# The principle is that reading and writing the hard disk will affect the gpg generation, which is also very critical, the purpose is to give the generator a better chance to obtain enough entropy.]:#
[# You can only create a new terminal, and then execute the following command, then another stuck program will end immediately. If it is not stuck, you can ignore it.]:#

## 制作 `SIGNATURES` 文件和上传

[## Mark `SIGNATURES` file and upload ]:#

```
# 制作需要校验的文件以input.tar.gz举例(不一定是tar,我只是为了和开源例子保持一样而已,任何文件都可以)
[root@centos-linux-k8s-sealos-5 ~]# echo "123456">input.txt
[root@centos-linux-k8s-sealos-5 ~]# tar -cf input.tar.gz input.txt
[root@centos-linux-k8s-sealos-5 ~]# ll
total 16
-rw-r--r--. 1 root root 10240 Oct 30 17:07 input.tar.gz
-rw-r--r--. 1 root root     7 Oct 30 17:07 input.txt

# 生成自己的公钥和私钥
[root@centos-linux-k8s-sealos-5 ~]# gpg --gen-key
gpg (GnuPG) 2.0.22; Copyright (C) 2013 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.


gpg: directory `/root/.gnupg' created
gpg: new configuration file `/root/.gnupg/gpg.conf' created
gpg: WARNING: options in `/root/.gnupg/gpg.conf' are not yet active during this run
gpg: keyring `/root/.gnupg/secring.gpg' created
gpg: keyring `/root/.gnupg/pubring.gpg' created
Please select what kind of key you want:
   (1) RSA and RSA (default)
   (2) DSA and Elgamal
   (3) DSA (sign only)
   (4) RSA (sign only)
Your selection? 1 # 加密和签名都使用RSA算法
RSA keys may be between 1024 and 4096 bits long.
What keysize do you want? (2048) 
Requested keysize is 2048 bits # 密钥越长越安全,默认是2048位
Please specify how long the key should be valid.
         0 = key does not expire
      <n>  = key expires in n days
      <n>w = key expires in n weeks
      <n>m = key expires in n months
      <n>y = key expires in n years
Key is valid for? (0) 0    # 0 = 密钥永不过期
Key does not expire at all
Is this correct? (y/N) y    # 输入y,系统就要求你提供个人信息


GnuPG needs to construct a user ID to identify your key.


Real name: zhang san    # 真实姓名(这个名字就是我们使用pgp加入KEYS要用到的名字,记住这个名字,本机导出公钥、私钥、加密、上传都会用到)
Email address: zhangsan@test.com    # 邮箱
Comment: zhangsan build # “说明”是可以留空着
You selected this USER-ID:
    "zhang san (zhangsan build) <zhangsan@test.com>"


Change (N)ame, (C)omment, (E)mail or (O)kay/(Q)uit? o    # 输入O表示"确定".
You need a Passphrase to protect your secret key. # 系统会让你设定一个私钥的密码(这里是12345678),会弹出一个对话框,提升你输入密码,这是为了防止误操作,或者系统被侵入时有人擅自动用私钥


We need to generate a lot of random bytes. It is a good idea to perform    # 开始生成,需要你做一些随机的举动,以生成一个随机数,比如(像是敲打键盘、移动鼠标、读写硬盘之类的)
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
gpg: /root/.gnupg/trustdb.gpg: trustdb created
gpg: key 7103FB49 marked as ultimately trusted    # gpg: 密钥 7103FB49 被标记为绝对信任(记住这个7103FB49,后续其它系统检验签名时会用到)
public and secret key created and signed.         # 公钥和私钥已经生成并经签名


gpg: checking the trustdb
gpg: 3 marginal(s) needed, 1 complete(s) needed, PGP trust model
gpg: depth: 0  valid:   1  signed:   0  trust: 0-, 0q, 0n, 0m, 0f, 1u
pub   2048R/7103FB49 2021-10-30
      Key fingerprint = ECC1 4051 0DE8 4FEC 3D09  42DD 952F D490 7103 FB49
uid                  zhang san (zhangsan build) <zhangsan@test.com>
sub   2048R/CF80F28A 2021-10-30

# 导出公钥和私钥的备份(私钥和密码千万不要泄露)
[root@centos-linux-k8s-sealos-5 ~]# gpg --batch --passphrase 12345678 --armor --output zhang-san-pubkey.asc --export 'zhang san'
[root@centos-linux-k8s-sealos-5 ~]# gpg --batch --passphrase 12345678 --armor --output zhang-san-privkey.asc --export-secret-keys 'zhang san'
[root@centos-linux-k8s-sealos-5 ~]# ll
total 24
-rw-r--r--. 1 root root 10240 Oct 30 17:07 input.tar.gz
-rw-r--r--. 1 root root     7 Oct 30 17:07 input.txt
-rw-r--r--. 1 root root  3606 Oct 30 17:46 zhang-san-privkey.asc
-rw-r--r--. 1 root root  1739 Oct 30 17:45 zhang-san-pubkey.asc

[root@centos-linux-k8s-sealos-5 ~]# cat zhang-san-pubkey.asc    # 以下内容我删除当中一些字符,避免泄露
-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v2.0.22 (GNU/Linux)


mQENBGF9uFEBCAC8BWBucNgIrlBTExw+N3U5qtYQps1MBbzmOF7D++I2OBE4mT0i
whcFtlTKESLCXCWMLDdiPOZ8NWcIzamJRsCbN2K6oBrEi1wxxj1h2dg7q3Ek6fSu
ZBl477eRGWmsCPXAY85W8ZzzrAlzr9f+YT0SMsLsIHAqUrZ8WHdlsxv7UzYAMbmE
e2Ia/o9eC+ST7iz6pzJ51B1CLRxI9c+pI0vhOcCp8DZvAOGZmOB8nPtZliSsYYOx
yBWoYtBU8bn2MDYpneKXvuWMmU2XbHWzmeBzu6jhUUMYXoeE4ptMquODJ+0Wc5lI
HbDY9acDK7X4ipwEEj6thxR4FP3FYM3Z6UgtABEBAAG0LnpoYW5nIHNhbiAoemhh
bmdzYW4gYnVpbGQpIDx6aGFuZ3NhbkB0ZXN0LmNvbT6JATkEEwECACMFAmF9uFEC
YSV52zpV8aJFQXWHfuS5Gz93+p4hi6rqlrCphdzRgNM2OtnUpNKn/fciOnTaWpMx
Ho5VtsuLI9UWZXcgnPxv9ZpRzRhorAO2NHYhFqC68UtG+etJQO2G75sID493BVb7
6wMgy+6lev+7p31JJWSw2vnGwQQMmUWzgaTS8pIuFiloSl8E3luLIaNZtypA6c1+
Rx68DBL3T0TsC5YMp7Pm+HDkTOgfcrKX3EKqo0SGC5LRbYxf6ylm5hAiffBee78x
IPeEGFDUtenjwAdNn5ZLoVoTqUDNALizSuv0ecYTz3KAqJeENu1NLV4guh9a7Mp3
U/f2bV32jY3tfujrj5+PpisAtxbCwDbGEPXxaPYFyc94m2T+yE5kUPrfYJGADCLT
IMLqcAbifSn8mVmlydIz2s6E0DChdHNBTs257qckg2R0TWFTdg/rABEBAAGJAR8E
GAECAAkFAmF9uFECGwwACgkQlS/UkHED+0n3pgf+IiRM7jnEZGDaZpOJbOG9R6CD
aNJdQItqfdeCxMswV7d1Js7TRTF67+S6CM747abgrt7WKSsLuoOxqoY9HwsExwHY
QC4OfymXprGNmH06zB9HOPfDY/28psSwleMkOAxRecGXJ3gJsFB0qdXTVGIk19sy
n3ZcerrkIboW4PYnw4w/bPuoK1fzlawJfnpmUEsYjgty7LbIwq0zSVvoFXaSxsCm
h+50YNShfmAviSxOuqIypyPI6Kmv+YKqss8+90jJZ3q76XTrLT8vGnxcZQjCEfos
FY604Zvv1U8gr/Qj6cXy5bIVYL/2/s1MjnxN8jg3Q2WAybCinS/yE9/np745hQ==
=aOmF
-----END PGP PUBLIC KEY BLOCK-----

[root@centos-linux-k8s-sealos-5 ~]# cat zhang-san-privkey.asc    # 以下内容我删除当中一些字符,避免泄露
-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: GnuPG v2.0.22 (GNU/Linux)


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
=u1ay
-----END PGP PRIVATE KEY BLOCK-----
[root@centos-linux-k8s-sealos-5 ~]# 

# 加密:一旦导出和密钥备份完成,我们现在可以加密和解密.tar.gz文件. 使用以下命令加密
[root@centos-linux-k8s-sealos-5 ~]# gpg --batch --passphrase 12345678 --encrypt --recipient 'zhang san' input.tar.gz
[root@centos-linux-k8s-sealos-5 ~]# ll
total 28
-rw-r--r--. 1 root root 10240 Oct 30 17:07 input.tar.gz
-rw-r--r--. 1 root root   453 Oct 30 17:50 input.tar.gz.gpg
-rw-r--r--. 1 root root     7 Oct 30 17:07 input.txt
-rw-r--r--. 1 root root  3606 Oct 30 17:46 zhang-san-privkey.asc
-rw-r--r--. 1 root root  1739 Oct 30 17:45 zhang-san-pubkey.asc

# 解密
[root@centos-linux-k8s-sealos-5 ~]# gpg --batch --passphrase 12345678 --output output.tar.gz --decrypt input.tar.gz.gpg
gpg: encrypted with 2048-bit RSA key, ID CF80F28A, created 2021-10-30
      "zhang san (zhangsan build) <zhangsan@test.com>"
[root@centos-linux-k8s-sealos-5 ~]# ll
total 40
-rw-r--r--. 1 root root 10240 Oct 30 17:07 input.tar.gz
-rw-r--r--. 1 root root   453 Oct 30 17:50 input.tar.gz.gpg
-rw-r--r--. 1 root root     7 Oct 30 17:07 input.txt
-rw-r--r--. 1 root root 10240 Oct 30 17:50 output.tar.gz
-rw-r--r--. 1 root root  3606 Oct 30 17:46 zhang-san-privkey.asc
-rw-r--r--. 1 root root  1739 Oct 30 17:45 zhang-san-pubkey.asc
[root@centos-linux-k8s-sealos-5 ~]# cp output.tar.gz /tmp/
[root@centos-linux-k8s-sealos-5 ~]# cd /tmp/
[root@centos-linux-k8s-sealos-5 tmp]# tar -xf output.tar.gz 

# 和input.tar.gz.gpg里面的input.txt原文一样
[root@centos-linux-k8s-sealos-5 tmp]# cat input.txt 
123456
[root@centos-linux-k8s-sealos-5 tmp]# 




# 根据源文件生成SIGNATURES文件,也就是input.tar.gz.asc,
[root@centos-linux-k8s-sealos-5 ~]# gpg --batch --passphrase 12345678  --armor --detach-sign input.tar.gz
[root@centos-linux-k8s-sealos-5 ~]# ll
total 44
-rw-r--r--. 1 root root 10240 Oct 30 17:07 input.tar.gz
-rw-r--r--. 1 root root   490 Oct 30 19:23 input.tar.gz.asc
-rw-r--r--. 1 root root   453 Oct 30 17:50 input.tar.gz.gpg
-rw-r--r--. 1 root root     7 Oct 30 17:07 input.txt
-rw-r--r--. 1 root root 10240 Oct 30 17:50 output.tar.gz
-rw-r--r--. 1 root root  3606 Oct 30 17:46 zhang-san-privkey.asc
-rw-r--r--. 1 root root  1739 Oct 30 17:45 zhang-san-pubkey.asc
[root@centos-linux-k8s-sealos-5 ~]# cat input.tar.gz.asc
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2.0.22 (GNU/Linux)


iQEcBAABAgAGBQJhfdP5AAoJEJUv1JBxA/tJ/dwH/04ij5GmNq1d01AMuuyORalG
rqQDjjtnU+ucbbw443VXbiPN/fe1Umkj2Px0Oaa89QgU6vhgXyrMj7XHHu8+anel
XlxsEOt5MOOk8ZyeNhwxAiev7TNkR2AbvY7OFeMm3KYNn/6kCMLXaxTQZMXG5Fl8
gf2COqC+GbpE51+h6OBc6Lja3AR2cSudc9sV5UiZMymkrJWeqIQsCh1qMh4uoRWg
kLG40RUDuuXXVKN91bUm84LRvSJDkjYyXM3hGn7LUHaWuNbqgAVllkU1pQE3wUCJ
hkOScEdrdWynR2cmFYtMR2sQZOH21eqNh/4KaSFYIwl3NkGuudPBNjN+koDGC4A=
=redV
-----END PGP SIGNATURE-----
[root@centos-linux-k8s-sealos-5 ~]# 

# 上传本机公钥文件(切记不要频繁上传不同的公钥,一旦上传,没法删除,这样全球提供的公钥平台就会出现好多无用的公钥,一个是浪费资源,一个是会泄露你的个人信息(比如名字、邮箱)等);
[root@centos-linux-k8s-sealos-5 ~]# gpg --keyserver keyserver.ubuntu.com  --send-keys ‘7103FB49' # 需要gpg的key,而不是名字:zhang san
gpg: sending key 7103FB49 to hkp server keyserver.ubuntu.com


[root@centos-linux-k8s-sealos-5 ~]# sha512sum input.tar.gz
cb211c54508c09950af44253503fd34dfec9a91f4a91f97787646a8f23fd8725ce9b0d997c8f373534d8168816f8034c87a980765e89632c249eb5d17941ea62  input.tar.gz

# 把源文件和签名文件发给对方centos-linux-k8s-sealos-6,我把公钥 zhang-san-pubkey.asc也发过去,这样开发者就可以直接导入即可,也不要去网上下载,两种方式我都会说到
[root@centos-linux-k8s-sealos-5 ~]# scp input.tar.gz input.tar.gz.asc zhang-san-pubkey.asc root@centos-linux-k8s-sealos-6:/root/
root@centos-linux-k8s-sealos-6's password: 
input.tar.gz                                                                          100%   10KB   9.4MB/s   00:00    
input.tar.gz.asc                                                                      100%  490   514.7KB/s   00:00    
zhang-san-pubkey.asc                                                                  100% 1739     2.4MB/s   00:00    
[root@centos-linux-k8s-sealos-5 ~]#
```

[# Take input.tar.gz as an example to make a file that needs to be verified (not necessarily tar, I just keep it the same as the open source example, any file is fine)]:#
[# Generate your own public and private keys]:#
[# Both encryption and signature use RSA algorithm]:#
[# The longer the key, the more secure it is, the default is 2048 bits]:#
[# 0 = key never expires]:#
[# Enter y, the system will ask you to provide personal information]:#
[# Real name (this name is the name we use to add KEYS using pgp, remember this name, it will be used for exporting public key, private key, encryption, and uploading)]:#
[# Mail]:#
[# "Description" can be left blank]:#
[# Enter O for "OK".]:#
[# The system will ask you to set a password for the private key (here is 12345678), and a dialog box will pop up to prompt you to enter the password. This is to prevent misoperation, or someone will use the private key without authorization when the system is hacked]:#
[# Start generating, you need to do some random actions to generate a random number, such as (such as typing on the keyboard, moving the mouse, reading and writing the hard disk, etc.)]:#
[# gpg: The key 7103FB49 is marked as absolute trust (remember this 7103FB49, it will be used by other systems to verify the signature later)]:#
[# Public and private keys have been generated and signed]:#
[# Export the backup of the public key and private key (the private key and password must not be leaked)]:#
[# I delete some characters in the following content to avoid leakage]:#
[# I delete some characters in the following content to avoid leakage]:#
[# Encryption: Once the export and key backup are complete, we can now encrypt and decrypt the .tar.gz file. Encrypt with the following commands]:#
[# decrypt]:#
[# is the same as the original input.txt in input.tar.gz.gpg]:#
[# Generate SIGNATURES files from source files, that is, input.tar.gz.asc,]:#
[# Upload the local public key file (remember not to upload different public keys frequently. Once uploaded, it cannot be deleted, so there will be a lot of useless public keys on the public key platform provided in the world, one is a waste of resources, the other will leak Your personal information (such as name, email), etc.);]:#
[# needs gpg key, not name: zhang san]:#
[# Send the source file and signature file to the other party centos-linux-k8s-sealos-6, I also send the public key zhang-san-pubkey.asc, so that the developer can import it directly without going online Download, I will talk about both ways]:#

## 检验发布者提供的源文件和 `SIGNATURES` 文件

[## Verify the source and `SIGNATURES` files privided by the publisher ]:#

```
[root@centos-linux-k8s-sealos-6 ~]# ll
total 20
-rw-r--r--. 1 root root 10240 Oct 30 12:22 input.tar.gz
-rw-r--r--. 1 root root   490 Oct 30 12:22 input.tar.gz.asc
-rw-r--r--. 1 root root  1739 Oct 30 12:22 zhang-san-pubkey.asc

# 为了环境干净,可以选择删除当前根目录下的gnupg文件
[root@centos-linux-k8s-sealos-6 ~]# rm -rf /root/.gnupg

# 检验HASHES,说明文件完整性没有问题
[root@centos-linux-k8s-sealos-6 ~]# sha512sum input.tar.gz
cb211c54508c09950af44253503fd34dfec9a91f4a91f97787646a8f23fd8725ce9b0d997c8f373534d8168816f8034c87a980765e89632c249eb5d17941ea62  input.tar.gz

# 检验签名,提示没找到7103FB49这个公钥
[root@centos-linux-k8s-sealos-6 ~]# gpg --verify  input.tar.gz.asc input.tar.gz
gpg: Signature made Sat 30 Oct 2021 07:23:37 PM EDT using RSA key ID 7103FB49
gpg: Can't check signature: No public key

# 搜索公共库
[root@centos-linux-k8s-sealos-5 ~]# gpg --keyserver  keyserver.ubuntu.com --search-keys 'zhang san'
gpg: searching for "zhang san" from hkp server keyserver.ubuntu.com
(1)       1024 bit DSA key 59D78701, created: 2017-10-17
(2)       2048 bit RSA key 25AF6782, created: 2018-11-18
(3)     Zhang San <zhangsana168@gmail.com>
          4096 bit RSA key BBD47D17, created: 2015-04-25
(4)     zhang san (zhangsan build) <zhangsan@test.com>    # 因为 "zhang san" 被多人都上传过,因此有多个,而我刚才上传的是第四这个,也就是 RSA key 7103FB49 这个
          2048 bit RSA key 7103FB49, created: 2021-10-30
(5)     zhang_san <zzdk733@163.com>
          2048 bit RSA key 981F080A, created: 2012-07-25
(6)       2048 bit RSA key 5F1185AF, created: 2013-03-16
Keys 1-6 of 6 for "zhang san".  Enter number(s), N)ext, or Q)uit > 4
gpg: requesting key 7103FB49 from hkp server keyserver.ubuntu.com
gpg: key 7103FB49 was created 17227 seconds in the future (time warp or clock problem)
gpg: key 7103FB49 was created 17227 seconds in the future (time warp or clock problem)
gpg: key 7103FB49 was created 17227 seconds in the future (time warp or clock problem)
gpg: key 7103FB49 was created 17227 seconds in the future (time warp or clock problem)
gpg: key 7103FB49: no valid user IDs
gpg: this may be caused by a missing self-signature
gpg: Total number processed: 1
gpg:           w/o user IDs: 1


# 导入签名(基于网络,和下面的基于文件导入模式任选其一即可)
[root@centos-linux-k8s-sealos-6 ~]# rm -rf /root/.gnupg
[root@centos-linux-k8s-sealos-6 ~]# gpg --keyserver keyserver.ubuntu.com --recv-key 7103FB49
gpg: directory `/root/.gnupg' created
gpg: new configuration file `/root/.gnupg/gpg.conf' created
gpg: WARNING: options in `/root/.gnupg/gpg.conf' are not yet active during this run
gpg: keyring `/root/.gnupg/secring.gpg' created
gpg: keyring `/root/.gnupg/pubring.gpg' created
gpg: requesting key 7103FB49 from hkp server keyserver.ubuntu.com
gpg: key 7103FB49 was created 16961 seconds in the future (time warp or clock problem)    # 时间对不上,导入失败
gpg: key 7103FB49 was created 16961 seconds in the future (time warp or clock problem)
gpg: key 7103FB49 was created 16961 seconds in the future (time warp or clock problem)
gpg: key 7103FB49 was created 16961 seconds in the future (time warp or clock problem)
gpg: key 7103FB49: no valid user IDs
gpg: this may be caused by a missing self-signature
gpg: Total number processed: 1
gpg:           w/o user IDs: 1

# 设置时间(因为我的时间不正确才有此步骤)
[root@centos-linux-k8s-sealos-6 ~]# date
Sat Oct 30 12:45:05 EDT 2021
Try 'date --help' for more information.
[root@centos-linux-k8s-sealos-6 ~]# date -s '2021-10-31 21:20:20'
Sun Oct 31 21:20:20 EDT 2021
[root@centos-linux-k8s-sealos-6 ~]# date
Sun Oct 31 21:20:22 EDT 2021

# 导入签名(基于网络)
[root@centos-linux-k8s-sealos-6 ~]# gpg --keyserver keyserver.ubuntu.com --recv-key 7103FB49
gpg: requesting key 7103FB49 from hkp server keyserver.ubuntu.com
gpg: /root/.gnupg/trustdb.gpg: trustdb created
gpg: key 7103FB49: public key "zhang san (zhangsan build) <zhangsan@test.com>" imported
gpg: Total number processed: 1
gpg:               imported: 1  (RSA: 1)

# 检验签名(签名正确)
[root@centos-linux-k8s-sealos-6 ~]# gpg --verify  input.tar.gz.asc input.tar.gz
gpg: Signature made Sat 30 Oct 2021 07:23:37 PM EDT using RSA key ID 7103FB49
gpg: Good signature from "zhang san (zhangsan build) <zhangsan@test.com>"
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: ECC1 4051 0DE8 4FEC 3D09  42DD 952F D490 7103 FB49

# 检验签名的签发者是否是zhang san
[root@centos-linux-k8s-sealos-6 ~]# gpg --fingerprint 7103FB49
pub   2048R/7103FB49 2021-10-30
      Key fingerprint = ECC1 4051 0DE8 4FEC 3D09  42DD 952F D490 7103 FB49
uid                  zhang san (zhangsan build) <zhangsan@test.com>
sub   2048R/CF80F28A 2021-10-30


# 导入签名(基于文件)
[root@centos-linux-k8s-sealos-6 ~]# rm -rf /root/.gnupg
[root@centos-linux-k8s-sealos-6 ~]# ll
total 20
-rw-r--r--. 1 root root 10240 Oct 30 12:22 input.tar.gz
-rw-r--r--. 1 root root   490 Oct 30 12:22 input.tar.gz.asc
-rw-r--r--. 1 root root  1739 Oct 30 12:22 zhang-san-pubkey.asc

[root@centos-linux-k8s-sealos-6 ~]# gpg --import /root/zhang-san-pubkey.asc
gpg: directory `/root/.gnupg' created
gpg: new configuration file `/root/.gnupg/gpg.conf' created
gpg: WARNING: options in `/root/.gnupg/gpg.conf' are not yet active during this run
gpg: keyring `/root/.gnupg/secring.gpg' created
gpg: keyring `/root/.gnupg/pubring.gpg' created
gpg: /root/.gnupg/trustdb.gpg: trustdb created
gpg: key 7103FB49: public key "zhang san (zhangsan build) <zhangsan@test.com>" imported
gpg: Total number processed: 1
gpg:               imported: 1  (RSA: 1)

# 检验签名(签名正确)
[root@centos-linux-k8s-sealos-6 ~]# gpg --verify  input.tar.gz.asc input.tar.gz
gpg: Signature made Sat 30 Oct 2021 07:23:37 PM EDT using RSA key ID 7103FB49
gpg: Good signature from "zhang san (zhangsan build) <zhangsan@test.com>"
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: ECC1 4051 0DE8 4FEC 3D09  42DD 952F D490 7103 FB49

# 检验签名的签发者是否是zhang san
[root@centos-linux-k8s-sealos-6 ~]#  gpg --fingerprint 7103FB49
pub   2048R/7103FB49 2021-10-30
      Key fingerprint = ECC1 4051 0DE8 4FEC 3D09  42DD 952F D490 7103 FB49
uid                  zhang san (zhangsan build) <zhangsan@test.com>
sub   2048R/CF80F28A 2021-10-30
```

[# For a clean environment, you can choose to delete the gnupg file in the current root directory]:#
[# Check HASHES, indication that there is no problem with the integrity of the file ]:#
[# Check the signature, prompting that the public key 7103FB49 was not found ]:#
[# search public repositories]:#
[# Decause "zhang san" has been uploaded by many people, there are many, and I just uploaded the fourth one, which is the RSA key 7103FB49 ]:#
[# Import signature (based on network, and choose one of file-based import mode below)]:#
[# Time does not match, import failed ]:#
[# Set the time (this step is required because my time is incorrect)]:#
[# Import signature (network based) ]:#
[# Verify signature (network based) ]:#
[# Check if the issuer of the signature is zhang san ]:#
[# Import signature (file based)]:#
[# Check the signature (the signature is correct)]:#
[# Check if the issuer of the signature is zhang san ]:#


# 结论

[# Conclusion ]:#

&nbsp;&nbsp;&nbsp;&nbsp; 通过以上三个方面(背景、检验、剖析),了解到什么是 Crypto files (`知其然`),也知道怎么去检验 Crypto files (`知其所以然`),也知道自己怎么制作 Crypto files 并发给别人 (`知其所以必然`).

[&nbsp;&nbsp;&nbsp;&nbsp; Through the above three aspects (background, inspection,analysis), understand what is Crypto files (`Know the truth`), and know how to verify Crypto files (`Know the reason`), Also know how to make Crypto files and send them to others (`Knowing why it is inevitable`). ]:#

<br>

### [back](./)
