package main

import (
	"bytes"
	"encoding/binary"
	"flag"
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
	countPtr := flag.Int("count", 67, "Number of records to generate.")
	versionPtr := flag.Int("version", 206, "Version to write to header.")
	var directory string
	flag.StringVar(&directory, "directory", "", "Location to place file")
	flag.Parse()
	magicString := "MPS7"
	magicBytes := []byte(magicString)
	bufHeader(*countPtr, *versionPtr, magicBytes)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < int(*countPtr); i++ {
		generateRecord()
	}
	writeFile(directory)
}

func bufHeader(c int, v int, m []byte) {
	binary.Write(&buf, binary.BigEndian, m)
	binary.Write(&buf, binary.BigEndian, uint8(v))
	binary.Write(&buf, binary.BigEndian, uint32(c))
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

func writeFile(d string) {
	file, err0 := os.Create(d + "txnlog.dat")
	defer file.Close()
	if err0 != nil {
		log.Fatal(err0)
	}

	_, err1 := file.Write(buf.Bytes())
	if err1 != nil {
		log.Fatal(err1)
	}
}
