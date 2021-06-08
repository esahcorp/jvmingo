package heap

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (dptor *MethodDescriptor) addParameterType(t string) {
	pLen := len(dptor.parameterTypes)
	if pLen == cap(dptor.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, dptor.parameterTypes)
		dptor.parameterTypes = s
	}

	dptor.parameterTypes = append(dptor.parameterTypes, t)
}
