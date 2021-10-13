### Buen Dia,

El presente documento se busca sirva de guía para la persona que desee utilizar la API diseñada, de antemano agradezco la oportunidad brindada.

1. Lo primero a tener en cuenta es revisar la documentación de Golang para proceder con la instalación según corresponda el sistema operativo, en mi caso tengo Windows y se realizo la configuración e instalación en dicha máquina. 

Posterior a contar con GO, se debe realizar la creación del contenedor creado con la imagen de MySQL para proceder con la creación de la base de datos:


docker run --name mysql-dvp -e MYSQL_ROOT_PASSWORD=prueba1 -d -p 3306:3306 MySQL: latest

Para prevenir la perdida de información en el contenedor, se debe realizar la configuración del volumen utilizado para almacenar el proceso con el siguiente comando: 


docker run -d -p 33060:3306 --name mysql-dvp  -e MYSQL_ROOT_PASSWORD=secret --mount src=mysql-db-data,dst=/var/lib/mysql mysql

Una ves realizado el proceso, se puede proceder a ingresar al contenedor previamente creado, una ves adentro se procede a iniciar la imagen configurada y se continua el proceso de creación y alistamiento de la base de datos.

create database tickets;
use tickets
create table Ticket (ID int unsigned not null auto_increment comment 'Llave primaria', USUARIO varchar (30) not null, FCH_CRE date, FCH_UPD date, ESTADO varchar (20), primary key (ID));

insert into Ticket(USUARIO,FCH_CRE ,FCH_UPD ,ESTADO) values ('gilberth','2021-12-24','2021-12-29','activo');

--------------------------------------------------------
## Es recomendable contar con VSC como editor de Código.

Posterior a contar con la BD se requiere realizar la instalación de los complementos que se relacionan, esto debido a que los módulos listados se utilizaron para el desarrollo de la AP.

go mod init Prueba-Tecnica -- 

go get -u github.com/gorilla/mux -- complemento para usar router como variable de enrutamiento
go get -u github.com/go-sql-driver/mysql ---  se ejecuta para un driver que permite que GO se conecte a una base de datos

go get -u added github.com/gin-gonic/gin
go get google.golang.org/grpc



Nota: En tal caso de no ejecutar el contenedor, es necesario abrir la conexión a la base de datos mediante una herramienta como es XAMPP, verificando que la IP del servidor de conexión del localhost concuerde con 127.0.0.1, en tal caso de no hacerlo es necesario se actualice la función conec().

una ves dentro del gestor de bd, permite restaurar a una versión anterior en tal caso de requerir información que se  elimino.


CompileDaemon -command="Prueba-Tecnica.exe" --  ejecutar el build



una ves verificada la instancia de conexión con la BD, se puede proceder a ejecutar el localhost:5000 para visualizar la API que se creo. en la cual se evidenciaría la CRUD de una simulación de mesa de ayuda en la que se gestionan ticket de forma básica.
