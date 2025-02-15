package gopoc

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/zan8in/afrog/pkg/poc"
	http2 "github.com/zan8in/afrog/pkg/protocols/http"
)

var pocName = "GO-phpinfo"

func pocInfo() poc.Poc {
	return poc.Poc{
		Id: pocName,
		Info: poc.Info{
			Name:        "phpinfo disclosure",
			Author:      "zan8in",
			Severity:    "low",
			Description: "phpinfo disclosure...",
			Reference: []string{
				"http://php.net",
			},
		},
	}
}

func check(args *GoPocArgs) (Result, error) {
	// init pocinfo & result
	args.SetPocInfo(pocInfo())
	result := Result{Gpa: args, IsVul: false}

	// start
	req, _ := http.NewRequest("GET", args.Target+"/info.php", nil)

	body, respraw, reqraw, status, url, err := http2.Gopochttp(req)
	if err != nil {
		return result, err
	}

	if status == 200 && bytes.Contains(body, []byte(`PHP Extension`)) && bytes.Contains(body, []byte(`PHP Version`)) {
		result.IsVul = true
		result.SetAllPocResult(true, url, reqraw, respraw)
		return result, nil
	}

	return result, errors.New("check result: no vul")
}

func init() {
	GoPocRegister(pocName, check)
}
