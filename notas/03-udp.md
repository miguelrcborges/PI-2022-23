# UDP

Implementar um protótipo rudimentar em UDP foi relativamente simples.

Comecei por fazer um script básico de python que comunica com o
servidor de Go. [Estão aqui os resultados](https://youtu.be/pu966tCpuGA).

Atualizei o webserver a também enviar a data para a frontend,
faltando só desenvolver uma interface que dê display de toda a informação.
Para futuro eu não me esquecer, foi necessário exportar os termos das structs
(começar por uma letra maiúscula) para elas serem convertidas a JSON.
A data do stream está a ser enviada em JSON uma vez que javascript irá
interpretar a informação imediatamente como se fosse um objeto, o que é
conveniente e praticamente impossível fazer uma solução mais otimizada.

Com esta alteração, os endpoints do webserver foram comentados, uma vez
que não representa a melhor forma do dispositivo, mas sim devem ser clientes
a comunicar em UDP com o servidor, da mesma forma os dispositivos iram 
o fazer.

### EDIT

Estive a melhorar o UI, e agora tenho uma repesentação melhor da situação
(tive de incluir as portas no IP uma vez que o IP é sempre o mesmo).
[Link](https://www.youtube.com/watch?v=IV-SHyz7uKo)


## Porque UDP em vez de TCP?

O UDP é um protocolo mais rápido e mais leve (precisa de menos processamento,
permitindo poupar mais energia). TCP em contrapartida, exige que uma conexão
mantenha viva.

A maior desvantagem de UDP é que não garante a transferência da informação.
em contrapartida, não é necessário receber a integridade da informação:
a perda de informação de um loop é compensada ao ser enviada no loop seguinte,
uma vez que só é necessário da informação do caso mais recente.

O MQTT é o protocolo standard no *IoT*. No entanto, ele é um protocolo feito
sobre TCP. Assim sendo, eu acredito que consiga desenvolver por mim mesmo
um protocolo de comunicação mais eficiente, feito sobre UDP.


# Referências

- https://www.spiceworks.com/tech/networking/articles/tcp-vs-udp/
- https://www.youtube.com/watch?v=uwoD5YsGACg
