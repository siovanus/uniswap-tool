package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Config object used by ontology-instance
type Config struct {
	JsonRpcAddress          string `json:"json_rpc_address"`
	ExchangeContractAddress string `json:"exchange_contract_address"`
	TokenContractAddress    string `json:"token_contract_address"`

	TokenAmount uint64 `json:"token_amount"`
	OntdAmount  uint64 `json:"ontd_amount"`
}

func NewConfig(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal Config:%s error:%s", data, err)
	}
	return cfg, nil
}
