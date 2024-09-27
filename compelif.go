package main

import (
    
    "strings"
)

// Function to get the population of a country (simplified data for illustration)
func compelif(country string) int {
    country = strings.ToLower(country)
    switch country {
    case "china":
        return 1402000000
    case "india":
        return 1380000000
    case "japan":
        return 125800000
    case "south korea":
        return 51780000
    case "indonesia":
        return 276400000
    case "thailand":
        return 70000000
    case "philippines":
        return 113000000
    case "vietnam":
        return 97000000
    default:
        return 0
    }
}