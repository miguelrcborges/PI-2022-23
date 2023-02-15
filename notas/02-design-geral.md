# Design geral

O dispositivo a construir, de forma a saber para onde tem de ir, tem de estar
em comunicação com o servidor. Esta necessidade acaba também, de certa forma,
conveniente, pois permite a monitorização dos utilizadores, sendo possível
ir em auxílio se for percebido que alguem está perdido.

Sendo a localização do utilizador feita por RFID, o dispositivo não vai saber
onde este estará localizado. A localização deste estará no computador/controlador
que esta conectado aos RFID readers (que, numa edicao futura, poderá
ser o mesmo computador que o servidor).

Para o desenvolvimento do projeto, uma vez que o sistema em RFID está fora dos
objetivos (uma colega está a trabalhar nesse objetivo), será feita uma API cuja
o controlador/computador desse sistema consiga enviar a informação da posição
dos dispositivos ao servidor.


## Design do servidor

O servidor vai ser escrito em Go. É uma linguagem simples de trabalhar, e que
na sua standard library ferramentas necessárias para o desenvolvimento.
