package interactions

import (
	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/model/clan"
	m "github.com/nwoik/calibotapi/model/member"
)

// func InClan(clan *c.Clan) bson.E {
// 	return bson.E{Key: "clanid", Value: clan.ClanID}
// }

type Predicate func(interface{}) bool

func And(predicates ...Predicate) Predicate {
	return func(item interface{}) bool {
		for _, pred := range predicates {
			if !pred(item) {
				return false
			}
		}
		return true
	}
}

func Or(predicates ...Predicate) Predicate {
	return func(item interface{}) bool {
		for _, pred := range predicates {
			if pred(item) {
				return true
			}
		}
		return false
	}
}

func Negate(pred Predicate) Predicate {
	return func(item interface{}) bool {
		return !pred(item)
	}
}

func FilterMembers(members []*m.Member, Predicate func(interface{}) bool) []*m.Member {
	filtered := make([]*m.Member, 0)

	for _, item := range members {
		if Predicate(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func InClan(clan *c.Clan) Predicate {
	return func(member interface{}) bool {
		return member.(*m.Member).ClanID == clan.ClanID
	}
}

func IsLeader(clan *c.Clan) Predicate {
	return func(member interface{}) bool {
		return member.(*m.Member).UserID == clan.LeaderID
	}
}

func IsOfficer(session *discordgo.Session, clan *c.Clan) Predicate {
	return func(member interface{}) bool {
		return IsRole(session, member.(*m.Member), clan, clan.OfficerRole)
	}
}

func IsMember(session *discordgo.Session, clan *c.Clan) Predicate {
	return func(member interface{}) bool {
		return IsRole(session, member.(*m.Member), clan, clan.MemberRole)
	}
}
