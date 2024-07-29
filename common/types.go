package common

import "time"

type Duration struct {
	time.Duration
}

func NewDuration(text string) *Duration {
	var d Duration
	d.UnmarshalText([]byte(text))
	return &d
}

func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func (d Duration) MarshalText() ([]byte, error) {
	return []byte(d.Duration.String()), nil
}

func (d *Duration) Value() time.Duration {
	if d == nil {
		return time.Duration(0)
	}
	return d.Duration
}

func (d *Duration) Seconds() int {
	return int(d.Value() / time.Second)
}
