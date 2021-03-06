// Code generated by "aststring"; DO NOT EDIT.
package ast

import "bytes"

func (node *Script) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("Script (")
	_, _ = buf.WriteString("\n")
	if node.ScriptBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.ScriptBody.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ScriptBody) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ScriptBody (")
	_, _ = buf.WriteString("\n")
	if node.StatementList != nil {
		_, _ = buf.WriteString(PrefixToString(node.StatementList.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}
