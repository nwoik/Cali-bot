package util

import (
	c "github.com/nwoik/calibotapi/clan"
	m "github.com/nwoik/calibotapi/member"
)

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

func Filter(array []interface{}, Predicate func(interface{}) bool) []interface{} {
	filtered := make([]interface{}, 0)

	for _, item := range array {
		if Predicate(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func FilterByClan(clan *c.Clan) Predicate {
	return func(member interface{}) bool {
		return member.(*m.Member).ClanID == clan.ClanID
	}
}
