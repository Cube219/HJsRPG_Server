package main

import (
	tc "TestClient"
//	ts "TestServer"
)

func main() {
	tc.ConnectToServer("127.0.0.1", 8888)

	//ts.OpenSocketServer(8888)
}
