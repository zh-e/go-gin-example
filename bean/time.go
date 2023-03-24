package bean

import "time"

type Time time.Time

const timeFormatter = "2023-03-24 00:00:00"

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	now, _ := time.ParseInLocation(`"`+timeFormatter+`"`, string(b), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormatter)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormatter)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormatter)
}
