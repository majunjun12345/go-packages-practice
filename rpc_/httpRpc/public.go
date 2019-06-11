package rpc1

import "errors"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B // 不甚明白，reply 是　*int，但是　*reply　是？
	return nil
}

func (t *Arith) Divide(args *Args, reply *Quotient) error {
	if args.B == 0 {
		return errors.New("divided by zero!")
	}
	reply.Quo = args.A / args.B
	reply.Rem = args.A % args.B
	return nil
}
