package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

var input = `
{
	"created_at" : "Thu May 31 00:00:01 +0000 2012"
}
`

type TimeStamp time.Time

func (t *TimeStamp) String() string {
	return time.Time(*t).String()
}

func (t *TimeStamp) UnmarshalJSON(bytes []byte) error {
	v, err := time.Parse(time.RubyDate, string(bytes[1:len(bytes)-1]))
	if err != nil {
		return err
	}

	*t = TimeStamp(v)
	return nil
}

func main() {
	var val map[string]TimeStamp

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Printf("k: %v, v: %v, type: %v", k, v.String(), reflect.TypeOf(v))
	}
}
