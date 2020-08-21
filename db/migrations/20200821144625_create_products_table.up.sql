create table `products`(  
    `id` int(10) unsigned not null AUTO_INCREMENT,
    `name` varchar(512) default null,
    `price` decimal default null,
    `quantity` int(10) default null,
    `sold` int(10) default null,
    `revenue` decimal default null,
    PRIMARY KEY(`id`),
    UNIQUE KEY `id`(`id`),
    KEY `name`(`name`),
    KEY `price`(`price`),
    KEY `quantity`(`quantity`),
    KEY `sold`(`sold`),
    KEY `revenue`(`revenue`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;