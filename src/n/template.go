package n

// import "n"

/*============================================================================*
 * {{{ Public api 
 *============================================================================*/

type ITemplate interface {
}

func Template() ITemplate {
  return &template {}
}

/*============================================================================*
 * }}}
 *============================================================================*/

/*============================================================================*
 * {{{ ITemplate api implementation
 *============================================================================*/

type template struct {
}

/*============================================================================*
 * }}} 
 *============================================================================*/

/*============================================================================*
 * {{{ Internal
 *============================================================================*/

/*============================================================================*
 * }}} 
 *============================================================================*/
