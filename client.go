package anvil

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client is an RPC client for anvil specific functions
type Client struct {
	cli *rpc.Client
}

// NewClient creates a new client with the given RPC client
func NewClient(cli *rpc.Client) *Client {
	return &Client{cli: cli}
}

// Dial creates a new client for the given URL.
func Dial(addr string) (*Client, error) {
	cli, err := rpc.Dial(addr)
	if err != nil {
		return nil, err
	}

	return NewClient(cli), nil
}

// SetBalance sets the balance of a given account.
//
// Equivalent to the `anvil_setBalance` RPC call.
func (c *Client) SetBalance(account common.Address, balance *big.Int) error {
	return c.cli.Call(nil, "anvil_setBalance", account, "0x"+balance.Text(16))
}

// SetNonce sets the transaction nonce for a given account.
//
// Equivalent to the `anvil_setNonce` RPC call.
func (c *Client) SetNonce(account common.Address, nonce uint64) error {
	return c.cli.Call(nil, "anvil_setNonce", account, fmt.Sprintf("0x%x", nonce))
}

// SetCode sets the EVM bytecode at the specified account.
//
// Equivalent to the `anvil_setCode` RPC call.
func (c *Client) SetCode(account common.Address, code []byte) error {
	return c.cli.Call(nil, "anvil_setCode", account, fmt.Sprintf("0x%x", code))
}

// SetStorageAt writes a single storage slot at a given account.
//
// Equivalent to the `anvil_setStorageAt` RPC call.
func (c *Client) SetStorageAt(account common.Address, slot string, value string) error {
	return c.cli.Call(nil, "anvil_setStorageAt", account, slot, value)
}

// SetMinGasPrice sets the minimum gas price for the node.
//
// Equivalent to the `anvil_setMinGasPrice` RPC call.
func (c *Client) SetMinGasPrice(price *big.Int) error {
	return c.cli.Call(nil, "anvil_setMinGasPrice", "0x"+price.Text(16))
}

// SetNextBlockBaseFeePerGas sets the base fee for the next block.
//
// Equivalent to the `anvil_setNextBlockBaseFeePerGas` RPC call.
func (c *Client) SetNextBlockBaseFeePerGas(fee *big.Int) error {
	return c.cli.Call(nil, "anvil_setNextBlockBaseFeePerGas", "0x"+fee.Text(16))
}

// SetChainId sets the chain ID of the node.
//
// Equivalent to the `anvil_setChainId` RPC call.
func (c *Client) SetChainId(id uint64) error {
	return c.cli.Call(nil, "anvil_setChainId", fmt.Sprintf("0x%x", id))
}

// SetCoinbase sets the coinbase address for block rewards.
//
// Equivalent to the `anvil_setCoinbase` RPC call.
func (c *Client) SetCoinbase(addr common.Address) error {
	return c.cli.Call(nil, "anvil_setCoinbase", addr)
}

// SetLoggingEnabled enables or disables logging output.
//
// Equivalent to the `anvil_setLoggingEnabled` RPC call.
func (c *Client) SetLoggingEnabled(enable bool) error {
	return c.cli.Call(nil, "anvil_setLoggingEnabled", enable)
}

// Reset resets the node state to the original or a new forked state.
//
// Equivalent to the `anvil_reset` RPC call.
func (c *Client) Reset(forkURL string) error {
	if forkURL == "" {
		return c.cli.Call(nil, "anvil_reset")
	}
	config := map[string]interface{}{"forking": map[string]string{"jsonRpcUrl": forkURL}}
	return c.cli.Call(nil, "anvil_reset", config)
}

// DumpState returns a hex-encoded snapshot of the entire chain state.
//
// Equivalent to the `anvil_dumpState` RPC call.
func (c *Client) DumpState(out *string) error {
	return c.cli.Call(out, "anvil_dumpState")
}

// LoadState merges a previously dumped state into the current chain state.
//
// Equivalent to the `anvil_loadState` RPC call.
func (c *Client) LoadState(state string) error {
	return c.cli.Call(nil, "anvil_loadState", state)
}

// NodeInfo retrieves the current node configuration parameters.
//
// Equivalent to the `anvil_nodeInfo` RPC call.
func (c *Client) NodeInfo(info *map[string]interface{}) error {
	return c.cli.Call(info, "anvil_nodeInfo")
}
