# Seleção de hardware

## Microcontrolador 

Existem três marcas principais de microcontroladores para *prototyping*:
- Arduino.
- ESP.
- Micro:bit.

O Arduino não tem nenhuma forma de comunicar wireless nativamente. O
micro:bit atual tem essas funcionalidades, porém só é capaz de comunicar
entre outros micro:bits, o que implicaria ter um ligado ao servidor para
receber as mensagens dos dispositivos. Os microcontroladores da ESP tendem
a incluir Wi-Fi incluído na placa de desenvolvimento.

Assim sendo, o ideal seria utilizar uma placa de desenvolvimento da ESP,
uma vez que diminui o trabalho necessário para a transmissão de informação.
No caso não houver disponível uma placa de desenvolvimento com capacidades
de ligar a uma rede Wi-Fi, é possível usar um outro controlador com um
módulo de Wi-Fi em separado.


## Displays

Os displays para microcontroladores tendem a comunicar de uma das três formas:
- UART.
- I2C.
- SPI.


Não creio que vale a pena discutir o funcionamento de cada até saber de que tipo
será o display que estiver a trabalhar sobre. No entando, serão averiguadas as
vantagens entre estes protocolos:

- UART tem a vantagem de suportar comunicação bidirecional assincrona. Este ponto
não é vantajoso para o nosso caso uma vez que o display não tem de enviar nenhuma
informação ao controlador.
- UART e I2C só precisam de 2 cabos, enquanto SPI necessita de 4.
- SPI é o protocolo que permite maior bandwith (parcialmente por não ter bits de
começo nem paragem), seguido pelo I2C.
- I2C é mais flexível (suporta múltiplos *masters* e *slaves*), enquanto o UART é
so entre dois dispositivos e o SPI entre um *master* e 4 *slaves*. Este ponto
também não faz grande diferença uma vez que serão só usados um microcontrolandor
ligado a um display.

Para o dispositivo a fazer, nenhum protocolo revela-se ser mais útil, sendo assim
utilizado o suportado do display que for arranjado.


## Sensor de orientação

No caso de ter-se como objetivo por no display por "uma setinha a apontar a
direção certa", será necessário saber a orientação do disposititvo, pois a
localização deste não é suficiente para apresentar a direção correta.


# Referências

- https://www.seeedstudio.com/blog/2019/11/07/arduino-communication-peripherals-uart-i2c-and-spi/