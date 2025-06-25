# Chess Game in Go

Este proyecto es una implementación de un juego de ajedrez en el lenguaje Go. El objetivo es simular las reglas y movimientos del ajedrez clásico, permitiendo jugar desde la terminal y sirviendo como base para aprender programación orientada a objetos y buenas prácticas en Go.

## Características
- **Tablero de ajedrez** representado en consola.
- **Movimientos válidos** para todas las piezas: peón, torre, alfil, caballo, reina y rey.
- **Turnos alternos** entre blancas y negras.
- **Validación de movimientos** según las reglas oficiales.
- **Código modular** y organizado por piezas y funciones.
- **Tests unitarios** para asegurar el correcto funcionamiento de cada pieza.

## Estructura del proyecto
```
Chess_Game_InGo/
├── main.go                # Entrada principal del juego
├── Functions/             # Funciones auxiliares y lógica de juego
├── Pieces/                # Lógica de cada tipo de pieza
├── inore/                 # (Carpeta de pruebas/experimentos, puede ser eliminada)
├── go.mod                 # Módulo de Go
└── README.md              # Este archivo
```

## Cómo jugar
1. Asegúrate de tener Go instalado (versión 1.18 o superior).
2. Clona este repositorio y entra a la carpeta del proyecto.
3. Ejecuta el juego con:
   ```bash
   go run main.go
   ```
4. Sigue las instrucciones en la terminal para mover las piezas.

## Cómo correr los tests
Para verificar que la lógica de las piezas funciona correctamente, ejecuta:
```bash
go test -v ./Pieces
```
Verás el resultado de los tests unitarios para cada pieza.

## Créditos
Proyecto creado por Caleb Seña Melo como ejercicio de aprendizaje y práctica de Go.

¡Disfruta programando y jugando ajedrez! ♟️
