package store_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"treeverse-lake/index/store"

	"github.com/dgraph-io/badger/options"

	"github.com/dgraph-io/badger"
)

type nullLogger struct{}

func (l nullLogger) Errorf(string, ...interface{})   {}
func (l nullLogger) Warningf(string, ...interface{}) {}
func (l nullLogger) Infof(string, ...interface{})    {}
func (l nullLogger) Debugf(string, ...interface{})   {}

func GetIndexStore(t *testing.T) (store.Store, func()) {
	dir, err := ioutil.TempDir("", "treeverse-tests-badger")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("badger directory: %s\n", dir)
	opts := badger.DefaultOptions(dir)
	opts.Logger = nullLogger{}
	opts.TableLoadingMode = options.LoadToRAM
	kv, err := badger.Open(opts)
	if err != nil {
		t.Fatal(err)
	}
	return store.NewKVStore(kv), func() {
		err := os.RemoveAll(dir)
		if err != nil {
			t.Fatal(err)
		}
	}
}
