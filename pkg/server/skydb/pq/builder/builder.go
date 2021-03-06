// Copyright 2015-present Oursky Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builder

import (
	"errors"

	"github.com/lib/pq"
)

var ErrCannotCompareUsingInOperator = errors.New(`cannot use "in" operator to compare the specified values`)

func fullQuoteIdentifier(aliasName string, columnName string) string {
	// If aliasName is empty, generate a identifier without qualifying
	// it with an alias name.
	if aliasName == "" {
		return pq.QuoteIdentifier(columnName)
	}
	return pq.QuoteIdentifier(aliasName) + "." + pq.QuoteIdentifier(columnName)
}
