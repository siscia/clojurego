To implement all the clojure data structure I believe that the way to go is via the clojure protocols.

Define all the clojure protocols as Go interface and then move from there.

Once we have all the interface we just need to implement the real data structures, those does not need to be fast, they just need to work.