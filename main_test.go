package main

import "testing"

func TestExecA(t *testing.T) {
	exp := "1+2"
	exec(exp)
}

func TestExecB(t *testing.T) {
	exp := "1+2-4"
	exec(exp)
}

func TestExecC(t *testing.T) {
	exp := "1+2-4*3-8"
	exec(exp)
}

func TestExecD(t *testing.T) {
	exp := "1+2-(4*3-8)"
	exec(exp)
}

func TestExecE(t *testing.T) {
	exp := "1+2-(4*3+(1-8))"
	exec(exp)
}

func TestExecF(t *testing.T) {
	exp := "1+(2-(4*3+(1-8)))"
	exec(exp)
}

func TestExecG(t *testing.T) {
	exp := "((1-2)*(3-8))*((((9+2222))))"
	exec(exp)
}

func TestExecH(t *testing.T) {
	exp := "0.8888-0.1 * 444         -0.2"
	exec(exp)
}

func TestExecI(t *testing.T) {
	exp := "0.8888-0.1 * (444         -0.2)"
	exec(exp)
}
