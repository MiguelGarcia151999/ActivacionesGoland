
## 📚 Proyecto de Aprendizaje en Go con Fiber

¡Bienvenido/a a este proyecto! 👋 Este repositorio es parte de mi camino de aprendizaje del lenguaje Go (Golang) utilizando el framework Fiber para construir APIs eficientes y rápidas.
### 📌 Objetivos

- Aprender los fundamentos de Go: sintaxis, estructuras de datos, concurrencia, etc.
- Explorar Fiber, un framework web inspirado en Express.js, pero optimizado para Go.
- Implementar buenas prácticas: manejo de errores, testing, estructura de proyectos, etc.
- Desarrollar una API funcional con endpoints básicos (CRUD, autenticación, etc.).
### 🚀 Tecnologías Usadas

1. Tecnología Descripción
2. Go (Golang) Lenguaje de programación eficiente y concurrente.
3. Fiber Framework web rápido y minimalista para Go.
4. Git Control de versiones.
5. VS Code Editor de código principal.

### **📂 Estructura del Proyecto**
📦 **ActivacionesGoland**  
├── 📂 **config**           # Configuración principal (archivos YAML/ENV)  
├── 📂 **database**         # Conexiones a bases de datos (MySQL, MongoDB)  
├── 📂 **internal**         # Lógica interna del proyecto  
│   ├── 📂 **handlers**     # Lógica de endpoints (API)  
│   ├── 📂 **models**       # Estructuras de datos (BD)  
│   └── 📂 **routes**       # Definición de rutas HTTP  
├── 📂 **middleware**       # Interceptores (auth, logs, etc.)  
├── 📂 **tmp**              # Archivos temporales  
├── 📄 **.air.toml**        # Configuración para autoreload (Air)  
├── 📄 **.env.example**     # Variables de entorno de ejemplo  
├── 📄 **.gitignore**       # Archivos excluidos de Git  
├── 📄 **go.mod**           # Módulo de Go (dependencias)  
├── 📄 **go.sum**           # Checksum de dependencias  
├── 📄 **main.go**          # Punto de entrada  
└── 📄 **README.md**        # Documentación principal  

### 🔧 Configuración y Uso
**1. Requisitos**

    1.1 Go 1.20+ (Descargar)
    1.2 Git (Descargar)

2. Clonar y ejecutar

# Clonar el repositorio
git clone https://github.com/MiguelGarcia151999/ActivacionesGoland.git
cd ActivacionesGoland

# Instalar dependencias
go mod download

# Ejecutar el proyecto
go run main.go
Nota: Si usas Fiber, el servidor correrá por defecto en http://localhost:19070.

**3. Endpoints de ejemplo**

| Método | Ruta           | Descripción                     |
|--------|----------------|---------------------------------|
| GET    | /api/health    | Verifica el estado del API.     |
| POST   | /api/users     | Crea un nuevo usuario.          |


### 📚 Recursos de Aprendizaje
Documentación oficial de Go

* Guía de Fiber
* Go by Example

### 🤝 Contribuciones
Si tienes sugerencias o mejoras, ¡siéntete libre de hacer un fork y enviar un PR! Este proyecto es educativo, así que cualquier aporte es bienvenido.

### 📜 Licencia
Este proyecto está sin linguna licencia puesto que solo tiene fines educativos

### ✨ ¡Gracias por visitar! ✨
Si estás aprendiendo Go como yo, ¡espero que este proyecto te ayude! 🚀

###### Nota: Este README.md es un template básico. Ajusta las rutas, endpoints y estructura según tu proyecto
