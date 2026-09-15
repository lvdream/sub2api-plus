CREATE TABLE IF NOT EXISTS claude_code_account_identities (
    account_id BIGINT PRIMARY KEY REFERENCES accounts(id) ON DELETE CASCADE,
    device_id CHAR(64) NOT NULL,
    default_session_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE claude_code_account_identities IS
    'Private stable Claude Code device and fallback session identity per upstream account.';
