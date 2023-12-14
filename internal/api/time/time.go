package time

import (
	"encoding/binary"
	"github.com/gofiber/fiber/v2"
	"net"
	"time"
)

type packet struct {
	Settings       uint8
	Stratum        uint8
	Poll           int8
	Precision      int8
	RootDelay      uint32
	RootDispersion uint32
	ReferenceID    uint32
	RefTimeSec     uint32
	RefTimeFrac    uint32
	OrigTimeSec    uint32
	OrigTimeFrac   uint32
	RxTimeSec      uint32
	RxTimeFrac     uint32
	TxTimeSec      uint32
	TxTimeFrac     uint32
}

func getTimeFromServer(server string) (time.Time, error) {
	// Ask the NTP server for a bit of magick :D
	conn, err := net.Dial("udp", server+":123")
	if err != nil {
		// return empty time to return ANYTHING
		return time.Time{}, err
	}
	defer conn.Close()

	// we need to tell the server what we want to get back
	// 1B is 00011011
	// 00 - no leap second
	// Version Number - 011 = 3
	// Mode - Client Request (3)
	req := &packet{Settings: 0x1B}

	// send a ntp request to ntp server in most significant bit order
	binary.Write(conn, binary.BigEndian, req)

	// just tabula rasa
	rsp := &packet{}

	// receive answer from ntp server and store inside struct
	binary.Read(conn, binary.BigEndian, rsp)

	// from NTP to unix time
	// subtract the amount seconds from 1900 to 1970
	secs := float64(rsp.TxTimeSec) - 2208988800

	// convert fractions of a second to nanoseconds
	// shift 32 bits to the right = divide by 2^32
	nanos := (int64(rsp.TxTimeFrac) * 1e9) >> 32 // pointer magick :D

	// convert to unix time (seconds and nanoseconds since 01.01.1970)
	return time.Unix(int64(secs), nanos), nil
}

func HandleCurrent(c *fiber.Ctx) error {
	// get current time
	currentTime, err := getTimeFromServer("ptbtime1.ptb.de")
	// if something is wrong return a server error as HTTP Status Code
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	// return as json to use in pwsh
	return c.JSON(fiber.Map{
		"time": currentTime,
	})
}
