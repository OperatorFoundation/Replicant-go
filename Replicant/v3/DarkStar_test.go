package replicant

import (
	"fmt"
	"testing"

	"github.com/OperatorFoundation/Replicant-go/Replicant/v3/polish"
	"github.com/OperatorFoundation/Replicant-go/Replicant/v3/toneburst"
)

func TestDarkStarPolish(t *testing.T) {
	serverPrivateKeyString := "RaHouPFVOazVSqInoMm8BSO9o/7J493y4cUVofmwXAU="

	polishServerConfig := polish.DarkStarPolishServerConfig{
		ServerAddress:    "127.0.0.1:2345",
		ServerPrivateKey: serverPrivateKeyString,
	}

	serverPublicKeyString := "6LukZ8KqZLQ7eOdaTVFkBVqMA8NS1AUxwqG17L/kHnQ="

	polishClientConfig := polish.DarkStarPolishClientConfig{
		ServerAddress:   "127.0.0.1:2345",
		ServerPublicKey: serverPublicKeyString,
	}

	serverConfig := ServerConfig{
		ServerAddress: "127.0.0.1:2345",
		Toneburst:     nil,
		Polish:        polishServerConfig,
		Transport:     "Replicant",
		BindAddress:   nil,
	}

	listener, listenError := serverConfig.Listen()
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
		ServerAddress: "127.0.0.1:2345",
		Toneburst:     nil,
		Polish:        polishClientConfig,
		Transport:     "Replicant",
	}

	clientConn, clientConnError := clientConfig.Dial()
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
	toneburstServerConfig := toneburst.StarburstConfig{
		Mode: "SMTPServer",
	}

	serverConfig := ServerConfig{
		ServerAddress: "127.0.0.1:2345",
		Toneburst:     toneburstServerConfig,
		Polish:        nil,
		Transport:     "Replicant",
		BindAddress:   nil,
	}

	listener, listenError := serverConfig.Listen()
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
		Mode: "SMTPClient",
	}

	clientConfig := ClientConfig{
		ServerAddress: "127.0.0.1:2345",
		Toneburst:     toneburstClientConfig,
		Polish:        nil,
		Transport:     "Replicant",
	}

	clientConn, clientConnError := clientConfig.Dial()
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
	serverPrivateKeyString := "RaHouPFVOazVSqInoMm8BSO9o/7J493y4cUVofmwXAU="

	polishServerConfig := polish.DarkStarPolishServerConfig{
		ServerAddress:    "127.0.0.1:2345",
		ServerPrivateKey: serverPrivateKeyString,
	}

	serverPublicKeyString := "6LukZ8KqZLQ7eOdaTVFkBVqMA8NS1AUxwqG17L/kHnQ="

	polishClientConfig := polish.DarkStarPolishClientConfig{
		ServerAddress:   "127.0.0.1:2345",
		ServerPublicKey: serverPublicKeyString,
	}

	toneburstServerConfig := toneburst.StarburstConfig{
		Mode: "SMTPServer",
	}

	serverConfig := ServerConfig{
		ServerAddress: "127.0.0.1:2345",
		Toneburst:     toneburstServerConfig,
		Polish:        polishServerConfig,
		Transport:     "Replicant",
		BindAddress:   nil,
	}

	listener, listenError := serverConfig.Listen()
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
		Mode: "SMTPClient",
	}

	clientConfig := ClientConfig{
		ServerAddress: "127.0.0.1:2345",
		Toneburst:     toneburstClientConfig,
		Polish:        polishClientConfig,
		Transport:     "Replicant",
	}

	clientConn, clientConnError := clientConfig.Dial()
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
