package utils

import "strings"

// RemoveZoneSuffix removes "_arrival" or "_departure" suffix from zoneID
// This is needed because the search airport endpoint now returns zone IDs with these suffixes,
// but the database still expects the original zone IDs without suffixes
func RemoveZoneSuffix(zoneID string) string {
	if strings.HasSuffix(zoneID, "_arrival") {
		return strings.TrimSuffix(zoneID, "_arrival")
	}
	if strings.HasSuffix(zoneID, "_departure") {
		return strings.TrimSuffix(zoneID, "_departure")
	}
	return zoneID
}
