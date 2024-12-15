package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

// Server ... struct to hold global variables
type Server struct {
	Router *gin.Engine
	Config Config // tambahkan field config
}

type Config struct {
	SASC struct {
		JamOperasional struct {
			SeninKamis struct {
				Mulai            string `yaml:"mulai"`
				IstirahatMulai   string `yaml:"istirahat_mulai"`
				IstirahatSelesai string `yaml:"istirahat_selesai"`
				Selesai          string `yaml:"selesai"`
			} `yaml:"senin_kamis"`
			Jumat struct {
				Mulai            string `yaml:"mulai"`
				IstirahatMulai   string `yaml:"istirahat_mulai"`
				IstirahatSelesai string `yaml:"istirahat_selesai"`
				Selesai          string `yaml:"selesai"`
			} `yaml:"jumat"`
			Sabtu struct {
				Mulai            string `yaml:"mulai"`
				IstirahatMulai   string `yaml:"istirahat_mulai"`
				IstirahatSelesai string `yaml:"istirahat_selesai"`
				Selesai          string `yaml:"selesai"`
			} `yaml:"sabtu"`
		} `yaml:"jam_operasional"`
		DurasiKonsultasi int `yaml:"durasi_konsultasi"`
		MaxMahasiswaHari int `yaml:"max_mahasiswa_hari"`
	} `yaml:"sasc"`
	Jurusan map[string]struct {
		Konsultasi map[string]struct {
			Konselor []string `yaml:"konselor"`
		}
	} `yaml:"jurusan"`
}

// Init ... Initializes the app
func (server *Server) Init(port string) {
	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.New()

	// Load Configuration from YAML
	server.loadConfig("./pengaturan.yaml")

	// disini untuk routing custom sesuai kebutuhan
	server.home()
	server.register()
	server.apiStatus()

	// Load HTML and Static files
	server.Router.LoadHTMLGlob("views/*.html")
	server.Router.Static("/css", "views/css")
	server.Router.Static("/fonts", "views/fonts")
	server.Router.Static("/img", "views/img")
	server.Router.Static("/js", "views/js")

	notify := fmt.Sprintf("Server started.\nAccess on:\n - http://0.0.0.0%v/\n - http://127.0.0.1%v/", port, port)
	fmt.Println(notify)
	server.Router.Run(port)
}

func (server *Server) loadConfig(filePath string) {

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &server.Config)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML file: %v", err)
	}

}
