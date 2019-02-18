package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
   THIS IS AN AUTOMATICALLY GENERATED FILE

   DO NOT EDIT!
*/

// Buckets returns updated buckets items matching key
// prefix=KeyPrefixBuckets.
func (ns *NestState) Buckets() ([]*Buckets, error) {
	bucketsSlice := []*Buckets{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixBuckets+".") {
			buckets := &Buckets{}
			if err := json.Unmarshal(v, &buckets); err != nil {
				return nil, fmt.Errorf("unmarshalling to Buckets: %v", err)
			}
			bucketsSlice = append(bucketsSlice, buckets)
		}
	}
	return bucketsSlice, nil
}

// DemandResponses returns updated buckets items matching key
// prefix=KeyPrefixDemandResponse.
func (ns *NestState) DemandResponses() ([]*DemandResponse, error) {
	demandResponsesSlice := []*DemandResponse{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixDemandResponse+".") {
			demandResponse := &DemandResponse{}
			if err := json.Unmarshal(v, &demandResponse); err != nil {
				return nil, fmt.Errorf("unmarshalling to DemandResponse: %v", err)
			}
			demandResponsesSlice = append(demandResponsesSlice, demandResponse)
		}
	}
	return demandResponsesSlice, nil
}

// DetectInfos returns updated buckets items matching key
// prefix=KeyPrefixDetectInfo.
func (ns *NestState) DetectInfos() ([]*DetectInfo, error) {
	detectInfosSlice := []*DetectInfo{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixDetectInfo+".") {
			detectInfo := &DetectInfo{}
			if err := json.Unmarshal(v, &detectInfo); err != nil {
				return nil, fmt.Errorf("unmarshalling to DetectInfo: %v", err)
			}
			detectInfosSlice = append(detectInfosSlice, detectInfo)
		}
	}
	return detectInfosSlice, nil
}

// DetectLocales returns updated buckets items matching key
// prefix=KeyPrefixDetectLocales.
func (ns *NestState) DetectLocales() ([]*DetectLocales, error) {
	detectLocalesSlice := []*DetectLocales{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixDetectLocales+".") {
			detectLocales := &DetectLocales{}
			if err := json.Unmarshal(v, &detectLocales); err != nil {
				return nil, fmt.Errorf("unmarshalling to DetectLocales: %v", err)
			}
			detectLocalesSlice = append(detectLocalesSlice, detectLocales)
		}
	}
	return detectLocalesSlice, nil
}

// DetectSelfTests returns updated buckets items matching key
// prefix=KeyPrefixDetectSelfTest.
func (ns *NestState) DetectSelfTests() ([]*DetectSelfTest, error) {
	detectSelfTestsSlice := []*DetectSelfTest{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixDetectSelfTest+".") {
			detectSelfTest := &DetectSelfTest{}
			if err := json.Unmarshal(v, &detectSelfTest); err != nil {
				return nil, fmt.Errorf("unmarshalling to DetectSelfTest: %v", err)
			}
			detectSelfTestsSlice = append(detectSelfTestsSlice, detectSelfTest)
		}
	}
	return detectSelfTestsSlice, nil
}

// DetectStates returns updated buckets items matching key
// prefix=KeyPrefixDetectState.
func (ns *NestState) DetectStates() ([]*DetectState, error) {
	detectStatesSlice := []*DetectState{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixDetectState+".") {
			detectState := &DetectState{}
			if err := json.Unmarshal(v, &detectState); err != nil {
				return nil, fmt.Errorf("unmarshalling to DetectState: %v", err)
			}
			detectStatesSlice = append(detectStatesSlice, detectState)
		}
	}
	return detectStatesSlice, nil
}

// DeviceAlertDialogs returns updated buckets items matching key
// prefix=KeyPrefixDeviceAlertDialog.
func (ns *NestState) DeviceAlertDialogs() ([]*DeviceAlertDialog, error) {
	deviceAlertDialogsSlice := []*DeviceAlertDialog{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixDeviceAlertDialog+".") {
			deviceAlertDialog := &DeviceAlertDialog{}
			if err := json.Unmarshal(v, &deviceAlertDialog); err != nil {
				return nil, fmt.Errorf("unmarshalling to DeviceAlertDialog: %v", err)
			}
			deviceAlertDialogsSlice = append(deviceAlertDialogsSlice, deviceAlertDialog)
		}
	}
	return deviceAlertDialogsSlice, nil
}

// GeofenceInfos returns updated buckets items matching key
// prefix=KeyPrefixGeofenceInfo.
func (ns *NestState) GeofenceInfos() ([]*GeofenceInfo, error) {
	geofenceInfosSlice := []*GeofenceInfo{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixGeofenceInfo+".") {
			geofenceInfo := &GeofenceInfo{}
			if err := json.Unmarshal(v, &geofenceInfo); err != nil {
				return nil, fmt.Errorf("unmarshalling to GeofenceInfo: %v", err)
			}
			geofenceInfosSlice = append(geofenceInfosSlice, geofenceInfo)
		}
	}
	return geofenceInfosSlice, nil
}

// HelloCams returns updated buckets items matching key
// prefix=KeyPrefixHelloCam.
func (ns *NestState) HelloCams() ([]*HelloCam, error) {
	helloCamsSlice := []*HelloCam{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixHelloCam+".") {
			helloCam := &HelloCam{}
			if err := json.Unmarshal(v, &helloCam); err != nil {
				return nil, fmt.Errorf("unmarshalling to HelloCam: %v", err)
			}
			helloCamsSlice = append(helloCamsSlice, helloCam)
		}
	}
	return helloCamsSlice, nil
}

// LearningThermostats returns updated buckets items matching key
// prefix=KeyPrefixLearningThermostat.
func (ns *NestState) LearningThermostats() ([]*LearningThermostat, error) {
	learningThermostatsSlice := []*LearningThermostat{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixLearningThermostat+".") {
			learningThermostat := &LearningThermostat{}
			if err := json.Unmarshal(v, &learningThermostat); err != nil {
				return nil, fmt.Errorf("unmarshalling to LearningThermostat: %v", err)
			}
			learningThermostatsSlice = append(learningThermostatsSlice, learningThermostat)
		}
	}
	return learningThermostatsSlice, nil
}

// LearningThermostatStates returns updated buckets items matching key
// prefix=KeyPrefixLearningThermostatState.
func (ns *NestState) LearningThermostatStates() ([]*LearningThermostatState, error) {
	learningThermostatStatesSlice := []*LearningThermostatState{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixLearningThermostatState+".") {
			learningThermostatState := &LearningThermostatState{}
			if err := json.Unmarshal(v, &learningThermostatState); err != nil {
				return nil, fmt.Errorf("unmarshalling to LearningThermostatState: %v", err)
			}
			learningThermostatStatesSlice = append(learningThermostatStatesSlice, learningThermostatState)
		}
	}
	return learningThermostatStatesSlice, nil
}

// Links returns updated buckets items matching key
// prefix=KeyPrefixLink.
func (ns *NestState) Links() ([]*Link, error) {
	linksSlice := []*Link{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixLink+".") {
			link := &Link{}
			if err := json.Unmarshal(v, &link); err != nil {
				return nil, fmt.Errorf("unmarshalling to Link: %v", err)
			}
			linksSlice = append(linksSlice, link)
		}
	}
	return linksSlice, nil
}

// Messages returns updated buckets items matching key
// prefix=KeyPrefixMessage.
func (ns *NestState) Messages() ([]*Message, error) {
	messagesSlice := []*Message{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixMessage+".") {
			message := &Message{}
			if err := json.Unmarshal(v, &message); err != nil {
				return nil, fmt.Errorf("unmarshalling to Message: %v", err)
			}
			messagesSlice = append(messagesSlice, message)
		}
	}
	return messagesSlice, nil
}

// MessageCenters returns updated buckets items matching key
// prefix=KeyPrefixMessageCenter.
func (ns *NestState) MessageCenters() ([]*MessageCenter, error) {
	messageCentersSlice := []*MessageCenter{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixMessageCenter+".") {
			messageCenter := &MessageCenter{}
			if err := json.Unmarshal(v, &messageCenter); err != nil {
				return nil, fmt.Errorf("unmarshalling to MessageCenter: %v", err)
			}
			messageCentersSlice = append(messageCentersSlice, messageCenter)
		}
	}
	return messageCentersSlice, nil
}

// Metadatas returns updated buckets items matching key
// prefix=KeyPrefixMetadata.
func (ns *NestState) Metadatas() ([]*Metadata, error) {
	metadatasSlice := []*Metadata{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixMetadata+".") {
			metadata := &Metadata{}
			if err := json.Unmarshal(v, &metadata); err != nil {
				return nil, fmt.Errorf("unmarshalling to Metadata: %v", err)
			}
			metadatasSlice = append(metadatasSlice, metadata)
		}
	}
	return metadatasSlice, nil
}

// Occupancys returns updated buckets items matching key
// prefix=KeyPrefixOccupancy.
func (ns *NestState) Occupancys() ([]*Occupancy, error) {
	occupancysSlice := []*Occupancy{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixOccupancy+".") {
			occupancy := &Occupancy{}
			if err := json.Unmarshal(v, &occupancy); err != nil {
				return nil, fmt.Errorf("unmarshalling to Occupancy: %v", err)
			}
			occupancysSlice = append(occupancysSlice, occupancy)
		}
	}
	return occupancysSlice, nil
}

// RCSSettings returns updated buckets items matching key
// prefix=KeyPrefixRCSSettings.
func (ns *NestState) RCSSettings() ([]*RCSSettings, error) {
	rCSSettingsSlice := []*RCSSettings{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixRCSSettings+".") {
			rCSSettings := &RCSSettings{}
			if err := json.Unmarshal(v, &rCSSettings); err != nil {
				return nil, fmt.Errorf("unmarshalling to RCSSettings: %v", err)
			}
			rCSSettingsSlice = append(rCSSettingsSlice, rCSSettings)
		}
	}
	return rCSSettingsSlice, nil
}

// SafetySummarys returns updated buckets items matching key
// prefix=KeyPrefixSafetySummary.
func (ns *NestState) SafetySummarys() ([]*SafetySummary, error) {
	safetySummarysSlice := []*SafetySummary{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixSafetySummary+".") {
			safetySummary := &SafetySummary{}
			if err := json.Unmarshal(v, &safetySummary); err != nil {
				return nil, fmt.Errorf("unmarshalling to SafetySummary: %v", err)
			}
			safetySummarysSlice = append(safetySummarysSlice, safetySummary)
		}
	}
	return safetySummarysSlice, nil
}

// Structures returns updated buckets items matching key
// prefix=KeyPrefixStructure.
func (ns *NestState) Structures() ([]*Structure, error) {
	structuresSlice := []*Structure{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixStructure+".") {
			structure := &Structure{}
			if err := json.Unmarshal(v, &structure); err != nil {
				return nil, fmt.Errorf("unmarshalling to Structure: %v", err)
			}
			structuresSlice = append(structuresSlice, structure)
		}
	}
	return structuresSlice, nil
}

// StructureHistorys returns updated buckets items matching key
// prefix=KeyPrefixStructureHistory.
func (ns *NestState) StructureHistorys() ([]*StructureHistory, error) {
	structureHistorysSlice := []*StructureHistory{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixStructureHistory+".") {
			structureHistory := &StructureHistory{}
			if err := json.Unmarshal(v, &structureHistory); err != nil {
				return nil, fmt.Errorf("unmarshalling to StructureHistory: %v", err)
			}
			structureHistorysSlice = append(structureHistorysSlice, structureHistory)
		}
	}
	return structureHistorysSlice, nil
}

// StructureMetadatas returns updated buckets items matching key
// prefix=KeyPrefixStructureMetadata.
func (ns *NestState) StructureMetadatas() ([]*StructureMetadata, error) {
	structureMetadatasSlice := []*StructureMetadata{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixStructureMetadata+".") {
			structureMetadata := &StructureMetadata{}
			if err := json.Unmarshal(v, &structureMetadata); err != nil {
				return nil, fmt.Errorf("unmarshalling to StructureMetadata: %v", err)
			}
			structureMetadatasSlice = append(structureMetadatasSlice, structureMetadata)
		}
	}
	return structureMetadatasSlice, nil
}

// TemperatureSensors returns updated buckets items matching key
// prefix=KeyPrefixTemperatureSensor.
func (ns *NestState) TemperatureSensors() ([]*TemperatureSensor, error) {
	temperatureSensorsSlice := []*TemperatureSensor{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixTemperatureSensor+".") {
			temperatureSensor := &TemperatureSensor{}
			if err := json.Unmarshal(v, &temperatureSensor); err != nil {
				return nil, fmt.Errorf("unmarshalling to TemperatureSensor: %v", err)
			}
			temperatureSensorsSlice = append(temperatureSensorsSlice, temperatureSensor)
		}
	}
	return temperatureSensorsSlice, nil
}

// ThermostatActivitys returns updated buckets items matching key
// prefix=KeyPrefixThermostatActivity.
func (ns *NestState) ThermostatActivitys() ([]*ThermostatActivity, error) {
	thermostatActivitysSlice := []*ThermostatActivity{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixThermostatActivity+".") {
			thermostatActivity := &ThermostatActivity{}
			if err := json.Unmarshal(v, &thermostatActivity); err != nil {
				return nil, fmt.Errorf("unmarshalling to ThermostatActivity: %v", err)
			}
			thermostatActivitysSlice = append(thermostatActivitysSlice, thermostatActivity)
		}
	}
	return thermostatActivitysSlice, nil
}

// Tracks returns updated buckets items matching key
// prefix=KeyPrefixTrack.
func (ns *NestState) Tracks() ([]*Track, error) {
	tracksSlice := []*Track{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixTrack+".") {
			track := &Track{}
			if err := json.Unmarshal(v, &track); err != nil {
				return nil, fmt.Errorf("unmarshalling to Track: %v", err)
			}
			tracksSlice = append(tracksSlice, track)
		}
	}
	return tracksSlice, nil
}

// Trips returns updated buckets items matching key
// prefix=KeyPrefixTrip.
func (ns *NestState) Trips() ([]*Trip, error) {
	tripsSlice := []*Trip{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixTrip+".") {
			trip := &Trip{}
			if err := json.Unmarshal(v, &trip); err != nil {
				return nil, fmt.Errorf("unmarshalling to Trip: %v", err)
			}
			tripsSlice = append(tripsSlice, trip)
		}
	}
	return tripsSlice, nil
}

// Tuneups returns updated buckets items matching key
// prefix=KeyPrefixTuneups.
func (ns *NestState) Tuneups() ([]*Tuneups, error) {
	tuneupsSlice := []*Tuneups{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixTuneups+".") {
			tuneups := &Tuneups{}
			if err := json.Unmarshal(v, &tuneups); err != nil {
				return nil, fmt.Errorf("unmarshalling to Tuneups: %v", err)
			}
			tuneupsSlice = append(tuneupsSlice, tuneups)
		}
	}
	return tuneupsSlice, nil
}

// Users returns updated buckets items matching key
// prefix=KeyPrefixUser.
func (ns *NestState) Users() ([]*User, error) {
	usersSlice := []*User{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixUser+".") {
			user := &User{}
			if err := json.Unmarshal(v, &user); err != nil {
				return nil, fmt.Errorf("unmarshalling to User: %v", err)
			}
			usersSlice = append(usersSlice, user)
		}
	}
	return usersSlice, nil
}

// UserAlertDialogs returns updated buckets items matching key
// prefix=KeyPrefixUserAlertDialog.
func (ns *NestState) UserAlertDialogs() ([]*UserAlertDialog, error) {
	userAlertDialogsSlice := []*UserAlertDialog{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixUserAlertDialog+".") {
			userAlertDialog := &UserAlertDialog{}
			if err := json.Unmarshal(v, &userAlertDialog); err != nil {
				return nil, fmt.Errorf("unmarshalling to UserAlertDialog: %v", err)
			}
			userAlertDialogsSlice = append(userAlertDialogsSlice, userAlertDialog)
		}
	}
	return userAlertDialogsSlice, nil
}

// UserSettings returns updated buckets items matching key
// prefix=KeyPrefixUserSettings.
func (ns *NestState) UserSettings() ([]*UserSettings, error) {
	userSettingsSlice := []*UserSettings{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixUserSettings+".") {
			userSettings := &UserSettings{}
			if err := json.Unmarshal(v, &userSettings); err != nil {
				return nil, fmt.Errorf("unmarshalling to UserSettings: %v", err)
			}
			userSettingsSlice = append(userSettingsSlice, userSettings)
		}
	}
	return userSettingsSlice, nil
}

// Wheres returns updated buckets items matching key
// prefix=KeyPrefixWhere.
func (ns *NestState) Wheres() ([]*Where, error) {
	wheresSlice := []*Where{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixWhere+".") {
			where := &Where{}
			if err := json.Unmarshal(v, &where); err != nil {
				return nil, fmt.Errorf("unmarshalling to Where: %v", err)
			}
			wheresSlice = append(wheresSlice, where)
		}
	}
	return wheresSlice, nil
}

// WidgetTracks returns updated buckets items matching key
// prefix=KeyPrefixWidgetTrack.
func (ns *NestState) WidgetTracks() ([]*WidgetTrack, error) {
	widgetTracksSlice := []*WidgetTrack{}
	for _, v := range ns.UpdatedBuckets {
		b := &Bucket{}
		if err := json.Unmarshal(v, b); err != nil {
			return nil, err
		}
		if strings.HasPrefix(b.ObjectKey, KeyPrefixWidgetTrack+".") {
			widgetTrack := &WidgetTrack{}
			if err := json.Unmarshal(v, &widgetTrack); err != nil {
				return nil, fmt.Errorf("unmarshalling to WidgetTrack: %v", err)
			}
			widgetTracksSlice = append(widgetTracksSlice, widgetTrack)
		}
	}
	return widgetTracksSlice, nil
}
