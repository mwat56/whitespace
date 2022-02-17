/*
   Copyright © 2019, 2022 M.Watermann, 10247 Berlin, Germany
                  All rights reserved
              EMail : <support@mwat.de>
*/

package whitespace

//lint:file-ignore ST1017 - I prefer Yoda conditions

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	in1 := []byte(`<hr />
	<p> Here is an example of AppleScript:</p>

	<pre class="AppleScript">

	tell application &quot;Foo&quot;
	  beep
	end tell

	</pre>

	<hr />`)
	out1 := []byte(`<hr /><p>Here is an example of AppleScript:</p><pre class="AppleScript">

	tell application &quot;Foo&quot;
	  beep
	end tell

	</pre><hr />`)
	in2 := []byte(`
	<p>
	bla bla bla
	</p>

	<dl class="faq">

	<dt>
	question one
	</dt>

	<dd>
	answer 1
	</dd>

	<dt>
	question two
	</dt>

	<dd>
	<pre>
		answer 2
	</pre>
	</dd>

	</dl>
	`)
	out2 := []byte(`<p>bla bla bla</p><dl class="faq"><dt>question one</dt><dd>answer 1</dd><dt>question two</dt><dd><pre>
		answer 2
	</pre></dd></dl>`)

	type args struct {
		aPage []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{" 1", args{in1}, out1},
		{" 2", args{in2}, out2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.aPage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = [%s],\nwant [%s]", got, tt.want)
			}
		})
	}
} // TestRemove()

/* _EoF_ */
