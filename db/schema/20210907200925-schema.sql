-- +migrate Up
create schema if not exists hermes CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
-- create user if not exists hermes@'localhost' identified by hermespw;
create user if not exists hermes@'%' identified by 'hermespw';
-- grant all on hermes.* to hermes@'localhost';
grant all on hermes.* to hermes@'%';

-- +migrate Down
revoke all on hermes.*  from hermes@'%';
-- revoke all on hermes.*  from hermes@'localhost';
drop user if exists hermes@'%';
-- drop user if exists hermes@'localhost';
drop schema if exists hermes;