(ns clojure.examples.hello
  (:gen-class))

(defn lat [])
;; This program displays Hello World
(defn Example []
  (def string1 (slurp "/etc/passwd"))
  (println (empty? ()))
  (println (first '(1 2 3)))
  (println (rest '(1 2 3)))
  (println '(a b c))
  (println (quote (println "hello")))
  (println *ns*))

;; go
(Example)