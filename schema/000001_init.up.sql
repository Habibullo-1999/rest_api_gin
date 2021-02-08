

CREATE TABLE users 
(
    id              SERIAL              NOT NULL UNIQUE,
    name            VARCHAR(255)        NOT NULL,
    username        VARCHAR(255)        NOT NULL UNIQUE,
    password_hash   VARCHAR(255)        NOT NULL
);

CREATE TABLE note_list
(
    id          SERIAL              NOT NULL UNIQUE,
    title       VARCHAR(255)        NOT NULL,
    content     VARCHAR(255)
);

CREATE TABLE users_list
(
    id          SERIAL              NOT NULL unique,
    user_id     INT references users(id) ON DELETE cascade NOT NULL,
    list_id     INT references note_list(id) on delete cascade NOT NULL
);

CREATE TABLE note_items
(
    id          SERIAL              NOT NULL unique,
    title       VARCHAR(255)        NOT NULL,
    content     VARCHAR(255)
);

CREATE TABLE lists_item
(
    id          SERIAL              NOT NULL unique,
    item_id     INT references note_items(id) on delete cascade not null,
    list_id     INT references note_list(id) on delete cascade not null
);

