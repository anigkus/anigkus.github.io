<h1 style="color:#606c71;text-align:center;" >Do you really understand MySQL's own libraries and tables?</h1><br/>

<center>
<img src="assets/images/do-you-really-understand-mysql-is-own-libraries-and-tables/figure-1.png" alt="Do you really understand MySQL's own libraries and tables" title="Github of Anigkus" >
</center>

> <br/>
> &nbsp;&nbsp;&nbsp;&nbsp; The history of MySQL can be traced back to 1979, and it has been continuously developed and iterated for decades. From the release of MySQL 1.0 in 1996 to the later 3.x, 4.x, 5.x and the latest 8.x, new functions have been continuously developed And improve the problem, it is liked and used by the public, but the built-in database and tables of the database have also changed a lot. I will use 5.7.19 (but the difference between the major version is still very large, and the minor version is basically the same) Let's find out what are the existing self-contained libraries and tables, what data are stored, what are they used for, and what convenience they can bring to our daily database maintenance or development work, let's take a look.<br/>
> <br/>

## INFORMATION_SCHEMA 
&nbsp;&nbsp;&nbsp;&nbsp;The information_schema database comes with MySQL and provides a way to access database metadata. What is metadata? Metadata is data about data, such as database or table names, data types of columns, or access rights, etc. Other terms sometimes used to describe this information include "data dictionary" and "system catalog". In MySQL, information_schema is regarded as a database, to be precise, an information database. It holds information about all other databases maintained by the MySQL server. Such as database name, database table, table column data type and access rights, etc. In INFORMATION_SCHEMA, there are several read-only tables. They are actually views, not base tables, so you won't be able to see any files related to them. Information_schema database table description:

| Serial Number | Table Name | Meaning |
| :--- | :----: | ---: |
| 1 | CHARACTER_SETS | The default correspondence table between characters and proofreading rules cannot be added, updated or deleted.|
| 2 | COLLATIONS | All correspondence tables between characters and collation rules, and there are information such as collation ID, default identification, whether to compile into the server (these are all MySQL extension information for the time being, and are useless now). This table cannot be added, updated, or deleted.|
| 3 | COLLATION_CHARACTER_SET_APPLICABILITY | All correspondence between characters and proofreading rules |
| 4 | COLUMNS |  |
| 5 | COLUMN_PRIVILEGES |  |
| 6 | ENGINES |  |
| 7 | EVENTS |  |
| 8 | FILES |  |
| 9 | GLOBAL_STATUS |  |
| 10 | GLOBAL_VARIABLES |  |
| 11 | KEY_COLUMN_USAGE |  |
| 12 | OPTIMIZER_TRACE |  |
| 13 | PARAMETERS |  |
| 14 | PARTITIONS |  |
| 15 | PLUGINS |  |
| 16 | PROCESSLIST |  |
| 17 | PROFILING |  |
| 18 | REFERENTIAL_CONSTRAINTS |  |
| 19 | ROUTINES |  |
| 20 | SCHEMATA |  |
| 21 | SCHEMA_PRIVILEGES |  |
| 22 | SESSION_STATUS |  |
| 23 | SESSION_VARIABLES |  |
| 24 | STATISTICS |  |
| 25 | TABLES |  |
| 26 | TABLESPACES |  |
| 27 | TABLE_CONSTRAINTS |  |
| 28 | TABLE_PRIVILEGES |  |
| 29 | TRIGGERS |  |
| 30 | USER_PRIVILEGES |  |
| 31 | VIEWS |  |
| 32 | INNODB_LOCKS |  |
| 33 | INNODB_TRX |  |
| 34 | INNODB_SYS_DATAFILES |  |
| 35 | INNODB_FT_CONFIG |  |
| 36 | INNODB_SYS_VIRTUAL |  |
| 37 | INNODB_CMP |  |
| 38 | INNODB_FT_BEING_DELETED |  |
| 39 | INNODB_CMP_RESET |  |
| 40 | INNODB_CMP_PER_INDEX |  |
| 41 | INNODB_CMPMEM_RESET |  |
| 42 |INNODB_FT_DELETED  |  |
| 43 | INNODB_BUFFER_PAGE_LRU |  |
| 44 | INNODB_LOCK_WAITS |  |
| 45 | INNODB_TEMP_TABLE_INFO |  |
| 46 | INNODB_SYS_INDEXES |  |
| 47 | INNODB_SYS_TABLES |  |
| 48 | INNODB_SYS_FIELDS |  |
| 49 | INNODB_CMP_PER_INDEX_RESET |  |
| 50 | INNODB_BUFFER_PAGE |  |
| 51 | INNODB_FT_DEFAULT_STOPWORD |  |
| 52 | INNODB_FT_INDEX_TABLE |  |
| 53 | INNODB_FT_INDEX_CACHE |  |
| 54 | INNODB_SYS_TABLESPACES |  |
| 55 | INNODB_METRICS |  |
| 56 | INNODB_SYS_FOREIGN_COLS |  |
| 57 | INNODB_CMPMEM |  |
| 58 | INNODB_BUFFER_POOL_STATS |  |
| 59 | INNODB_SYS_COLUMNS |  |
| 60 | INNODB_SYS_FOREIGN |  |
| 61 | INNODB_SYS_TABLESTATS |  |

## MYSQL 
&nbsp;&nbsp;&nbsp;&nbsp;The core database of mysql, similar to the master table in sql server, is mainly responsible for storing the control and management information that mysql needs to use, such as database users, permission settings, keywords, etc. mysql database table description: 

| Serial Number | Table Name | Meaning |
| :--- | :----: | ---: |
| 1 | columns_priv | The permission table of the columns in the table, when executing GRANT SELECT (HOST), ON `db1`.* TO 'test'@'%' identified by "123456"; there will be data.|
| 2 | db | The permission table of the database, when executing GRANT SELECT ON `db1`.* TO 'test'@'%' identified by "123456"; there will be data.|
| 3 | engine_cost | Cost model, engine_cost (IO cost), server_cost (CPU cost), mysql execution plan analysis three see explain, profiling, optimizer_trace. |
| 4 | event |  |
| 5 | func |  |
| 6 | general_log |  |
| 7 | gtid_executed |  |
| 8 | help_category |  |
| 9 | help_keyword |  |
| 10 | help_relation |  |
| 11 | help_topic |  |
| 12 | innodb_index_stats |  |
| 13 | innodb_table_stats |  |
| 14 | ndb_binlog_index |  |
| 15 | plugin |  |
| 16 | proc |  |
| 17 | procs_priv |  |
| 18 | proxies_priv |  |
| 19 | server_cost |  |
| 20 | servers |  |
| 21 | slave_master_info |  |
| 22 | slave_relay_log_info |  |
| 23 | slave_worker_info |  |
| 24 | slow_log |  |
| 25 | tables_priv |  |
| 26 | time_zone |  |
| 27 | time_zone_leap_second |  |
| 28 | time_zone_name |  |
| 29 | time_zone_transition |  |
| 30 | time_zone_transition_type |  |
| 31 | user |  |


## PERFORMANCE_SCHEMA 
&nbsp;&nbsp;&nbsp;&nbsp;Mainly used to collect database server performance parameters. And the storage engines of the tables in the library are all PERFORMANCE_SCHEMA, and users cannot create tables whose storage engine is PERFORMANCE_SCHEMA. MySQL5.7 is enabled by default, performance_schema database table description: 

| Serial Number | Table Name | Meaning |
| :--- | :----: | ---: |
| 1 | accounts | Record the current login information and total login statistics of each user corresponding to the host (disconnect one time and record it into the total number of connections), if USER and HOST are NULL, it means internal users.|
| 2 | cond_instances | List all conditions for all instances while the server is executing(visible).|
| 3 | events_stages_current | Stage summary table, containing currently executing stage events for each thread. |
| 4 | events_stages_history |  |
| 5 | events_stages_history_long |  |
| 6 | events_stages_summary_by_account_by_event_name |  |
| 7 | events_stages_summary_by_host_by_event_name |  |
| 8 | events_stages_summary_by_thread_by_event_name |  |
| 9 | events_stages_summary_by_user_by_event_name |  |
| 10 | events_stages_summary_global_by_event_name |  |
| 11 | events_statements_current |  |
| 12 | events_statements_history |  |
| 13 | events_statements_history_long |  |
| 14 | events_statements_summary_by_account_by_event_name |  |
| 15 | events_statements_summary_by_digest |  |
| 16 | events_statements_summary_by_host_by_event_name |  |
| 17 | events_statements_summary_by_program	 |  |
| 18 | events_statements_summary_by_thread_by_event_name |  |
| 19 | events_statements_summary_by_user_by_event_name |  |
| 20 | events_statements_summary_global_by_event_name |  |
| 21 | events_transactions_current |  |
| 22 | events_transactions_history |  |
| 23 | events_transactions_history_long |  |
| 24 | events_transactions_summary_by_account_by_event_name |  |
| 25 | events_transactions_summary_by_host_by_event_name |  |
| 26 | events_transactions_summary_by_thread_by_event_name |  |
| 27 | events_transactions_summary_by_user_by_event_name |  |
| 28 | events_transactions_summary_global_by_event_name |  |
| 29 | events_waits_current |  |
| 30 | events_waits_history |  |
| 31 | events_waits_history_long |  |
| 32 | events_waits_summary_by_account_by_event_name |  |
| 33 | events_waits_summary_by_host_by_event_name |  |
| 34 | events_waits_summary_by_instance |  |
| 35 | events_waits_summary_by_thread_by_event_name |  |
| 36 | events_waits_summary_by_user_by_event_name |  |
| 37 | events_waits_summary_global_by_event_name |  |
| 38 | file_instances |  |
| 39 | file_summary_by_event_name |  |
| 40 | file_summary_by_instance |  |
| 41 | global_status |  |
| 42 |global_variables  |  |
| 43 | host_cache |  |
| 44 | hosts |  |
| 45 | memory_summary_by_account_by_event_name |  |
| 46 | memory_summary_by_host_by_event_name |  |
| 47 | memory_summary_by_thread_by_event_name |  |
| 48 | memory_summary_by_user_by_event_name |  |
| 49 | memory_summary_global_by_event_name |  |
| 50 | metadata_locks |  |
| 51 | mutex_instances |  |
| 52 | objects_summary_global_by_type |  |
| 53 | performance_timers |  |
| 54 | prepared_statements_instances |  |
| 55 | replication_applier_configuration |  |
| 56 | replication_applier_status |  |
| 57 | replication_applier_status_by_coordinator |  |
| 58 | replication_applier_status_by_worker |  |
| 59 | replication_connection_configuration |  |
| 60 | replication_connection_status |  |
| 61 | replication_group_member_stats |  |
| 62 | replication_group_members |  |
| 63 | rwlock_instances |  |
| 64 | session_account_connect_attrs |  |
| 65 | session_connect_attrs |  |
| 66 | session_status |  |
| 67 | session_variables  |  |
| 68 | setup_actors |  |
| 69 | setup_consumers |  |
| 70 | setup_instruments |  |
| 71 | setup_objects |  |
| 72 | setup_timers |  |
| 73 | socket_instances |  |
| 74 | socket_summary_by_event_name |  |
| 75 | socket_summary_by_instance |  |
| 76 | status_by_account |  |
| 77 | status_by_host |  |
| 78 | status_by_thread |  |
| 79 | status_by_user |  |
| 80 | table_handles |  |
| 81 | table_io_waits_summary_by_index_usage |  |
| 82 | table_io_waits_summary_by_table |  |
| 83 | table_lock_waits_summary_by_table |  |
| 84 | threads |  |
| 85 | user_variables_by_thread |  |
| 86 | users |  |
| 87 | variables_by_thread |  |


## SYS 
&nbsp;&nbsp;&nbsp;&nbsp;All data sources of the Sys library come from: performance_schema. The goal is to reduce the complexity of performance_schema so that DBAs can better read the contents of this library. Let the DBA understand the operation of the DB faster. All the views in the sys library are views, so I will not do anything unnecessary. The MySQL Doc has a unified description:
### sys library table description 
[sys Schema Object Index](https://dev.mysql.com/doc/refman/5.7/en/sys-schema-object-index.html)


## TEST 
&nbsp;&nbsp;&nbsp;&nbsp;The is a test database created during installation. Like its name, it is a completely empty database without any tables and can be deleted(Not explained).

<br>

[back](./)