package replicant

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/OperatorFoundation/Replicant-go/Replicant/v3/polish"
	"github.com/OperatorFoundation/Replicant-go/Replicant/v3/toneburst"
)

func TestDarkStarPolish(t *testing.T) {
	addr := "127.0.0.1:1234"

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
		fmt.Printf("listener type: %T\n", listener)
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
		numBytesRead, readError := serverConn.Read(buffer)
		if readError != nil {
			fmt.Println(readError)
			t.Fail()
			return
		}
		fmt.Printf("number of bytes read on server: %d\n", numBytesRead)
		fmt.Printf("serverConn type: %T\n", serverConn)

		// Send a response back to person contacting us.
		numBytesWritten, writeError := serverConn.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF})
		if writeError != nil {
			fmt.Println(writeError)
			t.Fail()
			return
		}
		fmt.Printf("number of bytes written on server: %d\n", numBytesWritten)

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
	bytesWritten, writeError := clientConn.Write(writeBytes)
	if writeError != nil {
		fmt.Println(writeError)
		t.Fail()
		return
	}
	fmt.Printf("number of bytes written on client: %d\n", bytesWritten)

	readBuffer := make([]byte, 4)
	bytesRead, readError := clientConn.Read(readBuffer)
	if readError != nil {
		fmt.Println(readError)
		t.Fail()
		return
	}
	fmt.Printf("number of bytes read on client: %d\n", bytesRead)

	_ = clientConn.Close()
}

func TestStarburstToneburst(t *testing.T) {
	addr := "127.0.0.1:1234"

	toneburstServerConfig := toneburst.StarburstConfig{
		FunctionName: "SMTPServer",
	}

	serverConfig := ServerConfig{
		Toneburst: toneburstServerConfig,
		Polish:    nil,
	}

	listener, listenError := serverConfig.Listen(addr)
	if listenError != nil {
		fmt.Println(listenError)
		t.Fail()
		return
	}

	go func() {
		fmt.Printf("listener type: %T\n", listener)
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
		numBytesRead, readError := serverConn.Read(buffer)
		if readError != nil {
			fmt.Println(readError)
			t.Fail()
			return
		}
		fmt.Printf("number of bytes read on server: %d\n", numBytesRead)
		fmt.Printf("serverConn type: %T\n", serverConn)

		// Send a response back to person contacting us.
		numBytesWritten, writeError := serverConn.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF})
		if writeError != nil {
			fmt.Println(writeError)
			t.Fail()
			return
		}
		fmt.Printf("number of bytes written on server: %d\n", numBytesWritten)

		_ = listener.Close()
	}()

	toneburstClientConfig := toneburst.StarburstConfig{
		FunctionName: "SMTPClient",
	}

	clientConfig := ClientConfig{
		Toneburst: toneburstClientConfig,
		Polish:    nil,
	}

	clientConn, clientConnError := clientConfig.Dial(addr)
	if clientConnError != nil {
		fmt.Println(clientConnError)
		t.Fail()
		return
	}

	writeBytes := []byte{0x0A, 0x11, 0xB0, 0xB1}
	bytesWritten, writeError := clientConn.Write(writeBytes)
	if writeError != nil {
		fmt.Println(writeError)
		t.Fail()
		return
	}
	fmt.Printf("number of bytes written on client: %d\n", bytesWritten)

	readBuffer := make([]byte, 4)
	bytesRead, readError := clientConn.Read(readBuffer)
	if readError != nil {
		fmt.Println(readError)
		t.Fail()
		return
	}
	fmt.Printf("number of bytes read on client: %d\n", bytesRead)

	_ = clientConn.Close()
}

func TestStarburstToneburstDarkStarPolish(t *testing.T) {
	addr := "127.0.0.1:1234"

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

	toneburstServerConfig := toneburst.StarburstConfig{
		FunctionName: "SMTPServer",
	}

	serverConfig := ServerConfig{
		Toneburst: toneburstServerConfig,
		Polish:    polishServerConfig,
	}

	listener, listenError := serverConfig.Listen(addr)
	if listenError != nil {
		fmt.Println(listenError)
		t.Fail()
		return
	}

	go func() {
		fmt.Printf("listener type: %T\n", listener)
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
		numBytesRead, readError := serverConn.Read(buffer)
		if readError != nil {
			fmt.Println(readError)
			t.Fail()
			return
		}
		fmt.Printf("number of bytes read on server: %d\n", numBytesRead)
		fmt.Printf("serverConn type: %T\n", serverConn)

		// Send a response back to person contacting us.
		numBytesWritten, writeError := serverConn.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF})
		if writeError != nil {
			fmt.Println(writeError)
			t.Fail()
			return
		}
		fmt.Printf("number of bytes written on server: %d\n", numBytesWritten)

		_ = listener.Close()
	}()

	toneburstClientConfig := toneburst.StarburstConfig{
		FunctionName: "SMTPClient",
	}

	clientConfig := ClientConfig{
		Toneburst: toneburstClientConfig,
		Polish:    polishClientConfig,
	}

	clientConn, clientConnError := clientConfig.Dial(addr)
	if clientConnError != nil {
		fmt.Println(clientConnError)
		t.Fail()
		return
	}

	writeBytes := []byte{0x0A, 0x11, 0xB0, 0xB1}
	bytesWritten, writeError := clientConn.Write(writeBytes)
	if writeError != nil {
		fmt.Println(writeError)
		t.Fail()
		return
	}
	fmt.Printf("number of bytes written on client: %d\n", bytesWritten)

	readBuffer := make([]byte, 4)
	bytesRead, readError := clientConn.Read(readBuffer)
	if readError != nil {
		fmt.Println(readError)
		t.Fail()
		return
	}
	fmt.Printf("number of bytes read on client: %d\n", bytesRead)

	_ = clientConn.Close()
}
