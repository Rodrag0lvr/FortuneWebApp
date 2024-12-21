package web

import (
	"fortuna-express-web/pkg/domain/entities"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// TemplateBasePath es la ruta absoluta a la carpeta views
var TemplateBasePath string

type LiquidationsHandler interface {
	HomeView(user *entities.User, w http.ResponseWriter, r *http.Request) (map[string]interface{}, error)
	New(user *entities.User, w http.ResponseWriter, r *http.Request)
	Update(user *entities.User, w http.ResponseWriter, r *http.Request)
	Delete(user *entities.User, w http.ResponseWriter, r *http.Request)
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
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	// Ruta para /home
	r.GET("/home", func(c *gin.Context) {
		// Simula un usuario
		user := &entities.User{
			Role: "admin",
		}

		// Llama a HomeView para obtener los datos necesarios
		data, err := handlerLiquidation.HomeView(user, c.Writer, c.Request)
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

		err = tmpl.ExecuteTemplate(c.Writer, "base.html", data)
		if err != nil {
			log.Println("Error renderizando la plantilla:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	})
}
