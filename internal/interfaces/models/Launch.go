package models

type Launch struct {
	Id         string `dynamo:"id"`
	Name       string `dynamo:"name"`
	FailReason string `dynamo:"failReason"`
}
