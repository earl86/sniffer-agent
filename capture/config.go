package capture

import (
	"math/rand"
	"time"

	sd "sniffer-agent/session-dealer"

	log "github.com/golang/glog"
)

var (
	localIPAddr *string

	sessionPool = make(map[string]sd.ConnSession)
	// sessionPoolLock sync.Mutex
)

func init() {
	ipAddr, err := getLocalIPAddr()
	if err != nil {
		panic(err)
	}

	localIPAddr = &ipAddr

	rand.Seed(time.Now().UnixNano())
}

func ShowLocalIP() {
	log.Infof("parsed local ip address:%s", *localIPAddr)
}
