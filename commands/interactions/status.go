package interactions

type Status int

const (
	Accepted        Status = 1
	AlreadyAccepted Status = 2
	BlacklistedUser Status = 3

	ClanNotRegisteredWithGuild Status = 4

	ClanMemberRemoved  Status = 5
	ClanMemberNotFound Status = 6

	Blacklisted        Status = 7
	AlreadyBlacklisted Status = 8

	RoleAdded    Status = 9
	AlreadyAdded Status = 10

	Success               Status = 11
	InvalidID             Status = 12
	UserAlreadyRegistered Status = 13
	UserNotRegistered     Status = 14
	ClanAlreadyRegistered Status = 15
	ClanNotRegistered     Status = 16

	Failure Status = 17

	FailedDBConnection Status = 18
	FailedDBPing       Status = 19
)
