package structs

type Timer struct {
	Id         uint32 `json:"id,omitempty"`
	Timer      string `json:"timer"`
	StoredDate int64  `json:"date,omitempty"`
}
