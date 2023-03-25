CREATE TABLE `app_user`
(
    `id`          bigint      NOT NULL AUTO_INCREMENT,
    `name`        varchar(20) NOT NULL DEFAULT '' COMMENT '用户昵称',
    `sex`         varchar(4)  NOT NULL DEFAULT 'm' COMMENT '用户性别',
    `create_time` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='用户信息';