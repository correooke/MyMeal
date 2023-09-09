# MyMeal

## Descripción

El proyecto MyMeal es un servicio relacionado con platos (dish-service). A continuación, se presentan los archivos clave del repositorio:

1. **Modelo de Plato (Dish Model)**: [dish-service/api/model/dish.go](https://github.com/correooke/MyMeal/blob/main/dish-service/api/model/dish.go)
2. **Router**: [dish-service/api/router/router.go](https://github.com/correooke/MyMeal/blob/main/dish-service/api/router/router.go)
3. **Punto de entrada principal (Main Entry Point)**: [dish-service/cmd/main.go](https://github.com/correooke/MyMeal/blob/main/dish-service/cmd/main.go)
4. **Dockerfile**: [dish-service/docker/Dockerfile](https://github.com/correooke/MyMeal/blob/main/dish-service/docker/Dockerfile)
5. **Docker Compose**: [dish-service/docker/docker-compose.yml](https://github.com/correooke/MyMeal/blob/main/dish-service/docker/docker-compose.yml)
6. **Manejador de Plato (Dish Handler)**: [dish-service/internal/app/dish/handler/dish.go](https://github.com/correooke/MyMeal/blob/main/dish-service/internal/app/dish/handler/dish.go)
7. **Repositorio de Plato (Dish Repository)**: [dish-service/internal/app/dish/repository/dish.go](https://github.com/correooke/MyMeal/blob/main/dish-service/internal/app/dish/repository/dish.go)
8. **Pruebas del Repositorio de Plato (Dish Repository Tests)**: [dish-service/internal/app/dish/repository/dish_test.go](https://github.com/correooke/MyMeal/blob/main/dish-service/internal/app/dish/repository/dish_test.go)
9. **Servicio de Plato (Dish Service)**: [dish-service/internal/app/dish/service/dish.go](https://github.com/correooke/MyMeal/blob/main/dish-service/internal/app/dish/service/dish.go)
10. **Implementación del Servicio de Plato (Dish Service Implementation)**: [dish-service/internal/app/dish/service/dish_impl.go](https://github.com/correooke/MyMeal/blob/main/dish-service/internal/app/dish/service/dish_impl.go)
11. **Pruebas del Servicio de Plato (Dish Service Tests)**: [dish-service/internal/app/dish/service/dish_test.go](https://github.com/correooke/MyMeal/blob/main/dish-service/internal/app/dish/service/dish_test.go)
12. **Swagger JSON**: [dish-service/swagger.json](https://github.com/correooke/MyMeal/blob/main/dish-service/swagger.json)

## Arquitectura

El proyecto utiliza una arquitectura basada en microservicios, donde el servicio principal es el `dish-service`. Este servicio se encarga de la gestión de platos, incluyendo operaciones CRUD y otras funcionalidades relacionadas. La estructura del proyecto sigue un patrón de diseño modular, con separación clara entre el modelo, el router, el manejador, el repositorio y el servicio. Además, se utiliza Docker para la contenerización y despliegue del servicio.

## Inicio Rápido

Para ejecutar el proyecto, sigue los siguientes pasos:

1. Clona el repositorio:
   ```bash
   git clone https://github.com/correooke/MyMeal.git
   cd MyMeal/dish-service
   ```
2. Construye y ejecuta el servicio con Docker Compose:
   ```bash
   docker-compose up --build
   ```
3. Accede a la documentación de la API a través de Swagger en `http://localhost:8080/swagger`.

Espero que esta guía te ayude a comprender y ejecutar el proyecto. ¡Buena suerte!