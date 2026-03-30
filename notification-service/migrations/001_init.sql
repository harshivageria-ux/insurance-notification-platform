-- Initial migration: Create all notification admin portal tables

-- Languages table
CREATE TABLE IF NOT EXISTS languages (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    code VARCHAR(10) NOT NULL UNIQUE,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Priorities table
CREATE TABLE IF NOT EXISTS priorities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    level INT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Statuses table
CREATE TABLE IF NOT EXISTS statuses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Schedule Types table
CREATE TABLE IF NOT EXISTS schedule_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Categories table
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Channels table
CREATE TABLE IF NOT EXISTS channels (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    channel_type VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Channel Providers table
CREATE TABLE IF NOT EXISTS channel_providers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    provider_type VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Provider Settings table
CREATE TABLE IF NOT EXISTS provider_settings (
    id SERIAL PRIMARY KEY,
    provider_id INT NOT NULL,
    setting_key VARCHAR(255) NOT NULL,
    setting_value TEXT NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (provider_id) REFERENCES channel_providers(id) ON DELETE CASCADE,
    UNIQUE(provider_id, setting_key)
);

-- Template Groups table
CREATE TABLE IF NOT EXISTS template_groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Templates table
CREATE TABLE IF NOT EXISTS templates (
    id SERIAL PRIMARY KEY,
    template_group_id INT NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    content TEXT NOT NULL,
    variables JSON,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (template_group_id) REFERENCES template_groups(id) ON DELETE CASCADE
);

-- Routing Rules table
CREATE TABLE IF NOT EXISTS routing_rules (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    condition JSON NOT NULL,
    target_channel INT NOT NULL,
    priority INT DEFAULT 0,
    is_active BOOLEAN DEFAULT true,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (target_channel) REFERENCES channels(id) ON DELETE CASCADE
);

-- Create indexes for better query performance
CREATE INDEX idx_languages_status ON languages(status);
CREATE INDEX idx_priorities_status ON priorities(status);
CREATE INDEX idx_channels_is_active ON channels(is_active);
CREATE INDEX idx_providers_is_active ON channel_providers(is_active);
CREATE INDEX idx_templates_status ON templates(status);
CREATE INDEX idx_routing_rules_is_active ON routing_rules(is_active);

-- Insert default languages
INSERT INTO languages (name, code, status) VALUES 
    ('English', 'EN', 'Active'),
    ('Hindi', 'HI', 'Active'),
    ('Spanish', 'ES', 'Active')
ON CONFLICT (code) DO NOTHING;

-- ============================
-- Admin portal master tables
-- (used by Go repositories)
-- ============================

-- Languages (master)
CREATE TABLE IF NOT EXISTS languages_master (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    code VARCHAR(10) NOT NULL UNIQUE,
    created_by VARCHAR(100) NOT NULL,
    updated_by VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT false,
    version INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Priorities (master)
CREATE TABLE IF NOT EXISTS notification_priority_master (
    priority_id SMALLSERIAL PRIMARY KEY,
    priority_code VARCHAR(30) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT false,
    version INT NOT NULL DEFAULT 0
);

-- Statuses (master)
CREATE TABLE IF NOT EXISTS notification_statuses_master (
    status_id SMALLSERIAL PRIMARY KEY,
    status_code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    is_final BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT false,
    version INT NOT NULL DEFAULT 0
);

-- Schedule Types (master)
CREATE TABLE IF NOT EXISTS notification_schedule_type_master (
    schedule_type_id SMALLSERIAL PRIMARY KEY,
    schedule_code VARCHAR(30) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT false,
    version INT NOT NULL DEFAULT 0
);

-- Categories (master) - NO UUID, auto-increment INT
CREATE TABLE IF NOT EXISTS notification_categories_master (
    id SERIAL PRIMARY KEY,
    code VARCHAR(30) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT false,
    version INT NOT NULL DEFAULT 0
);

-- Channels (master)
CREATE TABLE IF NOT EXISTS notification_channels_master (
    id SERIAL PRIMARY KEY,
    code VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    version INT NOT NULL DEFAULT 0
);

-- Channel Providers (master)
CREATE TABLE IF NOT EXISTS channel_providers_master (
    id SERIAL PRIMARY KEY,
    channel_id INT NOT NULL REFERENCES notification_channels_master(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(30) NOT NULL,
    priority INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    version INT NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_channel_providers_channel_id ON channel_providers_master(channel_id);

-- Provider settings (master)
CREATE TABLE IF NOT EXISTS channel_provider_settings_master (
    id SERIAL PRIMARY KEY,
    provider_id INT NOT NULL REFERENCES channel_providers_master(id) ON DELETE CASCADE,
    setting_key VARCHAR(255) NOT NULL,
    setting_value TEXT NOT NULL,
    is_sensitive BOOLEAN NOT NULL DEFAULT false,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    version INT NOT NULL DEFAULT 0,
    UNIQUE(provider_id, setting_key)
);

-- Template Groups (master)
CREATE TABLE IF NOT EXISTS template_groups_master (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    category_id INT NOT NULL REFERENCES notification_categories_master(id) ON DELETE CASCADE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    version INT NOT NULL DEFAULT 0
);

-- Templates (master)
CREATE TABLE IF NOT EXISTS notification_templates_master (
    id SERIAL PRIMARY KEY,
    template_group_id INT NOT NULL REFERENCES template_groups_master(id) ON DELETE CASCADE,
    channel_id INT NOT NULL REFERENCES notification_channels_master(id) ON DELETE CASCADE,
    language_id BIGINT NOT NULL REFERENCES languages_master(id) ON DELETE CASCADE,
    title_template TEXT,
    message_template TEXT NOT NULL,
    has_variables BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    version INT NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_templates_group ON notification_templates_master(template_group_id);

-- Routing Rules (master)
CREATE TABLE IF NOT EXISTS provider_routing_rules_master (
    id SERIAL PRIMARY KEY,
    template_group_id INT NOT NULL REFERENCES template_groups_master(id) ON DELETE CASCADE,
    channel_id INT NOT NULL REFERENCES notification_channels_master(id) ON DELETE CASCADE,
    preferred_provider_id INT NOT NULL REFERENCES channel_providers_master(id) ON DELETE CASCADE,
    fallback_provider_id INT REFERENCES channel_providers_master(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    version INT NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_routing_rules_channel ON provider_routing_rules_master(channel_id);

-- Insert default languages for master tables
INSERT INTO languages_master (name, code, created_by, updated_by, is_active, version, created_at, updated_at)
VALUES
  ('English', 'EN', 'system', 'system', true, 1, NOW(), NOW()),
  ('Hindi', 'HI', 'system', 'system', true, 1, NOW(), NOW()),
  ('Spanish', 'ES', 'system', 'system', true, 1, NOW(), NOW())
ON CONFLICT (code) DO NOTHING;
