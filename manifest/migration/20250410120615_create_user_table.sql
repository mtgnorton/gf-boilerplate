-- +goose Up
-- 管理员表
CREATE TABLE `member` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `password` char(60) NOT NULL COMMENT '密码',
  `salt` char(30) NOT NULL COMMENT '密码盐',
  `status` char(9) NOT NULL DEFAULT 'normal' COMMENT '状态 normal:启用 disabled:禁用',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='管理员表';

CREATE TABLE `role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `code` varchar(50) NOT NULL COMMENT '角色权限字符串',
  `status` char(9) NOT NULL DEFAULT 'normal' COMMENT '状态 normal:启用 disabled:禁用',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色表';

CREATE TABLE `menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `pid` bigint unsigned NOT NULL COMMENT '父菜单ID',
  `level` int unsigned NOT NULL DEFAULT '1' COMMENT '关系树等级',
  `tree` varchar(255) NOT NULL COMMENT '关系树',
  `title` varchar(50) NOT NULL COMMENT '菜单名称',
  `name` varchar(50) NOT NULL COMMENT '菜单标识',
  `path` varchar(100) NOT NULL COMMENT '路由地址',
  `component` varchar(255) NOT NULL COMMENT '组件路径',
  `type` char(10) NOT NULL COMMENT '菜单类型 directory:目录 menu:菜单 button:按钮',
  `permissions` varchar(255) NOT NULL COMMENT '权限标识',
  `status` char(9) NOT NULL DEFAULT 'normal' COMMENT '状态 normal:启用 disabled:禁用',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`),
  KEY `idx_pid` (`pid`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单表';

CREATE TABLE `role_menu_mapping` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `menu_id` bigint unsigned NOT NULL COMMENT '菜单ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色菜单关联表';

CREATE TABLE `casbin` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `p_type` char(10) NOT NULL COMMENT '策略类型',
  `v0` varchar(256) NOT NULL COMMENT '主体',
  `v1` varchar(256) NOT NULL COMMENT '资源',
  `v2` varchar(256) NOT NULL COMMENT '操作',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='casbin规则表';

-- +goose Down
DROP TABLE `member`;
DROP TABLE `role`;
DROP TABLE `menu`;
DROP TABLE `role_menu_mapping`;
DROP TABLE `casbin`;
