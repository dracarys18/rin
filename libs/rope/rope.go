package rope

type RopeInterface interface {
	Concatenate
}

type Concatenate interface {
	Concatenate(Rope) Concatenate
}

type Substring interface {
	Substring(int) Substring
}

type Rope struct {
	data   string
	length int
	left   *Rope
	right  *Rope
}

func NewRope(data string) Rope {
	return Rope{
		data:   data,
		length: len(data),
	}
}

func (rope Rope) IsLeaf() bool {
	return rope.left == nil && rope.right == nil
}

func (rope Rope) String() string {
	if rope.IsLeaf() {
		return rope.data
	} else {
		return rope.left.String() + rope.right.String()
	}
}

func (rope Rope) Eq(other Rope) bool {
	return rope.String() == other.String()
}

func (rope Rope) Concatenate(second Rope) Rope {
	if rope.length == 0 {
		return second
	} else if second.length == 0 {
		return rope
	} else {
		return Rope{
			length: rope.length + second.length,
			left:   &rope,
			right:  &second,
		}
	}
}
