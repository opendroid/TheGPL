package web

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"github.com/opendroid/the-gpl/chapter3"
	"github.com/opendroid/the-gpl/logger"
)

// templates processed
var templates *template.Template

// init Parses all the go template files
func init() {
	templates = template.Must(template.ParseGlob("public/templates/*.gohtml"))
}

// Process the "index" pattern. It also processes data by GET method sent to post page,
// 	curl -X POST localhost:8080/post --data 'q="Hello Mr 	ROb Pike"&c="K&R"&cpp=Bjarne Stroustrup'
// 		POST /post HTTP/1.1
// 		Header[Content-Length]: [31]
// 		Header[Content-Type]: [application/x-www-form-urlencoded]
// 		Header[User-Agent]: [curl/7.64.1]
// 		Header[Accept]: [*/*]
// 		HOST: localhost:8080, Remote: [::1]:63738
// 		Form[q]: ["Github"]
// 		Form[r]: ["Opendroid"]
// 		Form[s]: [Gpl]
func indexHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Println("indexHandler.")
	header := map[string]string{}                                  // data needed to show to the user
	if path, err := url.PathUnescape(r.URL.String()); err != nil { // unescape path
		header["URL"] = fmt.Sprintf("%s %s %s: %v\n", r.Method, r.URL, r.Proto, err)
	} else {
		header["URL"] = fmt.Sprintf("%s %s %s\n", r.Method, path, r.Proto)
	}
	header["Host"] = r.Host
	header["RemoteAddr"] = r.RemoteAddr
	// Save header keys
	for k, values := range r.Header {
		v := strings.Join(values, ",")
		if hv, ok := header[k]; ok { // If key data already exists append it
			header[k] = hv + ";" + header[k]
		} else {
			header[k] = v
		}
	}

	// Parse form first, reduces scope of 'err'
	if err := r.ParseForm(); err != nil {
		logger.Log.Printf("indexHandler: form parse error: %v", err)
		return
	}
	for k, values := range r.Form {
		v := strings.Join(values, ",")
		if hv, ok := header[k]; ok { // If key data already exists append it followed by ;
			header[k] = hv + ";" + header[k]
		} else {
			header[k] = v
		}
	}

	// Execute the template
	data := IndexPageData{Active: Post.String(), Data: header}
	if err := templates.ExecuteTemplate(w, IndexPage, &data); err != nil {
		logger.Log.Printf("indexHandler: %v", err)
	}
}

// aboutHandler parses about templates and and presents to use
func aboutHandler(w http.ResponseWriter, _ *http.Request) {
	logger.Log.Println("aboutHandler.")
	if err := templates.ExecuteTemplate(w, AboutPage, &AboutPageData{Active: About.String(), Data: socialCards}); err != nil {
		logger.Log.Printf("indexHandler: %v", err)
	}
}

// imagesHandler provides a scaffolding page for Images generated
func imagesHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Printf("imagesHandler: %q", r.URL.Path)
	var activePage, imagePath, heading string
	switch ImagePath(r.URL.Path) {
	case lisPath:
		activePage = Lis.String()
		heading = LisImageHanding
		imagePath = lisImagePath
	case mandelPath:
		activePage = Mandel.String()
		heading = MandelImageHanding
		imagePath = mandelImagePath
	case mandelBWPath:
		activePage = MandelBW.String()
		heading = MandelBWImageHanding
		imagePath = mandelBWImagePath
	default:
		activePage = MandelBW.String()
		heading = MandelBWImageHanding
		imagePath = mandelBWImagePath
	}
	if err := templates.ExecuteTemplate(w, LisMandelPage,
		&ImagesPageData{
			Active:    activePage,
			ImageName: imagePath,
			Heading:   heading,
		}); err != nil {
		logger.Log.Printf("imagesHandler: %v", err)
	}
}

// surfaceSVGHandler shows a specific surface
func surfaceSVGHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Printf("sincHandler: %q", r.URL.Path)
	var buf bytes.Buffer
	var activePage, heading string
	switch SVGSurfacePath(r.URL.Path) {
	case sincPath:
		activePage = Sinc.String()
		heading = SincSurfaceHeading
		chapter3.SquaresHandlerSVG(&buf)
	case sqPath:
		activePage = Square.String()
		heading = SquareSurfaceHeading
		chapter3.SquaresHandlerSVG(&buf)
	case eggPath:
		activePage = Egg.String()
		heading = EggSurfaceHeading
		chapter3.EggHandlerSVG(&buf)
	default: // valleyPath
		activePage = Valley.String()
		heading = ValleySurfaceHeading
		chapter3.ValleyHandlerSVG(&buf)
	}

	// Execute the template
	if err := templates.ExecuteTemplate(w, SurfacesPage, &SVGPageData{
		Active:  activePage,
		Data:    template.HTML(buf.String()),
		Heading: heading,
	}); err != nil {
		logger.Log.Printf("surfaceSVGHandler: %v", err)
	}
}
