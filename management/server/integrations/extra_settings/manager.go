package extra_settings

import (
	"context"

	"github.com/Bee-Bros-Software/r-vpn/management/server/types"
)

type Manager interface {
	GetExtraSettings(ctx context.Context, accountID string) (*types.ExtraSettings, error)
	UpdateExtraSettings(ctx context.Context, accountID, userID string, extraSettings *types.ExtraSettings) (bool, error)
}
