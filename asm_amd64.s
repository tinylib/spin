// +build !race

#define NOSPLIT 4

TEXT ·Lock(SB),NOSPLIT,$0-8
	MOVQ   l+0(FP), BP

spin:
	MOVL   $1, AX
	LOCK          
	XCHGL  AX, 0(BP)        
	TESTL  AX, AX
	JNZ    wait
	RET

wait:
	PAUSE
	JMP    spin


TEXT ·Unlock(SB),NOSPLIT,$0-8
	MOVQ	l+0(FP), BP
	MOVL    $0, AX
	LOCK
	XCHGL   AX, 0(BP)
	RET

