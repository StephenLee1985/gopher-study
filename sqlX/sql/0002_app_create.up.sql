use fang;
create table if not exists application (
    id bigint unsigned primary key auto_increment,
    uid varchar(64) not null,
    cid varchar(64) not null,
    name varchar(64) not null,
    instances bigint not null,
    status tinyint(1) not null,
    json text,
    created timestamp
    ) character set utf8 collate utf8_general_ci;
