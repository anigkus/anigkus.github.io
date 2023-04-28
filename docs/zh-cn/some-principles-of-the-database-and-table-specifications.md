<iframe src="../detail-header.html" title="Github of Anigkus" style="height:0px,widht:0px;display:none" id="kusifreamheader"></iframe>

<h1 style="color:#606c71;text-align:center;" id="h1" >关于数据库和表的一些原则性规范</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Some principles of the database and table specifications</h1><br/>]:#

<center>
<img src="../assets/images/some-principles-of-the-database-and-table-specifications/figure-1.jpeg" alt="Some principles of the database and table specifications" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; 为了规范数据库操作,统一团队的沟通方式和方便后续项目维护,提高产品代码质量,指导广大软件开发人员编写出简洁、可维护、可靠、可测试、高效、可移植的SQL,参考公司情况和业界数据库规范,因此编写和整理本规范.<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp; In order to standardize database operations, unify team communication methods, facilitate subsequent project maintenance, improve product code quality, and guide software developers to write concise, maintainable, reliable, testable, efficient, and portable SQL, refer to company conditions and industry database specification, so this specification was written and codified.<br/>]:#
[> <br/>]:#

# 术语说明

* <font color="red">原则</font>：编程时必须坚持的指导思想.

* <font color="yellow">规则</font>：编程时强制必须遵守的约定.

* <font color="blue">建议</font>：编程时必须加以考虑的约定.

<mark>说明</mark>:对此原则/规则/建议进行必要的解释.

<mark>示例</mark>:对此原则/规则/建议从正、反两个方面给出例子. 延伸阅读材料：建议进一步阅读的参考材料

[# Term]:#

[* <font color="red">Principle</font>: The guiding ideology that must be adhered to when programming.]:#

[* <font color="yellow">Rules</font>: conventions that must be followed when programming.]:#

[* <font color="blue">Suggestions</font>: Conventions that must be considered when programming.]:#

[<mark>Explanation</mark>: Provide necessary explanation of this principle/rule/recommendation.]:#

[<mark>Examples</mark>: Give examples of this principle/rule/suggestion from both positive and negative aspects. Further reading materials: Reference materials for further reading.]:#

# 规范

[# Specification]:#

## <font color="red">原则</font>: 不允许使用timestrap格式的列

<mark>说明</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; timestrap有性能问题、隐式问题、时间范围等三大问题

<mark>示例</mark>:

[## <font color="red">Principle</font>: Columns in timestamp format are not allowed]:#

[<mark>Explanation</mark>:]:#

[&nbsp;&nbsp;&nbsp;&nbsp;Timestrap has three major problems: performance problems, implicit problems, and time ranges.]:#

[<mark>Example</mark>:]:#

```sql
created_time` timestrap NOT NULL COMMENT 'Created Time',
```

## <font color="red">原则</font>: 数据库&表的编码格式和校对规则

<mark>说明</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; 编码格式: utf8mb4

&nbsp;&nbsp;&nbsp;&nbsp; 校对规则: utf8mb4_unicode_ci

<mark>示例</mark>:

[## <font color="red">Principle</font>: Coding format and Collate rules for databases & tables]:#

[<mark>Explanation</mark>:]:#

[&nbsp;&nbsp;&nbsp;&nbsp; Encoding format: utf8mb4]:#

[&nbsp;&nbsp;&nbsp;&nbsp; Collate rule: utf8mb4_unicode_ci]:#

[<mark>Example</mark>:]:#

Database

```sql
CREATE DATABASE `t_name`  DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ;
```

Table

```sql
CREATE TABLE `d_name`.`t_name` (
  `id` bigint(20) NOT NULL COMMENT 'primary id',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'logical deletion, 0: not deleted, 1: deleted',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci  COMMENT='table comment';
```

## <font color="red">原则</font>: 表、列、字段、索引都要有相关注释

<mark>说明</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; 表COMMENT 'xxx',列COMMENT 'xxx',索引COMMENT 'xxx'

&nbsp;&nbsp;&nbsp;&nbsp; 普通索引:非唯一索引使用`IDX_field_name`来命名

&nbsp;&nbsp;&nbsp;&nbsp; 唯一索引:唯一索引使用`UNI_field_name`来命名

<mark>示例</mark>:

[## <font color="red">Principle</font>: Tables, columns, fields, and indexes must have relevant comments]:#

[<mark>Explanation</mark>:]:#

[&nbsp;&nbsp;&nbsp;&nbsp; Table COMMENT 'xxx', column COMMENT 'xxx', index COMMENT 'xxx']:#

[&nbsp;&nbsp;&nbsp;&nbsp; Normal indexes: non-unique indexes are named using `IDX_field_name`]:#

[&nbsp;&nbsp;&nbsp;&nbsp; Unique indexes: Unique indexes are named using `UNI_field_name`]:#

[<mark>Example</mark>:]:#

```sql
CREATE TABLE `d_name`.`t_name` (
  `id` bigint(20) NOT NULL COMMENT 'primary id',
  `application_id` bigint(20) NOT NULL COMMENT 'application id',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'logical deletion, 0: not deleted, 1: deleted',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
  PRIMARY KEY (`id`),
  KEY `IDX_application_id-is_deleted` (`application_id`,`is_deleted`) COMMENT 'logical delete index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci  COMMENT='table comment';
```

## <font color="red">原则</font>: 每张表中添加created_time和updated_time两个字段

[## <font color="red">Principle</font>: Add created_time and updated_time fields to each table]:#

<mark>说明</mark>:

[<mark>Explanation</mark>:]:#

```sql
`created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
`updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
```

<mark>示例</mark>:

[<mark>Example</mark>:]:#

```sql
CREATE TABLE `d_name`.`t_name` (
  `id` bigint(20) NOT NULL COMMENT 'primary id',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci  COMMENT='table comment';
```

## <font color="red">原则</font>: 增加数据表中新的列时不要插入位置标识

<mark>说明</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; `FIRST`:表示增加在某个列之前

&nbsp;&nbsp;&nbsp;&nbsp; `AFTER`:表示增加在某个列之后

<mark>示例</mark>:

[## <font color="red">Principle</font>: Do not insert position markers when adding new columns in the data table]:#

[<mark>Explanation</mark>:]:#

[&nbsp;&nbsp;&nbsp;&nbsp; `FIRST`:Indicates that the addition is before a column]:#

[&nbsp;&nbsp;&nbsp;&nbsp; `AFTER`:Indicates that the addition is after a column]:#

[<mark>Example</mark>:]:#

Error

```sql
ALTER TABLE `d_name`.`t_name` 
ADD COLUMN `is_deleted` VARCHAR(45) NULL AFTER `updated_by`  COMMENT 'logical deletion, 0: not deleted, 1: deleted';
```

Right

```sql
ALTER TABLE `d_name`.`t_name` 
ADD COLUMN `is_deleted` VARCHAR(45) NULL  COMMENT 'logical deletion, 0: not deleted, 1: deleted';
```

## <font color="red">原则</font>: 更新数据库记录要有备份语句

<mark>说明</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; 更新哪些条件数据的列就备份哪些条件的列

&nbsp;&nbsp;&nbsp;&nbsp; 更新哪些条件的数据就备份哪些条件的数据

<mark>示例</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; 比如我就根据id(1,2,3)更新列created_time,那么就只需要备份id(1,2,3)的created_time列.

[## <font color="red">Principle</font>: Updating database records requires backup statements]:#

[<mark>Explanation</mark>:]:#

[&nbsp;&nbsp;&nbsp;&nbsp; Which conditional data columns are updated, which conditional columns are backed up]:#

[&nbsp;&nbsp;&nbsp;&nbsp; The data of which conditions are updated, the data of which conditions are backed up]:#

[<mark>Example</mark>:]:#

[&nbsp;&nbsp;&nbsp;&nbsp; For example, if I update the column created_time based on id (1, 2, 3), then only the created_time column of id (1, 2, 3) needs to be backed up.]:#

```sql
UPDATE `d_name`.`t_name` SET created_time=NOW() WHERE id in(1,2,3);

SELECT id,created_time  FROM `d_name`.`t_name` WHERE id in(1,2,3) INTO OUTFILE '/tmp/mysql/d_name-t_name-YYYYMMDD-num.csv' FIELDS TERMINATED BY ',' ENCLOSED BY '"' LINES TERMINATED BY '\r\n';
```

## <font color="red">原则</font>: 不准删除表中已经存在的字段

<mark>说明</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; 不能使用`DROP COLUMN`删除列

<mark>示例</mark>:

[## <font color="red">Principle</font>: It is not allowed to delete fields that already exist in the table]:#

[<mark>Explanation</mark>:]:#

[&nbsp;&nbsp;&nbsp;&nbsp; Cannot drop column using `DROP COLUMN` ]:#

[<mark>Example</mark>:]:#

```sql
ALTER TABLE `d_name`.`t_name` 
DROP COLUMN `version`;
```

## <font color="red">原则</font>: 变更时统一提附件并DDL和DML语句分开两个文件

<mark>说明</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; DDL和DML语句分开两个文件

&nbsp;&nbsp;&nbsp;&nbsp; 文件命名:库名-表名-变更日期(YYYYMMDD-1)-ddl.SQL,统一都是.SQL后缀

&nbsp;&nbsp;&nbsp;&nbsp; 统一编码UTF-8，在备注里面写好执行顺序

&nbsp;&nbsp;&nbsp;&nbsp; 增、删、改的操作一定要先备份语句,参考:(更新数据库记录要有备份语句)

[## <font color="red">Principle</font>: When making changes, submit attachments uniformly and separate DDL and DML statements into two files]:#

[<mark>Explanation</mark>:]:#

[&nbsp;&nbsp;&nbsp;&nbsp; DDL and DML statements are separated into two files]:#

[&nbsp;&nbsp;&nbsp;&nbsp; file name: library name-table name-change date (YYYYMMDD-1)-ddl.SQL, all are .SQL suffix]:#

[&nbsp;&nbsp;&nbsp;&nbsp; Unified encoding UTF-8, write the execution order in the remarks]:#

[&nbsp;&nbsp;&nbsp;&nbsp; The operations of adding, deleting and modifying must first backup the statement, refer to: (update database records must have backup statements)]:#

<mark>示例</mark>:

[<mark>Example</mark>:]:#

Name:d_name-t_name-20200820-1-ddl.SQL

```sql
ALTER TABLE `d_name`.`t_name` 
ADD COLUMN `is_deleted` VARCHAR(45) NULL  COMMENT 'logical deletion, 0: not deleted, 1: deleted';
```

Name:d_name-t_name-20200820-1-dml.SQL

```sql
DELETE FROM `d_name`.`t_name` WHERE id= 1;
```

## <font color="red">原则</font>: 必须有唯一主键索引并自增长

[## <font color="red">Principle</font>: Must have a unique primary key index and auto increment]:#

<mark>说明</mark>:

[<mark>Explanation</mark>:]:#

```sql
`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'primary id'
```

<mark>示例</mark>:

[<mark>Example</mark>:]:#

```sql
CREATE TABLE `d_name`.`t_name` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'primary id',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci  COMMENT='table comment';
```

## <font color="red">原则</font>: 表中字段不能有CHAR类型的NULL值

<mark>说明</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; 当ROW_FORMAT=COMPACT的时候CHAR类型的NULL值是不占用空间的

&nbsp;&nbsp;&nbsp;&nbsp; 当ROW_FORMAT=REDUNDANT的时候CHAR类型的NULL值是占用空间的

&nbsp;&nbsp;&nbsp;&nbsp; VARCHAR在这两种行格式中都不占用存储空间

[## <font color="red">Principle</font>: The fields in the table cannot have NULL values of type CHAR]:#

[<mark>Explanation</mark>:]:#

[&nbsp;&nbsp;&nbsp;&nbsp; When ROW_FORMAT=COMPACT, the NULL value of CHAR type does not occupy space]:#

[&nbsp;&nbsp;&nbsp;&nbsp; When ROW_FORMAT=REDUNDANT, the NULL value of CHAR type takes up space]:#

[&nbsp;&nbsp;&nbsp;&nbsp; VARCHAR takes no storage space in either row format]:#

<mark>示例</mark>:

[<mark>Example</mark>:]:#

```sql
CREATE TABLE `d_name`.`t_name` (
  `id` bigint(20) NOT NULL COMMENT 'primary id',
  `application_id` bigint(20) NOT NULL COMMENT 'application id',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'logical deletion, 0: not deleted, 1: deleted',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='table comment';
```

## <font color="red">原则</font>: 不要使用内部关键字作为列名

[## <font color="red">Principle</font>: Do not use internal keywords as column names]:#

<mark>说明</mark>:

[<mark>Explanation</mark>:]:#

&nbsp;&nbsp;&nbsp;&nbsp; 比如:NAME,VALUE,SELECT,ALTER AND,DELETE等等.[参考](https://dev.mysql.com/doc/refman/8.0/en/keywords.html)

[&nbsp;&nbsp;&nbsp;&nbsp; For example:NAME,VALUE,SELECT,ALTER AND,DELETE, etc.,___Refernces+++(https://dev.mysql.com/doc/refman/8.0/en/keywords.html)]:#

<mark>示例</mark>:

[<mark>Example</mark>:]:#

Error:(delete column)

```sql
CREATE TABLE `d_name`.`t_name` (
  `id` bigint(20) NOT NULL COMMENT 'primary id',
  `application_id` bigint(20) NOT NULL COMMENT 'application id',
  `delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'logical deletion, 0: not deleted, 1: deleted',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='table comment';
```

Right:(delete column)

```sql
CREATE TABLE `d_name`.`t_name` (
  `id` bigint(20) NOT NULL COMMENT 'primary id',
  `application_id` bigint(20) NOT NULL COMMENT 'application id',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'logical deletion, 0: not deleted, 1: deleted',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='table comment';
```

## <font color="red">原则</font>: 不要删除任何列和表

[## <font color="red">Principle</font>: Do not delete any column and table]:#

<mark>说明</mark>:

[<mark>Explanation</mark>:]:#

&nbsp;&nbsp;&nbsp;&nbsp; 例如:DROP table

&nbsp;&nbsp;&nbsp;&nbsp; 例如:Alter table table_name drop column column_name;

[&nbsp;&nbsp;&nbsp;&nbsp; For example:DROP table]:#

[&nbsp;&nbsp;&nbsp;&nbsp; For example:Alter table table_name drop column column_name;]:#

<mark>示例</mark>:

[<mark>Example</mark>:]:#

Error:(Drop table & column)

```sql
DROP table tablea;
Alter table tablea drop column col;
```

## <font color="red">原则</font>: 不要更改任何已经存在的表名和列表

[## <font color="red">Principle</font>: Do not change any existing table names and column name]:#

<mark>说明</mark>:

[<mark>Explanation</mark>:]:#

&nbsp;&nbsp;&nbsp;&nbsp; 比如:ALTER TABLE `旧表名` RENAME AS `新表名`;

&nbsp;&nbsp;&nbsp;&nbsp; 比如:ALTER TABLE  `旧表名` change `旧列名` `新列名` column_type;

[&nbsp;&nbsp;&nbsp;&nbsp; For example: ALTER TABLE `old_table` RENAME AS `new_table`;]:#

[&nbsp;&nbsp;&nbsp;&nbsp; For example: ALTER TABLE  `old_table` change `old_column` `new_columb` column_type;]:#

<mark>示例</mark>:

[<mark>Example</mark>:]:#

Error

```sql
ALTER TABLE `old_table` RENAME AS `new_table`;
alter table `old_table` change `old_columnname` `new_columnname` varchar(20);
```

备份表举例子:

[Example of backup table:]:#

```sql
CREATE TABLE `test`.`user1` (
  `id` int(11) NOT NULL,
  `names` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='table user';

--Contains all data and structures (but not primary keys, indexes, etc.)
CREATE TABLE `test`.`user2` AS  
(  
    SELECT *  
    FROM test.user1
);

-- Contains all structure and primary key information (but not data, etc.)
CREATE TABLE `test`.`user3` LIKE `test`.`user1`; 

-- Insert original data
INSERT INTO `test`.`user3`  SELECT * FROM `test`.`user1`;

-- Contains all data and structures (but not primary keys, indexes, etc.)
CREATE TABLE `test`.`user4` SELECT * FROM `test`.`user1`;
```

## <font color="red">原则</font>: 表中删除行记录,产品没有特别说需物理删除,需逻辑标识删除

[## <font color="red">Principle</font>: Delete row records in the table, the product does not specifically say that physical deletion is required, and logical identification is required to delete]:#

<mark>说明</mark>:

[<mark>Explanation</mark>:]:#

```sql
`is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'logical deletion, 0: not deleted, 1: deleted',
```

<mark>示例</mark>:

[<mark>Example</mark>:]:#

```sql
CREATE TABLE `d_name`.`t_name` (
  `id` bigint(20) NOT NULL COMMENT 'primary id',
  `application_id` bigint(20) NOT NULL COMMENT 'application id',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'logical deletion, 0: not deleted, 1: deleted',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci  COMMENT='table comment';
```


<br>

### [back](./)
