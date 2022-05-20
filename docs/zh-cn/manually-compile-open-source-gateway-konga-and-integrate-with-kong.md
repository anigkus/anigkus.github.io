<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >手动编译开源网关konga与kong集成</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Manually compile open source gateway konga and integrate with kong</h1><br/>]:#

<center>
<img src="../assets/images/manually-compile-open-source-gateway-konga-and-integrate-with-kong/figure-1.jpg" alt="Manually compile open source gateway konga and integrate with kong" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; 文章简要说明.<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; Some general notes on article.<br/>]:#
[> <br/>]:#





##Help documentation
##https://download.konghq.com/gateway-1.x-centos-7/Packages/k/
##https://download.konghq.com/gateway-2.x-centos-7/Packages/k/
##https://docs.konghq.com/install/centos/
##https://github.com/pantsel/konga
##https://github.com/Kong/kong/blob/master/kong.conf.default
#https://docs.npmjs.com/cli/v7/commands/npm-run-script
#https://nodejs.org/en/download/

##Configuration postgres
##change postgres password
postgres-# \password postgres
Enter new password: 
Enter it again: 
postgres-# 
postgres=#  CREATE DATABASE kong1;
CREATE DATABASE
postgres=# CREATE DATABASE konga1;
CREATE DATABASE
postgres=#  CREATE ROLE kong1 WITH SUPERUSER LOGIN PASSWORD 'kong1';
CREATE ROLE
postgres=# GRANT ALL PRIVILEGES ON DATABASE kong1 to kong1;
GRANT
postgres=# GRANT ALL PRIVILEGES ON DATABASE konga1 to kong1;



##Configuration kong
#wget --no-check-certificate  https://download.konghq.com/gateway-2.x-centos-7/Packages/k/kong-2.6.0.el7.amd64.rpm
// https://download.konghq.com/gateway-2.x-centos-8/Packages/k/kong-2.8.1.el8.amd64.rpm
#yum install kong-2.6.0.el7.amd64.rpm 
# kong version
2.6.0
#cp /etc/kong/kong.conf.default /etc/kong/kong.conf
#vim /etc/kong/kong.conf
# grep "^\s*[^# \t].*$" kong.conf
admin_listen = 127.0.0.1:8001,10.0.11.93:8001
pg_host = 172.16.21.48             # Host of the Postgres server.
pg_user = kong1                  # Postgres user.
pg_password =kong1                # Postgres user's password.
pg_database = kong1              # The database name to connect to.
#kong migrations bootstrap
2021/10/13 16:31:30 [warn] ulimit is currently set to "1024". For better performance set it to at least "4096" using "ulimit -n"

Bootstrapping database...
migrating core on database 'kong1'...
core migrated up to: 000_base (executed)
core migrated up to: 003_100_to_110 (executed)
core migrated up to: 004_110_to_120 (executed)
core migrated up to: 005_120_to_130 (executed)
core migrated up to: 006_130_to_140 (executed)
core migrated up to: 007_140_to_150 (executed)
core migrated up to: 008_150_to_200 (executed)
core migrated up to: 009_200_to_210 (executed)
core migrated up to: 010_210_to_211 (executed)
core migrated up to: 011_212_to_213 (executed)
core migrated up to: 012_213_to_220 (executed)
core migrated up to: 013_220_to_230 (executed)
migrating acl on database 'kong1'...
acl migrated up to: 000_base_acl (executed)
acl migrated up to: 002_130_to_140 (executed)
acl migrated up to: 003_200_to_210 (executed)
acl migrated up to: 004_212_to_213 (executed)
migrating acme on database 'kong1'...
acme migrated up to: 000_base_acme (executed)
migrating basic-auth on database 'kong1'...
basic-auth migrated up to: 000_base_basic_auth (executed)
basic-auth migrated up to: 002_130_to_140 (executed)
basic-auth migrated up to: 003_200_to_210 (executed)
migrating bot-detection on database 'kong1'...
bot-detection migrated up to: 001_200_to_210 (executed)
migrating hmac-auth on database 'kong1'...
hmac-auth migrated up to: 000_base_hmac_auth (executed)
hmac-auth migrated up to: 002_130_to_140 (executed)
hmac-auth migrated up to: 003_200_to_210 (executed)
migrating ip-restriction on database 'kong1'...
ip-restriction migrated up to: 001_200_to_210 (executed)
migrating jwt on database 'kong1'...
jwt migrated up to: 000_base_jwt (executed)
jwt migrated up to: 002_130_to_140 (executed)
jwt migrated up to: 003_200_to_210 (executed)
migrating key-auth on database 'kong1'...
key-auth migrated up to: 000_base_key_auth (executed)
key-auth migrated up to: 002_130_to_140 (executed)
key-auth migrated up to: 003_200_to_210 (executed)
migrating oauth2 on database 'kong1'...
oauth2 migrated up to: 000_base_oauth2 (executed)
oauth2 migrated up to: 003_130_to_140 (executed)
oauth2 migrated up to: 004_200_to_210 (executed)
oauth2 migrated up to: 005_210_to_211 (executed)
migrating rate-limiting on database 'kong1'...
rate-limiting migrated up to: 000_base_rate_limiting (executed)
rate-limiting migrated up to: 003_10_to_112 (executed)
rate-limiting migrated up to: 004_200_to_210 (executed)
migrating response-ratelimiting on database 'kong1'...
response-ratelimiting migrated up to: 000_base_response_rate_limiting (executed)
migrating session on database 'kong1'...
session migrated up to: 000_base_session (executed)
session migrated up to: 001_add_ttl_index (executed)
41 migrations processed
41 executed
Database is up-to-date
# kong start
2021/10/13 16:31:52 [warn] ulimit is currently set to "1024". For better performance set it to at least "4096" using "ulimit -n"
Kong started

##Configuration node
#VERSION=v14.18.0
#DISTRO=linux-x64
#wget https://nodejs.org/download/release/v14.18.0/node-$VERSION-$DISTRO.tar.xz
#mkdir -p  /usr/local/lib/nodejs
#tar -xJvf node-$VERSION-$DISTRO.tar.xz -C /usr/local/lib/nodejs
#vi ~/.node_profile
VERSION=v14.18.0
DISTRO=linux-x64
export PATH=/usr/local/lib/nodejs/node-$VERSION-$DISTRO/bin:$PATH
#vi ~/.bash_profile 
source ~/.node_profile
#sourec ~/.bash_profile 
#node -v
v14.18.0
#npm -v
6.14.15


#Configuration konga
#yum install -g git
#git clone --branch 0.14.9 https://github.com/pantsel/konga.git
#cd konga
##npm command have times  not Unstable##
#npm i
#npm i pg -g
##OR Use yarn##
##Not required,but required HangKong or foreign internet
#npm i yarn -g
#yarn -v
1.22.15
#yarn
yarn install v1.22.15
warning package-lock.json found. Your project contains lock files generated by tools other than Yarn. It is advised not to mix package managers in order to avoid resolution inconsistencies caused by unsynchronized lock files. To clear this warning, remove package-lock.json.
[1/4] Resolving packages ...
...
...
...
success Already up-to-date.
$ bower --allow-root install
bower bootstrap-switch extra-resolution Unnecessary resolution: bootstrap-switch#~3.3.4
Done in 1.79s.
#yarn add pg -g
# vi .env_example 
PORT=1337
NODE_ENV=production
KONGA_HOOK_TIMEOUT=120000
DB_ADAPTER=postgres
DB_USER=kong1
DB_PASSWORD=kong1
DB_DATABASE=konga1
DB_URI=postgresql://172.16.21.48:5432/konga1
KONGA_LOG_LEVEL=warn
TOKEN_SECRET=some_secret_token

#node ./bin/konga.js  prepare --adapter postgres --uri postgresql://kong1:kong1@172.16.21.48:5432/konga1
debug: Preparing database...
Using postgres DB Adapter.

# npm run  production

> kongadmin@0.14.9 production /root/konga
> node --harmony app.js --prod

No DB Adapter defined. Using localDB...
(node:13866) Warning: Accessing non-existent property 'padLevels' of module exports inside circular dependency
(Use `node --trace-warnings ...` to show where the warning was created)
The default `sails-disk` adapter is not designed for use as a production database;
(it stores the entire contents of your database in memory)
Instead, please use another adapter; e.g. sails-postgresql or sails-mongo.
For more info, see: http://sailsjs.org/documentation/concepts/deployment
To hide this warning message, enable `sails.config.orm.skipProductionWarnings`.



<br>

### [back](./)
