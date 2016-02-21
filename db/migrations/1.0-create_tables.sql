-- +migrate Up
CREATE TABLE panic_devices (serial VARCHAR(128) PRIMARY KEY, birth_date TIMESTAMP NOT NULL, death_date TIMESTAMP);

CREATE TABLE elders (id VARCHAR(128) PRIMARY KEY, name VARCHAR(128) NOT NULL, last_name VARCHAR(128) NOT NULL, cellphone VARCHAR(32) NOT NULL, registration_date TIMESTAMP);

CREATE TABLE devices_elders (serial VARCHAR(128), elder_id VARCHAR(128), PRIMARY KEY (serial, elder_id), FOREIGN KEY (serial) REFERENCES panic_devices(serial), FOREIGN KEY (elder_id) REFERENCES elders(id));

CREATE TABLE devices_history (serial VARCHAR(128) NOT NULL, elder_id VARCHAR(128) NOT NULL, attachment_date TIMESTAMP NOT NULL, detachment_date TIMESTAMP, FOREIGN KEY (serial) REFERENCES panic_devices(serial), FOREIGN KEY (elder_id) REFERENCES elders(id));

CREATE TABLE relatives (id VARCHAR(128) PRIMARY KEY, name VARCHAR(128) NOT NULL, last_name VARCHAR(128) NOT NULL, email VARCHAR(128), cellphone VARCHAR(32) NOT NULL, relationship VARCHAR(128));

CREATE TABLE elders_relatives (elder_id VARCHAR(128), relative_id VARCHAR(128), PRIMARY KEY (elder_id, relative_id), FOREIGN KEY (elder_id) REFERENCES elders(id), FOREIGN KEY (relative_id) REFERENCES relatives(id));

CREATE TABLE alerts (id INT AUTO_INCREMENT PRIMARY KEY, serial VARCHAR(128) NOT NULL, elder_id VARCHAR(128) NOT NULL, latitude DECIMAL(10, 7), longitude DECIMAL(10, 7), date TIMESTAMP NOT NULL, FOREIGN KEY (serial) REFERENCES panic_devices(serial), FOREIGN KEY (elder_id) REFERENCES elders(id));

CREATE TABLE sending_results (id INT AUTO_INCREMENT PRIMARY KEY, alert_id INT, relative_id VARCHAR(128), was_successful BOOLEAN, result_message VARCHAR(512), FOREIGN KEY (alert_id) REFERENCES alerts(id), FOREIGN KEY (relative_id) REFERENCES relatives(id));
