# proyecto_final

Proyecto simple en Go para gestionar estudiantes, materias y calificaciones.

## Archivos principales
- Rutas (endpoints)
- Las rutas se definen en `cmd/routes/routes.go` (función `routes.SetupRoutes`). A continuación se describe por recurso los endpoints expuestos, su método y propósito:
    - Students
        - GET /students/ — Lista todos los estudiantes.
            - Handler: `handlers.StudentHandler.GetAllStudents`
            - Respuesta: JSON array con objetos estudiante.
        - GET /students/{id} — Obtiene un estudiante por su ID.
            - Handler: `handlers.StudentHandler.GetStudentByID`
            - Respuesta: JSON del estudiante o 404 si no existe.
        - POST /students/ — Crea un nuevo estudiante.
            - Handler: `handlers.StudentHandler.CreateStudent`
            - Body (JSON): { "name":string, "group_name":string, "email":string }
        - PUT /students/{id} — Actualiza un estudiante existente.
            - Handler: `handlers.StudentHandler.UpdateStudent`
            - Body (JSON): { "name":string, "group_name":string, "email":string }
        - DELETE /students/{id} — Elimina un estudiante.
            - Handler: `handlers.StudentHandler.DeleteStudent`

    - Subjects
        - GET /subjects/{id} — Obtiene una materia por ID.
            - Handler: `handlers.SubjectHandler.GetSubjectByID`
        - POST /subjects/ — Crea una nueva materia.
            - Handler: `handlers.SubjectHandler.CreateSubjects`
            - Body (JSON): { "name":string }
        - PUT /subjects/{id} — Actualiza una materia.
            - Handler: `handlers.SubjectHandler.UpdateSubject`
            - Body (JSON): { "name":string }
        - DELETE /subjects/{id} — Elimina una materia.
            - Handler: `handlers.SubjectHandler.DeleteSubject`

    - Grades
        - GET /grades/students/{student_id} — Lista todas las calificaciones de un estudiante.
            - Handler: `handlers.GradeHandler.GetAllGradesByStudentID`
        - GET /grades/{grade_id}/students/{student_id} — Obtiene una calificación específica (por grade_id y student_id).
            - Handler: `handlers.GradeHandler.GetGradeByStudentIDAndSubjectID`
        - POST /grades/ — Crea una nueva calificación.
            - Handler: `handlers.GradeHandler.CreateGrade`
            - Body (JSON): { "student_id":int, "subject_id":int, "grade":float }
        - PUT /grades/{id} — Actualiza una calificación.
            - Handler: `handlers.GradeHandler.UpdateGrade`
            - Body (JSON): { "grade":float }
        - DELETE /grades/{id} — Elimina una calificación.
            - Handler: `handlers.GradeHandler.DeleteGrade`

    - Nota: Las implementaciones de persistencia están en `cmd/repositories` (las implementaciones MySQL usan archivos `mysql_*.go`).
    - Ejemplos de peticiones para probar los endpoints están en la carpeta `rest/` (archivos `.rest`).

## Requisitos
- Go >= 1.25
- MySQL (base de datos)

## Instalación y ejecución

1. Clonar el repositorio.
2. Crear la base de datos y tablas ejecutando el script SQL:
   - Ejecutar [Script-2.sql](Script-2.sql) en tu server MySQL
3. Configurar conexión en [cmd/main.go](cmd/main.go) (cadena de conexión MySQL):
   - Actualmente: `root:Password1@tcp(127.0.0.1:3306)/School`
4. Ejecutar la aplicación:
```sh
go run cmd/main.go
```

## Ejemplos de peticiones

### Students
```
### Obtener todos los estudiantes
GET http://localhost:3000/students/


### Obtener estudiante por ID
GET http://localhost:3000/students/8

### Crear un nuevo estudiante
POST  http://localhost:3000/students/
Content-Type: application/json

{
    "name":"Carlos N",
    "group_name":"C",
    "email":"carlos2@gmail.com"
}

### Actualizar un estudiante existente
PUT  http://localhost:3000/students/8
Content-Type: application/json

{
    "name":"Alberto",
    "group_name":"B",
    "email":"nuevo@gmail.com"
}


### Eliminar un estudiante
DELETE http://localhost:3000/students/8
```

### Subjects
```
### Obtener materia por ID
GET http://localhost:3000/subjects/1

### Crear una nueva materia
POST  http://localhost:3000/subjects/
Content-Type: application/json

{
    "name":"Ciencias"
}

### Actualizar una materia existente
PUT  http://localhost:3000/subjects/5
Content-Type: application/json

{
    "name":"Ciencias Naturales"
}


### Eliminar una materia
DELETE http://localhost:3000/subjects/5
```

### Grades
```
### Obtener todas las calificaciones de un estudiante
GET http://localhost:3000/grades/students/1

### Obtener calificación por ID y estudiante
GET http://localhost:3000/grades/5/students/1


### Crear una nueva calificación
POST  http://localhost:3000/grades/
Content-Type: application/json

{
    "student_id":1,
    "subject_id":3,
    "grade":8.644
}

### Actualizar una calificación existente
PUT  http://localhost:3000/grades/5
Content-Type: application/json

{
    "grade":5.6
}

### Eliminar una calificación
DELETE http://localhost:3000/grades/5
```