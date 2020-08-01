package petname

import "errors"

var (
	// QueryEmptyParam is returned when no params are given in
	// a query that requires at least one
	ErrQueryEmptyParam = errors.New("Query param cannot be empty.")

	ErrQueryEmptyLock     = errors.New("Lock query param cannot be empty.")
	ErrQueryEmptyName     = errors.New("Name query param cannot be empty.")
	ErrQueryEmptyTemplate = errors.New("Template query param cannot be empty.")
	ErrQueryAmountInvalid = errors.New("Amount parameter must be positive")

	ErrSeedAddParamRequired = errors.New("AddSeed requires add least two parameters ?type={adj, adv, name}&value=v1,v2")
	ErrSeedAddValueRequired = errors.New("AddSeed requires at least one inserted value")
	ErrSeedAddTypeRequired  = errors.New("AddSeed requires a specified type in {adj, adv, name}")
	ErrSeedRmParamRequired  = errors.New("RemoveSeed requires add least two parameters ?type={adj, adv, name}&value=v1,v2")
	ErrSeedRmTypeRequired   = errors.New("RemoveSeed requires a specified type in {adj, adv, name}")
	ErrSeedRmValueRequired  = errors.New("RemoveSeed requires at least one value")

	ErrAPIAddDuplicate     = errors.New("Failed insert due to duplicate")
	ErrFailedGeneration    = errors.New("Generation failed")
	ErrFailedJSONUnmarshal = errors.New("Failed to unmarshal struct")
)
