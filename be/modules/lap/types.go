package lap

type GetLapsFilter struct {
	Date  []string `form:"date"`
	Game  []string `form:"game"`
	Car   []string `form:"car"`
	Track []string `form:"track"`
	Time  []string `form:"time"`
	Clear []string `form:"clear"`
}

type GetListLap struct {
	ID    int    `json:"id"`
	Date  string `json:"date"`
	Game  string `json:"game"`
	Car   string `json:"car"`
	Track string `json:"track"`
	Time  string `json:"time"`
	Clear bool   `json:"clear"`
}

type GetLapsResponse struct {
	Laps       []GetListLap `json:"laps"`
	TotalCount int          `json:"total_count"`
}
