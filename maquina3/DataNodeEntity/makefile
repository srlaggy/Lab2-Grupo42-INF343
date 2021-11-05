###############
### WINDOWS ###
###############

build:
	@echo Construyendo datanode
	@go build -o bin/datanode.exe datanode.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

run:
	@echo Iniciando Datanode
	@bin/datanode.exe

all:
	@echo Construyendo datanode
	@go build -o bin/datanode.exe datanode.go
	@echo Iniciando Datanode
	@bin/datanode.exe

#############
### LINUX ###
#############

# build:
# @echo Construyendo datanode
# @go build -o bin/datanode datanode.go
# @echo Terminado

# clean:
# @echo Eliminando archivos ejecutables
# @rm bin/*
# @echo Terminado

# run:
# @echo Iniciando Datanode
# @./bin/datanode

# all:
# @echo Construyendo datanode
# @go build -o bin/datanode datanode.go
# @echo Iniciando Datanode
# @./bin/datanode