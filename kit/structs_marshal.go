package kit

// Thanks to MMJ -- aka themmj on Github -- for help with this...

// TODO / FIXME: really we should have all the MarshalJSON() methods needed to allow the whole
// Frame to be converted back to its true format.

import (
	"encoding/json"
)

// Our Factory and Unit types have some embedded structs for user convenience, but this seems to
// mean (as far as I can tell) that we need custom unmarshalling for them... (?)

// WARNING - we actually cannot include a MarshalJSON() function for Pos as things stand, because
// we use Pos as an embedded field in our Unit and Factory structs, but such embedded fields promote
// their methods to the top level, so that MarshalJSON() would be called when marshalling either
// a Unit or a Factory.

type unit_tmp struct {
	TeamId					int								`json:"team_id"`
	UnitId					string							`json:"unit_id"`
	Power					int								`json:"power"`
	UnitType				string							`json:"unit_type"`
	Pos						[2]int							`json:"pos"`
	Cargo					Cargo							`json:"cargo"`
	ActionQueue				[][6]int						`json:"action_queue"`
}

func (u *Unit) UnmarshalJSON(data []byte) error {

	var v unit_tmp
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	u.TeamId = v.TeamId
	u.UnitId = v.UnitId
	u.Power = v.Power
	u.UnitType = v.UnitType
	u.Pos.X = v.Pos[0]
	u.Pos.Y = v.Pos[1]
	u.Cargo = v.Cargo
	
	for _, item := range v.ActionQueue {
		u.ActionQueue = append(u.ActionQueue, Action{
			Type:		ActionType(item[0]),
			Direction:	Direction(item[1]),
			Resource:	Resource(item[2]),
			Amount:		item[3],
			Recycle:	item[4],
			N:			item[5],
		})
	}

	return nil
}

type factory_tmp struct {
	TeamId					int								`json:"team_id"`
	UnitId					string							`json:"unit_id"`
	Power					int								`json:"power"`
	Pos						[2]int							`json:"pos"`
	Cargo					Cargo							`json:"cargo"`
	StrainId				int								`json:"strain_id"`
}

func (fc *Factory) UnmarshalJSON(data []byte) error {

	var v factory_tmp
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	fc.TeamId = v.TeamId
	fc.UnitId = v.UnitId
	fc.Power = v.Power
	fc.Pos.X = v.Pos[0]
	fc.Pos.Y = v.Pos[1]
	fc.Cargo = v.Cargo
	fc.StrainId = v.StrainId

	return nil
}
