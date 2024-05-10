package response

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func AcceptedMember(user *discordgo.User) *Data {
	return NewResponseData(fmt.Sprintf("%s has been added to the clan", user.Mention()))
}

func AlreadyAccepted() *Data {
	return NewResponseData("User is already in the clan.")
}

func BlacklistedUser() *Data {
	return NewResponseData("User is blacklisted and cannot be accepted into clan")
}

func UserNotRegistered() *Data {
	return NewResponseData("User is not registered with the bot. User `/register`")
}

func ClanNotRegisteredWithGuild() *Data {
	return NewResponseData("Clan is not registered with the server. User `register-clan`")
}

func ClanMemberRemoved() *Data {
	return NewResponseData("Member has been removed")
}

func ClanMemberNotFound() *Data {
	return NewResponseData("This user isn't in the clan.")
}

func Blacklisted(user *discordgo.User) *Data {
	return NewResponseData(fmt.Sprintf("%s has been blacklisted.", user.Mention()))
}

func AlreadyBlacklisted() *Data {
	return NewResponseData("This user is already blacklisted.")
}

func RemovedFromBlackList() *Data {
	return NewResponseData("User has been removed from clan blacklist")
}

func MemberNotBlacklisted() *Data {
	return NewResponseData("User is not blacklisted")
}

func RoleAdded() *Data {
	return NewResponseData("Role has been added to clan members")
}

func AlreadyAdded() *Data {
	return NewResponseData("Role is already to the clan members")
}

func RoleRemoved() *Data {
	return NewResponseData("Role has been removed from clan members")
}

func RoleNotFound() *Data {
	return NewResponseData("Role was never assigned to clan members")
}

func RegisteredMember(interaction *discordgo.InteractionCreate) *Data {
	return NewResponseData("Registered " + interaction.Member.User.Mention())
}

func InvalidMemberID(interaction *discordgo.InteractionCreate) *Data {
	return NewResponseData("Invalid Game-ID. Failed to register" + interaction.Member.User.Mention())
}

func UserAlreadyRegistered() *Data {
	return NewResponseData("User is already registered. Details were updated")
}

func RegisteredClan(name string) *Data {
	return NewResponseData("Registered Clan: " + name +
		"\nUse `/viewclan` to see details" +
		"\nMake sure to use `/leaderrole`, `/officerrole` and `/memberrole` for the roles you want members to have.")
}

func InvalidClanID(name string) *Data {
	return NewResponseData("Invalid Game-ID. Failed to register " + name)
}

func ClanAlreadyRegistered() *Data {
	return NewResponseData("Clan is already registered. Details were updated")
}

func ClanNotRegistered() *Data {
	return NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`")
}

func FailedDBConnection() *Data {
	return NewResponseData("Failed to connect to database. Please ping admins")
}

func FailedDBPing() *Data {
	return NewResponseData("Failed to ping database")
}

func FailedGetGuildMember() *Data {
	return NewResponseData("Error retrieving member information.")
}

func OfficerDemoted(user *discordgo.User) *Data {
	return NewResponseData(fmt.Sprintf("%s has been demoted :cry:", user.Mention()))
}

func MemberRoleResponse() *Data {
	return NewResponseData("Member role registered")
}

func OfficerRoleResponse() *Data {
	return NewResponseData("Officer role registered")
}

func LeaderRoleResponse() *Data {
	return NewResponseData("Leader role registered")
}

func Promote(user *discordgo.User) *Data {
	return NewResponseData(fmt.Sprintf("%s has been appointed as officer :man_police_officer:", user.Mention()))
}

func CantPromoteNonMember() *Data {
	return NewResponseData("This user is not registered with the bot." +
		"\nThey must register with the bot and clan to be an officer")
}
