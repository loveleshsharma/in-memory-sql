package table

type Constraint interface {
	String() string
}

type NotNullConstraint struct {
}

func (n NotNullConstraint) String() string {
	return "not null constraint"
}

type UniqueConstraint struct {
}

func (n UniqueConstraint) String() string {
	return "unique constraint"
}

var NotNullConstraintType NotNullConstraint
var UniqueConstraintType UniqueConstraint
