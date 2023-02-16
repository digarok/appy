* formatting test
        org   $300 ; start

main nop
 * what about this

            ; and this

  sta   :jo+1


:jo+1   lda $400
             lda   $400
                 rts

** 24 (bit) hex to 8 (nibble) / 4 byte BCD
** 24 (bit) hex to
HEXDEC	mx %11
                LDA #0		; Ensure the result is clear
                STA DEC8+0
                STA DEC8+1
                STA DEC8+2
                STA DEC8+3
ReallyThisisaLoooooongLabelwith stal $e12000,y  ; look at this long line

* TABS....
* $D5	$0008	 sequence [Application Specific]
                                                ;  ed: mc=40 oc=26 cc=48 ms=5 bs=2   <-  indentation modeline
