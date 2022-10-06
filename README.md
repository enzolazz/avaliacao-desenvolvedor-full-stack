# Avaliação estágio backend 
O desafio a seguir pode ser realizado em qualquer linguagem.

Além dos requisitos que serão descritos, você é livre para realizar qualquer alteração que vise a melhoria do produto. Tenha em mente que seu objetivo é criar o melhor serviço possível.

## Desafio 

Você deve criar um web service encurtador de URLs que será utilizado via chamadas REST.

O serviço deverá disponibilizar a possibilidade de cadastro e remoção de URLs, além do redirecionamento.

Ao cadastrada, a URL deve ter vinculada um **identificador não sequencial alfanumérico**, pelo qual ela será acessada.

As mesmas serão acessadas pela rota raiz: `http://.../{identificador}`, onde deve ocorrer o redirecionamento à rota original de acordo com o identificador informado.

À cada acesso, devem ser coletadas métricas de forma que seja possível as seguintes consultas:
- Quantidade de visitas na última hora
- Quantidade de visitas no último dia
- Quantidade total de visitas

## Entrega
O projeto deve ser entregue no prazo combinado na forma de um repositório git e com as instruções de execução/instalação.
