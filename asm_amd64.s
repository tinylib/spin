// +build !race

#define NOSPLIT 4

TEXT ·Lock(SB),NOSPLIT,$0-8
	MOVQ   l+0(FP), BP		// BP = l

spin:
	MOVL   $1, AX			// AX = 1      
	XCHGL  AX, 0(BP)        // swap(AX, *BP)
	TESTL  AX, AX			// if AX != 0
	JNZ    wait				// goto wait
	RET						// done

wait:
	PAUSE					// pause
	JMP    spin				// try again


TEXT ·Unlock(SB),NOSPLIT,$0-8
	MOVQ 	l+0(FP), BP		// BP = l
	XORL    AX, AX			// AX = 0
	MOVL    AX, 0(BP)		// set *l=0
	RET						// return

