package dbutils

type TrainResource struct {
	ID              int    `json:"id"`
	DriverName      string `json:"drivers_name"`
	OperatingStatus bool   `json:"operating_Status"`
}

type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

type Schedule struct {
	ID          int    `json:"id"`
	TrainID     int    `json:"train_id"`
	StationID   int    `json:"station_id"`
	ArrivalTime string `json:"arrival_time"`
}
