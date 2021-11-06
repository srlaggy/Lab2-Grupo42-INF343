###############
### WINDOWS ###
###############

# build:
# @echo Construyendo Pozo
# @go build -o bin/pozo.exe pozo.go
# @echo Terminado

# clean:
# @echo Eliminando archivos ejecutables
# @del bin\"*.exe"
# @echo Terminado

# run:
# @echo Iniciando Pozo
# @bin/pozo.exe

# all:
# @echo Construyendo Pozo
# @go build -o bin/pozo.exe pozo.go
# @echo Iniciando Pozo
# @bin/pozo.exe

#############
### LINUX ###
#############

build:
	@echo Construyendo Pozo
	@go build -o bin/pozo pozo.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@rm bin/*
	@echo Terminado

run:
	@echo Iniciando Pozo
	@./bin/pozo

all:
	@echo Construyendo Pozo
	@go build -o bin/pozo pozo.go
	@echo Iniciando Pozo
	@./bin/pozo