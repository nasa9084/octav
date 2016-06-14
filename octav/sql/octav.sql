CREATE TABLE users (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    eid CHAR(64) CHARACTER SET latin1 NOT NULL,
    first_name TEXT,
    last_name TEXT,
    nickname  CHAR(128) NOT NULL,
    email TEXT,
    auth_via CHAR(16) NOT NULL, /* github, facebook, twitter */
    auth_user_id TEXT NOT NULL, /* ID in the auth provider */
    avatar_url TEXT,
    is_admin TINYINT(1) NOT NULL DEFAULT 0,
    tshirt_size CHAR(4) CHARACTER SET latin1,
    created_on DATETIME NOT NULL,
    modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY(eid),
    UNIQUE KEY(nickname),
    UNIQUE KEY(auth_via, auth_user_id(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE venues (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    eid CHAR(64) CHARACTER SET latin1 NOT NULL,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    latitude DECIMAL(10,7),
    longitude DECIMAL(10,7),
    created_on DATETIME NOT NULL,
    modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY(eid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE rooms (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    venue_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    eid CHAR(64) CHARACTER SET latin1 NOT NULL,
    name TEXT NOT NULL,
    capacity INT UNSIGNED NOT NULL DEFAULT 0,
    created_on DATETIME NOT NULL,
    modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (eid),
    KEY(venue_id, eid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE conference_series (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    eid CHAR(64) CHARACTER SET latin1 NOT NULL,
    slug TEXT NOT NULL,
    created_on DATETIME NOT NULL,
    modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (eid),
    UNIQUE KEY (slug(255))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    

CREATE TABLE conferences (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    eid CHAR(64) CHARACTER SET latin1 NOT NULL,
    series_id CHAR(64) CHARACER SET latin1, -- may be null
    slug TEXT NOT NULL,
    title TEXT NOT NULL,
    sub_title TEXT,
    created_by CHAR(64) CHARACTER SET latin1 NOT NULL,
    created_on DATETIME NOT NULL,
    modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY(eid),
    UNIQUE KEY(series_id, slug(255))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE conference_dates (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    conference_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    date DATE NOT NULL,
    open TIME,
    close TIME,
    KEY(date),
    FOREIGN KEY (conference_id) REFERENCES conferences(eid) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE conference_administrators (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    conference_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    user_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    created_on DATETIME NOT NULL,
    modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (conference_id) REFERENCES conferences(eid) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(eid) ON DELETE CASCADE,
    UNIQUE KEY(conference_id, user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE conference_venues (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    conference_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    venue_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    created_on DATETIME NOT NULL,
    modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (conference_id) REFERENCES conferences(eid) ON DELETE CASCADE,
    FOREIGN KEY (venue_id) REFERENCES venues(eid) ON DELETE CASCADE,
    UNIQUE KEY(conference_id, venue_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE sessions (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    eid CHAR(64) CHARACTER SET latin1 NOT NULL,
    conference_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    room_id CHAR(64) CHARACTER SET latin1,
    speaker_id CHAR(64) NOT NULL,
    title TEXT NOT NULL,
    abstract TEXT,
    memo TEXT,
    starts_on DATETIME,
    duration INTEGER UNSIGNED,
    material_level TEXT,
    tags TEXT,
    category TEXT,
    spoken_language TEXT,
    slide_language TEXT,
    slide_subtitles TEXT,
    slide_url TEXT,
    video_url TEXT,
    photo_permission CHAR(16) NOT NULL DEFAULT "allow",
    video_permission CHAR(16) NOT NULL DEFAULT "allow",
    has_interpretation TINYINT(1) NOT NULL DEFAULT 0,
    status CHAR(16) NOT NULL DEFAULT "pending",
    sort_order INTEGER NOT NULL DEFAULT 0,
    confirmed TINYINT(1) NOT NULL DEFAULT 0,
    created_on DATETIME NOT NULL,
    modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (eid),
    KEY(eid, conference_id, room_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE localized_strings (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    parent_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    parent_type CHAR(64) CHARACTER SET latin1 NOT NULL,
    name CHAR(128) BINARY NOT NULL,
    language CHAR(32) BINARY NOT NULL,
    localized TEXT NOT NULL,
    KEY (parent_id, parent_type, name, language)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE questions (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    eid CHAR(64) CHARACTER SET latin1 NOT NULL,
    session_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    user_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    body TEXT NOT NULL,
    KEY (eid, session_id, user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- session survey responses.
CREATE TABLE session_survey_responses (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    eid CHAR(64) CHARACTER SET latin1 NOT NULL,
    session_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    user_id CHAR(64) CHARACTER SET latin1 NOT NULL,
    user_prior_knowledge SMALLINT UNSIGNED NOT NULL DEFAULT 0,
    speaker_knowledge SMALLINT UNSIGNED NOT NULL DEFAULT 0,
    speaker_presentation SMALLINT UNSIGNED NOT NULL DEFAULT 0,
    material_quality SMALLINT UNSIGNED NOT NULL DEFAULT 0,
    overall_rating SMALLINT UNSIGNED NOT NULL DEFAULT 0,
    comment_good TEXT,
    comment_improvement TEXT,
    created_on DATETIME NOT NULL,
    modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (eid),
    KEY(eid, session_id, user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- client stores data about clients that use our API
CREATE TABLE clients (
    oid INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    eid CHAR(64) CHARACTER SET latin1 NOT NULL, -- client ID
    secret CHAR(64) BINARY CHARACTER SET latin1 NOT NULL,
    name TEXT NOT NULL, -- name of the client
    created_on DATETIME NOT NULL,
    modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (eid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
