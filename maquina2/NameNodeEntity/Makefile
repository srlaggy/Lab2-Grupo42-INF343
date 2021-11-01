###############
### WINDOWS ###
###############

build:
	@echo Construyendo namenode
	@go build -o bin/namenode.exe namenode.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

run:
	@echo Iniciando servidor namenode
	@bin/namenode.exe

all:
	@echo Construyendo namenode
	@go build -o bin/namenode.exe namenode.go
	@echo Iniciando servidor namenode
	@bin/namenode.exe

#############
### LINUX ###
#############

# build:
# @echo Construyendo namenode
# @go build -o bin/namenode namenode.go
# @echo Terminado

# clean:
# @echo Eliminando archivos ejecutables
# @rm bin/*
# @echo Terminado

# run:
# @echo Iniciando servidor namenode
# @./bin/namenode

# all:
# @echo Construyendo namenode
# @go build -o bin/namenode namenode.go
# @echo Iniciando servidor namenode
# @./bin/namenode