package service

import (
	"context"

	"github.com/LuckyKuang/sub2api-plus/internal/pkg/logger"
)

// ClaudeCodeAccountIdentity contains the private, account-owned identity used
// only when forwarding Anthropic OAuth or setup-token traffic. Request IDs are
// deliberately excluded because they must remain unique per ingress request.
type ClaudeCodeAccountIdentity struct {
	DeviceID         string
	DefaultSessionID string
}

// ClaudeCodeIdentityRepository is optional so focused GatewayService tests can
// keep using their existing AccountRepository stubs. Production account storage
// implements it with an atomic database upsert.
type ClaudeCodeIdentityRepository interface {
	GetOrCreateClaudeCodeIdentity(ctx context.Context, accountID int64) (*ClaudeCodeAccountIdentity, error)
}

func (s *GatewayService) persistentClaudeCodeIdentity(ctx context.Context, account *Account) *ClaudeCodeAccountIdentity {
	if s == nil || account == nil || !account.IsAnthropicOAuthOrSetupToken() {
		return nil
	}
	repository, ok := s.accountRepo.(ClaudeCodeIdentityRepository)
	if !ok {
		return nil
	}
	identity, err := repository.GetOrCreateClaudeCodeIdentity(ctx, account.ID)
	if err != nil {
		logger.LegacyPrintf("service.gateway", "Failed to resolve persistent Claude Code identity: account_id=%d", account.ID)
		return nil
	}
	return identity
}

func (s *GatewayService) applyPersistentClaudeCodeDevice(ctx context.Context, account *Account, fp *Fingerprint) {
	if fp == nil {
		return
	}
	if identity := s.persistentClaudeCodeIdentity(ctx, account); identity != nil {
		fp.ClientID = identity.DeviceID
	}
}
