package ethapi

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/tomochain/tomochain/accounts"
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/common/hexutil"
	"github.com/tomochain/tomochain/consensus"
	"github.com/tomochain/tomochain/consensus/ethash"
	"github.com/tomochain/tomochain/core"
	"github.com/tomochain/tomochain/core/state"
	"github.com/tomochain/tomochain/core/types"
	"github.com/tomochain/tomochain/core/vm"
	"github.com/tomochain/tomochain/crypto"
	"github.com/tomochain/tomochain/eth/downloader"
	"github.com/tomochain/tomochain/ethclient"
	"github.com/tomochain/tomochain/ethdb"
	"github.com/tomochain/tomochain/event"
	"github.com/tomochain/tomochain/params"
	"github.com/tomochain/tomochain/rpc"
	"github.com/tomochain/tomochain/tomox"
	"github.com/tomochain/tomochain/tomox/tradingstate"
	"github.com/tomochain/tomochain/tomoxlending"
)

type testBackend struct {
	db      ethdb.Database
	chain   *core.BlockChain
	pending *types.Block
	TomoX   *tomox.TomoX
}

func (t testBackend) Downloader() *downloader.Downloader {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) ProtocolVersion() int {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) ChainDb() ethdb.Database {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) EventMux() *event.TypeMux {
	//TODO implement me
	panic("implement me")
}

func (b testBackend) AccountManager() *accounts.Manager {
	return &accounts.Manager{}
}

func (b testBackend) TomoxService() *tomox.TomoX {
	return b.TomoX
}

func (t testBackend) LendingService() *tomoxlending.Lending {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) SetHead(number uint64) {
	//TODO implement me
	panic("implement me")
}

func (b testBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	if blockNr == rpc.LatestBlockNumber {
		return b.chain.CurrentBlock().Header(), nil
	}
	return b.chain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b testBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	if number == rpc.LatestBlockNumber {
		return b.chain.CurrentBlock(), nil
	}
	if number == rpc.PendingBlockNumber {
		return b.pending, nil
	}
	return b.chain.GetBlockByNumber(uint64(number)), nil
}

func (b testBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.chain.StateAt(header.Root)
	return stateDb, header, err
}

func (t testBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return core.GetBlockReceipts(t.db, blockHash, core.GetBlockNumber(t.db, blockHash)), nil
}

func (t testBackend) GetTd(blockHash common.Hash) *big.Int {
	//TODO implement me
	panic("implement me")
}

func (b testBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, tomoxState *tradingstate.TradingStateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.chain, nil)
	return vm.NewEVM(context, state, tomoxState, b.chain.Config(), vmCfg), vmError, nil
}

func (t testBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetPoolTransactions() (types.Transactions, error) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) Stats() (pending int, queued int) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) SubscribeTxPreEvent(events chan<- core.TxPreEvent) event.Subscription {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) SendOrderTx(ctx context.Context, signedTx *types.OrderTransaction) error {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) OrderTxPoolContent() (map[common.Address]types.OrderTransactions, map[common.Address]types.OrderTransactions) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) OrderStats() (pending int, queued int) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) SendLendingTx(ctx context.Context, signedTx *types.LendingTransaction) error {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) ChainConfig() *params.ChainConfig {
	return t.chain.Config()
}

func (t testBackend) CurrentBlock() *types.Block {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetIPCClient() (*ethclient.Client, error) {
	//TODO implement me
	panic("implement me")
}

func (b testBackend) GetEngine() consensus.Engine {
	return b.chain.Engine()
}

func (t testBackend) GetRewardByHash(hash common.Hash) map[string]map[string]map[string]*big.Int {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetVotersRewards(address common.Address) map[common.Address]*big.Int {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetVotersCap(checkpoint *big.Int, masterAddr common.Address, voters []common.Address) map[common.Address]*big.Int {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetEpochDuration() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetMasternodesCap(checkpoint uint64) map[common.Address]*big.Int {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetBlocksHashCache(blockNr uint64) []common.Hash {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) AreTwoBlockSamePath(newBlock common.Hash, oldBlock common.Hash) bool {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetOrderNonce(address common.Hash) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func newTestBackend(t *testing.T, n int, gspec *core.Genesis, generator func(i int, b *core.BlockGen)) *testBackend {
	var (
		engine      = ethash.NewFaker()
		cacheConfig = &core.CacheConfig{
			TrieTimeLimit: 5 * time.Minute,
			TrieNodeLimit: 256 * 1024 * 1024,
		}
	)
	// Generate blocks for testing
	db, blocks, _ := core.GenerateChainWithGenesis(gspec, engine, n, generator)
	chain, err := core.NewBlockChain(db, cacheConfig, params.TestChainConfig, engine, vm.Config{})
	if err != nil {
		t.Fatalf("failed to create tester chain: %v", err)
	}
	if n, err := chain.InsertChain(blocks); err != nil {
		t.Fatalf("block %d: failed to insert into chain: %v", n, err)
	}

	tomo := tomox.New(&tomox.DefaultConfig)

	backend := &testBackend{db: db, chain: chain, TomoX: tomo}
	return backend
}

type Account struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}

func newAccounts(n int) (accounts []Account) {
	for i := 0; i < n; i++ {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		accounts = append(accounts, Account{key: key, addr: addr})
	}
	slices.SortFunc(accounts, func(a, b Account) int {
		return a.addr.Cmp(b.addr)
	})
	return accounts
}

func TestEstimateGas(t *testing.T) {
	t.Parallel()
	// Initialize test accounts
	var (
		accounts = newAccounts(2)
		genesis  = &core.Genesis{
			Config: params.TestChainConfig,
			Alloc: core.GenesisAlloc{
				accounts[0].addr: {Balance: big.NewInt(params.Ether)},
				accounts[1].addr: {Balance: big.NewInt(params.Ether)},
			},
		}
		genBlocks = 10
		signer    = types.HomesteadSigner{}
	)
	api := NewPublicBlockChainAPI(newTestBackend(t, genBlocks, genesis, func(i int, b *core.BlockGen) {
		// Transfer from account[0] to account[1]
		//    value: 1000 wei
		//    fee:   0 wei
		//tx, _ := types.SignTx(types.NewTx(&types.LegacyTx{Nonce: uint64(i), To: &accounts[1].addr, Value: big.NewInt(1000), Gas: params.TxGas, GasPrice: b.BaseFee(), Data: nil}), signer, accounts[0].key)
		tx, err := types.SignTx(types.NewTransaction(uint64(i), accounts[1].addr, big.NewInt(1000), params.TxGas, nil, nil), signer, accounts[0].key)
		if err != nil {
			panic(err)
		}
		b.AddTx(tx)
	}))
	var testSuite = []struct {
		blockNumber rpc.BlockNumber
		call        CallArgs
		expectErr   error
		want        uint64
	}{
		// simple transfer on latest block
		{
			blockNumber: rpc.LatestBlockNumber,
			call: CallArgs{
				From:  accounts[0].addr,
				To:    &accounts[1].addr,
				Value: (hexutil.Big)(*big.NewInt(1000)),
			},
			expectErr: nil,
			want:      21000,
		},
		// empty create
		{
			blockNumber: rpc.LatestBlockNumber,
			call:        CallArgs{},
			expectErr:   nil,
			want:        53000,
		},
	}
	for i, tc := range testSuite {
		result, err := api.EstimateGas(context.Background(), tc.call, &tc.blockNumber)
		if tc.expectErr != nil {
			if err == nil {
				t.Errorf("test %d: want error %v, have nothing", i, tc.expectErr)
				continue
			}
			if !errors.Is(err, tc.expectErr) {
				t.Errorf("test %d: error mismatch, want %v, have %v", i, tc.expectErr, err)
			}
			continue
		}
		if err != nil {
			t.Errorf("test %d: want no error, have %v", i, err)
			continue
		}
		if uint64(result) != tc.want {
			t.Errorf("test %d, result mismatch, have\n%v\n, want\n%v\n", i, uint64(result), tc.want)
		}
	}
}

func TestPublicBlockChainAPI_GetProof(t *testing.T) {
	t.Parallel()
	var (
		accounts = newAccounts(2)
		genesis  = &core.Genesis{
			Config: params.TestChainConfig,
			Alloc: core.GenesisAlloc{
				accounts[0].addr: {Balance: big.NewInt(params.Ether)},
				accounts[1].addr: {Balance: big.NewInt(params.Ether)},
			},
		}
		genBlocks = 10
		signer    = types.HomesteadSigner{}
	)

	backend := newTestBackend(t, genBlocks, genesis, func(i int, b *core.BlockGen) {
		tx, err := types.SignTx(types.NewTransaction(uint64(i), accounts[1].addr, big.NewInt(1000), params.TxGas, nil, nil), signer, accounts[0].key)
		if err != nil {
			t.Fatal(err)
		}
		b.AddTx(tx)
	})

	api := NewPublicBlockChainAPI(backend)

	testCases := []struct {
		name        string
		address     common.Address
		storageKeys []string
		blockNr     rpc.BlockNumber
		wantErr     bool
		errMsg      string
		expected    *AccountResult
	}{
		{
		    name:        "Valid account proof latest block",
		    address:     accounts[0].addr,
		    storageKeys: []string{},
		    blockNr:     rpc.LatestBlockNumber,
		    wantErr:     false,
		},
		{
		    name:        "Valid account with storage proof",
		    address:     accounts[0].addr,
		    storageKeys: []string{"0x0000000000000000000000000000000000000000000000000000000000000000"},
		    blockNr:     rpc.LatestBlockNumber,
		    wantErr:     false,
		},
		{
		    name:        "Non-existent account",
		    address:     common.HexToAddress("0x1234567890123456789012345678901234567890"),
		    storageKeys: []string{},
		    blockNr:     rpc.LatestBlockNumber,
		    wantErr:     false,
		},
		{
			name:        "Invalid block number",
			address:     accounts[0].addr,
			storageKeys: []string{},
			blockNr:     rpc.BlockNumber(-5), // Using -5 to ensure it's invalid
			wantErr:     false,
			expected:    nil,
		},
		// {
		//     name:        "Pending block",
		//     address:     accounts[0].addr,
		//     storageKeys: []string{},
		//     blockNr:     rpc.PendingBlockNumber,
		//     wantErr:     true,
		//     errMsg:      "proof not supported for pending block",
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := api.GetProof(context.Background(), tc.address, tc.storageKeys, tc.blockNr)

			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error containing '%s' but got none", tc.errMsg)
					return
				}
				if !strings.Contains(err.Error(), tc.errMsg) {
					t.Errorf("expected error containing '%s', got '%v'", tc.errMsg, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Verify result fields
			if result == nil && tc.expected == nil {
				return
			}

			if result.Address != tc.address {
				t.Errorf("address mismatch: got %v, want %v", result.Address, tc.address)
			}

			if len(result.AccountProof) == 0 {
			    t.Error("account proof should not be empty")
			}

			if result.Balance == nil {
			    t.Error("balance should not be nil")
			}

			if result.CodeHash == (common.Hash{}) {
			    t.Error("codehash should not be empty")
			}

			if result.StorageHash == (common.Hash{}) {
			    t.Error("storagehash should not be empty")
			}

			if len(tc.storageKeys) > 0 {
			    if len(result.StorageProof) != len(tc.storageKeys) {
			        t.Errorf("storage proof length mismatch: got %d, want %d",
			            len(result.StorageProof), len(tc.storageKeys))
			    }
			}
		})
	}
}
