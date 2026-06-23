### Sistema de Evaluación de Competencias con Strategy
### Justificación de Diseño
### 1. ¿Por qué Strategy en lugar de otros patrones?

Se eligió el patrón Strategy porque permite encapsular cada regla de evaluación en una estrategia independiente. De esta manera, las reglas pueden agregarse, modificarse o eliminarse sin afectar el resto del sistema. Además, el motor de evaluación trabaja únicamente con la interfaz de las estrategias, cumpliendo el principio Open/Closed.

No se utilizó Chain of Responsibility porque el objetivo principal no es pasar una solicitud entre varios objetos hasta encontrar quién la procesa, sino aplicar reglas específicas de forma controlada. Tampoco se utilizó Visitor porque este patrón requiere modificar la estructura cuando aparecen nuevos elementos a visitar, lo que reduce la flexibilidad para agregar nuevas reglas de negocio.

### 2. Manejo de prioridades

Cada estrategia implementa el método Prioridad(), que devuelve un valor numérico. Antes de ejecutar las reglas, estas se ordenan según su prioridad. Las reglas con menor valor se ejecutan primero.

Cuando dos reglas tienen la misma prioridad, se mantienen en el orden en que fueron registradas dentro de la colección de estrategias. Esto garantiza un comportamiento determinista y evita resultados inconsistentes entre ejecuciones.

### 3. Composición vs Agregación

Se utilizó composición en el objeto Contexto, ya que este contiene toda la información necesaria para realizar una evaluación específica: docente, competencia y evaluación.

La composición fue elegida porque estos elementos forman parte directa del proceso de evaluación. El contexto necesita poseer dichos objetos para ejecutar correctamente las estrategias. Por ejemplo, sin la información del docente o de la competencia no es posible determinar si una regla debe aplicarse.

### 4. Extensibilidad para historial

Para soportar reglas que dependan del historial del docente, se podría incorporar al contexto un servicio o repositorio encargado de consultar evaluaciones anteriores.

Por ejemplo:

Promedio histórico de evaluaciones.
Número de publicaciones realizadas.
Participación en proyectos institucionales.
Reconocimientos o sanciones previas.

Las nuevas estrategias accederían a esta información mediante interfaces, manteniendo el desacoplamiento y permitiendo seguir agregando reglas sin modificar el motor principal.
### 5. Análisis de performance

[Tu respuesta aquí]

## Ejecución

[Instrucciones para ejecutar el código]

## Resultados

[Capturas o salidas de los escenarios de prueba]
