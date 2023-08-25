package domain

type Project struct {
	ID          int64
	Label       string
	Description string
	Status      string
	HoursSpent  int64 `json:"hoursSpent"`
	Subprojects []int64
}
