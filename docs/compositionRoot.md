**<-** [Voltar para o README](../README.md)

# Application Composition Root
atualmente minha arquitetura está;
```
├── internal
│   ├── handlers
│   ├── models
│   │   └── primitives
│   ├── services
│   ├── store
│   ├── tools
│   └── utils
```

cada camada com sua responsabilidade, porém irei criar mais uma camada porém, camada apenas de configuração + build
essa camada será a app/

na app irá ter isso:

1. Configurator
o configurator irá carregar o env, yaml (se tiver), além de futuras configurações própias da minha aplicação
ele nao vai quebrar SRP pq tipo, a responsabilidade dele é exatemente servir como um CONFIGURATOR 
ele vai receber configurações
com essas configurações ele irá fazer as validações das configurações recebidas
porém, nao faz nenhuma lógica apenas recebe as configurações e retorna a si própio

2. AppBuilder
a responsabilidade do AppBuilder é construir e conectar todos os componentes da aplicação a partir de um configurator
recebe exatamente UM configurator
instancia stores, services, handlers...
injeta dependencias
o Build() irá fazer tudo isso e irá retornar um buildType que é tudo que ele fez
a funçãp exec(b: BuildType) irá rodar o projeto

no main.go iremos ter esse pseudo-code:
```go
package main

import (
    "log"

    "myapp/app"
)

func main() {
    // 1. CONFIGURATION 
    configurator, err := app.NewConfigurator().
        LoadEnv().
        LoadYAML().
        LoadAppConf().

    configurator.Validate()

    // 2. BUILD (composition root)
    build, err := app.NewAppBuilder(configurator).
        Build()
    if err != nil {
        log.Fatal(err)
    }

    // 3. EXECUTION
    if err := build.Exec(); err != nil {
        log.Fatal(err)
    }
}
```