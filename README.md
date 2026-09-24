# AlpinismoHoy
Desarrollo de una solución informatica para dar solución a la falta de información actualizada sobre las condiciones en sistemas españoles


## Problema a resolver

![Fotografía de la tarjeta del cliente](media/foto_cliente.jpg)

España es un país montañoso, que cuenta con multitud de lugares a lo largo y ancho de su geografía que son idoneos para la práctica de deportes de montaña, como pueden ser el senderismo, _parachuting_, alpinismo, ski o _mountain bike_ entre otros.

Cada uno de estos deportes requiere de una serie de condiciones naturales especificas que se deben dar para hacer posible su práctica. Sin embargo, hoy en dia es dificil conocer a donde deberiamos ir para encontrar lo que buscamos.

Este proyecto se centra concretamente en la problemática del alpinismo. Se pretende proporcionar una solución en forma de aplicación web que permita *extraer* de forma clara y concisa los diferentes *datos* de interes y *procesarlos* para obtener a partir de ellos una *información visual eficaz* que ayude al usuario en la tarea de decidir el lugar y la fecha ideales.

## ¿Por que me implica personalmente?
Soy y siempre he sido un gran aficionado a todo tipo de deportes de montaña y siempre los he prácticado habiendome topado con el problema que el cliente describe a cerca de ellos en numerosas ocasiones e implicandome por tanto personalmente en la resolución del mismo.

## ¿Cómo se obtienen los datos?
Conozco personalmente diferentes foros donde  se puede encontrar información respectiva al estado de los diferentes sistemas montañosos vinculada a la experiencia personal de otros montañeros.

Además de ello, existen los datos abiertos del satelite Copernicus, que porporcionan la posibilidad de descargar las imagenes de los lugares que interesen con cierta regularidad.

Existen aparte otros datos abiertos y descargables en formato de documento csv provenientes de plataformas similares a _suremet_.

## ¿Donde está la logica de negocio?
La aplicación realiza una serie de cálculos sobre los datos obtenidos que proporcionan un analisis final de la situación en los diferentes lugares y proporcionan un producto final a modo de lista en la que se refleja una puntuación (o algo similar que sirva de sistema de valor) de cada uno de los sitios para alpinismo.

## Configuración adicional
Puede verse en detalle la configuración adicional de este proyecto en: [Ver detalles de configuración](configuracion.md)