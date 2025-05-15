-- +goose Up
-- 超级管理员表
CREATE TABLE `super_admin` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `password_hash` char(60) NOT NULL COMMENT '密码',
  `salt` char(30) NOT NULL COMMENT '密码盐',
  `status` char(9) NOT NULL DEFAULT 'normal' COMMENT '状态 normal:启用 disabled:禁用',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='超级管理员表';




-- +goose Down
DROP TABLE `super_admin`;