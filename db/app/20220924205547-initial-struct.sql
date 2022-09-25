
-- +migrate Up

create table users (
    id varchar(36) default (uuid()) not null,
    first_name text collate utf8mb4_unicode_ci,
    last_name text collate utf8mb4_unicode_ci,
    email text collate utf8mb4_unicode_ci,
    metadata json default null,
    primary key (id)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_unicode_ci;

create table messages (
    id varchar(36) default (uuid()) not null,
    text text not null,
    metadata json default null,
    author varchar(36) default null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (author) references users (id)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_unicode_ci;

create table `read` (
    user varchar(36) not null,
    message varchar(36) not null,
    viewed timestamp default current_timestamp,
    foreign key (user) references users (id),
    foreign key (message) references messages (id)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- +migrate Down
drop table `read`;
drop table messages;
drop table users;