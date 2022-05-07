package model

type SceneRowsBO struct {
	Rows []*SceneBO `json:"rows"`
}

type SceneBO struct {
	Key   []string `json:"key"`
	Value float64  `json:"value"`
}

type SceneMetricsVO struct {
	Metrics []*SceneVO
}

type SceneVO struct {
	Id       int     `json:"id"`
	Location string  `json:"location"`
	Scores   float64 `json:"metrics"`
	Year     int     `json:"year,omitempty"`
}

type SceneRequest struct {
	Scene string `json:"topic"`
	Year  int    `json:"year"`
}
