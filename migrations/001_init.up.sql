-- 001_init: Create aihelper schema and tables

CREATE SCHEMA IF NOT EXISTS aihelper;

CREATE TABLE IF NOT EXISTS aihelper.templates (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    category TEXT NOT NULL DEFAULT '通用',
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS aihelper.prompts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    category TEXT DEFAULT '通用',
    tags TEXT[] DEFAULT '{}',
    variables TEXT[] DEFAULT '{}',
    is_favorite BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS aihelper.user_settings (
    user_id BIGINT PRIMARY KEY,
    api_key TEXT DEFAULT '',
    api_base TEXT DEFAULT '',
    model TEXT DEFAULT 'gpt-4o-mini',
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
