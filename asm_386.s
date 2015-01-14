// +build !race

#define NOSPLIT 4
#define MAXSPIN 1000

TEXT ·Lock(SB),NOSPLIT,$0-4
begin:
	MOVL 	l+0(SB), BP
	MOVL	$MAXSPIN, DX
	MOVL	0(BP), CX
	TESTL	CX, CX
	JNZ		spin
acquire:
	MOVL	$1, AX
	XCHGL 	AX, 0(BP)
	TESTL 	AX, AX
	JNZ   	spin
	RET
spin:
	PAUSE
	MOVL	0(BP), CX
	TESTL	CX, CX
	JZ		acquire	
	DECL	DX
	JNZ 	spin
sched:
	CALL	runtime.Gosched(SB)
	JMP		begin


TEXT ·Unlock(SB),NOSPLIT,$0-4
	MOVL  l+0(SB), BP
	XORL  AX, AX
	XCHGL AX, 0(BP)
	RET


TEXT ·TryLock(SB),NOSPLIT,$0-5
	MOVL	l+0(FP), BP
	MOVL	$1, AX
	XCHGL 	AX, 0(BP)
	TESTL	AX, AX
	SETEQ	swapped+4(FP)
	RET
