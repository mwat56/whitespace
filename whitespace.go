/*
   Copyright © 2019, 2020 M.Watermann, 10247 Berlin, Germany
                  All rights reserved
              EMail : <support@mwat.de>
*/

package whitespace

//lint:file-ignore ST1017 - I prefer Yoda conditions

/*
 * This files provides a function to remove redundant whitespace
 * and comments from a HTML page.
 */

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
)

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

type (
	// `tTrimWriter` embeds a `ResponseWriter` and includes
	// whitespace removal.
	tTrimWriter struct {
		http.ResponseWriter // used to construct a HTTP response.
	}
)

// Write writes the data to the connection as part of an HTTP reply.
//
//	`aData` is the data (usually text) to send to the remote client.
func (tw *tTrimWriter) Write(aData []byte) (int, error) {
	if 0 == len(aData) {
		return 0, nil
	}

	if txt := Remove(aData); 0 < len(txt) {
		// replace the standard text with our trimmed page:
		aData = txt
	}

	return tw.ResponseWriter.Write(aData)
} // Write()

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// Internal list of regular expressions used by
// the `RemoveWhiteSpace()` function.
type (
	tReItem struct {
		regEx   *regexp.Regexp
		replace string
	}
)

var (
	wsREs = []tReItem{
		// comments:
		{regexp.MustCompile(`(?s)<!--.*?-->`), ``},
		// HTML and HEAD elements:
		{regexp.MustCompile(`(?si)\s*(</?(body|\!DOCTYPE|head|html|link|meta|script|style|title)[^>]*>)\s*`), `$1`},
		// block elements:
		{regexp.MustCompile(`(?si)\s+(</?(article|blockquote|div|footer|h[1-6]|header|nav|p|section)[^>]*>)`), `$1`},
		{regexp.MustCompile(`(?si)(</?(article|blockquote|div|footer|h[1-6]|header|nav|p|section)[^>]*>)\s+`), `$1`},
		// lists:
		{regexp.MustCompile(`(?si)\s+(</?([dou]l|li|d[dt])[^>]*>)`), `$1`},
		{regexp.MustCompile(`(?si)(</?([dou]l|li|d[dt])[^>]*>)\s+`), `$1`},
		// table elements:
		{regexp.MustCompile(`(?si)\s+(</?(col|t(able|body|foot|head|[dhr]))[^>]*>)`), `$1`},
		{regexp.MustCompile(`(?si)(</?(col|t(able|body|foot|head|[dhr]))[^>]*>)\s+`), `$1`},
		// form elements:
		{regexp.MustCompile(`(?si)\s+(</?(form|fieldset|legend|opt(group|ion))[^>]*>)`), `$1`},
		{regexp.MustCompile(`(?si)(</?(form|fieldset|legend|opt(group|ion))[^>]*>)\s+`), `$1`},
		// BR / HR:
		{regexp.MustCompile(`(?i)\s*(<[bh]r[^>]*>)\s*`), `$1`},
		// whitespace after opened anchor:
		{regexp.MustCompile(`(?si)(<a\s+[^>]*>)\s+`), `$1`},
		// preserve empty table cells:
		{regexp.MustCompile(`(?i)(<td(\s+[^>]*)?>)\s+(</td>)`), `$1&#160;$3`},
		// remove empty paragraphs:
		{regexp.MustCompile(`(?i)<(p)(\s+[^>]*)?>\s*</$1>`), ``},
		// whitespace before closing GT:
		{regexp.MustCompile(`\s+>`), `>`},
	}

	// RegEx to find PREformatted parts in an HTML page.
	wsPreRE = regexp.MustCompile(`(?si)\s*<pre[^>]*>.*?</pre>\s*`)
)

// Remove returns `aPage` with HTML comments and
// unnecessary whitespace removed.
//
// This function removes all unneeded/redundant whitespace
// and HTML comments from the given `aPage`.
// This can reduce significantly the amount of data to send to the
// remote user agent thus saving both bandwidth and transfer time.
//
//	`aPage` The web page's HTML markup to process.
func Remove(aPage []byte) []byte {
	var repl, search string

	// (0) Check whether there are PREformatted parts:
	preMatches := wsPreRE.FindAll(aPage, -1)
	if (nil == preMatches) || (0 >= len(preMatches)) {
		// no PRE hence only the other REs to perform
		for _, re := range wsREs {
			aPage = re.regEx.ReplaceAll(aPage, []byte(re.replace))
		}
		return aPage
	}

	// (1) Make sure PREformatted parts remain as-is.
	// Replace the PRE parts with a dummy text:
	for lLen, cnt := len(preMatches), 0; cnt < lLen; cnt++ {
		search = fmt.Sprintf(`\s*%s\s*`,
			regexp.QuoteMeta(string(preMatches[cnt])))
		if re, err := regexp.Compile(search); nil == err {
			repl = fmt.Sprintf(`</-%d-%d-%d-%d-/>`, cnt, cnt, cnt, cnt)
			aPage = re.ReplaceAllLiteral(aPage, []byte(repl))
		}
	}

	// (2) Traverse through all the whitespace REs:
	for _, re := range wsREs {
		aPage = re.regEx.ReplaceAll(aPage, []byte(re.replace))
	}

	// (3) Replace the PRE dummies with the real markup:
	for lLen, cnt := len(preMatches), 0; cnt < lLen; cnt++ {
		search = fmt.Sprintf(`\s*</-%d-%d-%d-%d-/>\s*`, cnt, cnt, cnt, cnt)
		if re, err := regexp.Compile(search); nil == err {
			aPage = re.ReplaceAllLiteral(aPage,
				bytes.TrimSpace(preMatches[cnt]))
		}
	}

	return aPage
} // Remove()

// Wrap returns a handler function that removes superfluous whitespace
// wrapping the given `aHandler` and calling it internally.
//
//	`aHandler` responds to the actual HTTP request.
func Wrap(aHandler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(aWriter http.ResponseWriter, aRequest *http.Request) {
			tw := &tTrimWriter{
				aWriter,
			}
			aHandler.ServeHTTP(tw, aRequest)
		})
} // Wrap()

/* _EoF_ */
