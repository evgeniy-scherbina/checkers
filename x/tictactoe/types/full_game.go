package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func (s *StoredGame) XPlayerAddress() (sdk.Address, error) {
	addr, err := sdk.AccAddressFromBech32(s.XPlayer)
	if err != nil {
		return nil, err
	}

	return addr, nil
}

func (s *StoredGame) YPlayerAddress() (sdk.Address, error) {
	addr, err := sdk.AccAddressFromBech32(s.YPlayer)
	if err != nil {
		return nil, err
	}

	return addr, nil
}

func (s *StoredGame) Validate() error {
	if _, err := s.XPlayerAddress(); err != nil {
		return err
	}
	if _, err := s.YPlayerAddress(); err != nil {
		return err
	}

	return nil
}
