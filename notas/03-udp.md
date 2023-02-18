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
