build:
	@echo Construyendo lider
	@go build -o bin/lider.exe util.go lider.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

run:
	bin/liderMuertos.exe

all:
	@echo Construyendo lider
	@go build -o bin/lider.exe util.go lider.go
	@echo Iniciando servidor lider
	@bin/lider.exe