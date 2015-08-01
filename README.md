# Identita

> Identita es un sistema criptográfico de identificación personal.

Su especificación es abierta para que terceros propongan y desarrollen sistemas complementarios que añadan funcionalidades al ecosistema de Identita, como la revocación de identificadores, entre otras.

Para el usuario Identita consiste en un identificador que puede usar para votar en un referéndum ciudadano u otro proceso participativo que requiera un alto nivel de fiabilidad (evitar votos falsos, múltiples y/o en nombre de terceros) para garantizar la legitimidad de su resultado.

> El identificador es el resultado de codificar la mínima información necesaria del usuario, cuyos datos más sensibles son cifrados con contraseña, sólo conocida por el usuario, en una cadena alfanumérica de 30 a 50 caracteres.

Para el administrador Identita es un software que le permite ser emisor de identificadores, usando su propia clave privada (basada en [MQQ-SIG192](https://github.com/imdario/mqqsig192)) para firmar el contenido y evitar la forja de identificadores por parte de terceros.

Para el promotor de procesos participativos Identita es un sistema práctico para asegurarse que los votos electrónicos son únicos y de una persona real sin menoscabar su privacidad. Identita es más fácil de usar que el DNI electrónico, 