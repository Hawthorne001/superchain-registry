package manage

import (
	"os"
	"testing"

	"github.com/ethereum-optimism/superchain-registry/ops/internal/config"
	"github.com/ethereum-optimism/superchain-registry/ops/internal/paths"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

const (
	validAddressesPath   = "testdata/depsets_valid/addresses.json"
	invalidAddressesPath = "testdata/depsets_invalid/addresses.json"
)

// Helper function to load addresses.json
func loadAddresses(t *testing.T, path string) config.AddressesJSON {
	var addrs config.AddressesJSON
	err := paths.ReadJSONFile(path, &addrs)
	require.NoError(t, err)
	return addrs
}

// Helper function to load chain configs
func loadChainConfig(t *testing.T, path string) *config.Chain {
	var cfg config.Chain
	err := paths.ReadTOMLFile(path, &cfg)
	require.NoError(t, err)
	return &cfg
}

func TestDepsetChecker(t *testing.T) {
	lgr := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, false))

	t.Run("validate test depsets", func(t *testing.T) {
		addrs := loadAddresses(t, validAddressesPath)

		cfgs, err := CollectChainConfigs("testdata/depsets_valid")
		require.NoError(t, err)

		checker := NewDepsetChecker(lgr, cfgs, addrs)
		require.NoError(t, checker.Check())
	})

	t.Run("validate actual depsets", func(t *testing.T) {
		rootDir, err := paths.FindRepoRoot()
		require.NoError(t, err)
		require.NotNil(t, rootDir)

		superchainCfgsDir := paths.SuperchainConfigsDir(rootDir)
		require.DirExists(t, superchainCfgsDir)

		addrs := loadAddresses(t, paths.AddressesFile(rootDir))
		cfgs, err := CollectChainConfigs(superchainCfgsDir)
		require.NoError(t, err)

		checker := NewDepsetChecker(lgr, cfgs, addrs)
		require.NoError(t, checker.Check())
	})

	t.Run("invalid depsets (transience)", func(t *testing.T) {
		addrs := loadAddresses(t, "testdata/depsets_invalid/transience/addresses.json")
		chains, err := CollectChainConfigs("testdata/depsets_invalid/transience")
		require.NoError(t, err)

		checker := NewDepsetChecker(lgr, chains, addrs)
		err = checker.Check()
		require.Error(t, err)
		require.ErrorIs(t, err, errInconsistentDepsets)
	})
}

func TestDepsetChecker_checkOffchain(t *testing.T) {
	lgr := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, false))

	t.Run("no chain configs", func(t *testing.T) {
		checker := NewDepsetChecker(lgr, nil, nil)
		err := checker.checkOffchain(nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "no chain configs provided to checkOffchain")
	})

	t.Run("inconsistent depset lengths", func(t *testing.T) {
		addrs := loadAddresses(t, validAddressesPath)
		var cfg1 config.Chain
		err := paths.ReadTOMLFile("testdata/depsets_invalid/depset_length_1.toml", &cfg1)
		require.NoError(t, err)

		var cfg2 config.Chain
		err = paths.ReadTOMLFile("testdata/depsets_invalid/depset_length_2.toml", &cfg2)
		require.NoError(t, err)

		chains := []DiskChainConfig{
			{Config: &cfg1, Superchain: "test"},
			{Config: &cfg2, Superchain: "test"},
		}

		checker := NewDepsetChecker(lgr, chains, addrs)
		err = checker.checkOffchain(chains)
		require.Error(t, err)
		require.ErrorIs(t, err, errDepsetLengths)
	})
}

func TestDepsetChecker_checkOnchain(t *testing.T) {
	lgr := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, false))

	t.Run("no chain configs", func(t *testing.T) {
		checker := NewDepsetChecker(lgr, nil, nil)
		err := checker.checkOnchain(nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "no chain configs provided to checkOnchain")
	})

	t.Run("single chain config", func(t *testing.T) {
		addrs := loadAddresses(t, validAddressesPath)
		chain1 := loadChainConfig(t, "testdata/depsets_valid/chain1.toml")
		chains := []DiskChainConfig{
			{Config: chain1, Superchain: "test"},
		}

		checker := NewDepsetChecker(lgr, chains, addrs)
		require.NoError(t, checker.checkOnchain(chains))
	})

	t.Run("matching proxy addresses", func(t *testing.T) {
		addrs := loadAddresses(t, validAddressesPath)
		chain1 := loadChainConfig(t, "testdata/depsets_valid/chain1.toml")
		chain2 := loadChainConfig(t, "testdata/depsets_valid/chain2.toml")
		chains := []DiskChainConfig{
			{Config: chain1, Superchain: "test"},
			{Config: chain2, Superchain: "test"},
		}

		checker := NewDepsetChecker(lgr, chains, addrs)
		require.NoError(t, checker.checkOnchain(chains))
	})

	t.Run("mismatched proxy addresses", func(t *testing.T) {
		addrs := loadAddresses(t, invalidAddressesPath)
		chain1 := loadChainConfig(t, "testdata/depsets_valid/chain1.toml")
		chain2 := loadChainConfig(t, "testdata/depsets_valid/chain2.toml")
		chains := []DiskChainConfig{
			{Config: chain1, Superchain: "test"},
			{Config: chain2, Superchain: "test"},
		}

		checker := NewDepsetChecker(lgr, chains, addrs)
		err := checker.checkOnchain(chains)
		require.Error(t, err)
		require.Contains(t, err.Error(), "DisputeGameFactoryProxy address mismatch")
	})

	t.Run("missing proxy addresses", func(t *testing.T) {
		addrs := loadAddresses(t, invalidAddressesPath)
		chain1 := loadChainConfig(t, "testdata/depsets_invalid/missing_address_1.toml")
		chain2 := loadChainConfig(t, "testdata/depsets_invalid/missing_address_2.toml")
		chains := []DiskChainConfig{
			{Config: chain1, Superchain: "test"},
			{Config: chain2, Superchain: "test"},
		}

		checker := NewDepsetChecker(lgr, chains, addrs)
		err := checker.checkOnchain(chains)
		require.Error(t, err)
		require.ErrorIs(t, err, errMissingAddress)
	})

	t.Run("zero address", func(t *testing.T) {
		addrs := loadAddresses(t, invalidAddressesPath)
		chain1 := loadChainConfig(t, "testdata/depsets_valid/chain2.toml")
		chain2 := loadChainConfig(t, "testdata/depsets_valid/chain3.toml")
		chains := []DiskChainConfig{
			{Config: chain1, Superchain: "test"},
			{Config: chain2, Superchain: "test"},
		}

		checker := NewDepsetChecker(lgr, chains, addrs)
		err := checker.checkOnchain(chains)
		require.Error(t, err)
		require.ErrorIs(t, err, errMissingAddress)
	})
}
