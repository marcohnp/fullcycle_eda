# Desafios Curso Full Cycle 3.0 | EDA - Event Driven Architecture
**GitHub**: https://github.com/marcohnp/fullcycle_eda

### Desafio: Criação de um microsserviço
**Descrição**: Desenvolva um microsserviço em sua linguagem de preferência que seja capaz de receber via Kafka os eventos gerados pelo microsserviço "Wallet Core" e persistir no banco de dados os balances atualizados para cada conta.   

  

Microsserviço **balance** foi criado em linguagem GO, aproveitando os conhecimentos adquiridos nesta linguagem durante o curso.  
Estrutura do projeto: o projeto está dividido em dois microsserviços, balance (desenvolvido para este desafio) e walletcore (desenvolvido pela Full Cycle). No diretório "mysql-init" estão os dois scritps que devem rodar ao subir a imagem do mysql através do docker compose.

**Passo a passo**:  
1 - Subir o docker: ```docker compose up --build -d`` 
  
2 - Criar tópicos (balances, transactions) em Confluent Control Center: http://localhost:9021  
  
3 - Chamar api para criar transação para contas pré-configuradas: ```./walletcore/api/client.http``` - OBS: a request está montada de acordo com as duas contas salvas no banco ao subir o compose.
  
4 - Chamar api para recuperar Balance das contas pré-configuradas: ```./balance/api/balance.http``` - OBS: as requests estão montadas de acordo com as duas contas salvas no banco ao subir o compose.
