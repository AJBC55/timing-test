package main

import (
	"encoding/json"
	"fmt"
)

type TimingMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func (tm *TimingMessage) Unmarshal(data []byte) error {
	temp := struct {
		Type string
		Data json.RawMessage
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	// set type
	tm.Type = temp.Type
	switch temp.Type {
	//parse heart beat
	case "F":
		var data Heartbeat
		if err := json.Unmarshal(temp.Data, &data); err != nil {
			return err
		}
		tm.Data = data
	case "A":
		var data CompetitorInfo
		if err := json.Unmarshal(temp.Data, &data); err != nil {
			return err
		}
		tm.Data = data
	case "COMP":
		var data CompInfo
		if err := json.Unmarshal(temp.Data, &data); err != nil {
			return err
		}
		tm.Data = data
	case "B":
		var data RunInfo
		if err := json.Unmarshal(temp.Data, &data); err != nil {
			return err
		}
		tm.Data = data
	case "C":
		var data ClassInfo
		if err := json.Unmarshal(temp.Data, &data); err != nil {
			return err
		}
		tm.Data = data
	case "E":
		var data SettingInfo
		if err := json.Unmarshal(temp.Data, &data); err != nil {
			return err
		}
		tm.Data = data
	case "G":
		var data RaceInfo
		if err := json.Unmarshal(temp.Data, &data); err != nil {
			return err
		}
		tm.Data = data
	case "H":
		var data PracticeQualifyInfo
		if err := json.Unmarshal(temp.Data, &data); err != nil {
			return err
		}
		tm.Data = data
	case "I":
		var data InitRecord
		if err := json.Unmarshal(temp.Data, &data); err != nil {
			return err
		}
		tm.Data = data
	case "J":
		var data PassingInfo
		if err := json.Unmarshal(temp.Data, &data); err != nil {
			return err
		}
		tm.Data = data
	case "COR":
		var data CorrectedFinish
		if err := json.Unmarshal(temp.Data, &data); err != nil {
			return err
		}
		tm.Data = data
	default:
		return fmt.Errorf("TYPE %s NOT RECONIZED", temp.Type)
	}
	return nil

}
