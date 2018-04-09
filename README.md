# UUID [![Build Status](https://travis-ci.org/Tecnologer/UUID.svg?branch=master)](https://travis-ci.org/Tecnologer/UUID)
Generador de UUID

## Definicion

Un identificador único universal o universally unique identifier (UUID) es un número de 16 bytes (128 bits). Por tanto, el número de posibles UUID es de 1632, o unos 3,4 × 1038. En su forma canónica, un UUID se expresa mediante 32 dígitos hexadecimales, divididos en cinco grupos separados por guiones de la forma 8-4-4-4-12, lo que da un total de 36 caracteres (32 dígitos y 4 guiones). Por ejemplo:

`550e8400-e29b-41d4-a716-446655440000`

Un UUID puede ser usado también con un identificador específico "intencionalmente" y repetidamente usado para identificar la misma cosa en diferentes contextos. Por ejemplo, en Microsoft Component Object Model, todos los componentes deben implementar la interfaz IUnknown (interfaz desconocido), que es realizado creando un UUID representante de IUnknow. En todos los casos cuando IUnknown es usado, ya sea usado por un proceso intentando acceder a la interfaz IUnknow en un componente, o por un componente implementando la interfaz IUnknown, siempre es referenciado por el mismo identificador:

| Name	   | Length (bytes)	| Length (hex digits)	| Contents
|----------|----------------|-----------------------|-------------------------------|
| time_low |4	            | 8	                    | Entero obtenido de los 32 primeros bits del tiempo|
| time_mid | 2	            | 4	                    | Entero obtenido de los 16 bits del medio del tiempo|
| time_hi_and_version	| 2	| 4	                    | 4-bit para la "version" en los bits mas importantes, seguido de los 12 bits de la parte superior del tiempo|
| clock_seq_hi_and_res clock_seq_low|	2	|4	| 1-3 bit de la "variante" en los bits mas importantes, seguido por los 13-15 bit de la secuencia del reloj|
node	|6	|12	|los 48 bits del identificador del nodo|


Fuente: [Wikipedia][1], [Tabla][2]

## Uso

1. Descargarlo: `go get -u github.com/tecnologer/uuid`
2. Usarlo en tu codigo
    ```Go
    package main
    import(
        "fmt"
        "github.com/tecnologer/uuid"
    )

    func main(){
        myUUID := UUID.GetUUID()
        fmt.Println(myUUID)
    }
3. Resultado (debe de ser diferente): `5e94bfea-94bf-02f9-945e-300180000134`

## Correr pruebas

Las pruebas generan N numero de UUID y se puede elegir si se quieren comparar sus resultados, asi como tambien elegir el valor de N. Los resultados de la prueba quedaran en el archivo `results.txt`.
1. Correr pruebas definiendo el numero de UUID a generar, N tomara el valor de `1000`:<br>
   `go run test/test.go -c 1000`
2. Al finalizar las pruebas se puede indicar si quiere que se comparen los resultados almacenados en `results.txt`, para esto basta agregar la bander `-t` al momento de ejecutarlo.<br>
    `go run test/test.go -t`
3. Se pueden combinar ambas banderas y su orden no afecta la ejecucion.<br>
    `go run test/test.go -c 1000 -t`<br>
    `go run test/test.go -t -c 500`<br>

[1]: <https://es.wikipedia.org/wiki/Identificador_%C3%BAnico_universal>
[2]: <https://en.wikipedia.org/wiki/Universally_unique_identifier#Format>
