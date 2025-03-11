package station

type Station struct {
	Id   string `json:"nid"`
	Name string `json:"tittle"`
}

type StationResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
