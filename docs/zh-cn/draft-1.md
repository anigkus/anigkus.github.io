 参考地址
https://dev.mysql.com/doc/refman/5.7/en/validate-password-plugin-installation.html
https://dev.mysql.com/doc/refman/5.7/en/validate-password-options-variables.html
以validate_password举例

#下面那些变量可以通过my.cnf或者 SET GLOBAL 设置，前提是要安装并加载这个插件,如果加载了，查看这个插件就会显示如下信息
mysql> SHOW VARIABLES LIKE 'validate_password%';
+--------------------------------------+--------+
| Variable_name                        | Value  |
+--------------------------------------+--------+
 #1或者ON为打开，0或者OFF为关闭，意思就是是否在创建密码时把密码和当前用户去做一个比较，如果它们匹配则拒绝创建
| validate_password_check_user_name    | OFF    |    
#当 validate_password_policy 为2是，也就是STRONG时才有效果，当密码匹配这个文件里面每行的的数据时，文件内容全是小写，如果匹配，就拒绝设置，这个文件最后为绝对路径,并且启动mysql服务的那个用户能有这个文件的读取权限
| validate_password_dictionary_file    | /var/lib/mysql/password_file.txt       |
#密码的最小长度
| validate_password_length             | 8      |
#最少有一个大写字符
| validate_password_mixed_case_count   | 1      |
#最少有一个数字字符
| validate_password_number_count       | 1      |
#密码策略,0 or LOW、1 or MEDIUM、2 or STRONG
| validate_password_policy             | MEDIUM |
#最少有一个特殊字符
| validate_password_special_char_count | 1      |
+--------------------------------------+--------+


#文件内容
[root@localhost mysql]# cat password_file.txt
!@##x123r
123123w!u

1、查看plugin目录
     SHOW VARIABLES LIKE 'plugin_dir%';
     +---------------+--------------------------+
      | Variable_name | Value                    |
     +---------------+--------------------------+
      | plugin_dir    | /usr/lib64/mysql/plugin/ |
     +---------------+--------------------------+
     1 row in set (0.00 sec)
2、首先确认/usr/lib64/mysql/plugin/有validate_password.so的文件
     ll |  grep "validate_password"
     -rwxr-xr-x 1 root root   217202 Jun 22 23:01 validate_password.so
3、安装,有两个方式
     方式1:
     编辑/etc/my.cnf,保存,重启服务即可看到插件的插件状态
     [mysqld]
     #validate_password=OFF #我暂时注释掉禁用插件,使用下面加载插件并且默认启用,然后重启msql 服务
     plugin-load-add=validate_password.so #这种方式加载只能通过SHOW PLUGINS  ;
     进入mysql控制台,看下面已经激活了:
     方式2:
               进入mysql控制台,看下面已经激活了:
               INSTALL PLUGIN validate_password SONAME 'validate_password.so';  #只能查看mysql.plugin或者INFORMATION_SCHEMA.PLUGINS两张表，通过 SHOW PLUGINS ;也能看到
4、查看插件
     方式1
          mysql>  SHOW PLUGINS  ;
+----------------------------+----------+--------------------+----------------------+---------+
| Name                       | Status   | Type               | Library              | License |
+----------------------------+----------+--------------------+----------------------+---------+
| binlog                     | ACTIVE   | STORAGE ENGINE     | NULL                 | GPL     |
| mysql_native_password      | ACTIVE   | AUTHENTICATION     | NULL                 | GPL     |
| sha256_password            | ACTIVE   | AUTHENTICATION     | NULL                 | GPL     |
| MRG_MYISAM                 | ACTIVE   | STORAGE ENGINE     | NULL                 | GPL     |
| PERFORMANCE_SCHEMA         | ACTIVE   | STORAGE ENGINE     | NULL                 | GPL     |
| MEMORY                     | ACTIVE   | STORAGE ENGINE     | NULL                 | GPL     |
| InnoDB                     | ACTIVE   | STORAGE ENGINE     | NULL                 | GPL     |
| INNODB_TRX                 | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_LOCKS               | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_LOCK_WAITS          | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_CMP                 | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_CMP_RESET           | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_CMPMEM              | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_CMPMEM_RESET        | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_CMP_PER_INDEX       | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_CMP_PER_INDEX_RESET | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_BUFFER_PAGE         | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_BUFFER_PAGE_LRU     | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_BUFFER_POOL_STATS   | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_TEMP_TABLE_INFO     | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_METRICS             | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_FT_DEFAULT_STOPWORD | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_FT_DELETED          | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_FT_BEING_DELETED    | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_FT_CONFIG           | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_FT_INDEX_CACHE      | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_FT_INDEX_TABLE      | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_SYS_TABLES          | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_SYS_TABLESTATS      | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_SYS_INDEXES         | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_SYS_COLUMNS         | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_SYS_FIELDS          | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_SYS_FOREIGN         | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_SYS_FOREIGN_COLS    | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_SYS_TABLESPACES     | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_SYS_DATAFILES       | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| INNODB_SYS_VIRTUAL         | ACTIVE   | INFORMATION SCHEMA | NULL                 | GPL     |
| CSV                        | ACTIVE   | STORAGE ENGINE     | NULL                 | GPL     |
| MyISAM                     | ACTIVE   | STORAGE ENGINE     | NULL                 | GPL     |
| FEDERATED                  | DISABLED | STORAGE ENGINE     | NULL                 | GPL     |
| BLACKHOLE                  | ACTIVE   | STORAGE ENGINE     | NULL                 | GPL     |
| partition                  | ACTIVE   | STORAGE ENGINE     | NULL                 | GPL     |
| ARCHIVE                    | ACTIVE   | STORAGE ENGINE     | NULL                 | GPL     |
| ngram                      | ACTIVE   | FTPARSER           | NULL                 | GPL     |
| validate_password          | DISABLED | VALIDATE PASSWORD  | validate_password.so | GPL     |
| test_plugin_server         | ACTIVE   | AUTHENTICATION     | auth_test_plugin.so  | GPL     |
+----------------------------+----------+--------------------+----------------------+---------+
46 rows in set (0.00 sec)
     方式2
mysql> SELECT MPL.PLUGIN_NAME,MPL.PLUGIN_LIBRARY,MPL.PLUGIN_STATUS FROM  INFORMATION_SCHEMA.PLUGINS MPL WHERE MPL.PLUGIN_NAME  LIKE '%test_plugin_server%';
+--------------------+---------------------+---------------+
| PLUGIN_NAME        | PLUGIN_LIBRARY      | PLUGIN_STATUS |
+--------------------+---------------------+---------------+
| test_plugin_server | auth_test_plugin.so | ACTIVE        |
+--------------------+---------------------+---------------+
1 row in set (0.00 sec)
      方式3
          mysql> select * from mysql.plugin;
+--------------------+----------------------+
| name               | dl                   |
+--------------------+----------------------+
| test_plugin_server | auth_test_plugin.so  |
| validate_password  | validate_password.so |
+--------------------+----------------------+
2 rows in set (0.00 sec)
 
5、卸载插件
     方式1
          如果通过INSTALL PLUGIN 安装的插件，需要通过使用以下方式卸载:
          UNINSTALL PLUGIN validate_password; #这样在mysql.plugin表中就是空的了
     方式2:
          如果有在/etc/my.cnf显示加载插件,删除加载插件那行代码即可，然后重启服务，就看不到插件了 SHOW PLUGINS  ;
          
               