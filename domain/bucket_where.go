package domain

const KeyPrefixWhere = "where"

type Where struct {
	Bucket

	Value WhereValue `json:"value"`
}

type WhereValue struct {
	Wheres StructureLocations `json:"wheres"`
}

type StructureLocation struct {
	Name    string `json:"name"`
	WhereID string `json:"where_id"`
}

type StructureLocations []StructureLocation
