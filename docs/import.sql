CREATE TABLE IF NOT EXISTS `history`
(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `source` varchar(255) DEFAULT NULL,
    `destination` varchar(255) DEFAULT NULL,
    `original` varchar(255) DEFAULT NULL,
    `translation` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;
