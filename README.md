<h1 align="center"><b>Entropy-Recon</b></h1>

<p align="center">
  <a href="https://golang.org">
    <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  </a>
  <img src="https://img.shields.io/badge/Entropy%20_%20Recon-black?style=for-the-badge" />
  <a href="./LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-black?style=for-the-badge&labelColor=black" />
  </a>
  <img src="https://img.shields.io/badge/Status-IN%20DEVELOPMENT-black?style=for-the-badge&labelColor=red" />
</p>

<div style="text-align: center;">
  <img src="https://count.getloli.com/get/@entropy-recon-mrtokyo33?theme=booru-rfck" alt="counter" />
</div>

# O que é o **Entropy Recon**
No hacking, recon é provavelmente a etapa mais importante de qualquer ataque ou pentest. É quando você coleta informações sobre o alvo e começa a entender como as coisas realmente funcionam antes de tentar explorar qualquer coisa.

Um recon mal feito normalmente resulta em horas perdidas, falsos positivos e oportunidades óbvias passando despercebidas. Você acaba atacando hosts mortos, ignorando endpoints críticos ou simplesmente não entendendo a superfície de ataque real.

O Entropy Recon foi criado pra organizar esse processo de forma lógica e eficiente, priorizando a coleta de informações que realmente possuem valor. Em vez de apenas rodar várias ferramentas soltas e rezar pra encontrar algo, a ideia é fazer análise de forma consistente, evitar ao máximo falsos positivos e mapear cada parte do alvo de verdade.

O **Entropy Recon** foi criado com a ideia de organizar esse processo de recon de forma lógica e eficiente, priorizando a coleta de informações que realmente possuem valor. Em vez de apenas rodar várias ferramentas soltas iremos fazer analise de forma consistente para evitar ao máximo os falsos positivos e analisar cada parte do alvo.

## OSINT?
Embora o recon passivo (OSINT) seja extremamente útil pra coleta de informações do alvo, o foco atual do Entropy Recon é o recon ativo. Futuramente vou adicionar uma camada de OSINT completa, mas primeiro quero deixar o ativo rodando bem.

# Como funciona?
O Entropy Recon segue um fluxo lógico de reconhecimento:
- Descoberta de ativos → o que existe?
- Validação → o que está vivo?
- Mapeamento → como isso funciona?
- Expansão → o que já existiu?
- Análise de código → tem algo sensível?
- Detecção de falhas → isso é vulnerável?

# Ferramentas que irão ser integradas:
## **1. Subfinder**
_**Asset Discovery / Subdomain Enumeration**_
- descobre subdominios de um dominio
- usa fontes passivas (APIs, certs, DNS, etc)
- Não faz brute force pesado

## **2. Httpx**
_**Service Validation / HTTP Probing**_
- testa quais hosts respondem via HTTP/HTTPS
- coleta metadata
  - status code
  - title
  - server
  - tech
  - redirect
  - headers

## **3. Katana**
_**Web Crawler + JS Parser**_
- crawling de páginas web
- parsing de java script
- extrai:
  - endpointss
  - parâmetros
  - paths
  - arquivos

## **4. gau**
_**Historical URP Enumeration (passivo)**_
- Busca URLs antigas:
  - Wayback Machine
  - Common Crawl
- Encontra endpoints que não estão mais listados

## **5. gitleaks**
_**Secret Detection / Leak Scanner**_
- procura:
  - API keys
  - tokens
  - creds

## **6. nuclei**
_**Vulnerability Scanner**_
- Executa testes de vuln reais:
  - exposures
  - misconfigs
  - CVEs
  - bugs lógicos simples

---

# **Arquitetura**
se quiser ver como funciona a arquitetura clique aqui
**->** [Ver Arquitetura](docs/arquitetura.md)
