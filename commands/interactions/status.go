package interactions

type Status int

const (
	Accepted        Status = 1
	AlreadyAccepted Status = 2
	BlacklistedUser Status = 3
	NotRegistered   Status = 4

	Removed          Status = 5
	MemberNotPresent Status = 6

	Blacklisted        Status = 7
	AlreadyBlacklisted Status = 8

	RoleAdded    Status = 9
	AlreadyAdded Status = 10

	RoleRemoved  Status = 11
	RoleNotFound Status = 12

	Success           Status = 13
	InvalidID         Status = 14
	AlreadyRegistered Status = 15
	Failure           Status = 16
	UserNotRegistered Status = 17
)
