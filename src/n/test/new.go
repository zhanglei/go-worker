package test

import "testing"

/*============================================================================*
 * {{{ Internal api 
 *============================================================================*/

type factory struct {
  Assert func(t *testing.T) Assert
}

var New factory = factory {
  Assert : newAssert,
}

/*============================================================================*
 * }}}
 *============================================================================*/
