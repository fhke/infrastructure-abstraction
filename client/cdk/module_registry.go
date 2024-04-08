package cdk

// WithGenericModule adds a generic module to the stack, and returns a builder for the module.
func (s *StackBuilder[R]) WithGenericModule(name, module string) *GenericModuleBuilder {
	newMod := newGenericModuleBuilder(name, module)
	s.AddModuleBuilder(name, newMod)
	return newMod
}

// adds an Aurora module the stack, and returns a builder for the module.
func (s *StackBuilder[R]) WithAurora(name string) *AuroraModuleBuilder {
	mod := newAuroraModuleBuilder(name)
	s.AddModuleBuilder(name, mod)
	return mod
}

// adds an EKS module the stack, and returns a builder for the module.
func (s *StackBuilder[R]) WithEKS(name string) *EKSModuleBuilder {
	mod := newEKSModuleBuilder(name)
	s.AddModuleBuilder(name, mod)
	return mod
}

// adds a VPC module to the stack, and returns a builder for the module.
func (s *StackBuilder[R]) WithVPC(name string) *VPCModuleBuilder {
	mod := newVPCModuleBuilder(name)
	s.AddModuleBuilder(name, mod)
	return mod
}
