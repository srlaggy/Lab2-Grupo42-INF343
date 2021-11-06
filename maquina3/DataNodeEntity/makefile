###############
### WINDOWS ###
###############

# build:
# @del utils\playerRecord\"*.txt"
# @echo Construyendo datanode
# @go build -o bin/datanode.exe datanode.go
# @echo Terminado

# clean:
# @del utils\playerRecord\"*.txt"
# @echo Eliminando archivos ejecutables
# @del bin\"*.exe"
# @echo Terminado

# run:
# @del utils\playerRecord\"*.txt"
# @echo Iniciando Datanode
# @bin/datanode.exe

# all:
# @del utils\playerRecord\"*.txt"
# @echo Construyendo datanode
# @go build -o bin/datanode.exe datanode.go
# @echo Iniciando Datanode
# @bin/datanode.exe

#############
### LINUX ###
#############

build:
	@rm utils/playerRecord/*
	@echo Construyendo datanode
	@go build -o bin/datanode datanode.go
	@echo Terminado

clean:
	@rm utils/playerRecord/*
	@echo Eliminando archivos ejecutables
	@rm bin/*
	@echo Terminado

run:
	@rm utils/playerRecord/*
	@echo Iniciando Datanode
	@./bin/datanode

all:
	@rm utils/playerRecord/*
	@echo Construyendo datanode
	@go build -o bin/datanode datanode.go
	@echo Iniciando Datanode
	@./bin/datanode