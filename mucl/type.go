package mucl

func (t *Type) String() string {
	if t == nil {
		return ""
	}
	if t.Scalar > 0 {
		return t.Scalar.GoString()
	}
	if t.Reference != "" {
		return t.Reference
	}

	return ""
}
