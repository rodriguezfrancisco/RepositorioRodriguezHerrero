FUNCIONES PRINCIPALES

1. Nueva Rutina
2. Agregar Ejercicio
3. Listar Categorias
4. Consultar Rutina Creada
5. Modificar Rutina Creada
6. Salir

Se crean dos tipos de estructura:
- Ejercicios: representa un ejercicio individual en la rutina
- Rutina: representa una serie de ejercicios

Se usa un mapa para las categorías de ejercicios por tipo
var categorias = make(map[string][]Ejercicio)

Implementacion de funciones:
Nueva rutina: se usa para la nueva rutian las funciones: 
crearRutinaPersonalizada(): usa como variable rutina de la estructura Rutina, para crear un rutina nueva en base a ejercicios previamente creados.
Primero toma un string para asignarle un nombre a la rutina, permite agregar ejercicios a la rutina tomando categorias disponibles y nombre de ejercicio.
Retorna una rutina nueva con nombre, duracion, tipo, intensidad, calorias y descripcion de los ejercicios agregados.

Aregar Ejercicio: permite crear un ejercicio nuevo utilizando la estructura Ejercicio, creandolo con nombre duracion, tipo, intensidad, calorias y descripcion.

Listar Categorias: devuelve un mapa con categoria de ejercicios y que tiene como valor los ejercicios creados en cada categoria.
Retorna el mapa categoria con los ejercicios para cada categoria.

Consultar Rutina Creada: retorna la estructura de las rutinas ya creadas permitiendo primero seleccionar la rutina creada y luego los datos de la rutina.
(nombre de rutina, ejercicios y duracion total) En caso de no haber creado una rutina, retorna "Aun no has agregado ninguna rutina".

Modificar Rutina Creada: permite editar la rutina ya creada mediante 3 metodos que son:agregar ejercicio y eliminar ejercicio 
Agregar ejercicio: permite agregar un ejercicio nuevo a la rutina manipulando el objeto rutina mediante el append usando el metodo agregarEjercicioARutina()
Editar ejercicio: no esta listo todavia, la idea es que el usuario pueda elegir los ejercicios que quiere editar de una rutina. Así lo que haría es reasignarle los valores a los campos del tipo ejercicio.
Eliminar ejercicio: permite eliminar un ejercicio que ya fue agregado a una rutina, el usuario elige la rutina y luego el ejercicio a eliminar. Así lo que hace es crea dos slices a partir del arreglo ejercicios, uno que corta antes del ejercicio elegido y otro que empieza despues, para luego unirlos.

Salir: cierra la ejecucion del programa retornando un mensaje de salida

Por hacer:

1. Refactor para separar todas las funciones del main
2. Agregar Csv para manejo de datos (Se toma toda la pila y se tranforma en objeto)
3. Tests para todo
4. Agregar puntos por categoria (Si hay cardio junto a fuerza seria 1 y 1, si no 1 y 0)