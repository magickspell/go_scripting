package main

type ParentInterface interface {
	ParentMethod(in string) (out string)
}

type ChildInterface interface { // ChildInterface теперь имеет свой метод и метод ParentInterface
	ParentInterface
	SomeChildMethod(in string) (out string)
}
