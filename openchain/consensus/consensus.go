/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package consensus

import pb "github.com/openblockchain/obc-peer/protos"

// Consenter is implemented by every consensus plugin package
type Consenter interface {
	RecvMsg(msg *pb.OpenchainMessage) error
}

type Ledger interface {
	GetBlock(id uint64) (block *pb.Block, err error)
	GetCurrentStateHash() (stateHash []byte, err error)
	GetBlockchainSize() uint64
	HashBlock(block *pb.Block) ([]byte, error)
	PutBlock(blockNumber uint64, block *pb.Block)
	ApplyStateDelta(delta []byte, unapply bool)
	EmptyState()
	VerifyBlockchain(start, finish uint64) (uint64, error)

	GetRemoteBlocks(replicaId uint64, start, finish uint64) (<-chan *pb.SyncBlocks, error)
	GetRemoteStateSnapshot(replicaId uint64) (<-chan *pb.SyncStateSnapshot, error)
	GetRemoteStateDeltas(replicaId uint64, start, finish uint64) (<-chan *pb.SyncStateDeltas, error)
}

// CPI (Consensus Programming Interface) is the set of
// stack-facing methods available to the consensus plugin
type CPI interface {
	GetNetworkHandles() (self string, network []string, err error)
	GetReplicaHandle(id uint64) (handle string, err error)
	GetReplicaID(handle string) (id uint64, err error)

	Broadcast(msg *pb.OpenchainMessage) error
	Unicast(msg *pb.OpenchainMessage, receiverHandle string) error

	BeginTxBatch(id interface{}) error
	ExecTXs(txs []*pb.Transaction) ([]byte, []error)
	CommitTxBatch(id interface{}, transactions []*pb.Transaction, proof []byte) error
	RollbackTxBatch(id interface{}) error

	Ledger
}
