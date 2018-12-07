package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type ReadRecordLong struct {
	enum  byte
	time  uint32
	user  uint64
	money float64
}

type ReadRecordShort struct {
	enum byte
	time uint32
	user uint64
}

var rsLong []*ReadRecordLong
var rsShort []*ReadRecordShort
var directory []string
var totCreditAgg float64
var totDebitAgg float64
var apStartAgg int
var apEndAgg int

func main() {
	totCreditAgg = 0
	totDebitAgg = 0
	apStartAgg = 0
	apEndAgg = 0
	totCPtr := flag.Bool("totcredit", true, "Total all credit transactions.")
	totDPtr := flag.Bool("totdebit", true, "Total of all debit transactions.")
	totAPtr := flag.Bool("totall", false, "Total all monetary transactions.")
	apStartPtr := flag.Bool("autopaystart", true, "Total number of autopays started.")
	apEndPtr := flag.Bool("autopayend", true, "Total number of autopays ended.")
	var file string
	flag.StringVar(&file, "file", "txnlog.dat", "File to read.")
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	createData(file)
	if *totCPtr {
		totCredit()
		fmt.Println("Total credit transactions: " + strconv.FormatFloat(totCreditAgg, 'f', -1, 64))
	}
	if *totDPtr {
		totDebit()
		fmt.Println("Total debit transactions: " + strconv.FormatFloat(totDebitAgg, 'f', -1, 64))
	}
	if *totAPtr {
		totAll()
		fmt.Println("All monetary transactions: " + strconv.FormatFloat((totCreditAgg+totDebitAgg), 'f', -1, 64))
	}
	if *apStartPtr {
		autoPayStart()
		fmt.Println("Total number of autopays started: " + strconv.Itoa(apStartAgg))
	}
	if *apEndPtr {
		autoPayEnd()
		fmt.Println("Total number of autopays ended: " + strconv.Itoa(apEndAgg))
	}
}

func totCredit() {
	for i := 0; i <= len(rsLong)-1; i++ {
		if rsLong[i].enum == 1 {
			totCreditAgg += rsLong[i].money
		}
	}
}
func totDebit() {
	for i := 0; i <= len(rsLong)-1; i++ {
		if rsLong[i].enum == 0 {
			totDebitAgg += rsLong[i].money
		}
	}
}
func totAll() {
	if (totCreditAgg == 0) && (totDebitAgg == 0) {
		totCredit()
		totDebit()
	}
}
func autoPayStart() {
	for i := 0; i <= len(rsShort)-1; i++ {
		if rsShort[i].enum == 2 {
			apStartAgg++
		}
	}
}
func autoPayEnd() {
	for i := 0; i <= len(rsShort)-1; i++ {
		if rsShort[1].enum == 3 {
			apEndAgg++
		}
	}
}

func createData(f string) {
	file, err := os.Open(f)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	readBytes(*file, 4)
	readBytes(*file, 1)
	count := readBytes(*file, 4)
	countUint := binary.BigEndian.Uint32(count)
	makeRecordsStructs(countUint, *file)
}

func makeRecordsStructs(c uint32, f os.File) {
	for i := 1; i <= int(c); i++ {
		enum := readBytes(f, 1)
		if (enum[0] == 0) || (enum[0] == 1) {
			recordLong := new(ReadRecordLong)
			recordLong.enum = enum[0]
			recordLong.time = binary.BigEndian.Uint32(readBytes(f, 4))
			recordLong.user = binary.BigEndian.Uint64(readBytes(f, 8))
			bits := binary.BigEndian.Uint64(readBytes(f, 8))
			recordLong.money = math.Float64frombits(bits)
			rsLong = append(rsLong, recordLong)
		} else {
			recordShort := new(ReadRecordShort)
			recordShort.enum = enum[0]
			recordShort.time = binary.BigEndian.Uint32(readBytes(f, 4))
			recordShort.user = binary.BigEndian.Uint64(readBytes(f, 8))
			rsShort = append(rsShort, recordShort)
		}
	}
}

func dollarTotals(d float64, c float64) {

}

func readBytes(file os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
