package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyEnableFeedist = []byte("EnableFeedist")
	// TODO: Determine the default value
	DefaultEnableFeedist bool = true
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	enableFeedist bool,
) Params {
	return Params{
		EnableFeedist: enableFeedist,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultEnableFeedist,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyEnableFeedist, &p.EnableFeedist, validateEnableFeedist),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateEnableFeedist(p.EnableFeedist); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateEnableFeedist validates the EnableFeedist param
func validateEnableFeedist(v interface{}) error {
	enableFeedist, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = enableFeedist

	return nil
}
