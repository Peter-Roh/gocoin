package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ToBytes(i interface{}) []byte {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	HandleErr(enc.Encode(i))
	return buffer.Bytes()
}

func FromBytes(i interface{}, data []byte) {
	dec := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(dec.Decode(i))
}
