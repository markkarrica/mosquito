package bitarray

type BitArray struct {
	ByteArray []byte
	length    byte
}

func CreateBitArray(bitLength byte) BitArray {
	var byteLength byte
	if bitLength%8 == 0 {
		byteLength = bitLength / 8
	} else {
		byteLength = bitLength/8 + 1
	}
	return BitArray{make([]byte, byteLength), bitLength}
}

func (ba *BitArray) GetBit(globalBitIndex byte) (bool, error) {
	if ba.CheckBoundsError(globalBitIndex) {
		return false, OutOfBoundsError{ba.length, globalBitIndex}
	}
	byteIndex := globalBitIndex / 8
	bitIndex := globalBitIndex % 8
	return ba.ByteArray[byteIndex]&(1<<bitIndex) == (1 << bitIndex), nil
}

func (ba *BitArray) SetBit(globalBitIndex byte) error {
	if ba.CheckBoundsError(globalBitIndex) {
		return OutOfBoundsError{ba.length, globalBitIndex}
	}
	byteIndex := globalBitIndex / 8
	bitIndex := globalBitIndex % 8
	ba.ByteArray[byteIndex] |= (1 << bitIndex)
	return nil
}

func (ba *BitArray) ResetBit(globalBitIndex byte) error {
	if ba.CheckBoundsError(globalBitIndex) {
		return OutOfBoundsError{ba.length, globalBitIndex}
	}
	byteIndex := globalBitIndex / 8
	bitIndex := globalBitIndex % 8
	ba.ByteArray[byteIndex] &= ^(1 << bitIndex)
	return nil
}

func (ba *BitArray) SetBitValue(globalBitIndex byte, value bool) error {
	if value {
		return ba.SetBit(globalBitIndex)
	} else {
		return ba.ResetBit(globalBitIndex)
	}
}
