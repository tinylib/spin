// +build !race

#define NOSPLIT 4

// Lock(l *uint32)
TEXT ·Lock(SB),NOSPLIT,$0-8
	MOVQ   l+0(FP), BP
spin:
	MOVL   $1, AX     
	XCHGL  AX, 0(BP)
	TESTL  AX, AX
	JNZ    wait
	RET
wait:
	PAUSE
	JMP    spin

// Unlock(l *uint32)
TEXT ·Unlock(SB),NOSPLIT,$0-8
	MOVQ 	l+0(FP), BP
	XORL    AX, AX
	MOVL    AX, 0(BP)
	RET

// TryLock(l *uint32) bool
TEXT ·TryLock(SB),NOSPLIT,$0-9
	MOVQ	l+0(FP), BP
	MOVL	$1, AX
	XCHGL 	AX, 0(BP)
	TESTL	AX, AX
	SETEQ	swapped+8(FP)
	RET

