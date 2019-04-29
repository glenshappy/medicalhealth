SET FOREIGN_KEY_CHECKS=0

DROP TABLE IF EXISTS `address` CASCADE
;

DROP TABLE IF EXISTS `body_parts` CASCADE
;

DROP TABLE IF EXISTS `cart` CASCADE
;

DROP TABLE IF EXISTS `cart_goods` CASCADE
;

DROP TABLE IF EXISTS `category` CASCADE
;

DROP TABLE IF EXISTS `ceshi_order` CASCADE
;

DROP TABLE IF EXISTS `classify` CASCADE
;

DROP TABLE IF EXISTS `coupon_type` CASCADE
;

DROP TABLE IF EXISTS `disease` CASCADE
;

DROP TABLE IF EXISTS `disease_record` CASCADE
;

DROP TABLE IF EXISTS `disease_record_disease_syndrome_yangshengfangan` CASCADE
;

DROP TABLE IF EXISTS `disease_syndrome` CASCADE
;

DROP TABLE IF EXISTS `disease_syndrome_yangshengfangan` CASCADE
;

DROP TABLE IF EXISTS `disease_yangshengfangan` CASCADE
;

DROP TABLE IF EXISTS `goods` CASCADE
;

DROP TABLE IF EXISTS `goods` CASCADE
;

DROP TABLE IF EXISTS `goods_ attached` CASCADE
;

DROP TABLE IF EXISTS `goods_ auditing` CASCADE
;

DROP TABLE IF EXISTS `goods_service` CASCADE
;

DROP TABLE IF EXISTS `label` CASCADE
;

DROP TABLE IF EXISTS `mian_q` CASCADE
;

DROP TABLE IF EXISTS `mian_q` CASCADE
;

DROP TABLE IF EXISTS `news` CASCADE
;

DROP TABLE IF EXISTS `news_label` CASCADE
;

DROP TABLE IF EXISTS `order` CASCADE
;

DROP TABLE IF EXISTS `order_list` CASCADE
;

DROP TABLE IF EXISTS `preliminary` CASCADE
;

DROP TABLE IF EXISTS `preliminary_syndrome` CASCADE
;

DROP TABLE IF EXISTS `preliminary_syndrome_yangshengfangan` CASCADE
;

DROP TABLE IF EXISTS `price_history` CASCADE
;

DROP TABLE IF EXISTS `questionnaire` CASCADE
;

DROP TABLE IF EXISTS `questionnaire_answer` CASCADE
;

DROP TABLE IF EXISTS `questionnaire_category` CASCADE
;

DROP TABLE IF EXISTS `queue_push` CASCADE
;

DROP TABLE IF EXISTS `reg` CASCADE
;

DROP TABLE IF EXISTS `reg_token` CASCADE
;

DROP TABLE IF EXISTS `result_record` CASCADE
;

DROP TABLE IF EXISTS `result_record_preliminary_syndrome_yangshengfangan` CASCADE
;

DROP TABLE IF EXISTS `result_record_syndrome_yangshengfangan` CASCADE
;

DROP TABLE IF EXISTS `specifications` CASCADE
;

DROP TABLE IF EXISTS `syndrome` CASCADE
;

DROP TABLE IF EXISTS `syndrome_keywords` CASCADE
;

DROP TABLE IF EXISTS `syndrome_r_syndrome_keywords` CASCADE
;

DROP TABLE IF EXISTS `syndrome_record` CASCADE
;

DROP TABLE IF EXISTS `syndrome_score` CASCADE
;

DROP TABLE IF EXISTS `syndrome_tongue_coating` CASCADE
;

DROP TABLE IF EXISTS `syndrome_type` CASCADE
;

DROP TABLE IF EXISTS `syndrome_yangshengfangan` CASCADE
;

DROP TABLE IF EXISTS `temp_order` CASCADE
;

DROP TABLE IF EXISTS `temp_order_goods` CASCADE
;

DROP TABLE IF EXISTS `tiantiance` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_answer` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_answer_conflict` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_answer_conflict` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_answer_level` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_common_plan` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_common_plan_yangshengfangan` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_question` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_quetion_detail` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_report` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_tiantiance_answer` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_type` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_type_slogan` CASCADE
;

DROP TABLE IF EXISTS `tiantiance_type_tiantiance_answer` CASCADE
;

DROP TABLE IF EXISTS `tiantiankang` CASCADE
;

DROP TABLE IF EXISTS `tiantiankang_1` CASCADE
;

DROP TABLE IF EXISTS `tongue_coating` CASCADE
;

DROP TABLE IF EXISTS `user_collection` CASCADE
;

DROP TABLE IF EXISTS `xuewei` CASCADE
;

DROP TABLE IF EXISTS `yangshengfangan` CASCADE
;

DROP TABLE IF EXISTS `yangshengfangan` CASCADE
;

DROP TABLE IF EXISTS `yangshengfangan_body_parts` CASCADE
;

DROP TABLE IF EXISTS `yangshengfangan_goods` CASCADE
;

DROP TABLE IF EXISTS `yangshengfangan_xuewei` CASCADE
;

DROP TABLE IF EXISTS `zhongchengyao` CASCADE
;

CREATE TABLE `address`
(
	`id` INTEGER NOT NULL,
	`reg_id` INTEGER NOT NULL,
	CONSTRAINT `PK_address` PRIMARY KEY (`id`)
)
;

CREATE TABLE `body_parts`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`title` VARCHAR(255),
	`position` VARCHAR(255),
	`img` VARCHAR(255),
	`img2` VARCHAR(255),
	`img3` VARCHAR(255),
	`is_delete` INTEGER NOT NULL 0,
	CONSTRAINT `PK_body_parts` PRIMARY KEY (`id`)
)
;

CREATE TABLE `cart`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`user_id` INTEGER NOT NULL,
	CONSTRAINT `PK_cart` PRIMARY KEY (`id`)
)
;

CREATE TABLE `cart_goods`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`cart_id` INTEGER NOT NULL,
	`goods_id` INTEGER NOT NULL,
	`num` INTEGER NOT NULL,
	`create_time` TIMESTAMP,
	`memo` VARCHAR(50),
	`sale_price` DECIMAL(10,0),
	`is_delete` BOOL,
	CONSTRAINT `PK_cart_goods` PRIMARY KEY (`id`)
)
;

CREATE TABLE `category`
(
	`id` INTEGER NOT NULL,
	`name` CHAR(50),
	`updated_at` DATETIME,
	`created_at` DATETIME,
	CONSTRAINT `PK_Table3` PRIMARY KEY (`id`)
)
;

CREATE TABLE `ceshi_order`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	CONSTRAINT `PK_ceshi_order` PRIMARY KEY (`id`)
)
;

CREATE TABLE `classify`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	CONSTRAINT `PK_classify` PRIMARY KEY (`id`)
)
;

CREATE TABLE `coupon_type`
(
	`id` INTEGER AUTO_INCREMENT ,
	`order` VARCHAR(50) NOT NULL
)
;

CREATE TABLE `disease`
(
	`id` INTEGER NOT NULL,
	`disease_name` VARCHAR(100),
	`disease_type` INTEGER NOT NULL,
	`type` INTEGER,
	`synopsis` VARCHAR(255),
	`caneat` VARCHAR(255),
	`cannoteat` VARCHAR(255),
	`conditioning` VARCHAR(50),
	`xuwei` VARCHAR(255),
	`zhongyao` VARCHAR(255),
	`doctortime` VARCHAR(255),
	`suggestroom` VARCHAR(50),
	`delete` INTEGER 0,
	`guanzhurenshu` INTEGER,
	`showindex` INTEGER,
	`pics` VARCHAR(255),
	`orders` INTEGER,
	`cause` VARCHAR(255),
	`emotional_conditioning` VARCHAR(255),
	`living` VARCHAR(255),
	`can_sport` VARCHAR(255),
	`cannot_sport` VARCHAR(255),
	`principle` VARCHAR(255),
	`apply_sex` INTEGER 0,
	CONSTRAINT `PK_disease` PRIMARY KEY (`id`)
)
;

CREATE TABLE `disease_record`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`user_id` INTEGER NOT NULL,
	`test_sn` VARCHAR(50),
	`disease_id` VARCHAR(50),
	`disease_name` VARCHAR(50),
	`time` TIMESTAMP,
	`remark` VARCHAR(50),
	`is_main` BOOL NOT NULL 1,
	CONSTRAINT `PK_disease_record` PRIMARY KEY (`id`)
)
;

CREATE TABLE `disease_record_disease_syndrome_yangshengfangan`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`disease_record_id` INTEGER NOT NULL,
	`disease_syndrome_yangshengfangan_id` INTEGER NOT NULL,
	`create_time` TIMESTAMP,
	`is_use` INTEGER,
	`is_delete` BOOL,
	CONSTRAINT `PK_disease_record_disease_syndrome_yangshengfangan` PRIMARY KEY (`id`)
)
;

CREATE TABLE `disease_syndrome`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`disease_id` INTEGER NOT NULL,
	`syndrome_id` INTEGER NOT NULL,
	`create_time` TIMESTAMP,
	`is_delete` BOOL 0,
	CONSTRAINT `PK_disease_syndrome` PRIMARY KEY (`id`)
)
;

CREATE TABLE `disease_syndrome_yangshengfangan`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`disease_syndrome_id` INTEGER NOT NULL,
	`yangshengfangan_id` INTEGER NOT NULL,
	`create_time` TIMESTAMP,
	`is_delete` BOOL 0,
	CONSTRAINT `PK_disease_syndrome_yangshengfangan` PRIMARY KEY (`id`)
)
;

CREATE TABLE `disease_yangshengfangan`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`disease_id` INTEGER NOT NULL,
	`yangshengfangan_id` INTEGER NOT NULL,
	CONSTRAINT `PK_disease_yangshengfangan` PRIMARY KEY (`id`)
)
;

CREATE TABLE `goods`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	CONSTRAINT `PK_goods` PRIMARY KEY (`id`)
)
;

CREATE TABLE `goods`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`type` INTEGER NOT NULL,
	`parent_goods_id` INTEGER,
	`service_url` VARCHAR(255),
	`summary` VARCHAR(255),
	`ad_label` VARCHAR(255),
	CONSTRAINT `PK_goods` PRIMARY KEY (`id`)
)
;

CREATE TABLE `goods_ attached`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`goods_id` INTEGER NOT NULL,
	`partner_id` INTEGER,
	`is_into_yangshengfangan` BOOL,
	`yangshengfangan_summary` VARCHAR(255),
	CONSTRAINT `PK_partner_goods` PRIMARY KEY (`id`)
)
;

CREATE TABLE `goods_ auditing`
(
	`id` INTEGER NOT NULL,
	`auditing_reg_id` INTEGER NOT NULL,
	`auditing_time` DATETIME,
	`momo` VARCHAR(255),
	`level` INTEGER,
	`goods_id` INTEGER NOT NULL,
	CONSTRAINT `PK_goods_ auditing` PRIMARY KEY (`id`)
)
;

CREATE TABLE `goods_service`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`goods_id` INTEGER,
	CONSTRAINT `PK_goods_service` PRIMARY KEY (`id`)
)
;

CREATE TABLE `label`
(
	`id` INTEGER NOT NULL,
	`name` CHAR(50),
	`memo` VARCHAR(200),
	`created_at` DATETIME,
	`updated_at` DATETIME,
	CONSTRAINT `PK_label` PRIMARY KEY (`id`)
)
;

CREATE TABLE `mian_q`
(
	`id` INTEGER NOT NULL,
	`val` INTEGER,
	CONSTRAINT `PK_mian_q` PRIMARY KEY (`id`)
)
;

CREATE TABLE `mian_q`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	CONSTRAINT `PK_mian_q` PRIMARY KEY (`id`)
)
;

CREATE TABLE `news`
(
	`id` INTEGER NOT NULL,
	`title` CHAR(50),
	`content` TEXT,
	`description` VARCHAR(200),
	`url` VARCHAR(200),
	`category_id` INTEGER,
	`created_at` DATETIME,
	`updated_at` DATETIME,
	CONSTRAINT `PK_news` PRIMARY KEY (`id`)
)
;

CREATE TABLE `news_label`
(
	`id` INTEGER NOT NULL,
	`news_id` INTEGER,
	`label_id` INTEGER,
	`created_at` DATETIME,
	`updated_at` DATETIME,
	CONSTRAINT `PK_Table3` PRIMARY KEY (`id`)
)
;

CREATE TABLE `order`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`order` VARCHAR(50) NOT NULL,
	`adress_id` INTEGER,
	CONSTRAINT `PK_order` PRIMARY KEY (`id`)
)
;

CREATE TABLE `order_list`
(
	`id` INTEGER NOT NULL,
	`order` VARCHAR(50) NOT NULL,
	`goods_id` INTEGER NOT NULL,
	`specifications_id1` INTEGER,
	`specifications_id2` INTEGER,
	`specifications_id3` INTEGER,
	CONSTRAINT `PK_order_list` PRIMARY KEY (`id`)
)
;

CREATE TABLE `preliminary`
(
	`id` INTEGER NOT NULL,
	CONSTRAINT `PK_preliminary` PRIMARY KEY (`id`)
)
;

CREATE TABLE `preliminary_syndrome`
(
	`id` INTEGER NOT NULL,
	`preliminary_id` BIGINT NOT NULL,
	`syndrome_id` INTEGER NOT NULL,
	`create_time` TIMESTAMP,
	CONSTRAINT `PK_preliminary_syndrome` PRIMARY KEY (`id`)
)
;

CREATE TABLE `preliminary_syndrome_yangshengfangan`
(
	`id` INTEGER NOT NULL,
	`preliminary_syndrome_id` INTEGER NOT NULL,
	`yangshengfangan_id` INTEGER NOT NULL,
	`create_time` TIMESTAMP NOT NULL,
	`is_delete` BOOL,
	CONSTRAINT `PK_preliminary_syndrome_yangshengfangan` PRIMARY KEY (`id`)
)
;

CREATE TABLE `price_history`
(
	`id` INTEGER NOT NULL,
	`price_type` INTEGER NOT NULL,
	`goods_id` INTEGER,
	`specifications_id` INTEGER,
	`price` DECIMAL(10,2),
	`start_time` DATETIME,
	`end_time` DATETIME,
	`create_time` TIMESTAMP,
	CONSTRAINT `PK_price_history` PRIMARY KEY (`id`)
)
;

CREATE TABLE `questionnaire`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`title` NVARCHAR(255) NOT NULL,
	`questionnaire_category_id` INTEGER NOT NULL,
	CONSTRAINT `PK_questionnaire` PRIMARY KEY (`id`)
)
;

CREATE TABLE `questionnaire_answer`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`font_name` VARCHAR(50),
	`questionnaire_id` INTEGER NOT NULL,
	`syndrome_keywords_id` INTEGER NOT NULL,
	`mian_q_val` INTEGER NOT NULL,
	CONSTRAINT `PK_questionnaire_answer` PRIMARY KEY (`id`)
)
;

CREATE TABLE `questionnaire_category`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`category_name` VARCHAR(50) NOT NULL,
	CONSTRAINT `PK_questionnaire_category` PRIMARY KEY (`id`)
)
;

CREATE TABLE `queue_push`
(
	`id` INTEGER NOT NULL,
	`label` VARCHAR(50),
	`body` TEXT,
	`ways` VARCHAR(50),
	`user_id` INTEGER,
	`cratetime` DATETIME,
	`worker` VARCHAR(32),
	CONSTRAINT `PK_queue_push` PRIMARY KEY (`id`)
)
;

CREATE TABLE `reg`
(
	`id` INTEGER NOT NULL,
	CONSTRAINT `PK_reg` PRIMARY KEY (`id`)
)
;

CREATE TABLE `reg_token`
(
	`id` INTEGER NOT NULL,
	`user_id` INTEGER,
	`token` VARCHAR(50),
	`expireTime` INTEGER,
	`createTime` INTEGER,
	CONSTRAINT `PK_token` PRIMARY KEY (`id`)
)
;

CREATE TABLE `result_record`
(
	`id` INTEGER NOT NULL,
	`test_sn` VARCHAR(50) NOT NULL,
	CONSTRAINT `PK_result_record` PRIMARY KEY (`id`)
)
;

CREATE TABLE `result_record_preliminary_syndrome_yangshengfangan`
(
	`id` INTEGER NOT NULL,
	`test_sn` VARCHAR(50) NOT NULL,
	`preliminary_syndrome_yangshengfangan_id` INTEGER NOT NULL,
	`create_time` TIMESTAMP NOT NULL,
	`create_user` INTEGER,
	`is_use` INTEGER,
	CONSTRAINT `PK_result_record_preliminary_syndrome_yangshengfangan` PRIMARY KEY (`id`)
)
;

CREATE TABLE `result_record_syndrome_yangshengfangan`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`test_sn` VARCHAR(50) NOT NULL,
	`syndrome_yangshengfangan_id` INTEGER NOT NULL,
	`create_time` TIMESTAMP NOT NULL,
	`create_user` VARCHAR(50) NOT NULL,
	`cycle` INTEGER NOT NULL,
	`frequency` INTEGER NOT NULL,
	`is_use` INTEGER NOT NULL 0,
	CONSTRAINT `PK_result_record_syndrome_yangshengfangan` PRIMARY KEY (`id`)
)
;

CREATE TABLE `specifications`
(
	`id` INTEGER NOT NULL,
	`name` VARCHAR(50) NOT NULL,
	`is_main` BOOL,
	`price` DECIMAL(10,2),
	`memo` VARCHAR(255),
	`create_time` TIMESTAMP NOT NULL,
	`goods_id` INTEGER NOT NULL,
	`is_delete` BOOL NOT NULL,
	CONSTRAINT `PK_specifications` PRIMARY KEY (`id`)
)
;

CREATE TABLE `syndrome`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`name` VARCHAR(20) NOT NULL,
	`syndrome_type_id` INTEGER NOT NULL,
	`is_delete` BOOL NOT NULL,
	`introduce` NVARCHAR(255),
	`treatment_time` NVARCHAR(255),
	`department` NVARCHAR(255),
	`memo1` NVARCHAR(255),
	`memo2` NVARCHAR(255),
	`principle` VARCHAR(200),
	CONSTRAINT `PK_syndrome` PRIMARY KEY (`id`)
)
;

CREATE TABLE `syndrome_keywords`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`k_name` VARCHAR(50) NOT NULL,
	`tizhi_val` INTEGER,
	CONSTRAINT `PK_syndrome_keywords` PRIMARY KEY (`id`)
)
;

CREATE TABLE `syndrome_r_syndrome_keywords`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`syndrome_id` INTEGER NOT NULL,
	`syndrome_keywords_id` INTEGER NOT NULL,
	`score` INTEGER NOT NULL 0,
	CONSTRAINT `PK_syndrome_r_syndrome_keywords` PRIMARY KEY (`id`)
)
;

CREATE TABLE `syndrome_record`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`test_sn` VARCHAR(50) NOT NULL,
	`shexing` INTEGER NOT NULL,
	`shese` INTEGER NOT NULL,
	`taizhi` INTEGER NOT NULL,
	`taise` INTEGER NOT NULL,
	`other` VARCHAR(50) NOT NULL 0,
	`wenzhen` VARCHAR(100),
	CONSTRAINT `PK_syndrome_record` PRIMARY KEY (`id`)
)
;

CREATE TABLE `syndrome_score`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`test_sn` VARCHAR(50) NOT NULL,
	`syndrome_id` INTEGER NOT NULL,
	`main_score` INTEGER NOT NULL,
	`tongue_coating_score` INTEGER NOT NULL,
	`syndrome_keywords_score` INTEGER 0,
	`create_time` DATETIME CURRENT_TIMESTAMP,
	`memo` VARCHAR(50),
	`tongue_coating_score_ismain` INTEGER,
	CONSTRAINT `PK_syndrome_score` PRIMARY KEY (`id`)
)
;

CREATE TABLE `syndrome_tongue_coating`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`syndrome_id` INTEGER NOT NULL,
	`tongue_coating_id` INTEGER NOT NULL,
	`tongue_coating_id_ismain` BOOL NOT NULL false,
	`is_delete` BOOL NOT NULL false,
	`create_time` DATETIME CURRENT_TIMESTAMP,
	CONSTRAINT `PK_syndrome_tongue_coating` PRIMARY KEY (`id`)
)
;

CREATE TABLE `syndrome_type`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`type_name` VARCHAR(50) NOT NULL,
	CONSTRAINT `PK_syndrome_type` PRIMARY KEY (`id`)
)
;

CREATE TABLE `syndrome_yangshengfangan`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`syndrome_id` INTEGER NOT NULL,
	`yangshengfangan_id` INTEGER NOT NULL,
	`type` INTEGER NOT NULL 1,
	CONSTRAINT `PK_syndrome_yangshengfangan` PRIMARY KEY (`id`)
)
;

CREATE TABLE `temp_order`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`order` VARCHAR(50),
	CONSTRAINT `PK_temp_order` PRIMARY KEY (`id`)
)
;

CREATE TABLE `temp_order_goods`
(
	`id` INTEGER NOT NULL,
	`goods_id` INTEGER,
	`specifications_id1` INTEGER,
	`specifications_id2` INTEGER,
	`specifications_id3` INTEGER,
	`order` VARCHAR(50),
	CONSTRAINT `PK_temp_order_goods` PRIMARY KEY (`id`)
)
;

CREATE TABLE `tiantiance`
(
	`ID` INTEGER NOT NULL,
	`user_id` INTEGER NOT NULL,
	`photo_url` VARCHAR(50),
	`tiantiance_date` DATE NOT NULL,
	CONSTRAINT `PK_tiantiance` PRIMARY KEY (`ID`)
)
;

CREATE TABLE `tiantiance_answer`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`tiantiance_question_id` INTEGER NOT NULL,
	`name` VARCHAR(50) NOT NULL,
	`answer_type` INTEGER,
	`answer_value` DECIMAL(10,0),
	`importance_level` INTEGER NOT NULL 0,
	`font_color` VARCHAR(50),
	`for_business` BOOL NOT NULL,
	`can_next` BOOL,
	CONSTRAINT `PK_tiantiance_answer` PRIMARY KEY (`id`)
)
;

CREATE TABLE `tiantiance_answer_conflict`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`tiantiance_answer_id` INTEGER NOT NULL,
	`conflict_id` INTEGER NOT NULL,
	CONSTRAINT `PK_tiantiance_answer_conflict` PRIMARY KEY (`id`)
)
;

CREATE TABLE `tiantiance_answer_conflict`
(
	`ID` INTEGER NOT NULL,
	`tiantiance_answer_id` INTEGER NOT NULL,
	`tiantiance_question_ids` INTEGER,
	CONSTRAINT `PK_tiantiance_question_conflict` PRIMARY KEY (`ID`)
)
;

CREATE TABLE `tiantiance_answer_level`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`name` VARCHAR(50) NOT NULL,
	`code` VARCHAR(50),
	`color` VARCHAR(50),
	`tiantiance_answer_id` INTEGER NOT NULL,
	CONSTRAINT `PK_tiantiance_answer_level` PRIMARY KEY (`id`)
)
;

CREATE TABLE `tiantiance_common_plan`
(
	`ID` INTEGER NOT NULL,
	`tiantiance_answer_id` INTEGER NOT NULL,
	`start_day` INTEGER,
	`end_day` INTEGER,
	`memo` VARCHAR(255),
	`font_color` VARCHAR(50),
	CONSTRAINT `PK_tiantiance_common_plan` PRIMARY KEY (`ID`)
)
;

CREATE TABLE `tiantiance_common_plan_yangshengfangan`
(
	`id` INTEGER NOT NULL,
	`tiantiance_common_plan_id` INTEGER NOT NULL,
	`yangshengfangan_id` INTEGER NOT NULL,
	CONSTRAINT `PK_tiantiance_common_plan_yangshengfangan` PRIMARY KEY (`id`)
)
;

CREATE TABLE `tiantiance_question`
(
	`ID` INTEGER NOT NULL,
	`tiantiance_type_ID` INTEGER NOT NULL,
	`memo` VARCHAR(255),
	`note` VARCHAR(255),
	`is_single` BOOL NOT NULL false,
	`order` INTEGER NOT NULL,
	CONSTRAINT `PK_tiantiance_question` PRIMARY KEY (`ID`)
)
;

CREATE TABLE `tiantiance_quetion_detail`
(
	`ID` INTEGER NOT NULL,
	`tiantiance_question_id` INTEGER NOT NULL,
	`start_times` INTEGER,
	`end_times` INTEGER,
	`memo` VARCHAR(255),
	`note` VARCHAR(255),
	CONSTRAINT `PK_tiantiance_quetion_detail` PRIMARY KEY (`ID`)
)
;

CREATE TABLE `tiantiance_report`
(
	`id` INTEGER NOT NULL,
	`tiantiance_id` INTEGER,
	`body` TEXT,
	`createtime` DATETIME,
	CONSTRAINT `PK_tiantiance_report` PRIMARY KEY (`id`)
)
;

CREATE TABLE `tiantiance_tiantiance_answer`
(
	`id` INTEGER NOT NULL,
	`tiantiance_id` INTEGER NOT NULL,
	`tiantiance_answer_id` INTEGER NOT NULL,
	`answer_value` INTEGER,
	`answer_inc` INTEGER,
	CONSTRAINT `PK_tiantiance_tiantiance_question` PRIMARY KEY (`id`)
)
;

CREATE TABLE `tiantiance_type`
(
	`ID` INTEGER NOT NULL,
	`total_value` INTEGER,
	`name` VARCHAR(50) NOT NULL,
	`icon_img` VARCHAR(50),
	`slogan` VARCHAR(255),
	`slogan_extend_on` BOOL NOT NULL false,
	`order` INTEGER,
	CONSTRAINT `PK_tiantiance_type` PRIMARY KEY (`ID`)
)
;

CREATE TABLE `tiantiance_type_slogan`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`start_time` DATE,
	`end_time` DATE,
	`momo` VARCHAR(255),
	`tiantiance_type_id` INTEGER NOT NULL,
	CONSTRAINT `PK_tiantiance_type_slogan` PRIMARY KEY (`id`)
)
;

CREATE TABLE `tiantiance_type_tiantiance_answer`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`tiantiance_type_id` INTEGER NOT NULL,
	`tiantiance_answer_id` BIGINT NOT NULL,
	`start_times` INTEGER,
	`end_times` INTEGER,
	`memo` VARCHAR(50),
	CONSTRAINT `PK_tiantiance_type_tiantiance_answer` PRIMARY KEY (`id`)
)
;

CREATE TABLE `tiantiankang`
(
	`id` INTEGER NOT NULL,
	`tiantiance_id` INTEGER NOT NULL,
	CONSTRAINT `PK_tiantiankang` PRIMARY KEY (`id`)
)
;

CREATE TABLE `tiantiankang_1`
(
	`id` INTEGER NOT NULL,
	`yangshengfangan_id` INTEGER NOT NULL,
	`tiantiankang_id` INTEGER NOT NULL,
	CONSTRAINT `PK_tiantiankang_1` PRIMARY KEY (`id`)
)
;

CREATE TABLE `tongue_coating`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`features` VARCHAR(20) NOT NULL,
	`is_main` BOOL NOT NULL false,
	`code` INTEGER NOT NULL 0,
	`category` INTEGER NOT NULL,
	`is_delete` BOOL true,
	`tizi_val` INTEGER,
	CONSTRAINT `PK_tongue` PRIMARY KEY (`id`)
)
;

CREATE TABLE `user_collection`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`reg_id` INTEGER NOT NULL,
	`goods_id` INTEGER NOT NULL,
	CONSTRAINT `PK_user_collection` PRIMARY KEY (`id`)
)
;

CREATE TABLE `xuewei`
(
	`id` INTEGER NOT NULL,
	CONSTRAINT `PK_xuewei` PRIMARY KEY (`id`)
)
;

CREATE TABLE `yangshengfangan`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	CONSTRAINT `PK_yangshengfangan` PRIMARY KEY (`id`)
)
;

CREATE TABLE `yangshengfangan`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	CONSTRAINT `PK_yangshengfangan` PRIMARY KEY (`id`)
)
;

CREATE TABLE `yangshengfangan_body_parts`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`yangshengfangan_id` INTEGER NOT NULL,
	`body_parts_id` INTEGER,
	`create_time` TIMESTAMP,
	CONSTRAINT `PK_yangshengfangan_body_parts` PRIMARY KEY (`id`)
)
;

CREATE TABLE `yangshengfangan_goods`
(
	`id` INTEGER NOT NULL,
	`yangshengfangan_id` INTEGER NOT NULL,
	`goods_id` INTEGER NOT NULL,
	CONSTRAINT `PK_yangshengfangan_goods` PRIMARY KEY (`id`)
)
;

CREATE TABLE `yangshengfangan_xuewei`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	`xuewei_id` INTEGER NOT NULL,
	`yangshengfangan_id` INTEGER NOT NULL,
	`create_time` TIMESTAMP,
	CONSTRAINT `PK_yangshengfangan_xuewei` PRIMARY KEY (`id`)
)
;

CREATE TABLE `zhongchengyao`
(
	`id` INTEGER NOT NULL AUTO_INCREMENT ,
	CONSTRAINT `PK_zhongchengyao` PRIMARY KEY (`id`)
)
;

ALTER TABLE `body_parts` 
 ADD CONSTRAINT `UQ_body_parts_id` UNIQUE (`id`)
;

ALTER TABLE `goods_ attached` 
 ADD CONSTRAINT `UQ_partner_goods_id` UNIQUE (`id`)
;

ALTER TABLE `mian_q` 
 ADD CONSTRAINT `UQ_mian_q_val` UNIQUE (`val`)
;

ALTER TABLE `mian_q` 
 ADD CONSTRAINT `UQ_mian_q_id` UNIQUE (`id`)
;

ALTER TABLE `order` 
 ADD CONSTRAINT `UQ_order_order` UNIQUE (`order`)
;

ALTER TABLE `questionnaire` 
 ADD CONSTRAINT `UQ_questionnaire_id` UNIQUE (`id`)
;

ALTER TABLE `questionnaire` 
 ADD CONSTRAINT `UQ_questionnaire_questionnaire_category_id` UNIQUE (`questionnaire_category_id`)
;

ALTER TABLE `questionnaire_answer` 
 ADD CONSTRAINT `UQ_questionnaire_answer_syndrome_keywords_id` UNIQUE (`syndrome_keywords_id`)
;

ALTER TABLE `questionnaire_category` 
 ADD CONSTRAINT `UQ_questionnaire_category_category_name` UNIQUE (`category_name`)
;

ALTER TABLE `result_record` 
 ADD CONSTRAINT `UQ_result_record_test_sn` UNIQUE (`test_sn`)
;

ALTER TABLE `syndrome` 
 ADD CONSTRAINT `UQ_syndrome_id` UNIQUE (`id`)
;

ALTER TABLE `syndrome` 
 ADD CONSTRAINT `UQ_syndrome_name` UNIQUE (`name`)
;

ALTER TABLE `syndrome_keywords` 
 ADD CONSTRAINT `UQ_syndrome_keywords_id` UNIQUE (`id`)
;

ALTER TABLE `syndrome_keywords` 
 ADD CONSTRAINT `UQ_syndrome_keywords_k_name` UNIQUE (`k_name`)
;

ALTER TABLE `syndrome_r_syndrome_keywords` 
 ADD CONSTRAINT `UQ_syndrome_r_syndrome_keywords_id` UNIQUE (`id`)
;

ALTER TABLE `syndrome_record` 
 ADD CONSTRAINT `UQ_syndrome_record_test_sn` UNIQUE (`test_sn`)
;

ALTER TABLE `syndrome_tongue_coating` 
 ADD CONSTRAINT `UQ_syndrome_tongue_coating_id` UNIQUE (`id`)
;

ALTER TABLE `syndrome_type` 
 ADD CONSTRAINT `UQ_syndrome_type_id` UNIQUE (`id`)
;

ALTER TABLE `syndrome_type` 
 ADD CONSTRAINT `UQ_syndrome_type_type_name` UNIQUE (`type_name`)
;

ALTER TABLE `temp_order` 
 ADD CONSTRAINT `UQ_temp_order_order` UNIQUE (`order`)
;

ALTER TABLE `tongue_coating` 
 ADD CONSTRAINT `UQ_tongue_coating_category` UNIQUE (`category`)
;

ALTER TABLE `tongue_coating` 
 ADD CONSTRAINT `UQ_tongue_coating_code` UNIQUE (`code`)
;

ALTER TABLE `tongue_coating` 
 ADD CONSTRAINT `UQ_tongue_coating_features` UNIQUE (`features`)
;

ALTER TABLE `tongue_coating` 
 ADD CONSTRAINT `UQ_tongue_id` UNIQUE (`id`)
;

ALTER TABLE `address` 
 ADD CONSTRAINT `FK_reg_id`
	FOREIGN KEY (`reg_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `cart_goods` 
 ADD CONSTRAINT `FK_cart_id`
	FOREIGN KEY (`cart_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `cart_goods` 
 ADD CONSTRAINT `FK_goods_id`
	FOREIGN KEY (`goods_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `coupon_type` 
 ADD CONSTRAINT `FK_order`
	FOREIGN KEY (`order`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `disease_record` 
 ADD CONSTRAINT `FK_disease_id`
	FOREIGN KEY (`disease_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `disease_record_disease_syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_disease_record_id`
	FOREIGN KEY (`disease_record_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `disease_record_disease_syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_disease_syndrome_yangshengfangan_id`
	FOREIGN KEY (`disease_syndrome_yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `disease_syndrome` 
 ADD CONSTRAINT `FK_disease_id`
	FOREIGN KEY (`disease_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `disease_syndrome` 
 ADD CONSTRAINT `FK_syndrome_id`
	FOREIGN KEY (`syndrome_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `disease_syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_disease_syndrome_id`
	FOREIGN KEY (`disease_syndrome_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `disease_syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_yangshengfangan_id`
	FOREIGN KEY (`yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `disease_yangshengfangan` 
 ADD CONSTRAINT `FK_disease_id`
	FOREIGN KEY (`disease_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `disease_yangshengfangan` 
 ADD CONSTRAINT `FK_yangshengfangan_id`
	FOREIGN KEY (`yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `goods_ attached` 
 ADD CONSTRAINT `FK_goods_id`
	FOREIGN KEY (`goods_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `goods_ auditing` 
 ADD CONSTRAINT `FK_goods_id`
	FOREIGN KEY (`goods_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `goods_service` 
 ADD CONSTRAINT `FK_goods_id`
	FOREIGN KEY (`goods_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `order_list` 
 ADD CONSTRAINT `FK_goods_id`
	FOREIGN KEY (`goods_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `order_list` 
 ADD CONSTRAINT `FK_order`
	FOREIGN KEY (`order`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `preliminary_syndrome` 
 ADD CONSTRAINT `FK_preliminary_id`
	FOREIGN KEY (`preliminary_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `preliminary_syndrome` 
 ADD CONSTRAINT `FK_syndrome_id`
	FOREIGN KEY (`syndrome_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `preliminary_syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_preliminary_syndrome_id`
	FOREIGN KEY (`preliminary_syndrome_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `preliminary_syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_yangshengfangan_id`
	FOREIGN KEY (`yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `questionnaire` 
 ADD CONSTRAINT `FK_questionnaire_category_id`
	FOREIGN KEY (`questionnaire_category_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `questionnaire_answer` 
 ADD CONSTRAINT `FK_mian_q_val`
	FOREIGN KEY (`mian_q_val`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `questionnaire_answer` 
 ADD CONSTRAINT `FK_questionnaire_id`
	FOREIGN KEY (`questionnaire_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `questionnaire_answer` 
 ADD CONSTRAINT `FK_syndrome_keywords_id`
	FOREIGN KEY (`syndrome_keywords_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `reg_token` 
 ADD CONSTRAINT `FK_user_id`
	FOREIGN KEY () REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `result_record_preliminary_syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_preliminary_syndrome_yangshengfangan_id`
	FOREIGN KEY (`preliminary_syndrome_yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `result_record_preliminary_syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_test_sn`
	FOREIGN KEY (`test_sn`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `result_record_syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_syndrome_yangshengfangan_id`
	FOREIGN KEY (`syndrome_yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `result_record_syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_test_sn`
	FOREIGN KEY (`test_sn`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `specifications` 
 ADD CONSTRAINT `FK_goods_id`
	FOREIGN KEY (`goods_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `syndrome` 
 ADD CONSTRAINT `FK_syndrome_type_id`
	FOREIGN KEY (`syndrome_type_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `syndrome_r_syndrome_keywords` 
 ADD CONSTRAINT `FK_syndrome_id`
	FOREIGN KEY (`syndrome_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `syndrome_r_syndrome_keywords` 
 ADD CONSTRAINT `FK_syndrome_keywords_id`
	FOREIGN KEY (`syndrome_keywords_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `syndrome_record` 
 ADD CONSTRAINT `FK_test_sn`
	FOREIGN KEY (`test_sn`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `syndrome_score` 
 ADD CONSTRAINT `FK_syndrome_id`
	FOREIGN KEY (`syndrome_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `syndrome_tongue_coating` 
 ADD CONSTRAINT `FK_syndrome_id`
	FOREIGN KEY (`syndrome_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_syndrome_id`
	FOREIGN KEY (`syndrome_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `syndrome_yangshengfangan` 
 ADD CONSTRAINT `FK_yangshengfangan_id`
	FOREIGN KEY (`yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `temp_order_goods` 
 ADD CONSTRAINT `FK_goods_id`
	FOREIGN KEY (`goods_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance` 
 ADD CONSTRAINT `FK_user_id`
	FOREIGN KEY (`user_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_answer` 
 ADD CONSTRAINT `FK_tiantiance_question_id`
	FOREIGN KEY (`tiantiance_question_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_answer_conflict` 
 ADD CONSTRAINT `FK_tiantiance_answer_id`
	FOREIGN KEY (`tiantiance_answer_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_answer_conflict` 
 ADD CONSTRAINT `FK_tiantiance_question_id`
	FOREIGN KEY (`tiantiance_answer_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_answer_level` 
 ADD CONSTRAINT `FK_tiantiance_answer_id`
	FOREIGN KEY (`tiantiance_answer_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_common_plan` 
 ADD CONSTRAINT `FK_tiantiance_question_id`
	FOREIGN KEY (`tiantiance_answer_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_common_plan_yangshengfangan` 
 ADD CONSTRAINT `FK_tiantiance_common_plan_id`
	FOREIGN KEY (`tiantiance_common_plan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_common_plan_yangshengfangan` 
 ADD CONSTRAINT `FK_yangshengfangan_id`
	FOREIGN KEY (`yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_question` 
 ADD CONSTRAINT `FK_tiantiance_type_ID`
	FOREIGN KEY (`tiantiance_type_ID`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_quetion_detail` 
 ADD CONSTRAINT `FK_tiantiance_question_id`
	FOREIGN KEY (`tiantiance_question_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_report` 
 ADD CONSTRAINT `FK_tiantiance_id`
	FOREIGN KEY () REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_tiantiance_answer` 
 ADD CONSTRAINT `FK_tiantiance_id`
	FOREIGN KEY (`tiantiance_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_tiantiance_answer` 
 ADD CONSTRAINT `FK_tiantiance_question_id`
	FOREIGN KEY (`tiantiance_answer_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_type_slogan` 
 ADD CONSTRAINT `FK_tiantiance_type_id`
	FOREIGN KEY (`tiantiance_type_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_type_tiantiance_answer` 
 ADD CONSTRAINT `FK_tiantiance_answer_id`
	FOREIGN KEY (`tiantiance_answer_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiance_type_tiantiance_answer` 
 ADD CONSTRAINT `FK_tiantiance_type_id`
	FOREIGN KEY (`tiantiance_type_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiankang` 
 ADD CONSTRAINT `FK_tiantiance_id`
	FOREIGN KEY (`tiantiance_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiankang_1` 
 ADD CONSTRAINT `FK_tiantiankang_id`
	FOREIGN KEY (`tiantiankang_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `tiantiankang_1` 
 ADD CONSTRAINT `FK_yangshengfangan_id`
	FOREIGN KEY (`yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `user_collection` 
 ADD CONSTRAINT `FK_goods_id`
	FOREIGN KEY (`goods_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `user_collection` 
 ADD CONSTRAINT `FK_reg_id`
	FOREIGN KEY (`reg_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `yangshengfangan_body_parts` 
 ADD CONSTRAINT `FK_body_parts_id`
	FOREIGN KEY (`body_parts_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `yangshengfangan_body_parts` 
 ADD CONSTRAINT `FK_yangshengfangan_id`
	FOREIGN KEY (`yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `yangshengfangan_goods` 
 ADD CONSTRAINT `FK_goods_id`
	FOREIGN KEY (`goods_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `yangshengfangan_goods` 
 ADD CONSTRAINT `FK_yangshengfangan_id`
	FOREIGN KEY (`yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `yangshengfangan_xuewei` 
 ADD CONSTRAINT `FK_xuewei_id`
	FOREIGN KEY (`xuewei_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE `yangshengfangan_xuewei` 
 ADD CONSTRAINT `FK_yangshengfangan_id`
	FOREIGN KEY (`yangshengfangan_id`) REFERENCES  () ON DELETE No Action ON UPDATE No Action
;

SET FOREIGN_KEY_CHECKS=1
