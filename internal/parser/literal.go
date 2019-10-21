package parser

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/token"
)

func parseTemplateLiteral(i *isolate, p param) *ast.TemplateLiteral {
	chck := i.checkpoint()

	if t, ok := i.accept(token.NoSubstitutionTemplate); ok {
		return &ast.TemplateLiteral{
			NoSubstitutionTemplate: t.Value,
		}
	}

	if substitutionTemplate := parseSubstitutionTemplate(i, p.only(pYield|pAwait|pTagged)); substitutionTemplate != nil {
		return &ast.TemplateLiteral{
			SubstitutionTemplate: substitutionTemplate,
		}
	}

	i.restore(chck)
	return nil
}

func parseSubstitutionTemplate(i *isolate, p param) *ast.SubstitutionTemplate {
	chck := i.checkpoint()

	var templateHead string
	var expr *ast.Expression
	var templateSpans *ast.TemplateSpans

	t, ok := i.accept(token.TemplateHead)
	if !ok {
		i.restore(chck)
		return nil
	}
	templateHead = t.Value

	if expr = parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr == nil {
		i.restore(chck)
		return nil
	}

	if templateSpans = parseTemplateSpans(i, p.only(pYield|pAwait|pTagged)); templateSpans == nil {
		i.restore(chck)
		return nil
	}

	return &ast.SubstitutionTemplate{
		TemplateHead:  templateHead,
		Expression:    expr,
		TemplateSpans: templateSpans,
	}
}

func parseNumericLiteral(i *isolate, p param) *ast.NumericLiteral {
	chck := i.checkpoint()

	if t, ok := i.accept(token.DecimalLiteral); ok {
		return &ast.NumericLiteral{
			DecimalLiteral: t.Value,
		}
	} else if t, ok := i.accept(token.BinaryIntegerLiteral); ok {
		return &ast.NumericLiteral{
			BinaryIntegerLiteral: t.Value,
		}
	} else if t, ok := i.accept(token.OctalIntegerLiteral); ok {
		return &ast.NumericLiteral{
			OctalIntegerLiteral: t.Value,
		}
	} else if t, ok := i.accept(token.HexIntegerLiteral); ok {
		return &ast.NumericLiteral{
			HexIntegerLiteral: t.Value,
		}
	}

	i.restore(chck)
	return nil
}