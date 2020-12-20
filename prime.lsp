;; 生成序列到管道1 2 3 ...
(setq sequence
      (lambda (c min)
        (begin
         (-> c min)
         (sequence c (+ min 1)))))

;; 遍历打印管道
(setq printch
      (lambda (c max)
        (begin
         (setq num (<- c))
         (display num)
         (if (< num max)
             (printch c max)))))

;; 从管道中过滤num 其他数字输出到管道out
(setq filter
      (lambda (out in num)
        (begin
         (setq new (<- in))
         (if (not (eq 0 (mod new num)))
             (-> out new))
         (filter out in num))))

;; in 输入整数序列，out输出素数序列
(setq primefilter
      (lambda (out in)
        (begin
         (setq cur (<- in))
         (setq newout (ch))
         (go filter newout in cur)
         (-> out cur )
         (primefilter out newout))))

(setq out (ch))
(setq numbers (ch))
(go sequence numbers 2)
(go primefilter out numbers)
(printch out 1000)
