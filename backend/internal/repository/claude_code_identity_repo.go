package repository

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/LuckyKuang/sub2api-plus/internal/service"
	"github.com/google/uuid"
)

func (r *accountRepository) GetOrCreateClaudeCodeIdentity(ctx context.Context, accountID int64) (*service.ClaudeCodeAccountIdentity, error) {
	deviceID, err := randomClaudeCodeDeviceID()
	if err != nil {
		return nil, err
	}
	defaultSessionID := uuid.NewString()

	identity := &service.ClaudeCodeAccountIdentity{}
	err = scanSingleRow(ctx, r.sql, `
		INSERT INTO claude_code_account_identities (account_id, device_id, default_session_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (account_id) DO UPDATE SET account_id = EXCLUDED.account_id
		RETURNING device_id, default_session_id::text
	`, []any{accountID, deviceID, defaultSessionID}, &identity.DeviceID, &identity.DefaultSessionID)
	if err != nil {
		return nil, err
	}
	return identity, nil
}

func randomClaudeCodeDeviceID() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("generate Claude Code device identity: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}
