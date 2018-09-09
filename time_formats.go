package zadara

import (
	"time"
)

type ZadaraDate time.Time

func (j *ZadaraDate) UnmarshalJSON(b []byte) {
	println("sdsadas")
	t, _ := time.Parse("yyyy-MM-dd HH:mm:ss UTC", string(b))
	// if err != nil {
	// 	return err
	// }

	*j = ZadaraDate(t)
	return
}
