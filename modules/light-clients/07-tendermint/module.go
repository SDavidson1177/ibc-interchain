package tendermint

import (
	"github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
