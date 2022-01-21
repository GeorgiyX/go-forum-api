DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA public;
CREATE
EXTENSION IF NOT EXISTS CITEXT;

CREATE
UNLOGGED TABLE users
(
    nickname CITEXT UNIQUE PRIMARY KEY NOT NULL,
    email    CITEXT UNIQUE             NOT NULL,
    fullname TEXT                      NOT NULL,
    about    TEXT                      NOT NULL
);

CREATE
UNLOGGED TABLE forums
(
    id      SERIAL UNIQUE                                                          NOT NULL,
    slug    CITEXT UNIQUE PRIMARY KEY                                              NOT NULL,
    title   TEXT                                                                   NOT NULL,
    "user"  CITEXT REFERENCES users (nickname) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    posts   INT DEFAULT 0,
    threads INT DEFAULT 0
);

CREATE
UNLOGGED TABLE threads
(
    id      SERIAL UNIQUE PRIMARY KEY                                              NOT NULL,
    slug    CITEXT UNIQUE,
    author  CITEXT REFERENCES users (nickname) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    forum   CITEXT REFERENCES forums (slug) ON UPDATE CASCADE ON DELETE CASCADE    NOT NULL,
    title   TEXT                                                                   NOT NULL,
    message TEXT                                                                   NOT NULL,
    created TIMESTAMP WITH TIME ZONE DEFAULT NOW()                                 NOT NULL,
    votes   INT                      DEFAULT 0
);

CREATE
UNLOGGED TABLE votes
(
    nickname CITEXT REFERENCES users (nickname) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    thread   INT REFERENCES threads (id) ON UPDATE CASCADE ON DELETE CASCADE        NOT NULL,
    value    INT                                                                    NOT NULL,

    PRIMARY KEY (thread, nickname),
    UNIQUE (thread, nickname)
);

CREATE
UNLOGGED TABLE posts
(
    id       BIGSERIAL UNIQUE PRIMARY KEY                                              NOT NULL,
    parent   INT                      DEFAULT NULL,
    author   CITEXT REFERENCES users (nickname) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    forum    CITEXT REFERENCES forums (slug) ON UPDATE CASCADE ON DELETE CASCADE    NOT NULL,
    thread   INT REFERENCES threads (id) ON UPDATE CASCADE ON DELETE CASCADE        NOT NULL,
    created  TIMESTAMP WITH TIME ZONE DEFAULT NOW()                                 NOT NULL,
    isEdited BOOLEAN                  DEFAULT false                                 NOT NULL,
    message  TEXT                                                                   NOT NULL,
    path     BIGINT[]                                                               NOT NULL
);

CREATE
UNLOGGED TABLE IF NOT EXISTS forum_users
(
    forum    CITEXT REFERENCES forums (slug) ON UPDATE CASCADE ON DELETE CASCADE    NOT NULL,
    nickname CITEXT REFERENCES users (nickname) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,

    PRIMARY KEY (forum, nickname),
    UNIQUE (forum, nickname)
);

--Триггеры на голосование--
--1. Insert votes
CREATE OR REPLACE FUNCTION on_insert_vote() RETURNS TRIGGER AS $$
BEGIN
    UPDATE threads SET votes = votes + NEW.value WHERE id = NEW.thread;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER recount_votes_insert
    AFTER INSERT
    ON votes
    FOR EACH ROW
EXECUTE PROCEDURE on_insert_vote();

--2. Update votes
CREATE OR REPLACE FUNCTION on_update_vote() RETURNS TRIGGER AS $$
BEGIN
    UPDATE threads SET votes = votes - OLD.value + NEW.value WHERE id = NEW.thread;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER recount_votes_update
    AFTER UPDATE
    ON votes
    FOR EACH ROW
EXECUTE PROCEDURE on_update_vote();

--Триггер на path--
CREATE OR REPLACE FUNCTION add_path_to_post() RETURNS TRIGGER AS $$
DECLARE
    parent_path BIGINT[];
BEGIN
    IF NEW.parent IS NULL THEN
        NEW.path := NEW.path || NEW.id;
    ELSE
        SELECT path FROM posts WHERE id = NEW.parent AND thread = NEW.thread INTO parent_path;

        IF (COALESCE(ARRAY_LENGTH(parent_path, 1), NULL) IS NULL) THEN
            RAISE EXCEPTION 'add_path_to_post: parent not found';
        END IF;

        NEW.path := NEW.path || parent_path || NEW.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER add_path_to_post
    BEFORE INSERT
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE add_path_to_post();
