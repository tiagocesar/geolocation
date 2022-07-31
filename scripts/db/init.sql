DROP DATABASE IF EXISTS geolocation;

CREATE DATABASE geolocation;

\c geolocation;

create table location_info
(
    id            serial
        primary key,
    ip_address    inet             not null,
    country_code  varchar(10)      not null,
    country       varchar(50)      not null,
    city          varchar(50)      not null,
    latitude      double precision not null,
    longitude     double precision not null,
    mystery_value varchar(100)
);

alter table location_info
    owner to root;

create unique index location_info_ip_address_uindex
    on location_info (ip_address);