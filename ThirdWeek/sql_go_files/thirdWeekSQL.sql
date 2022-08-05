# 创建数据库
DROP DATABASE IF EXISTS third_week;
CREATE DATABASE third_week;
USE	third_week;

# 创建 仓库信息表
DROP TABLE IF EXISTS storehouse;
CREATE TABLE storehouse(
                           store_code VARCHAR(255) NOT NULL COMMENT '仓库编码',
                           capacity INT(255)	NULL	COMMENT '仓库容量',
                           PRIMARY KEY(store_code)
)ENGINE = INNODB CHARACTER SET = utf8mb4 COMMENT '仓库信息表';

# 创建 服装表
DROP TABLE IF EXISTS clothing_info;
CREATE TABLE clothing_info(
                              clothing_code VARCHAR(255) NOT NULL COMMENT '服装编码',
                              size VARCHAR(30)	NULL COMMENT	'服装尺码',
                              price INT(255)	NULL	COMMENT '销售价格',
                              clothing_type VARCHAR(255)	NULL COMMENT	'服装类型',
                              PRIMARY KEY(clothing_code)
)ENGINE = INNODB CHARACTER SET = utf8mb4 COMMENT '服装表';

# 创建 供应商表
DROP TABLE IF EXISTS supplier;
CREATE TABLE supplier(
                         supplier_code VARCHAR(255) NOT NULL COMMENT '供应商编码',
                         supplier_name VARCHAR(255) NULL COMMENT '供应商名称',
                         PRIMARY KEY(supplier_code)
)ENGINE = INNODB CHARACTER SET = utf8mb4 COMMENT '供应商表';

# 创建 供应情况表
DROP TABLE IF EXISTS supply_situation;
CREATE TABLE supply_situation(
                                 clothing_code VARCHAR(255) NOT NULL COMMENT '服装编码',
                                 supplier_code VARCHAR(255) NOT NULL COMMENT '供应商编码',
                                 quality VARCHAR(255) NULL COMMENT '服装质量',
                                 PRIMARY KEY(clothing_code, supplier_code)
)ENGINE = INNODB CHARACTER SET = utf8mb4 COMMENT '供应情况表';


# 向仓库信息表中添加数据
INSERT INTO storehouse VALUES('CK1001',9999),('CK1002',5000),('CK1003',8000),('CK1004',10000),('CK1005',7000);

# 向服装表中添加数据
INSERT INTO clothing_info VALUES('AFZ00001','S',110,'A'),('BFZ00002','L',95,'B'),('AFZ00003','M',120,'A'),('BFZ00004','S',90,'B'),('CFZ00005','S',85,'C');

# 向供应商表中添加数据
INSERT INTO supplier VALUES('GYS1001','供应商A'),('GYS1002','供应商B'),('GYS1003','供应商C'),('GYS1004','供应商D'),('GYS1005','供应商E');

# 向供应情况表中添加数据
INSERT INTO supply_situation VALUES('AFZ00001','GYS1001','合格'),('BFZ00002','GYS1002','不合格'),('AFZ00003','GYS1003','合格'),('BFZ00004','GYS1004','合格'),('CFZ00005','GYS1005','不合格');
