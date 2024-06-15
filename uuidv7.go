package uuidv7

import (
	"crypto/rand"
	"time"
)

func UUIDV7() ([16]byte, error) {
	// random bytes
	var value [16]byte
	_, err := rand.Read(value[5:])
	if err != nil {
		return value, err
	}

	// current timestamp in ms
	timestamp := uint64(time.Now().UnixNano() / int64(time.Millisecond))

	// timestamp
	value[0] = byte((timestamp >> 40) & 0xFF)
	value[1] = byte((timestamp >> 32) & 0xFF)
	value[2] = byte((timestamp >> 24) & 0xFF)
	value[3] = byte((timestamp >> 16) & 0xFF)
	value[4] = byte((timestamp >> 8) & 0xFF)
	value[5] = byte(timestamp & 0xFF)

	// version and variant
	value[6] = (value[6] & 0x0F) | 0x70
	value[8] = (value[8] & 0x3F) | 0x80

	return value, nil
}
