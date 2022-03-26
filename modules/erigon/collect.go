package erigon

import (
	"strings"

	"github.com/netdata/go.d.plugin/pkg/prometheus"
	"github.com/netdata/go.d.plugin/pkg/stm"
)

func (v *Erigon) collect() (map[string]int64, error) {
	pms, err := v.prom.Scrape()
	if err != nil {
		return nil, err
	}
	mx := v.collectErigon(pms)

	return stm.ToMap(mx), nil
}

func (g *Erigon) collectErigon(pms prometheus.Metrics) map[string]float64 {
	mx := make(map[string]float64)
	g.collectChainData(mx, pms)
	g.collectP2P(mx, pms)
	g.collectTxPool(mx, pms)
	g.collectRpc(mx, pms)
	return mx
}

func (v *Erigon) collectChainData(mx map[string]float64, pms prometheus.Metrics) {
	pms = pms.FindByNames(
		goRoutines,
		sync,
		reorgsInvalidTx,
		dbCommitSum,
		dbTotalSize,
		dbTableLoglSize,
		dbTableSCSSize,
		dbTableStateSize,
		dbTableTxSize,
	)
	v.processCollect(mx, pms)
}

func (v *Erigon) collectRpc(mx map[string]float64, pms prometheus.Metrics) {
	pms = pms.FindByNames(
		rpcRequests,
		rpcFailure,
	)
	v.processCollect(mx, pms)
}

func (v *Erigon) collectTxPool(mx map[string]float64, pms prometheus.Metrics) {
	pms = pms.FindByNames(
		txPoolInvalid,
		txPoolPending,
		txPoolLocal,
		txPoolPendingDiscard,
		txPoolNofunds,
		txPoolPendingRatelimit,
		txPoolPendingReplace,
		txPoolQueuedDiscard,
		txPoolQueuedEviction,
		txPoolQueuedEviction,
		txPoolQueuedRatelimit,
	)
	v.processCollect(mx, pms)
}

func (v *Erigon) collectP2P(mx map[string]float64, pms prometheus.Metrics) {
	pms = pms.FindByNames(
		p2pDials,
		p2pEgress,
		p2pIngress,
		p2pPeers,
		p2pServes,
	)
	v.processCollect(mx, pms)
}

func (v *Erigon) processCollect(mx map[string]float64, pms prometheus.Metrics) {
	for _, pm := range pms {
		metricName := pm.Name()
		if pm.Labels.Len() > 1 {
			values := make([]string, 0, pm.Labels.Len())
			for _, v := range pm.Labels.Map() {
				values = append(values, v)
			}
			metricName = strings.Join(values, "_")
		}
		mx[metricName] += pm.Value
	}
}
