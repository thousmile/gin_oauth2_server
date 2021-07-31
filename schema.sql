/*
注意： 这是mysql 数据表创建。
建议 其他数据库 根据这些字段创建表。
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for client_details
-- ----------------------------
DROP TABLE IF EXISTS `client_details`;
CREATE TABLE `client_details`
(
    `client_id`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT 'Client Id',
    `secret`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '客户端 密钥',
    `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '客户端名称',
    `logo`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '客户端 图标',
    `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '描述',
    `client_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '客户端类型，0.第三方应用  1.内部应用',
    `grant_types` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '授权类型 json数组格式 [\"sms\",\"password\"]',
    `domain_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '域名地址，如果是 授权码模式，\r\n必须校验回调地址是否属于此域名之下',
    `scope`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '授权作用域',
    `status`      tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态 0.禁用，1.正常',
    PRIMARY KEY (`client_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '[ OAuth2.0 ] 客户端详情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_infos
-- ----------------------------
DROP TABLE IF EXISTS `user_infos`;
CREATE TABLE `user_infos`
(
    `user_id`    char(19) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci     NOT NULL COMMENT '用户ID',
    `avatar`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '头像',
    `username`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '账号',
    `mobile`     varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '手机号码',
    `email`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '邮箱',
    `nickname`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名称',
    `password`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
    `status`     tinyint(1) NOT NULL COMMENT '状态 【0.禁用 1.正常 2.锁定 】',
    `user_type`  tinyint(1) NOT NULL COMMENT '用户类型 0.租户用户 1. 系统用户 ',
    `admin_flag` tinyint(1) NOT NULL COMMENT '0. 普通用户  1. 管理员',
    PRIMARY KEY (`user_id`) USING BTREE,
    UNIQUE INDEX `UK_ulo5s2i7qoksp54tgwl_mobile`(`mobile`) USING BTREE,
    UNIQUE INDEX `UK_6i5ixxulo5s2i7qoksp54tgwl_username`(`username`) USING BTREE,
    UNIQUE INDEX `UK_6i5ixxulo5s2i7984erhgwl_email`(`email`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '[ 系统 ] 用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_social
-- ----------------------------
DROP TABLE IF EXISTS `user_social`;
CREATE TABLE `user_social`
(
    `social_id`   bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '用户社交ID',
    `user_id`     varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '用户唯一ID',
    `open_id`     varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '社交账号唯一ID',
    `social_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT 'we_chat. 微信  tencent_qq. 腾讯QQ',
    PRIMARY KEY (`social_id`) USING BTREE,
    UNIQUE INDEX `UK_sfewfe49823_open_id`(`open_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '[ 系统 ] 用户社交平台登录' ROW_FORMAT = Dynamic;

SET
FOREIGN_KEY_CHECKS = 1;
