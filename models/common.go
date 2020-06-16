package models

import (
	"bytes"
	"unicode/utf8"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

var (
	minUnicodeRuneValue   = []byte(string(0))
	maxUnicodeRuneValue   = utf8.MaxRune
	compositeKeyNamespace = []byte("\x00")
)

var (
	primaryPrefix = []byte("pk")
	indexPrefix   = []byte("idx")
)

func init() {
	initLevelDB()
}

func initLevelDB() {
	// todo should set some necessary options
	// and should provided by config file
	db0, err := leveldb.OpenFile("testdb", nil)
	if err != nil {
		panic(err)
	} else if db != nil {
		panic("the db instance already inited")
	} else {
		db = db0
	}
}

// CreateCompositeKey ...
func CreateCompositeKey(objectType []byte, attributes ...[]byte) []byte {

	ckByte := [][]byte{compositeKeyNamespace, objectType}
	ckByte = append(ckByte, attributes...)

	ck := bytes.Join(ckByte, minUnicodeRuneValue)

	return ck
}

// Put ...
func Put(key, value []byte) error {
	return db.Put(key, value, nil)
}

// Get ...
func Get(key []byte) ([]byte, error) {
	return db.Get(key, nil)
}

// Record 对应于一条记录
type Record struct {
	pk     []byte
	indexs [][]byte

	data Marshable
}

// Marshable ...
type Marshable interface {
	Marshal() ([]byte, error)
}

// NewRecord ...NewRecordWithData
func NewRecord(pk []byte, data Marshable) Record {
	return Record{
		pk:   pk,
		data: data,
	}
}

// AddIndex ...
func (r *Record) AddIndex(idx []byte) {
	r.indexs = append(r.indexs, idx)
}

// Commit ...
func (r *Record) Commit() error {
	data, err := Get(r.pk)
	if err != nil {
		return err
	} else if data != nil {
		// data already exist
		// it's fine just cover
		// 覆盖掉就好了，数据已经存在也没关系
	}

	var batch leveldb.Batch

	data, err = r.data.Marshal()
	if err != nil {
		return err
	}

	// 主键指向具体的数据
	batch.Put(r.pk, data)

	// 索引指向主键
	for _, idxKey := range r.indexs {
		batch.Put(idxKey, r.pk)
	}

	return db.Write(&batch, nil)
}
