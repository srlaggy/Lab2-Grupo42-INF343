build:
	@echo Construyendo pozoProductorClient.exe
	@go build -o bin/pozoProductorClient.exe util.go pozoProductorClient.go
	@echo Construyendo pozoProductorServer.exe
	@go build -o bin/pozoConsumidorServer.exe util.go pozoConsumidorServer.go
	@echo Construyendo pozoClient.exe
	@go build -o bin/pozoClient.exe util.go pozoClient.go
	@echo Construyendo pozoServer.exe
	@go build -o bin/pozoServer.exe util.go pozoServer.go
	@echo Terminado

# asincronica
build-rabbit:
	@echo Construyendo pozoProductorClient.exe
	@go build -o bin/pozoProductorClient.exe util.go pozoProductorClient.go
	@echo Construyendo pozoProductorServer.exe
	@go build -o bin/pozoConsumidorServer.exe util.go pozoConsumidorServer.go
	@echo Terminado

# sincronica
build-grpc:
	@echo Construyendo pozoClient.exe
	@go build -o bin/pozoClient.exe util.go pozoClient.go
	@echo Construyendo pozoServer.exe
	@go build -o bin/pozoServer.exe util.go pozoServer.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

# rabbit client -> pozoProductorClient.exe
run-rc:
	bin/pozoProductorClient.exe

# rabbit server -> pozoConsumidorServer.exe
run-rs:
	bin/pozoConsumidorServer.exe

# grpc client -> pozoClient.exe
run-gc:
	bin/pozoClient.exe

# grpc server -> pozoServer.exe
run-gs:
	bin/pozoServer.exe