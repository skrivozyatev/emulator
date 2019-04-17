package emulator

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

type carrier struct {
	rnd    *rand.Rand
	info   *log.Logger
	sender *sender
}

func CreateCarrier() *carrier {
	c := new(carrier)
	c.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	c.info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	c.sender = CreateSender("172.16.100.54:65000")
	return c
}

func (c *carrier) Put(inputId int, parcel *parcel, wg *sync.WaitGroup) {
	defer wg.Done()
	c.info.Printf("Parcel %d start", parcel.id)
	c.sleep(1)
	c.info.Printf("Parcel %d request 1", parcel.id)
	c.sendRequest1(inputId, parcel)
	c.sleep(2)
	c.info.Printf("Parcel %d request 2", parcel.id)
	c.sendRequest2(inputId, parcel)
	c.sleep(1)
	c.info.Printf("Parcel %d request 3", parcel.id)
	c.sendRequest3(inputId, parcel)
	c.sleep(2)
	c.info.Printf("Parcel %d report", parcel.id)
	c.sendReport(parcel)
}

func (c *carrier) sendRequest1(inputId int, parcel *parcel) {
	c.sendRequest([]string{
		"IP;" + iif(inputId <= 3, "11", "12"),
		fmt.Sprintf("%d", parcel.id),
		fmt.Sprintf("%d", parcel.Length),
		fmt.Sprintf("%d", parcel.Height),
		fmt.Sprintf("%d", parcel.Width),
		fmt.Sprintf("%d", parcel.Weight),
		string(parcel.BarCodes)})
}

func (c *carrier) sendRequest2(inputId int, parcel *parcel) {
	c.sendRequest([]string{
		"IP;" + iif(inputId <= 3, "13", "14"),
		fmt.Sprintf("%d", parcel.id),
		string(parcel.BarCodes)})
}

func (c *carrier) sendRequest3(inputId int, parcel *parcel) {
	c.sendRequest([]string{
		"IP;" + iif(inputId <= 3, "15", "16"),
		fmt.Sprintf("%d", parcel.id),
		"1",
		string(parcel.BarCodes)})
}

func (c *carrier) sendReport(parcel *parcel) {
	c.sendRequest([]string{
		"RP;3", fmt.Sprintf("%d", parcel.id),
		fmt.Sprintf("%d", 1002+c.rnd.Intn(153)), "1"})
}

func (c *carrier) sendRequest(strings []string) {
	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)
	writer.Write(strings)
	writer.Flush()
	c.sender.Send(buf.Bytes()[:buf.Len()-1]) // приходится убирать последний байт - перевод строки
}

func iif(cond bool, trueStr string, falseString string) string {
	if cond {
		return trueStr
	} else {
		return falseString
	}
}

func (c *carrier) sleep(seconds int) {
	time.Sleep(time.Duration(seconds)*time.Second + time.Duration(c.rnd.Intn(1000))*time.Millisecond)
}
