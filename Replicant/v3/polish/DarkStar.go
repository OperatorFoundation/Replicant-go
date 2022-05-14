package polish

import (
	"crypto"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/OperatorFoundation/go-shadowsocks2/darkstar"
	"github.com/aead/ecdh"
	"net"
)

type DarkStarPolishServerConfig struct {
	ServerPrivateKey []byte
	Host             string
	Port             int
}

type DarkStarPolishClientConfig struct {
	ServerPublicKey []byte
	Host            string
	Port            int
}

type DarkStarPolishServer struct {
	darkStarServer darkstar.DarkStarServer
}

type DarkStarPolishServerConnection struct {
	darkStarServer darkstar.DarkStarServer
	conn           net.Conn
}

type DarkStarPolishClientConnection struct {
	darkStarClient darkstar.DarkStarClient
	chunkSize      int
}

func (serverConfig DarkStarPolishServerConfig) Construct() (Server, error) {
	return NewDarkStarServer(serverConfig), nil
}

func (clientConfig DarkStarPolishClientConfig) Construct() (Connection, error) {
	return NewDarkStarClient(clientConfig), nil
}

func (server DarkStarPolishServer) NewConnection(conn net.Conn) Connection {
	serverStreamConn := server.darkStarServer.StreamConn(conn)
	return &DarkStarPolishServerConnection{
		darkStarServer: server.darkStarServer,
		conn:           serverStreamConn,
	}
}

func (serverConn *DarkStarPolishServerConnection) Handshake(conn net.Conn) (net.Conn, error) {
	return serverConn.darkStarServer.StreamConn(conn), nil
}

func (clientConn *DarkStarPolishClientConnection) Handshake(conn net.Conn) (net.Conn, error) {
	return clientConn.darkStarClient.StreamConn(conn), nil
}

func NewDarkStarPolishClientConfigFromPrivate(serverPrivateKey crypto.PrivateKey, host string, port int) (*DarkStarPolishClientConfig, error) {
	serverPublicKey := crypto.PublicKey(serverPrivateKey)
	publicKeyBytes, keyError := darkstar.PublicKeyToBytes(serverPublicKey)
	if keyError != nil {
		return nil, keyError
	}

	return &DarkStarPolishClientConfig{
		ServerPublicKey: publicKeyBytes,
		Host:            host,
		Port:            port,
	}, nil
}

func NewDarkStarPolishClientConfig(serverPublicKey []byte, host string, port int) (*DarkStarPolishClientConfig, error) {
	return &DarkStarPolishClientConfig{
		ServerPublicKey: serverPublicKey,
		Host:            host,
		Port:            port,
	}, nil
}

func NewDarkStarPolishServerConfig(host string, port int) (*DarkStarPolishServerConfig, error) {
	keyExchange := ecdh.Generic(elliptic.P256())
	clientEphemeralPrivateKey, _, keyError := keyExchange.GenerateKey(rand.Reader)
	if keyError != nil {
		return nil, keyError
	}

	privateKeyBytes, ok := clientEphemeralPrivateKey.([]byte)
	if !ok {
		return nil, errors.New("could not convert private key to bytes")
	}

	return &DarkStarPolishServerConfig{
		ServerPrivateKey: privateKeyBytes,
		Host:             host,
		Port:             port,
	}, nil
}

func NewDarkStarClient(config DarkStarPolishClientConfig) Connection {
	publicKeyString := hex.EncodeToString(config.ServerPublicKey)
	darkStarClient := darkstar.NewDarkStarClient(publicKeyString, config.Host, config.Port)
	darkStarClientConnection := DarkStarPolishClientConnection{
		darkStarClient: *darkStarClient,
	}
	return &darkStarClientConnection
}

func NewDarkStarServer(config DarkStarPolishServerConfig) DarkStarPolishServer {
	privateKeyString := hex.EncodeToString(config.ServerPrivateKey)
	darkStarServer := darkstar.NewDarkStarServer(privateKeyString, config.Host, config.Port)
	darkStarPolishServer := DarkStarPolishServer{
		darkStarServer: *darkStarServer,
	}
	return darkStarPolishServer
}
