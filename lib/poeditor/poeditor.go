package poeditor

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-restit/lzjson"
)

// Context for the API calls
type Context struct {
	ProjectID string
	APIToken  string
}

// Values returns url.Values of the context
func (ctx Context) Values() url.Values {
	return url.Values{
		"id":        []string{ctx.ProjectID},
		"api_token": []string{ctx.APIToken},
	}
}

type Response struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Language of translation
type Language struct {
	Name         string  `json:"name"`
	Code         string  `json:"code"`
	Translations int     `json:"translations"`
	Percentage   float32 `json:"percentrage"`
	Updated      string  `json:"updated"`
}

// Term for translation
type Term struct {
	Term        string      `json:"term"`
	Context     string      `json:"context"`
	Plural      string      `json:"plural"`
	Created     string      `json:"created"`
	Updated     string      `json:"updated"`
	Translation Translation `json:"translation"`
	Reference   string      `json:"reference"`
	Tags        []string    `json:"tags"`
	Comment     string      `json:"comment"`
}

// Translation of a term
type Translation struct {
	// TODO: support plural Content:
	// {
	//    "one": "Un projet disponible",
	//    "other": "%d Projets disponibles"
	// },
	Content   string `json:"content"`
	Fuzzy     int    `json:"fuzzy"`
	ProofRead int    `json:"proofread"`
	Updated   string `json:"updated"`
}

// ListLanguages list language defined within the project context.
func ListLanguages(ctx Context) (langs []Language, err error) {
	var resp *http.Response
	var apiResp Response
	if resp, err = http.PostForm("https://api.poeditor.com/v2/languages/list", ctx.Values()); err != nil {
		return
	}

	// parse API response
	jsonRoot := lzjson.Decode(resp.Body)
	jsonRoot.Get("response").Unmarshal(&apiResp)
	if apiResp.Status != "success" {
		err = fmt.Errorf("api %s: %s", apiResp.Status, apiResp.Message)
		return
	}

	err = jsonRoot.Get("result").Get("languages").Unmarshal(&langs)
	return
}

// ListTerms list term of a specific language in the given context
func ListTerms(ctx Context, lang string) (terms []Term, err error) {
	var resp *http.Response
	var apiResp Response
	values := ctx.Values()
	values.Add("language", lang)
	if resp, err = http.PostForm("https://api.poeditor.com/v2/terms/list", values); err != nil {
		return
	}

	// parse API response
	jsonRoot := lzjson.Decode(resp.Body)
	jsonRoot.Get("response").Unmarshal(&apiResp)
	if apiResp.Status != "success" {
		err = fmt.Errorf("api %s: %s", apiResp.Status, apiResp.Message)
		return
	}
	termsNode := jsonRoot.Get("result").Get("terms")
	err = termsNode.Unmarshal(&terms)
	return
}

// GetExportURL gets the export URL (expires after 10 minutes) of
// a language in the context.
func GetExportURL(ctx Context, lang, fileType string) (u string, err error) {
	var resp *http.Response
	var apiResp Response
	values := ctx.Values()
	values.Add("language", lang)
	values.Add("type", fileType)
	if resp, err = http.PostForm("https://api.poeditor.com/v2/projects/export", values); err != nil {
		return
	}

	// parse API response
	jsonRoot := lzjson.Decode(resp.Body)
	jsonRoot.Get("response").Unmarshal(&apiResp)
	if apiResp.Status != "success" {
		err = fmt.Errorf("api %s: %s", apiResp.Status, apiResp.Message)
		return
	}

	u = jsonRoot.Get("result").Get("url").String()
	return
}

// UpdateLanguage updates language with terms to insert / update translation.
func UpdateLanguage(ctx Context, lang string, terms []Term) (err error) {
	for _, term := range terms {
		fmt.Printf("%#v\n", term)
	}
	return
}
