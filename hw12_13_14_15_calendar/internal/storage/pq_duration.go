package storage

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type PqDuration time.Duration

func (d *PqDuration) Scan(src interface{}) error {
	srcStr, ok := src.([]uint8) // convert to string, then parse as time.Duration
	if !ok {
		return fmt.Errorf("duration column was not []uint8; type %T", src)
	}

	v := string(srcStr)
	v = strings.Replace(v, ":", "h", 1)
	v = strings.Replace(v, ":", "m", 1)
	v += "s"

	dur, err := time.ParseDuration(v)
	if err != nil {
		return err
	}
	*d = PqDuration(dur)

	return nil
}

func (d PqDuration) Value() (driver.Value, error) {
	return time.Duration(d).String(), nil
}
