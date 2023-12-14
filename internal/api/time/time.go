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
	conn, err := net.Dial("udp", server+":123")
	if err != nil {
		return time.Time{}, err
	}
	defer conn.Close()

	req := &packet{Settings: 0x1B}
	binary.Write(conn, binary.BigEndian, req)

	rsp := &packet{}
	binary.Read(conn, binary.BigEndian, rsp)

	secs := float64(rsp.TxTimeSec) - 2208988800
	nanos := (int64(rsp.TxTimeFrac) * 1e9) >> 32

	return time.Unix(int64(secs), nanos), nil
}

func HandleCurrent(c *fiber.Ctx) error {
	currentTime, err := getTimeFromServer("ptbtime1.ptb.de")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"time": currentTime,
	})
}
