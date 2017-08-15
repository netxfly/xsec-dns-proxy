package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type DnsInfo struct {
	Id         int64
	DomainName string    `xorm:"Domain_name"`
	ClientIp   string `xorm:"client_ip"`
	DnsRecord  []string `xorm:"dns_record"`
	CreateAt   time.Time `xorm:"created"`
}

type MgoDnsInfo struct {
	Id         bson.ObjectId `bson:"_id"`
	DomainName string `bson:"domain_name"`
	ClientIp   string `bson:"client_ip"`
	DnsRecord  []string `bson:"dns_record"`
	CreateAt   time.Time `bson:"created"`
}

func NewDnsInfo(domain, client string, record []string) (*DnsInfo) {
	return &DnsInfo{DomainName: domain, ClientIp: client, DnsRecord: record}
}

func (d *DnsInfo) Insert() (int64, error) {
	return Engine.Insert(d)
}

func Query() ([]DnsInfo, error) {
	dnsInfos := make([]DnsInfo, 0)
	err := Engine.Table("dns_info").Desc("id").Limit(1000, 0).Find(&dnsInfos)
	return dnsInfos, err
}

func NewMgoDnsInfo(domain, client string, record []string) (*MgoDnsInfo) {
	now := time.Now()
	return &MgoDnsInfo{DomainName: domain, ClientIp: client, DnsRecord: record, CreateAt: now, Id: bson.NewObjectId()}
}

func (m *MgoDnsInfo) Insert() (error) {
	_, err := MongodbClient.Collection("dns_info").Insert(m)
	return err
}

func MgoQuery() ([]MgoDnsInfo, error) {
	mgoDnsInfos := make([]MgoDnsInfo, 0)
	res := MongodbClient.Collection("dns_info").Find("-_id").OrderBy().Limit(1000)
	err := res.All(&mgoDnsInfos)
	return mgoDnsInfos, err
}
