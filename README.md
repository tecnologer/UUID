# UUID

Generador de UUID

## Definicion

Un identificador único universal o universally unique identifier (UUID) es un número de 16 bytes (128 bits). Por tanto, el número de posibles UUID es de 1632, o unos 3,4 × 1038. En su forma canónica, un UUID se expresa mediante 32 dígitos hexadecimales, divididos en cinco grupos separados por guiones de la forma 8-4-4-4-12, lo que da un total de 36 caracteres (32 dígitos y 4 guiones). Por ejemplo:

`550e8400-e29b-41d4-a716-446655440000`

Un UUID puede ser usado también con un identificador específico "intencionalmente" y repetidamente usado para identificar la misma cosa en diferentes contextos. Por ejemplo, en Microsoft Component Object Model, todos los componentes deben implementar la interfaz IUnknown (interfaz desconocido), que es realizado creando un UUID representante de IUnknow. En todos los casos cuando IUnknown es usado, ya sea usado por un proceso intentando acceder a la interfaz IUnknow en un componente, o por un componente implementando la interfaz IUnknown, siempre es referenciado por el mismo identificador:

## Uso

1. Descargarlo: `go get -u github.com/tecnologer/UUID`
2. Usarlo en tu codigo
    ```Go
    import(
        "fmt"
        "github.com/tecnologer/UUID"
    )

    func main(){
        myUUID := UUID.GetUUID()
        fmt.Println(myUUID)
    }
    ```

Fuente: [Wikipedia][1]
[1]: <https://es.wikipedia.org/wiki/Identificador_%C3%BAnico_universal>