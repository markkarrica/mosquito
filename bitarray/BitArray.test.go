package bitarray

import (
	"testing"
)

func TestByteArrayLengthMultiple8(t *testing.T) {
	myByteArray := CreateBitArray(64)
	if len(myByteArray.ByteArray) != 8 {
		t.Errorf("Bit array of 64 should create 8 bytes (expected: 8, actual: %d)", len(myByteArray.ByteArray))
	}
}
func TestByteArrayLengthNotMultiple8(t *testing.T) {
	myByteArray := CreateBitArray(65)
	if len(myByteArray.ByteArray) != 9 {
		t.Errorf("Bit array of 65 should create 9 bytes (expected: 9, actual: %d)", len(myByteArray.ByteArray))
	}
}
func TestSetBitValue(t *testing.T) {
	myByteArray := CreateBitArray(8)
	myByteArray.SetBitValue(3, true)
	if myByteArray.ByteArray[0] != 8 {
		t.Errorf("Expected first byte to be 8 (expected: 8, actual: %d)", myByteArray.ByteArray[0])
	}
}
func TestSetLaterBitValue(t *testing.T) {
	myByteArray := CreateBitArray(16)
	myByteArray.SetBitValue(11, true)
	if myByteArray.ByteArray[1] != 8 {
		t.Errorf("Expected second byte to be 8 (expected: 8, actual: %d)", myByteArray.ByteArray[1])
	}
}
func TestResetBit(t *testing.T){
	myByteArray := CreateBitArray(8)
	for i := byte(0); i < 8; i++ {
		myByteArray.SetBit(i)
	}
	myByteArray.ResetBit(3)
	if myByteArray.ByteArray[0] != 247 {
		t.Errorf("Expected first byte to be 239 (expected: 239, actual: %d)", myByteArray.ByteArray[0])
	}
}