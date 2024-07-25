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
    file_url    VARCHAR(3000)   NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE insurance(
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

CREATE TABLE doctor (
    doctor_id       SERIAL          NOT NULL    PRIMARY KEY,
    doctor_name     VARCHAR(500)    NOT NULL,
    doctor_summary  VARCHAR(500)    NULL,
    exp_year        INTEGER         NOT NULL,
    edu_location    VARCHAR(500)    NOT NULL,
    degree          VARCHAR(500)    NOT NULL,
    description     VARCHAR(500)    NULL,
    file_url        VARCHAR(500)    NULL,
    clinic_id       INTEGER         NOT NULL,
    FOREIGN KEY (clinic_id) REFERENCES clinic(clinic_id)
);

CREATE TABLE doctor_profile(
    doctor_id       SERIAL          NOT NULL,
    work_location   INTEGER         NOT NULL,
    language        INTEGER         NOT NULL,
    FOREIGN KEY (doctor_id) REFERENCES doctor(doctor_id),
    FOREIGN KEY (work_location) REFERENCES work_location(work_location_id),
    FOREIGN KEY (language) REFERENCES language(language_id)
);

CREATE TABLE rating(
    rating_id       SERIAL          NOT NULL    PRIMARY KEY,
    star            INTEGER         NULL,
    description     VARCHAR(500)    NULL,
    doctor_id       INTEGER         NOT NULL,
    FOREIGN KEY (doctor_id) REFERENCES doctor(doctor_id)
);

CREATE TABLE surgery(
    surgery_id      SERIAL          NOT NULL    PRIMARY KEY,
    surgery_name    VARCHAR(50)     NOT NULL,
    surgery_addr    VARCHAR(50)     NOT NULL,
    surgery_time    VARCHAR(50)     NOT NULL,
    insurance       INTEGER         NOT NULL,
    FOREIGN KEY (insurance) REFERENCES insurance(insurance_id)
);

CREATE TABLE surgery_attachment(
    surgery_id      INTEGER         NOT NULL,
    file_url        VARCHAR(300)    NULL,
    FOREIGN KEY (surgery_id) REFERENCES surgery(surgery_id)
);

CREATE TABLE surgery_doctor(
    doctor_id       INTEGER         NOT NULL,
    surgery_id      INTEGER         NOT NULL,
    FOREIGN KEY (doctor_id) REFERENCES doctor(doctor_id)
);

CREATE TABLE payment(
    payment_id      SERIAL          NOT NULL    PRIMARY KEY,
    payment_name    VARCHAR(50)     NOT NULL,
    file_url        VARCHAR(300)    NULL
);

CREATE TABLE surgery_payment(
    surgery_id      INTEGER         NOT NULL,
    payment_id      INTEGER         NOT NULL,
    FOREIGN KEY (surgery_id) REFERENCES surgery(surgery_id),
    FOREIGN KEY(payment_id) REFERENCES payment(payment_id)
);

CREATE TABLE insurance_surgery(
    surgery_id      INTEGER         NOT NULL,
    insurance_id    INTEGER         NOT NULL,
    FOREIGN KEY (surgery_id) REFERENCES surgery(surgery_id),
    FOREIGN KEY (insurance_id) REFERENCES insurance(insurance_id)
);

INSERT INTO roles(role_name) VALUES ('User'), ('Admin'),('Doctor');

INSERT INTO language(language_name) VALUES ('Vietnamese') , ('English');

INSERT INTO clinic(clinic_name, clinic_image) VALUES ('Nhi Khoa', ''), ('Đa khoa', ''); 