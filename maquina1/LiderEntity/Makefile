###############
### WINDOWS ###
###############

build:
	@echo Construyendo lider
	@go build -o bin/lider.exe lider.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

run:
	@echo Iniciando servidor lider
	@bin/lider.exe

all:
	@echo Construyendo lider
	@go build -o bin/lider.exe lider.go
	@echo Iniciando servidor lider
	@bin/lider.exe

#############
### LINUX ###
#############

# build:
# @echo Construyendo lider
# @go build -o bin/lider lider.go
# @echo Terminado

# clean:
# @echo Eliminando archivos ejecutables
# @rm bin/*
# @echo Terminado

# run:
# @echo Iniciando servidor lider
# @./bin/lider

# all:
# @echo Construyendo lider
# @go build -o bin/lider lider.go
# @echo Iniciando servidor lider
# @./bin/lider