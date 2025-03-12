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

Una vez creado el proyecto en la carpeta de limpio procedemos copiar y renombrar el archivo de *docker-compose.yml* que se encuentra en la carpeta de *adobe* con el nombre de *limpio.yml*. En el caso de que el proyecto de limpio esta al mismo nivel que la carpeta de este proyecto, puede usar el siguiente comando para mover y renombar el archivo, en caso de que no, adapte el comando cambiando la direccion del proyecto limio.

```
cp adobe/limpio.yml ../limpio/docker-compose.yml
```

Despues se sugiere que copie el archivo script para ayudar con tareas comunes de magento y algunas acciones usuales con los contenedores del proyecto.

```
cp adobe/limpio.ckt ../limpio/ckt
```

Recuerde verificar los permisos del archivo para que se pueda ejecutar con un programa.

El siguiente paso a seguir es dar el nuevo dominio (en este caso es *limpio.local*) en su archivo de dominios locales, que en el caso de linux es el archivo */etc/hosts*. Se indica que el domino limpio va ser de caracter local agregando la siguiente linea al final del archivo.

```
127.0.0.l 	limpio.local
```

Una vez declarado que el dominio se resolvera de manera local, corrobarar que el puerto que se usara en el vhost es el mismo que se encuentra declarado en el docker-compose.yml del proyecto.

Comparar la siguiente linea del archivo *nginx/nginx/limpio.conf*.

```
upstream limpio.local { server 172.17.0.1:8006; }
```

Con el del archivo del *docker-compose.yml* del proyecto.

```
ports:
    - 8005:80
```

En caso de que vaya a levantar un nuevo proyecto, se recomiendo usar los archivos de limpio para referencia y a partir de estos archivos realizar las configuraciones de su ambiente.

El paso siguiente es tener una base de datos vacia. Recuerde que Adobe Ecommerce solo soporta mysql o en su version libre mariadb. Para facilitar esta tarea exiten dos comandos que puede emplear para crear la base de datos o vaciar una existente, para ello solo se requiere saber el id del contenedor, el cual puede copiar ejecutando el commando *docker ps* en terminal y copiando el id de su contenedor de mysql/mariadb.

Comando para crear una nueva base de datos llamada limpio en el contenedor especificado.

```
./ckt db create [id del contenedor]
```

En caso de que la base de datos limpio ya exista y quiera usarla, puede ejecutar el siguiente comando. Solo recuerde que todas las tablas en esta base de datos seran borradas permanentemente.


```
./ckt db reset [id del contenedor]
```

Una vez preparada la terminal y los servicios vivos en docker, ya deberia tener todo lo necesario para hacer la instalcion y configuracion de Adobe Ecommerce. Para iniciar la instalacion de Adobe Ecommerce puede usar el siguiente comando, el cual ya tiene las configuraciones necesaras y acorde al ambiente que se levanto.

```
php bin/magento setup:install \
--base-url="https://limpio.local/" \
--db-host="mysql" \
--db-name="limpio" \
--db-user="root" \
--db-password="mysql" \
--admin-firstname="admin" \
--admin-lastname="magento" \
--admin-email="admin@local.com" \
--admin-user="admin" \
--admin-password="local513" \
--use-rewrites="1" \
--backend-frontname="admin" \
--language=en_US \
--currency=USD \
--timezone=America/Chicago \
--opensearch-host=os-host.example.com \
--search-engine=elasticsearch7 \
--elasticsearch-host=elasticsearch7 \
--elasticsearch-port=9200
```

Para validar que la instalacion es correcta se sugiere hacer un *upgrade* para corroborar que el proyecto esta funcionando sin problemas. Para ellos puede usar el script y ejecutar el siguiente comando.

```
./ckt adobe upgrade
```

Si el comando se ejecuto bien, significa que tenemos nuestro ambiente funcionando y listo para trabajar.

## Instalacion de Sample Data en el proyecto limpio

Sample Data es informacion de prueba que oferece Adobe en su proyecto de Ecommerce para tener una tienda con informacion y visualizar como es que funciona la tienda para adaptarla a nuestras neceisdades. Para instalar la informacion de Sample Data hay que ejecutar un solo comando, el cual se recomienda antes de ejecutar que se haga un respaldo para la base de datos.

Para ejecutar el comando para instalar sample data debemos entrar al conetedor de Adobe Ecommerce, para ello ejecutamos la siguiente linea.

```
./ckt adobe
```

Una vez adentro del contenedor ejecutamos el siguiente comando.

```
php bin/magento sampledata:deploy
```

Una vez finalizado el comando realizamos un *upgrade* para que se reflejen los cambios. Este lo puede hacer aqui mismo o saliendo del contenedor con el comando *exit* y ejecutando el siguiente comando con nuestro archivo script.

```
./ckt adobe upgrade
```

Al finalizar el comando podremos ver que nuestro sitio luce diferente, ya contamos con mas informacion, exiten clientes, productos, ordenes y mas configuraciones que Adobe oferece para gestionar la tienda.

Recomendamos hacer un dump de la base de datos en este momento, ya que puede ser util si en algun futuro quiere comprobar una funcionalidad y quiere ver el flujo nativo de Adobe Ecommerce.

Para crear una cuenta de administrador puede ejecutar el siguiente comando.

```
./ckt adobe create-admin
```

Este comando le creara un usuario llamado *palomino* de caracter administrador y que podra acceder al administrador usando la siguiente url.

```
https://limpio.local/admin
```

En esta pagina tendremos un formmulario de tipo login para poder ingresar al administrador del proyecto. La contrasena del usuario es *Law200110513*.

Tomar en cuenta que el archivo script solo tienen comandos de uso frecuente, los cuales puede personalizar a su gusto y agregar mas gusta.

Eso es todo, Gracias.