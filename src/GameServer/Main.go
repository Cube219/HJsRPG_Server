package main

import (
	tc "GameServer/TestClient"
//	ts "GameServer/TestServer"
)

func main() {
	tc.ConnectToServer("127.0.0.1", 8888)

	//ts.OpenSocketServer(8888)
}
