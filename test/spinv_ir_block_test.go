// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"strings"
	"testing"

	"github.com/88250/lute"
)

var spinVditorIRBlockDOMTests = []*parseTest{

	{"11", "<p data-block=\"0\" data-node-id=\"20200920103718-6c7rbap\" data-type=\"p\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200920103713-2ahvw2c</span> <span class=\"vditor-ir__blockref\">\"*foo<wbr>\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span></p>", "<p data-block=\"0\" data-node-id=\"20200920103718-6c7rbap\" data-type=\"p\"><span data-type=\"block-ref-embed\" class=\"vditor-ir__node vditor-ir__node--expand\"><span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200920103713-2ahvw2c</span> <span class=\"vditor-ir__blockref\">\"*foo<wbr>\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span><span data-block-def-id=\"20200920103713-2ahvw2c\" data-render=\"2\" data-type=\"block-render\"></span></span></p>"},
	{"10", "<p data-block=\"0\" data-node-id=\"20200919110303-ki7u3yc\" data-type=\"p\">&lt;div&gt;<wbr></p>", "<div data-block=\"0\" data-node-id=\"20200919110303-ki7u3yc\" data-type=\"html-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"html-block\">&lt;div&gt;<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><div></pre></div>"},
	{"9", "<div data-block=\"0\" data-node-id=\"20200919105707-ofvauia\" data-type=\"code-block\" class=\"vditor-ir__node\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">​f<wbr></span><pre class=\"vditor-ir__marker--pre\"><code>\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div>", "<div data-block=\"0\" data-node-id=\"20200919105707-ofvauia\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200bf<wbr></span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code class=\"language-f\">\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code class=\"language-f\">\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"8", "<ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20200919102109-s3a5xwo\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\"><p data-block=\"0\" data-node-id=\"20200919105157-ucv97d9\" data-type=\"p\">foo</p><div data-block=\"0\" data-node-id=\"20200919105319-ie5iyfq\" data-type=\"math-block\" class=\"vditor-ir__node\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">bar<wbr></code></pre></div></li></ul>", "<ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20200919102109-s3a5xwo\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\"><p data-block=\"0\" data-node-id=\"20200919105157-ucv97d9\" data-type=\"p\">foo</p><div data-block=\"0\" data-node-id=\"20200919105319-ie5iyfq\" data-type=\"math-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">bar<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code data-type=\"math-block\" class=\"language-math\">bar</code></pre><span data-type=\"math-block-close-marker\">$$</span></div></li></ul>"},
	{"7", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200919102109-s3a5xwo\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\">foo<div data-block=\"0\" data-node-id=\"20200919102144-q0l1bco\" data-type=\"code-block\" class=\"vditor-ir__node\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">​</span><pre class=\"vditor-ir__marker--pre\"><code>bar<wbr></code></pre></div></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200919102109-s3a5xwo\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\">foo<div data-block=\"0\" data-node-id=\"20200919102144-q0l1bco\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>bar<wbr>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>bar\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div></li></ul>"},
	{"6", "<blockquote data-block=\"0\" data-node-id=\"20200919091557-m4ivx3j\" data-type=\"blockquote\"><p data-block=\"0\" data-node-id=\"20200919091706-6f175t8\" data-type=\"p\"><wbr><br></p></blockquote>", "<blockquote data-block=\"0\" data-node-id=\"20200919091557-m4ivx3j\" data-type=\"blockquote\"><p data-block=\"0\" data-node-id=\"20200919091706-6f175t8\" data-type=\"p\"><wbr></p></blockquote>"},
	{"5", "<p data-block=\"0\" data-node-id=\"20200916084414-0hk6l8v\" data-type=\"p\">![foo中文.png](assets/foo中文<wbr>.png)</p>", "<p data-block=\"0\" data-node-id=\"20200916084414-0hk6l8v\" data-type=\"p\"><span class=\"vditor-ir__node vditor-ir__node--expand\" data-type=\"img\"><span class=\"vditor-ir__marker\">!</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">[</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">foo 中文.png</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">]</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">assets/foo中文<wbr>.png</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><img src=\" http://127.0.0.1:6807/webdav/q0fk7yv/测试笔记/assets/foo中文.png\" alt=\"foo 中文.png\" /></span></p>"},
	{"4", "<h1 data-block=\"0\" class=\"vditor-ir__node\" data-node-id=\"test\" bookmark=\"bookmark\" data-type=\"h\" id=\"ir-foo\" data-marker=\"#\"><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo</h1>", "<h1 data-block=\"0\" class=\"vditor-ir__node\" data-node-id=\"test\" bookmark=\"bookmark\" data-type=\"h\" id=\"ir-foo\" data-marker=\"#\"><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo</h1>"},
	{"3", "<div data-block=\"0\" data-node-id=\"test\" bookmark=\"bookmark\" data-type=\"code-block\" class=\"vditor-ir__node\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>foo\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>foo\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div>", "<div data-block=\"0\" data-node-id=\"test\" bookmark=\"bookmark\" data-type=\"code-block\" class=\"vditor-ir__node\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>foo\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>foo\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"2", "<p data-block=\"0\" data-node-id=\"20200915173154-1wi2p2h\" data-type=\"p\">$$<wbr></p>", "<div data-block=\"0\" data-node-id=\"20200915173154-1wi2p2h\" data-type=\"math-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\"><wbr>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code data-type=\"math-block\" class=\"language-math\"></code></pre><span data-type=\"math-block-close-marker\">$$</span></div>"},
	{"1", "<p data-block=\"0\" data-node-id=\"20200915172226-iexs3bo\" data-type=\"p\">```<wbr></p>", "<div data-block=\"0\" data-node-id=\"20200915172226-iexs3bo\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b<wbr></span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code></code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"0", "<p data-block=\"0\" data-node-id=\"20200914181352-laa3jyd\" data-type=\"p\">foo<wbr></p>", "<p data-block=\"0\" data-node-id=\"20200914181352-laa3jyd\" data-type=\"p\">foo<wbr></p>"},
}

func TestSpinVditorIRBlockDOM(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.BlockRef = true
	luteEngine.KramdownIAL = true
	luteEngine.SetLinkBase(" http://127.0.0.1:6807/webdav/q0fk7yv/测试笔记/")

	for _, test := range spinVditorIRBlockDOMTests {
		html := luteEngine.SpinVditorIRBlockDOM(test.from)

		if "15" == test.name || "18" == test.name {
			// 去掉最后一个新生成的节点 ID，因为这个节点 ID 是随机生成，每次测试用例运行都不一样，比较没有意义，长度一致即可
			lastNodeIDIdx := strings.LastIndex(html, "data-node-id=")
			length := len("data-node-id=\"20200813190226-1234567\" ")
			html = html[:lastNodeIDIdx] + html[lastNodeIDIdx+length:]
			lastNodeIDIdx = strings.LastIndex(test.to, "data-node-id=")
			test.to = test.to[:lastNodeIDIdx] + test.to[lastNodeIDIdx+length:]
		}

		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
}
