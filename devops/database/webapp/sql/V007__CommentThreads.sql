CREATE TABLE comment_threads (
    id BIGSERIAL PRIMARY KEY   
);

CREATE TABLE comments (
    id BIGSERIAL PRIMARY KEY,
    thread_id BIGINT NOT NULL REFERENCES comment_threads(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    post_time TIMESTAMPTZ NOT NULL,
    content TEXT
);

CREATE INDEX ON comments(thread_id);
