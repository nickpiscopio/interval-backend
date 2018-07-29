package structs

type Timer struct {
	Id         uint32 `json:"id,omitempty"`
	Timer      string `json:"timer,omitempty"`
	DateCreated int64  `json:"date_created,omitempty"`
	DateUpdated int64  `json:"date_updated,omitempty"`
	DateLastUsed int64  `json:"date_last_used,omitempty"`
}
