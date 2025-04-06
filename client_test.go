package anvil_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sprintertech/go-anvil"
)

func TestClient(t *testing.T) {
	port := 8547
	chainid := 13451

	cli := anvil.New(
		anvil.WithPort(port),
		anvil.WithChainID(chainid),
	)

	defer cli.Stop()

	go func() {
		err := cli.Run()
		if err != nil {
			t.Fatal(err) // @TODO
		}
	}()

	time.Sleep(1 * time.Second)

	ethcli, err := ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%d", port))
	if err != nil {
		t.Fatal(err)
	}

	id, err := ethcli.ChainID(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if id.Uint64() != uint64(chainid) {
		t.Fatalf("chain ids do not match, actual: %s expected: %d", id, chainid)
	}
}
