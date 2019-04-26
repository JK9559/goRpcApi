package api

type Args struct {
	A, B int
}

type Arith struct {
}

type ArithReq struct {
	A int
	B int
}

type ArithRsp struct {
	Pro int
	Quo int
	Rem int
}
