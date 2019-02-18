package domain

const KeyPrefixStructure = "structure"

type Structure struct {
	Bucket

	Value StructureValue `json:"value"`
}

type StructureValue struct {
	AddressLines                  []string                 `json:"address_lines"`
	AuxPrimaryFabricID            string                   `json:"aux_primary_fabric_id"`
	Away                          bool                     `json:"away"`
	AwaySetter                    int                      `json:"away_setter"`
	City                          string                   `json:"city"`
	Clutches                      []interface{}            `json:"clutches"`
	CountryCode                   string                   `json:"country_code"`
	CreationTime                  int64                    `json:"creation_time"`
	DemandChargeEnabled           bool                     `json:"demand_charge_enabled"`
	Devices                       []string                 `json:"devices"`
	DiamondChangedLocation        bool                     `json:"diamond_changed_location"`
	DrReminderEnabled             bool                     `json:"dr_reminder_enabled"`
	EtaPreconditioningActive      bool                     `json:"eta_preconditioning_active"`
	FabricIDs                     []string                 `json:"fabric_ids"`
	GeofenceEnhancedAutoaway      GeofenceEnhancedAutoaway `json:"geofence_enhanced_autoaway"`
	GooseSensorEventsEnabled      bool                     `json:"goose_sensor_events_enabled"`
	HouseType                     string                   `json:"house_type"`
	HvacSafetyShutoffEnabled      bool                     `json:"hvac_safety_shutoff_enabled"`
	HvacSmokeSafetyShutoffEnabled bool                     `json:"hvac_smoke_safety_shutoff_enabled"`
	IFJPrimaryFabricID            string                   `json:"ifj_primary_fabric_id"`
	Latitude                      float64                  `json:"latitude"`
	Location                      string                   `json:"location"`
	Longitude                     float64                  `json:"longitude"`
	ManualAwayTimestamp           int                      `json:"manual_away_timestamp"`
	ManualEcoAll                  bool                     `json:"manual_eco_all"`
	ManualEcoTimestamp            int                      `json:"manual_eco_timestamp"`
	MeasurementScale              string                   `json:"measurement_scale"`
	Members                       StructureMembers         `json:"members"`
	Name                          string                   `json:"name"`
	NumThermostats                string                   `json:"num_thermostats"`
	PostalCode                    string                   `json:"postal_code"`
	RcsSensorSwarm                []string                 `json:"rcs_sensor_swarm"`
	RenovationDate                string                   `json:"renovation_date"`
	State                         string                   `json:"state"`
	StructureArea                 float64                  `json:"structure_area"`
	Swarm                         []string                 `json:"swarm"`
	TimeZone                      string                   `json:"time_zone"`
	TopazAway                     bool                     `json:"topaz_away"`
	TopazHushKey                  string                   `json:"topaz_hush_key"`
	TOUEnabled                    bool                     `json:"tou_enabled"`
	TouchedBy                     TouchedBy                `json:"touched_by"`
	User                          string                   `json:"user"`
	VacationMode                  bool                     `json:"vacation_mode"`
}

type GeofenceEnhancedAutoaway struct {
	Enabled bool `json:"enabled"`
	IsSet   bool `json:"is_set"`
}

type StructureMember struct {
	Roles []string `json:"roles"`
	User  string   `json:"user"`
}

type StructureMembers []StructureMember
