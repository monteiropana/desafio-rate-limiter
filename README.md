Passos:
para configurar e utilizar o desafio do sistema de rate limiter,que utiliza Redis para gerenciar limites de requisições.

Pré-requisitos:
- Docker e Docker Compose para subir uma instância do Redis.

Variáveis de Ambiente
Defina as seguintes variáveis no seu arquivo `.env` na pasta raiz do projeto ou exporte-as diretamente no seu ambiente:

1. **REDIS_ADDR**: Endereço do servidor Redis.
   - Exemplo: `REDIS_ADDR=redis:6379`

2. **TIME_EXP**: Tempo em segundos para a janela de rate limiting.
   - Exemplo: `TIME_EXP=15`

3. **API_KEY_NAME_LIST**: Lista de nomes de chaves de API permitidos nos cabeçalhos das requisições.
   - Exemplo: `API_KEY_NAME_LIST=api_key,api_key2,api_key3`

Usando Docker Compose:

Execute `docker-compose up` para iniciar tanto o Redis quanto a aplicação.

Manualmente:

Inicie manualmente ou use uma instância local do Redis, assegure-se de que o Redis esteja rodando no endereço definido em `REDIS_ADDR`.

Testando a Aplicação:

Dois scripts são fornecidos para testar a aplicação, podendo ser executado, no terminal, navegue até o diretório dos scripts:

bash noToken.bash
1. **noToken.bash**: Testa o limite de taxa baseado no endereço IP.

ou 

bash withToken.bash
2. **withToken.bash**: Testa o limite de taxa baseado no token de acesso.


