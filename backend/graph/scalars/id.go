package scalars

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// Lets redefine the base ID type to use an id from gorm's default uint type
func MarshalID(id uint) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
	})
}

// And the same for the unmarshaler
func UnmarshalID(v interface{}) (uint, error) {
	str, okString := v.(string)
	integ, okInt := v.(int)
	integ64, okInt64 := v.(int64)
	if !okString && !okInt && !okInt64 {
		return 0, fmt.Errorf("ids must be strings or ints")
	}

	if okInt64 {
		if integ64 < 0 {
			return 0, fmt.Errorf("ids must be positive numbers")
		}
		return uint(integ64), nil
	}

	if okInt {
		if integ < 0 {
			return 0, fmt.Errorf("ids must be positive numbers")
		}
		return uint(integ), nil
	}

	i, err := strconv.Atoi(str)
	return uint(i), err
}
