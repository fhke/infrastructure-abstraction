package controller

import "context"

func (c *controllerImpl) GetModuleNames(ctx context.Context) ([]string, error) {
	return c.moduleRepo.GetModuleNames(ctx)
}
