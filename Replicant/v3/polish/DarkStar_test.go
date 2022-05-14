package polish

import "testing"

func TestNewDarkStarServerConfig(t *testing.T) {
	_, configError := NewDarkStarPolishServerConfig("127.0.0.1", 1234)
	if configError != nil {
		t.Fail()
		return
	}
}

func TestNewDarkStarClientConfig(t *testing.T) {
	//serverConfig, serverConfigError := NewDarkStarPolishServerConfig("127.0.0.1", 1234)
	//if serverConfigError != nil {
	//	t.Fail()
	//	return
	//}

	//_, clientConfigError := NewDarkStarPolishClientConfig(serverConfig.ServerPrivateKey, "127.0.0.1", 1234)
	//if clientConfigError != nil {
	//	t.Fail()
	//	return
	//}
}
