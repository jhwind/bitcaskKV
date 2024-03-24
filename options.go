package bitcaskkv

import "os"

type Options struct {
	// 数据库数据目录
	DirPath string

	// 数据文件大小
	DataFileSize int64

	// 索引类型
	IndexType IndexerType

	// 每次写数据是否持久化
	SyncWrites bool

	// 累计写到多少字节后进行持久化
	BytesPerSync uint

	// 启动时是否使用 MMap 加载数据
	MMapAtStartup bool

	//	数据文件合并的阈值 无效数据上限百分比
	DataFileMergeRatio float32
}

// IteratorOptions 索引迭代器配置项
type IteratorOptions struct {
	// 遍历前缀为指定值的 Key，默认为空
	Prefix []byte
	// 是否反向遍历，默认 false 是正向
	Reverse bool
}

type WriteBatchOptions struct {
	// 一个批次中最大的数据量
	MaxBatchNum uint
	//  提交事务的时候是否进行 sync 持久化
	SyncWrites bool
}

type IndexerType = int8

const (
	// BTree 索引
	BTree IndexerType = iota + 1

	// ART
	ART

	// BPlusTree B+ 树索引，将索引存储到磁盘上
	BPlusTree
)

var DefaultOptions = Options{
	DirPath:       os.TempDir(),
	DataFileSize:  256 * 1024 * 1024, // 256MB
	SyncWrites:    false,
	BytesPerSync:  0,
	MMapAtStartup: true,
	// IndexType:    BPlusTree,
	// IndexType: ART,
	IndexType:          BTree,
	DataFileMergeRatio: 0.5,
}

var DefaultIteratorOptions = IteratorOptions{
	Prefix:  nil,
	Reverse: false,
}

var DefaultWriteBatchOptions = WriteBatchOptions{
	MaxBatchNum: 10000,
	SyncWrites:  true,
}
