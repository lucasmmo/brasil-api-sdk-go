# BRASIL API GOLANG SDK

## Install:

To install the brasil-api-sdk-go you need to run the command down below:

`
   $ go get github.com/lucasmmo/brasil-api-sdk-go/service/{service_name}
`

### Installing cep_v1 client:

`
    $ go get github.com/lucasmmo/brasil-api-sdk-go/service/cep_v1
`

## Usage:

```go
    func Foo() {
        // some generic function
        client := cepv1.NewClient()
        output, err := client.FindByCepWithMultiplesProviders("xxxxxxxx")
        // error handling
        // the output should will return a pointer of cepv1.CepOutput
    }
```

Or clone this project, navigate to example dir and have fun.

## Tests:

If you wanna to test in your machine, run the command:

`
    $ go test ./...
`

 This command should test all golang test files.

## Services:

- CEP (done)
- DDD
- Bank (done)
- CNPJ
- IBGE
- Feriados Nacionais
- Table FIPE
- ISBN
- Registro de domínio br
- Taxas

## License:

This project is licensed under the MIT License.

## Main maintainer:

- Lucas Mocerino Monteiro ([lucasmmo](https://github.com/lucasmmo))

