// Code generated by "aststring"; DO NOT EDIT.
package ast

import (
	"bytes"
	"fmt"
)

func (node *StatementList) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("StatementList (")
	_, _ = buf.WriteString("\n")
	for _, elem := range node.StatementListItems {
		_, _ = buf.WriteString(PrefixToString(elem.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *StatementListItem) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("StatementListItem (")
	_, _ = buf.WriteString("\n")
	if node.Statement != nil {
		_, _ = buf.WriteString(PrefixToString(node.Statement.String(), "  "))
	}
	if node.Declaration != nil {
		_, _ = buf.WriteString(PrefixToString(node.Declaration.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *Statement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("Statement (")
	_, _ = buf.WriteString("\n")
	if node.BlockStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.BlockStatement.String(), "  "))
	}
	if node.VariableStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.VariableStatement.String(), "  "))
	}
	if node.EmptyStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.EmptyStatement.String(), "  "))
	}
	if node.ExpressionStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.ExpressionStatement.String(), "  "))
	}
	if node.IfStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.IfStatement.String(), "  "))
	}
	if node.BreakableStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.BreakableStatement.String(), "  "))
	}
	if node.ContinueStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.ContinueStatement.String(), "  "))
	}
	if node.BreakStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.BreakStatement.String(), "  "))
	}
	if node.ReturnStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.ReturnStatement.String(), "  "))
	}
	if node.WithStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.WithStatement.String(), "  "))
	}
	if node.LabelledStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.LabelledStatement.String(), "  "))
	}
	if node.ThrowStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.ThrowStatement.String(), "  "))
	}
	if node.TryStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.TryStatement.String(), "  "))
	}
	if node.DebuggerStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.DebuggerStatement.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *BlockStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("BlockStatement (")
	_, _ = buf.WriteString("\n")
	if node.Block != nil {
		_, _ = buf.WriteString(PrefixToString(node.Block.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *Block) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("Block (")
	_, _ = buf.WriteString("\n")
	if node.StatementList != nil {
		_, _ = buf.WriteString(PrefixToString(node.StatementList.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *VariableStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("VariableStatement (")
	_, _ = buf.WriteString("\n")
	if node.VariableDeclarationList != nil {
		_, _ = buf.WriteString(PrefixToString(node.VariableDeclarationList.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *EmptyStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("EmptyStatement (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ExpressionStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ExpressionStatement (")
	_, _ = buf.WriteString("\n")
	if node.Expression != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *IfStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("IfStatement (")
	_, _ = buf.WriteString("\n")
	if node.Expression != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression.String(), "  "))
	}
	if node.Statement != nil {
		_, _ = buf.WriteString(PrefixToString(node.Statement.String(), "  "))
	}
	if node.ElseStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.ElseStatement.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *BreakableStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("BreakableStatement (")
	_, _ = buf.WriteString("\n")
	if node.IterationStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.IterationStatement.String(), "  "))
	}
	if node.SwitchStatement != nil {
		_, _ = buf.WriteString(PrefixToString(node.SwitchStatement.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *IterationStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("IterationStatement (")
	_, _ = buf.WriteString("\n")
	if node.Statement != nil {
		_, _ = buf.WriteString(PrefixToString(node.Statement.String(), "  "))
	}
	if node.Expression1 != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression1.String(), "  "))
	}
	if node.Expression2 != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression2.String(), "  "))
	}
	if node.Expression3 != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression3.String(), "  "))
	}
	if node.VariableDeclarationList != nil {
		_, _ = buf.WriteString(PrefixToString(node.VariableDeclarationList.String(), "  "))
	}
	if node.LexicalDeclaration != nil {
		_, _ = buf.WriteString(PrefixToString(node.LexicalDeclaration.String(), "  "))
	}
	if node.LeftHandSideExpression != nil {
		_, _ = buf.WriteString(PrefixToString(node.LeftHandSideExpression.String(), "  "))
	}
	if node.ForBinding != nil {
		_, _ = buf.WriteString(PrefixToString(node.ForBinding.String(), "  "))
	}
	if node.ForDeclaration != nil {
		_, _ = buf.WriteString(PrefixToString(node.ForDeclaration.String(), "  "))
	}
	if node.AssignmentExpression != nil {
		_, _ = buf.WriteString(PrefixToString(node.AssignmentExpression.String(), "  "))
	}
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Do", node.Do), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "While", node.While), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "For", node.For), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "In", node.In), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Var", node.Var), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Of", node.Of), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Await", node.Await), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *SwitchStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("SwitchStatement (")
	_, _ = buf.WriteString("\n")
	if node.Expression != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression.String(), "  "))
	}
	if node.CaseBlock != nil {
		_, _ = buf.WriteString(PrefixToString(node.CaseBlock.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *CaseBlock) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("CaseBlock (")
	_, _ = buf.WriteString("\n")
	for _, elem := range node.CaseClauses {
		_, _ = buf.WriteString(PrefixToString(elem.String(), "  "))
	}
	if node.DefaultClause != nil {
		_, _ = buf.WriteString(PrefixToString(node.DefaultClause.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *CaseClause) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("CaseClause (")
	_, _ = buf.WriteString("\n")
	if node.Expression != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression.String(), "  "))
	}
	if node.StatementList != nil {
		_, _ = buf.WriteString(PrefixToString(node.StatementList.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *DefaultClause) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("DefaultClause (")
	_, _ = buf.WriteString("\n")
	if node.StatementList != nil {
		_, _ = buf.WriteString(PrefixToString(node.StatementList.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ContinueStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ContinueStatement (")
	_, _ = buf.WriteString("\n")
	if node.LabelIdentifier != nil {
		_, _ = buf.WriteString(PrefixToString(node.LabelIdentifier.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *BreakStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("BreakStatement (")
	_, _ = buf.WriteString("\n")
	if node.LabelIdentifier != nil {
		_, _ = buf.WriteString(PrefixToString(node.LabelIdentifier.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ReturnStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ReturnStatement (")
	_, _ = buf.WriteString("\n")
	if node.Expression != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *WithStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("WithStatement (")
	_, _ = buf.WriteString("\n")
	if node.Expression != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression.String(), "  "))
	}
	if node.Statement != nil {
		_, _ = buf.WriteString(PrefixToString(node.Statement.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *LabelledStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("LabelledStatement (")
	_, _ = buf.WriteString("\n")
	if node.LabelIdentifier != nil {
		_, _ = buf.WriteString(PrefixToString(node.LabelIdentifier.String(), "  "))
	}
	if node.LabelledItem != nil {
		_, _ = buf.WriteString(PrefixToString(node.LabelledItem.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *LabelledItem) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("LabelledItem (")
	_, _ = buf.WriteString("\n")
	if node.Statement != nil {
		_, _ = buf.WriteString(PrefixToString(node.Statement.String(), "  "))
	}
	if node.FunctionDeclaration != nil {
		_, _ = buf.WriteString(PrefixToString(node.FunctionDeclaration.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ThrowStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ThrowStatement (")
	_, _ = buf.WriteString("\n")
	if node.Expression != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *TryStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("TryStatement (")
	_, _ = buf.WriteString("\n")
	if node.Block != nil {
		_, _ = buf.WriteString(PrefixToString(node.Block.String(), "  "))
	}
	if node.Catch != nil {
		_, _ = buf.WriteString(PrefixToString(node.Catch.String(), "  "))
	}
	if node.Finally != nil {
		_, _ = buf.WriteString(PrefixToString(node.Finally.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *Catch) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("Catch (")
	_, _ = buf.WriteString("\n")
	if node.CatchParameter != nil {
		_, _ = buf.WriteString(PrefixToString(node.CatchParameter.String(), "  "))
	}
	if node.Block != nil {
		_, _ = buf.WriteString(PrefixToString(node.Block.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *CatchParameter) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("CatchParameter (")
	_, _ = buf.WriteString("\n")
	if node.BindingIdentifier != nil {
		_, _ = buf.WriteString(PrefixToString(node.BindingIdentifier.String(), "  "))
	}
	if node.BindingPattern != nil {
		_, _ = buf.WriteString(PrefixToString(node.BindingPattern.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *Finally) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("Finally (")
	_, _ = buf.WriteString("\n")
	if node.Block != nil {
		_, _ = buf.WriteString(PrefixToString(node.Block.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *DebuggerStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("DebuggerStatement (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}