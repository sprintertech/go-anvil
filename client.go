package anvil

import (
	"errors"
	"fmt"
	"os/exec"
	"sync/atomic"
)

// Node represents an active Anvil client
type Node struct {
	running atomic.Bool
	cmd     *exec.Cmd
}

// New creates a new Node configured with the passed options
func New(opts ...Option) *Node {
	var args []string
	for _, opt := range opts {
		args = append(args, opt...)
	}

	return &Node{
		cmd: exec.Command("anvil", args...),
	}
}

// Start starts the anvil node
// @TODO: should return available accounts on startup or store them
func (n *Node) Start() error {
	defer n.running.Store(false)

	if n.running.Load() {
		return errors.New("node is running")
	}

	n.running.Store(true)

	return n.cmd.Start()
}

// Stop stops the anvil node
func (n *Node) Stop() error {
	n.running.Store(false)
	return n.cmd.Process.Kill()
}

// An Option configures a [Node]
type Option []string

// WithBlockTime sets the block time in seconds for interval mining.
//
// Equivalent to the `-b, --block-time <BLOCK-TIME>` flag
func WithBlockTime(seconds int) Option {
	return []string{"--block-time", fmt.Sprintf("%d", seconds)}
}

// WithBalance sets the initial balance of accounts.
//
// Equivalent to the `--balance <BALANCE>` flag
func WithBalance(balance int) Option {
	return []string{"--balance", fmt.Sprintf("%d", balance)}
}

// WithDerivationPath sets the derivation path for HD wallets.
//
// Equivalent to the `--derivation-path <DERIVATION_PATH>` flag
func WithDerivationPath(path string) Option {
	return []string{"--derivation-path", path}
}

// WithHelp prints help information.
//
// Equivalent to the `-h, --help` flag
func WithHelp() Option {
	return []string{"--help"}
}

// WithHardfork sets the EVM hardfork to use.
//
// Equivalent to the `--hardfork <HARDFORK>` flag
func WithHardfork(name string) Option {
	return []string{"--hardfork", name}
}

// WithInit initializes the genesis block using a genesis.json file.
//
// Equivalent to the `--init <PATH>` flag
func WithInit(path string) Option {
	return []string{"--init", path}
}

// WithMnemonic sets the BIP39 mnemonic phrase for generating accounts.
//
// Equivalent to the `-m, --mnemonic <MNEMONIC>` flag
func WithMnemonic(mnemonic string) Option {
	return []string{"--mnemonic", mnemonic}
}

// WithNoMining disables auto and interval mining.
//
// Equivalent to the `--no-mining` flag
func WithNoMining() Option {
	return []string{"--no-mining"}
}

// WithOrder sets the transaction ordering strategy in the mempool.
//
// Equivalent to the `--order <ORDER>` flag
func WithOrder(order string) Option {
	return []string{"--order", order}
}

// WithPort sets the listening port.
//
// Equivalent to the `-p, --port <PORT>` flag
func WithPort(port int) Option {
	return []string{"--port", fmt.Sprintf("%d", port)}
}

// WithStepsTracing enables steps tracing for geth-style traces.
//
// Equivalent to the `--steps-tracing` flag (alias: `--tracing`)
func WithStepsTracing() Option {
	return []string{"--steps-tracing"}
}

// WithIPC starts an IPC endpoint at a given path (optional).
//
// Equivalent to the `--ipc [<PATH>]` flag
func WithIPC(path string) Option {
	if path == "" {
		return []string{"--ipc"}
	}
	return []string{"--ipc", path}
}

// WithSilent disables all startup logs.
//
// Equivalent to the `--silent` flag
func WithSilent() Option {
	return []string{"--silent"}
}

// WithTimestamp sets the timestamp of the genesis block.
//
// Equivalent to the `--timestamp <TIMESTAMP>` flag
func WithTimestamp(ts int64) Option {
	return []string{"--timestamp", fmt.Sprintf("%d", ts)}
}

// WithVersion prints version information.
//
// Equivalent to the `-V, --version` flag
func WithVersion() Option {
	return []string{"--version"}
}

// WithDisableDefaultCreate2Deployer disables the default CREATE2 factory.
//
// Equivalent to the `--disable-default-create2-deployer` flag
func WithDisableDefaultCreate2Deployer() Option {
	return []string{"--disable-default-create2-deployer"}
}

// WithForkURL enables state forking from a remote endpoint.
//
// Equivalent to the `-f, --fork-url <URL>` flag
func WithForkURL(url string) Option {
	return []string{"--fork-url", url}
}

// WithForkBlockNumber forks from a specific block number.
//
// Equivalent to the `--fork-block-number <BLOCK>` flag
func WithForkBlockNumber(block int) Option {
	return []string{"--fork-block-number", fmt.Sprintf("%d", block)}
}

// WithForkRetryBackoff sets initial retry backoff on fork errors.
//
// Equivalent to the `--fork-retry-backoff <BACKOFF>` flag
func WithForkRetryBackoff(backoff int) Option {
	return []string{"--fork-retry-backoff", fmt.Sprintf("%d", backoff)}
}

// WithForkTransactionHash forks state from a specific transaction hash.
//
// Equivalent to the `--fork-transaction-hash <TRANSACTION>` flag
func WithForkTransactionHash(tx string) Option {
	return []string{"--fork-transaction-hash", tx}
}

// WithRetries sets the number of retry attempts for network issues.
//
// Equivalent to the `--retries <RETRIES>` flag
func WithRetries(count int) Option {
	return []string{"--retries", fmt.Sprintf("%d", count)}
}

// WithTimeout sets the request timeout for forking mode in ms.
//
// Equivalent to the `--timeout <TIMEOUT>` flag
func WithTimeout(ms int) Option {
	return []string{"--timeout", fmt.Sprintf("%d", ms)}
}

// WithComputeUnitsPerSecond sets the assumed compute units per second.
//
// Equivalent to the `--compute-units-per-second <CUPS>` flag
func WithComputeUnitsPerSecond(cups int) Option {
	return []string{"--compute-units-per-second", fmt.Sprintf("%d", cups)}
}

// WithNoRateLimit disables rate limiting for the node’s provider.
//
// Equivalent to the `--no-rate-limit` flag
func WithNoRateLimit() Option {
	return []string{"--no-rate-limit"}
}

// WithNoStorageCaching disables RPC storage slot caching.
//
// Equivalent to the `--no-storage-caching` flag
func WithNoStorageCaching() Option {
	return []string{"--no-storage-caching"}
}

// WithBaseFee sets the base fee in a block.
//
// Equivalent to the `--base-fee <FEE>` flag
func WithBaseFee(fee int) Option {
	return []string{"--base-fee", fmt.Sprintf("%d", fee)}
}

// WithBlockBaseFeePerGas sets the block base fee per gas.
//
// Equivalent to the `--block-base-fee-per-gas <FEE>` flag
func WithBlockBaseFeePerGas(fee int) Option {
	return []string{"--block-base-fee-per-gas", fmt.Sprintf("%d", fee)}
}

// WithChainID sets the chain ID.
//
// Equivalent to the `--chain-id <CHAIN_ID>` flag
func WithChainID(id int) Option {
	return []string{"--chain-id", fmt.Sprintf("%d", id)}
}

// WithCodeSizeLimit sets the EIP-170 code size limit in bytes.
//
// Equivalent to the `--code-size-limit <CODE_SIZE>` flag
func WithCodeSizeLimit(limit int) Option {
	return []string{"--code-size-limit", fmt.Sprintf("%d", limit)}
}

// WithGasLimit sets the block gas limit.
//
// Equivalent to the `--gas-limit <GAS_LIMIT>` flag
func WithGasLimit(limit int) Option {
	return []string{"--gas-limit", fmt.Sprintf("%d", limit)}
}

// WithGasPrice sets the gas price.
//
// Equivalent to the `--gas-price <GAS_PRICE>` flag
func WithGasPrice(price int) Option {
	return []string{"--gas-price", fmt.Sprintf("%d", price)}
}

// WithAllowOrigin sets the CORS allowed origin.
//
// Equivalent to the `--allow-origin <ALLOW-ORIGIN>` flag
func WithAllowOrigin(origin string) Option {
	return []string{"--allow-origin", origin}
}

// WithNoCORS disables CORS support.
//
// Equivalent to the `--no-cors` flag
func WithNoCORS() Option {
	return []string{"--no-cors"}
}

// WithHost sets the host address to listen on.
//
// Equivalent to the `--host <HOST>` flag
func WithHost(host string) Option {
	return []string{"--host", host}
}

// WithConfigOut writes the Anvil output to a given file as JSON.
//
// Equivalent to the `--config-out <OUT_FILE>` flag
func WithConfigOut(path string) Option {
	return []string{"--config-out", path}
}

// WithPruneHistory disables full chain history retention.
//
// Equivalent to the `--prune-history` flag
func WithPruneHistory() Option {
	return []string{"--prune-history"}
}

// WithNoRequestSizeLimit disables the default request size limit (2MB).
//
// Equivalent to the `--no-request-size-limit` flag
func WithNoRequestSizeLimit() Option {
	return []string{"--no-request-size-limit"}
}
