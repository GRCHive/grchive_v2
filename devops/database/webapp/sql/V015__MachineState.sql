CREATE TABLE machine_state (
    id BIGSERIAL PRIMARY KEY,
    engagement_id BIGINT NOT NULL REFERENCES engagements(id) ON DELETE CASCADE,
    unique_id VARCHAR NOT NULL,
    hypervisor TEXT,
    operating_system TEXT,
    UNIQUE(engagement_id, unique_id)
);

ALTER TABLE inventory_servers
DROP COLUMN operating_system;

ALTER TABLE inventory_servers
DROP COLUMN hypervisor;

ALTER TABLE inventory_desktops
DROP COLUMN operating_system;

ALTER TABLE inventory_laptops
DROP COLUMN operating_system;

ALTER TABLE inventory_mobile
DROP COLUMN operating_system;

ALTER TABLE inventory_embedded
DROP COLUMN operating_system;

ALTER TABLE inventory_servers
ADD COLUMN state_id BIGINT REFERENCES machine_state(id) ON DELETE CASCADE;

ALTER TABLE inventory_desktops
ADD COLUMN state_id BIGINT REFERENCES machine_state(id) ON DELETE CASCADE;

ALTER TABLE inventory_laptops
ADD COLUMN state_id BIGINT REFERENCES machine_state(id) ON DELETE CASCADE;

ALTER TABLE inventory_mobile
ADD COLUMN state_id BIGINT REFERENCES machine_state(id) ON DELETE CASCADE;

ALTER TABLE inventory_embedded
ADD COLUMN state_id BIGINT REFERENCES machine_state(id) ON DELETE CASCADE;
