package response

const (
	QueryEmptyParam    = "Query param cannot be empty."
	QueryEmptyLock     = "Lock query param cannot be empty."
	QueryEmptyName     = "Name query param cannot be empty."
	QueryEmptyTemplate = "Template query param cannot be empty."
	QueryAmountInvalid = "Amount parameter must be positive"

	SeedAddParamRequired = "AddSeed requires add least two parameters ?type={adj, adv, name}&value=v1,v2"
	SeedAddValueRequired = "AddSeed requires at least one inserted value"
	SeedAddTypeRequired  = "AddSeed requires a specified type in {adj, adv, name}"
	SeedRmParamRequired  = "RemoveSeed requires add least two parameters ?type={adj, adv, name}&value=v1,v2"
	SeedRmTypeRequired   = "RemoveSeed requires a specified type in {adj, adv, name}"
	SeedRmValueRequired  = "RemoveSeed requires at least one value"

	APIAddDuplicateError = "Failed insert due to duplicate"
)
