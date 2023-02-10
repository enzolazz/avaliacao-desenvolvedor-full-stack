# Avaliação backend 
O desafio a seguir pode ser realizado em qualquer linguagem.

Além dos requisitos que serão descritos, você é livre e encorajado à realizar qualquer alteração que vise a melhoria do produto. Tenha em mente que seu objetivo é criar o melhor serviço que conseguir.

## Parte 1

Você deve criar um web service encurtador de URLs que será utilizado via chamadas REST.

O serviço deverá disponibilizar a possibilidade de cadastro e remoção de URLs, além do redirecionamento.

Ao cadastrada, a URL deve ter vinculada um **identificador não sequencial alfanumérico**, pelo qual ela será acessada.

As mesmas serão acessadas pela rota raiz: `http://.../{identificador}`, onde deve ocorrer o redirecionamento à rota original de acordo com o identificador informado.

## Parte 2

Com o aumento da popularidade da aplicação, é vital de coletar informações úteis para os seus usuários.

Assim, à cada acesso, devem ser coletadas métricas de forma que seja possível as seguintes consultas:
- Quantidade de visitas no último dia
- Quantidade de visitas na última hora
- Quantidade de visitas no último mês

## Parte 3

Foi detectado que uma grande quantidade dos links cadastrados são de websites inexistentes ou que se tornaram inativos.
Desta maneira, foi decidido que é necessário healizar um *healthcheck* dos links cadastrados.

Implemente uma solução que realize essa checagem no momento do cadastro do link, bloqueando se falhar.
Além disso, é necessário fazer a verificação **periódica** dos links já cadastrados, de forma que: se houver **5 falhas consecutivas** esse link seja desativado.

## Parte 4

A utilização do serviço atingiu uma alta jamais imaginada para o que havia sido desenvolvido.
Após análise, identificamos que a funcionalidade de *healthcheck* está atrapalhando o desempenho do sistema como um todo e deve ser movida para um *worker* separado.

Divida o serviço atual em dois serviços diferentes:
- Serviço principal: cadastro e redirecionamento de links
- Serviço validador: responsável por realizar o *healthcheck* e informar quando ultrapassaram o limite de falhas

# Critérios

Serão utilizados os seguintes critérios para avaliação:
- Clareza do código
- Facilidade de manutenção
- Eficiência da solução
- Escalabilidade
- Resiliência

# Entrega
O projeto deve ser entregue no prazo combinado na forma de um repositório git e com as instruções de instalação e execução.
