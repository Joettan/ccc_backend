package model

type SportsSceneBO struct {
	Rows []*SportsBO `json:"rows"`
}

type SportsBO struct {
	Key   []string `json:"key"`
	Value float64  `json:"value"`
}

type SportsSceneVO struct {
	Metrics []*SportVO
}

type SportVO struct {
	Id       int     `json:"id"`
	Location string  `json:"location"`
	Scores   float64 `json:"metrics"`
	Year     int     `json:"year"`
}
