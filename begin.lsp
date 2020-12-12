(setq double
      (lambda (x)
        (begin
         (setq y (* x 20))
         (display y)
         y)))
(double 33)
