package replicant

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/OperatorFoundation/Replicant-go/Replicant/v3/polish"
)

//func TestNewDarkStarServerConfig(t *testing.T) {
//	_, configError := polish.NewDarkStarPolishServerConfig("127.0.0.1", 1234)
//	if configError != nil {
//		t.Fail()
//		return
//	}
//}

func TestNewDarkStarClientConfig(t *testing.T) {
	addr := "127.0.0.1:1234"

	// polishServerConfig, serverConfigError := polish.NewDarkStarPolishServerConfig("127.0.0.1", 1234)
	// if serverConfigError != nil {
	// 	fmt.Println(serverConfigError)
	// 	t.Fail()
	// 	return
	// }

	serverPrivateKeyBytes, keyError := hex.DecodeString("dd5e9e88d13e66017eb2087b128c1009539d446208f86173e30409a898ada148")
	if keyError != nil {
		fmt.Println(keyError)
		t.Fail()
		return
	}

	polishServerConfig := polish.DarkStarPolishServerConfig{
		Host:             "127.0.0.1",
		Port:             1234,
		ServerPrivateKey: serverPrivateKeyBytes,
	}

	// polishClientConfig, clientConfigError := polish.NewDarkStarPolishClientConfigFromPrivate(polishServerConfig.ServerPrivateKey, "127.0.0.1", 1234)
	// if clientConfigError != nil {
	// 	fmt.Println(clientConfigError)
	// 	t.Fail()
	// 	return
	// }

	serverPublicKeyBytes, keyError := hex.DecodeString("d089c225ef8cda8d477a586f062b31a756270124d94944e458edf1a9e1e41ed6")
	if keyError != nil {
		fmt.Println(keyError)
		t.Fail()
		return
	}

	polishClientConfig := polish.DarkStarPolishClientConfig{
		Host:            "127.0.0.1",
		Port:            1234,
		ServerPublicKey: serverPublicKeyBytes,
	}

	serverConfig := ServerConfig{
		Toneburst: nil,
		Polish:    polishServerConfig,
	}

	listener, listenError := serverConfig.Listen(addr)
	if listenError != nil {
		fmt.Println(listenError)
		t.Fail()
		return
	}

	go func() {
		serverConn, serverConnError := listener.Accept()
		if serverConnError != nil {
			fmt.Println(serverConnError)
			t.Fail()
			return
		}
		if serverConn == nil {
			fmt.Println("serverConn is nil")
			t.Fail()
			return
		}

		buffer := make([]byte, 4)
		_, readError := serverConn.Read(buffer)
		if readError != nil {
			fmt.Println(readError)
			t.Fail()
			return
		}

		// Send a response back to person contacting us.
		_, writeError := serverConn.Write([]byte("Message received."))
		if writeError != nil {
			fmt.Println(writeError)
			t.Fail()
			return
		}

		_ = listener.Close()
	}()

	clientConfig := ClientConfig{
		Toneburst: nil,
		Polish:    polishClientConfig,
	}

	clientConn, clientConnError := clientConfig.Dial(addr)
	if clientConnError != nil {
		fmt.Println(clientConnError)
		t.Fail()
		return
	}

	writeBytes := []byte{0x0A, 0x11, 0xB0, 0xB1}
	_, writeError := clientConn.Write(writeBytes)
	if writeError != nil {
		fmt.Println(writeError)
		t.Fail()
		return
	}

	readBuffer := make([]byte, 17)
	_, readError := clientConn.Read(readBuffer)
	if readError != nil {
		fmt.Println(readError)
		t.Fail()
		return
	}

	_ = clientConn.Close()
}
