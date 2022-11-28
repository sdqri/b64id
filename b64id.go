package b64id

import (
	"encoding/base64"
	"math"
	"math/rand"
	"time"
)

type IDLength int

const (
	IDL2  IDLength = 2
	IDL3  IDLength = 3
	IDL4  IDLength = 4
	IDL8  IDLength = 8
	IDL10 IDLength = 10
	IDL11 IDLength = 11
	IDL12 IDLength = 12
	IDL14 IDLength = 14
	IDL15 IDLength = 15
	IDL16 IDLength = 16
)

func GenerateB64ID(l IDLength) string {
	// Seeding by time
	rand.Seed(time.Now().UnixNano())

	// length of output for b64 string without padding = ceil(n × 4 / 3)  where n = length of bytes
	// to generate an out put with length = l we should find n
	// Finding n
	n := int(math.Floor((float64(l) * 3.0) / 4.0))

	b := make([]byte, n)

	// Read size number of bytes into b
	rand.Read(b)

	id := base64.RawURLEncoding.EncodeToString(b)

	return id
}
