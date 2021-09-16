package code

func Execute(client *identity.Client, fn *Function, arguments []string) ([]Result, []SideEffect, error)
