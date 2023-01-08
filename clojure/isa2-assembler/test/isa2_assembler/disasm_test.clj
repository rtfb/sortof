(ns isa2-assembler.disasm-test
  (:require [clojure.test :refer :all]
            [isa2-assembler.disasm :refer :all]))

(deftest test-disassembler-returns-null
  (testing "The disassembler entry point returns null"
    (is (= nil (run [0x9a])))))
