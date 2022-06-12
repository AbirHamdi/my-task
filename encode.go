package encode

import (
	"bytes"
	"fmt"
    "log"
	"github.com/HewlettPackard/structex"
)

type Nas5GSUpdateType struct {
	IEI             uint8                // Byte 0
	Length          uint8                // Byte 1
	SMS_requested   uint8 `bitfield:"1"` // Byte 2
	NG_RAN_RCU      uint8 `bitfield:"1"`
	fiveGS_PNB_CIoT uint8 `bitfield:"2"`
	EPS_PNB_CIoT    uint8 `bitfield:"2"`
	Spare           uint8 `bitfield:"2,reserved"`
}

func (ie Nas5GSUpdateType) Encode(buffer *bytes.Buffer) {
	if error := structex.Encode(buffer, ie); error != nil {
        log.Fatal(error)
	}

	iei := fmt.Sprintf("%#02x", buffer.Bytes()[0])
	length := fmt.Sprintf("%#02x", buffer.Bytes()[1])
	bitfield := fmt.Sprintf("%#02x", buffer.Bytes()[2])
	fmt.Println("ByteStream", iei, ",", length, ",", bitfield)
}