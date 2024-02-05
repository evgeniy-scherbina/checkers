mock-expected-keepers:
	mockgen -source=x/checkers/types/expected_keepers.go \
		-package testutil \
		-destination=x/checkers/testutil/expected_keepers_mocks.go

oddnum-mock-expected-keepers:
	mockgen -source=x/oddnum/types/expected_keepers.go \
		-package testutil \
		-destination=x/oddnum/testutil/expected_keepers_mocks.go

monobank-mock-expected-keepers:
	mockgen -source=x/monobank/types/expected_keepers.go \
		-package testutil \
		-destination=x/monobank/testutil/expected_keepers_mocks.go
