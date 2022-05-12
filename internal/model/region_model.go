package model

type RegionRowsDO struct {
	Rows []*RegionDO `json:"rows"`
}

type RegionDO struct {
	Key   []string `json:"key"`
	Value Value    `json:"value"`
}

type Value struct {
	Sum    float64 `json:"sum"`
	Count  float64 `json:"count"`
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
	SumSQR float64 `json:"sum_sqr"`
}

type SportsVO struct {
	BarsNumber  float64 `json:"barsNumber"`
	CafesNUmber float64 `json:"cafesNUmber""`
	Ratio       float64 `json:"ratio"`
	LocationPid float64 `json:"locationPid"`
}

type FoodsVO struct {
	BarsScore   float64 `json:"barsScore"`
	CafesScore  float64 `json:"cafesScore""`
	Ratio       float64 `json:"ratio"`
	LocationPid string  `json:"locationPid"`
}
