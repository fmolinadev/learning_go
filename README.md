# Learning Go (Golang)

Este es el repositorio donde voy almacenando info a medida que voy aprendiendo las bases de Go

## Para instalar Go
Podes descargar el archivo de instalación de la página [https://golang.org/dl/](https://golang.org/dl/). 

## Configuración de Variables de Entorno en Windows:

- Precionar la tecla "Windows", buscar y seleccione: Sistema (Panel de control)
- Haga clic en el enlace Configuración avanzada del sistema (del lado derecho esta en una lista de opciones).
- Haga clic en Variables de entorno. En la sección Variables del sistema busque:

- GOROOT: En esta variable de entorno va la ruta de la instalación. Por ejm: `C:\Program Files\Go`

- PATH: En esta variable de entorno va la ruta de la instalación pero la ruta de la carpeta bin: `C:\Program Files\Go\bin`

- GOPATH: En esta variable de entorno va la ruta de tu espacio de trabajo. Por ejm: `C:\User\Francisco\go`

- GOBIN: En esta variable de entorno va la ruta de la carpeta bin que tienes que tener o crear dentro de tu espacio de trabajo, por ejemplo. `C:\User\Francisco\go\bin` 

## Estructura de espacio de trabajo de GO

- BIN: Guarda todo los ejecutables que utilicemos o creemos.

- PKG: Guarda paquetes o librerias que vas a utilizar en tu proyecto 

- SRC: Aqui vas crear todo tu codigo o tus proyectos y tambien la librerias que vas a utilizar de terceros. 

```
C:\User\Francisco\go\
    .bin
    .pkg
    .src
```

## Configuración de Variables de Entorno en Linux

Descargar el archivo para linux luego realiza la instalación con indica en la pagina, para darle permiso como administrado usa `sudo`. 

```
sudo tar -C /usr/local -xzf go1.16.7.linux-amd64.tar.gz
```

### Configuiración de Variables de entorno 
Abre el archivo oculto en la dirección de usuario que es `.bashrc` abre con algun editor y configura los siguientes variables de entorno al final de archivo. 

```
export GOPATH=/home/francisco/workspace/go
export GOBIN=$GOPATH/bin
export GOROOT=/usr/local/go
export PATH=$PATH:$GOBIN:$GOROOT/bin
```
