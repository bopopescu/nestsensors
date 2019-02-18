package domain

import (
	"encoding/json"
	"errors"
)

var (
	ErrUnknownID = errors.New("unknown id")
)

type NestState struct {
	TwoFaEnabled         bool                 `json:"2fa_enabled"`
	ServiceURLs          ServiceURLs          `json:"service_urls"`
	UpdatedBuckets       []json.RawMessage    `json:"updated_buckets"`
	WeatherForStructures WeatherForStructures `json:"weather_for_structures"`
}

type ServiceURLs struct {
	Limits Limits `json:"limits"`
	URLs   URLs   `json:"urls"`
	Weave  Weave  `json:"weave"`
}

type Limits struct {
	SmokeDetectors             int `json:"smoke_detectors"`
	SmokeDetectorsPerStructure int `json:"smoke_detectors_per_structure"`
	Structures                 int `json:"structures"`
	Thermostats                int `json:"thermostats"`
	ThermostatsPerStructure    int `json:"thermostats_per_structure"`
}

type URLs struct {
	CzfeURL            string `json:"czfe_url"`
	DirectTransportURL string `json:"direct_transport_url"`
	LogUploadURL       string `json:"log_upload_url"`
	RubyapiURL         string `json:"rubyapi_url"`
	SupportURL         string `json:"support_url"`
	TransportURL       string `json:"transport_url"`
	WeatherURL         string `json:"weather_url"`
}

type Weave struct {
	ServiceConfig string `json:"service_config"`
	PairingToken  string `json:"pairing_token"`
	AccessToken   string `json:"access_token"`
}

type UpdatedBucket struct {
	ObjectKey       string                 `json:"object_key"`
	ObjectRevision  int                    `json:"object_revision"`
	ObjectTimestamp int64                  `json:"object_timestamp"`
	Value           map[string]interface{} `json:"value"` // Depends entirely on ObjectRevision.
}

type UpdatedBuckets []UpdatedBucket

type WeatherForStructure struct {
	Current  WeatherCurrent  `json:"current"`
	Location WeatherLocation `json:"location"`
}

type WeatherForStructures map[string]WeatherForStructure

type WeatherCurrent struct {
	Icon    string `json:"icon"`
	Sunrise string `json:"sunrise"`
	Sunset  string `json:"sunset"`
	TempC   string `json:"temp_c"`
}

type WeatherLocation struct {
	City    string `json:"city"`
	Country string `json:"country"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

// WhereToName looks up the human friendly structure name for a WhereID UUID.
func (ns *NestState) WhereToName(id string) (string, error) {
	wheres, err := ns.Wheres()
	if err != nil {
		return "", err
	}
	for _, w := range wheres {
		for _, strucLoc := range w.Value.Wheres {
			if strucLoc.WhereID == id {
				return strucLoc.Name, nil
			}
		}
	}
	return "", ErrUnknownID
}

/*func (ns *NestState) Temps() (map[string]float64, error) {
	wheresIface, err := ns.Lookup(KeyPrefixWhere)
	if err != nil {
		return nil, err
	}
	wheres := wheresIface.([]*Where)

	fmt.Printf("%+v\n", wheres)

	// sensorIface, err := ns.Lookup(KeyPrefixTemperatureSensor)
	// if err != nil {
	// 	return err
	// }
	// sensor := sensorIface.(*TemperatureSensor)

	// sensor.Value.

	m := map[string]float64{}
	return m, nil
}

func (ns *NestState) TemperatureSensors() ([]*TemperatureSensor, error) {
	temperatureSensors := []*TemperatureSensor{}
	ifaces, err := ns.Lookup(KeyPrefixTemperatureSensor)
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		temperatureSensor, ok := iface.(*TemperatureSensor)
		if !ok {
			return nil, fmt.Errorf("failed to cast %[1]T/%[1]# v to type 'TemperatureSensor'", iface)
		}
	}
	return temperatureSensors, nil
}*/

// // Lookup finds and returns the a pointer of type <bucket item> corresponding
// // with the named key.
// func (ns *NestState) Lookup(key string) ([]interface{}, error) {
// 	out := []interface{}{}
// 	for _, v := range ns.UpdatedBuckets {
// 		b := &Bucket{}
// 		if err := json.Unmarshal(v, b); err != nil {
// 			return nil, err
// 		}
// 		if strings.HasPrefix(b.ObjectKey, key+".") {
// 			obj := keys[key]
// 			if err := json.Unmarshal(v, &obj); err != nil {
// 				return nil, err
// 			}
// 			fmt.Printf("Found obj type=%v %T\n", key, obj)
// 			out = append(out, obj)
// 		}
// 	}
// 	return out, nil
// }

/*
{
    "2fa_enabled": true,
    "service_urls": {
        "limits": {
            "smoke_detectors": 54,
            "smoke_detectors_per_structure": 18,
            "structures": 3,
            "thermostats": 60,
            "thermostats_per_structure": 20
        },
        "urls": {
            "czfe_url": "https://czfe25-front01-iad01.transport.home.nest.com",
            "direct_transport_url": "https://transport04-rts02-iad01.transport.home.nest.com:443",
            "log_upload_url": "https://logsink.home.nest.com/upload/user",
            "rubyapi_url": "https://home.nest.com/",
            "support_url": "https://nest.secure.force.com/support/webapp?",
            "transport_url": "https://czfe25-front01-iad01.transport.home.nest.com",
            "weather_url": "https://apps-weather.nest.com/weather/v1?query="
        },
        "weave": {
            "access_token": "lQkANQEwAQgKkeRnuqLsfyQCBDcDLAEIMTMyNjg3MjMYJgQHPnskJgUH56A3NwYsAQgxMzI2ODcyMxgkBwImCCUAWiMwCjkEm+UGG09SH5pHGY3mEEkvMvQ574bVZ+kY19/mk0fc4kWLU9+I1XJHiJSxcU0xYs9et24ebNEuWUA1gykBGDWCKQEkAgUYNYQpATYCBAIEARgYNYEwAghMginnfigomhg1gDACCEyCKed+KCiaGDUMMAEcNl3fIFGF4dzgzjlbWjt0AACG/kQdbFGBPPY88DACHEnlpe1jAkRxcJkWlnmtn+3W5olu/GaIM9fm4roYGDUCJgElAFojMAIcMiyRPoOuE3RPhBt2LYld+Y+D/Zgvyab4Jp2NDRgY",
            "pairing_token": "wu.gbgaNqMLiM0i3xfdbOSWm/65HNb6cgNHKTx2sEI/7osI2fAQLSxz6LtWAv+Oe1rDwaj/6U04C6NpbSHo1FvhWx2K+0k=",
            "service_config": "1QAADwABADYBFTABCEdYfNbKQ+Z8JAIENwMnEwEAAADuMLQYGCYERUgWGiYFRQb7STcGJxMBAAAA7jC0GBgkBwImCCUAWiMwCjkEFJtArVoBj0T+fpKZeV1TBT0osdlSUx1FRK1ZTcslNfDQSMUwiG3iH5tQjeGhrxbfr0kphGthRBA1gykBKQIYNYIpASQCYBg1gTACCEM09xLfX5HPGDWAMAIIQzT3Et9fkc8YNQwwAR0AlvB51cbUag/T//0P7EqwJNQATpT0hWNT7odJgzACHQCPTm4jL9hwiKnIKvpqXtz6LnCd0ifQrU9apo0jGBgVMAEICpHkZ7qi7H8kAgQ3AywBCDEzMjY4NzIzGCYEBz57JCYFB+egNzcGLAEIMTMyNjg3MjMYJAcCJgglAFojMAo5BJvlBhtPUh+aRxmN5hBJLzL0Oe+G1WfpGNff5pNH3OJFi1PfiNVyR4iUsXFNMWLPXrduHmzRLllANYMpARg1gikBJAIFGDWEKQE2AgQCBAEYGDWBMAIITIIp534oKJoYNYAwAghMginnfigomhg1DDABHDZd3yBRheHc4M45W1o7dAAAhv5EHWxRgTz2PPAwAhxJ5aXtYwJEcXCZFpZ5rZ/t1uaJbvxmiDPX5uK6GBgYNQInAQEAAAACMLQYNgIVLAESZnJvbnRkb29yLm5lc3QuY29tJQJXKxgYGBg="
        }
    },
    "updated_buckets": [
        {
            "object_key": "widget_track.6416660000C66097",
            "object_revision": -12090,
            "object_timestamp": 1550395282030,
            "value": {
                "last_connection": 1550395282030,
                "last_ip": "/172.16.47.159",
                "online": true
            }
        },
        {
            "object_key": "safety.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
            "object_revision": -25710,
            "object_timestamp": 1549867870993,
            "value": {
                "audio_self_test_enabled": true,
                "audio_self_test_end_utc_offset_secs": 0,
                "audio_self_test_participants": [
                    "6416660000C66097"
                ],
                "audio_self_test_start_utc_offset_secs": 68400,
                "last_manual_self_test_cancelled": true,
                "last_manual_self_test_end_utc_secs": 1549867865,
                "last_manual_self_test_start_utc_secs": 1549867865,
                "manual_self_test_in_progress": false
            }
        },
        {
            "object_key": "shared.09AA01AC421818BS",
            "object_revision": 4206,
            "object_timestamp": 1550451921590,
            "value": {
                "auto_away": 1,
                "auto_away_learning": "ready",
                "can_cool": false,
                "can_heat": true,
                "compressor_lockout_enabled": false,
                "compressor_lockout_timeout": 0,
                "current_temperature": 20.18999,
                "hvac_ac_state": false,
                "hvac_alt_heat_state": false,
                "hvac_alt_heat_x2_state": false,
                "hvac_aux_heater_state": false,
                "hvac_cool_x2_state": false,
                "hvac_cool_x3_state": false,
                "hvac_emer_heat_state": false,
                "hvac_fan_state": false,
                "hvac_heat_x2_state": false,
                "hvac_heat_x3_state": false,
                "hvac_heater_state": true,
                "name": "",
                "target_change_pending": false,
                "target_temperature": 20.16695,
                "target_temperature_high": 24,
                "target_temperature_low": 20,
                "target_temperature_type": "heat",
                "touched_by": {
                    "touched_by": 3,
                    "touched_id": "Thermostat on IFTTT",
                    "touched_user_id": "client.10f540d7-88a2-45b3-86d8-f7e4e4f428d2"
                }
            }
        },
        {
            "object_key": "metadata.09AA01AC421818BS",
            "object_revision": -1,
            "object_timestamp": 1356998400000,
            "value": {
                "last_connection": 1356998400000,
                "last_ip": "127.0.0.1"
            }
        },
        {
            "object_key": "kryptonite.18B430CEBD24D3CF",
            "object_revision": 20867,
            "object_timestamp": 1550452819734,
            "value": {
                "battery_level": 99,
                "current_temperature": 20.29999,
                "last_updated_at": 1550452818,
                "model": "KR1",
                "serial_number": "22AA01AC291806C3",
                "structure_id": "a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                "where_id": "00000000-0000-0000-0000-000100000008"
            }
        },
        {
            "object_key": "kryptonite.18B430CFD99E3ABC",
            "object_revision": -21881,
            "object_timestamp": 1550452520332,
            "value": {
                "battery_level": 100,
                "current_temperature": 19.2,
                "last_updated_at": 1550452518,
                "model": "KR1",
                "serial_number": "22AA01AC291806CT",
                "structure_id": "a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                "where_id": "00000000-0000-0000-0000-00010000000d"
            }
        },
        {
            "object_key": "occupancy.6416660000C66097",
            "object_revision": 11908,
            "object_timestamp": 1550395281740,
            "value": {
                "element_duration_seconds": 300,
                "occupancy": [
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    1,
                    1,
                    1,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    1,
                    0,
                    1,
                    1,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    1,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    1,
                    1,
                    0,
                    1,
                    1,
                    0,
                    1,
                    1,
                    1,
                    1,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    1,
                    0,
                    0,
                    1,
                    1,
                    1,
                    0,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0,
                    1,
                    1,
                    0,
                    0,
                    0,
                    0,
                    0,
                    0
                ]
            }
        },
        {
            "object_key": "user_settings.13268723",
            "object_revision": -11097,
            "object_timestamp": 1547768789262,
            "value": {
                "app_swu_minimum_version": {
                    "pairing": 2
                },
                "email_verified": true,
                "has_location_history": true,
                "lang": "en_US",
                "location_primary_device": "d28e6d35-1e3b-4f90-9e0a-db90b887bb8c",
                "max_kryptonites": 54,
                "max_kryptonites_per_structure": 18,
                "max_kryptonites_per_thermostat": 6,
                "max_smoke_detectors": 54,
                "max_smoke_detectors_per_structure": 18,
                "max_structures": 3,
                "max_thermostats": 60,
                "max_thermostats_per_structure": 20,
                "max_wwn_devices": 20,
                "max_wwn_devices_per_structure": 10,
                "receive_marketing_emails": false,
                "receive_nest_emails": false,
                "receive_support_emails": true,
                "state_2fa": "none",
                "tos_accepted_version": 1527058800000,
                "tos_current_version": 1527058800000,
                "tos_minimum_version": 1527058800000
            }
        },
        {
            "object_key": "rcs_settings.09AA01AC421818BS",
            "object_revision": -22419,
            "object_timestamp": 1550417959998,
            "value": {
                "active_rcs_sensors": [
                    "kryptonite.18B430C79EF7406D"
                ],
                "associated_rcs_sensors": [
                    "kryptonite.18B430C49E3DA2A2",
                    "kryptonite.18B430C79EF7406D",
                    "kryptonite.18B430CEBD24D3CF",
                    "kryptonite.18B430CFD99E3ABC"
                ],
                "multiroom_active": false,
                "multiroom_temperature": 17.67999,
                "rcs_control_setting": "SCHEDULE",
                "sensor_insights": {
                    "insights": [
                        {
                            "created_at": 1550390642,
                            "range_test_result": "STRONG",
                            "sensor_id": "kryptonite.18B430CFD99E3ABC",
                            "value": 18
                        },
                        {
                            "created_at": 1550390642,
                            "range_test_result": "STRONG",
                            "sensor_id": "kryptonite.18B430C79EF7406D",
                            "value": 21
                        },
                        {
                            "created_at": 1550390642,
                            "range_test_result": "STRONG",
                            "sensor_id": "kryptonite.18B430C49E3DA2A2",
                            "value": 21
                        },
                        {
                            "created_at": 1550390642,
                            "range_test_result": "STRONG",
                            "sensor_id": "kryptonite.18B430CEBD24D3CF",
                            "value": 20
                        }
                    ]
                },
                "sensor_schedule": {
                    "intervals": [
                        {
                            "end_day": 6,
                            "end_time": 39600,
                            "multiroom_active": false,
                            "sensors": [
                                "kryptonite.18B430C79EF7406D"
                            ],
                            "start_day": 0,
                            "start_time": 25201
                        },
                        {
                            "end_day": 6,
                            "end_time": 57600,
                            "multiroom_active": false,
                            "sensors": [
                                "kryptonite.18B430C79EF7406D"
                            ],
                            "start_day": 0,
                            "start_time": 39601
                        },
                        {
                            "end_day": 6,
                            "end_time": 75600,
                            "multiroom_active": false,
                            "sensors": [
                                "kryptonite.18B430C79EF7406D"
                            ],
                            "start_day": 0,
                            "start_time": 57601
                        },
                        {
                            "end_day": 6,
                            "end_time": 25200,
                            "multiroom_active": false,
                            "sensors": [
                                "kryptonite.18B430CFD99E3ABC"
                            ],
                            "start_day": 0,
                            "start_time": 75601
                        }
                    ]
                },
                "thermostat_alert": "OFF"
            }
        },
        {
            "object_key": "buckets.13268723",
            "object_revision": -23656,
            "object_timestamp": 1550163787107,
            "value": {
                "buckets": [
                    "quartz.f67bb3d58fbf4602b0ace3a27382aae6",
                    "kryptonite.18B430C79EF7406D",
                    "topaz_resource.5",
                    "user.13268723",
                    "user_alert_dialog.13268723",
                    "message_center.13268723",
                    "user_settings.13268723",
                    "buckets.13268723",
                    "kryptonite.18B430CFD99E3ABC",
                    "kryptonite.18B430CEBD24D3CF",
                    "device.09AA01AC421818BS",
                    "demand_response.09AA01AC421818BS",
                    "device_alert_dialog.09AA01AC421818BS",
                    "link.09AA01AC421818BS",
                    "where.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "utility.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "structure_history.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "trip.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "wwn_security.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "structure_metadata.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "safety.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "safety_summary.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "structure.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "geofence_info.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "partner_programs.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "kryptonite.18B430C49E3DA2A2",
                    "message.09AA01AC421818BS",
                    "metadata.09AA01AC421818BS",
                    "schedule.09AA01AC421818BS",
                    "shared.09AA01AC421818BS",
                    "hvac_partner.09AA01AC421818BS",
                    "track.09AA01AC421818BS",
                    "tuneups.09AA01AC421818BS",
                    "energy_latest.09AA01AC421818BS",
                    "energy_weekly.09AA01AC421818BS",
                    "tou.09AA01AC421818BS",
                    "demand_charge.09AA01AC421818BS",
                    "demand_response_fleet.09AA01AC421818BS",
                    "rcs_settings.09AA01AC421818BS",
                    "cloud_algo.09AA01AC421818BS",
                    "diagnostics.09AA01AC421818BS",
                    "hvac_issues.09AA01AC421818BS",
                    "topaz.6416660000C66097",
                    "occupancy.6416660000C66097",
                    "widget_track.6416660000C66097",
                    "topaz_history.6416660000C66097",
                    "delayed_topaz.6416660000C66097"
                ]
            }
        },
        {
            "object_key": "tuneups.09AA01AC421818BS",
            "object_revision": 28298,
            "object_timestamp": 1547786576459,
            "value": {}
        },
        {
            "object_key": "device.09AA01AC421818BS",
            "object_revision": -27612,
            "object_timestamp": 1550452521566,
            "value": {
                "alt_heat_delivery": "forced-air",
                "alt_heat_source": "gas",
                "alt_heat_x2_delivery": "forced-air",
                "alt_heat_x2_source": "gas",
                "auto_away_enable": true,
                "auto_away_reset": false,
                "auto_dehum_enabled": false,
                "auto_dehum_state": false,
                "aux_heat_delivery": "forced-air",
                "aux_heat_source": "electric",
                "aux_lockout_leaf": 10,
                "available_locales": "en_US,fr_CA,es_US,en_GB,fr_FR,nl_NL,es_ES,it_IT,de_DE",
                "away_temperature_high": 24.44444,
                "away_temperature_high_adjusted": 24.44444,
                "away_temperature_high_enabled": false,
                "away_temperature_low": 16.409557,
                "away_temperature_low_adjusted": 16.40955,
                "away_temperature_low_enabled": true,
                "backplate_bsl_info": "BSL",
                "backplate_bsl_version": "3.1",
                "backplate_model": "Backplate-5.4",
                "backplate_mono_info": "TFE (BP_D3) 1.5.9 (jenkins-slave@jenkins-agent-017-v17-emb-prod) 2018-08-27 18:18:13",
                "backplate_mono_version": "1.5.9",
                "backplate_serial_number": "09DA02AC42180D74",
                "backplate_temperature": 21.42999,
                "battery_level": 3.93,
                "capability_level": 5.93,
                "click_sound": "on",
                "compressor_lockout_leaf": -17.79999,
                "cooling_delivery": "unknown",
                "cooling_source": "electric",
                "cooling_x2_delivery": "unknown",
                "cooling_x2_source": "electric",
                "cooling_x3_delivery": "unknown",
                "cooling_x3_source": "electric",
                "country_code": "US",
                "creation_time": 1547786576087,
                "current_humidity": 47,
                "current_schedule_mode": "HEAT",
                "current_version": "5.9.3-6",
                "dehumidifier_orientation_selected": "unknown",
                "dehumidifier_state": false,
                "dehumidifier_type": "unknown",
                "demand_charge_icon": false,
                "device_locale": "en_US",
                "dual_fuel_breakpoint": -1,
                "dual_fuel_breakpoint_override": "none",
                "eco": {
                    "mode": "schedule",
                    "mode_update_timestamp": 1550428973,
                    "touched_by": 1
                },
                "eco_onboarding_needed": false,
                "emer_heat_delivery": "forced-air",
                "emer_heat_enable": false,
                "emer_heat_source": "electric",
                "error_code": "",
                "fan_capabilities": "none",
                "fan_control_state": false,
                "fan_cooling_enabled": true,
                "fan_cooling_readiness": "not ready",
                "fan_cooling_state": false,
                "fan_current_speed": "off",
                "fan_duty_cycle": 3600,
                "fan_duty_end_time": 0,
                "fan_duty_start_time": 0,
                "fan_heat_cool_speed": "auto",
                "fan_mode": "auto",
                "fan_schedule_speed": "stage1",
                "fan_timer_duration": 900,
                "fan_timer_speed": "stage1",
                "fan_timer_timeout": 0,
                "farsight_screen": "target_temp",
                "filter_changed_date": 0,
                "filter_changed_set_date": 0,
                "filter_reminder_enabled": false,
                "filter_reminder_level": 0,
                "filter_replacement_needed": false,
                "filter_replacement_threshold_sec": 1800000,
                "filter_runtime_sec": 671820,
                "forced_air": true,
                "gear_threshold_high": 0,
                "gear_threshold_low": 0,
                "has_air_filter": true,
                "has_alt_heat": false,
                "has_aux_heat": false,
                "has_dehumidifier": false,
                "has_dual_fuel": false,
                "has_emer_heat": false,
                "has_fan": false,
                "has_fossil_fuel": true,
                "has_heat_pump": false,
                "has_hot_water_control": false,
                "has_hot_water_temperature": false,
                "has_humidifier": false,
                "has_x2_alt_heat": false,
                "has_x2_cool": false,
                "has_x2_heat": false,
                "has_x3_cool": false,
                "has_x3_heat": false,
                "heat_link_connection": 0,
                "heat_pump_aux_threshold": 10,
                "heat_pump_aux_threshold_enabled": true,
                "heat_pump_comp_threshold": -31.5,
                "heat_pump_comp_threshold_enabled": false,
                "heat_x2_delivery": "forced-air",
                "heat_x2_source": "gas",
                "heat_x3_delivery": "forced-air",
                "heat_x3_source": "gas",
                "heater_delivery": "forced-air",
                "heater_source": "gas",
                "heatpump_ready": false,
                "heatpump_savings": "off",
                "heatpump_setback_active": false,
                "home_away_input": true,
                "hot_water_active": true,
                "hot_water_away_active": false,
                "hot_water_away_enabled": true,
                "hot_water_boiling_state": true,
                "humidifier_state": false,
                "humidifier_type": "unknown",
                "humidity_control_lockout_enabled": false,
                "humidity_control_lockout_end_time": 0,
                "humidity_control_lockout_start_time": 0,
                "hvac_pins": "W1,C,Rc",
                "hvac_safety_shutoff_active": false,
                "hvac_smoke_safety_shutoff_active": false,
                "hvac_staging_ignore": false,
                "hvac_wires": "Heat,Common Wire,Rc",
                "is_furnace_shutdown": false,
                "kryptonite_range_test_timestamp": 1548131675,
                "last_software_update_utc_secs": 1548131457,
                "leaf": false,
                "leaf_away_high": 28.87999,
                "leaf_away_low": 16.7903,
                "leaf_schedule_delta": 1.10999,
                "leaf_threshold_cool": 0,
                "leaf_threshold_heat": 19.05696,
                "learning_mode": true,
                "local_ip": "192.168.225.57",
                "logging_priority": "informational",
                "lower_safety_temp": 4.44443,
                "lower_safety_temp_enabled": true,
                "mac_address": "6416668c51ec",
                "maint_band_lower": 0.39,
                "maint_band_upper": 0.39,
                "max_nighttime_preconditioning_seconds": 3600,
                "model_version": "Display-3.4",
                "nlclient_state": "",
                "ob_orientation": "O",
                "ob_persistence": true,
                "oob_interview_completed": true,
                "oob_startup_completed": true,
                "oob_summary_completed": true,
                "oob_temp_completed": true,
                "oob_test_completed": true,
                "oob_where_completed": true,
                "oob_wifi_completed": true,
                "oob_wires_completed": true,
                "pin_c_description": "power",
                "pin_g_description": "none",
                "pin_ob_description": "none",
                "pin_rc_description": "power",
                "pin_rh_description": "none",
                "pin_star_description": "none",
                "pin_w1_description": "heat",
                "pin_w2aux_description": "none",
                "pin_y1_description": "none",
                "pin_y2_description": "none",
                "postal_code": "94025",
                "preconditioning_active": false,
                "preconditioning_enabled": true,
                "preconditioning_ready": true,
                "pro_id": "",
                "radiant_control_enabled": false,
                "rcs_capable": true,
                "rssi": 73,
                "safety_state": "none",
                "safety_state_time": 0,
                "safety_temp_activating_hvac": false,
                "schedule_learning_reset": false,
                "schedules": [],
                "serial_number": "09AA01AC421818BS",
                "should_wake_on_approach": true,
                "smoke_shutoff_supported": true,
                "star_type": "unknown",
                "sunlight_correction_active": false,
                "sunlight_correction_enabled": true,
                "sunlight_correction_ready": true,
                "target_humidity": 35,
                "target_humidity_enabled": false,
                "temperature_lock": false,
                "temperature_lock_high_temp": 22.22223,
                "temperature_lock_low_temp": 20,
                "temperature_lock_pin_hash": "",
                "temperature_scale": "F",
                "time_to_target": 0,
                "time_to_target_training": "ready",
                "tou_icon": false,
                "touched_by": {},
                "upper_safety_temp": 35,
                "upper_safety_temp_enabled": false,
                "weave_device_id_self_reported": "64166600009892A8",
                "where_id": "00000000-0000-0000-0000-000100000009",
                "wiring_error": "",
                "wiring_error_timestamp": 1550122274,
                "y2_type": "unknown"
            }
        },
        {
            "object_key": "geofence_info.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
            "object_revision": 232,
            "object_timestamp": 1550428971425,
            "value": {
                "combined_presence": {
                    "presence": 0,
                    "presence_evaluation_timestamp": 1550428970410,
                    "raw_presence": 0
                },
                "device_events": [
                    {
                        "device_id": "d28e6d35-1e3b-4f90-9e0a-db90b887bb8c",
                        "events": [
                            {
                                "fence_id": "a08e1370-1ab1-11e9-a955-0e166ab4c5b8:GZGmAKka",
                                "state": 0,
                                "time_of_state_transition": 1550428970043,
                                "time_state_reported": 1550428970410,
                                "user_id": "13268723"
                            }
                        ],
                        "is_attached_to_occupancy": true,
                        "is_primary_device": true
                    },
                    {
                        "device_id": "7C931E7C-BD8D-4208-84CC-3BBA5B0E3555",
                        "events": [
                            {
                                "fence_id": "a08e1370-1ab1-11e9-a955-0e166ab4c5b8:GZGmAKka",
                                "state": 1,
                                "time_of_state_transition": 1550425600000,
                                "time_state_reported": 1550425600000,
                                "user_id": "1923020"
                            }
                        ],
                        "is_attached_to_occupancy": true,
                        "is_primary_device": true
                    }
                ],
                "fences": [
                    {
                        "direction": [
                            0,
                            1
                        ],
                        "fence_id": "a08e1370-1ab1-11e9-a955-0e166ab4c5b8:GZGmAKka",
                        "latitude": 37.43855462778808,
                        "longitude": -122.18863986432552,
                        "radius": 200
                    }
                ]
            }
        },
        {
            "object_key": "tou.09AA01AC421818BS",
            "object_revision": -1,
            "object_timestamp": 1,
            "value": {}
        },
        {
            "object_key": "safety_summary.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
            "object_revision": -8354,
            "object_timestamp": 1547768582893,
            "value": {}
        },
        {
            "object_key": "quartz.f67bb3d58fbf4602b0ace3a27382aae6",
            "object_revision": 23071,
            "object_timestamp": 1550398210129,
            "value": {
                "activation_time": 1550163786696,
                "audio_input_enabled": true,
                "camera_type": 12,
                "capabilities": [
                    "statusled.watching.on_off",
                    "image.properties",
                    "force_idr",
                    "streaming.data-usage-tier.120",
                    "streaming.cameraprofile.VIDEO_H264_100KBIT_L30",
                    "spoken_locale.da_DK",
                    "aspect_ratio_4x3",
                    "streaming.start-stop",
                    "audio.microphone",
                    "spoken_locale.en_US",
                    "night_vision",
                    "indoor_chime",
                    "spoken_locale.sv_SE",
                    "streaming.params",
                    "streaming.cameraprofile.VIDEO_H264_2MBIT_L40",
                    "stranger_detection",
                    "spoken_locale.es_US",
                    "spoken_locale.de_DE",
                    "statusled",
                    "spoken_locale.it_IT",
                    "streaming.data-usage-tier.300",
                    "streaming.cameraprofile.VIDEO_H264_530KBIT_L31",
                    "spoken_locale.fr_FR",
                    "sharing",
                    "backfill",
                    "watermark",
                    "spoken_locale.nb_NO",
                    "spoken_locale.es_ES",
                    "labs",
                    "spoken_locale.fr_CA",
                    "audio.speaker",
                    "talkback.chime.on_off",
                    "dptz",
                    "irled",
                    "notify.email",
                    "standby",
                    "detectors.on_camera",
                    "spoken_locale.en_GB",
                    "spoken_locale.fi_FI",
                    "spoken_locale.nl_NL",
                    "spoken_locale",
                    "dptz.8x",
                    "streaming.data-usage-tier.30",
                    "statusled.watching",
                    "statusled.on_off"
                ],
                "cvr_enrolled": "5-days",
                "description": "",
                "direct_nexustalk_host": "oculus3524-us1.dropcam.com",
                "download_host": "oculus3524-us1.dropcam.com:80",
                "fabric_id": "ABEA30BD2353598A",
                "ip_address": "192.168.225.69",
                "last_connect_time": 1550398207000,
                "last_disconnect_reason": "no_reason",
                "last_disconnect_time": 1550398207000,
                "live_stream_host": "oculus3524-us1.dropcam.com:1935",
                "mac_address": "6416666b504e",
                "model": "Nest Hello",
                "nexus_api_http_server_url": "https://nexusapi-us1.dropcam.com",
                "preview_streaming_enabled": true,
                "public_share_enabled": false,
                "recorded_stream_host": "oculus3524-us1.dropcam.com:1935",
                "serial_number": "19AA01AE3018081X",
                "snapshot_url": {
                    "snapshot_url_prefix": "https://www.dropcam.com/api/wwn.get_snapshot",
                    "snapshot_url_suffix": ""
                },
                "software_version": "4080019",
                "streaming_state": "streaming-enabled",
                "structure_id": "a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                "weave_device_id": "64166600006B504E",
                "websocket_nexustalk_host": "oculus3524-us1.dropcam.com:80",
                "where_id": "00000000-0000-0000-0000-00010000001b",
                "wwn_stream_host": "stream-us1-alfa.dropcam.com"
            }
        },
        {
            "object_key": "track.09AA01AC421818BS",
            "object_revision": -7599,
            "object_timestamp": 1550452819736,
            "value": {
                "last_connection": 1550452819736,
                "last_ip": "/172.16.42.140",
                "online": true
            }
        },
        {
            "object_key": "device_alert_dialog.09AA01AC421818BS",
            "object_revision": -22638,
            "object_timestamp": 1548128703035,
            "value": {
                "dialog_data": "",
                "dialog_id": "confirm-pairing"
            }
        },
        {
            "object_key": "kryptonite.18B430C79EF7406D",
            "object_revision": 9367,
            "object_timestamp": 1550450720109,
            "value": {
                "battery_level": 100,
                "current_temperature": 20.2,
                "last_updated_at": 1550450718,
                "model": "KR1",
                "serial_number": "22AA01AC291806C0",
                "structure_id": "a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                "where_id": "00000000-0000-0000-0000-00010000000b"
            }
        },
        {
            "object_key": "schedule.09AA01AC421818BS",
            "object_revision": -13921,
            "object_timestamp": 1550390650362,
            "value": {
                "days": {
                    "0": {
                        "0": {
                            "entry_type": "continuation",
                            "temp": 16.91875,
                            "time": 0,
                            "touched_at": 1550153744,
                            "touched_by": 1,
                            "touched_tzo": -28800,
                            "type": "HEAT"
                        },
                        "1": {
                            "entry_type": "setpoint",
                            "temp": 18.43181,
                            "time": 28800,
                            "touched_at": 1549436255,
                            "touched_by": 5,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "2": {
                            "entry_type": "setpoint",
                            "temp": 20.16695,
                            "time": 36000,
                            "touched_at": 1549995682,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "3": {
                            "entry_type": "setpoint",
                            "temp": 19.81856,
                            "time": 65700,
                            "touched_at": 1548911707,
                            "touched_by": 6,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.1923020",
                            "type": "HEAT"
                        },
                        "4": {
                            "entry_type": "setpoint",
                            "temp": 16.7903,
                            "time": 79200,
                            "touched_at": 1548919394,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        }
                    },
                    "1": {
                        "0": {
                            "entry_type": "continuation",
                            "temp": 16.7903,
                            "time": 0,
                            "touched_at": 1550153744,
                            "touched_by": 1,
                            "touched_tzo": -28800,
                            "type": "HEAT"
                        },
                        "1": {
                            "entry_type": "setpoint",
                            "temp": 18.31363,
                            "time": 28800,
                            "touched_at": 1549436262,
                            "touched_by": 5,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "2": {
                            "entry_type": "setpoint",
                            "temp": 20.25208,
                            "time": 36000,
                            "touched_at": 1549995679,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "3": {
                            "entry_type": "setpoint",
                            "temp": 20.16698,
                            "time": 65700,
                            "touched_at": 1548919230,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "4": {
                            "entry_type": "setpoint",
                            "temp": 16.91875,
                            "time": 79200,
                            "touched_at": 1548919405,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        }
                    },
                    "2": {
                        "0": {
                            "entry_type": "continuation",
                            "temp": 16.91875,
                            "time": 0,
                            "touched_at": 1550153744,
                            "touched_by": 1,
                            "touched_tzo": -28800,
                            "type": "HEAT"
                        },
                        "1": {
                            "entry_type": "setpoint",
                            "temp": 18.11661,
                            "time": 28800,
                            "touched_at": 1549436271,
                            "touched_by": 5,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "2": {
                            "entry_type": "setpoint",
                            "temp": 20.16695,
                            "time": 36000,
                            "touched_at": 1549995688,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "3": {
                            "entry_type": "setpoint",
                            "temp": 20.25208,
                            "time": 65700,
                            "touched_at": 1548919234,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "4": {
                            "entry_type": "setpoint",
                            "temp": 16.91875,
                            "time": 79200,
                            "touched_at": 1548919396,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        }
                    },
                    "3": {
                        "0": {
                            "entry_type": "continuation",
                            "temp": 16.91875,
                            "time": 0,
                            "touched_at": 1550153744,
                            "touched_by": 1,
                            "touched_tzo": -28800,
                            "type": "HEAT"
                        },
                        "1": {
                            "entry_type": "setpoint",
                            "temp": 18.55002,
                            "time": 28800,
                            "touched_at": 1549436276,
                            "touched_by": 5,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "2": {
                            "entry_type": "setpoint",
                            "temp": 20.16695,
                            "time": 36000,
                            "touched_at": 1549995686,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "3": {
                            "entry_type": "setpoint",
                            "temp": 20.25208,
                            "time": 65700,
                            "touched_at": 1548919237,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "4": {
                            "entry_type": "setpoint",
                            "temp": 16.91875,
                            "time": 79200,
                            "touched_at": 1548919400,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        }
                    },
                    "4": {
                        "0": {
                            "entry_type": "continuation",
                            "temp": 16.91875,
                            "time": 0,
                            "touched_at": 1550153744,
                            "touched_by": 1,
                            "touched_tzo": -28800,
                            "type": "HEAT"
                        },
                        "1": {
                            "entry_type": "setpoint",
                            "temp": 18.39243,
                            "time": 28800,
                            "touched_at": 1549436283,
                            "touched_by": 5,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "2": {
                            "entry_type": "setpoint",
                            "temp": 20.16695,
                            "time": 36000,
                            "touched_at": 1549995690,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "3": {
                            "entry_type": "setpoint",
                            "temp": 20.25208,
                            "time": 65700,
                            "touched_at": 1548919241,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "4": {
                            "entry_type": "setpoint",
                            "temp": 16.91875,
                            "time": 79200,
                            "touched_at": 1548919397,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        }
                    },
                    "5": {
                        "0": {
                            "entry_type": "continuation",
                            "temp": 16.91875,
                            "time": 0,
                            "touched_at": 1550153744,
                            "touched_by": 1,
                            "touched_tzo": -28800,
                            "type": "HEAT"
                        },
                        "1": {
                            "entry_type": "setpoint",
                            "temp": 17.40738,
                            "time": 28800,
                            "touched_at": 1549995712,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "2": {
                            "entry_type": "setpoint",
                            "temp": 20.16695,
                            "time": 36000,
                            "touched_at": 1549995695,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "3": {
                            "entry_type": "setpoint",
                            "temp": 20.25208,
                            "time": 65700,
                            "touched_at": 1548919246,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "4": {
                            "entry_type": "setpoint",
                            "temp": 16.91875,
                            "time": 79200,
                            "touched_at": 1548919402,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        }
                    },
                    "6": {
                        "0": {
                            "entry_type": "continuation",
                            "temp": 16.91875,
                            "time": 0,
                            "touched_at": 1550153744,
                            "touched_by": 1,
                            "touched_tzo": -28800,
                            "type": "HEAT"
                        },
                        "1": {
                            "entry_type": "setpoint",
                            "temp": 17.04488,
                            "time": 28800,
                            "touched_at": 1549995709,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "2": {
                            "entry_type": "setpoint",
                            "temp": 20.16695,
                            "time": 36000,
                            "touched_at": 1549995693,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "3": {
                            "entry_type": "setpoint",
                            "temp": 20.25208,
                            "time": 65700,
                            "touched_at": 1548919249,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        },
                        "4": {
                            "entry_type": "setpoint",
                            "temp": 16.91875,
                            "time": 79200,
                            "touched_at": 1548919398,
                            "touched_by": 4,
                            "touched_tzo": -28800,
                            "touched_user_id": "user.13268723",
                            "type": "HEAT"
                        }
                    }
                },
                "name": "Current Schedule",
                "schedule_mode": "HEAT",
                "ver": 2
            }
        },
        {
            "object_key": "topaz.6416660000C66097",
            "object_revision": 12060,
            "object_timestamp": 1550395281518,
            "value": {
                "auto_away": true,
                "battery_health_state": 0,
                "battery_level": 5287,
                "buzzer_test_results": 32769,
                "capability_level": 2,
                "certification_body": 1,
                "co_blame_duration": 0,
                "co_blame_threshold": 0,
                "co_previous_peak": 0,
                "co_sequence_number": 0,
                "co_status": 0,
                "component_als_test_passed": true,
                "component_buzzer_test_passed": true,
                "component_co_test_passed": true,
                "component_heat_test_passed": true,
                "component_hum_test_passed": true,
                "component_led_test_passed": true,
                "component_pir_test_passed": true,
                "component_smoke_test_passed": true,
                "component_speaker_test_passed": true,
                "component_temp_test_passed": true,
                "component_us_test_passed": true,
                "component_wifi_test_passed": true,
                "creation_time": 1548988193992,
                "device_born_on_date_utc_secs": 1541808000,
                "device_external_color": "white",
                "device_locale": "en_US",
                "fabric_id": "ABEA30BD2353598A",
                "factory_loaded_languages": "en_US,es_US",
                "gesture_hush_enable": false,
                "heads_up_enable": true,
                "home_alarm_link_capable": false,
                "home_away_input": true,
                "hushed_state": false,
                "installed_locale": "en_US",
                "kl_software_version": "3.0.16",
                "latest_manual_test_cancelled": true,
                "latest_manual_test_end_utc_secs": 1549867865,
                "latest_manual_test_start_utc_secs": 1549867865,
                "line_power_present": false,
                "model": "Topaz-2.7",
                "night_light_continuous": false,
                "night_light_enable": false,
                "ntp_green_led_brightness": 2,
                "ntp_green_led_enable": true,
                "product_id": 9,
                "replace_by_date_utc_secs": 1857427200,
                "resource_id": "topaz_resource.5",
                "serial_number": "06AA01AC451808T3",
                "smoke_sequence_number": 0,
                "smoke_status": 0,
                "software_version": "3.1.4rc3",
                "speaker_test_results": 32769,
                "spoken_where_id": "00000000-0000-0000-0000-000100000002",
                "steam_detection_enable": true,
                "structure_id": "a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                "thread_ip_address": [
                    "fe80:0000:0000:0000:6616:6600:00c6:6097",
                    "fdbd:2353:598a:0002:6616:6600:00c6:6097"
                ],
                "thread_mac_address": "6416660000c66097",
                "where_id": "00000000-0000-0000-0000-000100000002",
                "wifi_ip_address": "192.168.225.23",
                "wifi_mac_address": "641666ca6eaa",
                "wifi_regulatory_domain": "A2",
                "wired_led_enable": true,
                "wired_or_battery": 1
            }
        },
        {
            "object_key": "topaz_resource.5",
            "object_revision": -1,
            "object_timestamp": 1503446400000,
            "value": {
                "supported_locales": [
                    "en_US",
                    "es_US",
                    "en_GB",
                    "fr_CA",
                    "fr_FR",
                    "nl_NL",
                    "de_DE",
                    "it_IT",
                    "es_ES",
                    "da_DK",
                    "sv_SE",
                    "fi_FI",
                    "nb_NO"
                ],
                "wheres": [
                    "00000000-0000-0000-0000-000100000000",
                    "00000000-0000-0000-0000-000100000001",
                    "00000000-0000-0000-0000-000100000002",
                    "00000000-0000-0000-0000-000100000003",
                    "00000000-0000-0000-0000-000100000004",
                    "00000000-0000-0000-0000-000100000005",
                    "00000000-0000-0000-0000-000100000006",
                    "00000000-0000-0000-0000-000100000007",
                    "00000000-0000-0000-0000-000100000008",
                    "00000000-0000-0000-0000-000100000009",
                    "00000000-0000-0000-0000-00010000000a",
                    "00000000-0000-0000-0000-00010000000b",
                    "00000000-0000-0000-0000-00010000000c",
                    "00000000-0000-0000-0000-00010000000d",
                    "00000000-0000-0000-0000-00010000000e",
                    "00000000-0000-0000-0000-00010000000f",
                    "00000000-0000-0000-0000-000100000010"
                ]
            }
        },
        {
            "object_key": "delayed_topaz.6416660000C66097",
            "object_revision": 27328,
            "object_timestamp": 1550395281640,
            "value": {
                "NI_6416660000C66097": {
                    "object_timestamp": 1550395281518,
                    "value": {
                        "auto_away": true,
                        "battery_health_state": 0,
                        "battery_level": 5287,
                        "buzzer_test_results": 32769,
                        "capability_level": 2,
                        "certification_body": 1,
                        "co_blame_duration": 0,
                        "co_blame_threshold": 0,
                        "co_previous_peak": 0,
                        "co_sequence_number": 0,
                        "co_status": 0,
                        "component_als_test_passed": true,
                        "component_buzzer_test_passed": true,
                        "component_co_test_passed": true,
                        "component_heat_test_passed": true,
                        "component_hum_test_passed": true,
                        "component_led_test_passed": true,
                        "component_pir_test_passed": true,
                        "component_smoke_test_passed": true,
                        "component_speaker_test_passed": true,
                        "component_temp_test_passed": true,
                        "component_us_test_passed": true,
                        "component_wifi_test_passed": true,
                        "creation_time": 1548988193992,
                        "device_born_on_date_utc_secs": 1541808000,
                        "device_external_color": "white",
                        "device_locale": "en_US",
                        "fabric_id": "ABEA30BD2353598A",
                        "factory_loaded_languages": "en_US,es_US",
                        "gesture_hush_enable": false,
                        "heads_up_enable": true,
                        "home_alarm_link_capable": false,
                        "home_away_input": true,
                        "hushed_state": false,
                        "installed_locale": "en_US",
                        "kl_software_version": "3.0.16",
                        "latest_manual_test_cancelled": true,
                        "latest_manual_test_end_utc_secs": 1549867865,
                        "latest_manual_test_start_utc_secs": 1549867865,
                        "line_power_present": false,
                        "model": "Topaz-2.7",
                        "night_light_continuous": false,
                        "night_light_enable": false,
                        "ntp_green_led_brightness": 2,
                        "ntp_green_led_enable": true,
                        "product_id": 9,
                        "replace_by_date_utc_secs": 1857427200,
                        "resource_id": "topaz_resource.5",
                        "serial_number": "06AA01AC451808T3",
                        "smoke_sequence_number": 0,
                        "smoke_status": 0,
                        "software_version": "3.1.4rc3",
                        "speaker_test_results": 32769,
                        "spoken_where_id": "00000000-0000-0000-0000-000100000002",
                        "steam_detection_enable": true,
                        "structure_id": "a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                        "thread_ip_address": [
                            "fe80:0000:0000:0000:6616:6600:00c6:6097",
                            "fdbd:2353:598a:0002:6616:6600:00c6:6097"
                        ],
                        "thread_mac_address": "6416660000c66097",
                        "where_id": "00000000-0000-0000-0000-000100000002",
                        "wifi_ip_address": "192.168.225.23",
                        "wifi_mac_address": "641666ca6eaa",
                        "wifi_regulatory_domain": "A2",
                        "wired_led_enable": true,
                        "wired_or_battery": 1
                    }
                }
            }
        },
        {
            "object_key": "user_alert_dialog.13268723",
            "object_revision": 21167,
            "object_timestamp": 1548128703035,
            "value": {
                "dialog_data": "",
                "dialog_id": "confirm-pairing"
            }
        },
        {
            "object_key": "structure_metadata.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
            "object_revision": -12855,
            "object_timestamp": 1550399071442,
            "value": {
                "entitlements_api_path": "/api/0.1/structure/a08e1370-1ab1-11e9-a955-0e166ab4c5b8/entitlements?v=a5b1cc36-0e02-47de-9127-69ac338c90fd",
                "structure_history_api_path": "/api/0.1/structure/a08e1370-1ab1-11e9-a955-0e166ab4c5b8/history?v=1550399071439"
            }
        },
        {
            "object_key": "message_center.13268723",
            "object_revision": -7738,
            "object_timestamp": 1550163909459,
            "value": {
                "messages": [
                    {
                        "dismissed": true,
                        "id": "f2c61740-1ab2-11e9-925e-0aae12655e82",
                        "key": "family_accounts_invite_claimed",
                        "parameters": [
                            "a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                            "Hobart Chalet",
                            "Kirsten",
                            "kirstenm.stewart@gmail.com",
                            "Jayz"
                        ],
                        "priority": 1,
                        "read": true,
                        "thread_id": 0,
                        "timestamp": 1547769149
                    },
                    {
                        "dismissed": true,
                        "id": "3f387b80-1e1a-11e9-925e-0aae12655e82",
                        "key": "family_accounts_invite_claimed",
                        "parameters": [
                            "a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                            "Hobart Chalet",
                            "Cindy",
                            "cindy@cindybarrett.net",
                            "Jayz"
                        ],
                        "priority": 1,
                        "read": true,
                        "thread_id": 0,
                        "timestamp": 1548143369
                    },
                    {
                        "dismissed": true,
                        "id": "fa413850-3061-11e9-af79-160c72d19936",
                        "key": "protect_smoke_warn",
                        "parameters": [
                            "6416660000C66097",
                            "00000000-0000-0000-0000-000100000002",
                            "Hallway",
                            "",
                            "Hobart Chalet",
                            2,
                            1,
                            1,
                            2,
                            1,
                            1,
                            "a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                            "America/Los_Angeles"
                        ],
                        "priority": 1,
                        "read": true,
                        "thread_id": 1550153297990,
                        "timestamp": 1550153298
                    },
                    {
                        "dismissed": true,
                        "id": "20820170-3062-11e9-af79-160c72d19936",
                        "key": "protect_smoke_warn_clear",
                        "parameters": [
                            "6416660000C66097",
                            "00000000-0000-0000-0000-000100000002",
                            "Hallway",
                            "",
                            "Hobart Chalet",
                            0,
                            0,
                            0,
                            2,
                            1,
                            1,
                            "a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                            "America/Los_Angeles"
                        ],
                        "priority": 1,
                        "read": true,
                        "thread_id": 1550153297990,
                        "timestamp": 1550153362
                    }
                ]
            }
        },
        {
            "object_key": "kryptonite.18B430C49E3DA2A2",
            "object_revision": -1543,
            "object_timestamp": 1550452520332,
            "value": {
                "battery_level": 100,
                "current_temperature": 20.29999,
                "last_updated_at": 1550452518,
                "model": "KR1",
                "serial_number": "22AA01AC29180EAL",
                "structure_id": "a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                "where_id": "00000000-0000-0000-0000-000100000010"
            }
        },
        {
            "object_key": "structure_history.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
            "object_revision": 15226,
            "object_timestamp": 1547768582893,
            "value": {
                "events": [],
                "products": [],
                "schema_version": "1.0"
            }
        },
        {
            "object_key": "user.13268723",
            "object_revision": -8345,
            "object_timestamp": 1550167255253,
            "value": {
                "acknowledged_onboarding_screens": [
                    "family_account_education_a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
                    "rcs",
                    "camera_chime_assist"
                ],
                "email": "nest@bf2e.com",
                "is_merged_with_gaia": false,
                "name": "nest@bf2e.com",
                "short_name": "Jayz",
                "structure_memberships": [
                    {
                        "roles": [
                            "owner"
                        ],
                        "structure": "structure.a08e1370-1ab1-11e9-a955-0e166ab4c5b8"
                    }
                ],
                "structures": [
                    "structure.a08e1370-1ab1-11e9-a955-0e166ab4c5b8"
                ]
            }
        },
        {
            "object_key": "trip.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
            "object_revision": 12669,
            "object_timestamp": 1547768582893,
            "value": {
                "trips": []
            }
        },
        {
            "object_key": "link.09AA01AC421818BS",
            "object_revision": -24201,
            "object_timestamp": 1548128703035,
            "value": {
                "structure": "structure.a08e1370-1ab1-11e9-a955-0e166ab4c5b8"
            }
        },
        {
            "object_key": "demand_response.09AA01AC421818BS",
            "object_revision": -4225,
            "object_timestamp": 1547786576459,
            "value": {}
        },
        {
            "object_key": "message.09AA01AC421818BS",
            "object_revision": 9766,
            "object_timestamp": 1548128707197,
            "value": {}
        },
        {
            "object_key": "structure.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
            "object_revision": 30230,
            "object_timestamp": 1550429151221,
            "value": {
                "address_lines": [
                    "595 Hobart St"
                ],
                "aux_primary_fabric_id": "ABEA30BD2353598A",
                "away": false,
                "away_setter": 0,
                "city": "Menlo Park",
                "clutches": [],
                "country_code": "US",
                "creation_time": 1547768581672,
                "demand_charge_enabled": true,
                "devices": [
                    "device.09AA01AC421818BS"
                ],
                "diamond_changed_location": false,
                "dr_reminder_enabled": true,
                "eta_preconditioning_active": false,
                "fabric_ids": [
                    "ABEA30BD2353598A"
                ],
                "geofence_enhanced_autoaway": {
                    "enabled": true,
                    "is_set": true
                },
                "goose_sensor_events_enabled": true,
                "house_type": "family",
                "hvac_safety_shutoff_enabled": true,
                "hvac_smoke_safety_shutoff_enabled": false,
                "ifj_primary_fabric_id": "ABEA30BD2353598A",
                "latitude": 37.43855462778808,
                "location": "West Menlo Park, CA",
                "longitude": -122.18863986432552,
                "manual_away_timestamp": 1550429150,
                "manual_eco_all": false,
                "manual_eco_timestamp": 0,
                "measurement_scale": "imperial",
                "members": [
                    {
                        "roles": [
                            "owner"
                        ],
                        "user": "user.13268723"
                    },
                    {
                        "roles": [
                            "member"
                        ],
                        "user": "user.13333542"
                    },
                    {
                        "roles": [
                            "member"
                        ],
                        "user": "user.1923020"
                    }
                ],
                "name": "Hobart Chalet",
                "num_thermostats": "1",
                "postal_code": "94025",
                "rcs_sensor_swarm": [
                    "kryptonite.18B430C79EF7406D",
                    "kryptonite.18B430C49E3DA2A2",
                    "kryptonite.18B430CEBD24D3CF",
                    "kryptonite.18B430CFD99E3ABC"
                ],
                "renovation_date": "pre-1940",
                "state": "CA",
                "structure_area": 74.323,
                "swarm": [
                    "kryptonite.18B430C79EF7406D",
                    "topaz.6416660000C66097",
                    "kryptonite.18B430CEBD24D3CF",
                    "quartz.f67bb3d58fbf4602b0ace3a27382aae6",
                    "device.09AA01AC421818BS",
                    "kryptonite.18B430CFD99E3ABC",
                    "kryptonite.18B430C49E3DA2A2"
                ],
                "time_zone": "America/Los_Angeles",
                "topaz_away": false,
                "topaz_hush_key": "ZCHkgEMxG0adK8n1Si5DEw==",
                "tou_enabled": true,
                "touched_by": {
                    "touched_by": 3,
                    "touched_id": "",
                    "touched_user_id": "user.1923020"
                },
                "user": "user.13268723",
                "vacation_mode": false
            }
        },
        {
            "object_key": "where.a08e1370-1ab1-11e9-a955-0e166ab4c5b8",
            "object_revision": -12901,
            "object_timestamp": 1549513557324,
            "value": {
                "wheres": [
                    {
                        "name": "Attic",
                        "where_id": "00000000-0000-0000-0000-000100000004"
                    },
                    {
                        "name": "Back Door",
                        "where_id": "00000000-0000-0000-0000-00010000001d"
                    },
                    {
                        "name": "Backyard",
                        "where_id": "00000000-0000-0000-0000-000100000011"
                    },
                    {
                        "name": "Basement",
                        "where_id": "00000000-0000-0000-0000-000100000001"
                    },
                    {
                        "name": "Bathroom",
                        "where_id": "00000000-0000-0000-0000-000100000009"
                    },
                    {
                        "name": "Bedroom",
                        "where_id": "00000000-0000-0000-0000-00010000000d"
                    },
                    {
                        "name": "Deck",
                        "where_id": "00000000-0000-0000-0000-000100000017"
                    },
                    {
                        "name": "Den",
                        "where_id": "00000000-0000-0000-0000-000100000003"
                    },
                    {
                        "name": "Dining Room",
                        "where_id": "00000000-0000-0000-0000-000100000010"
                    },
                    {
                        "name": "Downstairs",
                        "where_id": "00000000-0000-0000-0000-000100000006"
                    },
                    {
                        "name": "Driveway",
                        "where_id": "00000000-0000-0000-0000-000100000012"
                    },
                    {
                        "name": "Entryway",
                        "where_id": "00000000-0000-0000-0000-000100000000"
                    },
                    {
                        "name": "Family Room",
                        "where_id": "00000000-0000-0000-0000-00010000000b"
                    },
                    {
                        "name": "Front Door",
                        "where_id": "00000000-0000-0000-0000-00010000001b"
                    },
                    {
                        "name": "Front Yard",
                        "where_id": "00000000-0000-0000-0000-000100000013"
                    },
                    {
                        "name": "Garage",
                        "where_id": "00000000-0000-0000-0000-000100000007"
                    },
                    {
                        "name": "Guest House",
                        "where_id": "00000000-0000-0000-0000-000100000015"
                    },
                    {
                        "name": "Guest Room",
                        "where_id": "00000000-0000-0000-0000-00010000001a"
                    },
                    {
                        "name": "Hallway",
                        "where_id": "00000000-0000-0000-0000-000100000002"
                    },
                    {
                        "name": "Kids Room",
                        "where_id": "00000000-0000-0000-0000-000100000008"
                    },
                    {
                        "name": "Kitchen",
                        "where_id": "00000000-0000-0000-0000-00010000000a"
                    },
                    {
                        "name": "Living Room",
                        "where_id": "00000000-0000-0000-0000-00010000000c"
                    },
                    {
                        "name": "Master Bedroom",
                        "where_id": "00000000-0000-0000-0000-000100000005"
                    },
                    {
                        "name": "Office",
                        "where_id": "00000000-0000-0000-0000-00010000000e"
                    },
                    {
                        "name": "Outside",
                        "where_id": "00000000-0000-0000-0000-000100000014"
                    },
                    {
                        "name": "Patio",
                        "where_id": "00000000-0000-0000-0000-000100000018"
                    },
                    {
                        "name": "Shed",
                        "where_id": "00000000-0000-0000-0000-000100000016"
                    },
                    {
                        "name": "Side Door",
                        "where_id": "00000000-0000-0000-0000-00010000001c"
                    },
                    {
                        "name": "Upstairs",
                        "where_id": "00000000-0000-0000-0000-00010000000f"
                    },
                    {
                        "name": "Mud Room",
                        "where_id": "00000000-0000-0000-5e73-405f2b691208"
                    }
                ]
            }
        }
    ],
    "weather_for_structures": {
        "structure.a08e1370-1ab1-11e9-a955-0e166ab4c5b8": {
            "current": {
                "icon": "partlycloudy",
                "sunrise": "1550415360",
                "sunset": "1550454600",
                "temp_c": "9.4"
            },
            "location": {
                "city": "West Menlo Park",
                "country": "US",
                "state": "CA",
                "zip": "94025"
            }
        }
    }
}
*/
