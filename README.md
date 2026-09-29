# AlpinismoHoy
Desarrollo de la solución para dar respuesta a la falta de información actualizada sobre las condiciones del medio natural en sistemas montañosos españoles.


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
Para la obtención de los datos, trataremos de construir una representación del estado de la montaña a partir de la previsión meteorológica, estaciones cercanas a la montaña, cobertura de nieve, nivel de agua de arroyos, pendiente, historico de precipitación y reportes humanos.

He reunido las siguientes fuentes:

- [AEMET](https://opendata.aemet.es/centrodedescargas/inicio) -> Mediante la API _OpenData_ de _AEMET_ se pueden obtener datos meteorológicos de los muicipios mas cercanos y estaciones.

- [Meteoexploration](https://meteoexploration.com/es/forecasts/) -> Una web que ofrece pronosticos de montaña en retrospectiva y a futuro, no tengo claro si se pueden obtener, estoy investigandolo.

- [Copernicus](https://land.copernicus.eu/en/products/snow/snow-cover-extent-europe-v1-0-500m) -> [Fotografía](c:\Users\javic\Downloads\2026-06-26-00_00_2026-06-26-23_59_FSC_Europe_20m_Daily_V2_FSC_OG.png) de alta resolución por satelite de Copernicus de la capa de nieve y [API estadística.](https://documentation.dataspace.copernicus.eu/APIs/SentinelHub/Statistical.html)

- [SAIH](https://www.chguadalquivir.es/saih/) -> Datos del caudal de los rios, se pbtendría mediantre scrapping de los datos de la web para los valores históricos.

- [IGN / CING](https://centrodedescargas.cnig.es/CentroDescargas/buscar-mapa) -> Para obtener mapa con elevada resolución de la zona.

- [Junta de Andalucía](https://www.juntadeandalucia.es/medioambiente/portal/web/ventanadelvisitante/detalle-buscador-mapa/-/asset_publisher/Jlbxh2qB3NwR/content/subida-a-pico-m%C3%A1gina-y-miramundos) -> Aqui se obtiene una versión en GML de las posibles rutas.

- [Wikiloc](https://es.wikiloc.com/rutas-senderismo/mata-bejid-torres-18668294) -> Scrapping para en caso de haber rutas recientes, obtener información adicional de rutas realizadas por personas.

## ¿Donde está la logica de negocio?
Se proporcionan imagenes y datos reales que solucionan el problema descrito al proporcionar toda la información requerida por el usuario para elegir el lugar concreto o saber que debe llevar a la salida que se va a realizar.

## Configuración adicional
Puede verse en detalle la configuración adicional de este proyecto en: [Ver detalles de configuración](configuracion.md)