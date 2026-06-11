CREATE TABLE IF NOT EXISTS ai_pronunciation_results (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    child_id UUID NOT NULL REFERENCES children(id) ON DELETE CASCADE,
    audio_url TEXT NOT NULL,
    transcription TEXT,
    score DECIMAL(5, 2),
    feedback TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
