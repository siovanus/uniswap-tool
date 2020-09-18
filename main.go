package main

import (
	"fmt"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/password"
	"github.com/siovanus/uniswap-tool/config"
	"time"
)

func main() {
	config, err := config.NewConfig("config.json")
	if err != nil {
		fmt.Println("parse config failed, err:", err)
		return
	}

	sdk := ontology_go_sdk.NewOntologySdk()
	sdk.NewRpcClient().SetAddress(config.JsonRpcAddress)
	account, _ := GetAccountByPassword(sdk, "wallet.dat")
	exchangeAddress, err := common.AddressFromHexString(config.ExchangeContractAddress)
	if err != nil {
		fmt.Println(err)
	}
	tokenAddress, err := common.AddressFromHexString(config.TokenContractAddress)
	if err != nil {
		fmt.Println(err)
	}

	txhash1, err := sdk.NeoVM.InvokeNeoVMContract(2500, 20000, account, account, tokenAddress,
		[]interface{}{"approve", []interface{}{account.Address, exchangeAddress, config.TokenAmount}})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(txhash1.ToHexString())
	}
	time.Sleep(10 * time.Second)

	txhash2, err := sdk.NeoVM.InvokeNeoVMContract(2500, 200000, account, account, exchangeAddress,
		[]interface{}{"addLiquidity", []interface{}{1, config.TokenAmount, 1634450069, account.Address[:], config.OntdAmount}})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(txhash2.ToHexString())
	}
}

func GetAccountByPassword(sdk *ontology_go_sdk.OntologySdk, path string) (*ontology_go_sdk.Account, bool) {
	wallet, err := sdk.OpenWallet(path)
	if err != nil {
		fmt.Println("open wallet error:", err)
		return nil, false
	}
	pwd, err := password.GetPassword()
	if err != nil {
		fmt.Println("getPassword error:", err)
		return nil, false
	}
	user, err := wallet.GetDefaultAccount(pwd)
	if err != nil {
		fmt.Println("getDefaultAccount error:", err)
		return nil, false
	}
	return user, true
}
