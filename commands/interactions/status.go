package interactions

type Status int

const (
	Accepted        Status = 1
	AlreadyAccepted Status = 2
	BlacklistedUser Status = 3
	NotRegistered   Status = 4

	Removed  Status = 5
	NotFound Status = 6

	Blacklisted        Status = 7
	AlreadyBlacklisted Status = 8

	RoleAdded    Status = 9
	AlreadyAdded Status = 10

	Success           Status = 11
	InvalidID         Status = 12
	AlreadyRegistered Status = 13
	Failure           Status = 14
	UserNotRegistered Status = 15
)
