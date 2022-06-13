<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >Some principles of the database and table specifications</h1><br/>

<center>
<img src="assets/images/some-principles-of-the-database-and-table-specifications/figure-1.jpeg" alt="Some principles of the database and table specifications" title="Github of Anigkus" >
</center>

> <br/>&nbsp;&nbsp;&nbsp;&nbsp; In order to standardize database operations, unify team communication methods, facilitate subsequent project maintenance, improve product code quality, and guide software developers to write concise, maintainable, reliable, testable, efficient, and portable SQL, refer to company conditions and industry database specification, so this specification was written and codified.<br/>
> <br/>

# term description

* <font color="red">Principle</font>: The guiding ideology that must be adhered to when programming.

* <font color="yellow">Rules</font>: conventions that must be followed when programming.

* <font color="blue">Suggestions</font>: Conventions that must be considered when programming.

<mark>Explanation</mark>: Provide necessary explanation of this principle/rule/recommendation.

<mark>Examples</mark>: Give examples of this principle/rule/suggestion from both positive and negative aspects. Further reading materials: Reference materials for further reading.

# Specification

## <font color="red">Principle</font>: Columns in timestamp format are not allowed

<mark>Explanation</mark>:

&nbsp;&nbsp;&nbsp;&nbsp;Timestrap has three major problems: performance problems, implicit problems, and time ranges.

<mark>Example</mark>:

```sql
created_time` timestrap NOT NULL COMMENT 'Created Time',
```

## <font color="red">Principle</font>: Coding format and Collate rules for databases & tables

<mark>Explanation</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; Encoding format: utf8mb4

&nbsp;&nbsp;&nbsp;&nbsp; Collate rule: utf8mb4_unicode_ci

<mark>Example</mark>:

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

## <font color="red">Principle</font>: Tables, columns, fields, and indexes must have relevant comments

<mark>Explanation</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; Table COMMENT 'xxx', column COMMENT 'xxx', index COMMENT 'xxx'

&nbsp;&nbsp;&nbsp;&nbsp; Normal indexes: non-unique indexes are named using `IDX_field_name`

&nbsp;&nbsp;&nbsp;&nbsp; Unique indexes: Unique indexes are named using `UNI_field_name`

<mark>Example</mark>:

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

## <font color="red">Principle</font>: Add created_time and updated_time fields to each table

<mark>Explanation</mark>:

```sql
`created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
`updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
```

<mark>Example</mark>:

```sql
CREATE TABLE `d_name`.`t_name` (
  `id` bigint(20) NOT NULL COMMENT 'primary id',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci  COMMENT='table comment';
```

## <font color="red">Principle</font>: Do not insert position markers when adding new columns in the data table

<mark>Explanation</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; `FIRST`:Indicates that the addition is before a column

&nbsp;&nbsp;&nbsp;&nbsp; `AFTER`:Indicates that the addition is after a column

<mark>Example</mark>:

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

## <font color="red">Principle</font>: Updating database records requires backup statements

<mark>Explanation</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; Which conditional data columns are updated, which conditional columns are backed up

&nbsp;&nbsp;&nbsp;&nbsp; The data of which conditions are updated, the data of which conditions are backed up

<mark>Example</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; For example, if I update the column created_time based on id (1, 2, 3), then only the created_time column of id (1, 2, 3) needs to be backed up.

```sql
UPDATE `d_name`.`t_name` SET created_time=NOW() WHERE id in(1,2,3);

SELECT id,created_time  FROM `d_name`.`t_name` WHERE id in(1,2,3) INTO OUTFILE '/tmp/mysql/d_name-t_name-YYYYMMDD-num.csv' FIELDS TERMINATED BY ',' ENCLOSED BY '"' LINES TERMINATED BY '\r\n';
```

## <font color="red">原则</font>: It is not allowed to delete fields that already exist in the table

<mark>Explanation</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; Cannot drop column using `DROP COLUMN` 

<mark>Example</mark>:

```sql
ALTER TABLE `d_name`.`t_name` 
DROP COLUMN `version`;
```

## <font color="red">Principle</font>: When making changes, submit attachments uniformly and separate DDL and DML statements into two files

<mark>Explanation</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; DDL and DML statements are separated into two files

&nbsp;&nbsp;&nbsp;&nbsp; file name: library name-table name-change date (YYYYMMDD-1)-ddl.SQL, all are .SQL suffix

&nbsp;&nbsp;&nbsp;&nbsp; Unified encoding UTF-8, write the execution order in the remarks

&nbsp;&nbsp;&nbsp;&nbsp; The operations of adding, deleting and modifying must first backup the statement, refer to: (update database records must have backup statements)

<mark>Example</mark>:

Name:d_name-t_name-20200820-1-ddl.SQL

```sql
ALTER TABLE `d_name`.`t_name` 
ADD COLUMN `is_deleted` VARCHAR(45) NULL  COMMENT 'logical deletion, 0: not deleted, 1: deleted';
```

Name:d_name-t_name-20200820-1-dml.SQL

```sql
DELETE FROM `d_name`.`t_name` WHERE id= 1;
```

## <font color="red">Principle</font>: Must have a unique primary key index and auto increment

<mark>Explanation</mark>:

```sql
`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'primary id'
```

<mark>Example</mark>:

```sql
CREATE TABLE `d_name`.`t_name` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'primary id',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci  COMMENT='table comment';
```

## <font color="red">Principle</font>: The fields in the table cannot have NULL values of type CHAR

<mark>Explanation</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; When ROW_FORMAT=COMPACT, the NULL value of CHAR type does not occupy space

&nbsp;&nbsp;&nbsp;&nbsp; When ROW_FORMAT=REDUNDANT, the NULL value of CHAR type takes up space

&nbsp;&nbsp;&nbsp;&nbsp; VARCHAR takes no storage space in either row format

<mark>Example</mark>:

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

## <font color="red">Principle</font>: Do not use internal keywords as column names

<mark>Explanation</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; For example:NAME,VALUE,SELECT,ALTER AND,DELETE, etc.,[Refernces](https://dev.mysql.com/doc/refman/8.0/en/keywords.html)

<mark>Example</mark>:

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

## <font color="red">Principle</font>: Do not delete any column and table

<mark>Explanation</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; For example:DROP table

&nbsp;&nbsp;&nbsp;&nbsp; For example:Alter table table_name drop column column_name;

<mark>Example</mark>:

Error:(Drop table & column)

```sql
DROP table tablea;
Alter table tablea drop column col;
```

## <font color="red">Principle</font>: Do not change any existing table names and column name

<mark>Explanation</mark>:

&nbsp;&nbsp;&nbsp;&nbsp; For example: ALTER TABLE `old_table` RENAME AS `new_table`;

&nbsp;&nbsp;&nbsp;&nbsp; For example: ALTER TABLE  `old_table` change `old_column` `new_columb` column_type;

<mark>示例</mark>:

<mark>Example</mark>:

Error

```sql
ALTER TABLE `old_table` RENAME AS `new_table`;
alter table `old_table` change `old_columnname` `new_columnname` varchar(20);
```

Example of backup table:

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

## <font color="red">Principle</font>: Delete row records in the table, the product does not specifically say that physical deletion is required, and logical identification is required to delete

<mark>Explanation</mark>:

```sql
`is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'logical deletion, 0: not deleted, 1: deleted',
```

<mark>Example</mark>:

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
