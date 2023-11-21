ALTER TABLE user_details
ADD CONSTRAINT fk_reporting_to FOREIGN KEY (reporting_to) REFERENCES registration_details(id);