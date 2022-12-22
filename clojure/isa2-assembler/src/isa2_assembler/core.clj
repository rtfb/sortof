(ns isa2-assembler.core
  (:require [isa2-assembler.disasm :refer :all])
  (:gen-class))

(defn disassemble [file]
  (do
    (println "hello from disassembler")
    (println (isa2-assembler.disasm/do [0x91]))))

(defn assemble [file]
  (println "hello from assembler"))

(defn -main
  "Calls the disassembler if -d flag is present or assembler otherwise."
  [& args]
  (if (empty? args)
    (do
      (println "Not enough arguments: need at least filename")
      (System/exit 1))
    (if (= (first args) "-d")
      (disassemble (second args))
      (assemble (first args)))))
