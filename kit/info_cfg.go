package kit

import (
	"fmt"
	"math"
)

func CfgInt(key string) (int, error) {

	val, ok := msg.Info.EnvCfg[key]
	if ok == false {
		return 0, fmt.Errorf("Key \"%s\" not in env_cfg dict")
	}

	switch n := val.(type) {
	case float64:
		return int(math.Round(n)), nil
	default:
		return 0, fmt.Errorf("Value was not a number")
	}
}

func CfgFloat(key string) (float64, error) {

	val, ok := msg.Info.EnvCfg[key]
	if ok == false {
		return 0, fmt.Errorf("Key \"%s\" not in env_cfg dict")
	}

	switch n := val.(type) {
	case float64:
		return n, nil
	default:
		return 0, fmt.Errorf("Value was not a number")
	}
}
