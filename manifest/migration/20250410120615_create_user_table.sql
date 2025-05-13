-- +goose Up
-- 管理员表
CREATE TABLE `member` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `password_hash` char(60) NOT NULL COMMENT '密码',
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

INSERT INTO `role` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '超级管理员', 'super_admin', 'normal', '2025-04-30 00:00:00', "2025-04-30 00:00:00", NULL);

CREATE TABLE `menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `pid` bigint unsigned NOT NULL COMMENT '父菜单ID,pid=0时为顶级菜单',
  `level` int unsigned NOT NULL DEFAULT '1' COMMENT '关系树等级,pid=0的menu等级为1,然后根据父子关系依次递增',
  `tree` varchar(255) NOT NULL COMMENT '关系树,将当前menu的所有父级menu的id用逗号分隔拼接而成,顺序从直接上级到最顶级',
  `name` varchar(50) NOT NULL COMMENT '菜单名称',
  `code` varchar(50) NOT NULL COMMENT '菜单标识',
  `path` varchar(100) NOT NULL COMMENT '前端路由地址',
  `component` varchar(255) NOT NULL COMMENT '前端组件路径',
  `type` char(10) NOT NULL COMMENT '菜单类型 directory:目录 menu:菜单 button:按钮',
  `permissions` varchar(255) NOT NULL COMMENT '该menu对应的后端路由(权限标识),如果一个菜单页面具有多个后端路由,用逗号分隔',
  `icon` varchar(255) NOT NULL COMMENT '菜单图标',
  `redirect` varchar(255) NOT NULL COMMENT '重定向地址',
  `bar_active_code`  varchar(50) NOT NULL COMMENT '当前页面高亮哪个code对应的左侧菜单,例如,在/user/edit页面时,左侧菜单仍然高亮/user/list',
  `is_external` bit(1) NOT NULL DEFAULT 0 COMMENT '是否为外链,0:否 1:是',
  `external_url` varchar(255) NOT NULL COMMENT '外链地址',
  `status` char(9) NOT NULL DEFAULT 'normal' COMMENT '状态 normal:启用 disabled:禁用',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`),
  KEY `idx_pid` (`pid`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单表';

INSERT INTO `menu` (`id`, `pid`, `level`, `tree`, `name`, `code`, `path`, `component`, `type`, `permissions`, `icon`, `redirect`, `bar_active_code`, `is_external`, `external_url`, `status`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 0, 1, '', '系统管理', 'system', '/system', 'Layout', 'directory', '', 'setting', '/system/user', 'system', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(2, 1, 2, '1', '用户管理', 'system:user', '/system/user', 'system/user/index', 'menu', 'auth.user.getList,auth.user.create,auth.user.update,auth.user.delete', 'user', '', 'system:user', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(3, 2, 3, '2,1', '用户新增', 'system:user:create', '', '', 'button', 'auth.user.create', '', '', '', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(4, 2, 3, '2,1', '用户编辑', 'system:user:update', '', '', 'button', 'auth.user.update', '', '', '', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(5, 2, 3, '2,1', '用户删除', 'system:user:delete', '', '', 'button', 'auth.user.delete', '', '', '', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(6, 1, 2, '1', '角色管理', 'system:role', '/system/role', 'system/role/index', 'menu', 'auth.role.getList,auth.role.create,auth.role.update,auth.role.delete', 'peoples', '', 'system:role', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(7, 6, 3, '6,1', '角色新增', 'system:role:create', '', '', 'button', 'auth.role.create', '', '', '', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(8, 6, 3, '6,1', '角色编辑', 'system:role:update', '', '', 'button', 'auth.role.update', '', '', '', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(9, 6, 3, '6,1', '角色删除', 'system:role:delete', '', '', 'button', 'auth.role.delete', '', '', '', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(10, 1, 2, '1', '菜单管理', 'system:menu', '/system/menu', 'system/menu/index', 'menu', 'auth.menu.getList,auth.menu.create,auth.menu.update,auth.menu.delete', 'tree-table', '', 'system:menu', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(11, 10, 3, '10,1', '菜单新增', 'system:menu:create', '', '', 'button', 'auth.menu.create', '', '', '', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(12, 10, 3, '10,1', '菜单编辑', 'system:menu:update', '', '', 'button', 'auth.menu.update', '', '', '', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL),
(13, 10, 3, '10,1', '菜单删除', 'system:menu:delete', '', '', 'button', 'auth.menu.delete', '', '', '', b'0', '', 'normal', 0, '2025-04-30 00:00:00', '2025-04-30 00:00:00', NULL);


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
