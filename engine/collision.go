package engine

func (e *Engine) checkCollision() {
	onCollisions := e.getOnCollisions()
	hasCollisions := e.getHasCollisions()
	for _, collision := range onCollisions {
		cords := collision.CheckForCollisionCords()
		for _, hasCollision := range hasCollisions {
			if hasCollision.HasCollision(cords) {
				collision.OnCollision(hasCollision)
			}
		}
	}
}

func (e *Engine) getOnCollisions() []OnCollision {
	var onCollisions []OnCollision
	for _, object := range e.objects {
		c, ok := object.(OnCollision)
		if ok {
			onCollisions = append(onCollisions, c)
		}
	}
	return onCollisions
}

func (e *Engine) getHasCollisions() []HasCollision {
	var hasCollisions []HasCollision
	for _, object := range e.objects {
		c, ok := object.(HasCollision)
		if ok {
			hasCollisions = append(hasCollisions, c)
		}
	}
	return hasCollisions
}
