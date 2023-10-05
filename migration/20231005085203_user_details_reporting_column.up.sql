ALTER TABLE user_details 
ADD COLUMN reporting_to varchar(36);

-- ALTER TABLE user_details
-- ADD CONSTRAINT fk_reporting_to FOREIGN KEY (reporting_to) REFERENCES registration_details(id);