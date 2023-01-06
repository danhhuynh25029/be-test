start_server_node1:
	go run ./cmd/server/main.go 8000
start_client_node12:
	go run ./cmd/client/main.go 127.0.0.1:8001
start_client_node13: 
	go run ./cmd/client/main.go 127.0.0.1:8002
start_server_node2:
	go run ./cmd/server/main.go 8001
start_client_node21:
	go run ./cmd/client/main.go 127.0.0.1:8000
start_client_node23:
	go run ./cmd/client/main.go 127.0.0.1:8002
start_server_node3:
	go run ./cmd/server/main.go 8002
start_client_node31:
	go run ./cmd/client/main.go 127.0.0.1:8000
start_client_node32:
	go run ./cmd/client/main.go 127.0.0.1:8001
start_client:
	go run ./cmd/client/main.go 127.0.0.1:8000 client