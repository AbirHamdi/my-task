package encode

import (
	"bytes"
	"fmt"
	"testing"
)
func TestEncode(t *testing.T) {

    var buffer bytes.Buffer
	// Define the input structure
	m := Nas5GSUpdateType{1, 2, 1, 1, 0, 0, 0}
	// Print the defined structure with Variable Name
	fmt.Println("Nas5GSUpdateType ")
	fmt.Printf("%+v\n", m)
	//Encode
	m.Encode(&buffer)

	//the byte stream we expect 
	expected := []string{"0x01", "0x02", "0x03"}
	//the byte stream we accually got after encoding the input struct 
	actual := []string{fmt.Sprintf("%#02x", buffer.Bytes()[0]), fmt.Sprintf("%#02x", buffer.Bytes()[1]), fmt.Sprintf("%#02x", buffer.Bytes()[2])}
	
	for i := 0; i < 3; i++ {
		if expected[i] != actual[i] {
			t.Errorf("Invalid array pack: Index: %d Expected: %#02x Actual: %#02x", i, expected[i], actual[i])
		}
	}

}