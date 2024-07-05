package sharder

import (
	"encoding/binary"
	"errors"
	"fmt"
	"sort"
	"sync"

	"go.uber.org/zap"
	"golang.org/x/crypto/blake2b"

	"github.com/syth0le/dialog-service/internal/storage"
)

const replicationFactor = 10

var ErrNoHosts = errors.New("no hosts added")

type Host struct {
	Name string
	Load int64
}

type ConsistentSharder struct {
	shards    map[uint64]storage.Storage
	sortedSet []uint64

	logger *zap.Logger

	sync.RWMutex
}

func NewConsistentSharder(logger *zap.Logger) *ConsistentSharder {
	return &ConsistentSharder{
		shards:    map[uint64]storage.Storage{},
		sortedSet: []uint64{},
		logger:    logger,
	}
}

func (c *ConsistentSharder) Add(storage storage.Storage) {
	c.Lock()
	defer c.Unlock()

	for i := 0; i < replicationFactor; i++ {
		h := c.hash(fmt.Sprintf("%s%d", storage.Salt(), i))
		c.shards[h] = storage
		c.sortedSet = append(c.sortedSet, h)
	}
	// sort hashes ascendingly
	sort.Slice(c.sortedSet, func(i int, j int) bool {
		if c.sortedSet[i] < c.sortedSet[j] {
			return true
		}
		return false
	})
}

// Get Returns the host that owns `key`.
// As described in https://en.wikipedia.org/wiki/Consistent_hashing
// It returns ErrNoHosts if the ring has no hosts in it.
func (c *ConsistentSharder) Get(key string) (storage.Storage, error) {
	c.RLock()
	defer c.RUnlock()

	if len(c.shards) == 0 {
		return nil, ErrNoHosts
	}

	h := c.hash(key)
	idx := c.search(h)

	shard := c.shards[c.sortedSet[idx]]
	c.logger.Sugar().Infof("get shard: %s", shard.Salt())

	return shard, nil
}

// Remove Deletes host from the ring
func (c *ConsistentSharder) Remove(shard storage.Storage) bool {
	c.Lock()
	defer c.Unlock()

	for i := 0; i < replicationFactor; i++ {
		h := c.hash(fmt.Sprintf("%s%d", shard.Salt(), i))
		delete(c.shards, h)
		c.delSlice(h)
	}

	return true
}

// Hosts Return the list of hosts in the ring
func (c *ConsistentSharder) Hosts() (hosts []string) {
	c.RLock()
	defer c.RUnlock()

	for _, v := range c.shards {
		hosts = append(hosts, v.Hosts()...)
	}

	return hosts
}

func (c *ConsistentSharder) search(key uint64) int {
	idx := sort.Search(len(c.sortedSet), func(i int) bool {
		return c.sortedSet[i] >= key
	})

	if idx >= len(c.sortedSet) {
		idx = 0
	}
	return idx
}

func (c *ConsistentSharder) delSlice(val uint64) {
	idx := -1
	l := 0
	r := len(c.sortedSet) - 1
	for l <= r {
		m := (l + r) / 2
		if c.sortedSet[m] == val {
			idx = m
			break
		} else if c.sortedSet[m] < val {
			l = m + 1
		} else if c.sortedSet[m] > val {
			r = m - 1
		}
	}
	if idx != -1 {
		c.sortedSet = append(c.sortedSet[:idx], c.sortedSet[idx+1:]...)
	}
}

func (c *ConsistentSharder) hash(key string) uint64 {
	out := blake2b.Sum512([]byte(key))
	return binary.LittleEndian.Uint64(out[:])
}
