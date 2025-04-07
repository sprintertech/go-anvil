package anvil_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lmittmann/w3"
	"github.com/sprintertech/go-anvil"
)

func TestClient(t *testing.T) {
	port := 8547
	chainid := 13451

	cli := anvil.New(
		anvil.WithPort(port),
		anvil.WithChainID(chainid),
	)

	err := cli.Start()
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Stop()

	ethcli, err := ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%d", port))
	if err != nil {
		t.Fatal(err)
	}

	addr := common.HexToAddress("0xc0de000000000000000000000000000000000000")
	balance := w3.I("53 eth")
	err = cli.SetBalance(addr, balance)
	if err != nil {
		t.Fatal(err)
	}

	bal, err := ethcli.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		t.Fatal(err)
	}

	if bal.Cmp(balance) != 0 {
		t.Fatalf("unexpected balance actual: %s, expected: %s", bal, balance)
	}

	id, err := ethcli.ChainID(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if id.Uint64() != uint64(chainid) {
		t.Fatalf("chain ids do not match, actual: %s expected: %d", id, chainid)
	}
}
