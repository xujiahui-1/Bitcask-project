package index

import (
	"Bitcask-project/data"
	"bytes"
	"github.com/google/btree"
	"sync"
)

// BTree 实现btree数据结构来存储索引(基于内存的)
// BTree google btree github.com/google/btree
type BTree struct {
	tree *btree.BTree
	lock *sync.RWMutex
}

// NewBTree  初始化BTree 索引结构
func NewBTree() *BTree {

	//TODO 这里的32表示叶子节点的数量，我们可以提供让用户选择次参数，
	return &BTree{tree: btree.New(32), lock: new(sync.RWMutex)}
}

// Item 实现Item interface  定义自己的key 对比结构
type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}

func (bt *BTree) Put(key []byte, pos *data.LogRecordPos) bool {
	item := &Item{key: key, pos: pos}
	//存数据前对其进行加锁，因为谷歌的btree同步写是不安全的
	bt.lock.Lock()
	bt.tree.ReplaceOrInsert(item)
	bt.lock.Unlock()
	return true
}
func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	item := &Item{key: key}
	btreeItem := bt.tree.Get(item)
	if btreeItem == nil {
		return nil
	}
	return btreeItem.(*Item).pos
}
func (bt *BTree) Delete(key []byte) bool {
	if len(key) < 0 {
		return false
	}
	itme := &Item{key: key}
	bt.lock.Lock()
	oldItem := bt.tree.Delete(itme)
	bt.lock.Unlock()
	if oldItem == nil {
		return false
	}
	return true
}
