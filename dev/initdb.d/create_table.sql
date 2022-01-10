CREATE TABLE manager (
    `email`   VARCHAR(30)  NOT NULL,
    `pw`      BLOB         NOT NULL,
    `circle`  INT          NOT NULL,
    `pubkey`  BLOB         NOT NULL,
    PRIMARY KEY (circle)
);
CREATE TABLE circle (
    `student_id` VARCHAR(12)  NOT NULL,
    `circle`     INT          NOT NULL,
    `major`      VARCHAR(12)  NOT NULL,
    `name`       BLOB         NOT NULL,
    `year`       INT          NOT NULL,
    `email`      BLOB         NOT NULL,
    `phone`      BLOB         NOT NULL,
    `pay`        BOOL         NOT NULL,
    `status`     INT          NOT NULL,
    PRIMARY KEY (student_id)
);
