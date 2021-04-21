package keycache

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/json"

	// ."github.com/golang/mock"
)

type Data struct {
	MyLongData1 string
	MyLongData2 string
}

type MyTestItem struct {
	id   int
	md5  []byte
	Data *Data
	Find bool
}

func (it *MyTestItem) ID() int {
	return it.id
}

func GetMD5(it interface{}) []byte {
	js, err := json.Marshal(it)
	if err != nil {
		return nil
	}

	s := md5.Sum(js)
	return s[:]
}

func (it *MyTestItem) Sign() []byte {

	if it.md5 == nil || len(it.md5) == 0 {
		it.md5 = GetMD5(it.Data)
	}

	return it.md5
}

func (it *MyTestItem) Decode(data []byte) error {
	it.Data = &Data{}
	it.Find = true
	return gob.NewDecoder(bytes.NewBuffer(data)).Decode(it.Data)
}

func (it *MyTestItem) Encode() ([]byte, error) {
	network := &bytes.Buffer{}
	if err := gob.NewEncoder(network).Encode(it.Data); err != nil {
		return nil, err
	}

	return network.Bytes(), nil
}
