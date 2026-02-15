create extension if not exists pg_trgm;

create table if not exists cockpits
(
    id         serial primary key,
    name       varchar(100) not null unique,
    is_default boolean   default false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create index idx_cockpits_name_gin on cockpits using gin (name gin_trgm_ops);

create table if not exists games
(
    id         serial primary key,
    name       varchar(100) not null unique,
    image      varchar,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create index idx_games_name_gin on games using gin (name gin_trgm_ops);

create table if not exists pedals
(
    id         serial primary key,
    name       varchar(100) not null unique,
    is_default boolean   default false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create index idx_pedals_name_gin on pedals using gin (name gin_trgm_ops);

create table if not exists tracks
(
    id          serial primary key,
    name        varchar(100) not null unique,
    image       varchar,
    description text,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create index idx_tracks_name_gin on tracks using gin (name gin_trgm_ops);

create table if not exists wheels
(
    id         serial primary key,
    name       varchar(100) not null unique,
    is_default boolean   default false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create index idx_wheels_name_gin on wheels using gin (name gin_trgm_ops);

create table if not exists laps
(
    id                     serial primary key,
    car_id                 int              not null,
    track_id               int              not null,
    game_id                int              not null,
    wheel_id               int              not null,
    cockpit_id             int              not null,
    pedals_id              int              not null,
    time                   double precision not null,
    is_clear               boolean   default false,
    has_significant_errors boolean   default false,
    created_at             TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    foreign key (car_id) references cars (id) on delete cascade,
    foreign key (track_id) references tracks (id) on delete cascade,
    foreign key (game_id) references games (id) on delete cascade,
    foreign key (wheel_id) references wheels (id) on delete cascade,
    foreign key (cockpit_id) references cockpits (id) on delete cascade,
    foreign key (pedals_id) references pedals (id) on delete cascade
);

create index idx_laps_sort on laps (track_id, car_id, time);
