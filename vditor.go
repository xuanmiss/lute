// Lute - A structured markdown engine.
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.

// +build javascript

package lute

// RenderVditorDOM 用于渲染 Vditor DOM，start 和 end 是光标位置，从 0 开始。
func (lute *Lute) RenderVditorDOM(markdownText string, startOffset, endOffset int) (html string, err error) {
	var tree *Tree
	lute.VditorWYSIWYG = true
	markdownText = lute.endNewline(markdownText)
	tree, err = lute.parse("", []byte(markdownText))
	if nil != err {
		return
	}

	renderer := lute.newVditorRenderer(tree.Root)
	startOffset, endOffset = renderer.byteOffset(markdownText, startOffset, endOffset)
	renderer.mapSelection(tree.Root, startOffset, endOffset)
	var output []byte
	output, err = renderer.Render()
	html = string(output)
	return
}

// VditorOperation 用于在 markdownText 中 startOffset、endOffset 指定的选段上做出 operation 操作。
// 支持的 operation 如下：
//   * newline 换行
func (lute *Lute) VditorOperation(markdownText string, startOffset, endOffset int, operation string) (html string, err error) {
	var tree *Tree
	lute.VditorWYSIWYG = true
	markdownText = lute.endNewline(markdownText)
	tree, err = lute.parse("", []byte(markdownText))
	if nil != err {
		return
	}

	renderer := lute.newVditorRenderer(tree.Root)
	startOffset, endOffset = renderer.byteOffset(markdownText, startOffset, endOffset)

	var nodes []*Node
	for c := tree.Root.firstChild; nil != c; c = c.next {
		renderer.findSelection(c, startOffset, endOffset, &nodes)
	}

	if 1 > len(nodes) {
		// 当且仅当渲染空 Markdown 时
		nodes = append(nodes, tree.Root)
	}

	// TODO: 仅实现了光标插入位置节点获取，选段映射待实现

	en := renderer.nearest(nodes, endOffset)
	baseOffset := 0
	if 0 < len(en.tokens) {
		baseOffset = en.tokens[0].Offset()
	}
	endOffset = endOffset - baseOffset

	if NodeThematicBreak == en.typ { // 如果光标所处节点是分隔线节点的话
		en.typ = NodeText // 将分隔线转为文本，后续会按文本节点进行换行处理
		// 构造段落父节点
		paragraph := &Node{typ: NodeParagraph}
		en.parent.AppendChild(paragraph)
		paragraph.AppendChild(en)
	}

	newTree := &Node{typ: en.typ, tokens: en.tokens[endOffset:]} // 生成新的节点树，内容是当前选中节点的后半部分
	en.tokens = en.tokens[:endOffset]                            // 当前选中节点内容为前半部分
	// 保持排版格式并实现换行
	for parent := en.parent; ; parent = parent.parent { // 从当前选中节点开始向父节点方向迭代
		if NodeDocument == parent.typ || NodeList == parent.typ || NodeBlockquote == parent.typ {
			// 遇到这几种块容器说明迭代结束
			break
		}

		newParent := &Node{typ: parent.typ, listData: parent.listData} // 生成新的父节点
		left := true                                                   // 用于标记左边兄弟节点是否迭代结束
		child := parent.firstChild
		for { // 从左到右迭代子节点
			next := child.next                   // AppendChild 会断开，所以这里需要临时保存
			if child == newTree || child == en { // 如果遍历到当前节点
				newParent.AppendChild(newTree) // 将当前节点挂到新的父节点上
				left = false                   // 标记左边兄弟节点迭代结束
			} else if child.isMarker() { // 如果遍历到的是排版标记节点
				newParent.AppendChild(&Node{typ: child.typ, tokens: child.tokens}) // 生成新的标记节点以便保持排版格式
			} else if !left { // 如果遍历到右边兄弟节点
				newParent.AppendChild(child) // 将右边的兄弟节点断开并挂到新的父节点上
			}
			if child = next; nil == child {
				break
			}
		}
		parent.InsertAfter(newParent) // 将新的父节作为当前迭代节点的右兄弟节点挂好
		newTree = newParent           // 设置当前节点以便下一次迭代
	}

	// 进行最终渲染
	var output []byte
	renderer = lute.newVditorRenderer(tree.Root)
	var child *Node
	for child = newTree.firstChild; nil != child.firstChild; child = child.firstChild {
	}
	if 1 > len(child.tokens) {
		child.tokens = items{newItem(itemNewline, 0, 0, 0)}
	}
	child.caretStartOffset = "0"
	child.caretEndOffset = "0"
	renderer.expand(child)
	output, err = renderer.Render()
	html = string(output)
	return
}

// VditorDOMMarkdown 用于将 Vditor DOM 转换为 Markdown 文本。
// TODO：改为解析标准 DOM
func (lute *Lute) VditorDOMMarkdown(html string) (markdown string, err error) {
	tree, err := lute.parseVditorDOM(html)
	if nil != err {
		return
	}

	var formatted []byte
	renderer := lute.newFormatRenderer(tree.Root)
	formatted, err = renderer.Render()
	if nil != err {
		return
	}
	markdown = bytesToStr(formatted)
	return
}
