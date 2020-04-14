package server

import (
	"context"
	"github.com/jinzhu/gorm"
	"twirp/db_proxy/datasource"
	"twirp/db_proxy/module"
	"twirp/pb"
)

type DbProxy struct {
}

func (d *DbProxy) AddUrl(ctx context.Context, req *pb.UrlReq) (*pb.UrlResp, error) {
	err := datasource.PgSqlConn.Create(&module.Url{Key: req.Key, Url: req.Url}).Error
	if err != nil {
		return nil,err
	}
	return &pb.UrlResp{Success: true}, nil
}

func (d *DbProxy) CheckUrlEx(ctx context.Context, req *pb.CheckUrlExReq) (*pb.CheckUrlExResp, error) {
	item := &module.Url{}
	err := datasource.PgSqlConn.Where("url = ?", req.Url).First(item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &pb.CheckUrlExResp{Key: ""},nil
		}
		return nil,err
	}
	return &pb.CheckUrlExResp{Key: item.Key},nil
}

func (d *DbProxy) GetUrl(ctx context.Context, req *pb.GetUrlReq) (*pb.GetUrlResp, error) {
	item := &module.Url{}
	err := datasource.PgSqlConn.Where("key = ?", req.Key).First(item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &pb.GetUrlResp{Addr: ""},nil
		}
		return nil,err
	}
	return &pb.GetUrlResp{Addr: item.Url},nil
}
