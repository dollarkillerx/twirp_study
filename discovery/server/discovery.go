package server

import (
	"context"
	"sync"
	"time"
	"twirp/define"
	"twirp/pb"
	"twirp/utils"
)

type Discovery struct {
	dbMu sync.Mutex
	db   map[string][]*pb.Server
}

func New() *Discovery {
	return &Discovery{
		db: map[string][]*pb.Server{},
	}
}

func (d *Discovery) Registry(ctx context.Context, req *pb.RegistryReq) (*pb.RegistryResp, error) {
	id := utils.RandId()
	d.dbMu.Lock()
	defer d.dbMu.Unlock()
ki: // 使用了万恶的goto
	it, ex := d.db[req.Server]
	if !ex {
		d.db[req.Server] = make([]*pb.Server, 0)
		goto ki
	}
	resp := &pb.RegistryResp{
		Id:      id,
		Success: true,
	}
	it = append(it, &pb.Server{
		Id:      id,
		Server:  req.Server,
		Addr:    req.Addr,
		Load:    0,
		Timeout: time.Now().Add(time.Second).Unix(),
	})
	d.db[req.Server] = it
	return resp, nil
}

func (d *Discovery) Discovery(ctx context.Context, req *pb.DiscoveryReq) (*pb.DiscoveryResp, error) {
	resp := &pb.DiscoveryResp{Servers: make([]*pb.Server, 0)}
	d.dbMu.Lock()
	defer d.dbMu.Unlock()
	servers, ex := d.db[req.Server]
	if !ex {
		return resp, nil
	}

	now := time.Now().Unix()
	for _, v := range servers {
		if v.Timeout > now {
			resp.Servers = append(resp.Servers, v)
		}
	}

	if len(resp.Servers) != len(servers) {
		// 把其中过期的填充调
		d.db[req.Server] = resp.Servers
	}

	return resp, nil
}

func (d *Discovery) Heartbeat(ctx context.Context, req *pb.HeartbeatReq) (*pb.HeartbeatResp, error) {
	d.dbMu.Lock()
	defer d.dbMu.Unlock()
	servers, ex := d.db[req.Server]
	if !ex {
		return nil, define.DiscoveryIsNull
	}

	su := make([]*pb.Server, 0)
	for _, v := range servers {
		if v.Id == req.Id {
			v.Timeout = time.Now().Add(time.Second).Unix()
			v.Load = req.Load
		}
		su = append(su, v)
	}

	// 太懒了 就这样实现把
	d.db[req.Server] = su
	return &pb.HeartbeatResp{}, nil
}
