package n

/*============================================================================*
 * {{{ Internal api 
 *============================================================================*/

type factory struct {
  Resolver func() Resolver
  Str func() Str
}

var New factory = factory {
  Resolver : newResolver,
  Str : newStr,
}

/*============================================================================*
 * }}}
 *============================================================================*/
