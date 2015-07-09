package db

import (
	"fmt"
	"path"
	"sync"
	"sync/atomic"

	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/cloudfoundry/gunk/workpool"
	"github.com/pivotal-golang/lager"
)

const maxDesiredLRPGetterWorkPoolSize = 50
const DesiredLRPSchemaRoot = DataSchemaRoot + "desired"

func DesiredLRPSchemaPath(lrp models.DesiredLRP) string {
	return DesiredLRPSchemaPathByProcessGuid(lrp.GetProcessGuid())
}

func DesiredLRPSchemaPathByProcessGuid(processGuid string) string {
	return path.Join(DesiredLRPSchemaRoot, processGuid)
}

func (db *ETCDDB) DesiredLRPs(logger lager.Logger) (*models.DesiredLRPs, error) {
	root, err := db.fetchRecursiveRaw(DesiredLRPSchemaRoot, logger)
	if err != nil {
		return &models.DesiredLRPs{}, nil
	}
	if root.Nodes.Len() == 0 {
		return &models.DesiredLRPs{}, nil
	}

	desiredLRPs := models.DesiredLRPs{}

	lrpsLock := sync.Mutex{}
	var workErr atomic.Value
	works := []func(){}

	for _, node := range root.Nodes {
		node := node

		works = append(works, func() {
			var lrp models.DesiredLRP
			deserializeErr := models.FromJSON([]byte(node.Value), &lrp)
			if deserializeErr != nil {
				logger.Error("failed-parsing-desired-lrp", deserializeErr)
				workErr.Store(fmt.Errorf("cannot parse lrp JSON for key %s: %s", node.Key, deserializeErr.Error()))
				return
			}

			lrpsLock.Lock()
			desiredLRPs.DesiredLrps = append(desiredLRPs.DesiredLrps, &lrp)
			lrpsLock.Unlock()
		})
	}

	throttler, err := workpool.NewThrottler(maxDesiredLRPGetterWorkPoolSize, works)
	if err != nil {
		logger.Error("failed-constructing-throttler", err, lager.Data{"max-workers": maxDesiredLRPGetterWorkPoolSize, "num-works": len(works)})
		return &models.DesiredLRPs{}, err
	}

	logger.Debug("performing-deserialization-work")
	throttler.Work()
	if err, ok := workErr.Load().(error); ok {
		logger.Error("failed-performing-deserialization-work", err)
		return &models.DesiredLRPs{}, err
	}
	logger.Debug("succeeded-performing-deserialization-work", lager.Data{"num-desired-lrps": len(desiredLRPs.GetDesiredLrps())})

	return &desiredLRPs, nil
}