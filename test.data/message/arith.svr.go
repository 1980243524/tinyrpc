// Code generated by protoc-gen-tinyrpc. DO NOT EDIT.

package message

// ArithServiceServer You need to define a struct to implement these methods and then call them
type ArithServiceServer interface {
	Add(args *ArithRequest, reply *ArithResponse) error
	Sub(args *ArithRequest, reply *ArithResponse) error
	Mul(args *ArithRequest, reply *ArithResponse) error
	Div(args *ArithRequest, reply *ArithResponse) error
}

// ArithService Defining Computational Digital Services
type ArithService struct{}

// Add addition
func (this *ArithService) Add(args *ArithRequest, reply *ArithResponse) error {
	// define your service ...
	return nil
}

// Sub subtraction
func (this *ArithService) Sub(args *ArithRequest, reply *ArithResponse) error {
	// define your service ...
	return nil
}

// Mul multiplication
func (this *ArithService) Mul(args *ArithRequest, reply *ArithResponse) error {
	// define your service ...
	return nil
}

// Div division
func (this *ArithService) Div(args *ArithRequest, reply *ArithResponse) error {
	// define your service ...
	return nil
}
