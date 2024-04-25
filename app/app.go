package app

import (
	"fmt"
	"io"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	tmlog "github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v7/packetforward"
	packetforwardkeeper "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v7/packetforward/keeper"
	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v7/packetforward/types"
	transfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v7/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	solomachine "github.com/cosmos/ibc-go/v7/modules/light-clients/06-solomachine"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	evmante "github.com/evmos/ethermint/app/ante"
	ethermintconfig "github.com/evmos/ethermint/server/config"
	"github.com/evmos/ethermint/x/evm"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/evmos/ethermint/x/feemarket"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"

	"github.com/0glabs/0g-chain/app/ante"
	chainparams "github.com/0glabs/0g-chain/app/params"
	"github.com/0glabs/0g-chain/chaincfg"
	"github.com/0glabs/0g-chain/precompiles/dasigners"

	evmutil "github.com/0glabs/0g-chain/x/evmutil"
	evmutilkeeper "github.com/0glabs/0g-chain/x/evmutil/keeper"
	evmutiltypes "github.com/0glabs/0g-chain/x/evmutil/types"

	committee "github.com/0glabs/0g-chain/x/committee/v1"
	committeekeeper "github.com/0glabs/0g-chain/x/committee/v1/keeper"
	committeetypes "github.com/0glabs/0g-chain/x/committee/v1/types"
	das "github.com/0glabs/0g-chain/x/das/v1"
	daskeeper "github.com/0glabs/0g-chain/x/das/v1/keeper"
	dastypes "github.com/0glabs/0g-chain/x/das/v1/types"
	validatorvesting "github.com/0glabs/0g-chain/x/validator-vesting"
	validatorvestingrest "github.com/0glabs/0g-chain/x/validator-vesting/client/rest"
	validatorvestingtypes "github.com/0glabs/0g-chain/x/validator-vesting/types"
)

var (
	// ModuleBasics manages simple versions of full app modules.
	// It's used for things such as codec registration and genesis file verification.
	ModuleBasics = module.NewBasicManager(
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic([]govclient.ProposalHandler{
			paramsclient.ProposalHandler,
			upgradeclient.LegacyProposalHandler,
			upgradeclient.LegacyCancelProposalHandler,
			ibcclientclient.UpdateClientProposalHandler,
			ibcclientclient.UpgradeProposalHandler,
		}),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ibctm.AppModuleBasic{},
		solomachine.AppModuleBasic{},
		packetforward.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		evm.AppModuleBasic{},
		feemarket.AppModuleBasic{},
		validatorvesting.AppModuleBasic{},
		evmutil.AppModuleBasic{},
		mint.AppModuleBasic{},
		consensus.AppModuleBasic{},
		committee.AppModuleBasic{},
		das.AppModuleBasic{},
	)

	// module account permissions
	// If these are changed, the permissions stored in accounts
	// must also be migrated during a chain upgrade.
	mAccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		evmtypes.ModuleName:            {authtypes.Minter, authtypes.Burner}, // used for secure addition and subtraction of balance using module account
		evmutiltypes.ModuleName:        {authtypes.Minter, authtypes.Burner},
		minttypes.ModuleName:           {authtypes.Minter},
	}
)

// Verify app interface at compile time
var (
	_ servertypes.Application = (*App)(nil)
)

// Options bundles several configuration params for an App.
type Options struct {
	SkipLoadLatest        bool
	SkipUpgradeHeights    map[int64]bool
	SkipGenesisInvariants bool
	InvariantCheckPeriod  uint
	MempoolEnableAuth     bool
	MempoolAuthAddresses  []sdk.AccAddress
	EVMTrace              string
	EVMMaxGasWanted       uint64
}

// DefaultOptions is a sensible default Options value.
var DefaultOptions = Options{
	EVMTrace:        ethermintconfig.DefaultEVMTracer,
	EVMMaxGasWanted: ethermintconfig.DefaultMaxTxGasWanted,
}

// App is the 0gChain ABCI application.
type App struct {
	*baseapp.BaseApp

	// codec
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers from all the modules
	accountKeeper         authkeeper.AccountKeeper
	bankKeeper            bankkeeper.Keeper
	capabilityKeeper      *capabilitykeeper.Keeper
	stakingKeeper         *stakingkeeper.Keeper
	distrKeeper           distrkeeper.Keeper
	govKeeper             govkeeper.Keeper
	paramsKeeper          paramskeeper.Keeper
	authzKeeper           authzkeeper.Keeper
	crisisKeeper          crisiskeeper.Keeper
	slashingKeeper        slashingkeeper.Keeper
	ibcKeeper             *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	packetForwardKeeper   *packetforwardkeeper.Keeper
	evmKeeper             *evmkeeper.Keeper
	evmutilKeeper         evmutilkeeper.Keeper
	feeMarketKeeper       feemarketkeeper.Keeper
	upgradeKeeper         upgradekeeper.Keeper
	evidenceKeeper        evidencekeeper.Keeper
	transferKeeper        ibctransferkeeper.Keeper
	mintKeeper            mintkeeper.Keeper
	consensusParamsKeeper consensusparamkeeper.Keeper
	CommitteeKeeper       committeekeeper.Keeper
	DasKeeper             daskeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// configurator
	configurator module.Configurator
}

// NewApp returns a reference to an initialized App.
func NewApp(
	logger tmlog.Logger,
	db dbm.DB,
	homePath string,
	traceStore io.Writer,
	encodingConfig chainparams.EncodingConfig,
	options Options,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(chaincfg.AppName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		packetforwardtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		ibcexported.StoreKey,
		upgradetypes.StoreKey,
		evidencetypes.StoreKey,
		ibctransfertypes.StoreKey,
		evmtypes.StoreKey,
		feemarkettypes.StoreKey,
		authzkeeper.StoreKey,
		capabilitytypes.StoreKey,
		evmutiltypes.StoreKey,
		minttypes.StoreKey,
		consensusparamtypes.StoreKey,
		crisistypes.StoreKey,
		committeetypes.StoreKey,
		dastypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey, evmtypes.TransientKey, feemarkettypes.TransientKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// Authority for gov proposals, using the x/gov module account address
	govAuthAddr := authtypes.NewModuleAddress(govtypes.ModuleName)
	govAuthAddrStr := govAuthAddr.String()

	app := &App{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	// init params keeper and subspaces
	app.paramsKeeper = paramskeeper.NewKeeper(
		appCodec,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)
	authSubspace := app.paramsKeeper.Subspace(authtypes.ModuleName)
	bankSubspace := app.paramsKeeper.Subspace(banktypes.ModuleName)
	stakingSubspace := app.paramsKeeper.Subspace(stakingtypes.ModuleName)
	distrSubspace := app.paramsKeeper.Subspace(distrtypes.ModuleName)
	slashingSubspace := app.paramsKeeper.Subspace(slashingtypes.ModuleName)
	govSubspace := app.paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable())
	crisisSubspace := app.paramsKeeper.Subspace(crisistypes.ModuleName)
	ibcSubspace := app.paramsKeeper.Subspace(ibcexported.ModuleName)
	ibctransferSubspace := app.paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	packetforwardSubspace := app.paramsKeeper.Subspace(packetforwardtypes.ModuleName).WithKeyTable(packetforwardtypes.ParamKeyTable())
	feemarketSubspace := app.paramsKeeper.Subspace(feemarkettypes.ModuleName)
	evmSubspace := app.paramsKeeper.Subspace(evmtypes.ModuleName)
	evmutilSubspace := app.paramsKeeper.Subspace(evmutiltypes.ModuleName)
	mintSubspace := app.paramsKeeper.Subspace(minttypes.ModuleName)

	// set the BaseApp's parameter store
	app.consensusParamsKeeper = consensusparamkeeper.NewKeeper(appCodec, keys[consensusparamtypes.StoreKey], govAuthAddrStr)
	bApp.SetParamStore(&app.consensusParamsKeeper)

	app.capabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	scopedIBCKeeper := app.capabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	scopedTransferKeeper := app.capabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	app.capabilityKeeper.Seal()

	// add keepers
	app.accountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		mAccPerms,
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
		govAuthAddrStr,
	)
	app.bankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.accountKeeper,
		app.loadBlockedMaccAddrs(),
		govAuthAddrStr,
	)
	app.stakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		govAuthAddrStr,
	)
	app.authzKeeper = authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey],
		appCodec,
		app.BaseApp.MsgServiceRouter(),
		app.accountKeeper,
	)
	app.distrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		app.stakingKeeper,
		authtypes.FeeCollectorName,
		govAuthAddrStr,
	)
	app.slashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		app.legacyAmino,
		keys[slashingtypes.StoreKey],
		app.stakingKeeper,
		govAuthAddrStr,
	)
	app.crisisKeeper = *crisiskeeper.NewKeeper(
		appCodec,
		keys[crisistypes.StoreKey],
		options.InvariantCheckPeriod,
		app.bankKeeper,
		authtypes.FeeCollectorName,
		govAuthAddrStr,
	)
	app.upgradeKeeper = *upgradekeeper.NewKeeper(
		options.SkipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
		govAuthAddrStr,
	)
	app.evidenceKeeper = *evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		app.stakingKeeper,
		app.slashingKeeper,
	)

	app.ibcKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibcexported.StoreKey],
		ibcSubspace,
		app.stakingKeeper,
		app.upgradeKeeper,
		scopedIBCKeeper,
	)

	// Create Ethermint keepers
	app.feeMarketKeeper = feemarketkeeper.NewKeeper(
		appCodec,
		govAuthAddr,
		keys[feemarkettypes.StoreKey],
		tkeys[feemarkettypes.TransientKey],
		feemarketSubspace,
	)

	app.evmutilKeeper = evmutilkeeper.NewKeeper(
		app.appCodec,
		keys[evmutiltypes.StoreKey],
		evmutilSubspace,
		app.bankKeeper,
		app.accountKeeper,
	)

	evmBankKeeper := evmutilkeeper.NewEvmBankKeeper(app.evmutilKeeper, app.bankKeeper, app.accountKeeper)
	// precopmiles
	precompiles := make(map[common.Address]vm.PrecompiledContract)
	daSignersPrecompile, err := dasigners.NewDASignersPrecompile()
	if err != nil {
		panic("initialize precompile failed")
	}
	precompiles[daSignersPrecompile.Address()] = daSignersPrecompile
	// evm keeper
	app.evmKeeper = evmkeeper.NewKeeper(
		appCodec, keys[evmtypes.StoreKey], tkeys[evmtypes.TransientKey],
		govAuthAddr,
		app.accountKeeper, evmBankKeeper, app.stakingKeeper, app.feeMarketKeeper,
		options.EVMTrace,
		evmSubspace,
		precompiles,
	)
	fmt.Println(app.evmKeeper.GetPrecompiles())

	app.evmutilKeeper.SetEvmKeeper(app.evmKeeper)

	// It's important to note that the PFM Keeper must be initialized before the Transfer Keeper
	app.packetForwardKeeper = packetforwardkeeper.NewKeeper(
		appCodec,
		keys[packetforwardtypes.StoreKey],
		nil, // will be zero-value here, reference is set later on with SetTransferKeeper.
		app.ibcKeeper.ChannelKeeper,
		app.distrKeeper,
		app.bankKeeper,
		app.ibcKeeper.ChannelKeeper,
		govAuthAddrStr,
	)

	app.transferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		ibctransferSubspace,
		app.packetForwardKeeper,
		app.ibcKeeper.ChannelKeeper,
		&app.ibcKeeper.PortKeeper,
		app.accountKeeper,
		app.bankKeeper,
		scopedTransferKeeper,
	)
	app.packetForwardKeeper.SetTransferKeeper(app.transferKeeper)
	transferModule := transfer.NewAppModule(app.transferKeeper)

	// allow ibc packet forwarding for ibc transfers.
	// transfer stack contains (from top to bottom):
	// - Packet Forward Middleware
	// - Transfer
	var transferStack ibcporttypes.IBCModule
	transferStack = transfer.NewIBCModule(app.transferKeeper)
	transferStack = packetforward.NewIBCMiddleware(
		transferStack,
		app.packetForwardKeeper,
		0, // retries on timeout
		packetforwardkeeper.DefaultForwardTransferPacketTimeoutTimestamp,
		packetforwardkeeper.DefaultRefundTransferPacketTimeoutTimestamp,
	)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferStack)
	app.ibcKeeper.SetRouter(ibcRouter)

	app.mintKeeper = mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		app.stakingKeeper,
		app.accountKeeper,
		app.bankKeeper,
		authtypes.FeeCollectorName,
		govAuthAddrStr,
	)

	// register the staking hooks
	app.stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			app.distrKeeper.Hooks(),
			app.slashingKeeper.Hooks(),
		))

	// create gov keeper with router
	// NOTE this must be done after any keepers referenced in the gov router (ie committee) are defined
	govRouter := govv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.paramsKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(&app.upgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.ibcKeeper.ClientKeeper))

	govConfig := govtypes.DefaultConfig()
	govKeeper := govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		app.stakingKeeper,
		app.MsgServiceRouter(),
		govConfig,
		govAuthAddrStr,
	)
	govKeeper.SetLegacyRouter(govRouter)
	app.govKeeper = *govKeeper

	app.CommitteeKeeper = committeekeeper.NewKeeper(
		keys[committeetypes.StoreKey], appCodec, app.stakingKeeper,
	)
	app.DasKeeper = daskeeper.NewKeeper(keys[dastypes.StoreKey], appCodec, app.stakingKeeper)

	// create the module manager (Note: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.)
	app.mm = module.NewManager(
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(appCodec, app.accountKeeper, authsims.RandomGenesisAccounts, authSubspace),
		bank.NewAppModule(appCodec, app.bankKeeper, app.accountKeeper, bankSubspace),
		capability.NewAppModule(appCodec, *app.capabilityKeeper, false), // todo: confirm if this is okay to not be sealed
		staking.NewAppModule(appCodec, app.stakingKeeper, app.accountKeeper, app.bankKeeper, stakingSubspace),
		distr.NewAppModule(appCodec, app.distrKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper, distrSubspace),
		gov.NewAppModule(appCodec, &app.govKeeper, app.accountKeeper, app.bankKeeper, govSubspace),
		params.NewAppModule(app.paramsKeeper),
		crisis.NewAppModule(&app.crisisKeeper, options.SkipGenesisInvariants, crisisSubspace),
		slashing.NewAppModule(appCodec, app.slashingKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper, slashingSubspace),
		consensus.NewAppModule(appCodec, app.consensusParamsKeeper),
		ibc.NewAppModule(app.ibcKeeper),
		packetforward.NewAppModule(app.packetForwardKeeper, packetforwardSubspace),
		evm.NewAppModule(app.evmKeeper, app.accountKeeper),
		feemarket.NewAppModule(app.feeMarketKeeper, feemarketSubspace),
		upgrade.NewAppModule(&app.upgradeKeeper),
		evidence.NewAppModule(app.evidenceKeeper),
		transferModule,
		vesting.NewAppModule(app.accountKeeper, app.bankKeeper),
		authzmodule.NewAppModule(appCodec, app.authzKeeper, app.accountKeeper, app.bankKeeper, app.interfaceRegistry),
		validatorvesting.NewAppModule(app.bankKeeper),
		evmutil.NewAppModule(app.evmutilKeeper, app.bankKeeper, app.accountKeeper),
		mint.NewAppModule(appCodec, app.mintKeeper, app.accountKeeper, nil, mintSubspace),
	)

	// Warning: Some begin blockers must run before others. Ensure the dependencies are understood before modifying this list.
	app.mm.SetOrderBeginBlockers(
		// Upgrade begin blocker runs migrations on the first block after an upgrade. It should run before any other module.
		upgradetypes.ModuleName,
		// Capability begin blocker runs non state changing initialization.
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		// During begin block slashing happens after distr.BeginBlocker so that
		// there is nothing left over in the validator fee pool, so as to keep the
		// CanWithdrawInvariant invariant.
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		feemarkettypes.ModuleName,
		evmtypes.ModuleName,
		ibcexported.ModuleName,
		// Add all remaining modules with an empty begin blocker below since cosmos 0.45.0 requires it
		vestingtypes.ModuleName,
		validatorvestingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		ibctransfertypes.ModuleName,
		paramstypes.ModuleName,
		authz.ModuleName,
		evmutiltypes.ModuleName,
		consensusparamtypes.ModuleName,
		packetforwardtypes.ModuleName,

		committeetypes.ModuleName,
		dastypes.ModuleName,
	)

	// Warning: Some end blockers must run before others. Ensure the dependencies are understood before modifying this list.
	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		evmtypes.ModuleName,
		// fee market module must go after evm module in order to retrieve the block gas used.
		feemarkettypes.ModuleName,
		// Add all remaining modules with an empty end blocker below since cosmos 0.45.0 requires it
		capabilitytypes.ModuleName,
		slashingtypes.ModuleName,
		distrtypes.ModuleName,
		upgradetypes.ModuleName,
		evidencetypes.ModuleName,
		vestingtypes.ModuleName,
		ibcexported.ModuleName,
		validatorvestingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		genutiltypes.ModuleName,
		ibctransfertypes.ModuleName,
		paramstypes.ModuleName,
		authz.ModuleName,
		evmutiltypes.ModuleName,
		minttypes.ModuleName,
		consensusparamtypes.ModuleName,
		packetforwardtypes.ModuleName,
		committeetypes.ModuleName,
		dastypes.ModuleName,
	)

	// Warning: Some init genesis methods must run before others. Ensure the dependencies are understood before modifying this list
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName, // initialize capabilities, run before any module creating or claiming capabilities in InitGenesis
		authtypes.ModuleName,       // loads all accounts, run before any module with a module account
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName, // iterates over validators, run after staking
		govtypes.ModuleName,
		minttypes.ModuleName,
		ibcexported.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		ibctransfertypes.ModuleName,
		evmtypes.ModuleName,
		feemarkettypes.ModuleName,
		evmutiltypes.ModuleName,
		genutiltypes.ModuleName, // runs arbitrary txs included in genisis state, so run after modules have been initialized
		// Add all remaining modules with an empty InitGenesis below since cosmos 0.45.0 requires it
		vestingtypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		validatorvestingtypes.ModuleName,
		consensusparamtypes.ModuleName,
		packetforwardtypes.ModuleName,
		crisistypes.ModuleName, // runs the invariants at genesis, should run after other modules
		committeetypes.ModuleName,
		dastypes.ModuleName,
	)

	app.mm.RegisterInvariants(&app.crisisKeeper)

	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.RegisterServices(app.configurator)

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	// It needs to be called after `app.mm` and `app.configurator` are set.
	app.RegisterUpgradeHandlers()

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: This is not required for apps that don't use the simulator for fuzz testing
	// transactions.
	// TODO
	// app.sm = module.NewSimulationManager(
	// 	auth.NewAppModule(app.accountKeeper),
	// 	bank.NewAppModule(app.bankKeeper, app.accountKeeper),
	// 	gov.NewAppModule(app.govKeeper, app.accountKeeper, app.accountKeeper, app.bankKeeper),
	// 	mint.NewAppModule(app.mintKeeper),
	// 	distr.NewAppModule(app.distrKeeper, app.accountKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
	//  staking.NewAppModule(app.stakingKeeper, app.accountKeeper, app.accountKeeper, app.bankKeeper),
	//  evm.NewAppModule(app.evmKeeper, app.accountKeeper),
	// 	slashing.NewAppModule(app.slashingKeeper, app.accountKeeper, app.stakingKeeper),
	// )
	// app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize the app
	var fetchers []ante.AddressFetcher
	if options.MempoolEnableAuth {
		fetchers = append(fetchers,
			func(sdk.Context) []sdk.AccAddress { return options.MempoolAuthAddresses },
		)
	}

	anteOptions := ante.HandlerOptions{
		AccountKeeper:          app.accountKeeper,
		BankKeeper:             app.bankKeeper,
		EvmKeeper:              app.evmKeeper,
		IBCKeeper:              app.ibcKeeper,
		FeeMarketKeeper:        app.feeMarketKeeper,
		SignModeHandler:        encodingConfig.TxConfig.SignModeHandler(),
		SigGasConsumer:         evmante.DefaultSigVerificationGasConsumer,
		MaxTxGasWanted:         options.EVMMaxGasWanted,
		AddressFetchers:        fetchers,
		ExtensionOptionChecker: nil,
		TxFeeChecker:           nil,
	}

	antehandler, err := ante.NewAnteHandler(anteOptions)
	if err != nil {
		panic(fmt.Sprintf("failed to create antehandler: %s", err))
	}

	app.SetAnteHandler(antehandler)
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// load store
	if !options.SkipLoadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			panic(fmt.Sprintf("failed to load latest version: %s", err))
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper

	return app
}

func (app *App) RegisterServices(cfg module.Configurator) {
	app.mm.RegisterServices(cfg)
}

// BeginBlocker contains app specific logic for the BeginBlock abci call.
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker contains app specific logic for the EndBlock abci call.
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer contains app specific logic for the InitChain abci call.
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	// Store current module versions in 0gChain-10 to setup future in-place upgrades.
	// During in-place migrations, the old module versions in the store will be referenced to determine which migrations to run.
	app.upgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads the app state for a particular height.
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range mAccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// InterfaceRegistry returns the app's InterfaceRegistry.
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// SimulationManager implements the SimulationApp interface.
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx

	// Register custom REST routes
	validatorvestingrest.RegisterRoutes(clientCtx, apiSvr.Router)

	// Register GRPC Gateway routes
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Swagger API configuration is ignored
}

// RegisterTxService implements the Application.RegisterTxService method.
// It registers transaction related endpoints on the app's grpc server.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
// It registers the standard tendermint grpc endpoints on the app's grpc server.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

func (app *App) RegisterNodeService(clientCtx client.Context) {
	nodeservice.RegisterNodeService(clientCtx, app.BaseApp.GRPCQueryRouter())
}

// loadBlockedMaccAddrs returns a map indicating the blocked status of each module account address
func (app *App) loadBlockedMaccAddrs() map[string]bool {
	modAccAddrs := app.ModuleAccountAddrs()
	allowedMaccs := map[string]bool{}

	for addr := range modAccAddrs {
		// Set allowed module accounts as unblocked
		if allowedMaccs[addr] {
			modAccAddrs[addr] = false
		}
	}
	return modAccAddrs
}

// GetMaccPerms returns a mapping of the application's module account permissions.
func GetMaccPerms() map[string][]string {
	perms := make(map[string][]string)
	for k, v := range mAccPerms {
		perms[k] = v
	}
	return perms
}
