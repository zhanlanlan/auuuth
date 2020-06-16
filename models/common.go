package models

import (
	"strings"
	"unicode/utf8"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

const (
	minUnicodeRuneValue   = 0
	maxUnicodeRuneValue   = utf8.MaxRune
	compositeKeyNamespace = "\x00"
)

const (
	primaryPrefix = "pk"
	indexPrefix   = "idx"
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
func CreateCompositeKey(objectType string, attributes ...string) []byte {
	ck := compositeKeyNamespace + objectType + strings.Join(attributes, string(minUnicodeRuneValue))
	return []byte(ck)
}

// Put ...
func Put(key, value []byte) error {
	return db.Put(key, value, nil)
}

// Get ...
func Get(key []byte) ([]byte, error) {
	return db.Get(key, nil)
}

// WriteBatch ...
type WriteBatch struct {
	leveldb.Batch
}

// NewBatch ...
func NewBatch() WriteBatch {
	return WriteBatch{
		leveldb.Batch{},
	}
}

// Write ...
func (w *WriteBatch) Write() error {
	return db.Write(&w.Batch, nil)
}
