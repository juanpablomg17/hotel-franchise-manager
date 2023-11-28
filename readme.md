# Prueba técnica ClubHub - Backend

Este proyecto es una solución para la administración de franquicias hoteleras en ClubHub, construido utilizando Golang y siguiendo diversas metodologías y tecnologías. A continuación se detallan las características principales de la implementación:

## Tecnologías utilizadas

- **Gin**: Se utilizó Gin como el framework web para construir el API REST.
- **Arquitectura hexagonal**: Se implementó una arquitectura hexagonal para separar las responsabilidades y facilitar el mantenimiento y escalabilidad del código.
- **Enfoque de DDD**: Se siguió una aproximación basada en el dominio utilizando el Diseño Dirigido por el Dominio (DDD) para modelar el sistema.
- **Programación concurrente**: Se aplicó programación concurrente para mejorar la eficiencia y capacidad de respuesta del sistema.
- **Lógica de reintentos**: Se implementó lógica de reintentos para manejar posibles errores de integración con librerías y llamadas a API externas.
- **MongoDB**: Se utilizó MongoDB como el motor de base de datos NoSQL.


## Levantar el proyecto

Para levantar el proyecto, sigue los siguientes pasos:

1. Clona el repositorio desde GitHub.

2. Asegúrate de tener Docker y Docker Compose instalados en tu máquina.

3. Navega hasta el directorio raíz del proyecto.

4. Ejecuta el siguiente comando en la terminal para construir y levantar los contenedores de Docker:

   ```
   docker-compose up -d --build
   ```

   Esto creará y ejecutará los contenedores de Docker necesarios para ejecutar la aplicación.

5. Una vez que los contenedores estén en funcionamiento, podrás acceder al API REST en `http://localhost:8080`. Puedes probar los diferentes endpoints utilizando una herramienta como Postman.

## Diagrama de arquitectura y flujo de la solución

A continuación se muestra un diagrama de la arquitectura y el flujo de la solución:

![Diagrama de arquitectura](architecture_diagram.png)

El API REST construido con Gin recibe las solicitudes del cliente y las dirige al controlador correspondiente. El controlador interactúa con los servicios de aplicación, que implementan la lógica de negocio y realizan operaciones en los repositorios de datos. Los repositorios se comunican con la capa de persistencia, que está compuesta por una base de datos SQL o NoSQL, según se haya elegido.






