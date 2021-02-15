package net

var RegexpCenter = `(.+?)`

func (*ObjNet) RegexpWebBodyBlocks(tagName string) string {
	return `<` + tagName + `[^>]*?>[\w\W]*?<\/` + tagName + `>`
}
