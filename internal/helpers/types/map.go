package types

func PatchMap(original map[string]interface{}, patch map[string]interface{}, reserveDefault bool) map[string]interface{} {
	patched := make(map[string]interface{})

	// Copy original map into the patched map
	for key, value := range original {
		patched[key] = value
	}

	// Apply patches from the patch map
	for key, value := range patch {
		if oldValue, exists := patched[key]; exists {
			if originalSubMap, originalIsMap := oldValue.(map[string]interface{}); originalIsMap {
				if patchSubMap, patchIsMap := value.(map[string]interface{}); patchIsMap {
					patched[key] = PatchMap(originalSubMap, patchSubMap, reserveDefault) // Recursively patch nested maps
				} else {
					if patched[key] != nil && reserveDefault {
						continue // Skip if key doesn't exist and reserveDefault is true
					}
					patched[key] = value // Update non-map value directly
				}
			} else {
				if patched[key] != nil && reserveDefault {
					continue // Skip if key doesn't exist and reserveDefault is true
				}
				patched[key] = value // Update non-map value directly
			}
		} else {
			patched[key] = value // Add new key-value pair if key doesn't exist
		}
	}

	return patched
}
