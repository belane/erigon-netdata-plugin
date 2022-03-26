package erigon

import "github.com/netdata/go.d.plugin/agent/module"

type (
	Charts = module.Charts
	Chart  = module.Chart
	Dims   = module.Dims
	Dim    = module.Dim
)

var charts = Charts{
	chartDbCommitSum.Copy(),
	chartChainSize.Copy(),
	chartChainSync.Copy(),
	chartP2PNetwork.Copy(),
	chartNumberOfPeers.Copy(),
	chartp2pDialsServes.Copy(),
	chartReorgsTxs.Copy(),
	chartGoRoutines.Copy(),
	chartTxPoolCurrent.Copy(),
	chartTxPoolQueued.Copy(),
	chartTxPoolPending.Copy(),
	chartRpcInformation.Copy(),
}

var (
	chartDbCommitSum = Chart{
		ID:    "db_commit_sum",
		Title: "DB commit on disk",
		Units: "bytes/s",
		Fam:   "db",
		Ctx:   "erigon.db_commit_sum",
		Dims: Dims{
			{ID: dbCommitSumSync, Name: "sync"},
			{ID: dbCommitSumWrite, Name: "writes"},
			{ID: dbCommitSumTotal, Name: "total"},
		},
	}
	chartChainSize = Chart{
		ID:    "chain_db_size",
		Title: "Chain Tables Size",
		Units: "bytes",
		Fam:   "chain",
		Ctx:   "erigon.chain_db_size",
		Dims: Dims{
			{ID: dbTableLoglSize, Name: "log"},
			{ID: dbTableSCSSize, Name: "scs"},
			{ID: dbTableStateSize, Name: "state"},
			{ID: dbTableTxSize, Name: "tx"},
			{ID: dbTotalSize, Name: "total"},
		},
	}
	chartChainSync = Chart{
		ID:    "chain_sync",
		Title: "Chain Sync",
		Units: "block",
		Fam:   "chain",
		Ctx:   "erigon.chain_sync",
		Dims: Dims{
			{ID: syncHeaders, Name: "headers"},
			{ID: syncFinish, Name: "finish"},
			{ID: syncExecution, Name: "execution"},
		},
	}
	chartP2PNetwork = Chart{
		ID:    "p2p_network",
		Title: "P2P bandwidth",
		Units: "bytes/s",
		Fam:   "p2p_bandwidth",
		Ctx:   "erigon.p2p_bandwidth",
		Dims: Dims{
			{ID: p2pIngress, Name: "ingress", Algo: "incremental"},
			{ID: p2pEgress, Name: "egress", Mul: -1, Algo: "incremental"},
		},
	}
	chartNumberOfPeers = Chart{
		ID:    "p2p_peers_number",
		Title: "Number of Peers",
		Units: "peers",
		Fam:   "p2p_peers",
		Ctx:   "erigon.p2p_peers",
		Dims: Dims{
			{ID: p2pPeers, Name: "Peers"},
		},
	}

	chartp2pDialsServes = Chart{
		ID:    "p2p_dials_serves",
		Title: "P2P Serves and Dials",
		Units: "calls/s",
		Fam:   "p2p_peers",
		Ctx:   "erigon.p2p_peers",
		Dims: Dims{
			{ID: p2pDials, Name: "dials", Algo: "incremental"},
			{ID: p2pServes, Name: "serves", Algo: "incremental"},
		},
	}
	chartReorgsTxs = Chart{
		ID:    "reorgs_tx",
		Title: "Invalid Transactions from Reorg",
		Units: "transactions",
		Fam:   "reorgs",
		Ctx:   "erigon.reorgs_tx",
		Dims: Dims{
			{ID: reorgsInvalidTx, Name: "tx"},
		},
	}
	chartGoRoutines = Chart{
		ID:    "goroutines",
		Title: "Number of goroutines",
		Units: "goroutines",
		Fam:   "goroutines",
		Ctx:   "erigon.goroutines",
		Dims: Dims{
			{ID: goRoutines, Name: "goroutines"},
		},
	}
	chartTxPoolCurrent = Chart{
		ID:    "txpoolcurrent",
		Title: "Transaction Pool",
		Units: "transactions",
		Fam:   "tx_pool",
		Ctx:   "erigon.tx_pool_current",
		Dims: Dims{
			{ID: txPoolInvalid, Name: "invalid"},
			{ID: txPoolPending, Name: "pending"},
			{ID: txPoolLocal, Name: "local"},
			{ID: txPoolNofunds, Name: "pool"},
		},
	}
	chartTxPoolQueued = Chart{
		ID:    "txpoolqueued",
		Title: "Queued Transaction Pool",
		Units: "transactions",
		Fam:   "tx_pool",
		Ctx:   "erigon.tx_pool_queued",
		Dims: Dims{
			{ID: txPoolQueuedDiscard, Name: "discard"},
			{ID: txPoolQueuedEviction, Name: "eviction"},
			{ID: txPoolQueuedNofunds, Name: "no_funds"},
			{ID: txPoolQueuedRatelimit, Name: "ratelimit"},
		},
	}
	chartTxPoolPending = Chart{
		ID:    "txpoolpending",
		Title: "Pending Transaction Pool",
		Units: "transactions",
		Fam:   "tx_pool",
		Ctx:   "erigon.tx_pool_pending",
		Dims: Dims{
			{ID: txPoolInvalid, Name: "invalid"},
			{ID: txPoolPending, Name: "pending"},
			{ID: txPoolLocal, Name: "local"},
			{ID: txPoolPendingDiscard, Name: " discard"},
			{ID: txPoolNofunds, Name: "no funds"},
			{ID: txPoolPendingRatelimit, Name: "ratelimit"},
			{ID: txPoolPendingReplace, Name: "replace"},
		},
	}
	chartRpcInformation = Chart{
		ID:    "rpc_calls",
		Title: "rpc calls",
		Units: "calls/s",
		Fam:   "rpc",
		Ctx:   "erigon.rpc_calls",
		Dims: Dims{
			{ID: rpcRequests, Name: "total", Algo: "incremental"},
			{ID: rpcFailure, Name: "failed", Algo: "incremental"},
		},
	}
)
