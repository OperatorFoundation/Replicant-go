package replicant

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net"
	"net/smtp"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/OperatorFoundation/Replicant-go/Replicant/v3/polish"
	"github.com/OperatorFoundation/Replicant-go/Replicant/v3/toneburst"
	"github.com/kataras/golog"
	"golang.org/x/net/proxy"
)

func TestMain(m *testing.M) {
	runReplicantServer()
	runReplicantFactoryServer()
	os.Exit(m.Run())
}

func TestEmptyConfigs(t *testing.T) {
	clientConfig := ClientConfig{
		Toneburst: nil,
		Polish:    nil,
	}

	serverConfig := ServerConfig{
		Toneburst: nil,
		Polish:    nil,
	}

	replicantConnection(clientConfig, serverConfig, t)
}

func TestConfigNew(t *testing.T) {

	toneburstServerConfig := toneburst.StarburstConfig{
		Type: "Starburst",
		Mode: "SMTPServer",
	}

	serverPrivateKeyString := "RaHouPFVOazVSqInoMm8BSO9o/7J493y4cUVofmwXAU="

	polishServerConfig := polish.DarkStarPolishServerConfig{
		ServerAddress:   "127.0.0.1:1234",
		ServerPrivateKey: serverPrivateKeyString,
	}

	serverConfig := ServerConfig{
		ServerAddress: "127.0.0.1:1234",
		Toneburst:     toneburstServerConfig,
		Polish:        polishServerConfig,
		Transport:     "Replicant",
		BindAddress:   nil,
	}

	jsonString, marshalError := serverConfig.ToJsonString()
	if marshalError != nil {
		t.Fail()
	}

	fmt.Println(jsonString)
}

// Test an SMTP client connection that is not a valid replicant connection
func TestSMTPConnection(t *testing.T) {

	toneburstServerConfig := toneburst.StarburstConfig{
		Type: "Starburst",
		Mode: "SMTPServer",
	}

	serverPrivateKeyString := "RaHouPFVOazVSqInoMm8BSO9o/7J493y4cUVofmwXAU="

	polishServerConfig := polish.DarkStarPolishServerConfig{
		ServerAddress:    "127.0.0.1:2525",
		ServerPrivateKey: serverPrivateKeyString,
	}

	serverConfig := ServerConfig{
		ServerAddress: "127.0.0.1:2525",
		Toneburst:     toneburstServerConfig,
		Polish:        polishServerConfig,
		Transport:     "Replicant",
		BindAddress:   nil,
	}

	go func() {
		if err := listenAndServe(serverConfig); err != nil {
			fmt.Println("serve error: ", err)
		}
	}()

	c, err := smtp.Dial("127.0.0.1:2525")
	if err != nil {
		fmt.Println("SMTP Dial error")
		panic(err)
	}
	defer c.Close()

	if err := c.StartTLS(&tls.Config{InsecureSkipVerify: true}); err != nil {
		fmt.Println("StartTLS error: ", err)
		//panic(err)
	}
}

func listenAndServe(serverConfig ServerConfig) error {
	l, err := serverConfig.Listen()
	if err != nil {
		fmt.Println("Server listen error")
		panic(err)
	}
	defer l.Close()

	acceptErr := make(chan error)
	go func() {
		for {
			conn, err := l.Accept()

			fmt.Println("Accepted a connection!")
			if err != nil {
				acceptErr <- err
				fmt.Println("Listener accept error: ", err)
				return
			}

			go func(c net.Conn) {
				defer c.Close()

				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("read error: ", err)
					return
				}

				fmt.Println("from client: ", base64.StdEncoding.EncodeToString(buf[:n]))
			}(conn)
		}
	}()

	return <-acceptErr
}

func runReplicantServer() {
	serverStarted := make(chan bool)
	serverConfig := ServerConfig{
		ServerAddress: "127.0.0.1:1234",
		Toneburst:     nil,
		Polish:        nil,
		Transport:     "Replicant",
		BindAddress:   nil,
	}

	go func() {
		listener, _ := serverConfig.Listen()
		serverStarted <- true

		lConn, lConnError := listener.Accept()
		if lConnError != nil {
			return
		}

		lBuffer := make([]byte, 4)
		_, lReadError := lConn.Read(lBuffer)
		if lReadError != nil {
			return
		}

		// Send a response back to person contacting us.
		_, lWriteError := lConn.Write([]byte("Message received."))
		if lWriteError != nil {
			return
		}
	}()

	serverFinishedStarting := <-serverStarted
	if serverFinishedStarting == false {
		return
	}
}

func runReplicantFactoryServer() {
	MakeLog()
	serverStarted := make(chan bool)
	serverConfig := ServerConfig{
		ServerAddress: "127.0.0.1:3001",
		Toneburst:     nil,
		Polish:        nil,
		Transport:     "Replicant",
		BindAddress:   nil,
	}

	server := NewServer(serverConfig, proxy.Direct)

	go func() {
		listener, listenError := server.Listen()
		if listenError != nil {
			return
		}
		serverStarted <- true

		lConn, lConnError := listener.Accept()
		if lConnError != nil {
			return
		}

		lBuffer := make([]byte, 4)
		_, lReadError := lConn.Read(lBuffer)
		if lReadError != nil {
			return
		}

		// Send a response back to person contacting us.
		_, lWriteError := lConn.Write([]byte("Message received."))
		if lWriteError != nil {
			return
		}
	}()

	serverFinishedStarting := <-serverStarted
	if serverFinishedStarting == false {
		return
	}
}

func replicantConnection(clientConfig ClientConfig, serverConfig ServerConfig, t *testing.T) {
	serverStarted := make(chan bool)

	// Get a random port
	rand.Seed(time.Now().UnixNano())
	min := 1025
	max := 65535
	portNumber := min + rand.Intn(max-min+1)
	portString := strconv.Itoa(portNumber)
	addr := "127.0.0.1:"
	addr += portString

	go func() {
		listener, _ := serverConfig.Listen()
		serverStarted <- true

		lConn, lConnError := listener.Accept()
		if lConnError != nil {
			t.Fail()
			return
		}

		lBuffer := make([]byte, 4)
		_, lReadError := lConn.Read(lBuffer)
		if lReadError != nil {
			t.Fail()
			return
		}

		// Send a response back to person contacting us.
		_, lWriteError := lConn.Write([]byte("Message received."))
		if lWriteError != nil {
			t.Fail()
			return
		}

		_ = listener.Close()
	}()

	serverFinishedStarting := <-serverStarted
	if !serverFinishedStarting {
		t.Fail()
		return
	}

	cConn, connErr := clientConfig.Dial()
	if connErr != nil {
		t.Fail()
		return
	}

	writeBytes := []byte{0x0A, 0x11, 0xB0, 0xB1}
	_, cWriteError := cConn.Write(writeBytes)
	if cWriteError != nil {
		t.Fail()
		return
	}

	readBuffer := make([]byte, 17)
	_, cReadError := cConn.Read(readBuffer)
	if cReadError != nil {
		t.Fail()
		return
	}

	_ = cConn.Close()
}

func MakeLog() {
	golog.SetLevel("debug")
	golog.SetOutput(os.Stderr)
}
