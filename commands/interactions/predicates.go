package interactions

import (
	"go.mongodb.org/mongo-driver/bson"
)

func Pred(key string, value string) bson.E {
	return bson.E{Key: key, Value: value}
}

func Negate(predicate bson.E) bson.E {
	predicate.Value = bson.D{{Key: "$ne", Value: predicate.Value}}
	return predicate
}
