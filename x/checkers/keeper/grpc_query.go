package keeper

import (
	"github.com/nikolay-stoev/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
