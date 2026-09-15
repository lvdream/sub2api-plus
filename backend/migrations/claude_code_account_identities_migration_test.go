package migrations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClaudeCodeAccountIdentitiesMigration(t *testing.T) {
	content, err := FS.ReadFile("266_claude_code_account_identities.sql")
	require.NoError(t, err)
	migration := string(content)
	require.Contains(t, migration, "CREATE TABLE IF NOT EXISTS claude_code_account_identities")
	require.Contains(t, migration, "account_id BIGINT PRIMARY KEY REFERENCES accounts(id) ON DELETE CASCADE")
	require.Contains(t, migration, "device_id CHAR(64) NOT NULL")
	require.Contains(t, migration, "default_session_id UUID NOT NULL")
}
