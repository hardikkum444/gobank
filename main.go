package main

func main() {
	server := NewAPIServer(":8000")
	server.Run()
}
