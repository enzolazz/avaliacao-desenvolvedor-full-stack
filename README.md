# Avaliação Full Stack  

Este desafio pode ser realizado em qualquer linguagem ou framework.  

Além dos requisitos que serão descritos, você é livre e encorajado a realizar melhorias que visem a evolução do produto.  
Tenha em mente que o objetivo é criar o melhor serviço possível, considerando funcionalidades, eficiência, e experiência do usuário.  

A divisão em 4 partes é um direcionamento para ajudar na progressão da complexidade.  
**Não é necessário concluir todas as partes**, e as mesmas podem ser realizadas em qualquer ordem.  

## Parte 1  

Você deve criar um serviço **encurtador de URLs** que será acessado via chamadas REST.  

O serviço deverá permitir o cadastro e redirecionamento de URLs.  

Ao ser cadastrada, cada URL deve ter um **identificador não sequencial alfanumérico**, que será usado para acessá-la.  

As URLs serão acessadas pela rota raiz: `http://.../{identificador}`. Ao acessar essa rota, deve ocorrer o redirecionamento para o link original vinculado ao identificador.  

Além disso, um frontend deve ser implementado para interagir com o serviço. Este deve incluir:  
- Um **campo de entrada** para o cadastro de URLs a serem encurtadas.  
- Uma tabela que exiba as **URLs encurtadas** junto com seus identificadores.  
- Um **estilo visual moderno e responsivo**, priorizando usabilidade e experiência do usuário.  

## Parte 2  

Com o aumento da popularidade da aplicação, torna-se necessário coletar informações úteis para os usuários.  

O serviço deve coletar métricas a cada acesso, permitindo as seguintes consultas:  
- Quantidade de visitas no último dia  
- Quantidade de visitas na última hora  
- Quantidade de visitas no último mês  

O frontend deve exibir essas métricas em uma tabela ou gráfico interativo, organizado por identificadores de URLs.  

## Parte 3  

Foi identificado que muitas URLs cadastradas apontam para websites inexistentes ou que se tornaram inativos.  

Para resolver isso:  
- Implemente um **healthcheck** das URLs no momento do cadastro, bloqueando URLs inválidas.  
- Realize uma **verificação periódica** nas URLs já cadastradas. Caso uma URL falhe em 5 verificações consecutivas, ela deve ser **desativada**.  

No frontend:  
- Indique na tabela quais URLs estão ativas e quais foram desativadas, com um **sinal visual** como cores ou ícones.  

## Parte 4  

Com o aumento da utilização, a funcionalidade de *healthcheck* impacta o desempenho geral do sistema.  

Divida o serviço em dois:  
1. **Serviço Principal**: responsável pelo cadastro e redirecionamento de links.  
2. **Serviço Validador**: realiza o *healthcheck* e comunica ao serviço principal quando uma URL ultrapassa o limite de falhas.  

Garanta que os dois serviços possam operar de forma independente e escalável.  

No frontend:  
- Atualize as informações sobre o status das URLs em tempo real, sempre que houver alterações devido ao serviço validador.  

## Critérios de Avaliação  

O desafio será avaliado considerando os seguintes aspectos:  
- **Clareza do código**: organização, nomenclatura e documentação.  
- **Facilidade de manutenção**: modularidade e separação de responsabilidades.  
- **Eficiência**
