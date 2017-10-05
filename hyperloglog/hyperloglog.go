package hyperloglog

import(

)

/**
 * The hyperloglog sketch can be used to estimate the cardinality of a multiset
 * while using space O(loglog n) where n is the cardinality of the set
 */
type Hyperloglog struct {

}

func (hll *Hyperloglog) Add (input string) {

}

func (hll Hyperloglog) DistinctCount () int {

}

func Merge (hll1 Hyperloglog, hll2 Hyperloglog) *Hyperloglog {

}
