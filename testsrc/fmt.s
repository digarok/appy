* formatting test
                  org   $300                    ; start

main              nop

                  * what about this

                                                ; and this

                  sta   :jo+1


:jo+1             lda   $400
                  lda   $400
                  rts


ReallyThisisaLoooooongLabelwith stal $e12000,y  ; look at this long line
