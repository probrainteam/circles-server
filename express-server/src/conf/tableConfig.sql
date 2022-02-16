CREATE TABLE manager (
    `email`   VARCHAR(30)  NOT NULL,
    `pw`      BLOB         NOT NULL,
    `circle`  INT          default 123,
    `pubkey`  BLOB         NOT NULL,
    `auth`    BOOL         default false,
    PRIMARY KEY (email)
);
CREATE TABLE probrain (
    `student_id` VARCHAR(12)  NOT NULL,
    `major`      VARCHAR(12)  NOT NULL,
    `name`       BLOB         NOT NULL,
    `year`       INT          NOT NULL,
    `email`      BLOB         NOT NULL,
    `phone`      BLOB         NOT NULL,
    `paid`       BOOL         NOT NULL,
    `status`     INT          NOT NULL,
    PRIMARY KEY (student_id)
);
CREATE TABLE grow (
    `student_id` VARCHAR(12)  NOT NULL,
    `major`      VARCHAR(12)  NOT NULL,
    `name`       BLOB         NOT NULL,
    `year`       INT          NOT NULL,
    `email`      BLOB         NOT NULL,
    `phone`      BLOB         NOT NULL,
    `paid`       BOOL         NOT NULL,
    `status`     INT          NOT NULL,
    PRIMARY KEY (student_id)
);
CREATE TABLE argos (
    `student_id` VARCHAR(12)  NOT NULL,
    `major`      VARCHAR(12)  NOT NULL,
    `name`       BLOB         NOT NULL,
    `year`       INT          NOT NULL,
    `email`      BLOB         NOT NULL,
    `phone`      BLOB         NOT NULL,
    `paid`       BOOL         NOT NULL,
    `status`     INT          NOT NULL,
    PRIMARY KEY (student_id)
);
CREATE TABLE adm2n (
    `student_id` VARCHAR(12)  NOT NULL,
    `major`      VARCHAR(12)  NOT NULL,
    `name`       BLOB         NOT NULL,
    `year`       INT          NOT NULL,
    `email`      BLOB         NOT NULL,
    `phone`      BLOB         NOT NULL,
    `paid`       BOOL         NOT NULL,
    `status`     INT          NOT NULL,
    PRIMARY KEY (student_id)
);
CREATE TABLE ana (
    `student_id` VARCHAR(12)  NOT NULL,
    `major`      VARCHAR(12)  NOT NULL,
    `name`       BLOB         NOT NULL,
    `year`       INT          NOT NULL,
    `email`      BLOB         NOT NULL,
    `phone`      BLOB         NOT NULL,
    `paid`       BOOL         NOT NULL,
    `status`     INT          NOT NULL,
    PRIMARY KEY (student_id)
);
CREATE TABLE motion (
    `student_id` VARCHAR(12)  NOT NULL,
    `major`      VARCHAR(12)  NOT NULL,
    `name`       BLOB         NOT NULL,
    `year`       INT          NOT NULL,
    `email`      BLOB         NOT NULL,
    `phone`      BLOB         NOT NULL,
    `paid`       BOOL         NOT NULL,
    `status`     INT          NOT NULL,
    PRIMARY KEY (student_id)
);
CREATE TABLE spg (
    `student_id` VARCHAR(12)  NOT NULL,
    `major`      VARCHAR(12)  NOT NULL,
    `name`       BLOB         NOT NULL,
    `year`       INT          NOT NULL,
    `email`      BLOB         NOT NULL,
    `phone`      BLOB         NOT NULL,
    `paid`       BOOL         NOT NULL,
    `status`     INT          NOT NULL,
    PRIMARY KEY (student_id)
);
CREATE TABLE pai (
    `student_id` VARCHAR(12)  NOT NULL,
    `major`      VARCHAR(12)  NOT NULL,
    `name`       BLOB         NOT NULL,
    `year`       INT          NOT NULL,
    `email`      BLOB         NOT NULL,
    `phone`      BLOB         NOT NULL,
    `paid`        BOOL         NOT NULL,
    `status`     INT          NOT NULL,
    PRIMARY KEY (student_id)
);