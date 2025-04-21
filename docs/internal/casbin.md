CREATE TABLE `hg_admin_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `pid` bigint DEFAULT '0' COMMENT '父菜单ID',

  // Level 表示菜单层级深度，如 1 代表一级菜单,菜单深度增加，Level 值增加
  `level` int NOT NULL DEFAULT '1' COMMENT '关系树等级',

 // 存储该菜单的上级菜单到根菜单的路径，如: 当前菜单2089,上级菜单2087, tree: tr_2093 tr_2074 tr_2087 
  `tree` varchar(255) NOT NULL COMMENT '关系树',
  `title` varchar(64) NOT NULL COMMENT '菜单名称',

  // 唯一标识
  `name` varchar(128) NOT NULL COMMENT '名称编码', 

  // 后端路由地址
  `path` varchar(200) DEFAULT NULL COMMENT '路由地址',

  `icon` varchar(128) DEFAULT NULL COMMENT '菜单图标',
  
  // 1目录: 作为菜单分组，通常只需要显示权限
  // 2菜单: 需要访问权限，控制页面是否可访问
  // 3按钮: 需要操作权限，控制具体功能是否可用
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '菜单类型（1目录 2菜单 3按钮）',
  
  // Redirect: 当访问该菜单时重定向到指定路径
  `redirect` varchar(255) DEFAULT NULL COMMENT '重定向地址','


  // Permissions: 存储该菜单关联的API权限列表，多个权限用逗号分隔
  // 
  // 用户管理菜单的权限集合示例
  // "permissions": "/admin/member/add,/admin/member/edit,/admin/member/delete"
  `permissions` varchar(512) DEFAULT NULL COMMENT '菜单包含权限集合',

   //  PermissionName: 权限的显示名称，用于前端展示
  `permission_name` varchar(64) DEFAULT NULL COMMENT '权限名称',

  //Component: 指定该菜单对应的前端组件
  `component` varchar(255) NOT NULL COMMENT '组件路径',
  `always_show` tinyint(1) DEFAULT '0' COMMENT '取消自动计算根路由模式',
  `active_menu` varchar(255) DEFAULT NULL COMMENT '高亮菜单编码',
  `is_root` tinyint(1) DEFAULT '0' COMMENT '是否跟路由',

   //IsFrame + FrameSrc: 用于在页面中嵌入外部网页
  `is_frame` tinyint(1) DEFAULT '1' COMMENT '是否内嵌',
  `frame_src` varchar(512) DEFAULT NULL COMMENT '内联外部地址',

  // 控制页面是否需要缓存
  `keep_alive` tinyint(1) DEFAULT '0' COMMENT '缓存该路由',
  `hidden` tinyint(1) DEFAULT '0' COMMENT '是否隐藏',


  `affix` tinyint(1) DEFAULT '0' COMMENT '是否固定',
  `sort` int DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) DEFAULT '1' COMMENT '菜单状态',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `pid` (`pid`),
  KEY `status` (`status`),
  KEY `type` (`type`)
) ENGINE=InnoDB AUTO_INCREMENT=2433 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理员_菜单权限';