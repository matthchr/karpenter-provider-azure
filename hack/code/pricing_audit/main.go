package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	pricingAPIVersion  = "2023-01-01-preview"
	pricingURL         = "https://prices.azure.com/api/retail/prices?api-version=" + pricingAPIVersion
	pricingResponseDir = "hack/code/pricing_audit/pricing_responses"
	skuResponseDir     = "hack/code/pricing_audit/sku_responses"

	pricingMismatchReportFile = "hack/code/pricing_audit/pricing_mismatch_report.md"
	skuVsPricingReportFile    = "hack/code/pricing_audit/sku_vs_pricing_report.md"
	noPricingReportFile       = "hack/code/pricing_audit/skus_no_pricing_report.md"
)

var regions = []string{
	"australiacentral", "australiacentral2", "australiaeast", "australiasoutheast",
	"austriaeast", "brazilsouth", "brazilsoutheast", "canadacentral", "canadaeast",
	"centralindia", "centralus", "chilecentral", "eastasia", "eastus", "eastus2",
	"francecentral", "francesouth", "germanynorth", "germanywestcentral",
	"indonesiacentral", "israelcentral", "italynorth", "japaneast", "japanwest",
	"jioindiacentral", "jioindiawest", "koreacentral", "koreasouth", "malaysiawest",
	"mexicocentral", "newzealandnorth", "northcentralus", "northeurope", "norwayeast",
	"norwaywest", "polandcentral", "qatarcentral", "southafricanorth", "southafricawest",
	"southcentralus", "southeastasia", "southindia", "spaincentral", "swedencentral",
	"swedensouth", "switzerlandnorth", "switzerlandwest", "uaecentral", "uaenorth",
	"uksouth", "ukwest", "westcentralus", "westeurope", "westindia", "westus", "westus2",
	"westus3",
}

// ============================================================
// Shared types
// ============================================================

// Full pricing item (used by pricing mismatch report)
type PricingItemFull struct {
	CurrencyCode         string  `json:"currencyCode"`
	TierMinimumUnits     float64 `json:"tierMinimumUnits"`
	RetailPrice          float64 `json:"retailPrice"`
	UnitPrice            float64 `json:"unitPrice"`
	ArmRegionName        string  `json:"armRegionName"`
	Location             string  `json:"location"`
	EffectiveStartDate   string  `json:"effectiveStartDate"`
	MeterID              string  `json:"meterId"`
	MeterName            string  `json:"meterName"`
	ProductID            string  `json:"productId"`
	SkuID                string  `json:"skuId"`
	AvailabilityID       any     `json:"availabilityId"`
	ProductName          string  `json:"productName"`
	SkuName              string  `json:"skuName"`
	ServiceName          string  `json:"serviceName"`
	ServiceID            string  `json:"serviceId"`
	ServiceFamily        string  `json:"serviceFamily"`
	UnitOfMeasure        string  `json:"unitOfMeasure"`
	Type                 string  `json:"type"`
	IsPrimaryMeterRegion bool    `json:"isPrimaryMeterRegion"`
	ArmSkuName           string  `json:"armSkuName"`
	EffectiveEndDate     string  `json:"effectiveEndDate,omitempty"`
	ReservationTerm      string  `json:"reservationTerm,omitempty"`
}

type ProductsPricePage struct {
	BillingCurrency    string            `json:"BillingCurrency"`
	CustomerEntityID   string            `json:"CustomerEntityId"`
	CustomerEntityType string            `json:"CustomerEntityType"`
	Items              []PricingItemFull `json:"Items"`
	NextPageLink       string            `json:"NextPageLink"`
	Count              int               `json:"Count"`
}

type RegionPricingFull struct {
	Region    string            `json:"region"`
	Items     []PricingItemFull `json:"items"`
	FetchedAt string            `json:"fetchedAt"`
}

// Lightweight pricing item (used by SKU cross-ref and no-pricing reports when reading cached data)
type PricingItemLight struct {
	ArmSkuName  string  `json:"armSkuName"`
	SkuName     string  `json:"skuName"`
	ProductName string  `json:"productName"`
	MeterName   string  `json:"meterName"`
	RetailPrice float64 `json:"retailPrice"`
}

type RegionPricingLight struct {
	Region string             `json:"region"`
	Items  []PricingItemLight `json:"items"`
}

// SKU API types
type SKUCapability struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SKULocationInfo struct {
	Location string   `json:"location"`
	Zones    []string `json:"zones"`
}

type SKURestrictionInfo struct {
	Locations []string `json:"locations"`
	Zones     []string `json:"zones"`
}

type SKURestriction struct {
	Type       string             `json:"type"`
	Values     []string           `json:"values"`
	ReasonCode string             `json:"reasonCode"`
	Info       SKURestrictionInfo `json:"restrictionInfo"`
}

type ResourceSKU struct {
	Name         string            `json:"name"`
	ResourceType string            `json:"resourceType"`
	Tier         string            `json:"tier"`
	Size         string            `json:"size"`
	Family       string            `json:"family"`
	Locations    []string          `json:"locations"`
	LocationInfo []SKULocationInfo `json:"locationInfo"`
	Capabilities []SKUCapability   `json:"capabilities"`
	Restrictions []SKURestriction  `json:"restrictions"`
}

type SKUListResponse struct {
	Value    []ResourceSKU `json:"value"`
	NextLink string        `json:"nextLink"`
}

type CachedSKUResponse struct {
	Region    string        `json:"region"`
	SKUs      []ResourceSKU `json:"skus"`
	FetchedAt string        `json:"fetchedAt"`
}

type regionSKUResult struct {
	region string
	skus   []ResourceSKU
	err    error
}

// ============================================================
// main
// ============================================================

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: go run ./hack/code/pricing_audit/ <command>

Commands:
  pricing    Fetch retail pricing data and generate the pricing mismatch report
  skus       Fetch SKU data and generate the SKU vs pricing cross-reference report
  nopricing  Generate the "SKUs with no pricing" report (requires cached data)
  all        Run all three reports in sequence
  fetch      Fetch/cache both pricing and SKU data without generating reports
`)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	cmd := os.Args[1]
	switch cmd {
	case "pricing":
		runPricingReport()
	case "skus":
		runSKUReport()
	case "nopricing":
		runNoPricingReport()
	case "all":
		runPricingReport()
		runSKUReport()
		runNoPricingReport()
	case "fetch":
		fetchAllPricingData()
		fetchAllSKUData()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		usage()
	}
}

// ============================================================
// Shared helpers
// ============================================================

func getSubscriptionID() string {
	out, err := exec.Command("az", "account", "show", "--query", "id", "-o", "tsv").Output()
	if err != nil {
		log.Fatalf("failed to get subscription ID: %v", err)
	}
	return strings.TrimSpace(string(out))
}

func getAccessToken() string {
	out, err := exec.Command("az", "account", "get-access-token", "--query", "accessToken", "-o", "tsv").Output()
	if err != nil {
		log.Fatalf("failed to get access token: %v", err)
	}
	return strings.TrimSpace(string(out))
}

func getCapability(sku ResourceSKU, name string) string {
	for _, cap := range sku.Capabilities {
		if cap.Name == name {
			return cap.Value
		}
	}
	return ""
}

func skuInRegion(sku ResourceSKU, region string) bool {
	for _, loc := range sku.Locations {
		if strings.EqualFold(loc, region) {
			return true
		}
	}
	return false
}

func sortedKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// ============================================================
// Pricing data fetching
// ============================================================

type pricingRegionResult struct {
	region       string
	onDemandSKUs map[string]float64
	spotSKUs     map[string]float64
	err          error
}

func fetchAllPricingData() []pricingRegionResult {
	if err := os.MkdirAll(pricingResponseDir, 0755); err != nil {
		log.Fatalf("failed to create output dir: %v", err)
	}

	sem := make(chan struct{}, 8)
	var wg sync.WaitGroup
	results := make([]pricingRegionResult, len(regions))

	for i, region := range regions {
		wg.Add(1)
		go func(idx int, r string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			results[idx] = fetchPricingRegion(r)
		}(i, region)
	}
	wg.Wait()
	return results
}

func fetchPricingRegion(region string) pricingRegionResult {
	cacheFile := filepath.Join(pricingResponseDir, region+".json")

	if data, err := os.ReadFile(cacheFile); err == nil {
		log.Printf("[%s] using cached pricing response", region)
		var cached RegionPricingFull
		if err := json.Unmarshal(data, &cached); err == nil {
			onDemand, spot := categorizePricingItems(cached.Items)
			return pricingRegionResult{region: region, onDemandSKUs: onDemand, spotSKUs: spot}
		}
		log.Printf("[%s] cache parse error, re-fetching: %v", region, err)
	}

	log.Printf("[%s] fetching pricing data...", region)

	filterStr := fmt.Sprintf(
		"priceType eq 'Consumption' and currencyCode eq 'USD' and serviceFamily eq 'Compute' and serviceName eq 'Virtual Machines' and armRegionName eq '%s'",
		region,
	)

	nextURL := pricingURL + "&$filter=" + url.QueryEscape(filterStr)

	var allItems []PricingItemFull
	pageNum := 0
	for nextURL != "" {
		pageNum++
		items, next, err := fetchPricingPage(nextURL)
		if err != nil {
			success := false
			for attempt := 1; attempt <= 5; attempt++ {
				log.Printf("[%s] page %d attempt %d failed: %v, retrying in 5s...", region, pageNum, attempt, err)
				time.Sleep(5 * time.Second)
				items, next, err = fetchPricingPage(nextURL)
				if err == nil {
					success = true
					break
				}
			}
			if !success {
				return pricingRegionResult{region: region, err: fmt.Errorf("failed after retries on page %d: %v", pageNum, err)}
			}
		}
		allItems = append(allItems, items...)
		nextURL = next
	}

	log.Printf("[%s] fetched %d pricing items across %d pages", region, len(allItems), pageNum)

	cached := RegionPricingFull{
		Region:    region,
		Items:     allItems,
		FetchedAt: time.Now().UTC().Format(time.RFC3339),
	}
	data, err := json.MarshalIndent(cached, "", "  ")
	if err != nil {
		log.Printf("[%s] warning: failed to marshal cache: %v", region, err)
	} else {
		if err := os.WriteFile(cacheFile, data, 0644); err != nil {
			log.Printf("[%s] warning: failed to write cache: %v", region, err)
		}
	}

	onDemand, spot := categorizePricingItems(allItems)
	return pricingRegionResult{region: region, onDemandSKUs: onDemand, spotSKUs: spot}
}

func fetchPricingPage(pageURL string) ([]PricingItemFull, string, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, "", fmt.Errorf("status %d: %s", resp.StatusCode, string(body[:min(len(body), 200)]))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("reading body: %w", err)
	}

	var page ProductsPricePage
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, "", fmt.Errorf("parsing JSON: %w", err)
	}

	return page.Items, page.NextPageLink, nil
}

func categorizePricingItems(items []PricingItemFull) (onDemand map[string]float64, spot map[string]float64) {
	onDemand = make(map[string]float64)
	spot = make(map[string]float64)

	for _, item := range items {
		if strings.HasSuffix(item.ProductName, " Windows") {
			continue
		}
		if strings.HasSuffix(item.MeterName, " Low Priority") {
			continue
		}
		if strings.HasSuffix(item.SkuName, " Spot") {
			spot[item.ArmSkuName] = item.RetailPrice
		} else {
			onDemand[item.ArmSkuName] = item.RetailPrice
		}
	}
	return onDemand, spot
}

// ============================================================
// SKU data fetching
// ============================================================

func fetchAllSKUData() []regionSKUResult {
	if err := os.MkdirAll(skuResponseDir, 0755); err != nil {
		log.Fatalf("failed to create SKU output dir: %v", err)
	}

	subID := getSubscriptionID()
	token := getAccessToken()
	log.Printf("Using subscription: %s", subID)

	sem := make(chan struct{}, 6)
	var wg sync.WaitGroup
	results := make([]regionSKUResult, len(regions))

	for i, region := range regions {
		wg.Add(1)
		go func(idx int, r string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			skus, err := fetchSKUs(r, subID, token)
			results[idx] = regionSKUResult{region: r, skus: skus, err: err}
		}(i, region)
	}
	wg.Wait()
	return results
}

func fetchSKUs(region, subID, token string) ([]ResourceSKU, error) {
	cacheFile := filepath.Join(skuResponseDir, region+".json")

	if data, err := os.ReadFile(cacheFile); err == nil {
		log.Printf("[%s] using cached SKU response", region)
		var cached CachedSKUResponse
		if err := json.Unmarshal(data, &cached); err == nil {
			return cached.SKUs, nil
		}
		log.Printf("[%s] cache parse error, re-fetching", region)
	}

	log.Printf("[%s] fetching SKU data...", region)

	baseURL := fmt.Sprintf(
		"https://management.azure.com/subscriptions/%s/providers/Microsoft.Compute/skus?api-version=2021-07-01&$filter=location%%20eq%%20'%s'",
		subID, region,
	)

	var allSKUs []ResourceSKU
	nextURL := baseURL
	for nextURL != "" {
		skus, next, err := fetchSKUPage(nextURL, token)
		if err != nil {
			success := false
			for attempt := 1; attempt <= 3; attempt++ {
				log.Printf("[%s] attempt %d failed: %v, retrying...", region, attempt, err)
				time.Sleep(5 * time.Second)
				skus, next, err = fetchSKUPage(nextURL, token)
				if err == nil {
					success = true
					break
				}
			}
			if !success {
				return nil, fmt.Errorf("failed after retries: %v", err)
			}
		}
		allSKUs = append(allSKUs, skus...)
		nextURL = next
	}

	// Filter to virtualMachines only
	var vmSKUs []ResourceSKU
	for _, sku := range allSKUs {
		if sku.ResourceType == "virtualMachines" {
			vmSKUs = append(vmSKUs, sku)
		}
	}

	log.Printf("[%s] fetched %d VM SKUs", region, len(vmSKUs))

	cached := CachedSKUResponse{
		Region:    region,
		SKUs:      vmSKUs,
		FetchedAt: time.Now().UTC().Format(time.RFC3339),
	}
	data, _ := json.MarshalIndent(cached, "", "  ")
	os.WriteFile(cacheFile, data, 0644)

	return vmSKUs, nil
}

func fetchSKUPage(pageURL, token string) ([]ResourceSKU, string, error) {
	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, "", fmt.Errorf("status %d: %s", resp.StatusCode, string(body[:min(len(body), 300)]))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("reading body: %w", err)
	}

	var page SKUListResponse
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, "", fmt.Errorf("parsing JSON: %w", err)
	}

	return page.Value, page.NextLink, nil
}

// ============================================================
// Report 1: Pricing Mismatch (on-demand vs spot)
// ============================================================

func runPricingReport() {
	results := fetchAllPricingData()
	generatePricingMismatchReport(results)
}

func generatePricingMismatchReport(results []pricingRegionResult) {
	var sb strings.Builder
	sb.WriteString("# Pricing Mismatch Report\n\n")
	sb.WriteString(fmt.Sprintf("Generated: %s\n\n", time.Now().UTC().Format(time.RFC3339)))
	sb.WriteString("This report identifies VM sizes where:\n")
	sb.WriteString("- A region has **on-demand** pricing but **NO spot** pricing\n")
	sb.WriteString("- A region has **spot** pricing but **NO on-demand** pricing\n\n")

	var errorRegions []string
	for _, r := range results {
		if r.err != nil {
			errorRegions = append(errorRegions, fmt.Sprintf("- %s: %v", r.region, r.err))
		}
	}
	if len(errorRegions) > 0 {
		sb.WriteString("## Errors\n\n")
		for _, e := range errorRegions {
			sb.WriteString(e + "\n")
		}
		sb.WriteString("\n")
	}

	type mismatch struct {
		Region        string
		SKU           string
		HasOnDemand   bool
		HasSpot       bool
		OnDemandPrice float64
		SpotPrice     float64
	}

	var onDemandOnly []mismatch
	var spotOnly []mismatch

	for _, r := range results {
		if r.err != nil {
			continue
		}
		for sku, price := range r.onDemandSKUs {
			if _, hasSpot := r.spotSKUs[sku]; !hasSpot {
				onDemandOnly = append(onDemandOnly, mismatch{
					Region: r.region, SKU: sku, HasOnDemand: true, OnDemandPrice: price,
				})
			}
		}
		for sku, price := range r.spotSKUs {
			if _, hasOnDemand := r.onDemandSKUs[sku]; !hasOnDemand {
				spotOnly = append(spotOnly, mismatch{
					Region: r.region, SKU: sku, HasSpot: true, SpotPrice: price,
				})
			}
		}
	}

	sort.Slice(onDemandOnly, func(i, j int) bool {
		if onDemandOnly[i].SKU != onDemandOnly[j].SKU {
			return onDemandOnly[i].SKU < onDemandOnly[j].SKU
		}
		return onDemandOnly[i].Region < onDemandOnly[j].Region
	})
	sort.Slice(spotOnly, func(i, j int) bool {
		if spotOnly[i].SKU != spotOnly[j].SKU {
			return spotOnly[i].SKU < spotOnly[j].SKU
		}
		return spotOnly[i].Region < spotOnly[j].Region
	})

	totalRegions := 0
	for _, r := range results {
		if r.err == nil {
			totalRegions++
		}
	}
	sb.WriteString("## Summary\n\n")
	sb.WriteString(fmt.Sprintf("- Regions analyzed: %d\n", totalRegions))
	sb.WriteString(fmt.Sprintf("- SKUs with on-demand but NO spot: %d entries across regions\n", len(onDemandOnly)))
	sb.WriteString(fmt.Sprintf("- SKUs with spot but NO on-demand: %d entries across regions\n\n", len(spotOnly)))

	// On-demand only
	sb.WriteString("## On-Demand Only (no spot pricing)\n\n")
	if len(onDemandOnly) == 0 {
		sb.WriteString("None found.\n\n")
	} else {
		onDemandBySKU := map[string][]mismatch{}
		for _, m := range onDemandOnly {
			onDemandBySKU[m.SKU] = append(onDemandBySKU[m.SKU], m)
		}
		skus := make([]string, 0, len(onDemandBySKU))
		for sku := range onDemandBySKU {
			skus = append(skus, sku)
		}
		sort.Strings(skus)

		sb.WriteString(fmt.Sprintf("| SKU | Regions (%d unique SKUs) | On-Demand Price (first region) |\n", len(skus)))
		sb.WriteString("|-----|---------|--------|\n")
		for _, sku := range skus {
			matches := onDemandBySKU[sku]
			regionList := make([]string, 0, len(matches))
			for _, m := range matches {
				regionList = append(regionList, m.Region)
			}
			sort.Strings(regionList)
			price := matches[0].OnDemandPrice
			if len(regionList) <= 5 {
				sb.WriteString(fmt.Sprintf("| %s | %s | $%.4f |\n", sku, strings.Join(regionList, ", "), price))
			} else {
				sb.WriteString(fmt.Sprintf("| %s | %s (+%d more) | $%.4f |\n", sku, strings.Join(regionList[:5], ", "), len(regionList)-5, price))
			}
		}
		sb.WriteString("\n")
	}

	// Spot only
	sb.WriteString("## Spot Only (no on-demand pricing)\n\n")
	if len(spotOnly) == 0 {
		sb.WriteString("None found.\n\n")
	} else {
		spotBySKU := map[string][]mismatch{}
		for _, m := range spotOnly {
			spotBySKU[m.SKU] = append(spotBySKU[m.SKU], m)
		}
		skus := make([]string, 0, len(spotBySKU))
		for sku := range spotBySKU {
			skus = append(skus, sku)
		}
		sort.Strings(skus)

		sb.WriteString(fmt.Sprintf("| SKU | Regions (%d unique SKUs) | Spot Price (first region) |\n", len(skus)))
		sb.WriteString("|-----|---------|--------|\n")
		for _, sku := range skus {
			matches := spotBySKU[sku]
			regionList := make([]string, 0, len(matches))
			for _, m := range matches {
				regionList = append(regionList, m.Region)
			}
			sort.Strings(regionList)
			price := matches[0].SpotPrice
			if len(regionList) <= 5 {
				sb.WriteString(fmt.Sprintf("| %s | %s | $%.4f |\n", sku, strings.Join(regionList, ", "), price))
			} else {
				sb.WriteString(fmt.Sprintf("| %s | %s (+%d more) | $%.4f |\n", sku, strings.Join(regionList[:5], ", "), len(regionList)-5, price))
			}
		}
		sb.WriteString("\n")
	}

	// Per-region table
	sb.WriteString("## Per-Region Mismatch Counts\n\n")
	sb.WriteString("| Region | On-Demand Only | Spot Only | Total SKUs |\n")
	sb.WriteString("|--------|---------------|-----------|------------|\n")
	for _, r := range results {
		if r.err != nil {
			sb.WriteString(fmt.Sprintf("| %s | ERROR | ERROR | ERROR |\n", r.region))
			continue
		}
		odOnly := 0
		for sku := range r.onDemandSKUs {
			if _, ok := r.spotSKUs[sku]; !ok {
				odOnly++
			}
		}
		sOnly := 0
		for sku := range r.spotSKUs {
			if _, ok := r.onDemandSKUs[sku]; !ok {
				sOnly++
			}
		}
		totalSKUs := len(r.onDemandSKUs)
		if len(r.spotSKUs) > totalSKUs {
			totalSKUs = len(r.spotSKUs)
		}
		sb.WriteString(fmt.Sprintf("| %s | %d | %d | %d |\n", r.region, odOnly, sOnly, totalSKUs))
	}

	report := sb.String()
	if err := os.WriteFile(pricingMismatchReportFile, []byte(report), 0644); err != nil {
		log.Fatalf("failed to write report: %v", err)
	}
	fmt.Print(report)
	log.Printf("Report saved to %s", pricingMismatchReportFile)
}

// ============================================================
// Report 2: SKU vs Pricing Cross-Reference
// ============================================================

func runSKUReport() {
	skuResults := fetchAllSKUData()
	generateSKUCrossRefReport(skuResults)
}

func generateSKUCrossRefReport(skuResults []regionSKUResult) {
	var sb strings.Builder
	sb.WriteString("# SKU Capabilities vs Pricing: Cross-Reference Report\n\n")
	sb.WriteString(fmt.Sprintf("Generated: %s\n\n", time.Now().UTC().Format(time.RFC3339)))
	sb.WriteString("Cross-references `LowPriorityCapable` from the Resource SKUs API with spot pricing from the Retail Pricing API.\n\n")

	type regionData struct {
		region                string
		skuLowPriorityCapable map[string]bool
		skuNames              map[string]bool
		pricingHasOnDemand    map[string]bool
		pricingHasSpot        map[string]bool
	}

	var allData []regionData

	for _, sr := range skuResults {
		if sr.err != nil {
			log.Printf("[%s] SKU fetch error: %v", sr.region, sr.err)
			continue
		}

		rd := regionData{
			region:                sr.region,
			skuLowPriorityCapable: make(map[string]bool),
			skuNames:              make(map[string]bool),
			pricingHasOnDemand:    make(map[string]bool),
			pricingHasSpot:        make(map[string]bool),
		}

		for _, sku := range sr.skus {
			if !skuInRegion(sku, sr.region) {
				continue
			}
			rd.skuNames[sku.Name] = true
			lowPri := getCapability(sku, "LowPriorityCapable")
			rd.skuLowPriorityCapable[sku.Name] = strings.EqualFold(lowPri, "True")
		}

		pricingFile := filepath.Join(pricingResponseDir, sr.region+".json")
		if data, err := os.ReadFile(pricingFile); err == nil {
			var pricing RegionPricingLight
			if err := json.Unmarshal(data, &pricing); err == nil {
				for _, item := range pricing.Items {
					if strings.HasSuffix(item.ProductName, " Windows") {
						continue
					}
					if strings.HasSuffix(item.MeterName, " Low Priority") {
						continue
					}
					if strings.HasSuffix(item.SkuName, " Spot") {
						rd.pricingHasSpot[item.ArmSkuName] = true
					} else {
						rd.pricingHasOnDemand[item.ArmSkuName] = true
					}
				}
			}
		} else {
			log.Printf("[%s] no pricing data found", sr.region)
		}

		allData = append(allData, rd)
	}

	type mismatchEntry struct {
		SKU    string
		Region string
	}

	var (
		lowPriCapableNoSpot     []mismatchEntry
		notLowPriCapableHasSpot []mismatchEntry
		expectedNoSpot          int
		totalChecked            int
		totalMatch              int
	)

	type regionSummary struct {
		region                  string
		totalVMSKUs             int
		lowPriCapableCount      int
		hasSpotPricing          int
		hasOnDemandPricing      int
		lowPriCapableNoSpot     int
		notLowPriCapableHasSpot int
		onDemandNoSpotExpected  int
		onDemandNoSpotSurprise  int
	}
	var regionSummaries []regionSummary

	for _, rd := range allData {
		rs := regionSummary{region: rd.region, totalVMSKUs: len(rd.skuNames)}

		for sku := range rd.skuNames {
			isLowPriCapable := rd.skuLowPriorityCapable[sku]
			hasSpot := rd.pricingHasSpot[sku]
			hasOD := rd.pricingHasOnDemand[sku]

			if isLowPriCapable {
				rs.lowPriCapableCount++
			}
			if hasSpot {
				rs.hasSpotPricing++
			}
			if hasOD {
				rs.hasOnDemandPricing++
			}

			if hasOD || hasSpot {
				totalChecked++
				if isLowPriCapable && hasSpot {
					totalMatch++
				} else if !isLowPriCapable && !hasSpot {
					totalMatch++
					expectedNoSpot++
					rs.onDemandNoSpotExpected++
				} else if isLowPriCapable && !hasSpot {
					lowPriCapableNoSpot = append(lowPriCapableNoSpot, mismatchEntry{SKU: sku, Region: rd.region})
					rs.lowPriCapableNoSpot++
					rs.onDemandNoSpotSurprise++
				} else if !isLowPriCapable && hasSpot {
					notLowPriCapableHasSpot = append(notLowPriCapableHasSpot, mismatchEntry{SKU: sku, Region: rd.region})
					rs.notLowPriCapableHasSpot++
				}
			}
		}
		regionSummaries = append(regionSummaries, rs)
	}

	sb.WriteString("## Summary\n\n")
	sb.WriteString(fmt.Sprintf("- Total SKU+region combinations checked (with pricing data): %d\n", totalChecked))
	sb.WriteString(fmt.Sprintf("- Matches (capability aligns with pricing): %d (%.1f%%)\n", totalMatch, float64(totalMatch)/float64(totalChecked)*100))
	sb.WriteString(fmt.Sprintf("- `LowPriorityCapable=True` but NO spot price: **%d**\n", len(lowPriCapableNoSpot)))
	sb.WriteString(fmt.Sprintf("- `LowPriorityCapable=False` but HAS spot price: **%d**\n", len(notLowPriCapableHasSpot)))
	sb.WriteString(fmt.Sprintf("- `LowPriorityCapable=False` and no spot (expected): %d\n\n", expectedNoSpot))

	sb.WriteString("## LowPriorityCapable=True but NO Spot Pricing\n\n")
	sb.WriteString("These SKUs claim to support spot/low-priority but have no spot price in the retail pricing API.\n\n")
	if len(lowPriCapableNoSpot) == 0 {
		sb.WriteString("None found.\n\n")
	} else {
		grouped := map[string][]string{}
		for _, m := range lowPriCapableNoSpot {
			grouped[m.SKU] = append(grouped[m.SKU], m.Region)
		}
		skus := sortedKeys(grouped)
		sb.WriteString(fmt.Sprintf("| SKU | Regions (%d unique SKUs, %d total entries) |\n", len(skus), len(lowPriCapableNoSpot)))
		sb.WriteString("|-----|--------|\n")
		for _, sku := range skus {
			regions := grouped[sku]
			sort.Strings(regions)
			sb.WriteString(fmt.Sprintf("| %s | %s |\n", sku, strings.Join(regions, ", ")))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("## LowPriorityCapable=False but HAS Spot Pricing\n\n")
	sb.WriteString("These SKUs claim NOT to support spot but have a spot price in the retail pricing API.\n\n")
	if len(notLowPriCapableHasSpot) == 0 {
		sb.WriteString("None found.\n\n")
	} else {
		grouped := map[string][]string{}
		for _, m := range notLowPriCapableHasSpot {
			grouped[m.SKU] = append(grouped[m.SKU], m.Region)
		}
		skus := sortedKeys(grouped)
		sb.WriteString(fmt.Sprintf("| SKU | Regions (%d unique SKUs, %d total entries) |\n", len(skus), len(notLowPriCapableHasSpot)))
		sb.WriteString("|-----|--------|\n")
		for _, sku := range skus {
			regions := grouped[sku]
			sort.Strings(regions)
			sb.WriteString(fmt.Sprintf("| %s | %s |\n", sku, strings.Join(regions, ", ")))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("## Per-Region Summary\n\n")
	sb.WriteString("| Region | VM SKUs | LowPriCapable | Has Spot Price | LowPri=True,NoSpot | LowPri=False,HasSpot |\n")
	sb.WriteString("|--------|---------|---------------|----------------|--------------------|-----------------------|\n")
	for _, rs := range regionSummaries {
		sb.WriteString(fmt.Sprintf("| %s | %d | %d | %d | %d | %d |\n",
			rs.region, rs.totalVMSKUs, rs.lowPriCapableCount, rs.hasSpotPricing,
			rs.lowPriCapableNoSpot, rs.notLowPriCapableHasSpot))
	}

	report := sb.String()
	if err := os.WriteFile(skuVsPricingReportFile, []byte(report), 0644); err != nil {
		log.Fatalf("failed to write report: %v", err)
	}
	fmt.Print(report)
	log.Printf("\nReport saved to %s", skuVsPricingReportFile)
}

// ============================================================
// Report 3: SKUs with No Pricing
// ============================================================

func runNoPricingReport() {
	type regionStats struct {
		region       string
		skuCount     int
		pricingCount int
		missingCount int
	}

	allMissing := map[string][]string{} // SKU -> regions where missing
	var stats []regionStats

	for _, region := range regions {
		log.Printf("Processing %s...", region)

		skuFile := filepath.Join(skuResponseDir, region+".json")
		skuData, err := os.ReadFile(skuFile)
		if err != nil {
			log.Printf("[%s] no SKU data: %v", region, err)
			continue
		}
		var skuResp CachedSKUResponse
		if err := json.Unmarshal(skuData, &skuResp); err != nil {
			log.Printf("[%s] SKU parse error: %v", region, err)
			continue
		}
		skuNames := map[string]bool{}
		for _, sku := range skuResp.SKUs {
			if sku.ResourceType != "" && sku.ResourceType != "virtualMachines" {
				continue
			}
			if !skuInRegion(sku, region) {
				continue
			}
			skuNames[sku.Name] = true
		}

		pricingFile := filepath.Join(pricingResponseDir, region+".json")
		pricingData, err := os.ReadFile(pricingFile)
		if err != nil {
			log.Printf("[%s] no pricing data: %v", region, err)
			continue
		}
		var pricingResp RegionPricingLight
		if err := json.Unmarshal(pricingData, &pricingResp); err != nil {
			log.Printf("[%s] pricing parse error: %v", region, err)
			continue
		}
		pricingNames := map[string]bool{}
		for _, item := range pricingResp.Items {
			if strings.HasSuffix(item.ProductName, " Windows") {
				continue
			}
			if strings.HasSuffix(item.MeterName, " Low Priority") {
				continue
			}
			if !strings.HasPrefix(item.ArmSkuName, "Standard_") {
				continue
			}
			pricingNames[item.ArmSkuName] = true
		}

		missing := 0
		for name := range skuNames {
			if !pricingNames[name] {
				allMissing[name] = append(allMissing[name], region)
				missing++
			}
		}

		stats = append(stats, regionStats{
			region:       region,
			skuCount:     len(skuNames),
			pricingCount: len(pricingNames),
			missingCount: missing,
		})
	}

	var sb strings.Builder
	sb.WriteString("# SKUs with No Pricing Report\n\n")
	sb.WriteString(fmt.Sprintf("Generated: %s\n\n", time.Now().UTC().Format(time.RFC3339)))
	sb.WriteString("VM sizes listed in the Resource SKUs API but with **NO pricing** (neither on-demand nor spot) in the Retail Pricing API.\n\n")

	totalEntries := 0
	for _, rgns := range allMissing {
		totalEntries += len(rgns)
	}

	sb.WriteString("## Summary\n\n")
	sb.WriteString(fmt.Sprintf("- Regions analyzed: %d\n", len(stats)))
	sb.WriteString(fmt.Sprintf("- Unique SKUs missing pricing: %d\n", len(allMissing)))
	sb.WriteString(fmt.Sprintf("- Total SKU+region entries missing pricing: %d\n\n", totalEntries))

	sb.WriteString("## Per-Region Counts\n\n")
	sb.WriteString("| Region | SKUs in API | SKUs with Pricing | Missing Pricing | %% Missing |\n")
	sb.WriteString("|--------|-------------|-------------------|-----------------|----------|\n")
	for _, rs := range stats {
		pct := float64(rs.missingCount) / float64(rs.skuCount) * 100
		sb.WriteString(fmt.Sprintf("| %s | %d | %d | %d | %.1f%% |\n",
			rs.region, rs.skuCount, rs.pricingCount, rs.missingCount, pct))
	}
	sb.WriteString("\n")

	type skuEntry struct {
		name    string
		regions []string
	}
	var entries []skuEntry
	for name, rgns := range allMissing {
		sort.Strings(rgns)
		entries = append(entries, skuEntry{name: name, regions: rgns})
	}
	sort.Slice(entries, func(i, j int) bool {
		if len(entries[i].regions) != len(entries[j].regions) {
			return len(entries[i].regions) > len(entries[j].regions)
		}
		return entries[i].name < entries[j].name
	})

	sb.WriteString("## SKUs Missing Pricing in Many Regions (>40)\n\n")
	sb.WriteString("| SKU | # Regions | Regions |\n")
	sb.WriteString("|-----|-----------|---------|\n")
	for _, e := range entries {
		if len(e.regions) <= 40 {
			break
		}
		sb.WriteString(fmt.Sprintf("| %s | %d | %s |\n", e.name, len(e.regions), strings.Join(e.regions, ", ")))
	}
	sb.WriteString("\n")

	sb.WriteString("## SKUs Missing Pricing in Some Regions (5-40)\n\n")
	sb.WriteString("| SKU | # Regions | Regions |\n")
	sb.WriteString("|-----|-----------|---------|\n")
	for _, e := range entries {
		if len(e.regions) > 40 || len(e.regions) < 5 {
			continue
		}
		sb.WriteString(fmt.Sprintf("| %s | %d | %s |\n", e.name, len(e.regions), strings.Join(e.regions, ", ")))
	}
	sb.WriteString("\n")

	sb.WriteString("## SKUs Missing Pricing in Few Regions (1-4)\n\n")
	sb.WriteString("| SKU | Regions |\n")
	sb.WriteString("|-----|---------|\n")
	fewCount := 0
	for _, e := range entries {
		if len(e.regions) > 4 {
			continue
		}
		fewCount++
		sb.WriteString(fmt.Sprintf("| %s | %s |\n", e.name, strings.Join(e.regions, ", ")))
	}
	sb.WriteString(fmt.Sprintf("\n(%d SKUs in this category)\n\n", fewCount))

	sb.WriteString("## L-series (Storage Optimized) Missing Pricing\n\n")
	sb.WriteString("| SKU | # Regions | Regions |\n")
	sb.WriteString("|-----|-----------|---------|\n")
	for _, e := range entries {
		if !strings.HasPrefix(e.name, "Standard_L") {
			continue
		}
		sb.WriteString(fmt.Sprintf("| %s | %d | %s |\n", e.name, len(e.regions), strings.Join(e.regions, ", ")))
	}
	sb.WriteString("\n")

	report := sb.String()
	if err := os.WriteFile(noPricingReportFile, []byte(report), 0644); err != nil {
		log.Fatalf("failed to write report: %v", err)
	}
	log.Printf("Report saved to %s (%d lines)", noPricingReportFile, strings.Count(report, "\n"))
}
