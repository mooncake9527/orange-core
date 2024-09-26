package gLocal

import (
	"fmt"
	"runtime"
	"sync"
)

// 作为协程私有空间

type GLocal struct {
	AppKey    string
	ReqId     string
	UserId    string
	CompanyId string
}

type goroutineLocalColl struct {
	Coll map[int64]*GLocal
	l    *sync.RWMutex
}

var gLocalColl = &goroutineLocalColl{Coll: make(map[int64]*GLocal), l: &sync.RWMutex{}}

func SetCUId(userId, companyId string) {
	gid := GetGoID()
	gLocalColl.l.Lock()
	defer gLocalColl.l.Unlock()
	if store, ok := gLocalColl.Coll[gid]; ok {
		store.UserId = userId
		store.CompanyId = companyId
	} else {
		gLocalColl.Coll[gid] = &GLocal{UserId: userId}
	}
}

func SetReqId(reqId string) {
	gid := GetGoID()
	gLocalColl.l.Lock()
	defer gLocalColl.l.Unlock()
	if store, ok := gLocalColl.Coll[gid]; ok {
		store.ReqId = reqId
	} else {
		gLocalColl.Coll[gid] = &GLocal{ReqId: reqId}
	}
}

func SetAppKey(key string) {
	gid := GetGoID()
	gLocalColl.l.Lock()
	defer gLocalColl.l.Unlock()
	if store, ok := gLocalColl.Coll[gid]; ok {
		store.AppKey = key
	} else {
		gLocalColl.Coll[gid] = &GLocal{AppKey: key}
	}
}

func GetIds() (string, string, string, string) {
	gid := GetGoID()
	gLocalColl.l.RLock()
	defer gLocalColl.l.RUnlock()
	if store, ok := gLocalColl.Coll[gid]; ok {
		return store.ReqId, store.UserId, store.CompanyId, store.AppKey
	} else {
		return "", "", "", ""
	}
}

func GetGoID() int64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	var id int64
	fmt.Sscanf(string(b), "goroutine %d", &id)
	return id
}
