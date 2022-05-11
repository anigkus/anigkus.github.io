<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >Do you really understand MySQL's own libraries and tables?</h1><br/>

<center>
<img src="assets/images/do-you-really-understand-mysql-is-own-libraries-and-tables/figure-1.png" alt="Do you really understand MySQL's own libraries and tables" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; The history of MySQL can be traced back to 1979, and it has been continuously developed and iterated for decades. From the release of MySQL 1.0 in 1996 to the later 3.x, 4.x, 5.x and the latest 8.x, new functions have been continuously developed And improve the problem, it is liked and used by the public, but the built-in database and tables of the database have also changed a lot. I will use 5.7.19 (but the difference between the major version is still very large, and the minor version is basically the same) Let's find out what are the existing self-contained libraries and tables, what data are stored, what are they used for, and what convenience they can bring to our daily database maintenance or development work, let's take a look.<br/>
> <br/>

## INFORMATION_SCHEMA 
&nbsp;&nbsp;&nbsp;&nbsp;The information_schema database comes with MySQL and provides a way to access database metadata. What is metadata? Metadata is data about data, such as database or table names, data types of columns, or access rights, etc. Other terms sometimes used to describe this information include "data dictionary" and "system catalog". In MySQL, information_schema is regarded as a database, to be precise, an information database. It holds information about all other databases maintained by the MySQL server. Such as database name, database table, table column data type and access rights, etc. In INFORMATION_SCHEMA, there are several read-only tables. They are actually views, not base tables, so you won't be able to see any files related to them. Information_schema database table description:

| Serial Number | Table Name | Meaning |
| :--- | :----: | ---: |
| 1 | CHARACTER_SETS | The default correspondence table between characters and proofreading rules cannot be added, updated or deleted.|
| 2 | COLLATIONS | All correspondence tables between characters and collation rules, and there are information such as collation ID, default identification, whether to compile into the server (these are all MySQL extension information for the time being, and are useless now). This table cannot be added, updated, or deleted.|
| 3 | COLLATION_CHARACTER_SET_APPLICABILITY | All correspondence between characters and proofreading rules |
| 4 | COLUMNS | Information about the columns in all data tables of the current MySQL server, such as type, whether it is empty, length, character encoding, proofreading encoding and many other useful information; |
| 5 | COLUMN_PRIVILEGES | Almost the same as mysql.columns_priv, there is a column IS_COLUMN_PRIVILEGES: used to indicate whether the user can continue to authorize others. |
| 6 | ENGINES | All storage engines currently served by MySQL, engine status(supported, default, disabled, and not supportd, supported but disabled), which supports transactions,XA transactions, partica transactions,etc. |
| 7 | EVENTS | The MySQL timed task tanle is almost the same as mysql.event, and has similar fuctions, but there will be a little difference in data performance, mainly in terms of time. |
| 8 | FILES | MySQL file table,mainly store innodb and NDB file information. |
| 9 | GLOBAL_STATUS | The global system status table, a table reserved by MySQL for compatibility with mysql5.6, has the same effect as the SHOW GLOBAL STATUS command and SELECT * FROM performance_schema.global_status (specific to 5.7), but it may be removed later, more compatible tables. |
| 10 | GLOBAL_VARIABLES | The global system variable table, a table reserved by MySQL for compatibility with mysql5.6, has the same effect as the SHOW GLOBAL VARIABLES command and SELECT * FROM performance_schema.global_variables (specific to 5.7), but it may be removed later. |
| 11 | KEY_COLUMN_USAGE | Primary key constraint record table for all MySQL libraries. |
| 12 | OPTIMIZER_TRACE | If you want to see the entire execution plan and how to choose between multiple indexing schemes? This function is supported in MySQL 5.6, optimizer_trace, which is closed by default, session level, and three carriages with the MySQL execution plan of explain, profile and optimizer_trace. |
| 13 | PARAMETERS | MySQL stored procedure or function parameter record table, more detailed than mysql.proc. |
| 14 | PARTITIONS | MySQL table space record table, the default innodb is shared storage, that is, the stored data information and index information are in the same *.ibd, and the Myisam storage engine, it uses an independent table space by default, but it can be changed by parameters. |
| 15 | PLUGINS | The MySQL plugin table is the same as the command SHOW PLUGINS, but it is more detailed than this command, such as the company, version, author, loading method, etc. of the plugin. |
| 16 | PROCESSLIST | The MySQL thread running table is the same as the SHOW FULL PROCESSLIST command, but different from the SHOW PROCESSLIST command. It seems to be related to permissions. The performance_schema.threads table can also view threads. The official website says that PROCESSLIST has a negative performance impact, but threads do not Yes, I haven't tested the obvious difference for the time being. |
| 17 | PROFILING | QUERY_ID is equivalent to show profile for query QUERY_ID; this, show profile only displays the last time, through SET profiling = 1; turn on the switch, Show profiles displays all records of the current session. |
| 18 | REFERENTIAL_CONSTRAINTS | MySQL reference constraint table, but through ADD CONSTRAINT there will be records in this table. |
| 19 | ROUTINES | MySQL contains stored procedures and function tables, excluding UDFs function records. |
| 20 | SCHEMATA | All MySQL library-defined data tables, including library names, characters, collation rules, etc. |
| 21 | SCHEMA_PRIVILEGES | MySQL database level privilege table, the same as mysql.db. |
| 22 | SESSION_STATUS | MySQL 5.7 has the same effect as show status in order to be compatible with the reserved table of 5.6. |
| 23 | SESSION_VARIABLES | MySQL 5.7 has the same effect as show VARIABLES in order to be compatible with the reserved table of 5.6. |
| 24 | STATISTICS | All index records on the MySQL server table, query a single index situation: SHOW INDEX FROM TABLE_NAME; and SHOW INDEX FROM TABLE_NAME FROM DB_NAME; |
| 25 | TABLES | All tables on the MySQL server record the table, which may be out of sync with the current table contents, but can be updated by running ANALYZE. |
| 26 | TABLESPACES | MySQL table space record table, this table does not provide information about InnoDB table space, you need to query the INFORMATION_SCHEMA.FILES table. |
| 27 | TABLE_CONSTRAINTS | MySQL constraint tables, including unique constraints, primary key constraints, foreign key constraints related record tables. |
| 28 | TABLE_PRIVILEGES | MySQL table-level privilege table, almost the same as mysql.tables_priv. |
| 29 | TRIGGERS | MySQL trigger archive table. |
| 30 | USER_PRIVILEGES | MySQL user permission table, almost the same as mysql.user. |
| 31 | VIEWS | MySQL view table. |
| 32 | INNODB_LOCKS | MySQL innodb transaction lock waiting table, used to analyze lock waiting, (transaction related table). |
| 33 | INNODB_TRX | All open transaction tables in MySQL innodb disappear automatically when closed, (transaction-related tables). |
| 34 | INNODB_SYS_DATAFILES | MySQL innodb data file table, only contains the disk path data of *.idb files, the default is relative to the mysql data directory, you can specify an absolute path when creating. |
| 35 | INNODB_FT_CONFIG | MySQL full-text index configuration table, only when the full-text index is defined on the column of the field type, when set global innodb_ft_aux_table='DB_NAME/TABLE_NAME'; there will be records, how to use it in the follow-up research. |
| 36 | INNODB_SYS_VIRTUAL | MySQL virtual field record table, when it is listed as a virtual column, there will be records. |
| 37 | INNODB_CMP | Compression related tables, compression status, mysql table data compression, innodb table compression, how to set the compression of mysql innodb table, InnoDB record compression and usage analysis, innodb_cmp saves historical summary data, and innodb_cmp_reset records a more real-time statistics value. |
| 38 | INNODB_FT_BEING_DELETED | MySQL full-text index ready to delete record table. |
| 39 | INNODB_CMP_RESET | Compression related tables, compression status, mysql table data compression, innodb table compression, how to set the compression of mysql innodb table, InnoDB record compression and usage analysis. |
| 40 | INNODB_CMP_PER_INDEX | Compress related tables, which provide the compression of each table and index, innodb compression. |
| 41 | INNODB_CMPMEM_RESET | Compress related tables, cache compressed data mysql table data compression, innodb table compression, how to set the compression of mysql innodb table, InnoDB record compression and usage analysis. |
| 42 |INNODB_FT_DELETED  | MySQL full-text index deleted records table. |
| 43 | INNODB_BUFFER_PAGE_LRU | Information of each page in the buffer pool, this table to observe the specific information of each page in the LRU (Least Recently Used Page Replacement Algorithm) list. |
| 44 | INNODB_LOCK_WAITS | Transaction-related tables, lock waiting tables. |
| 45 | INNODB_TEMP_TABLE_INFO | This table represents information about temporary tables. |
| 46 | INNODB_SYS_INDEXES | This table represents index metadata in all tables of the system. |
| 47 | INNODB_SYS_TABLES | This table represents metadata for all tables in the system. |
| 48 | INNODB_SYS_FIELDS | This table represents the mapping logic of fields and index positions in all tables of the system |
| 49 | INNODB_CMP_PER_INDEX_RESET | (Compression related tables) innodb_cmp_per_index_reset is similar to innodb_cmp_per_index, the difference is that the previous statistics are reset each time a query is made. |
| 50 | INNODB_BUFFER_PAGE | Information about each page in the buffer pool. |
| 51 | INNODB_FT_DEFAULT_STOPWORD | MySQL full-text index disabled keyword table. |
| 52 | INNODB_FT_INDEX_TABLE | MySQL full-text index table. |
| 53 | INNODB_FT_INDEX_CACHE | MySQL full-text index cache table. |
| 54 | INNODB_SYS_TABLESPACES | MySQL tablespace usage table. |
| 55 | INNODB_METRICS | MySQL monitoring table, introduced in 5.6, this table is used to monitor whether InnoDB is running normally. |
| 56 | INNODB_SYS_FOREIGN_COLS | Mysql foreign key reference relational table is somewhat similar to REFERENTIAL_CONSTRAINTS, but REFERENTIAL_CONSTRAINTS is more detailed, and REFERENTIAL_CONSTRAINTS cannot indicate location, and this table can. |
| 57 | INNODB_CMPMEM | (compress related tables), cache compressed data mysql table data compression, innodb table compression, how to set the compression of mysql innodb table, InnoDB record compression and usage analysis.  |
| 58 | INNODB_BUFFER_POOL_STATS | The status information in the buffer pool is similar to some data in the SHOW ENGINE INNODB STATUS command, but this table provides more information). |
| 59 | INNODB_SYS_COLUMNS | Corresponding table ID of all columns in the MySQL database (the relationship between the access ID and table name of this table INNODB_SYS_TABLES), column name, position in the table, column type, length and other related information. |
| 60 | INNODB_SYS_FOREIGN | All foreign key column relationship information tables in the library are somewhat similar to INNODB_SYS_FOREIGN_COLS and REFERENTIAL_CONSTRAINTS. |
| 61 | INNODB_SYS_TABLESTATS | The number of all data rows in all tables in the MySQL database will be updated after DELETE\UPDATE, and may be inaccurate if uncommitted transactions are being inserted or deleted from the table. What is the difference between NUN_ROWS in this table and TABLE_ROWS in TABLES, because sometimes I find that the data in TABLE_ROWS does not match the actual data in the table, is it because the transaction is not committed? |

## MYSQL 
&nbsp;&nbsp;&nbsp;&nbsp;The core database of mysql, similar to the master table in sql server, is mainly responsible for storing the control and management information that mysql needs to use, such as database users, permission settings, keywords, etc. mysql database table description: 

| Serial Number | Table Name | Meaning |
| :--- | :----: | ---: |
| 1 | columns_priv | The permission table of the columns in the table, when executing GRANT SELECT (HOST), ON `db1`.* TO 'test'@'%' identified by "123456"; there will be data.|
| 2 | db | The permission table of the database, when executing GRANT SELECT ON `db1`.* TO 'test'@'%' identified by "123456"; there will be data.|
| 3 | engine_cost | Cost model, engine_cost (IO cost), server_cost (CPU cost), mysql execution plan analysis three see explain, profiling, optimizer_trace. |
| 4 | event | Scheduled task table, when you pass create event event_now..., there will be data. Note that the scheduled task will only be in the event table if it is valid. By default, the event will be released after execution. If the event is executed immediately, the event will be automatically deleted after execution. |
| 5 | func | UDF: User-defined function, not Fuction, a higher-level function. |
| 6 | general_log | MySQL general query log, when log_output=TABLE and start general_log, all mysql execution logs will be written to this table. |
| 7 | gtid_executed | MySQL Global Transaction ID Persistent Table.  |
| 8 | help_category | Help information table, information about help topic categories, through help 'keyword'; to view help information, search through the name field in help_category, help_keyword, help_topic. |
| 9 | help_keyword | Help information sheet, keyword information related to help topics. |
| 10 | help_relation | Help information table, mapping between help keyword information and topic information |
| 11 | help_topic | Help information sheet, details of help topics. |
| 12 | innodb_index_stats | Statistics table, which stores statistics of index dimensions, which is quite effective for analyzing SQL parsing and query optimization. |
| 13 | innodb_table_stats | Statistics table, which stores statistics for table dimensions. |
| 14 | ndb_binlog_index | DB Cluster replication is used, the demonstration is too troublesome, follow-up tests. |
| 15 | plugin | The plugin table that MySQL has loaded, the plugin loaded through my.cnf or INSTALL PLUGIN will have a record in this table. |
| 16 | proc | MySQL stored procedures and custom Function record tables, when PROCEDURE or FUNCTION are created, there will be records. The permissions of Mysq.5.7.19 functions and stored procedures are much stronger than before, and the syntax has also changed a little. |
| 17 | procs_priv | Stored procedure permission table, when executing GRANT EXECUTE ON PROCEDURE `db1`.procedure_name TO 'someuser'@'%';; there will be data. |
| 18 | proxies_priv | Proxy user permission table, which only appears in this table when the proxy user is authorized through grant proxy
Translated with Google (Chinese → English)  |
| 19 | server_cost | Cost model, engine_cost (IO cost), server_cost (CPU cost), mysql execution plan analysis three see explain, profiling, optimizer_trace
Translated with Google (Chinese → English) |
| 20 | servers | The remote data user table, which is used when using the federated engine, can share remote data and local data, but the function is not complete. Use CREATE SERVER. . . To create, one thing to note is that the information must be written in the first time, otherwise it seems that it cannot be modified by commands in the future. Only by manually deleting the data in the servers table and restarting the service can the same server_name be created again, and then used again, FEDERATED The function of the engine is that the local machine can have the same data table as the remote, and the table name can be different, but the local machine does not save data, only the table definition, and the data is on the remote database. |
| 21 | slave_master_info | MySQL server main binlog information, when the binlog log is enabled and master_info_repository=TABLE, the relevant information will be recorded to this table.|
| 22 | slave_relay_log_info | The repeater binlog information of the MySQL server slave, when the slave opens the binlog log and relay_log_info_repository=TABLE., the relevant information will be recorded to this table. |
| 23 | slave_worker_info | MySQL slave service multi-threaded management table, when slave server slave_parallel_type='logical_clock'; (configured based on logical clock), the default is DATABASE (that is, a library has only one thread to pull dump log information from the master server), and global variables slave_parallel_workers>0 will show data in this table (mysql.slave_worker_info), the default performance_schema.replication_applier_status_by_worker has only one single-threaded data.|
| 24 | slow_log | MySQL slow query log, when log_output=TABLE and start slow_log, all mysql execution logs will be written to this table. |
| 25 | tables_priv | Data table permission table, when executing GRANT SELECT ON `db1`.tb1 TO 'test'@'%' identified by "123456"; there will be data .|
| 26 | time_zone | MySQL time zone related table, load the relevant time zone information through mysql_tzinfo_to_sql, there will be data, the Linux system time zone file is in the /usr/share/zoneinfo/ directory, my system was CentOS_6.6_Final_2.6.32-504.23.4.el6.x86_64, This table provides the mapping relationship data between query time zone IDs and jump seconds. |
| 27 | time_zone_leap_second | MySQL time zone related table, load related time zone information through mysql_tzinfo_to_sql, there will be data, this table provides query jump second machine correction value information. |
| 28 | time_zone_name | MySQL time zone related table, there will be data by loading related time zone information through mysql_tzinfo_to_sql, this table provides the mapping relationship between the name list of the query time zone and the time zone ID. |
| 29 | time_zone_transition | MySQL time zone related table, load related time zone information through mysql_tzinfo_to_sql, there will be data, this table provides the jump second data of query time zone. |
| 30 | time_zone_transition_type | MySQL time zone related table, load related time zone information through mysql_tzinfo_to_sql, there will be data, this table provides query specific jump second information and corresponding data with time zone.|
| 31 | user | MySQL user table, all users created are in this table. |


## PERFORMANCE_SCHEMA 
&nbsp;&nbsp;&nbsp;&nbsp;Mainly used to collect database server performance parameters. And the storage engines of the tables in the library are all PERFORMANCE_SCHEMA, and users cannot create tables whose storage engine is PERFORMANCE_SCHEMA. MySQL5.7 is enabled by default, performance_schema database table description: 

| Serial Number | Table Name | Meaning |
| :--- | :----: | ---: |
| 1 | accounts | Record the current login information and total login statistics of each user corresponding to the host (disconnect one time and record it into the total number of connections), if USER and HOST are NULL, it means internal users.|
| 2 | cond_instances | List all conditions for all instances while the server is executing(visible).|
| 3 | events_stages_current | Stage summary table, containing currently executing stage events for each thread. |
| 4 | events_stages_history | Stage summary table, containing N (system variable: performance_schema_events_stages_history_size) stage events that each thread has ended. |
| 5 | events_stages_history_long | Stage summary table, containing N (system variable: performance_schema_events_stages_history_long_size) stage events that end globally in all threads. |
| 6 | events_stages_summary_by_account_by_event_name | The stage summary table is aggregated according to the statement of account, host and event stage name, and the statistical information of each dimension is aggregated, with USER, HOST, EVENT_NAME (stage) as the dimension. |
| 7 | events_stages_summary_by_host_by_event_name | The stage summary table is aggregated according to the statement of the host and event stage name, and the statistical information of each dimension is aggregated, with HOST and EVENT_NAME (stage) as the dimensions. |
| 8 | events_stages_summary_by_thread_by_event_name | The stage summary table, which is aggregated according to the statement of the thread and event stage name, aggregates the statistical information of each dimension, and uses THREAD_ID and EVENT_NAME (stage) as the dimensions. |
| 9 | events_stages_summary_by_user_by_event_name | The stage summary table, which is aggregated according to the statement of the user and event stage name, aggregates the statistical information of each dimension, and uses USER and EVENT_NAME (stage) as the dimensions. |
| 10 | events_stages_summary_global_by_event_name | The stage summary table, aggregated according to the event stage name (global), aggregated the statistical information of each dimension, with EVENT_NAME (stage) as the dimension. |
| 11 | events_statements_current | Statement summary table, containing currently executing statement events for each thread. |
| 12 | events_statements_history | Statement summary table, containing N (system variable: performance_schema_events_statements_history_size) statement events that each thread has ended. |
| 13 | events_statements_history_long | Statement summary table, containing N (system variable: performance_schema_events_statements_history_long_size) statement events that end globally in all threads. |
| 14 | events_statements_summary_by_account_by_event_name | Statement summary table, which is aggregated according to the statement of account, host and event statement name, and aggregates the statistical information of each dimension, with USER, HOST, EVENT_NAME (statement) as the dimension. |
| 15 | events_statements_summary_by_digest | Statement summary table, statistical information table of SQL dimension, can count the statistical information of a certain type of SQL statement in each dimension (for example: the number of executions, the number of times of sorting, the use of temporary tables, etc.), with SCHEMA_NAME (statement), DIGEST, DIGEST_TEXT as the dimensions. |
| 16 | events_statements_summary_by_host_by_event_name | Statement summary table, which is aggregated according to the statement of host and event statement name, and aggregates the statistical information of each dimension, with HOST, EVENT_NAME (statement) as the dimension. |
| 17 | events_statements_summary_by_program	 | The statement summary table is aggregated according to the statement of the thread and event statement function name, and the statistical information of each dimension is aggregated, with OBJECT_TYPE, OBJECT_SCHEMA, and OBJECT_NAME as the dimensions. |
| 18 | events_statements_summary_by_thread_by_event_name | Statement summary table, which is aggregated according to the statement of thread and event statement name, and aggregates the statistical information of each dimension, with THREAD_ID and EVENT_NAME (statement) as the dimensions. |
| 19 | events_statements_summary_by_user_by_event_name | Statement summary table, which is aggregated according to the statement of user and event statement name, and aggregates the statistical information of each dimension, with USER and EVENT_NAME (statement) as the dimensions. |
| 20 | events_statements_summary_global_by_event_name | Statement summary table, which is aggregated according to the event statement name (globally), aggregates the statistical information of each dimension, and uses EVENT_NAME (statement) as the dimension. |
| 21 | events_transactions_current | Transaction summary table, containing currently executing transaction events for each thread. |
| 22 | events_transactions_history | Transaction summary table, containing N (system variable: performance_schema_events_transactions_history_size) transaction events that each thread has ended. |
| 23 | events_transactions_history_long | Transaction summary table, containing N (system variable: performance_schema_events_transactions_history_long_size) transaction events that end globally in all threads. |
| 24 | events_transactions_summary_by_account_by_event_name | The transaction summary table is aggregated according to the statement of account, host and event transaction name, and aggregates the statistical information of each dimension, with USER, HOST, EVENT_NAME (transaction) as the dimension. |
| 25 | events_transactions_summary_by_host_by_event_name | The transaction summary table is aggregated according to the statement of the host and event transaction name, and the statistical information of each dimension is aggregated, with HOST and EVENT_NAME (transaction) as the dimensions. |
| 26 | events_transactions_summary_by_thread_by_event_name | The transaction summary table is aggregated according to the statement of the thread and event transaction name, and the statistical information of each dimension is aggregated, with THREAD_ID and EVENT_NAME (transaction) as the dimensions. |
| 27 | events_transactions_summary_by_user_by_event_name | The transaction summary table is aggregated according to the statement of the user and event transaction name, and the statistical information of each dimension is aggregated, with USER and EVENT_NAME (transaction) as the dimensions. |
| 28 | events_transactions_summary_global_by_event_name | The transaction summary table is aggregated according to the event name (global), and the statistical information of each dimension is aggregated, with EVENT_NAME (transaction) as the dimension. |
| 29 | events_waits_current | Wait summary table, containing currently executing wait events for each thread. |
| 30 | events_waits_history | Wait summary table, containing N (system variable: performance_schema_events_waits_history_size) wait events that each thread has ended. |
| 31 | events_waits_history_long | Wait summary table, containing N (system variable: performance_schema_events_waits_history_long_size) wait events that end globally in all threads. |
| 32 | events_waits_summary_by_account_by_event_name | The waiting summary table is aggregated according to the statement of account, host and event waiting name, and the statistical information of each dimension is aggregated, and USER, HOST, EVENT_NAME (waiting) are listed as dimensions.  |
| 33 | events_waits_summary_by_host_by_event_name | The waiting summary table is aggregated according to the statement of the host and the event waiting name, and the statistical information of each dimension is aggregated, and the HOST and EVENT_NAME (waiting) are listed as the dimensions. |
| 34 | events_waits_summary_by_instance | The waiting summary table is aggregated according to the statement of thread and event waiting function name, and the statistical information of each dimension is aggregated, and EVENT_NAME (waiting) and OBJECT_INSTANCE_BEGIN are listed as dimensions. |
| 35 | events_waits_summary_by_thread_by_event_name | The waiting summary table is aggregated according to the statement of the thread and event waiting name, and the statistical information of each dimension is aggregated, with THREAD_ID and EVENT_NAME (waiting) as the dimensions. |
| 36 | events_waits_summary_by_user_by_event_name | The waiting summary table is aggregated according to the statement of the user and the event waiting name, and the statistical information of each dimension is aggregated, with USER and EVENT_NAME (waiting) as the dimensions. |
| 37 | events_waits_summary_global_by_event_name | The waiting summary table is aggregated according to the event waiting name (global), and the statistical information of each dimension is aggregated, with EVENT_NAME (waiting) as the dimension. |
| 38 | file_instances | List all files (open) of all instances while the server is executing. |
| 39 | file_summary_by_event_name | Performance file I/O summary table, each row counts events for a given event name, with EVENT_NAME (event file) as the dimension. |
| 40 | file_summary_by_instance | Performance file I/O summary table, each row counts events for a given file and event name, with FILE_NAME (file) and EVENT_NAME (event file) as dimensions. |
| 41 | global_status | The global system status table, that is, the values ​​we usually display with showglobalstatuslike'% parameter name%' are stored in this table. |
| 42 |global_variables  | The global system variable table, that is, the values ​​that we usually display with showglobalvariableslike'% parameter name%' are stored in this table. |
| 43 | host_cache | Host cache table statistics, users speed up dns resolution. |
| 44 | hosts | Statistics on the number of current connections and total connections by host dimension. |
| 45 | memory_summary_by_account_by_event_name | The memory summary table is aggregated according to the statement of account, host and event memory object variable name, and the statistical information of each dimension is aggregated, and USER, HOST, EVENT_NAME (memory object variable) is listed as the dimension. |
| 46 | memory_summary_by_host_by_event_name | The memory summary table is aggregated according to the statement of the host and event memory object variable names, and the statistical information of each dimension is aggregated, with HOST and EVENT_NAME (memory object variable) as the dimensions.|
| 47 | memory_summary_by_thread_by_event_name | The memory summary table is aggregated according to the statement of thread and event memory object variable name, and the statistical information of each dimension is aggregated, with THREAD_ID and EVENT_NAME (memory object variable) as the dimensions. |
| 48 | memory_summary_by_user_by_event_name | The memory summary table is aggregated according to the statement of the user and event memory object variable names, and the statistical information of each dimension is aggregated, with USER and EVENT_NAME (memory object variable) as the dimensions. |
| 49 | memory_summary_global_by_event_name | The memory summary table is aggregated according to the event memory object variable name (global), and the statistical information of each dimension is aggregated, and EVENT_NAME (memory object variable) is listed as the dimension. |
| 50 | metadata_locks | Use a record-locked table that uses metadata locking to manage concurrent access to database objects and ensure data consistency (records only happen when they happen). |
| 51 | mutex_instances | See which other thread currently owns the mutex. |
| 52 | objects_summary_global_by_type | Object wait summary table, which is used to summarize the object wait event table. |
| 53 | performance_timers | Used to query the table of available event timer types. |
| 54 | prepared_statements_instances | Prepared Statement Examples and Statistics. |
| 55 | replication_applier_configuration | Check whether each channel is configured with replication delay. |
| 56 | replication_applier_status | Check whether each channel is replicated normally (service_state) and the number of transaction reconnections. |
| 57 | replication_applier_status_by_coordinator | Check whether each channel is copied normally, and copy the wrong code, message and time. |
| 58 | replication_applier_status_by_worker | Check whether the replication of each channel is normal, and copy the work number in parallel, copy the wrong code, SQL and time. |
| 59 | replication_connection_configuration | View the connection configuration information of each channel: host, port, user, auto_position, etc. |
| 60 | replication_connection_status | View the connection information of each channel. |
| 61 | replication_group_member_stats | Only has value when GroupReplication full synchronous replication mode is enabled, the user provides group level information related to the authentication process. |
| 62 | replication_group_members | Only has value when GroupReplication full synchronous replication mode is enabled, which is used to monitor the status of different server instances that are members of the group. |
| 63 | rwlock_instances | The table lists all rwlock (read-write lock) instances related records of the server. |
| 64 | session_account_connect_attrs | A connection properties table containing only the current session and other sessions associated with this session account. |
| 65 | session_connect_attrs | All session connection properties table. |
| 66 | session_status | The user status table at the current session level, that is, the values ​​that we usually use show[session]statuslike'% parameter name%' to display are stored in this table. |
| 67 | session_variables  | The user variable table of the current session level, that is, the values ​​that we usually display with show[session]variableslike'% parameter name%' are stored in this table. |
| 68 | setup_actors | The settings table provides information about the current instrumentation, how to initialize monitoring for new foreground threads, to change the number control of table records, modify at server startup (system variable: performance_schema_setup_actors_size). |
| 69 | setup_consumers | The settings table provides information about the current instrumentation, listing the consumer types that can store event information and which consumers are enabled. |
| 70 | setup_instruments | The settings table provides information about the current instrumentation, classes of instrumented objects that can collect events. |
| 71 | setup_objects | The settings table provides information about the current detection, and controls whether to monitor specific objects, such as changing the number of records table conditions, please modify at server startup (system variable: performance_schema_setup_objects_size). |
| 72 | setup_timers | The settings table provides information about the current detection, the table shows the currently selected event timer. |
| 73 | socket_instances | This table provides a real-time snapshot of active connections to the MySQL server. |
| 74 | socket_summary_by_event_name | Socket summary table, aggregates events by event name, with EVENT_NAME (event name) as the column dimension, provides some additional information, such as socket operations and the number of bytes transmitted and received by the network. |
| 75 | socket_summary_by_instance | Socket summary table, aggregates events by event name, with EVENT_NAME (event name), OBJECT_INSTANCE_BEGIN columns as dimensions, provides some additional information, such as socket operations and the number of bytes transmitted and received by the network. |
| 76 | status_by_account | Status variable summary table with USER, HOST, VARIABLE_NAME columns to summarize status variables by account. |
| 77 | status_by_host | State variable summary table with HOST, VARIABLE_NAME columns for summarizing state variables by the host connected to the client. |
| 78 | status_by_thread |State variable summary table with THREAD_ID, VARIABLE_NAME columns, summary session state variables for each active user. |
| 79 | status_by_user | Status variable summary table with USER, VARIABLE_NAME columns to summarize status variables by client username. |
| 80 | table_handles | Table locks detect the contents of records, this information shows which table handles are opened by the server, how they are locked and by which session, such as changing the number of records table conditions, please modify at server startup (system variable: performance_schema_max_table_handles). |
| 81 | table_io_waits_summary_by_index_usage | This table summarizes all table index I/O wait events generated by wait/io/table/sql/handler, grouped by table index. |
| 82 | table_io_waits_summary_by_table | This table summarizes all table I/O wait events generated by wait/io/table/sql/handler, grouped by table. |
| 83 | table_lock_waits_summary_by_table | This table summarizes all table lock wait events generated by wait/lock/table/sql/handler, grouped by table. |
| 84 | threads | All threads statistics for the current server, each row contains information about the thread and indicates whether monitoring and historical event logging are enabled for that thread. |
| 85 | user_variables_by_thread | Change the table to record some user-defined variables, these are variables defined in a specific session, that is, use set@v1=10;, and then select@v1fromdual; in the current connection. |
| 86 | users | The value records the statistics of the logged in user, does not record the host, the host is in the accounts table. |
| 87 | variables_by_thread | Session system variables for each active session, the session variable table (session_variables, variables_by_thread) only contains information for active sessions, not for terminated sessions. |


## SYS 
&nbsp;&nbsp;&nbsp;&nbsp;All data sources of the Sys library come from: performance_schema. The goal is to reduce the complexity of performance_schema so that DBAs can better read the contents of this library. Let the DBA understand the operation of the DB faster. All the views in the sys library are views, so I will not do anything unnecessary. The MySQL Doc has a unified description:
### sys library table description 
[sys Schema Object Index](https://dev.mysql.com/doc/refman/5.7/en/sys-schema-object-index.html)


## TEST 
&nbsp;&nbsp;&nbsp;&nbsp;The is a test database created during installation. Like its name, it is a completely empty database without any tables and can be deleted(Not explained).

<br>

[back](./)