mock-expected-keepers:
    mockgen -source=x/checkers/types/expected_keepers.go \
        -package testutil \
        -destination=x/checkers/testutil/expected_keepers_mocks.go

mock-expected-keepers-lessernum:
    mockgen -source=x/lessernum/types/expected_keepers.go \
        -package testutil \
        -destination=x/lessernum/testutil/expected_keepers_mocks.go
