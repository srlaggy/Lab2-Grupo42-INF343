build:
	@echo Construyendo Pozo
	@go build -o bin/pozoServer.exe util.go pozoServer.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

run:
	bin/pozoServer.exe

all:
	@echo Construyendo Pozo
	@go build -o bin/pozoServer.exe util.go pozoServer.go
	@echo Iniciando servidor Pozo
	@bin/pozoServer.exe