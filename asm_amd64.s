// +build !race

#define NOSPLIT 4
#define MAXSPIN 1000

// Lock(l *uint32)
TEXT ·Lock(SB),NOSPLIT,$0-8
begin:
	MOVQ 	l+0(FP), BP
	MOVL	$MAXSPIN, DX
	MOVL	0(BP), CX
	TESTL	CX, CX
	JNZ		spin
acquire:
	MOVL 	$1, AX
	XCHGL 	AX, 0(BP)
	TESTL 	AX, AX
	JNZ 	spin
	RET
spin:
	PAUSE
	MOVL   	0(BP), CX	// spin on dirty read
	TESTL  	CX, CX
	JZ 		acquire
	DECL	DX 
	JNZ 	spin
sched:
	CALL 	runtime·Gosched(SB)	// bailout
	JMP		begin


// Unlock(l *uint32)
TEXT ·Unlock(SB),NOSPLIT,$0-8
	MOVQ 	l+0(FP), BP
	XORL    AX, AX
	MOVL    AX, 0(BP)
	RET


// TryLock(l *uint32) bool
TEXT ·TryLock(SB),NOSPLIT,$0-9
	MOVQ	l+0(FP), BP
	MOVL	0(BP), CX
	TESTL	CX, CX
	JZ		try
	XORQ	AX, AX
	MOVB	AX,swapped+8(FP)
	RET
try:
	MOVL	$1, AX
	XCHGL 	AX, 0(BP)
	TESTL	AX, AX
	SETEQ	swapped+8(FP)
	RET

