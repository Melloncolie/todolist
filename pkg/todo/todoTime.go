package todo

import (
	"encoding/json"
	"time"
)

func (times *TodoTime) String() string {
	return times.Time.Format("2006-01-02 15:04:05")
}

func (times *TodoTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(times.String())
}

func (times *TodoTime) UnmarshalJSON(data []byte) (err error) {
	var (
		str       *string
		parseTime time.Time
	)
	err = json.Unmarshal(data, &str)
	if err != nil {
		return
	}

	parseTime, err = time.Parse("2006-01-02 15:04:05", *str)
	if err != nil {
		return
	}

	times.Time = &parseTime
	return
}
