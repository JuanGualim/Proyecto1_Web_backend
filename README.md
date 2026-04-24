# 🎬 Series Tracker API (Backend)

API REST desarrollada en Go que permite gestionar series, episodios, ratings e imágenes.  
Este backend expone un conjunto de endpoints JSON consumidos por un cliente web independiente.

---

## 🌐 Deploy

Backend desplegado en Railway:

 https://proyecto1webbackend-production-d655.up.railway.app

---

## ⚙️ Cómo correr el proyecto localmente

### 1. Clonar repositorio

```bash
git clone https://github.com/JuanGualim/Proyecto1_Web_backend
cd Proyecto1_Web_backend
```

---

### 2. Configurar variable de entorno

Se necesita una base de datos PostgreSQL.

```bash
export DATABASE_URL="postgresql://usuario:password@host:port/database"
```

---

### 3. Ejecutar el servidor

```bash
go run main.go
```

---

### 4. Acceder

```
http://localhost:8080
```

---

## 📡 Endpoints principales

### Series

- `GET /series` → Listar series  
- `GET /series/:id` → Obtener serie por ID  
- `POST /series` → Crear serie  
- `PUT /series/:id` → Editar serie  
- `DELETE /series/:id` → Eliminar serie  

---

### Rating

- `POST /series/:id/rating` → Agregar rating  
- `GET /series/:id/rating` → Obtener promedio  

---

### Imágenes

- `POST /upload` → Subir imagen  

---

## 🔍 Funcionalidades avanzadas

El endpoint `GET /series` soporta:

### Paginación

```
/series?page=1&limit=10
```

### Búsqueda

```
/series?q=breaking
```

### Ordenamiento

```
/series?sort=name&order=asc
```

---

## 📄 Swagger / OpenAPI

La documentación de la API está disponible en:

```
/docs/
```

Incluye:

- Visualización de endpoints  
- Ejecución directa desde navegador  
- Contrato de API en YAML  

---

## 🧠 CORS

CORS (Cross-Origin Resource Sharing) es una política de seguridad del navegador que restringe solicitudes entre distintos orígenes.

Se configuró en el backend permitiendo todos los orígenes mediante los headers:

- `Access-Control-Allow-Origin: *`  
- `Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS`  
- `Access-Control-Allow-Headers: Content-Type`  

Esto permite que el frontend (Vercel) consuma la API sin bloqueos.

---

## 🧩 Challenges implementados

### API y Backend

- ✔ OpenAPI / Swagger spec completa  
- ✔ Swagger UI servido desde el backend  
- ✔ Códigos HTTP correctos (201, 204, 404, 400)  
- ✔ Validación server-side con errores JSON  
- ✔ Paginación (`page`, `limit`)  
- ✔ Búsqueda (`q`)  
- ✔ Ordenamiento (`sort`, `order`)  

---

### Challenges adicionales

- ✔ Sistema de rating (tabla y endpoints propios)  
- ✔ Subida de imágenes (con límite de 1MB)  

---

## 🗄️ Base de datos

Se utiliza PostgreSQL para persistencia de datos.

Tablas principales:

- `series`  
- `ratings`  

Relación:

- Una serie puede tener múltiples ratings  
- El promedio se calcula dinámicamente  

---

##  Reflexión

Este proyecto permitió comprender la separación real entre frontend y backend utilizando una API REST. El uso de Go facilitó la construcción de un servidor eficiente, con una estructura clara y control total sobre el manejo de HTTP.

La integración con PostgreSQL aportó una capa sólida de persistencia, mientras que la implementación de validaciones, códigos HTTP correctos y Swagger permitió trabajar bajo estándares más cercanos a entornos profesionales.

Uno de los aspectos más importantes fue entender CORS y el proceso de deploy en Railway, ya que representan problemas reales en aplicaciones distribuidas.

Definitivamente utilizaría esta arquitectura nuevamente en proyectos futuros por su escalabilidad, claridad y flexibilidad.
