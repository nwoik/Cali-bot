package interactions

import (
	r "calibot/commands/response"
	e "calibot/components/embeds"
	"calibot/globals"
	"fmt"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/model/clan"
	m "github.com/nwoik/calibotapi/model/member"
	"go.mongodb.org/mongo-driver/bson"
)

func AddClan(interaction *discordgo.InteractionCreate) *r.Data {
	client := globals.CLIENT

	clanCollection := client.Database("calibot").Collection("clan")
	clanRepo := c.NewClanRepo(clanCollection)

	memberCollection := client.Database("calibot").Collection("member")
	memberRepo := m.NewMemberRepo(memberCollection)

	userid := interaction.Member.User.ID
	guildid := interaction.GuildID

	member, err := memberRepo.Get(userid)

	if err != nil {
		return r.UserNotRegistered()
	}

	args := interaction.ApplicationCommandData().Options
	name := GetArgument(args, "name").StringValue()
	clanid := GetArgument(args, "clanid").StringValue()

	if len(clanid) < 7 {
		return r.InvalidClanID(name)
	}

	clan, err := clanRepo.Get(guildid)

	if err != nil {
		clan = c.CreateClan(name, clanid, interaction.GuildID).
			SetLeaderID(userid)
		clanRepo.Insert(clan)

		member.ClanID = clan.ClanID
		memberRepo.Update(member)
		return r.RegisteredClan(name)
	}

	return r.ClanAlreadyRegistered()
}

func AddMember(interaction *discordgo.InteractionCreate) *r.Data {
	client := globals.CLIENT

	memberCollection := client.Database("calibot").Collection("member")
	memberRepo := m.NewMemberRepo(memberCollection)

	args := interaction.ApplicationCommandData().Options
	gameid := GetArgument(args, "gameid").StringValue()
	ign := GetArgument(args, "ign").StringValue()

	if len(gameid) < 7 {
		return r.InvalidMemberID(interaction)
	}

	userid := interaction.Member.User.ID
	username := interaction.Member.User.Username

	member, err := memberRepo.Get(userid)

	if err != nil {
		member = m.CreateMember(username, ign, gameid, userid)
		memberRepo.Insert(member)
		return r.RegisteredMember(interaction)
	}

	return r.UserAlreadyRegistered()
}

func AddClanMember(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Data {
	client := globals.CLIENT

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)

	memberCollection := client.Database("calibot").Collection("member")
	memberRepo := m.NewMemberRepo(memberCollection)
	member, err := memberRepo.Get(user.ID)

	if err != nil {
		return r.UserNotRegistered()
	}

	clan, err := GetClan(interaction.GuildID)

	if err != nil {
		return r.ClanNotRegisteredWithGuild()
	}

	if clan.ClanID != member.ClanID {
		if !IsBlacklisted(clan, member.UserID) {
			member.ClanID = clan.ClanID
			memberRepo.Update(member)
			AddRole(session, interaction, member, clan.MemberRole)
			for _, role := range clan.ExtraRoles {
				AddRole(session, interaction, member, role)
			}
			return r.AcceptedMember(user)
		}
		return r.BlacklistedUser()
	}

	return r.AlreadyAccepted()

}

func AddExtraRole(clan *c.Clan, id string) *r.Data {
	client := globals.CLIENT
	clanCollection := client.Database("calibot").Collection("clan")
	clanRepo := c.NewClanRepo(clanCollection)
	clan, err := clanRepo.Get(clan.GuildID)

	if err != nil {
		return r.ClanNotRegistered()
	}

	for _, roleid := range clan.ExtraRoles {
		if roleid == id {
			return r.AlreadyAdded()
		}
	}
	clan.ExtraRoles = append(clan.ExtraRoles, id)
	clanRepo.Update(clan)

	return r.RoleAdded()
}

func AddRole(session *discordgo.Session, interaction *discordgo.InteractionCreate, member *m.Member, role string) {
	err := session.GuildMemberRoleAdd(interaction.GuildID, member.UserID, role)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error changing member role: %s", role), err)
	}
}

func BlacklistUser(clan *c.Clan, session *discordgo.Session, interaction *discordgo.InteractionCreate) (*c.Clan, *r.Data) {
	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)

	if !IsBlacklisted(clan, user.ID) {
		clan.Blacklist = append(clan.Blacklist, user.ID)
		return clan, r.Blacklisted(user)
	}
	return clan, r.AlreadyBlacklisted()
}

func GetMembersWithCond(predicates ...bson.E) ([]*m.Member, error) {
	client := globals.CLIENT

	memberCollection := client.Database("calibot").Collection("member")
	memberRepo := m.NewMemberRepo(memberCollection)

	return memberRepo.Filter(predicates...)
}

func GetArgument(options []*discordgo.ApplicationCommandInteractionDataOption, name string) *discordgo.ApplicationCommandInteractionDataOption {
	for _, option := range options {
		if option.Name == name {
			return option
		}
	}
	return nil
}

func GetClan(id string) (*c.Clan, error) {
	client := globals.CLIENT

	clanCollection := client.Database("calibot").Collection("clan")
	clanRepo := c.NewClanRepo(clanCollection)

	return clanRepo.Get(id)
}

func GetMember(id string) (*m.Member, error) {
	client := globals.CLIENT

	memberCollection := client.Database("calibot").Collection("member")
	memberRepo := m.NewMemberRepo(memberCollection)

	return memberRepo.Get(id)
}

func GetMembers() ([]*m.Member, error) {
	client := globals.CLIENT

	memberCollection := client.Database("calibot").Collection("member")
	memberRepo := m.NewMemberRepo(memberCollection)

	return memberRepo.GetAll()
}

func GetGuild(session *discordgo.Session, guildID string) *discordgo.Guild {
	guild, _ := session.Guild(guildID)
	return guild
}

func GetGuildMember(session *discordgo.Session, guildID string, memberID string) (*discordgo.Member, *r.Data) {
	guildMember, err := session.GuildMember(guildID, memberID)
	if err != nil {
		fmt.Println("Error retrieving member information:", err)
		return nil, r.FailedGetGuildMember()
	}
	return guildMember, nil
}

func MemberEmbed(member *m.Member, guildMember *discordgo.Member, discordUser *discordgo.User) *e.Embed {
	embed := e.NewRichEmbed(member.Nick, "User Info", 0x08d052c)
	embed.SetThumbnail(guildMember.AvatarURL(""))

	embed.AddField("**IGN: **", member.IGN, false)
	embed.AddField("**ID: **", member.IGID, false)

	if member.ClanID != "" {
		clan, err := GetClan(member.ClanID)
		if err == nil {
			embed.AddField("**Clan: **", clan.Name, true)
		}
	}

	embed.SetFooter(fmt.Sprintf("Requested by %s", discordUser.Username), discordUser.AvatarURL(""))

	return embed
}

func PingUser(userid string) string {
	return fmt.Sprintf("<@%s>", userid)
}

func PingRole(id string) string {
	if len(id) < 10 {
		return ""
	}
	return fmt.Sprintf("<@&%s>", id)
}

func PrintRole(id string, roleInClan bool) string {

	if roleInClan {
		return PingRole(id)
	}

	return ""
}

func PrintBlacklist(clan *c.Clan) string {
	var output string

	for _, id := range clan.Blacklist {
		output += PingUser(id) + "\n"
	}

	if output == "" {
		output = "None"
	}

	return output
}

func PrintMembers(members []*m.Member) string {
	var output string

	for _, member := range members {
		output += fmt.Sprintf("%s **IGN: **%s **ID: **%s\n", PingUser(member.UserID), member.IGN, member.IGID)
	}

	if output == "" {
		output = "None"
	}

	return output
}

func PrintExtraRoles(clan *c.Clan, roleInClan bool) string {
	var output string

	if !roleInClan {
		return ""
	}

	for _, id := range clan.ExtraRoles {
		output += PingRole(id) + " "
	}

	if output == "" {
		output = "None"
	}

	return output
}

func RemoveClanMember(clan *c.Clan, member *m.Member, session *discordgo.Session, interaction *discordgo.InteractionCreate) (*m.Member, *r.Data) {
	if clan.ClanID == member.ClanID {
		member.ClanID = ""

		guildMember, _ := GetGuildMember(session, interaction.GuildID, member.UserID)
		RemoveRoles(session, interaction, guildMember)

		return member, r.ClanMemberRemoved()
	}
	return member, r.ClanMemberNotFound()
}

func Remove(slice []string, value string) ([]string, Status) {
	index := -1
	for i, v := range slice {
		if v == value {
			index = i
			break
		}
	}

	if index != -1 {
		return append(slice[:index], slice[index+1:]...), Removed
	}
	return slice, NotFound
}

func RemoveRole(session *discordgo.Session, interaction *discordgo.InteractionCreate, guildMember *discordgo.Member, roleid string) {
	err := session.GuildMemberRoleRemove(interaction.GuildID, guildMember.User.ID, roleid)
	if err != nil {
		fmt.Println("Error removing role from member: ", err)
	}
}

func RemoveRoles(session *discordgo.Session, interaction *discordgo.InteractionCreate, guildMember *discordgo.Member) {
	for _, roleID := range guildMember.Roles {
		RemoveRole(session, interaction, guildMember, roleID)
	}
}

func UpdateMember(interaction *discordgo.InteractionCreate) *r.Data {
	client := globals.CLIENT

	memberCollection := client.Database("calibot").Collection("member")
	memberRepo := m.NewMemberRepo(memberCollection)

	args := interaction.ApplicationCommandData().Options
	gameid := GetArgument(args, "gameid").StringValue()
	ign := GetArgument(args, "ign").StringValue()

	if len(gameid) < 7 {
		return r.InvalidMemberID(interaction)
	}

	userid := interaction.Member.User.ID

	member, err := memberRepo.Get(userid)

	if err == nil {
		member.IGN = ign
		member.IGID = gameid
		memberRepo.Update(member)

		return r.DetailsUpdated()
	}
	return r.UserNotRegistered()
}

func IsRole(session *discordgo.Session, member *m.Member, clan *c.Clan, clanRole string) bool {
	guildMember, _ := GetGuildMember(session, clan.GuildID, member.UserID)

	for _, role := range guildMember.Roles {
		if role == clanRole {
			return true
		}
	}
	return false
}

func IsBlacklisted(clan *c.Clan, userid string) bool {
	for _, id := range clan.Blacklist {
		if id == userid {
			return true
		}
	}

	return false
}
