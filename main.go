package main

import (
	"fmt"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/password"
	"time"
)

func main() {
	sdk := ontology_go_sdk.NewOntologySdk()
	sdk.NewRpcClient().SetAddress("http://polaris2.ont.io:20336")
	account, _ := GetAccountByPassword(sdk, "wallet.dat")
	contractAddress, err := common.AddressFromHexString("1b5ed9d9cee8e7ae90515fd76b166e203a4ef6d9")
	if err != nil {
		fmt.Println(err)
	}
	tokenAddress, err := common.AddressFromHexString("57ed5666dafabcd4fc69082b4e5c503105ec85d4")
	if err != nil {
		fmt.Println(err)
	}

	txhash1, err := sdk.NeoVM.InvokeNeoVMContract(2500, 20000, account, account, tokenAddress,
		[]interface{}{"approve", []interface{}{account.Address, contractAddress, 81000000000000000}})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(txhash1.ToHexString())
	}
    time.Sleep(30*time.Second)

	txhash2, err := sdk.NeoVM.InvokeNeoVMContract(2500, 200000, account, account, contractAddress,
		[]interface{}{"addLiquidity", []interface{}{1, 81000000000000000, 1634450069, account.Address[:], 100000000000}})
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