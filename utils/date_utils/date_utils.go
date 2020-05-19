package date_utils

import "time"

const (
	API_DATE_LAYOUT    = "2006-01-02T15:04:052"
	API_DB_DATE_LAYOUT = "2006-01-02 15:04:052"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(API_DATE_LAYOUT)
}

func GetNowDbFormat() string {
	return GetNow().Format(API_DB_DATE_LAYOUT)
}
