package pas

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

const (
	TypeString     = PropertyType("string")
	TypeInteger    = PropertyType("integer")
	TypeBigInteger = PropertyType("big_integer")
	TypeFloat      = PropertyType("float")
	TypeDouble     = PropertyType("double")
	TypeDecimal    = PropertyType("decimal")
	TypeBoolean    = PropertyType("boolean")
	TypeDate       = PropertyType("date")
	TypeDateTime   = PropertyType("datetime")
)

var propertyTypes = map[PropertyType]string{
	TypeString:     "varchar(2000)",
	TypeInteger:    "int",
	TypeBigInteger: "bigint",
	TypeFloat:      "float(23, 2)",
	TypeDouble:     "double(53, 2)",
	TypeDecimal:    "decimal",
	TypeBoolean:    "tinyint(1)",
	TypeDate:       "date",
	TypeDateTime:   "datetime",
}

type Property struct {
	Type      PropertyType `json:"type"`
	Name      PropertyName `json:"name"`
	Value     interface{}  `json:"value"`
	Precision int          `json:"precision"`
	Scale     int          `json:"scale"`
}

func (p Property) dbType() string {
	t := propertyTypes[p.Type]
	if p.Type == TypeDecimal {
		return t + "(" + strconv.Itoa(p.Precision) + "," + strconv.Itoa(p.Scale) + ")"
	}
	return t
}

type PropertyName string

var propertyNameRegex = regexp.MustCompile(`[a-z]+[a-z_0-9 \-]*`)

func (n *PropertyName) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	if len(*n) > 255 {
		return errors.New("property name too big")
	}
	if !propertyNameRegex.MatchString(s) {
		return errors.New("invalid property name")
	}
	*n = PropertyName(s)
	return nil
}

type PropertyType string

func (t *PropertyType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = PropertyType(s)
	_, ok := propertyTypes[*t]
	if !ok {
		return errors.New("unknown property type")
	}
	return nil
}
