USE alllogs;

CREATE TABLE IF NOT EXISTS `alllogs`.`profile` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `description` TEXT NULL,
    `avatar` TEXT NULL,
    `token` VARCHAR(300) NULL,
    PRIMARY KEY(id),
    UNIQUE (`token`),
    FOREIGN KEY (user_id) REFERENCES `auth`(id) ON DELETE CASCADE
);

DELIMITER //

CREATE TRIGGER set_default_avatar BEFORE INSERT ON `profile`
FOR EACH ROW
BEGIN
    IF NEW.avatar IS NULL THEN
        SET NEW.avatar = '/storage/avatars/default.jpg';
    END IF;
END//

DELIMITER ;