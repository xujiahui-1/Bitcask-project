package index

import "Bitcask-project/data"

// Indexer Different data structures are implemented by implementing this interface
type Indexer interface {
	//Put 向索引中存储key对应的信息
	Put(key []byte, pos *data.LogRecordPos) bool

	//Get 根据key获取btree中的索引信息
	Get(key []byte) *data.LogRecordPos

	//Delete 根据key删除索引
	Delete(key []byte) bool
}
