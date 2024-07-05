-- Modify Max Connection
ALTER SYSTEM SET max_connections = 9999;

-- CREATE TABLES
CREATE TABLE roles(
    role_id     SERIAL      NOT NULL PRIMARY KEY,
    role_name   VARCHAR(10) NOT NULL
);

CREATE TABLE users(
    user_id     SERIAL          NOT NULL    PRIMARY KEY,
    email       VARCHAR(60)     NOT NULL,
    full_name   VARCHAR(60)     NOT NULL,
    password    VARCHAR(300)    NOT NULL,
    role_id     INTEGER         NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(role_id)
);

CREATE TABLE user_attachment(
    user_id     INTEGER         NOT NULL,
    file_url    VARCHAR(300)    NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE insurance_info(
    insurance_id    SERIAL          NOT NULL    PRIMARY KEY,
    insurance_name  VARCHAR(50)     NOT NULL,
    file_url        VARCHAR(300)    NULL
);

CREATE TABLE clinic(
    clinic_id       SERIAL          NOT NULL    PRIMARY KEY,
    clinic_name     VARCHAR(50)     NOT NULL,
    clinic_image    VARCHAR(300)    NULL
);

CREATE TABLE language(
    language_id     SERIAL          NOT NULL    PRIMARY KEY,
    language_name   VARCHAR(50)     NOT NULL
);

CREATE TABLE work_location(
    work_location_id    SERIAL          NOT NULL    PRIMARY KEY,
    work_location_name  VARCHAR(50)     NOT NULL
);

CREATE TABLE doctor_profile(
    doctor_id       SERIAL          NOT NULL    PRIMARY KEY,
    docker_summary  VARCHAR(500)    NULL,
    exp_year        INTEGER         NOT NULL,
    clinic_id       INTEGER         NOT NULL,
    edu_location    VARCHAR(500)    NOT NULL,
    degree          VARCHAR(500)    NOT NULL,
    work_location   INTEGER         NULL,
    language        INTEGER         NULL,
    description     VARCHAR(500)    NULL,
    FOREIGN KEY (clinic_id) REFERENCES clinic(clinic_id),
    FOREIGN KEY (language) REFERENCES language(language_id),
    FOREIGN KEY (work_location) REFERENCES work_location(work_location_id)
);

CREATE TABLE rating(
    rating_id       SERIAL          NOT NULL    PRIMARY KEY,
    star            INTEGER         NULL,
    description     VARCHAR(500)    NULL,
    doctor_id       INTEGER         NOT NULL,
    FOREIGN KEY (doctor_id) REFERENCES doctor_profile(doctor_id)
);

CREATE TABLE surgery_info(
    surgery_id      SERIAL          NOT NULL    PRIMARY KEY,
    surgery_name    VARCHAR(50)     NOT NULL,
    surgery_addr    VARCHAR(50)     NOT NULL,
    surgery_time    VARCHAR(50)     NOT NULL,
    doctor          INTEGER         NOT NULL,
    insurance       INTEGER         NOT NULL,
    FOREIGN KEY (doctor) REFERENCES doctor_profile(doctor_id),
    FOREIGN KEY (insurance) REFERENCES insurance_info(insurance_id)
);
CREATE TABLE surgery_attachment(
    surgery_id      INTEGER         NOT NULL,
    file_url        VARCHAR(300)    NULL,
    FOREIGN KEY (surgery_id) REFERENCES surgery_info(surgery_id)
);
CREATE TABLE payment(
    surgery_id      INTEGER         NOT NULL,
    payment_name    VARCHAR(50)     NOT NULL,
    file_url        VARCHAR(300)    NULL,
    FOREIGN KEY (surgery_id) REFERENCES surgery_info(surgery_id)
);