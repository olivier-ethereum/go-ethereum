package state

import (
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Log interface {
	Address() common.Address
	Topics() []common.Hash
	Data() []byte

	Number() uint64
}

type StateLog struct {
	address common.Address
	topics  []common.Hash
	data    []byte
	number  uint64
}

func NewLog(address common.Address, topics []common.Hash, data []byte, number uint64) *StateLog {
	return &StateLog{address, topics, data, number}
}

func (self *StateLog) Address() common.Address {
	return self.address
}

func (self *StateLog) Topics() []common.Hash {
	return self.topics
}

func (self *StateLog) Data() []byte {
	return self.data
}

func (self *StateLog) Number() uint64 {
	return self.number
}

/*
func NewLogFromValue(decoder *common.Value) *StateLog {
	var extlog struct {

	}

	log := &StateLog{
		address: decoder.Get(0).Bytes(),
		data:    decoder.Get(2).Bytes(),
	}

	it := decoder.Get(1).NewIterator()
	for it.Next() {
		log.topics = append(log.topics, it.Value().Bytes())
	}

	return log
}
*/

func (self *StateLog) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{self.address, self.topics, self.data})
}

/*
func (self *StateLog) RlpData() interface{} {
	return []interface{}{self.address, common.ByteSliceToInterface(self.topics), self.data}
}
*/

func (self *StateLog) String() string {
	return fmt.Sprintf(`log: %x %x %x`, self.address, self.topics, self.data)
}

type Logs []Log

/*
func (self Logs) RlpData() interface{} {
	data := make([]interface{}, len(self))
	for i, log := range self {
		data[i] = log.RlpData()
	}

	return data
}
*/

func (self Logs) String() (ret string) {
	for _, log := range self {
		ret += fmt.Sprintf("%v", log)
	}

	return "[" + ret + "]"
}
