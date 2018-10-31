package paracross

import (
	"gitlab.33.cn/chain33/chain33/plugin/dapp/paracross/commands"
	"gitlab.33.cn/chain33/chain33/plugin/dapp/paracross/executor"
	"gitlab.33.cn/chain33/chain33/plugin/dapp/paracross/rpc"
	ty "gitlab.33.cn/chain33/chain33/plugin/dapp/paracross/types"
	"gitlab.33.cn/chain33/chain33/pluginmgr"
	"gitlab.33.cn/chain33/chain33/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     "paracross",
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.ParcCmd,
		RPC:      rpc.Init,
	})
	types.RegisterDappFork(ty.ParaX, "Enable", 0)
}
