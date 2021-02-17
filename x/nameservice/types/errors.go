package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrNameDoesNotExist is a constant
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 1, "name does not exist")
)
