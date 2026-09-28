# AlpinismoHoy
Desarrollo de una solución informática para dar respuesta a la falta de información actualizada sobre las condiciones del medio natural en cordilleras españolas.


## Problema a resolver

![Fotografía de la tarjeta del cliente](media/foto_cliente.jpg)

España es un país montañoso, que cuenta con multitud de lugares a lo largo y ancho de su geografía que son idoneos para la práctica de deportes de montaña, como pueden ser el senderismo, _parachuting_, alpinismo, ski o _mountain bike_ entre otros.

Cada uno de estos deportes requiere de una serie de condiciones naturales especificas que se deben dar para hacer posible su práctica. Sin embargo, hoy en dia es dificil conocer a donde deberiamos ir para encontrar lo que buscamos.

Concretamente en el alpinismo, existe el problema de que muchas veces se dan situaciones comprometidas en la montaña en las que una persona se halla haciendo una ruta y se encuntra con condiciones adversas que le impiden progresar adecuadamente o que lo ponen en una situacion de riesgo debido a que no se ha elegido bien el destino, ya que no se ha podido acceder a suficiente información sobre el sitio al que se fué.

Dichas situaciones adversas o simplemente no deseadas mencionadas, no se reducen simplemente al tiempo atmosférico que se pueda dar en ese momento, si no también a las influencias que lo acontecido en dias anteriores puede tener. Por ejemplo, malas condiciones de la nieve, causando desprendimientos, abalanchas o riadas, poca nieve en zonas donde se esperaba, dando lugar a requerir un equipamiento diferente, o demasiada nieve, provocando las mismas consecuencias. No existe, por mucha organización que se tenga, una manera de saber, para sitios concretos como van a ser las condiciones sin volverse loco buscando (A menos que se trate de un lugar muy famoso como algunas zonas de sierra Nevada o los Pirineos).

Por ejemplo, un sitio donde yo he vivido este problema es Sierra Magina, en el sur de Jaén. Pese a presentar condiciones que la situan como un lugar adecuado para realizar actividades de alta montaña, no se dispone de una fuente de información detallada del estado de la montaña como si ocurre en el caso de Sierra Nevada, a la hora de realizar una expedición invernal como podría ser el canuto de peña Jaén no es posible saber el estado en que se encontrará el manto nivoso o siquiera si lo habrá, por tanto surge el problema de no saber si se va a poder realizar o no la actividad que tenemos en mente. Esta situación se repite con muchas de las montañas de características similares de Andalucía, como La Sagra en la Sierra de la Sagra, la sierra de Castril o La Maroma, en Málaga.


## ¿Por que me implica personalmente?
Soy y siempre he sido un gran aficionado a todo tipo de deportes de montaña y siempre los he prácticado habiendome topado con el problema que el cliente describe a cerca de ellos en numerosas ocasiones e implicandome por tanto personalmente en la resolución del mismo.

## ¿Cómo se obtienen los datos?
Conozco diferentes foros donde  se puede encontrar información respectiva al estado de los diferentes sistemas montañosos, entendiendose por sistema montañoso, cordilleras y parques naturales de España donde se práctica alpinísmo o senderismo de alta montaña. 

Además de ello, existen los datos abiertos del satelite Copernicus, que porporcionan la posibilidad de descargar las imagenes de los lugares que interesen con cierta regularidad.

Existen aparte otros datos abiertos y descargables en formato de documento csv provenientes de plataformas similares a _suremet_ o desde la propia _aemet_.

## ¿Donde está la logica de negocio?
Se proporcionan imagenes y datos reales que solucionan el problema descrito al proporcionar toda la información requerida por el usuario para elegir el lugar concreto o saber que debe llevar a la salida que se va a realizar.

## Configuración adicional
Puede verse en detalle la configuración adicional de este proyecto en: [Ver detalles de configuración](configuracion.md)