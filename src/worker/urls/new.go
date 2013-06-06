package urls

/*============================================================================*
 * {{{ Internal api 
 *============================================================================*/

type factory struct {
  UrlHelper func() *UrlHelper
}

var New factory = factory {
  UrlHelper : newUrlHelper,
}

/*============================================================================*
 * }}}
 *============================================================================*/
