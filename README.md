# entropy-recon

o entropy recon é o projeto que eu quero fazer desde criança, ele nao vai ser apenas um app que usa várias ferramentas, eu quero que ele seja um motor de descoberta reativo. a lógica central baseia-se no conceito de que cada nova informação encontrada sobre um alvo aumenta a superfície da ataque, e consequentemente a "entropia" que o sistema precisa organizar em grafo

a lógica funcional é:
## 1. Seed
o app recebe uma seed, que pode ser um dominio, um CIDR block ou um IP
- normalização: antes de qualquer ação, o sistema higieniza o input, indentifica o tipo de alvo e carrega as configurações de ambiente sem valores fixos no código.

## 2. Modules & Events
em vez de uma sequencia linear (A -> B -> C...) o app funciona por Reatividade
- descoberta: O sistema dispara um móodulo de enumeração
- emissão de evento: Cada subdominio encontrado é publicado como um "Evento de Ativo"
- Reação automática: Outros módulos estão "Ouvindo". Assim que um ativo aparece, eles iniciam suas tarefas de forma independente

Exemplo: Se o módulo de portas descobre a porta 80/443, ele emite um evento que ativa o "Web-Crawler". Se descobre a porta 3306, ativa o "Brute-Force-Checker" e por aí vai

## 3. Data Correlation
Enquanto as ferramentas rodam, o core do app faz a Correlaçaõ de Dados
- ele identifica que o IP pertence a um subdominio, que por sua vez está hospedade em um Bucket S3 exposto
- Lógica de Deduplicação: O sistema usa estrutras matematicas para garantir que o alvo nunca seja escaneado duas vezes, mesmo que venha de fontes diferentes

## 4. Heat Calculator
O app calcula o "peso" de cada nó no grafo
- um nó com muitas portas abertas e serviços desatualizados ganha uma pontuação de heatmap
- isso permite que o usuário veja visualmente onde as vulnerabilidades estão concentradas no gráfico

### 5. Output
ao final, a lógica de output transforma o mapa mental do app em três produtos:
- Grafo interativo (Qt Interface): onde vc clica nos nós para ver os detalhes técnicos
- Logs: A ordem cronologica de como umd ado levou a outro
- Relatorio Estruturado

---

### Seed
Seed é o ponto de partida do seu recon. ou seja é a condição inicial de um sistema
- ela é o input bruto. Pode ser um dominio(target.com), um IP (192.168.1.1), um CIDR block(10.0.0.0/24) ou um handle de rede social...
- na lógica do app: A Seed é o priemiro objeto criado no seu banco de dados que desencadeia a primeira onde de "entropia"

### Eventos e módulos
Aqui entra o SOLID puro. Em vez de um código gigante, vamos separar em funções
1. Módulo: É uma unidade isolada de funcionalidade
- Regra: Um módulo não sabe que outro existe. Eles só sabem ler e escrever eventos

2. Evento: É o "mensageiro". Quando um módulo de subdominio encontrar um subdominio, ele nao chama o PortScanner. Apena avisa ao sistema que encontrou um. 
- o sistema(engine) recebe esse evento e vê quem está interessado nele. O modulo de PortScanner diz que ele está interessado e começa a trabalhar

vamos usar Arquitetura Orientada a Eventos(EDA)

### Peso (Heat)
o Peso seria uma métrica de criticidade calculada por uma algoritmo
- Lógica: definimos regras de pesos
- No Grafo: O peso define o tamanho do nó e a cor(Heatmap)

---

### Unindo EDA(Event-Driven Arquitecture) e Clean Arquitecture
o componente A chama o componente B (acomplamento forte). Na EDA, o componente A apenas anuncia que algo aconteceu e não se importa com quem vai ler
- produtor: O módulo que descobriu algo
- Evento: O pacote de informação
- Bus/Broker: O canal por onde o evento viaja
- Consumidor: O módulo que estava esperando por aquele tipo de informação para agir

#### Estrutura do evento (contrato)
para evitar MAGIC NUMBERS e Hardcoding, o evento deve ser genérico o suficiente para carregar qualquer dado, mas rígido o suficiente para ser tipado

##### Event vai possuir:
- ID
- Type (EventType)
- Source (string)
- TimeStamp (time.Time)
- Payload (any) 
- Heat (int)

### Clean Arch
para seguir o Uncle Bob, o EventBus deve ser uma interface na camada de UseCases. A implementação real fica na camada de Infrastructure

1. Camada Domain 
na camada mais interna, ela nao conhece ninguem de fora dela. É onde noós definiremos as "leis" do programa
- Entities: São as sturcts de dados puras. exemplo: Port não é apenas um número, ela é um objeto que tem estado (Open/Closed), Serviços(HTTP/SSH) e um Heat
- Payload do Event: o Event vai ser um envelope. O Payload é o conéudo. Se é do tipo "Port Founded", o contéudo é a entitie Port

2. Camada UseCases
aqui é onde a "mágica" da automação acontece. é o cérebro que decide o fluxo do Pentest
- O orquestrador: Ele define o que deve ser feito, mas não como
- Interfaces: Em vez de chamar o nmap diretamente, o UseCase chama uma interface. Quem é o NMAP? o usecase nao sabe e nao quer saber. Ele só sabe que ao receber um Subdomain, ele deve comandar um PortScanner
- Fluxo de Dados: é aqui que definimos as regras de negócio: "Se o Heat do alvo for maior que X, dispare o módulo Y" -> VAI SER MAIS AVANÇADO QUE ISSO MAS DÁ PRA ENTENDER

3. Camada Infrastructure
está é a camada mais externa. é onde o mundo real(bin de terceiros, redes, banco de dados) interage com o resto
- Implementação do bus/broker: O sistema de events precisa de um motor para rodar. Na infrastructure, criamos o que realemnte entrega as mensagens de um lado pro outro
- Adapters: é aqui que escrevemos o "tradutor" para o Qt e para as ferramntas de hacking. Se o subdfinder cospe um texto sujo é aqui que limpamos e transformamos em uma entitie do Domain

---
### State vs Event
se o PortScanner ouvir um evento NewSubdomain. se dois módulos acharem o mesmo domínio, dois eventos serão disparados.
- Na camada UseCase, antes de processar um event, verificamos se aquele dado já existe no DB. Isso garante que o sistema seja viável (ou seja processar o mesmo dado duas vezes tem o mesmo efeito de processar só uma vez)

### Ciclo de Vida do event no Broker
usaremos Go Channels para criar essa comunicação. mas um canal simples é "um para um" (quem pega a mensagem primerio, tira ela do canal)
- Se o PortScanner e o GraphGenerator ambos querem ouvir o evento NewSubdomain, precisamos usar um padrão Fan-Out
- o broker da camada de Infrastructure deve manter uma lista de inscritos para cada tipo de evento e enviar uma cópia para cada um, como se fosse um wifi

### Definição da "Normalização" (Cleaners)
na Infrastructure vamos "limpar" o texto sujo das ferramentas
- Criaremos uma subpasta chamada parsers dentro de infrastructure. Cada Hacking Tool tera seu própio parser que transforma RAW string ou RAW JSON na entitie do Domain em algo mais limpo e fácil de testar

---
### Service Discovery de módulos
para evitar hardcoding de ter que registrar cada ferramenta manualemnte no código principal, vou fazer um sistema de "auto-discovery" ou um "registry"
Para iso usarei registry:
- cada módulo e etc terá uma função init() que se "anuncia" para o seu engine
- o engine mantém um mapa interno: map[string]Scanner

###para evitar Magic Numbers, cada hacking tool terá um manifesto(YAML ou uma struct interna). Esse manifesto dirá ao sistema:
1. name: "nmap-stealth"
2. Trigger Event: TypeNewIP (diz que só vai ser chamado se apareçer um IP)
3. Intensity: ajudará a calcular o Heat
4. Binary Path: Onde o executavel real está no sistema

O Orquestrador será um gerente de contratos
1. ele lê as configurações (ENVs e manifestos)
2. Ele inicia o Bus com suporte Fan-out
3. Ele "carrega" os módulos que passarem na validação
4. Ele fica em loop esperando a Seed Inicial

(usarei viper???)
---

# INTEGRAÇÕES FUTURAS:
- bot de discord HHEHEHEHEH
