package util

// 存储目前的权限状态
type BinaryCheck struct {
	Flag int
}

// 设置状态
func (c *BinaryCheck) SetStatus(status int) {
	c.Flag = status
}

// 添加一种或多种状态
func (c *BinaryCheck) AddStatus(status int) {
	c.Flag |= status
}

// 删除一种或者多种状态
func (c *BinaryCheck) DeleteStatus(status int) {
	/**
	go 不支持取反符号~
	c.Flag &= ~status
	取反 ^status
	*/
	c.Flag &= ^status
}

// 是否具有某些状态
func (c *BinaryCheck) HasStatus(status int) bool {
	return (c.Flag & status) == status
}

// 是否不具有某些状态
func (c *BinaryCheck) NotHasStatus(status int) bool {
	return (c.Flag & status) == 0
}

// 是否仅仅具有某些状态
func (c *BinaryCheck) OnlyHas(status int) bool {
	return c.Flag == status
}
