package replicant

import (
	"github.com/OperatorFoundation/Replicant-go/Replicant/v3/toneburst"
)

type ClientJsonConfig struct {
	ServerAddress string    				     `json:"serverAddress"`
	Toneburst     toneburst.StarburstConfig      `json:"toneburst"`    
	Polish        DarkStarPolishClientJsonConfig `json:"polish"`       
	Transport     string    		 			 `json:"transport"`
}

type ServerJsonConfig struct {
	ServerAddress string    				     `json:"serverAddress"`
	Toneburst     toneburst.StarburstConfig		 `json:"toneburst"`    
	Polish        DarkStarPolishServerJsonConfig `json:"polish"`       
	Transport     string    		  			 `json:"transport"`
	BindAddress	  *string						 `json:"bindAddress"`
}

type DarkStarPolishServerJsonConfig struct {
	ServerAddress    string `json:"serverAddress"`
	ServerPrivateKey string `json:"serverPrivateKey"`
}

type DarkStarPolishClientJsonConfig struct {
	ServerAddress    string `json:"serverAddress"`
	ServerPublicKey string `json:"serverPublicKey"`
}