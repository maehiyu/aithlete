
CREATE TABLE IF NOT EXISTS participants (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    role TEXT NOT NULL,
    sports TEXT[] NOT NULL,
    icon_url TEXT
);

CREATE TABLE IF NOT EXISTS chats (
    id TEXT PRIMARY KEY,
    started_at TIMESTAMP NOT NULL,
    last_active_at TIMESTAMP NOT NULL,
    title TEXT,
    participant_ids TEXT[] NOT NULL
);

CREATE TABLE IF NOT EXISTS questions (
    id TEXT PRIMARY KEY,
    chat_id TEXT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    participant_id TEXT NOT NULL REFERENCES participants(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS answers (
    id TEXT PRIMARY KEY,
    chat_id TEXT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    question_id TEXT NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    participant_id TEXT NOT NULL REFERENCES participants(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS attachments (
    id TEXT PRIMARY KEY,
    type TEXT NOT NULL,
    url TEXT NOT NULL,
    thumbnail TEXT,
    pose_id TEXT,
    meta TEXT,
    original_id TEXT,
    question_id TEXT REFERENCES questions(id) ON DELETE CASCADE,
    answer_id TEXT REFERENCES answers(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS posedata (
    id TEXT PRIMARY KEY,
    participant_ids TEXT[] NOT NULL,
    score DOUBLE PRECISION
);

-- テストデータ
INSERT INTO participants (id, name, email, role, sports, icon_url) VALUES
    ('user1', 'ユーザー1', 'user1@example.com', 'user', ARRAY['soccer'], NULL),
    ('coach1', 'コーチ1', 'coach1@example.com', 'coach', ARRAY['tennis'], NULL)
ON CONFLICT (id) DO NOTHING;

INSERT INTO chats (id, started_at, last_active_at, title, participant_ids) VALUES
    ('chat1', NOW(), NOW(), 'テストチャット', ARRAY['user1','coach1'])
ON CONFLICT (id) DO NOTHING;

INSERT INTO questions (id, chat_id, participant_id, content, created_at) VALUES
    ('q1', 'chat1', 'user1', '最初の質問', NOW())
ON CONFLICT (id) DO NOTHING;

INSERT INTO answers (id, chat_id, question_id, participant_id, content, created_at) VALUES
    ('a1', 'chat1', 'q1', 'coach1', '最初の回答', NOW())
ON CONFLICT (id) DO NOTHING;
