/*
This Source Code Form is subject to the terms of the Mozilla Public License, v.2.0.

If a copy of the MPL was not distributed with this file, You can obtain one at
http://mozilla.org/MPL/2.0/.
*/
package version

import "fmt"

const Version = "0.2.4"

var (
	Name      string
	GitCommit string

	HumanVersion = fmt.Sprintf("%s v%s (%s)", Name, Version, GitCommit)
)
