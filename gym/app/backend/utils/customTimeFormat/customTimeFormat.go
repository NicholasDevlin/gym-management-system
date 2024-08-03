package customtimeformat

import "time"

type CustomTime struct {
	time.Time
}

const customTimeFormat = "2006-01-02T15:04:05.999Z"

func (ct *CustomTime) UnmarshalText(text []byte) error {
	parsedTime, err := time.Parse(customTimeFormat, string(text))
	if err != nil {
		return err
	}
	*ct = CustomTime{parsedTime}
	return nil
}