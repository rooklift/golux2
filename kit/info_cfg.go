package kit

import (
	"math"
)

// This is all necessary because we do a lazy unmarshal of the env_cfg into a map[string]interface{}

func (self *Frame) CfgFloat(key string) float64 {
	if iface, ok1 := self.Info.EnvCfg[key]; ok1 {
		if as_float, ok2 := iface.(float64); ok2 {
			return as_float
		}
	}
	return 0
}

func (self *Frame) CfgInt(key string) int {
	return int(math.Round(self.CfgFloat(key)))
}

func (self *Frame) RobotCfgFloat(rtype int, key string) float64 {

	var robotkey string
	if rtype == LIGHT {
		robotkey = "LIGHT"
	} else if rtype == HEAVY {
		robotkey = "HEAVY"
	} else {
		panic("Supplied robot type was neither LIGHT nor HEAVY")
	}

	cfg_robots := self.Info.EnvCfg["ROBOTS"].(map[string]interface{})
	cfg_this_weight := cfg_robots[robotkey].(map[string]interface{})

	if iface, ok1 := cfg_this_weight[key]; ok1 {
		if as_float, ok2 := iface.(float64); ok2 {
			return as_float
		}
	}
	return 0
}

func (self *Frame) RobotCfgInt(rtype int, key string) int {
	return int(math.Round(self.RobotCfgFloat(rtype, key)))
}
