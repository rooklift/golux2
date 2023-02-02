package kit

// Thanks to MMJ -- aka themmj on Github -- for help with this...

import (
	"encoding/json"
)

// Our Factory and Unit types have some embedded structs for user convenience, but this seems to
// mean (as far as I can tell) that we need custom unmarshalling for them... (?)
//
// Be careful about implementing MarshalJSON() for inner types like Pos, this can cause issues
// if you're not careful: https://github.com/golang/go/issues/39175

type unit_tmp struct {
	TeamId					int								`json:"team_id"`
	UnitId					string							`json:"unit_id"`
	Power					int								`json:"power"`
	UnitType				string							`json:"unit_type"`
	Pos						[2]int							`json:"pos"`
	Cargo					Cargo							`json:"cargo"`
	ActionQueue				[][6]int						`json:"action_queue"`
}

func (u Unit) MarshalJSON() ([]byte, error) {

	var tmp unit_tmp

	tmp.TeamId = u.TeamId
	tmp.UnitId = u.UnitId
	tmp.Power = u.Power
	tmp.UnitType = u.UnitType
	tmp.Pos = [2]int{u.X, u.Y}
	tmp.Cargo = u.Cargo

	for _, item := range u.ActionQueue {
		tmp.ActionQueue = append(tmp.ActionQueue, [6]int{
			int(item.Type),
			int(item.Direction),
			int(item.Resource),
			item.Amount,
			item.Recycle,
			item.N,
		})
	}

	return json.Marshal(tmp)
}

func (u *Unit) UnmarshalJSON(data []byte) error {

	var tmp unit_tmp
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	u.TeamId = tmp.TeamId
	u.UnitId = tmp.UnitId
	u.Power = tmp.Power
	u.UnitType = tmp.UnitType
	u.Pos.X = tmp.Pos[0]
	u.Pos.Y = tmp.Pos[1]
	u.Cargo = tmp.Cargo

	for _, item := range tmp.ActionQueue {
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

func (fc Factory) MarshalJSON() ([]byte, error) {

	var tmp factory_tmp

	tmp.TeamId = fc.TeamId
	tmp.UnitId = fc.UnitId
	tmp.Power = fc.Power
	tmp.Pos = [2]int{fc.X, fc.Y}
	tmp.Cargo = fc.Cargo
	tmp.StrainId = fc.StrainId

	return json.Marshal(tmp)
}

func (fc *Factory) UnmarshalJSON(data []byte) error {

	var tmp factory_tmp
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	fc.TeamId = tmp.TeamId
	fc.UnitId = tmp.UnitId
	fc.Power = tmp.Power
	fc.Pos.X = tmp.Pos[0]
	fc.Pos.Y = tmp.Pos[1]
	fc.Cargo = tmp.Cargo
	fc.StrainId = tmp.StrainId

	return nil
}
