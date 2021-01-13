package main

import "fmt"
import "encoding/hex"
import "./bn256"

func main() {
	Export()
}

func Export() {

	var p bn256.G1
	var input = "a01f9bcc1208dee302769931ad378a4c0c4b2c21b0cfb3e752607e12d2b6fa6425"
	var input_uncompressed = "001f9bcc1208dee302769931ad378a4c0c4b2c21b0cfb3e752607e12d2b6fa642510483ad6191b3bdfd1ba5f610a24b56ab7ec72c38e00bb84953ab3481b98e74e"

	data, err := hex.DecodeString(input_uncompressed)
	if err != nil {
		panic(err)
	}

	if err := p.DecodeUncompressed(data); err != nil {
		panic(fmt.Sprintf("Decode point err %s err", err))
	}

	// now compressing the point again will result in corruption
	compressed := p.EncodeCompressed()

	fmt.Printf("Actual   %x\n", compressed)
	fmt.Printf("Expected %s\n", input)

}
