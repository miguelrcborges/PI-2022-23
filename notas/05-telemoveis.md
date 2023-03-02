# Telemóveis

Outra alternativa a usar microcontroladores será usar telemóveis.

O desenvolvimento da aplicação para telemovel (se considerarmos
que não faço alterações no servidor) não pode ser uma PWA 
(*progressive web app*) uma vez que não é possível fazer
conexões UDP nem TCP com um servidor.

Para adicionar suporte para telemóveis tenho de fazzer uma das duas:
- Adicionar endpoints no webserver que permitem fazer websockets, e
desenvolver uma PWA.
- Desenvolver uma aplicação nativa que comunica para o mesmo endpoint
que os microcontroladores comunicam para.

A segunda solução é a que dá mais trabalho, mas é a que dá uma solução
mais robusta, uma vez que diminui a área de erro (se for preciso fazer
alguma atualização no protocolo de comunicação, ou no caso seja
apareça um problema no processamento da informação, a área do código
do servidor que tem de ser trabalhada é menor), bem como permite o uso
mais eficiente dos recursos dos telemóveis.
