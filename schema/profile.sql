USE alllogs;

CREATE TABLE IF NOT EXISTS `alllogs`.`profile` (
    `user_id` INT NOT NULL UNIQUE,
    `description` TEXT NULL,
    `avatar` TEXT NULL,
    `token` VARCHAR(300) NULL,
    UNIQUE (`token`),
    FOREIGN KEY (user_id) REFERENCES `auth`(id) ON DELETE CASCADE
);

DELIMITER //

CREATE TRIGGER IF NOT EXISTS set_default_avatar BEFORE INSERT ON `profile`
FOR EACH ROW
BEGIN
    IF NEW.avatar IS NULL THEN
        SET NEW.avatar = '/storage/avatars/default.jpg';
    END IF;
END//

DELIMITER ;

CREATE TABLE IF NOT EXISTS notifications (
    id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    type ENUM('info', 'group_invite', 'project') NOT NULL,
    payload JSON NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY (user_id) REFERENCES `auth`(id) ON DELETE CASCADE
);
