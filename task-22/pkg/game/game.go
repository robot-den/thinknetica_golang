package game

import "task-22/pkg/model"

type Calculator interface {
	Setup(string)
	Iterate(int)
	Result() *model.CalculationsData
}
