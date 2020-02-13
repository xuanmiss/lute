// Code generated by "stringer -type=nodeType"; DO NOT EDIT.

package lute

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NodeDocument-0]
	_ = x[NodeParagraph-1]
	_ = x[NodeHeading-2]
	_ = x[NodeHeadingC8hMarker-3]
	_ = x[NodeThematicBreak-4]
	_ = x[NodeBlockquote-5]
	_ = x[NodeBlockquoteMarker-6]
	_ = x[NodeList-7]
	_ = x[NodeListItem-8]
	_ = x[NodeHTMLBlock-9]
	_ = x[NodeInlineHTML-10]
	_ = x[NodeCodeBlock-11]
	_ = x[NodeCodeBlockFenceOpenMarker-12]
	_ = x[NodeCodeBlockFenceCloseMarker-13]
	_ = x[NodeCodeBlockFenceInfoMarker-14]
	_ = x[NodeCodeBlockCode-15]
	_ = x[NodeText-16]
	_ = x[NodeEmphasis-17]
	_ = x[NodeEmA6kOpenMarker-18]
	_ = x[NodeEmA6kCloseMarker-19]
	_ = x[NodeEmU8eOpenMarker-20]
	_ = x[NodeEmU8eCloseMarker-21]
	_ = x[NodeStrong-22]
	_ = x[NodeStrongA6kOpenMarker-23]
	_ = x[NodeStrongA6kCloseMarker-24]
	_ = x[NodeStrongU8eOpenMarker-25]
	_ = x[NodeStrongU8eCloseMarker-26]
	_ = x[NodeCodeSpan-27]
	_ = x[NodeCodeSpanOpenMarker-28]
	_ = x[NodeCodeSpanContent-29]
	_ = x[NodeCodeSpanCloseMarker-30]
	_ = x[NodeHardBreak-31]
	_ = x[NodeSoftBreak-32]
	_ = x[NodeLink-33]
	_ = x[NodeImage-34]
	_ = x[NodeBang-35]
	_ = x[NodeOpenBracket-36]
	_ = x[NodeCloseBracket-37]
	_ = x[NodeOpenParen-38]
	_ = x[NodeCloseParen-39]
	_ = x[NodeLinkText-40]
	_ = x[NodeLinkDest-41]
	_ = x[NodeLinkTitle-42]
	_ = x[NodeLinkSpace-43]
	_ = x[NodeTaskListItemMarker-100]
	_ = x[NodeStrikethrough-101]
	_ = x[NodeStrikethrough1OpenMarker-102]
	_ = x[NodeStrikethrough1CloseMarker-103]
	_ = x[NodeStrikethrough2OpenMarker-104]
	_ = x[NodeStrikethrough2CloseMarker-105]
	_ = x[NodeTable-106]
	_ = x[NodeTableHead-107]
	_ = x[NodeTableRow-108]
	_ = x[NodeTableCell-109]
	_ = x[NodeEmoji-200]
	_ = x[NodeEmojiUnicode-201]
	_ = x[NodeEmojiImg-202]
	_ = x[NodeEmojiAlias-203]
	_ = x[NodeMathBlock-300]
	_ = x[NodeMathBlockOpenMarker-301]
	_ = x[NodeMathBlockContent-302]
	_ = x[NodeMathBlockCloseMarker-303]
	_ = x[NodeInlineMath-304]
	_ = x[NodeInlineMathOpenMarker-305]
	_ = x[NodeInlineMathContent-306]
	_ = x[NodeInlineMathCloseMarker-307]
	_ = x[NodeBackslash-400]
	_ = x[NodeBackslashContent-401]
	_ = x[NodeFootnotesDef-500]
	_ = x[NodeFootnotesRef-501]
	_ = x[NodeToC-600]
}

const (
	_nodeType_name_0 = "NodeDocumentNodeParagraphNodeHeadingNodeHeadingC8hMarkerNodeThematicBreakNodeBlockquoteNodeBlockquoteMarkerNodeListNodeListItemNodeHTMLBlockNodeInlineHTMLNodeCodeBlockNodeCodeBlockFenceOpenMarkerNodeCodeBlockFenceCloseMarkerNodeCodeBlockFenceInfoMarkerNodeCodeBlockCodeNodeTextNodeEmphasisNodeEmA6kOpenMarkerNodeEmA6kCloseMarkerNodeEmU8eOpenMarkerNodeEmU8eCloseMarkerNodeStrongNodeStrongA6kOpenMarkerNodeStrongA6kCloseMarkerNodeStrongU8eOpenMarkerNodeStrongU8eCloseMarkerNodeCodeSpanNodeCodeSpanOpenMarkerNodeCodeSpanContentNodeCodeSpanCloseMarkerNodeHardBreakNodeSoftBreakNodeLinkNodeImageNodeBangNodeOpenBracketNodeCloseBracketNodeOpenParenNodeCloseParenNodeLinkTextNodeLinkDestNodeLinkTitleNodeLinkSpace"
	_nodeType_name_1 = "NodeTaskListItemMarkerNodeStrikethroughNodeStrikethrough1OpenMarkerNodeStrikethrough1CloseMarkerNodeStrikethrough2OpenMarkerNodeStrikethrough2CloseMarkerNodeTableNodeTableHeadNodeTableRowNodeTableCell"
	_nodeType_name_2 = "NodeEmojiNodeEmojiUnicodeNodeEmojiImgNodeEmojiAlias"
	_nodeType_name_3 = "NodeMathBlockNodeMathBlockOpenMarkerNodeMathBlockContentNodeMathBlockCloseMarkerNodeInlineMathNodeInlineMathOpenMarkerNodeInlineMathContentNodeInlineMathCloseMarker"
	_nodeType_name_4 = "NodeBackslashNodeBackslashContent"
	_nodeType_name_5 = "NodeFootnotesDefNodeFootnotesRef"
	_nodeType_name_6 = "NodeToC"
)

var (
	_nodeType_index_0 = [...]uint16{0, 12, 25, 36, 56, 73, 87, 107, 115, 127, 140, 154, 167, 195, 224, 252, 269, 277, 289, 308, 328, 347, 367, 377, 400, 424, 447, 471, 483, 505, 524, 547, 560, 573, 581, 590, 598, 613, 629, 642, 656, 668, 680, 693, 706}
	_nodeType_index_1 = [...]uint8{0, 22, 39, 67, 96, 124, 153, 162, 175, 187, 200}
	_nodeType_index_2 = [...]uint8{0, 9, 25, 37, 51}
	_nodeType_index_3 = [...]uint8{0, 13, 36, 56, 80, 94, 118, 139, 164}
	_nodeType_index_4 = [...]uint8{0, 13, 33}
	_nodeType_index_5 = [...]uint8{0, 16, 32}
)

func (i nodeType) String() string {
	switch {
	case 0 <= i && i <= 43:
		return _nodeType_name_0[_nodeType_index_0[i]:_nodeType_index_0[i+1]]
	case 100 <= i && i <= 109:
		i -= 100
		return _nodeType_name_1[_nodeType_index_1[i]:_nodeType_index_1[i+1]]
	case 200 <= i && i <= 203:
		i -= 200
		return _nodeType_name_2[_nodeType_index_2[i]:_nodeType_index_2[i+1]]
	case 300 <= i && i <= 307:
		i -= 300
		return _nodeType_name_3[_nodeType_index_3[i]:_nodeType_index_3[i+1]]
	case 400 <= i && i <= 401:
		i -= 400
		return _nodeType_name_4[_nodeType_index_4[i]:_nodeType_index_4[i+1]]
	case 500 <= i && i <= 501:
		i -= 500
		return _nodeType_name_5[_nodeType_index_5[i]:_nodeType_index_5[i+1]]
	case i == 600:
		return _nodeType_name_6
	default:
		return "nodeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
