package domain

const KeyPrefixUser = "user"

type User struct {
	Bucket

	Value UserValue `json:"value"`
}

type UserValue struct {
	AcknowledgedOnboardingScreens []string             `json:"acknowledged_onboarding_screens"`
	Email                         string               `json:"email"`
	IsMergedWithGaia              bool                 `json:"is_merged_with_gaia"`
	Name                          string               `json:"name"`
	ShortName                     string               `json:"short_name"`
	StructureMemberships          StructureMemberships `json:"structure_memberships"`
	Structures                    []string             `json:"structures"`
}

type StructureMembership struct {
	Roles     []string `json:"roles"`
	Structure string   `json:"structure"`
}

type StructureMemberships []StructureMembership
