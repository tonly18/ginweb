/* *
 * model
 * error code: 20000000 ` 20000999
 */

package model

import "errors"

var (
	ErrorNoRows = errors.New("sql: no rows in result set(model)")
)
