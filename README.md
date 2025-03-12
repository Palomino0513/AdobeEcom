# Manejador de los serivcios de la plataforma para Adobe Ecommerce

## Validar archivo script

Se creo un script para manejar los servicios mas facil, el nombre del script es *ckt*, el cual debe de tener permisos para ser ejecutable. En caso de que el archivo script no tenga los permisos, ejecutar el siguiente comando.

```
chmod +x ckt
```

## Iniciar servicios

Para levnatar los sercivios con el archivo script puedes ejecutar el siguiente comando.

```
./ckt up
```

Tomar en cuenta que se debe generar el certificado ssl para que los proyectos usen urls seguras. En caso de no tener este certificado seguir los siguientes pasos.

### Generar certificado SSL

Para generar el certificado SSL ejecute los siguientes comandos.

## Finalizar sevicios

Para finalizar los servicios se debe ejecutar el comando *down* mediante el archivo script. Primero validar que el archivo script sea ejecutable con los pasos de **Validar archivo script**. El comando es el siguiente.

```
./ckt down
```

## Recargar sevicios

Para recargar los servicios lo mejor siempre va ser finalizarlos e iniciarlos para evitar datos cachados. En este caso el script tiene el comando *restart* para facilitar esta tarea. A continuacion el comando descrito.

```
./ckt restart
```

# Crear un proyecto limpio con sample data

Para crear un proyecto nuevo se puede ejecutar la siguiente linea.

```
composer create-project --repository-url=https://repo.magento.com/ magento/project-community-edition limpio --ignore-platform-reqs

```

