// Copyright 2020 TiKV Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package statistics

// StoreStatKind represents the statistics type of store.
type StoreStatKind int

// Different store statistics kinds.
const (
	StoreReadBytes StoreStatKind = iota
	StoreReadKeys
	StoreWriteBytes
	StoreWriteKeys
	StoreCPUUsage
	StoreDiskReadRate
	StoreDiskWriteRate

	StoreStatCount
)

func (k StoreStatKind) String() string {
	switch k {
	case StoreReadBytes:
		return "store_read_bytes"
	case StoreReadKeys:
		return "store_read_keys"
	case StoreWriteBytes:
		return "store_write_bytes"
	case StoreWriteKeys:
		return "store_write_keys"
	case StoreCPUUsage:
		return "store_cpu_usage"
	case StoreDiskReadRate:
		return "store_disk_read_rate"
	case StoreDiskWriteRate:
		return "store_disk_write_rate"
	}

	return "unknown StoreStatKind"
}
