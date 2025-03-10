use ferry;
DROP TABLE IF EXISTS `sys_user_info`;
CREATE TABLE `sys_user_info` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `user_id` int NOT NULL ,
                                 `username` varchar(64) DEFAULT NULL,
                                 `password` varchar(128) DEFAULT NULL,
                                 `phone` varchar(11) DEFAULT NULL,
                                 `sex` varchar(255) DEFAULT NULL,
                                 `create_by` varchar(128) DEFAULT NULL,
                                 `update_by` varchar(128) DEFAULT NULL,
                                 `remark` varchar(255) DEFAULT NULL,
                                 `status` int DEFAULT NULL,
                                 `create_time` timestamp NULL DEFAULT NULL,
                                 `update_time` timestamp NULL DEFAULT NULL,
                                 `delete_time` timestamp NULL DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 KEY `idx_sys_user_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
