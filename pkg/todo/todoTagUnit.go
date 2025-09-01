package todo

import "errors"

func (tagPointer *TodoTagUnit) getTagUnit(tagID int) (err error) {
	switch tagID {
	case 0:
		tagPointer.Title = "Home"
		tagPointer.Class = "Green"
	case 1:
		tagPointer.Title = "Work"
		tagPointer.Class = "Red"
	case 2:
		tagPointer.Title = "Default"
		tagPointer.Class = "Blue"
	default:
		return errors.New("unknow TagID 'getTagUnit 17'")
	}
	return
}
