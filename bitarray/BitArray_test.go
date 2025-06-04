package bitarray

import (
	"testing"
)

func TestByteArrayLengthMultiple8(t *testing.T) {
	myByteArray := CreateBitArray(64)
	if len(myByteArray.byteArray) != 8 {
		t.Errorf("Bit array of 64 should create 8 bytes (expected: 8, actual: %d)", len(myByteArray.byteArray))
	}
}
func TestByteArrayLengthNotMultiple8(t *testing.T) {
	myByteArray := CreateBitArray(65)
	if len(myByteArray.byteArray) != 9 {
		t.Errorf("Bit array of 65 should create 9 bytes (expected: 9, actual: %d)", len(myByteArray.byteArray))
	}
}
func TestSetBitValue(t *testing.T) {
	myByteArray := CreateBitArray(8)
	myByteArray.SetBitValue(3, true)
	if myByteArray.byteArray[0] != 8 {
		t.Errorf("Expected first byte to be 8 (expected: 8, actual: %d)", myByteArray.byteArray[0])
	}
}
func TestSetLaterBitValue(t *testing.T) {
	myByteArray := CreateBitArray(16)
	myByteArray.SetBitValue(11, true)
	if myByteArray.byteArray[1] != 8 {
		t.Errorf("Expected second byte to be 8 (expected: 8, actual: %d)", myByteArray.byteArray[1])
	}
}
func TestResetBit(t *testing.T){
	myByteArray := CreateBitArray(8)
	for i := byte(0); i < 8; i++ {
		myByteArray.SetBit(i)
	}
	myByteArray.ResetBit(3)
	if myByteArray.byteArray[0] != 247 {
		t.Errorf("Expected first byte to be 239 (expected: 239, actual: %d)", myByteArray.byteArray[0])
	}
}