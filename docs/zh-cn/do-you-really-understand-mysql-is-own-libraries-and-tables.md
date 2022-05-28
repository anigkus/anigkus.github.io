<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >MySQL自带库和表你真的理解?</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Do you really understand MySQL's own libraries and tables?</h1><br/>]:#

<center>
<img src="../assets/images/do-you-really-understand-mysql-is-own-libraries-and-tables/figure-1.jpeg" alt="MySQL自带库和表你真的理解?" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp;MySQL的历史可以追溯到1979年,后来几十年不断开发和迭代,从1996年MySQL 1.0发布到后来的3.x、4.x、5.x及最新发布的8.x,不断开发新功能和完善问题,获得大众的喜欢和使用,但是数据库自带内置的数据库和表也变化了不少,我就以5.7.19(不过大版本的差异还是很大,小版本基本都差不多)这个版本来了解下现在的存在的自带库和表都有哪些,都来存储哪些数据,分别有什么用,能给我们日常数据库维护或者开发工作中带来什么便利,让我们一起来看看吧.<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; The history of MySQL can be traced back to 1979, and it has been continuously developed and iterated for decades. From the release of MySQL 1.0 in 1996 to the later 3.x, 4.x, 5.x and the latest 8.x, new functions have been continuously developed And improve the problem, it is liked and used by the public, but the built-in database and tables of the database have also changed a lot. I will use 5.7.19 (but the difference between the major version is still very large, and the minor version is basically the same) Let's find out what are the existing self-contained libraries and tables, what data are stored, what are they used for, and what convenience they can bring to our daily database maintenance or development work, let's take a look.<br/>]:#
[> <br/>]:#

# INFORMATION_SCHEMA 
[# INFORMATION_SCHEMA ]:#

&nbsp;&nbsp;&nbsp;&nbsp;information_schema数据库是MySQL自带的,它提供了访问数据库元数据的方式.什么是元数据呢？元数据是关于数据的数据,如数据库名或表名,列的数据类型,或访问权限等.有些时候用于表述该信息的其他术语包括“数据词典”和“系统目录”.在MySQL中,把 information_schema 看作是一个数据库,确切说是信息数据库.其中保存着关于MySQL服务器所维护的所有其他数据库的信息.如数据库名,数据库的表,表栏的数据类型与访问权 限等.在INFORMATION_SCHEMA中,有数个只读表.它们实际上是视图,而不是基本表,因此,你将无法看到与之相关的任何文件.information_schema数据库表说明:

[&nbsp;&nbsp;&nbsp;&nbsp;The information_schema database comes with MySQL and provides a way to access database metadata. What is metadata? Metadata is data about data, such as database or table names, data types of columns, or access rights, etc. Other terms sometimes used to describe this information include "data dictionary" and "system catalog". In MySQL, information_schema is regarded as a database, to be precise, an information database. It holds information about all other databases maintained by the MySQL server. Such as database name, database table, table column data type and access rights, etc. In INFORMATION_SCHEMA, there are several read-only tables. They are actually views, not base tables, so you won't be able to see any files related to them. Information_schema database table description:]:#

| 序号 | 表名 | 含义 |
| :--- | :---  | :---  |
| 1 | CHARACTER_SETS | 字符和校对规则默认对应关系表,这个表是无法增加、更新、删除的.|
| 2 | COLLATIONS | 字符和校对规则所有对应关系表,并且有校对ID,默认标识、是否编译进服务器等信息（这些暂时都是MySQL扩展信息,现在没用）,这个表是无法增加、更新、删除的.|
| 3 | COLLATION_CHARACTER_SET_APPLICABILITY | 字符和校对规则所有对应关系表. |
| 4 | COLUMNS | MySQL当前服务器所有数据表中列的信息,如类型、是否为空、长度、字符编码、校对编码等好多有用信息. |
| 5 | COLUMN_PRIVILEGES | 和mysql.columns_priv几乎一样,就都了一列IS_GRANTABLE：用来说明用户已是否可以继续授权他人. |
| 6 | ENGINES | 当前MySQL服务所有的存储引擎,引擎状态(支持、默认、禁用并且不支持、支持但禁用)、那个支持事务、XA事务、 部分事务等. |
| 7 | EVENTS | MySQL定时任务表,和mysql.event几乎一样,功能也差不多,但是在数据上表现上会有点差别,主要是时间上. |
| 8 | FILES | MySQL文件表,只要存储innodb和NDB文件信息. |
| 9 | GLOBAL_STATUS | 全局系统状态表,MySQL为了兼容mysql5.6而特意保留的表,和 SHOW GLOBAL STATUS命令及 SELECT * FROM performance_schema.global_status  (5.7特有的)是一样的效果,不过后续可能会去掉,更多兼容表. |
| 10 | GLOBAL_VARIABLES | 全局系统变量表,MySQL为了兼容mysql5.6而特意保留的表,和 SHOW GLOBAL VARIABLES 命令及 SELECT * FROM performance_schema.global_variables  (5.7特有的)是一样的效果,不过后续可能会去掉. |
| 11 | KEY_COLUMN_USAGE | MySQL所有库的主键约束记录表. |
| 12 | OPTIMIZER_TRACE | 如果想看整个执行计划以及对于多种索引方案之间是如何选择的？MySQL5.6中支持这个功能,optimizer_trace,默认关闭的,session级别,并与explain、profile、optimizer_trace 之MySQL执行计划三大马车. |
| 13 | PARAMETERS | MySQL存储过程或函数参数记录表,比mysqL.proc更详细. |
| 14 | PARTITIONS | MySQL表空间记录表,默认innodb是共享存储的,就是存储数据信息和索引信息都是在同一个*.ibd里面,而Myisam 存储引擎,它默认使用独立表空间,不过可以通过参数改变. |
| 15 | PLUGINS | MySQL插件表,和SHOW PLUGINS这个命令一样,但是比这个命令更详细点,如表示这个插件所属公司、版本、作者、加载方式等. |
| 16 | PROCESSLIST | MySQL线程运行表,和SHOW FULL PROCESSLIST这个命令一样,但是和SHOW PROCESSLIST这个命令是不一样的,好像和权限有关,performance_schema.threads表也可以查看线程,官网说PROCESSLIST有负面的性能影响,但是 thread不会有,暂时没特别测试明显的区别. |
| 17 | PROFILING | QUERY_ID相当于show profile for query QUERY_ID;这个,show profile只显示最后一次,通过SET profiling = 1;打开开关,Show profiles显示当前session所有的记录.. |
| 18 | REFERENTIAL_CONSTRAINTS | MySQL引用约束表,但通过ADD CONSTRAINT就会在此表有记录了. |
| 19 | ROUTINES | MySQL包含存储过程和函数表,不包括UDFs函数记录. |
| 20 | SCHEMATA | MySQL所有的库定义的数据表,包括库名字、字符、校对规则等. |
| 21 | SCHEMA_PRIVILEGES | MySQL数据库级别权限表,和mysql.db是一样的. |
| 22 | SESSION_STATUS | MySQL5.7为了兼容5.6的保留表,和show status一样的效果. |
| 23 | SESSION_VARIABLES | MySQL5.7为了兼容5.6的保留表,和show VARIABLES 一样的效果. |
| 24 | STATISTICS | MySQL服务器上所有索引记录表,查询单个索引情况：SHOW INDEX FROM TABLE_NAME; 和SHOW INDEX FROM TABLE_NAME FROM DB_NAME. |
| 25 | TABLES | MySQL服务器上所有表记录表,可能与当前表内容不同步,但可以通过运行ANALYZE来更新它. |
| 26 | TABLESPACES | MySQL表空间记录表,此表不提供有关InnoDB表空间的信息,需要查询INFORMATION_SCHEMA.FILES表. |
| 27 | TABLE_CONSTRAINTS | MySQL约束表,包括唯一约束、主键约束 、外键约束相关记录表. |
| 28 | TABLE_PRIVILEGES | MySQL数表级别权限表,和mysql.tables_priv差不多一样的. |
| 29 | TRIGGERS |MySQL触发器存档表. |
| 30 | USER_PRIVILEGES | MySQL用户权限表,和 mysql.user差不多一样的. |
| 31 | VIEWS | MySQL视图表. |
| 32 | INNODB_LOCKS | MySQL innodb事务锁等待表,用来分析锁等待,(事务相关的表). |
| 33 | INNODB_TRX | MySQL innodb所有已经打开的事务表,关闭就自动消失了,(事务相关的表). |
| 34 | INNODB_SYS_DATAFILES | MySQL innodb 数据文件表,只包含*.idb文件的磁盘路径数据,默认是相对mysql data目录的,创建时可以指定绝对路径. |
| 35 | INNODB_FT_CONFIG | MySQL全文索引配置表,只有当字段类型的列上定义全文索引,当执行set global innodb_ft_aux_table='DB_NAME/TABLE_NAME';就会有记录,具体怎么用后续在研究. |
| 36 | INNODB_SYS_VIRTUAL |MySQL虚拟字段记录表,当列为虚拟列时就会有记录. |
| 37 | INNODB_CMP | 压缩相关的表,压缩状态, mysql表数据压缩,innodb表压缩,如何设置mysql innodb 表的压缩,InnoDB记录压缩及使用分析, 其中innodb_cmp保存历史汇总数据,而innodb_cmp_reset则记录的是一个较为实时的统计值. |
| 38 | INNODB_FT_BEING_DELETED | MySQL全文索引准备删除的记录表. |
| 39 | INNODB_CMP_RESET | 压缩相关的表,压缩状态, mysql表数据压缩,innodb表压缩,如何设置mysql innodb 表的压缩,InnoDB记录压缩及使用分析. |
| 40 | INNODB_CMP_PER_INDEX | 压缩相关的表,该表提供每一张表和索引的压缩情况,innodb压缩. |
| 41 | INNODB_CMPMEM_RESET | 压缩相关的表,缓存压缩数据 mysql表数据压缩,innodb表压缩,如何设置mysql innodb 表的压缩,InnoDB记录压缩及使用分析. |
| 42 |INNODB_FT_DELETED  | MySQL全文索引已经删除的记录表. |
| 43 | INNODB_BUFFER_PAGE_LRU | 缓冲池中每个页面的信息,此表来观察LRU（最近最少使用的页面置换算法）列表中每个页的具体信息. |
| 44 | INNODB_LOCK_WAITS | 事务相关的表,锁等待表. |
| 45 | INNODB_TEMP_TABLE_INFO | 此表表示临时表的相关信息. |
| 46 | INNODB_SYS_INDEXES | 此表表示系统所有表中索引元数据. |
| 47 | INNODB_SYS_TABLES | 此表表示系统所有表的元数据. |
| 48 | INNODB_SYS_FIELDS | T此表表示系统所有表中字段和索引位置的映射逻辑情况. |
| 49 | INNODB_CMP_PER_INDEX_RESET | (压缩相关的表) innodb_cmp_per_index_reset跟 innodb_cmp_per_index差不多,区别就是每次查询时会重置之前的统计. |
| 50 | INNODB_BUFFER_PAGE | 缓冲池中每个页面的信息. |
| 51 | INNODB_FT_DEFAULT_STOPWORD | MySQL全文索引停用的关键字表. |
| 52 | INNODB_FT_INDEX_TABLE |MySQL全文索引表. |
| 53 | INNODB_FT_INDEX_CACHE | MySQL全文索引缓存表. |
| 54 | INNODB_SYS_TABLESPACES | MySQL表空间使用情况表. |
| 55 | INNODB_METRICS | MySQL监控表,5.6引入的,此表用来监控InnoDB运行是否正常. |
| 56 | INNODB_SYS_FOREIGN_COLS | Mysql外键引用关系表和REFERENTIAL_CONSTRAINTS有点相似,但是REFERENTIAL_CONSTRAINTS更详细一点,而REFERENTIAL_CONSTRAINTS没法表示位置,而这个表可以. |
| 57 | INNODB_CMPMEM | (压缩相关的表), 缓存压缩数据mysql表数据压缩,innodb表压缩,如何设置mysql innodb 表的压缩,InnoDB记录压缩及使用分析.  |
| 58 | INNODB_BUFFER_POOL_STATS | 缓冲池中状态信息和SHOW ENGINE INNODB STATUS命令中的一些数据相似,不过此表提供更多的信息）. |
| 59 | INNODB_SYS_COLUMNS | MySQL库中所有列的对应的表ID(此表INNODB_SYS_TABLES存取ID和表名的关系)、列名、表中的位置、列的类型、长度等相关信息. |
| 60 | INNODB_SYS_FOREIGN | 库中所有外键列关系信息表,和INNODB_SYS_FOREIGN_COLS、REFERENTIAL_CONSTRAINTS都有点相似. |
| 61 | INNODB_SYS_TABLESTATS | MySQL库中所有表的所有数据行的数量情况,当DELETE\UPDATE后会更新,如果未提交的事务正在插入或从表中删除,可能会不准确.这个表中的NUN_ROWS和TABLES中TABLE_ROWS有啥区别呢,因为有时我发现TABLE_ROWS中的数据对不上表中实际的数据,难道是事务未提交的关系? |

###### &nbsp;
[| Serial Number | Table Name | Meaning |]:#
[| :--- | :---  | :---  |]:#
[| 1 | CHARACTER_SETS | The default correspondence table between characters and proofreading rules cannot be added, updated or deleted.|]:#
[| 2 | COLLATIONS | All correspondence tables between characters and collation rules, and there are information such as collation ID, default identification, whether to compile into the server (these are all MySQL extension information for the time being, and are useless now). This table cannot be added, updated, or deleted.|]:#
[| 3 | COLLATION_CHARACTER_SET_APPLICABILITY | All correspondence between characters and proofreading rules |]:#
[| 4 | COLUMNS | Information about the columns in all data tables of the current MySQL server, such as type, whether it is empty, length, character encoding, proofreading encoding and many other useful information. |]:#
[| 5 | COLUMN_PRIVILEGES | Almost the same as mysql.columns_priv, there is a column IS_COLUMN_PRIVILEGES: used to indicate whether the user can continue to authorize others. |]:#
[| 6 | ENGINES | All storage engines currently served by MySQL, engine status(supported, default, disabled, and not supportd, supported but disabled), which supports transactions,XA transactions, partica transactions,etc. |]:#
[| 7 | EVENTS | The MySQL timed task tanle is almost the same as mysql.event, and has similar fuctions, but there will be a little difference in data performance, mainly in terms of time. |]:#
[| 8 | FILES | MySQL file table,mainly store innodb and NDB file information. |]:#
[| 9 | GLOBAL_STATUS | The global system status table, a table reserved by MySQL for compatibility with mysql5.6, has the same effect as the SHOW GLOBAL STATUS command and SELECT * FROM performance_schema.global_status (specific to 5.7), but it may be removed later, more compatible tables. |]:#
[| 10 | GLOBAL_VARIABLES | The global system variable table, a table reserved by MySQL for compatibility with mysql5.6, has the same effect as the SHOW GLOBAL VARIABLES command and SELECT * FROM performance_schema.global_variables (specific to 5.7), but it may be removed later. |]:#
[| 11 | KEY_COLUMN_USAGE | Primary key constraint record table for all MySQL libraries. |]:#
[| 12 | OPTIMIZER_TRACE | If you want to see the entire execution plan and how to choose between multiple indexing schemes? This function is supported in MySQL 5.6, optimizer_trace, which is closed by default, session level, and three carriages with the MySQL execution plan of explain, profile and optimizer_trace. |]:#
[| 13 | PARAMETERS | MySQL stored procedure or function parameter record table, more detailed than mysql.proc. |]:#
[| 14 | PARTITIONS | MySQL table space record table, the default innodb is shared storage, that is, the stored data information and index information are in the same *.ibd, and the Myisam storage engine, it uses an independent table space by default, but it can be changed by parameters. |]:#
[| 15 | PLUGINS | The MySQL plugin table is the same as the command SHOW PLUGINS, but it is more detailed than this command, such as the company, version, author, loading method, etc. of the plugin. |]:#
[| 16 | PROCESSLIST | The MySQL thread running table is the same as the SHOW FULL PROCESSLIST command, but different from the SHOW PROCESSLIST command. It seems to be related to permissions. The performance_schema.threads table can also view threads. The official website says that PROCESSLIST has a negative performance impact, but threads do not Yes, I haven't tested the obvious difference for the time being. |]:#
[| 17 | PROFILING | QUERY_ID is equivalent to show profile for query QUERY_ID; this, show profile only displays the last time, through SET profiling = 1; turn on the switch, Show profiles displays all records of the current session. |]:#
[| 18 | REFERENTIAL_CONSTRAINTS | MySQL reference constraint table, but through ADD CONSTRAINT there will be records in this table. |]:#
[| 19 | ROUTINES | MySQL contains stored procedures and function tables, excluding UDFs function records. |]:#
[| 20 | SCHEMATA | All MySQL library-defined data tables, including library names, characters, collation rules, etc. |]:#
[| 21 | SCHEMA_PRIVILEGES | MySQL database level privilege table, the same as mysql.db. |]:#
[| 22 | SESSION_STATUS | MySQL 5.7 has the same effect as show status in order to be compatible with the reserved table of 5.6. |]:#
[| 23 | SESSION_VARIABLES | MySQL 5.7 has the same effect as show VARIABLES in order to be compatible with the reserved table of 5.6. |]:#
[| 24 | STATISTICS | All index records on the MySQL server table, query a single index situation: SHOW INDEX FROM TABLE_NAME; and SHOW INDEX FROM TABLE_NAME FROM DB_NAME. |]:#
[| 25 | TABLES | All tables on the MySQL server record the table, which may be out of sync with the current table contents, but can be updated by running ANALYZE. |]:#
[| 26 | TABLESPACES | MySQL table space record table, this table does not provide information about InnoDB table space, you need to query the INFORMATION_SCHEMA.FILES table. |]:#
[| 27 | TABLE_CONSTRAINTS | MySQL constraint tables, including unique constraints, primary key constraints, foreign key constraints related record tables. |]:#
[| 28 | TABLE_PRIVILEGES | MySQL table-level privilege table, almost the same as mysql.tables_priv. |]:#
[| 29 | TRIGGERS | MySQL trigger archive table. |]:#
[| 30 | USER_PRIVILEGES | MySQL user permission table, almost the same as mysql.user. |]:#
[| 31 | VIEWS | MySQL view table. |]:#
[| 32 | INNODB_LOCKS | MySQL innodb transaction lock waiting table, used to analyze lock waiting, (transaction related table). |]:#
[| 33 | INNODB_TRX | All open transaction tables in MySQL innodb disappear automatically when closed, (transaction-related tables). |]:#
[| 34 | INNODB_SYS_DATAFILES | MySQL innodb data file table, only contains the disk path data of *.idb files, the default is relative to the mysql data directory, you can specify an absolute path when creating. |]:#
[| 35 | INNODB_FT_CONFIG | MySQL full-text index configuration table, only when the full-text index is defined on the column of the field type, when set global innodb_ft_aux_table='DB_NAME/TABLE_NAME'; there will be records, how to use it in the follow-up research. |]:#
[| 36 | INNODB_SYS_VIRTUAL | MySQL virtual field record table, when it is listed as a virtual column, there will be records. |]:#
[| 37 | INNODB_CMP | Compression related tables, compression status, mysql table data compression, innodb table compression, how to set the compression of mysql innodb table, InnoDB record compression and usage analysis, innodb_cmp saves historical summary data, and innodb_cmp_reset records a more real-time statistics value. |]:#
[| 38 | INNODB_FT_BEING_DELETED | MySQL full-text index ready to delete record table. |]:#
[| 39 | INNODB_CMP_RESET | Compression related tables, compression status, mysql table data compression, innodb table compression, how to set the compression of mysql innodb table, InnoDB record compression and usage analysis. |]:#
[| 40 | INNODB_CMP_PER_INDEX | Compress related tables, which provide the compression of each table and index, innodb compression. |]:#
[| 41 | INNODB_CMPMEM_RESET | Compress related tables, cache compressed data mysql table data compression, innodb table compression, how to set the compression of mysql innodb table, InnoDB record compression and usage analysis. |]:#
[| 42 |INNODB_FT_DELETED  | MySQL full-text index deleted records table. |]:#
[| 43 | INNODB_BUFFER_PAGE_LRU | Information of each page in the buffer pool, this table to observe the specific information of each page in the LRU (Least Recently Used Page Replacement Algorithm) list. |]:#
[| 44 | INNODB_LOCK_WAITS | Transaction-related tables, lock waiting tables. |]:#
[| 45 | INNODB_TEMP_TABLE_INFO | This table represents information about temporary tables. |]:#
[| 46 | INNODB_SYS_INDEXES | This table represents index metadata in all tables of the system. |]:#
[| 47 | INNODB_SYS_TABLES | This table represents metadata for all tables in the system. |]:#
[| 48 | INNODB_SYS_FIELDS | This table represents the mapping logic of fields and index positions in all tables of the system. |]:#
[| 49 | INNODB_CMP_PER_INDEX_RESET | (Compression related tables) innodb_cmp_per_index_reset is similar to innodb_cmp_per_index, the difference is that the previous statistics are reset each time a query is made. |]:#
[| 50 | INNODB_BUFFER_PAGE | Information about each page in the buffer pool. |]:#
[| 51 | INNODB_FT_DEFAULT_STOPWORD | MySQL full-text index disabled keyword table. |]:#
[| 52 | INNODB_FT_INDEX_TABLE | MySQL full-text index table. |]:#
[| 53 | INNODB_FT_INDEX_CACHE | MySQL full-text index cache table. |]:#
[| 54 | INNODB_SYS_TABLESPACES | MySQL tablespace usage table. |]:#
[| 55 | INNODB_METRICS | MySQL monitoring table, introduced in 5.6, this table is used to monitor whether InnoDB is running normally. |]:#
[| 56 | INNODB_SYS_FOREIGN_COLS | Mysql foreign key reference relational table is somewhat similar to REFERENTIAL_CONSTRAINTS, but REFERENTIAL_CONSTRAINTS is more detailed, and REFERENTIAL_CONSTRAINTS cannot indicate location, and this table can. |]:#
[| 57 | INNODB_CMPMEM | (compress related tables), cache compressed data mysql table data compression, innodb table compression, how to set the compression of mysql innodb table, InnoDB record compression and usage analysis.  |]:#
[| 58 | INNODB_BUFFER_POOL_STATS | The status information in the buffer pool is similar to some data in the SHOW ENGINE INNODB STATUS command, but this table provides more information). |]:#
[| 59 | INNODB_SYS_COLUMNS | Corresponding table ID of all columns in the MySQL database (the relationship between the access ID and table name of this table INNODB_SYS_TABLES), column name, position in the table, column type, length and other related information. |]:#
[| 60 | INNODB_SYS_FOREIGN | All foreign key column relationship information tables in the library are somewhat similar to INNODB_SYS_FOREIGN_COLS and REFERENTIAL_CONSTRAINTS. |]:#
[| 61 | INNODB_SYS_TABLESTATS | The number of all data rows in all tables in the MySQL database will be updated after DELETE\UPDATE, and may be inaccurate if uncommitted transactions are being inserted or deleted from the table. What is the difference between NUN_ROWS in this table and TABLE_ROWS in TABLES, because sometimes I find that the data in TABLE_ROWS does not match the actual data in the table, is it because the transaction is not committed? |]:#


# MYSQL 
[# MYSQL]:#

&nbsp;&nbsp;&nbsp;&nbsp;mysql的核心数据库,类似于sql server中的master表,主要负责存储数据库的用户、权限设置、关键字等mysql自己需要使用的控制和管理信息,mysql 数据库表说明:

[&nbsp;&nbsp;&nbsp;&nbsp;The core database of mysql, similar to the master table in sql server, is mainly responsible for storing the control and management information that mysql needs to use, such as database users, permission settings, keywords, etc. mysql database table description:]:#

| 序号 | 表名 | 含义 |
| :--- | :---  | :---  |
| 1 | columns_priv | 表中列的权限表,当执行GRANT SELECT (HOST), ON `db1`.* TO 'test'@'%' identified by "123456"; 就会有数据.|
| 2 | db | 数据库的权限表,当执行 GRANT SELECT  ON `db1`.* TO 'test'@'%' identified by "123456";就会有数据.|
| 3 | engine_cost | 代价模型 ,engine_cost  (IO代价)、 server_cost( CPU代价), mysql 执行计划分析三看explain,profiling,optimizer_trace. |
| 4 | event | 定时任务表,当通过create event event_now...,就会有数据,注意定时任务只是有效才会在event表中,默认事件执行完即释放,如立即执行事件,执行完后,事件便自动删除. |
| 5 | func | UDF：用户定义函数,不是Fuction哦,更高级别的函数. |
| 6 | general_log | MySQL 通用查询日志,当log_output=TABLE并且开始general_log时,会把所有的mysql  执行日志写入到这张表中. |
| 7 | gtid_executed | MySQL全局事务标识持久表.  |
| 8 | help_category | 帮助信息表,关于帮助主题类别的信息,通过在help '关键字';查看帮助信息,通过help_category、help_keyword、help_topic里面的name字段搜索. |
| 9 | help_keyword | 帮助信息表,与帮助主题相关的关键字信息. |
| 10 | help_relation | 帮助信息表,帮助关键字信息和主题信息之间的映射. |
| 11 | help_topic | 帮助信息表,帮助主题的详细内容. |
| 12 | innodb_index_stats | 统计信息表,存储索引维度的统计信息,用来分析SQL解析和查询优化相当有效. |
| 13 | innodb_table_stats | 统计信息表 ,存储表维度的统计信息. |
| 14 | ndb_binlog_index | DB Cluster replication是用,演示太麻烦,后续测试. |
| 15 | plugin | MySQL已经加载的插件表,通过my.cnf或者INSTALL PLUGIN方式加载的插件就会在这个表有记录. |
| 16 | proc | MySQL存储过程和自定义Function记录表,当创建PROCEDURE或FUNCTION就会有记录,Mysq.5.7.19函数和存储过程的权限比以前强大多了,语法上也有一点点变化了. |
| 17 | procs_priv | 存储过程权限表,当执行GRANT EXECUTE ON PROCEDURE `db1`.procedure_name TO 'someuser'@'%';；就会有数据. |
| 18 | proxies_priv | 代理用户权限表,当通过grant proxy授权代理用户时才会出现在这张表中 . |
| 19 | server_cost | 代价模型 ,engine_cost  (IO代价)、 server_cost( CPU代价), mysql 执行计划分析三看explain,profiling,optimizer_trace.|
| 20 | servers | 远程数据用户表,当使用federated引擎时用到,可以做到远程数据和本地数据共享,但是功能还不全.使用CREATE SERVER...创建,有个注意的地方就是信息一定需要在第一次写全,不然后续好像没法通过命令修改,只有手动删除servers表的数据,并重启服务才能再次创建相同的server_name,然后再次使用,FEDERATED引擎的作用就是本机可以和远程有相同的数据表,表名可以不一样,但是本机不保存数据,只有表定义,数据在远程数据库上面. |
| 21 | slave_master_info | MySQL服务器主 binlog信息,当开启binlog日志时并且master_info_repository=TABLE时会把相关信息记录到这张表. |
| 22 | slave_relay_log_info | MySQL服务器从的中继器binlog信息,当从开启binlog日志时并且relay_log_info_repository=TABLE.时会把相关信息记录到这张表. |
| 23 | slave_worker_info | MySQL从服务多线程管理表,当从服务器slave_parallel_type='logical_clock';（配置成基于逻辑时钟的方式）,默认是DATABASE（就是一个库只有一个线程到主服务器拉取dump日志信息）,并且全局变量slave_parallel_workers>0才会在这张表（mysql.slave_worker_info）出现数据,默认performance_schema.replication_applier_status_by_worker只有一个单线程数据.|
| 24 | slow_log | MySQL 慢查询日志,当log_output=TABLE并且开始slow_log 时,会把所有的mysql  执行日志写入到这张表中. |
| 25 | tables_priv | 数据表权限表,当执行 GRANT SELECT  ON `db1`.tb1 TO 'test'@'%' identified by "123456";就会有数据. |
| 26 | time_zone | MySQL时区相关表,通过mysql_tzinfo_to_sql加载相关时区信息就会有数据,Linux系统时区文件在/usr/share/zoneinfo/这个目录里面,我当时系统是CentOS_6.6_Final_2.6.32-504.23.4.el6.x86_64,该表提供查询时区ID和跳秒之间的映射关系数据. |
| 27 | time_zone_leap_second | MySQL时区相关表,通过mysql_tzinfo_to_sql加载相关时区信息就会有数据,该表提供查询跳秒机器修正值信息. |
| 28 | time_zone_name |MySQL时区相关表,通过mysql_tzinfo_to_sql加载相关时区信息就会有数据,该表提供查询时区的名称列表和时区ID的映射关系. |
| 29 | time_zone_transition | MySQL时区相关表,通过mysql_tzinfo_to_sql加载相关时区信息就会有数据,该表提供查询时区的跳秒数据. |
| 30 | time_zone_transition_type | MySQL时区相关表,通过mysql_tzinfo_to_sql加载相关时区信息就会有数据,该表提供查询具体的跳秒信息以及与时区的对应数据.|
| 31 | user | MySQL用户表,创建的所有的用户都在这个表中. |

###### &nbsp;
[| Serial Number | Table Name | Meaning |]:#
[| :--- | :---  | :---  |]:#
[| 1 | columns_priv | The permission table of the columns in the table, when executing GRANT SELECT (HOST), ON `db1`.* TO 'test'@'%' identified by "123456"; there will be data.|]:#
[| 2 | db | The permission table of the database, when executing GRANT SELECT ON `db1`.* TO 'test'@'%' identified by "123456"; there will be data.|]:#
[| 3 | engine_cost | Cost model, engine_cost (IO cost), server_cost (CPU cost), mysql execution plan analysis three see explain, profiling, optimizer_trace. |]:#
[| 4 | event | Scheduled task table, when you pass create event event_now..., there will be data. Note that the scheduled task will only be in the event table if it is valid. By default, the event will be released after execution. If the event is executed immediately, the event will be automatically deleted after execution. |]:#
[| 5 | func | UDF: User-defined function, not Fuction, a higher-level function. |]:#
[| 6 | general_log | MySQL general query log, when log_output=TABLE and start general_log, all mysql execution logs will be written to this table. |]:#
[| 7 | gtid_executed | MySQL Global Transaction ID Persistent Table.  |]:#
[| 8 | help_category | Help information table, information about help topic categories, through help 'keyword'; to view help information, search through the name field in help_category, help_keyword, help_topic. |]:#
[| 9 | help_keyword | Help information sheet, keyword information related to help topics. |]:#
[| 10 | help_relation | Help information table, mapping between help keyword information and topic information. |]:#
[| 11 | help_topic | Help information sheet, details of help topics. |]:#
[| 12 | innodb_index_stats | Statistics table, which stores statistics of index dimensions, which is quite effective for analyzing SQL parsing and query optimization. |]:#
[| 13 | innodb_table_stats | Statistics table, which stores statistics for table dimensions. |]:#
[| 14 | ndb_binlog_index | DB Cluster replication is used, the demonstration is too troublesome, follow-up tests. |]:#
[| 15 | plugin | The plugin table that MySQL has loaded, the plugin loaded through my.cnf or INSTALL PLUGIN will have a record in this table. |]:#
[| 16 | proc | MySQL stored procedures and custom Function record tables, when PROCEDURE or FUNCTION are created, there will be records. The permissions of Mysq.5.7.19 functions and stored procedures are much stronger than before, and the syntax has also changed a little. |]:#
[| 17 | procs_priv | Stored procedure permission table, when executing GRANT EXECUTE ON PROCEDURE `db1`.procedure_name TO 'someuser'@'%';; there will be data. |]:#
[| 18 | proxies_priv | Proxy user permission table, which only appears in this table when the proxy user is authorized through grant proxy. |]:#
[| 19 | server_cost | Cost model, engine_cost (IO cost), server_cost (CPU cost), mysql execution plan analysis three see explain, profiling, optimizer_trace.|]:#
[| 20 | servers | The remote data user table, which is used when using the federated engine, can share remote data and local data, but the function is not complete. Use CREATE SERVER. . . To create, one thing to note is that the information must be written in the first time, otherwise it seems that it cannot be modified by commands in the future. Only by manually deleting the data in the servers table and restarting the service can the same server_name be created again, and then used again, FEDERATED The function of the engine is that the local machine can have the same data table as the remote, and the table name can be different, but the local machine does not save data, only the table definition, and the data is on the remote database. |]:#
[| 21 | slave_master_info | MySQL server main binlog information, when the binlog log is enabled and master_info_repository=TABLE, the relevant information will be recorded to this table.|]:#
[| 22 | slave_relay_log_info | The repeater binlog information of the MySQL server slave, when the slave opens the binlog log and relay_log_info_repository=TABLE., the relevant information will be recorded to this table. |]:#
[| 23 | slave_worker_info | MySQL slave service multi-threaded management table, when slave server slave_parallel_type='logical_clock'; (configured based on logical clock), the default is DATABASE (that is, a library has only one thread to pull dump log information from the master server), and global variables slave_parallel_workers>0 will show data in this table (mysql.slave_worker_info), the default performance_schema.replication_applier_status_by_worker has only one single-threaded data.|]:#
[| 24 | slow_log | MySQL slow query log, when log_output=TABLE and start slow_log, all mysql execution logs will be written to this table. |]:#
[| 25 | tables_priv | Data table permission table, when executing GRANT SELECT ON `db1`.tb1 TO 'test'@'%' identified by "123456"; there will be data .|]:#
[| 26 | time_zone | MySQL time zone related table, load the relevant time zone information through mysql_tzinfo_to_sql, there will be data, the Linux system time zone file is in the /usr/share/zoneinfo/ directory, my system was CentOS_6.6_Final_2.6.32-504.23.4.el6.x86_64, This table provides the mapping relationship data between query time zone IDs and jump seconds. |]:#
[| 27 | time_zone_leap_second | MySQL time zone related table, load related time zone information through mysql_tzinfo_to_sql, there will be data, this table provides query jump second machine correction value information. |]:#
[| 28 | time_zone_name | MySQL time zone related table, there will be data by loading related time zone information through mysql_tzinfo_to_sql, this table provides the mapping relationship between the name list of the query time zone and the time zone ID. |]:#
[| 29 | time_zone_transition | MySQL time zone related table, load related time zone information through mysql_tzinfo_to_sql, there will be data, this table provides the jump second data of query time zone. |]:#
[| 30 | time_zone_transition_type | MySQL time zone related table, load related time zone information through mysql_tzinfo_to_sql, there will be data, this table provides query specific jump second information and corresponding data with time zone.|]:#
[| 31 | user | MySQL user table, all users created are in this table. |]:#


# PERFORMANCE_SCHEMA 
[# PERFORMANCE_SCHEMA]:#

&nbsp;&nbsp;&nbsp;&nbsp;主要用于收集数据库服务器性能参数.并且库里表的存储引擎均为PERFORMANCE_SCHEMA,而用户是不能创建存储引擎为PERFORMANCE_SCHEMA的表.MySQL5.7默认是开启的,performance_schema 数据库表说明:

[&nbsp;&nbsp;&nbsp;&nbsp;Mainly used to collect database server performance parameters. And the storage engines of the tables in the library are all PERFORMANCE_SCHEMA, and users cannot create tables whose storage engine is PERFORMANCE_SCHEMA. MySQL5.7 is enabled by default, performance_schema database table description:]:#

| 序号 | 表名 | 含义 |
| :--- | :---  | :---  |
| 1 | accounts | 记录每个用户对应主机当前登录信息和总的登录统计数据(断开一次记录到总的连接数里面去),为USER和HOST为NULL表示内部用户.|
| 2 | cond_instances | 列出服务器正在执行时所有实例的所有条件(可见的).|
| 3 | events_stages_current | 阶段摘要表,包含每个线程的当前正在执行的阶段事件. |
| 4 | events_stages_history | 阶段摘要表,包含每个线程已结束的N(系统变量:performance_schema_events_stages_history_size)个阶段事件. |
| 5 | events_stages_history_long | 阶段摘要表,包含所有线程中全局结束的N(系统变量:performance_schema_events_stages_history_long_size)个阶段事件. |
| 6 | events_stages_summary_by_account_by_event_name | 阶段摘要表,按照账号和主机及事件阶段名称的语句进行聚合,聚集了各个维度的统计信息,以USER、HOST、EVENT_NAME(阶段)列为维度. |
| 7 | events_stages_summary_by_host_by_event_name | 阶段摘要表,按照主机和事件阶段名称的语句进行聚合,聚集了各个维度的统计信息,以HOST、EVENT_NAME(阶段)列为维度. |
| 8 | events_stages_summary_by_thread_by_event_name | 阶段摘要表,按照线程和事件阶段名称的语句进行聚合,聚集了各个维度的统计信息,以THREAD_ID、EVENT_NAME(阶段)列为维度. |
| 9 | events_stages_summary_by_user_by_event_name | 阶段摘要表,按照用户和事件阶段名称的语句进行聚合,聚集了各个维度的统计信息,以USER、EVENT_NAME(阶段)列为维度. |
| 10 | events_stages_summary_global_by_event_name | 阶段摘要表,按照事件阶段名称进行聚合(全局),聚集了各个维度的统计信息,以EVENT_NAME(阶段)列为维度. |
| 11 | events_statements_current | 语句摘要表,包含每个线程的当前正在执行的语句事件. |
| 12 | events_statements_history | 语句摘要表,包含每个线程已结束的N(系统变量:performance_schema_events_statements_history_size)个语句事件. |
| 13 | events_statements_history_long | 语句摘要表,包含所有线程中全局结束的N(系统变量:performance_schema_events_statements_history_long_size)个语句事件. |
| 14 | events_statements_summary_by_account_by_event_name | 语句摘要表,按照账号和主机及事件语句名称的语句进行聚合,聚集了各个维度的统计信息,以USER、HOST、EVENT_NAME(语句)列为维度. |
| 15 | events_statements_summary_by_digest | 语句摘要表,SQL维度的统计信息表,可以统计某类SQL语句在各个维度的统计信息(比如：执行次数,排序次数,使用临时表等),以SCHEMA_NAME(语句)、DIGEST、DIGEST_TEXT列为维度. |
| 16 | events_statements_summary_by_host_by_event_name | 语句摘要表,按照主机和事件语句名称的语句进行聚合,聚集了各个维度的统计信息,以HOST、EVENT_NAME(语句)列为维度. |
| 17 | events_statements_summary_by_program	 | 语句摘要表,按照线程和事件语句函数名称的语句进行聚合,聚集了各个维度的统计信息,以OBJECT_TYPE、OBJECT_SCHEMA、OBJECT_NAME列为维度. |
| 18 | events_statements_summary_by_thread_by_event_name | 语句摘要表,按照线程和事件语句名称的语句进行聚合,聚集了各个维度的统计信息,以THREAD_ID、EVENT_NAME(语句)列为维度. |
| 19 | events_statements_summary_by_user_by_event_name | 语句摘要表,按照用户和事件语句名称的语句进行聚合,聚集了各个维度的统计信息,以USER、EVENT_NAME(语句)列为维度. |
| 20 | events_statements_summary_global_by_event_name | 语句摘要表,按照事件语句名称进行聚合(全局),聚集了各个维度的统计信息,以EVENT_NAME(语句)列为维度. |
| 21 | events_transactions_current | 事务摘要表,包含每个线程的当前正在执行的事务事件. |
| 22 | events_transactions_history | 事务摘要表,包含每个线程已结束的N(系统变量:performance_schema_events_transactions_history_size)个事务事件. |
| 23 | events_transactions_history_long | 事务摘要表,包含所有线程中全局结束的N(系统变量:performance_schema_events_transactions_history_long_size)个事务事件. |
| 24 | events_transactions_summary_by_account_by_event_name | 事务摘要表,按照账号和主机及事件事务名称的语句进行聚合,聚集了各个维度的统计信息,以USER、HOST、EVENT_NAME(事务)列为维度. |
| 25 | events_transactions_summary_by_host_by_event_name | 事务摘要表,按照主机和事件事务名称的语句进行聚合,聚集了各个维度的统计信息,以HOST、EVENT_NAME(事务)列为维度. |
| 26 | events_transactions_summary_by_thread_by_event_name | 事务摘要表,按照线程和事件事务名称的语句进行聚合,聚集了各个维度的统计信息,以THREAD_ID、EVENT_NAME(事务)列为维度. |
| 27 | events_transactions_summary_by_user_by_event_name | 事务摘要表,按照用户和事件事务名称的语句进行聚合,聚集了各个维度的统计信息,以USER、EVENT_NAME(事务)列为维度. |
| 28 | events_transactions_summary_global_by_event_name | 事务摘要表,按照事件事物名称进行聚合(全局),聚集了各个维度的统计信息,以EVENT_NAME(事务)列为维度. |
| 29 | events_waits_current | 等待摘要表,包含每个线程的当前正在执行的等待事件. |
| 30 | events_waits_history | 等待摘要表,包含每个线程已结束的N(系统变量:performance_schema_events_waits_history_size)个等待事件. |
| 31 | events_waits_history_long | 等待摘要表,包含所有线程中全局结束的N(系统变量:performance_schema_events_waits_history_long_size)个等待事件. |
| 32 | events_waits_summary_by_account_by_event_name | 等待摘要表,按照账号和主机及事件等待名称的语句进行聚合,聚集了各个维度的统计信息,以USER、HOST、EVENT_NAME(等待)列为维度.  |
| 33 | events_waits_summary_by_host_by_event_name | 等待摘要表,按照主机和事件等待名称的语句进行聚合,聚集了各个维度的统计信息,以HOST、EVENT_NAME(等待)列为维度. |
| 34 | events_waits_summary_by_instance | 等待摘要表,按照线程和事件等待函数名称的语句进行聚合,聚集了各个维度的统计信息,以EVENT_NAME(等待)、OBJECT_INSTANCE_BEGIN列为维度. |
| 35 | events_waits_summary_by_thread_by_event_name | 等待摘要表,按照线程和事件等待名称的语句进行聚合,聚集了各个维度的统计信息,以THREAD_ID、EVENT_NAME(等待)列为维度. |
| 36 | events_waits_summary_by_user_by_event_name | 等待摘要表,按照用户和事件等待名称的语句进行聚合,聚集了各个维度的统计信息,以USER、EVENT_NAME(等待)列为维度. |
| 37 | events_waits_summary_global_by_event_name | 等待摘要表,按照事件等待名称进行聚合(全局),聚集了各个维度的统计信息,以EVENT_NAME(等待)列为维度. |
| 38 | file_instances | 列出服务器正在执行时所有实例的所有文件(打开的). |
| 39 | file_summary_by_event_name | 性能文件I/O摘要表,每行统计了给定事件名称的事件,以EVENT_NAME(事件文件)列为维度. |
| 40 | file_summary_by_instance | 性能文件I/O摘要表,每行统计了给定文件和事件名称的事件,以FILE_NAME(文件)、EVENT_NAME(事件文件)列为维度. |
| 41 | global_status | 全局系统状态表,就是平时我们用showglobalstatuslike'%参数名称%'显示的值都存在这张表中. |
| 42 |global_variables  | 全局系统变量表,就是平时我们用 `show global variables like '%parameter%'` 显示的值都存在这张表中. |
| 43 | host_cache | 主机缓存表统计数据,用户加快dns解析的. |
| 44 | hosts | 按主机维度统计当前连接数和总连接的数据. |
| 45 | memory_summary_by_account_by_event_name | 内存摘要表,按照账号和主机及事件内存对象变量名称的语句进行聚合,聚集了各个维度的统计信息,以USER、HOST、EVENT_NAME(内存对象变量)列为维度. |
| 46 | memory_summary_by_host_by_event_name | 内存摘要表,按照主机和事件内存对象变量名称的语句进行聚合,聚集了各个维度的统计信息,以HOST、EVENT_NAME(内存对象变量)列为维度.|
| 47 | memory_summary_by_thread_by_event_name | 内存摘要表,按照线程和事件内存对象变量名称的语句进行聚合,聚集了各个维度的统计信息,以THREAD_ID、EVENT_NAME(内存对象变量)列为维度. |
| 48 | memory_summary_by_user_by_event_name | 内存摘要表,按照用户和事件内存对象变量名称的语句进行聚合,聚集了各个维度的统计信息,以USER、EVENT_NAME(内存对象变量)列为维度. |
| 49 | memory_summary_global_by_event_name | 内存摘要表,按照事件内存对象变量名称进行聚合(全局),聚集了各个维度的统计信息,以EVENT_NAME(内存对象变量)列为维度. |
| 50 | metadata_locks | 使用用元数据锁定来管理对数据库对象的并发访问并确保数据一致性的记录锁表(发生时才有记录). |
| 51 | mutex_instances | 查看哪个其他线程当前拥有互斥锁. |
| 52 | objects_summary_global_by_type | 对象等待摘要表,用于汇总对象等待事件表. |
| 53 | performance_timers | 用来查询有哪些可用的事件计时器类型表. |
| 54 | prepared_statements_instances | 预处理语句实例和统计信息. |
| 55 | replication_applier_configuration | 查看各个channel是否配置了复制延迟 |
| 56 | replication_applier_status | 查看各个channel是否复制正常（service_state）以及事务重连的次数. |
| 57 | replication_applier_status_by_coordinator | 查看各个channel是否复制正常,以及复制错误的code、message和时间. |
| 58 | replication_applier_status_by_worker | 查看各个channel是否复制正常,以及并行复制work号,复制错误的code、SQL和时间. |
| 59 | replication_connection_configuration | 查看各个channel的连接配置信息：host、port、user、auto_position等. |
| 60 | replication_connection_status | 查看各个channel的连接信息. |
| 61 | replication_group_member_stats | 当GroupReplication全同步复制模式启用时才会有值,用户提供了与认证过程相关的组级别信息. |
| 62 | replication_group_members | 当GroupReplication全同步复制模式启用时才会有值,用于监视属于该组成员的不同服务器实例的状态. |
| 63 | rwlock_instances | 表列出了服务器所有rwlock(读写锁)实例相关记录. |
| 64 | session_account_connect_attrs | 仅包含当前会话以及与该会话帐户关联的其他会话的连接属性表. |
| 65 | session_connect_attrs | 所有会话连接属性表. |
| 66 | session_status | 当前会话级别的用户状态表,就是平时我们用 `show ('session') status like '%parameter%'` 显示的值都存在这张表中. |
| 67 | session_variables  | 当前会话级别的用户变量表,就是平时我们用 `show ('session') variables like '%parameter%'` 显示的值都存在这张表中. |
| 68 | setup_actors | 设置表提供有关当前检测的信息,如何初始化对新前台线程的监视,要更改表记录的数量控制,请在服务器启动时修改(系统变量:performance_schema_setup_actors_size). |
| 69 | setup_consumers | 设置表提供有关当前检测的信息,列出可以存储event信息且启用了哪些消费者的消费者类型. |
| 70 | setup_instruments | 设置表提供有关当前检测的信息,可以收集events的已检测objects的类. |
| 71 | setup_objects | 设置表提供有关当前检测的信息,控制是否监视特定的objects,如更改记录表的数量条件,请在服务器启动时修改(系统变量:performance_schema_setup_objects_size). |
| 72 | setup_timers | 设置表提供有关当前检测的信息,表显示当前选定的事件计时器. |
| 73 | socket_instances | 该表提供了与MySQL服务器的活动连接的实时快照. |
| 74 | socket_summary_by_event_name | 套接字摘要表,按照事件名称的事件的进行聚合,以EVENT_NAME(事件名称)列为维度,提供了一些附加信息,例如像socket操作以及网络传输和接收的字节数. |
| 75 | socket_summary_by_instance | 套接字摘要表,按照事件名称的事件的进行聚合,以EVENT_NAME(事件名称)、OBJECT_INSTANCE_BEGIN列为维度,提供了一些附加信息,例如像socket操作以及网络传输和接收的字节数. |
| 76 | status_by_account | 状态变量汇总表,具有USER、HOST、VARIABLE_NAME列,以按帐户汇总状态变量. |
| 77 | status_by_host | 状态变量汇总表,具有HOST、VARIABLE_NAME列,用于通过与客户端连接的主机来汇总状态变量. |
| 78 | status_by_thread | 状态变量汇总表,具有THREAD_ID、VARIABLE_NAME列,每个活动用户的汇总会话状态变量. |
| 79 | status_by_user | S状态变量汇总表,具有USER、VARIABLE_NAME列,用于按客户端用户名汇总状态变量. |
| 80 | table_handles | T表锁检测记录的内容,此信息显示服务器打开了哪个表句柄,如何锁定它们以及通过哪个会话,如更改记录表的数量条件,请在服务器启动时修改(系统变量:performance_schema_max_table_handles). |
| 81 | table_io_waits_summary_by_index_usage | 该表汇总了由wait/io/table/sql/handler生成的所有表索引I/O等待事件,按表索引进行分组. |
| 82 | table_io_waits_summary_by_table | 该表汇总了由wait/io/table/sql/handler生成的所有表I/O等待事件,按表分组. |
| 83 | table_lock_waits_summary_by_table | 该表汇总了由wait/lock/table/sql/handler生成的所有表锁等待事件,按表分组. |
| 84 | threads | 当前服务器所有的线程统计数据,每行包含有关线程的信息,并指示是否为该线程启用了监视和历史事件日志记录. |
| 85 | user_variables_by_thread | 改表记录了用户自定义的一些变量,这些是在特定会话中定义的变量,就是使用set@v1=10;,然后select@v1fromdual;的当次连接中. |
| 86 | users | 值记录登录用户的统计数据,不记录主机,主机在accounts表. |
| 87 | variables_by_thread | 每个活动会话的会话系统变量,会话变量表(session_variables,variables_by_thread)仅包含活动会话的信息,而不包含终止的会话的信息. |

###### &nbsp;
[| Serial Number | Table Name | Meaning |]:#
[| :--- | :---  | :---  |]:#
[| 1 | accounts | Record the current login information and total login statistics of each user corresponding to the host (disconnect one time and record it into the total number of connections), if USER and HOST are NULL, it means internal users.|]:#
[| 2 | cond_instances | List all conditions for all instances while the server is executing(visible).|]:#
[| 3 | events_stages_current | Stage summary table, containing currently executing stage events for each thread. |]:#
[| 4 | events_stages_history | Stage summary table, containing N (system variable: performance_schema_events_stages_history_size) stage events that each thread has ended. |]:#
[| 5 | events_stages_history_long | Stage summary table, containing N (system variable: performance_schema_events_stages_history_long_size) stage events that end globally in all threads. |]:#
[| 6 | events_stages_summary_by_account_by_event_name | The stage summary table is aggregated according to the statement of account, host and event stage name, and the statistical information of each dimension is aggregated, with USER, HOST, EVENT_NAME (stage) as the dimension. |]:#
[| 7 | events_stages_summary_by_host_by_event_name | The stage summary table is aggregated according to the statement of the host and event stage name, and the statistical information of each dimension is aggregated, with HOST and EVENT_NAME (stage) as the dimensions. |]:#
[| 8 | events_stages_summary_by_thread_by_event_name | The stage summary table, which is aggregated according to the statement of the thread and event stage name, aggregates the statistical information of each dimension, and uses THREAD_ID and EVENT_NAME (stage) as the dimensions. |]:#
[| 9 | events_stages_summary_by_user_by_event_name | The stage summary table, which is aggregated according to the statement of the user and event stage name, aggregates the statistical information of each dimension, and uses USER and EVENT_NAME (stage) as the dimensions. |]:#
[| 10 | events_stages_summary_global_by_event_name | The stage summary table, aggregated according to the event stage name (global), aggregated the statistical information of each dimension, with EVENT_NAME (stage) as the dimension. |]:#
[| 11 | events_statements_current | Statement summary table, containing currently executing statement events for each thread. |]:#
[| 12 | events_statements_history | Statement summary table, containing N (system variable: performance_schema_events_statements_history_size) statement events that each thread has ended. |]:#
[| 13 | events_statements_history_long | Statement summary table, containing N (system variable: performance_schema_events_statements_history_long_size) statement events that end globally in all threads. |]:#
[| 14 | events_statements_summary_by_account_by_event_name | Statement summary table, which is aggregated according to the statement of account, host and event statement name, and aggregates the statistical information of each dimension, with USER, HOST, EVENT_NAME (statement) as the dimension. |]:#
[| 15 | events_statements_summary_by_digest | Statement summary table, statistical information table of SQL dimension, can count the statistical information of a certain type of SQL statement in each dimension (for example: the number of executions, the number of times of sorting, the use of temporary tables, etc.), with SCHEMA_NAME (statement), DIGEST, DIGEST_TEXT as the dimensions. |]:#
[| 16 | events_statements_summary_by_host_by_event_name | Statement summary table, which is aggregated according to the statement of host and event statement name, and aggregates the statistical information of each dimension, with HOST, EVENT_NAME (statement) as the dimension. |]:#
[| 17 | events_statements_summary_by_program	 | The statement summary table is aggregated according to the statement of the thread and event statement function name, and the statistical information of each dimension is aggregated, with OBJECT_TYPE, OBJECT_SCHEMA, and OBJECT_NAME as the dimensions. |]:#
[| 18 | events_statements_summary_by_thread_by_event_name | Statement summary table, which is aggregated according to the statement of thread and event statement name, and aggregates the statistical information of each dimension, with THREAD_ID and EVENT_NAME (statement) as the dimensions. |]:#
[| 19 | events_statements_summary_by_user_by_event_name | Statement summary table, which is aggregated according to the statement of user and event statement name, and aggregates the statistical information of each dimension, with USER and EVENT_NAME (statement) as the dimensions. |]:#
[| 20 | events_statements_summary_global_by_event_name | Statement summary table, which is aggregated according to the event statement name (globally), aggregates the statistical information of each dimension, and uses EVENT_NAME (statement) as the dimension. |]:#
[| 21 | events_transactions_current | Transaction summary table, containing currently executing transaction events for each thread. |]:#
[| 22 | events_transactions_history | Transaction summary table, containing N (system variable: performance_schema_events_transactions_history_size) transaction events that each thread has ended. |]:#
[| 23 | events_transactions_history_long | Transaction summary table, containing N (system variable: performance_schema_events_transactions_history_long_size) transaction events that end globally in all threads. |]:#
[| 24 | events_transactions_summary_by_account_by_event_name | The transaction summary table is aggregated according to the statement of account, host and event transaction name, and aggregates the statistical information of each dimension, with USER, HOST, EVENT_NAME (transaction) as the dimension. |]:#
[| 25 | events_transactions_summary_by_host_by_event_name | The transaction summary table is aggregated according to the statement of the host and event transaction name, and the statistical information of each dimension is aggregated, with HOST and EVENT_NAME (transaction) as the dimensions. |]:#
[| 26 | events_transactions_summary_by_thread_by_event_name | The transaction summary table is aggregated according to the statement of the thread and event transaction name, and the statistical information of each dimension is aggregated, with THREAD_ID and EVENT_NAME (transaction) as the dimensions. |]:#
[| 27 | events_transactions_summary_by_user_by_event_name | The transaction summary table is aggregated according to the statement of the user and event transaction name, and the statistical information of each dimension is aggregated, with USER and EVENT_NAME (transaction) as the dimensions. |]:#
[| 28 | events_transactions_summary_global_by_event_name | The transaction summary table is aggregated according to the event name (global), and the statistical information of each dimension is aggregated, with EVENT_NAME (transaction) as the dimension. |]:#
[| 29 | events_waits_current | Wait summary table, containing currently executing wait events for each thread. |]:#
[| 30 | events_waits_history | Wait summary table, containing N (system variable: performance_schema_events_waits_history_size) wait events that each thread has ended. |]:#
[| 31 | events_waits_history_long | Wait summary table, containing N (system variable: performance_schema_events_waits_history_long_size) wait events that end globally in all threads. |]:#
[| 32 | events_waits_summary_by_account_by_event_name | The waiting summary table is aggregated according to the statement of account, host and event waiting name, and the statistical information of each dimension is aggregated, and USER, HOST, EVENT_NAME (waiting) are listed as dimensions.  |]:#
[| 33 | events_waits_summary_by_host_by_event_name | The waiting summary table is aggregated according to the statement of the host and the event waiting name, and the statistical information of each dimension is aggregated, and the HOST and EVENT_NAME (waiting) are listed as the dimensions. |]:#
[| 34 | events_waits_summary_by_instance | The waiting summary table is aggregated according to the statement of thread and event waiting function name, and the statistical information of each dimension is aggregated, and EVENT_NAME (waiting) and OBJECT_INSTANCE_BEGIN are listed as dimensions. |]:#
[| 35 | events_waits_summary_by_thread_by_event_name | The waiting summary table is aggregated according to the statement of the thread and event waiting name, and the statistical information of each dimension is aggregated, with THREAD_ID and EVENT_NAME (waiting) as the dimensions. |]:#
[| 36 | events_waits_summary_by_user_by_event_name | The waiting summary table is aggregated according to the statement of the user and the event waiting name, and the statistical information of each dimension is aggregated, with USER and EVENT_NAME (waiting) as the dimensions. |]:#
[| 37 | events_waits_summary_global_by_event_name | The waiting summary table is aggregated according to the event waiting name (global), and the statistical information of each dimension is aggregated, with EVENT_NAME (waiting) as the dimension. |]:#
[| 38 | file_instances | List all files (open) of all instances while the server is executing. |]:#
[| 39 | file_summary_by_event_name | Performance file I/O summary table, each row counts events for a given event name, with EVENT_NAME (event file) as the dimension. |]:#
[| 40 | file_summary_by_instance | Performance file I/O summary table, each row counts events for a given file and event name, with FILE_NAME (file) and EVENT_NAME (event file) as dimensions. |]:#
[| 41 | global_status | The global system status table, that is, the values ​​we usually display with showglobalstatuslike'% parameter name%' are stored in this table. |]:#
[| 42 |global_variables  | The global system variable table, that is, the values ​​that we usually display with `show global variables like '%parameter%'` are stored in this table. |]:#
[| 43 | host_cache | Host cache table statistics, users speed up dns resolution. |]:#
[| 44 | hosts | Statistics on the number of current connections and total connections by host dimension. |]:#
[| 45 | memory_summary_by_account_by_event_name | The memory summary table is aggregated according to the statement of account, host and event memory object variable name, and the statistical information of each dimension is aggregated, and USER, HOST, EVENT_NAME (memory object variable) is listed as the dimension. |]:#
[| 46 | memory_summary_by_host_by_event_name | The memory summary table is aggregated according to the statement of the host and event memory object variable names, and the statistical information of each dimension is aggregated, with HOST and EVENT_NAME (memory object variable) as the dimensions.|]:#
[| 47 | memory_summary_by_thread_by_event_name | The memory summary table is aggregated according to the statement of thread and event memory object variable name, and the statistical information of each dimension is aggregated, with THREAD_ID and EVENT_NAME (memory object variable) as the dimensions. |]:#
[| 48 | memory_summary_by_user_by_event_name | The memory summary table is aggregated according to the statement of the user and event memory object variable names, and the statistical information of each dimension is aggregated, with USER and EVENT_NAME (memory object variable) as the dimensions. |]:#
[| 49 | memory_summary_global_by_event_name | The memory summary table is aggregated according to the event memory object variable name (global), and the statistical information of each dimension is aggregated, and EVENT_NAME (memory object variable) is listed as the dimension. |]:#
[| 50 | metadata_locks | Use a record-locked table that uses metadata locking to manage concurrent access to database objects and ensure data consistency (records only happen when they happen). |]:#
[| 51 | mutex_instances | See which other thread currently owns the mutex. |]:#
[| 52 | objects_summary_global_by_type | Object wait summary table, which is used to summarize the object wait event table. |]:#
[| 53 | performance_timers | Used to query the table of available event timer types. |]:#
[| 54 | prepared_statements_instances | Prepared Statement Examples and Statistics. |]:#
[| 55 | replication_applier_configuration | Check whether each channel is configured with replication delay. |]:#
[| 56 | replication_applier_status | Check whether each channel is replicated normally (service_state) and the number of transaction reconnections. |]:#
[| 57 | replication_applier_status_by_coordinator | Check whether each channel is copied normally, and copy the wrong code, message and time. |]:#
[| 58 | replication_applier_status_by_worker | Check whether the replication of each channel is normal, and copy the work number in parallel, copy the wrong code, SQL and time. |]:#
[| 59 | replication_connection_configuration | View the connection configuration information of each channel: host, port, user, auto_position, etc. |]:#
[| 60 | replication_connection_status | View the connection information of each channel. |]:#
[| 61 | replication_group_member_stats | Only has value when GroupReplication full synchronous replication mode is enabled, the user provides group level information related to the authentication process. |]:#
[| 62 | replication_group_members | Only has value when GroupReplication full synchronous replication mode is enabled, which is used to monitor the status of different server instances that are members of the group. |]:#
[| 63 | rwlock_instances | The table lists all rwlock (read-write lock) instances related records of the server. |]:#
[| 64 | session_account_connect_attrs | A connection properties table containing only the current session and other sessions associated with this session account. |]:#
[| 65 | session_connect_attrs | All session connection properties table. |]:#
[| 66 | session_status | The user status table at the current session level, that is, the values ​​that we usually use `show ('session') status like '%parameter%'` to display are stored in this table. |]:#
[| 67 | session_variables  | The user variable table of the current session level, that is, the values ​​that we usually display with `show ('session') variables like '%parameter%'` are stored in this table. |]:#
[| 68 | setup_actors | The settings table provides information about the current instrumentation, how to initialize monitoring for new foreground threads, to change the number control of table records, modify at server startup (system variable: performance_schema_setup_actors_size). |]:#
[| 69 | setup_consumers | The settings table provides information about the current instrumentation, listing the consumer types that can store event information and which consumers are enabled. |]:#
[| 70 | setup_instruments | The settings table provides information about the current instrumentation, classes of instrumented objects that can collect events. |]:#
[| 71 | setup_objects | The settings table provides information about the current detection, and controls whether to monitor specific objects, such as changing the number of records table conditions, please modify at server startup (system variable: performance_schema_setup_objects_size). |]:#
[| 72 | setup_timers | The settings table provides information about the current detection, the table shows the currently selected event timer. |]:#
[| 73 | socket_instances | This table provides a real-time snapshot of active connections to the MySQL server. |]:#
[| 74 | socket_summary_by_event_name | Socket summary table, aggregates events by event name, with EVENT_NAME (event name) as the column dimension, provides some additional information, such as socket operations and the number of bytes transmitted and received by the network. |]:#
[| 75 | socket_summary_by_instance | Socket summary table, aggregates events by event name, with EVENT_NAME (event name), OBJECT_INSTANCE_BEGIN columns as dimensions, provides some additional information, such as socket operations and the number of bytes transmitted and received by the network. |]:#
[| 76 | status_by_account | Status variable summary table with USER, HOST, VARIABLE_NAME columns to summarize status variables by account. |]:#
[| 77 | status_by_host | State variable summary table with HOST, VARIABLE_NAME columns for summarizing state variables by the host connected to the client. |]:#
[| 78 | status_by_thread |State variable summary table with THREAD_ID, VARIABLE_NAME columns, summary session state variables for each active user. |]:#
[| 79 | status_by_user | Status variable summary table with USER, VARIABLE_NAME columns to summarize status variables by client username. |]:#
[| 80 | table_handles | Table locks detect the contents of records, this information shows which table handles are opened by the server, how they are locked and by which session, such as changing the number of records table conditions, please modify at server startup (system variable: performance_schema_max_table_handles). |]:#
[| 81 | table_io_waits_summary_by_index_usage | This table summarizes all table index I/O wait events generated by wait/io/table/sql/handler, grouped by table index. |]:#
[| 82 | table_io_waits_summary_by_table | This table summarizes all table I/O wait events generated by wait/io/table/sql/handler, grouped by table. |]:#
[| 83 | table_lock_waits_summary_by_table | This table summarizes all table lock wait events generated by wait/lock/table/sql/handler, grouped by table. |]:#
[| 84 | threads | All threads statistics for the current server, each row contains information about the thread and indicates whether monitoring and historical event logging are enabled for that thread. |]:#
[| 85 | user_variables_by_thread | Change the table to record some user-defined variables, these are variables defined in a specific session, that is, use set@v1=10;, and then select@v1fromdual; in the current connection. |]:#
[| 86 | users | The value records the statistics of the logged in user, does not record the host, the host is in the accounts table. |]:#
[| 87 | variables_by_thread | Session system variables for each active session, the session variable table (session_variables, variables_by_thread) only contains information for active sessions, not for terminated sessions. |]:#

# SYS 
[# SYS]:#

&nbsp;&nbsp;&nbsp;&nbsp;Sys库所有的数据源来自：performance_schema.目标是把performance_schema的把复杂度降低,让DBA能更好的阅读这个库里的内容.让DBA更快的了解DB的运行情况,sys库下全部是视图,我就不多此一举了,MySQL Doc有统一的说明:

[&nbsp;&nbsp;&nbsp;&nbsp;All data sources of the Sys library come from: performance_schema. The goal is to reduce the complexity of performance_schema so that DBAs can better read the contents of this library. Let the DBA understand the operation of the DB faster. All the views in the sys library are views, so I will not do anything unnecessary. The MySQL Doc has a unified description:]:#

# sys库表说明
[# sys library table description]:#
[sys Schema Object Index](https://dev.mysql.com/doc/refman/5.7/en/sys-schema-object-index.html)


# TEST 
[# TEST ]:# 

&nbsp;&nbsp;&nbsp;&nbsp这个是安装时候创建的一个测试数据库,和它的名字一样,是一个完全的空数据库,没有任何表,可以删除(就不说明了).

[&nbsp;&nbsp;&nbsp;&nbsp;The is a test database created during installation. Like its name, it is a completely empty database without any tables and can be deleted(Not explained).]:#

<br>

[back](./)