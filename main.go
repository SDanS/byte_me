package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type RecordLong struct {
	enum  uint8
	time  uint32
	user  uint64
	money float64
}

type RecordShort struct {
	enum uint8
	time uint32
	user uint64
}

var buf bytes.Buffer

func main() {
	magicString := "MPS7"
	magicBytes := []byte(magicString)
	rand.Seed(time.Now().UnixNano())
	binary.Write(&buf, binary.BigEndian, magicBytes)
	binary.Write(&buf, binary.BigEndian, uint8(255))
	binary.Write(&buf, binary.BigEndian, uint32(67))
	fmt.Println(buf.Bytes())
	for i := 0; i < int(67); i++ {
		generateRecord()
	}
	fmt.Println(buf.Bytes())
	writeFile()
}

func generateRecord() {
	enum := randomEnum(0, 4)
	if (enum == 0) || (enum == 1) {
		bufwriteLong(enum)
	} else {
		bufWriteShort(enum)
	}
}

func bufwriteLong(enum int) {
	binary.Write(&buf, binary.BigEndian, &RecordLong{uint8(enum), uint32(time.Now().Unix()), rand.Uint64(), randomFloat64(.01, 10000)})
}

func bufWriteShort(enum int) {
	binary.Write(&buf, binary.BigEndian, &RecordShort{uint8(enum), uint32(time.Now().Unix()), rand.Uint64()})
}
func randomEnum(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func writeFile() {
	file, err0 := os.Create("txnlog.dat")
	defer file.Close()
	if err0 != nil {
		log.Fatal(err0)
	}

	_, err1 := file.Write(buf.Bytes())
	if err1 != nil {
		log.Fatal(err1)
	}
}
