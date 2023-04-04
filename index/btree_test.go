package index

import (
	"Bitcask-project/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试需要断言，所以 github.com/stretchr/testify/assert 这个的断言等功能
func TestBTree_Put(t *testing.T) {
	bt := NewBTree()
	res := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res)

	res2 := bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 2})
	assert.True(t, res2)
}

func TestBTree_Get(t *testing.T) {
	bt := NewBTree()
	res := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res)
	pos1 := bt.Get(nil)
	assert.Equal(t, uint32(1), pos1.Fid)

	res2 := bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 2})
	assert.True(t, res2)
	pos2 := bt.Get([]byte("a"))
	assert.Equal(t, uint64(2), pos2.Offset)
}

func TestBTree_Delete(t *testing.T) {

}
