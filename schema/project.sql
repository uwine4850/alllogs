USE alllogs;

CREATE TABLE IF NOT EXISTS `alllogs`.`project`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `name` VARCHAR(200) NOT NULL,
    `description` VARCHAR(200) NULL,
    PRIMARY KEY(id),
    FOREIGN KEY (user_id) REFERENCES `auth`(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `alllogs`.`project_log_group`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `project_id` INT NOT NULL,
    `name` VARCHAR(200) NOT NULL,
    `description` VARCHAR(200) NULL,
    PRIMARY KEY(id),
    FOREIGN KEY (project_id) REFERENCES `project`(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `alllogs`.`log_item`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `log_group_id` INT NOT NULL,
    `text` VARCHAR(200) NOT NULL,
    type ENUM('INFO', 'WARN', 'ERROR') NOT NULL,
    `tag` VARCHAR(200) NULL,
    `datatime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY (log_group_id) REFERENCES `project_log_group`(id) ON DELETE CASCADE
);
