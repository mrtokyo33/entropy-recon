**<-** [Voltar para o README](../README.md)

# ARQUITETURA
irei usar a arquitetura **Service + Model + Handler**

ideia central:
- **Model** -> dados puros _(structs)_
- **Service** -> as regras de negócio
- **Handlers** -> recebe input e chama services

### **Models**
**Models** é a camada responsável por definir o esquema interno de dados do sistema. Ela estabelece quais atributos existem, seus tipos e suas relações, servindo como um contrato entre todas as outras camadas.

Models não contém fluxo, decisãao, nem interação externa. Ele apenas descreve o estado possivel do dominio que esta sendo analisado.

**primitives** é um subconjunto de Models, voltado para _value objects fundamentals_, possui o significado invariavel e bem definido, com regras próṕias de validade e identidade.

Eles existem para eliminar a ambiguidade semântica, garantir normalização estrutural e permitir comparação correta entre dados que vem de origens diferentes, com formatos e modelos distintos.

### **Services**
**Services** é a camada que concentra a lógica operacional e estratégica do sistema. Ela define o fluxo de execução, coordena chamadas e ferramenas, aplica as regras de negócio, faz a correlação e a expansão de dados, e decide quando um estado é promovido, descartado ou atualizado. 

Services é a úncia camada autorizada a modificar o estado dos models de forma consciente, sendo responsável por transformar dados brutos em informação estruturada e de valor.


### **Handlers**
***Handlers*** é a camada de entrada do sistema e atua como boundary entre interfaces externas e os usecases internos. Eles recebem inputs externos, validam parâmetros estruturais, inicializam o contexto de execução e disparam o service apropriado.

Handlers não contêm lógica de domínio nem executam ferramentas externas. sua responsabilidade é apenas conectar interfaces externas ao núcleo funcional do sistema.

### **Tools**
**Tools** é a camada que representa diretamente as ferramentas de hacking utilizadas pelo sistema. Essa camada é chamada pelo Service, que recebe o retorno bruto e o transforma em estruturas técnicas neutras, sem aplicar regras ou conceitos de domínio.. Tools não mantêm estado de domínio nem aplicam lógica de negócio; elas apenas expõem os resultados brutos das ferramentas externas de forma padronizada para consumo pelos services.