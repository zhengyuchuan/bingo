package db

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/op/go-logging"
	"github.com/samuel/go-zookeeper/zk"
	"sync"
	"time"
)

type proxyInfo struct {
	ProtoType string
	ProxyAddr string
}

type Pool struct {
	nextIdx   int
	pools     []redis.Conn
	zk        zk.Conn
	ZkServers []string
	ZkTimeout time.Duration
	ZkDir     string
	Dial      func(network, address string) (redis.Conn, error)
	mu        *sync.Mutex
}

var (
	Log = logging.MustGetLogger("Codis")
)

func initCodis() *Pool {
	p := &Pool{
		ZkServers: []string{"locallhost:8021"},
		ZkTimeout: time.Second * 60,
		ZkDir:     "/codis/proxy",
		Dial: func(network, address string) (redis.Conn, error) {
			conn, err := redis.Dial(network, address)
			if err == nil {
				conn.Send("AUTH", "PASSWORD")
			}
			return conn, err
		},
		mu: &sync.Mutex{},
	}
	return p
}

// initFromZk 从zk获取proxy连接
func (this *Pool) initFromZk() {
	this.initZk()
	this.pools = []redis.Conn{}
	children, _, err := this.zk.Children(this.ZkDir) // 获取子节点

	if err != nil {
		Log.Fatal(err)
	}

	for _, child := range children {
		data, _, err := this.zk.Get(this.ZkDir + "/" + child)

		if err != nil {
			continue
		}

		var p proxyInfo

		json.Unmarshal(data, &p)
		conn, err := this.Dial(p.ProtoType, p.ProxyAddr)
		if err != nil {
			Log.Errorf("Create redis connection failed: %s", err.Error())
			continue
		}
		this.pools = append(this.pools, conn)
	}

	go this.watch(this.ZkDir)
}

func (this *Pool) watch(node string) {
	for {
		_, _, ch, err := this.zk.ChildrenW(node) // 获取子节点并监控子节点Event
		if err != nil {
			Log.Error(err)
			return
		}
		evt := <-ch

		if evt.Type == zk.EventSession {
			if evt.State == zk.StateConnecting {
				continue
			}
			if evt.State == zk.StateExpired {
				this.zk.Close()
				Log.Info("Zookeeper session expired, reconnecting...")
				this.initZk()
			}
		}
		if evt.State == zk.StateConnected {
			switch evt.Type {
			case
				zk.EventNodeCreated,
				zk.EventNodeDeleted,
				zk.EventNodeChildrenChanged,
				zk.EventNodeDataChanged:
				this.initFromZk()
				return
			}
			continue
		}
	}
}

// initZk 连接zk
func (this *Pool) initZk() {
	zkConn, _, err := zk.Connect(this.ZkServers, this.ZkTimeout)
	if err != nil {
		Log.Fatalf("Failed to connect to zookeeper: %+v", err)
	}
	this.zk = *zkConn
}

func (this *Pool) Get() redis.Conn {
	this.mu.Lock()
	if len(this.pools) == 0 {
		this.initFromZk()
	}
	this.nextIdx += 1
	if this.nextIdx >= len(this.pools) {
		this.nextIdx = 0
	}
	if len(this.pools) == 0 {
		this.mu.Unlock()
		err := errors.New("Proxy list empty")
		Log.Error(err)
		return errorConnection{err: err}
	} else {
		c := this.pools[this.nextIdx]
		this.mu.Unlock()
		return c
	}
}

func (this *Pool) Close() error {
	this.zk.Close()
	this.pools = []redis.Conn{}
	return nil
}

type errorConnection struct{ err error }

func (ec errorConnection) Do(string, ...interface{}) (interface{}, error) { return nil, ec.err }
func (ec errorConnection) Send(string, ...interface{}) error              { return ec.err }
func (ec errorConnection) Err() error                                     { return ec.err }
func (ec errorConnection) Close() error                                   { return ec.err }
func (ec errorConnection) Flush() error                                   { return ec.err }
func (ec errorConnection) Receive() (interface{}, error)                  { return nil, ec.err }
