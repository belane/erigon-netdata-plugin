package erigon

// sync
const (
	syncExecution = "sync_execution"
	syncFinish    = "sync_finish"
	syncHeaders   = "sync_headers"
	sync          = "sync"
)

// rate
const (
	dbCommitSum      = "db_commit_seconds_sum"
	dbCommitSumTotal = "db_commit_seconds_sum_total"
	dbCommitSumSync  = "db_commit_seconds_sum_sync"
	dbCommitSumWrite = "db_commit_seconds_sum_write"

	dbTotalSize      = "db_size"
	dbTableLoglSize  = "table_log_size"
	dbTableSCSSize   = "table_scs_size"
	dbTableStateSize = "table_state_size"
	dbTableTxSize    = "table_tx_size"

	txPoolInvalid          = "txpool_invalid"
	txPoolPending          = "txpool_pending"
	txPoolLocal            = "txpool_local"
	txPoolPendingDiscard   = "txpool_pending_discard"
	txPoolNofunds          = "txpool_pending_nofunds"
	txPoolPendingRatelimit = "txpool_pending_ratelimit"
	txPoolPendingReplace   = "txpool_pending_replace"
	txPoolQueuedDiscard    = "txpool_queued_discard"
	txPoolQueuedEviction   = "txpool_queued_eviction"
	txPoolQueuedNofunds    = "txpool_queued_nofunds"
	txPoolQueuedRatelimit  = "txpool_queued_ratelimit"
)

const (
	p2pEgress  = "p2p_egress"
	p2pIngress = "p2p_ingress"

	p2pPeers  = "p2p_peers"
	p2pServes = "p2p_serves"
	p2pDials  = "p2p_dials"

	rpcRequests = "rpc_total"
	rpcFailure  = "rpc_failure"

	reorgsInvalidTx = "chain_reorg_invalidTx"

	goRoutines = "go_goroutines"
)
