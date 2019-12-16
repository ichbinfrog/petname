package response

import (
	"errors"
)

var (
	QueryEmptyParam    = errors.New("Query param cannot be empty.")
	QueryEmptyLock     = errors.New("Lock query param cannot be empty.")
	QueryEmptyName     = errors.New("Name query param cannot be empty.")
	QueryEmptyTemplate = errors.New("Template query param cannot be empty.")
	QueryAmountInvalid = errors.New("Amount parameter must be positive")
)

var (
	SeedAddParamRequired = errors.New("AddSeed requires add least two parameters ?type={adj, adv, name}&value=v1,v2")
	SeedAddValueRequired = errors.New("AddSeed requires at least one inserted value")
	SeedAddTypeRequired  = errors.New("AddSeed requires a specified type in {adj, adv, name}")
	SeedRmParamRequired  = errors.New("RemoveSeed requires add least two parameters ?type={adj, adv, name}&value=v1,v2")
	SeedRmTypeRequired   = errors.New("RemoveSeed requires a specified type in {adj, adv, name}")
	SeedRmValueRequired  = errors.New("RemoveSeed requires at least one value")
)

var (
	APIAddDuplicateError = errors.New("Failed insert due to duplicate")
)
