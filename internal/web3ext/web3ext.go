// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// package web3ext contains getrue specific web3.js extensions.
package web3ext

var Modules = map[string]string{
	"admin":      Admin_JS,
	"chequebook": Chequebook_JS,
	"clique":     Clique_JS,
	"debug":      Debug_JS,
	"etrue":      Etrue_JS,
	"net":        Net_JS,
	"personal":   Personal_JS,
	"rpc":        RPC_JS,
	"shh":        Shh_JS,
	"swarmfs":    SWARMFS_JS,
	"txpool":     TxPool_JS,
	"fruitpool":  FruitPool_JS,
	"impawn":     Impawn_JS,
}

const Chequebook_JS = `
web3._extend({
	property: 'chequebook',
	methods: [
		new web3._extend.Method({
			name: 'deposit',
			call: 'chequebook_deposit',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Property({
			name: 'balance',
			getter: 'chequebook_balance',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.Method({
			name: 'cash',
			call: 'chequebook_cash',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'issue',
			call: 'chequebook_issue',
			params: 2,
			inputFormatter: [null, null]
		}),
	]
});
`

const Clique_JS = `
web3._extend({
	property: 'clique',
	methods: [
		new web3._extend.Method({
			name: 'getSnapshot',
			call: 'clique_getSnapshot',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getSnapshotAtHash',
			call: 'clique_getSnapshotAtHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getSigners',
			call: 'clique_getSigners',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getSignersAtHash',
			call: 'clique_getSignersAtHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'propose',
			call: 'clique_propose',
			params: 2
		}),
		new web3._extend.Method({
			name: 'discard',
			call: 'clique_discard',
			params: 1
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'proposals',
			getter: 'clique_proposals'
		}),
	]
});
`

const Admin_JS = `
web3._extend({
	property: 'admin',
	methods: [
		new web3._extend.Method({
			name: 'addPeer',
			call: 'admin_addPeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'removePeer',
			call: 'admin_removePeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'exportChain',
			call: 'admin_exportChain',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'importChain',
			call: 'admin_importChain',
			params: 1
		}),
		new web3._extend.Method({
			name: 'sleepBlocks',
			call: 'admin_sleepBlocks',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startRPC',
			call: 'admin_startRPC',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.Method({
			name: 'stopRPC',
			call: 'admin_stopRPC'
		}),
		new web3._extend.Method({
			name: 'startWS',
			call: 'admin_startWS',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.Method({
			name: 'stopWS',
			call: 'admin_stopWS'
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'nodeInfo',
			getter: 'admin_nodeInfo'
		}),
		new web3._extend.Property({
			name: 'peers',
			getter: 'admin_peers'
		}),
		new web3._extend.Property({
			name: 'datadir',
			getter: 'admin_datadir'
		}),
	]
});
`

const Debug_JS = `
web3._extend({
	property: 'debug',
	methods: [
		new web3._extend.Method({
			name: 'printBlock',
			call: 'debug_printBlock',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getBlockRlp',
			call: 'debug_getBlockRlp',
			params: 1
		}),
		new web3._extend.Method({
			name: 'setHead',
			call: 'debug_setHead',
			params: 1
		}),
		new web3._extend.Method({
			name: 'seedHash',
			call: 'debug_seedHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'dumpBlock',
			call: 'debug_dumpBlock',
			params: 1
		}),
		new web3._extend.Method({
			name: 'chaindbProperty',
			call: 'debug_chaindbProperty',
			params: 1,
			outputFormatter: console.log
		}),
		new web3._extend.Method({
			name: 'chaindbCompact',
			call: 'debug_chaindbCompact',
		}),
		new web3._extend.Method({
			name: 'metrics',
			call: 'debug_metrics',
			params: 1
		}),
		new web3._extend.Method({
			name: 'verbosity',
			call: 'debug_verbosity',
			params: 1
		}),
		new web3._extend.Method({
			name: 'vmodule',
			call: 'debug_vmodule',
			params: 1
		}),
		new web3._extend.Method({
			name: 'backtraceAt',
			call: 'debug_backtraceAt',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'stacks',
			call: 'debug_stacks',
			params: 0,
			outputFormatter: console.log
		}),
		new web3._extend.Method({
			name: 'freeOSMemory',
			call: 'debug_freeOSMemory',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'setGCPercent',
			call: 'debug_setGCPercent',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'memStats',
			call: 'debug_memStats',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'gcStats',
			call: 'debug_gcStats',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'cpuProfile',
			call: 'debug_cpuProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startCPUProfile',
			call: 'debug_startCPUProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'stopCPUProfile',
			call: 'debug_stopCPUProfile',
			params: 0
		}),
		new web3._extend.Method({
			name: 'goTrace',
			call: 'debug_goTrace',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startGoTrace',
			call: 'debug_startGoTrace',
			params: 1
		}),
		new web3._extend.Method({
			name: 'stopGoTrace',
			call: 'debug_stopGoTrace',
			params: 0
		}),
		new web3._extend.Method({
			name: 'blockProfile',
			call: 'debug_blockProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'setBlockProfileRate',
			call: 'debug_setBlockProfileRate',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeBlockProfile',
			call: 'debug_writeBlockProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'mutexProfile',
			call: 'debug_mutexProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'setMutexProfileFraction',
			call: 'debug_setMutexProfileFraction',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeMutexProfile',
			call: 'debug_writeMutexProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeMemProfile',
			call: 'debug_writeMemProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'traceBlock',
			call: 'debug_traceBlock',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockFromFile',
			call: 'debug_traceBlockFromFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockByNumber',
			call: 'debug_traceBlockByNumber',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockByHash',
			call: 'debug_traceBlockByHash',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceTransaction',
			call: 'debug_traceTransaction',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'preimage',
			call: 'debug_preimage',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getBadBlocks',
			call: 'debug_getBadBlocks',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'storageRangeAt',
			call: 'debug_storageRangeAt',
			params: 5,
		}),
		new web3._extend.Method({
			name: 'getModifiedAccountsByNumber',
			call: 'debug_getModifiedAccountsByNumber',
			params: 2,
			inputFormatter: [null, null],
		}),
		new web3._extend.Method({
			name: 'getModifiedAccountsByHash',
			call: 'debug_getModifiedAccountsByHash',
			params: 2,
			inputFormatter:[null, null],
		}),
	],
	properties: []
});
`

const Etrue_JS = `
web3._extend({
	property: 'etrue',
	methods: [
		new web3._extend.Method({
			name: 'chainId',
			call: 'etrue_chainId',
			params: 0
		}),
		new web3._extend.Method({
			name: 'sign',
			call: 'etrue_sign',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, null]
		}),
		new web3._extend.Method({
			name: 'resend',
			call: 'etrue_resend',
			params: 3,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter, web3._extend.utils.fromDecimal, web3._extend.utils.fromDecimal]
		}),
		new web3._extend.Method({
			name: 'signTransaction',
			call: 'etrue_signTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.Method({
			name: 'submitTransaction',
			call: 'etrue_submitTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.Method({
			name: 'getRawTransaction',
			call: 'etrue_getRawTransactionByHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getRawTransactionFromBlock',
			call: function(args) {
				return (web3._extend.utils.isString(args[0]) && args[0].indexOf('0x') === 0) ? 'etrue_getRawTransactionByBlockHashAndIndex' : 'etrue_getRawTransactionByBlockNumberAndIndex';
			},
			params: 2,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter, web3._extend.utils.toHex]
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'pendingTransactions',
			getter: 'etrue_pendingTransactions',
			outputFormatter: function(txs) {
				var formatted = [];
				for (var i = 0; i < txs.length; i++) {
					formatted.push(web3._extend.formatters.outputTransactionFormatter(txs[i]));
					formatted[i].blockHash = null;
				}
				return formatted;
			}
		}),
	]
});
`

const Net_JS = `
web3._extend({
	property: 'net',
	methods: [],
	properties: [
		new web3._extend.Property({
			name: 'version',
			getter: 'net_version'
		}),
	]
});
`

const Personal_JS = `
web3._extend({
	property: 'personal',
	methods: [
		new web3._extend.Method({
			name: 'importRawKey',
			call: 'personal_importRawKey',
			params: 2
		}),
		new web3._extend.Method({
			name: 'sign',
			call: 'personal_sign',
			params: 3,
			inputFormatter: [null, web3._extend.formatters.inputAddressFormatter, null]
		}),
		new web3._extend.Method({
			name: 'ecRecover',
			call: 'personal_ecRecover',
			params: 2
		}),
		new web3._extend.Method({
			name: 'openWallet',
			call: 'personal_openWallet',
			params: 2
		}),
		new web3._extend.Method({
			name: 'deriveAccount',
			call: 'personal_deriveAccount',
			params: 3
		}),
		new web3._extend.Method({
			name: 'signTransaction',
			call: 'personal_signTransaction',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter, null]
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'listWallets',
			getter: 'personal_listWallets'
		}),
	]
})
`

const RPC_JS = `
web3._extend({
	property: 'rpc',
	methods: [],
	properties: [
		new web3._extend.Property({
			name: 'modules',
			getter: 'rpc_modules'
		}),
	]
});
`

const Shh_JS = `
web3._extend({
	property: 'shh',
	methods: [
	],
	properties:
	[
		new web3._extend.Property({
			name: 'version',
			getter: 'shh_version',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.Property({
			name: 'info',
			getter: 'shh_info'
		}),
	]
});
`

const SWARMFS_JS = `
web3._extend({
	property: 'swarmfs',
	methods:
	[
		new web3._extend.Method({
			name: 'mount',
			call: 'swarmfs_mount',
			params: 2
		}),
		new web3._extend.Method({
			name: 'unmount',
			call: 'swarmfs_unmount',
			params: 1
		}),
		new web3._extend.Method({
			name: 'listmounts',
			call: 'swarmfs_listmounts',
			params: 0
		}),
	]
});
`

const TxPool_JS = `
web3._extend({
	property: 'txpool',
	methods: [],
	properties:
	[
		new web3._extend.Property({
			name: 'content',
			getter: 'txpool_content'
		}),
		new web3._extend.Property({
			name: 'inspect',
			getter: 'txpool_inspect'
		}),
		new web3._extend.Property({
			name: 'status',
			getter: 'txpool_status',
			outputFormatter: function(status) {
				status.pending = web3._extend.utils.toDecimal(status.pending);
				status.queued = web3._extend.utils.toDecimal(status.queued);
				return status;
			}
		}),
	]
});
`

const FruitPool_JS = `
web3._extend({
	property: 'fruitpool',
	methods: [],
	properties:
	[
		new web3._extend.Property({
			name: 'content',
			getter: 'fruitpool_content'
		}),
		new web3._extend.Property({
			name: 'inspect',
			getter: 'fruitpool_inspect'
		}),
		new web3._extend.Property({
			name: 'status',
			getter: 'fruitpool_status'
		}),
	]
});
`

const Impawn_JS = `
web3._extend({
	property: 'impawn',
	methods: [
		new web3._extend.Method({
			name: 'getAllStakingAccount',
			call: 'impawn_getAllStakingAccount',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputDefaultBlockNumberFormatter],
			outputFormatter: function(sas) {
				if (sas !== null) {
					for (var i = 0; i < sas.stakers.length; i++) {
						var sa = sas.stakers[i];
						if(sa.unit.value !== null) {
							for (var j = 0; j < sa.unit.value.length; j++) {
								sa.unit.value[j].amount = web3._extend.utils.toBigNumber(sa.unit.value[j].amount);
								sa.unit.value[j].height = web3._extend.utils.toBigNumber(sa.unit.value[j].height);
							}						
						}
						if(sa.unit.redeemInfo !== null) {
							for (var j = 0; j < sa.unit.redeemInfo.length; j++) {
								sa.unit.redeemInfo[j].amount = web3._extend.utils.toBigNumber(sa.unit.redeemInfo[j].amount);
							}						
						}
						if(sa.delegation !== null) {
							for (var m = 0; m < sa.delegation.length; m++) {
								if(sa.delegation[m].unit.value !== null) {
									for (var j = 0; j < sa.delegation[m].unit.value.length; j++) {
										sa.delegation[m].unit.value[j].amount = web3._extend.utils.toBigNumber(sa.delegation[m].unit.value[j].amount);
										sa.delegation[m].unit.value[j].height = web3._extend.utils.toBigNumber(sa.delegation[m].unit.value[j].height);
									}						
								}
								if(sa.delegation[m].unit.redeemInfo !== null) {
									for (var j = 0; j < sa.delegation[m].unit.redeemInfo.length; j++) {
										sa.delegation[m].unit.redeemInfo[j].amount = web3._extend.utils.toBigNumber(sa.delegation[m].unit.redeemInfo[j].amount);
									}						
								}
							}						
						}
					}
				}
				return sas;
			}
		}),
		new web3._extend.Method({
			name: 'getStakingAsset',
			call: 'impawn_getStakingAsset',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter,web3._extend.formatters.inputDefaultBlockNumberFormatter],
			outputFormatter: function(sas) {
				var formatted = [];
				for (var i = 0; i < sas.length; i++) {
					if(sas[i].stakingValue !== null) {
						for (var j = 0; j < sas[i].stakingValue.length; j++) {
							sas[i].stakingValue[j].amount = web3._extend.utils.toBigNumber(sas[i].stakingValue[j].amount);
							sas[i].stakingValue[j].height = web3._extend.utils.toDecimal(sas[i].stakingValue[j].height);
						}						
					}
					formatted.push(sas[i]);
				}
				return formatted;
			}
		}),
		new web3._extend.Method({
			name: 'getLockedAsset',
			call: 'impawn_getLockedAsset',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter,web3._extend.formatters.inputDefaultBlockNumberFormatter],
			outputFormatter: function(las) {
				var formatted = [];
				for (var i = 0; i < las.length; i++) {
					if(las[i].lockValue !== null) {
						for (var j = 0; j < las[i].lockValue.length; j++) {
							las[i].lockValue[j].amount = web3._extend.utils.toBigNumber(las[i].lockValue[j].amount);
							las[i].lockValue[j].epochID = web3._extend.utils.toDecimal(las[i].lockValue[j].epochID);
							las[i].lockValue[j].height = web3._extend.utils.toBigNumber(las[i].lockValue[j].height);
						}
					}
					formatted.push(las[i]);
				}
				return formatted;
			}
		}),
		new web3._extend.Method({
			name: 'getAllCancelableAsset',
			call: 'impawn_getAllCancelableAsset',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter,web3._extend.formatters.inputDefaultBlockNumberFormatter],
			outputFormatter: function(cas) {
				var formatted = [];
				for (var i = 0; i < cas.length; i++) {
					cas[i].value = web3._extend.utils.toBigNumber(cas[i].value)
					formatted.push(cas[i]);
				}
				return formatted;
			}
		}),
		new web3._extend.Method({
			name: 'getStakingAccount',
			call: 'impawn_getStakingAccount',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter,web3._extend.formatters.inputDefaultBlockNumberFormatter],
			outputFormatter: function(sa) {
				if (sa.unit != null) {
					if(sa.unit.value !== null) {
						for (var j = 0; j < sa.unit.value.length; j++) {
							sa.unit.value[j].amount = web3._extend.utils.toBigNumber(sa.unit.value[j].amount);
							sa.unit.value[j].height = web3._extend.utils.toBigNumber(sa.unit.value[j].height);
						}						
					}
					if(sa.unit.redeemInfo !== null) {
						for (var j = 0; j < sa.unit.redeemInfo.length; j++) {
							sa.unit.redeemInfo[j].amount = web3._extend.utils.toBigNumber(sa.unit.redeemInfo[j].amount);
						}						
					}
				}
				if(sa.delegation !== null) {
					for (var m = 0; m < sa.delegation.length; m++) {
						if(sa.delegation[m].unit.value !== null) {
							for (var j = 0; j < sa.delegation[m].unit.value.length; j++) {
								sa.delegation[m].unit.value[j].amount = web3._extend.utils.toBigNumber(sa.delegation[m].unit.value[j].amount);
								sa.delegation[m].unit.value[j].height = web3._extend.utils.toBigNumber(sa.delegation[m].unit.value[j].height);
							}						
						}
						if(sa.delegation[m].unit.redeemInfo !== null) {
							for (var j = 0; j < sa.delegation[m].unit.redeemInfo.length; j++) {
								sa.delegation[m].unit.redeemInfo[j].amount = web3._extend.utils.toBigNumber(sa.delegation[m].unit.redeemInfo[j].amount);
							}						
						}
					}						
				}
				return sa;
			}
		}),
		new web3._extend.Method({
			name: 'getImpawnSummay',
			call: 'impawn_getImpawnSummay',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputDefaultBlockNumberFormatter],
			outputFormatter: function(infos) {
				if(infos.currentAllStaking != null) {
					infos.currentAllStaking = web3._extend.utils.toBigNumber(infos.currentAllStaking);
				}
				for (var i = 0;i < infos.EpochInfos.length;i++) {
					if (infos.EpochInfos[i].AllAmount != null) {
						infos.EpochInfos[i].AllAmount = web3._extend.utils.toBigNumber(infos.EpochInfos[i].AllAmount);
					}
				}
				return infos;
			}
		}),
	]
});
`
