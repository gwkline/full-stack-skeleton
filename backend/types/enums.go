package types

type Operator string

const (
	OperatorUnset              Operator = ""
	OperatorGreaterThan        Operator = ">"
	OperatorLessThan           Operator = "<"
	OperatorGreaterThanOrEqual Operator = ">="
	OperatorLessThanOrEqual    Operator = "<="
	OperatorEqual              Operator = "="
	OperatorIn                 Operator = "IN"
	OperatorIncludes           Operator = "@>"
	OperatorILike              Operator = "ILIKE"
)

type Filter struct {
	Key      string
	Value    any
	Operator Operator
}

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type LoadersString string

const LoadersKey = LoadersString("dataloaders")

type AscOrDesc string

const (
	AscOrDescAsc  AscOrDesc = "ASC"
	AscOrDescDesc AscOrDesc = "DESC"
)

type AppleVariety string

const (
	Fuji            AppleVariety = "FUJI"
	Gala            AppleVariety = "GALA"
	Honeycrisp      AppleVariety = "HONEYCRISP"
	GoldenDelicious AppleVariety = "GOLDEN_DELICIOUS"
	RedDelicious    AppleVariety = "RED_DELICIOUS"
	GrannySmith     AppleVariety = "GRANNY_SMITH"
	Braeburn        AppleVariety = "BRAEBURN"
	Jonagold        AppleVariety = "JONAGOLD"
	CrippsPink      AppleVariety = "CRIPPS_PINK"
	McIntosh        AppleVariety = "MCINTOSH"
	Empire          AppleVariety = "EMPIRE"
	Jonathan        AppleVariety = "JONATHAN"
	Cortland        AppleVariety = "CORTLAND"
	Winesap         AppleVariety = "WINESAP"
	Ambrosia        AppleVariety = "AMBROSIA"
	CosmicCrisp     AppleVariety = "COSMIC_CRISP"
	Envy            AppleVariety = "ENVY"
	Jazz            AppleVariety = "JAZZ"
)
