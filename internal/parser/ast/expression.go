package ast

type Expression struct {
	AssignmentExpressions []*AssignmentExpression
}

type AssignmentExpression struct {
	ConditionalExpression  *ConditionalExpression
	YieldExpression        *YieldExpression
	ArrowFunction          *ArrowFunction
	AsyncArrowFunction     *AsyncArrowFunction
	LeftHandSideExpression *LeftHandSideExpression
	Assign                 bool
	AssignmentOperator     string
	AssignmentExpression   *AssignmentExpression
}

type ConditionalExpression struct {
	LogicalORExpression   *LogicalORExpression
	AssignmentExpression1 *AssignmentExpression
	AssignmentExpression2 *AssignmentExpression
}

type LogicalORExpression struct {
	LogicalANDExpression *LogicalANDExpression
	LogicalORExpression  *LogicalORExpression
}

type LogicalANDExpression struct {
	BitwiseORExpression  *BitwiseORExpression
	LogicalANDExpression *LogicalANDExpression
}

type BitwiseORExpression struct {
	BitwiseORExpression  *BitwiseORExpression
	BitwiseXORExpression *BitwiseXORExpression
}

type BitwiseANDExpression struct {
	BitwiseANDExpression *BitwiseANDExpression
	EqualityExpression   *EqualityExpression
}

type BitwiseXORExpression struct {
	BitwiseANDExpression *BitwiseANDExpression
	BitwiseXORExpression *BitwiseXORExpression
}

type EqualityExpression struct {
	EqualityExpression   *EqualityExpression
	RelationalExpression *RelationalExpression
	Equals               bool
	StrictEquals         bool
	NotEquals            bool
	StrictNotEquals      bool
}

type RelationalExpression struct {
	ShiftExpression      *ShiftExpression
	RelationalExpression *RelationalExpression
	LessThan             bool
	GreaterThan          bool
	LessThanOrEqualTo    bool
	GreaterThanOrEqualTo bool
	Instanceof           bool
	In                   bool
}

type ShiftExpression struct {
	ShiftExpression    *ShiftExpression
	AdditiveExpression *AdditiveExpression
	LeftShift          bool
	RightShift         bool
	UnsignedRightShift bool
}

type AdditiveExpression struct {
	MultiplicativeExpression *MultiplicativeExpression
	AdditiveExpression       *AdditiveExpression
	Plus                     bool
	Minus                    bool
}

type MultiplicativeExpression struct {
	ExponentiationExpression *ExponentiationExpression
	MultiplicativeExpression *MultiplicativeExpression
	Asterisk                 bool
	Slash                    bool
	Modulo                   bool
}

type ExponentiationExpression struct {
	UnaryExpression          *UnaryExpression
	UpdateExpression         *UpdateExpression
	ExponentiationExpression *ExponentiationExpression
}

type UpdateExpression struct {
	LeftHandSideExpression *LeftHandSideExpression
	UnaryExpression        *UnaryExpression
	PlusPlus               bool
	MinusMinus             bool
}

type UnaryExpression struct {
	UpdateExpression *UpdateExpression
	UnaryExpression  *UnaryExpression
	AwaitExpression  *AwaitExpression
	Delete           bool
	Void             bool
	Typeof           bool
	Plus             bool
	Minus            bool
	Tilde            bool
	ExclamationMark  bool
}

type AwaitExpression struct {
	UnaryExpression *UnaryExpression
}

type YieldExpression struct {
	AssignmentExpression *AssignmentExpression
	Asterisk             bool
}

type LeftHandSideExpression struct {
	NewExpression  *NewExpression
	CallExpression *CallExpression
}

type NewExpression struct {
	MemberExpression *MemberExpression
	NewExpression    *NewExpression
}

type MemberExpression struct {
	PrimaryExpression *PrimaryExpression
	MemberExpression  *MemberExpression
	Expression        *Expression
	IdentifierName    string
	TemplateLiteral   *TemplateLiteral
	SuperProperty     *SuperProperty
	MetaProperty      *MetaProperty
	Arguments         *Arguments
}

type PrimaryExpression struct {
	This                                              bool
	IdentifierReference                               *IdentifierReference
	Literal                                           *Literal
	ArrayLiteral                                      *ArrayLiteral
	ObjectLiteral                                     *ObjectLiteral
	FunctionExpression                                *FunctionExpression
	ClassExpression                                   *ClassExpression
	GeneratorExpression                               *GeneratorExpression
	AsyncFunctionExpression                           *AsyncFunctionExpression
	AsyncGeneratorExpression                          *AsyncGeneratorExpression
	RegularExpressionLiteral                          *RegularExpressionLiteral
	TemplateLiteral                                   *TemplateLiteral
	CoverParenthesizedExpressionAndArrowParameterList *CoverParenthesizedExpressionAndArrowParameterList
}

type FunctionExpression struct {
	BindingIdentifier *BindingIdentifier
	FormalParameters  *FormalParameters
	FunctionBody      *FunctionBody
}

type ClassExpression struct {
	BindingIdentifier *BindingIdentifier
	ClassTail         *ClassTail
}

type ClassTail struct {
	ClassHeritage *ClassHeritage
	ClassBody     *ClassBody
}

type ClassHeritage struct {
	LeftHandSideExpression *LeftHandSideExpression
}

type ClassBody struct {
	ClassElementList *ClassElementList
}

type ClassElementList struct {
	ClassElements []*ClassElement
}

type ClassElement struct {
	MethodDefinition *MethodDefinition
	Static           bool
}

type GeneratorExpression struct {
	BindingIdentifier *BindingIdentifier
	FormalParameters  *FormalParameters
	GeneratorBody     *GeneratorBody
}

type AsyncFunctionExpression struct {
	FormalParameters  *FormalParameters
	AsyncFunctionBody *AsyncFunctionBody
	BindingIdentifier *BindingIdentifier
}

type AsyncGeneratorExpression struct {
	BindingIdentifier  *BindingIdentifier
	FormalParameters   *FormalParameters
	AsyncGeneratorBody *AsyncGeneratorBody
}

type CallExpression struct {
	CoverCallExpressionAndAsyncArrowHead *CoverCallExpressionAndAsyncArrowHead
	SuperCall                            *SuperCall
	CallExpression                       *CallExpression
	Arguments                            *Arguments
	Expression                           *Expression
	IdentifierName                       string
	TemplateLiteral                      *TemplateLiteral
}

type SuperCall struct {
	Arguments *Arguments
}
