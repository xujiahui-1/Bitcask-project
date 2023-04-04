package data

// LogRecordPos Memory index, describing the location of memory on disk(内存上的索引位置)
type LogRecordPos struct {
	Fid    uint32 //文件id
	Offset uint64 //偏移量
}
