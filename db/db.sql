DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA public;
CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE UNLOGGED TABLE users
(
    nickname CITEXT UNIQUE PRIMARY KEY NOT NULL,
    email    CITEXT UNIQUE             NOT NULL,
    fullname TEXT                      NOT NULL,
    about    TEXT                      NOT NULL
);

CREATE UNLOGGED TABLE forums
(
    slug    CITEXT UNIQUE PRIMARY KEY                                              NOT NULL,
    title   TEXT                                                                   NOT NULL,
    "user"  CITEXT REFERENCES users (nickname) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    posts   INT DEFAULT 0,
    threads INT DEFAULT 0
);

CREATE UNLOGGED TABLE threads
(
    id      SERIAL UNIQUE PRIMARY KEY                                              NOT NULL,
    slug    CITEXT UNIQUE                                                          NOT NULL,
    author  CITEXT REFERENCES users (nickname) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    forum   CITEXT REFERENCES forums (slug) ON UPDATE CASCADE ON DELETE CASCADE    NOT NULL,
    title   TEXT                                                                   NOT NULL,
    message TEXT,
    created TIMESTAMP WITH TIME ZONE DEFAULT NOW()                                 NOT NULL,
    votes   INT                      DEFAULT 0
);

CREATE UNLOGGED TABLE votes
(
    nickname CITEXT REFERENCES users (nickname) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    thread   INT REFERENCES threads (id) ON UPDATE CASCADE ON DELETE CASCADE        NOT NULL,
    value    INT                                                                    NOT NULL,

    PRIMARY KEY (thread, nickname),
    UNIQUE (thread, nickname)
);

CREATE UNLOGGED TABLE posts
(
    id       SERIAL UNIQUE PRIMARY KEY                                              NOT NULL,
    parent   INT                      DEFAULT NULL,
    author   CITEXT REFERENCES users (nickname) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    forum    CITEXT REFERENCES forums (slug) ON UPDATE CASCADE ON DELETE CASCADE    NOT NULL,
    thread   INT REFERENCES threads (id) ON UPDATE CASCADE ON DELETE CASCADE        NOT NULL,
    created  TIMESTAMP WITH TIME ZONE DEFAULT NOW()                                 NOT NULL,
    isEdited BOOLEAN                  DEFAULT false                                 NOT NULL,
    message  TEXT                                                                   NOT NULL
);

CREATE UNLOGGED TABLE IF NOT EXISTS forum_users
(
    forum    CITEXT REFERENCES forums (slug) ON UPDATE CASCADE ON DELETE CASCADE    NOT NULL,
    nickname CITEXT REFERENCES users (nickname) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,

    PRIMARY KEY (forum, nickname),
    UNIQUE (forum, nickname)
);
