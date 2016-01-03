-- +migrate Up
CREATE TABLE panic_devices (serial VARCHAR(128) PRIMARY KEY, birth_date TIMESTAMP NOT NULL, death_date TIMESTAMP);
CREATE TABLE devices_elders (serial VARCHAR(128), elder_id VARCHAR(128), PRIMARY KEY (serial, elder_id));
CREATE TABLE devices_history (serial VARCHAR(128) NOT NULL, elder_id VARCHAR(128) NOT NULL, attachmentDate TIMESTAMP NOT NULL, detachmentDate TIMESTAMP);
