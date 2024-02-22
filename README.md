# go_htmx_base
En este proyecto, nos sumergimos en el mundo de las tecnologías Go y HTMX a través de una maqueta diseñada para el aprendizaje y la aplicación práctica. Go, conocido por su eficiencia, y HTMX, que facilita la creación de interfaces web dinámicas, se combinan para explorar nuevas formas de desarrollo. Esta introducción detalla nuestro viaje en la construcción de esta maqueta, donde buscamos comprender y aprovechar al máximo estas tecnologías emergentes.

# Tecnologias
En este proyecto, utilizaremos Go como lenguaje de programación principal debido a su eficiencia, simplicidad y robustez. Chi, un marco ligero para construir aplicaciones web en Go, será nuestra elección para la gestión de rutas y la creación de servicios web. Complementando esto, integraremos HTMX, una biblioteca JavaScript que nos permitirá agregar interactividad a nuestras interfaces web de manera sencilla y eficiente. Esta combinación de tecnologías nos proporciona una base sólida y versátil para el desarrollo de nuestra maqueta, permitiéndonos explorar y aprovechar al máximo las capacidades de Go y HTMX en la creación de aplicaciones web interactivas y dinámicas.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white) ![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)

# Historial de Cambios
> [!TIP]
>  Aquí se detallara la evolución del proyecto y que se realizo en cada commit

| Version| Descripcion |
|:-:|:-|
| 0.1.0 | Se crea middleware que busca la coockie que identifica al usuario, de no existir lo envia a la pagina de `/login`. Se ajusta la llamada POST del formulario para que cunado sea exitosa redireciones a la base del proyecto.|
| 0.2.0 | Se agrega el manejo de sesiones en el backend, de esta forma se lograra tener un mayor control de las mismas. Se agrega tambien variables de entorno para obtener de forma segura la key que encripta las sesiones. Se cambia tambien el middleware encargado de revisar si la sesion se encuentra activa mantieniendo la misma funcinalidad del anterior.|
| 0.2.1 | Se crea la vista para el `/login` |