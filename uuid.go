package UUID

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
	"sync/atomic"
	"time"
)

func GetUUID() string {
	seed := time.Now().UTC().UnixNano()

	var data = make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(seed))

	timeLow := getTimeLow(data)
	timeMid := getTimeMiddle(seed)
	timeHiNVersion := getTimeHiNVersion(seed)
	clockSeqPart := getClockSeq(seed)
	node := getNode()

	return fmt.Sprintf("%s-%s-%s-%s-%s", timeLow, timeMid, timeHiNVersion, clockSeqPart, node)
}

func getTimeLow(data []byte) string {
	return hex.EncodeToString(data[:4])
}

func getTimeMiddle(s int64) string {
	data := []byte{
		byte(s >> 8),
		byte(s >> 16),
	}
	return hex.EncodeToString(data)
}

func getTimeHiNVersion(s int64) string {
	data := []byte{
		byte(s & 0x23),
		byte(s >> 12),
	}
	return hex.EncodeToString(data)
}

func getClockSeq(s int64) string {
	var clockSeq uint32
	clock := atomic.AddUint32(&clockSeq, uint32(s))
	data := []byte{
		byte(clock >> 8),
		byte(clock),
	}
	return hex.EncodeToString(data)
}
func getNode() string {
	mac := getMacAddr()
	return strings.Replace(mac, ":", "", -1)
}

// getMacAddr gets the MAC hardware
// address of the host machine
func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}
