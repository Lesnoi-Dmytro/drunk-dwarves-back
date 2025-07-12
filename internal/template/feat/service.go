package feat

const ServiceFileTemplate = `package %[1]s

func get%[2]s() []%[3]sDto {
	%[1]s := []%[3]sDto{}[:]

	return %[1]s
}
`
