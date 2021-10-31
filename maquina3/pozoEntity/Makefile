###############
### WINDOWS ###
###############

build:
	@echo Construyendo Pozo
	@go build -o bin/pozoServer.exe pozoServer.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

run:
	@echo Iniciando servidor Pozo
	@bin/pozoServer.exe

all:
	@echo Construyendo Pozo
	@go build -o bin/pozoServer.exe pozoServer.go
	@echo Iniciando servidor Pozo
	@bin/pozoServer.exe

#############
### LINUX ###
#############

# build:
# @echo Construyendo Pozo
# @go build -o bin/pozoServer pozoServer.go
# @echo Terminado

# clean:
# @echo Eliminando archivos ejecutables
# @rm bin/*
# @echo Terminado

# run:
# @echo Iniciando servidor Pozo
# @./bin/pozoServer

# all:
# @echo Construyendo Pozo
# @go build -o bin/pozoServer pozoServer.go
# @echo Iniciando servidor Pozo
# @./bin/pozoServer