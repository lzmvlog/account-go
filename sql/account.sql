/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : 127.0.0.1:3306
 Source Schema         : account

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 11/12/2021 20:31:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for subjects
-- ----------------------------
DROP TABLE IF EXISTS `subjects`;
CREATE TABLE `subjects`  (
  `id` int(0) NULL DEFAULT NULL,
  `sub_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `parent_id` int(0) NULL DEFAULT NULL COMMENT '父级id',
  `direction` int(0) NULL DEFAULT NULL COMMENT '0 借 1 贷',
  `code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '编码',
  `is_enable` int(0) NULL DEFAULT NULL COMMENT '0 启用 1 禁用'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '会计科目表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of subjects
-- ----------------------------
INSERT INTO `subjects` VALUES (1, '扫码支付', NULL, 1, '1', 0);
INSERT INTO `subjects` VALUES (2, '支付宝', 1, 1, '1-1', 0);
INSERT INTO `subjects` VALUES (3, '微信', 1, 1, '1-2', 0);

SET FOREIGN_KEY_CHECKS = 1;
