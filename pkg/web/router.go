package web

import (
	"fortuna-express-web/pkg/domain/entities"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var sessionToken bool

func SetSessionToken(value bool) {
	sessionToken = value
}

func GetSessionToken() bool {
	return sessionToken
}

var TemplateBasePath string

type LiquidationsHandler interface {
	HomeView(user *entities.User, w http.ResponseWriter, r *http.Request) (map[string]interface{}, error)
	New(user *entities.User, w http.ResponseWriter, r *http.Request)
	NewView(user *entities.User, w http.ResponseWriter, r *http.Request)
	Update(user *entities.User, w http.ResponseWriter, r *http.Request)
	Delete(user *entities.User, w http.ResponseWriter, r *http.Request)
	Get(user *entities.User, w http.ResponseWriter, r *http.Request) (map[string]interface{}, error)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	LoginForm(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

func init() {
	// Calcula la ruta absoluta a la carpeta views dentro de web
	absPath, err := filepath.Abs("../../pkg/web/views")
	if err != nil {
		log.Fatalf("Error setting TemplateBasePath: %v", err)
	}

	TemplateBasePath = absPath + string(filepath.Separator) // Asegura que termine con un separador
	log.Println("Template base path set to:", TemplateBasePath)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verifica si la cookie de sesión existe
		sessionToken, err := c.Cookie("session_token")
		if err != nil || sessionToken != "authenticated" {
			// Si no hay cookie o no es válida, bloquea el acceso
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		// Continúa con la solicitud
		c.Next()
	}
}

func SetupRouter(r *gin.Engine, handlerLiquidation LiquidationsHandler) {
	// Ruta absoluta a las carpetas public y views
	publicDir, err := filepath.Abs("../../public")
	if err != nil {
		log.Fatal("Error obteniendo la ruta de public:", err)
	}
	viewsDir, err := filepath.Abs("../../pkg/web/views")
	if err != nil {
		log.Fatal("Error obteniendo la ruta de views:", err)
	}

	log.Println("Ruta de la carpeta public:", publicDir)
	log.Println("Ruta de la carpeta views:", viewsDir)

	// Servir archivos estáticos
	r.Static("/assets", filepath.Join(publicDir, "assets"))

	r.GET("/new", func(c *gin.Context) {
		if GetSessionToken() == false {
			c.Redirect(http.StatusFound, "/login")
			return
		} else {

			data := map[string]interface{}{
				"IsSessionActive": GetSessionToken(),
			}
			// Simula un usuario
			user := &entities.User{
				Role: "admin",
			}

			// Llama al método del handler para obtener los datos necesarios para el formulario
			handlerLiquidation.NewView(user, c.Writer, c.Request)
			if err != nil {
				log.Println("Error obteniendo datos desde NewView:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			// Cargar y renderizar las plantillas
			tmpl, err := template.ParseFiles(
				filepath.Join(publicDir, "layouts", "base.html"), // Base Layout
				filepath.Join(viewsDir, "new.html"),              // Vista específica
			)
			if err != nil {
				log.Println("Error cargando las plantillas:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			// Renderizar la plantilla con los datos obtenidos
			err = tmpl.ExecuteTemplate(c.Writer, "base.html", data)
			if err != nil {
				log.Println("Error renderizando la plantilla:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}
	})
	r.POST("/new", func(c *gin.Context) {

		if GetSessionToken() == false {
			c.Redirect(http.StatusFound, "/login")
			return
		} else {

			// Simula un usuario
			user := &entities.User{
				Role: "admin",
			}

			// Llama al método del handler para crear una nueva liquidación
			handlerLiquidation.New(user, c.Writer, c.Request)
			if err != nil {
				log.Println("Error creando una nueva liquidación:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}

	})

	r.GET("liquidations/:id", func(c *gin.Context) {
		if GetSessionToken() == false {
			c.Redirect(http.StatusFound, "/login")
			return
		} else {

			// Obtiene el parámetro de la URL
			id := c.Param("id")
			c.Request.URL.RawQuery = "id=" + id

			// Simula un usuario
			user := &entities.User{
				Role: "admin",
			}

			// Llama al método del handler para obtener los datos de una liquidación
			dataInfo2, err := handlerLiquidation.Get(user, c.Writer, c.Request)
			if err != nil {
				log.Println("Error obteniendo datos de la liquidación:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			data := map[string]interface{}{
				"data":            dataInfo2,
				"IsSessionActive": GetSessionToken(),
			}
			// Cargar y renderizar las plantillas
			tmpl, err := template.ParseFiles(
				filepath.Join(publicDir, "layouts", "base.html"),    // Base Layout
				filepath.Join(viewsDir, "liquidations/detail.html"), // Vista específica
			)
			if err != nil {
				log.Println("Error cargando las plantillas:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			// Renderizar la plantilla con los datos obtenidos
			err = tmpl.ExecuteTemplate(c.Writer, "base.html", data)
			if err != nil {
				log.Println("Error renderizando la plantilla:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}

	})

	// Ruta para /home
	r.GET("/home", func(c *gin.Context) {
		// Simula un usuario

		if GetSessionToken() == false {
			c.Redirect(http.StatusFound, "/login")
			return
		} else {
			user := &entities.User{
				Role: "admin",
			}
			// Llama a HomeView para obtener los datos necesarios
			dataHome, err := handlerLiquidation.HomeView(user, c.Writer, c.Request)
			if err != nil {
				log.Println("Error obteniendo datos desde HomeView:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			// Cargar y renderizar las plantillas
			tmpl, err := template.ParseFiles(
				filepath.Join(publicDir, "layouts", "base.html"), // Base Layout
				filepath.Join(viewsDir, "home.html"),             // Vista específica
			)
			if err != nil {
				log.Println("Error cargando las plantillas:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			data := map[string]interface{}{
				"user":            user,
				"data":            dataHome,
				"IsSessionActive": GetSessionToken(),
			}

			err = tmpl.ExecuteTemplate(c.Writer, "base.html", data)
			if err != nil {
				log.Println("Error renderizando la plantilla:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}

		}
	})
	r.GET("/login", func(c *gin.Context) {

		// Cargar y renderizar las plantillas
		tmpl, err := template.ParseFiles(
			filepath.Join(publicDir, "layouts", "base.html"), // Base Layout
			filepath.Join(viewsDir, "login.html"),            // Vista específica
		)
		if err != nil {
			log.Println("Error cargando las plantillas:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(c.Writer, "base.html", nil)
		if err != nil {
			log.Println("Error renderizando la plantilla:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	})
	r.POST("/login", func(c *gin.Context) {
		// Simula un usuario

		handlerLiquidation.LoginForm(c.Writer, c.Request)
		if err != nil {
			log.Println("Error al inciar sesion:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

	})
	r.POST("/logout", func(c *gin.Context) {
		handlerLiquidation.Logout(c.Writer, c.Request)
		if err != nil {
			log.Println("Error al cerrar sesion:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

	})

}
