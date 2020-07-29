CREATE TABLE risks_controls (
    id BIGSERIAL PRIMARY KEY,
    risk_id BIGINT NOT NULL REFERENCES risks(id) ON DELETE CASCADE,
    control_id BIGINT NOT NULL REFERENCES controls(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX ON risks_controls(risk_id, control_id);
CREATE UNIQUE INDEX ON risks_controls(control_id, risk_id);
