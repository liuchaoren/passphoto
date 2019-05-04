package common

import (
	"math"
	"testing"
)

func TestReadConfig(t *testing.T) {
	cropConfig := ReadConfig("crop.json")
	headTop := cropConfig["head_top"].(float64)
	headBottom := cropConfig["head_bottom"].(float64)
	if math.Abs(headTop-0.091837) >= 1.0e-5 {
		t.Error("head_top read from config is wrong")
	}
	if math.Abs(headBottom-0.755) >= 1.0e-5 {
		t.Error("head_bottom read from config is wrong")
	}

}
