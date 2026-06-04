# SKU Capabilities vs Pricing: Cross-Reference Report

Generated: 2026-09-01T21:40:24Z

Cross-references `LowPriorityCapable` from the Resource SKUs API with spot pricing from the Retail Pricing API.

## Summary

- Total SKU+region combinations checked (with pricing data): 51640
- Matches (capability aligns with pricing): 50269 (97.3%)
- `LowPriorityCapable=True` but NO spot price: **735**
- `LowPriorityCapable=False` but HAS spot price: **636**
- `LowPriorityCapable=False` and no spot (expected): 1722

## LowPriorityCapable=True but NO Spot Pricing

These SKUs claim to support spot/low-priority but have no spot price in the retail pricing API.

| SKU | Regions (546 unique SKUs, 735 total entries) |
|-----|--------|
| Standard_A1_v2 | eastus2, westus2 |
| Standard_A4_v2 | brazilsouth |
| Standard_A4m_v2 | eastus |
| Standard_B12ms | austriaeast, chilecentral, indonesiacentral, italynorth, malaysiawest, mexicocentral, newzealandnorth, polandcentral, spaincentral |
| Standard_B16als_v2 | westus2 |
| Standard_B16as_v2 | brazilsouth, westus2 |
| Standard_B16ms | austriaeast, chilecentral, indonesiacentral, italynorth, malaysiawest, mexicocentral, newzealandnorth, polandcentral, spaincentral |
| Standard_B1ls | austriaeast, chilecentral, indonesiacentral, italynorth, malaysiawest, mexicocentral, newzealandnorth, polandcentral, spaincentral |
| Standard_B1ms | austriaeast, chilecentral, indonesiacentral, italynorth, malaysiawest, mexicocentral, newzealandnorth, polandcentral, spaincentral |
| Standard_B1s | austriaeast, chilecentral, indonesiacentral, italynorth, malaysiawest, mexicocentral, newzealandnorth, polandcentral, spaincentral |
| Standard_B20ms | austriaeast, chilecentral, indonesiacentral, italynorth, malaysiawest, mexicocentral, newzealandnorth, polandcentral, spaincentral |
| Standard_B2as_v2 | eastus2 |
| Standard_B2ms | austriaeast, chilecentral, indonesiacentral, italynorth, malaysiawest, mexicocentral, newzealandnorth, polandcentral, spaincentral |
| Standard_B2pts_v2 | eastus2, westcentralus |
| Standard_B2s | austriaeast, chilecentral, indonesiacentral, italynorth, malaysiawest, mexicocentral, newzealandnorth, polandcentral, spaincentral |
| Standard_B2ts_v2 | eastus2 |
| Standard_B32s_v2 | westus2 |
| Standard_B4ms | austriaeast, chilecentral, indonesiacentral, italynorth, malaysiawest, mexicocentral, newzealandnorth, polandcentral, spaincentral |
| Standard_B4pls_v2 | uksouth |
| Standard_B4ps_v2 | westus2 |
| Standard_B8as_v2 | mexicocentral |
| Standard_B8ls_v2 | brazilsouth, westus2 |
| Standard_B8ms | austriaeast, chilecentral, indonesiacentral, italynorth, malaysiawest, mexicocentral, newzealandnorth, polandcentral, spaincentral |
| Standard_B8pls_v2 | eastus |
| Standard_B8ps_v2 | westcentralus |
| Standard_D1 | eastus |
| Standard_D128alds_v7 | eastus |
| Standard_D128ls_v6 | westus2 |
| Standard_D128nds_v6 | westcentralus |
| Standard_D128nlds_v6 | brazilsouth, westus2 |
| Standard_D128s_v7 | eastus2, westus2 |
| Standard_D12_v2 | eastus |
| Standard_D13_v2 | westus2 |
| Standard_D160alds_v7 | eastus |
| Standard_D16a_v4 | eastus2 |
| Standard_D16ads_v5 | uksouth |
| Standard_D16ads_v6 | eastus2 |
| Standard_D16ads_v7 | eastus2 |
| Standard_D16alds_v6 | westus2 |
| Standard_D16als_v7 | westus2 |
| Standard_D16as_v5 | mexicocentral |
| Standard_D16as_v6 | westcentralus |
| Standard_D16ds_v6 | westcentralus |
| Standard_D16lds_v6 | eastus2 |
| Standard_D16ls_v5 | eastus2 |
| Standard_D16ls_v6 | uksouth |
| Standard_D16ls_v7 | eastus2 |
| Standard_D16nlds_v6 | westcentralus |
| Standard_D16pls_v6 | westus2 |
| Standard_D192s_v6 | westus2 |
| Standard_D1_v2 | westcentralus |
| Standard_D2 | eastus2, northcentralus |
| Standard_D2_v2 | brazilsouth |
| Standard_D2ads_v5 | westus2 |
| Standard_D2ads_v6 | westcentralus |
| Standard_D2ads_v7 | westus2 |
| Standard_D2alds_v6 | westus2 |
| Standard_D2als_v7 | westus2 |
| Standard_D2as_v6 | brazilsouth |
| Standard_D2as_v7 | westus2 |
| Standard_D2d_v4 | westus2 |
| Standard_D2ds_v6 | eastus2 |
| Standard_D2lds_v5 | eastus2 |
| Standard_D2lds_v6 | westus2 |
| Standard_D2ls_v7 | eastus |
| Standard_D2nlds_v6 | westus2 |
| Standard_D2nls_v6 | uksouth |
| Standard_D2ns_v6 | westcentralus, westus2 |
| Standard_D2plds_v6 | westus2 |
| Standard_D2ps_v6 | eastus2 |
| Standard_D2s_v3 | eastus2 |
| Standard_D32_v4 | brazilsouth |
| Standard_D32_v5 | westus2 |
| Standard_D32a_v4 | brazilsouth, mexicocentral |
| Standard_D32ads_v5 | brazilsouth, westus3 |
| Standard_D32alds_v7 | westus2 |
| Standard_D32als_v7 | eastus2 |
| Standard_D32as_v6 | brazilsouth |
| Standard_D32d_v5 | eastus |
| Standard_D32ds_v5 | eastus |
| Standard_D32lds_v5 | westcentralus |
| Standard_D32ls_v6 | westus2 |
| Standard_D32nds_v6 | eastus2 |
| Standard_D32nlds_v6 | brazilsouth, westcentralus |
| Standard_D32plds_v5 | eastus2 |
| Standard_D32ps_v5 | westus2 |
| Standard_D32ps_v6 | mexicocentral |
| Standard_D32s_v3 | westus2 |
| Standard_D32s_v6 | westus2 |
| Standard_D32s_v7 | westus2 |
| Standard_D3_v2 | eastus2 |
| Standard_D48_v4 | eastus2 |
| Standard_D48ads_v7 | westus2 |
| Standard_D48alds_v6 | westus2 |
| Standard_D48als_v6 | westcentralus, westus2 |
| Standard_D48as_v4 | eastus |
| Standard_D48as_v6 | westus2 |
| Standard_D48ds_v4 | westcentralus |
| Standard_D48lds_v5 | uksouth |
| Standard_D48lds_v6 | westus2 |
| Standard_D48lds_v7 | westus2 |
| Standard_D48ls_v6 | brazilsouth |
| Standard_D48nds_v6 | brazilsouth |
| Standard_D48nlds_v6 | eastus2 |
| Standard_D48nls_v6 | brazilsouth, eastus2 |
| Standard_D48ns_v6 | brazilsouth |
| Standard_D48plds_v5 | westus2 |
| Standard_D48plds_v6 | uksouth, westcentralus |
| Standard_D48pls_v5 | westus2 |
| Standard_D48s_v3 | brazilsouth, eastus2 |
| Standard_D48s_v5 | eastus, eastus2 |
| Standard_D48s_v6 | eastus, westcentralus |
| Standard_D4_v2 | westcentralus |
| Standard_D4_v4 | westcentralus, westus2 |
| Standard_D4_v5 | brazilsoutheast, eastus2 |
| Standard_D4a_v4 | mexicocentral, westus2 |
| Standard_D4ads_v5 | eastus2 |
| Standard_D4ads_v6 | brazilsouth |
| Standard_D4ads_v7 | westus2 |
| Standard_D4alds_v7 | eastus2 |
| Standard_D4as_v5 | eastus2 |
| Standard_D4as_v6 | mexicocentral |
| Standard_D4ds_v5 | eastus |
| Standard_D4ds_v6 | brazilsouth, westus2 |
| Standard_D4ls_v6 | eastus2, mexicocentral |
| Standard_D4ls_v7 | westus2 |
| Standard_D4nds_v6 | northcentralus |
| Standard_D4nlds_v6 | westus2 |
| Standard_D4pds_v5 | westus2 |
| Standard_D4plds_v6 | eastus2 |
| Standard_D4pls_v6 | brazilsouth |
| Standard_D4ps_v5 | westus2 |
| Standard_D4ps_v6 | westus2 |
| Standard_D4s_v5 | canadacentral |
| Standard_D64_v3 | westus2 |
| Standard_D64_v5 | brazilsoutheast |
| Standard_D64ads_v5 | westus2 |
| Standard_D64als_v6 | uksouth |
| Standard_D64as_v5 | brazilsoutheast, westus2 |
| Standard_D64as_v6 | brazilsoutheast |
| Standard_D64as_v7 | westus2 |
| Standard_D64ds_v5 | westus2 |
| Standard_D64ds_v6 | uksouth, westus2 |
| Standard_D64lds_v5 | brazilsouth |
| Standard_D64lds_v6 | brazilsouth |
| Standard_D64ls_v5 | westus2 |
| Standard_D64nds_v6 | eastus2 |
| Standard_D64pds_v6 | westus2 |
| Standard_D64plds_v5 | eastus2, westus2 |
| Standard_D64pls_v6 | eastus2 |
| Standard_D64ps_v5 | westus2 |
| Standard_D64s_v3 | eastus2 |
| Standard_D64s_v6 | westus2 |
| Standard_D64s_v7 | westus2 |
| Standard_D8_v5 | eastus, eastus2, westus2 |
| Standard_D8a_v4 | eastus2 |
| Standard_D8ads_v5 | westus2 |
| Standard_D8ads_v6 | mexicocentral |
| Standard_D8als_v7 | eastus2 |
| Standard_D8as_v4 | brazilsouth, eastus |
| Standard_D8as_v6 | eastus2, westus2 |
| Standard_D8ds_v4 | westcentralus |
| Standard_D8ds_v6 | eastus2, westus2 |
| Standard_D8lds_v5 | uksouth |
| Standard_D8ls_v5 | eastus2 |
| Standard_D8nlds_v6 | westcentralus, westus2 |
| Standard_D8ns_v6 | brazilsouth, eastus2 |
| Standard_D8pds_v6 | westcentralus |
| Standard_D8plds_v5 | canadacentral, westus2 |
| Standard_D8plds_v6 | eastus2 |
| Standard_D8pls_v6 | mexicocentral |
| Standard_D8s_v6 | eastus2 |
| Standard_D8s_v7 | eastus2 |
| Standard_D96a_v4 | westus2 |
| Standard_D96ads_v5 | westus2 |
| Standard_D96als_v6 | westus2 |
| Standard_D96als_v7 | westus2 |
| Standard_D96as_v5 | brazilsouth, brazilsoutheast |
| Standard_D96as_v6 | brazilsouth |
| Standard_D96as_v7 | eastus2 |
| Standard_D96ds_v6 | westus2 |
| Standard_D96lds_v6 | uksouth, westus2 |
| Standard_D96ls_v5 | westus2 |
| Standard_D96nlds_v6 | westus2 |
| Standard_D96nls_v6 | westus2 |
| Standard_D96ps_v6 | eastus2 |
| Standard_DC16as_v6 | brazilsoutheast |
| Standard_DC24ds_v3 | eastus2 |
| Standard_DC24s_v3 | westus2 |
| Standard_DC2ads_v6 | eastus2, northcentralus |
| Standard_DC2as_v5 | eastus |
| Standard_DC2as_v6 | eastus2, westus2 |
| Standard_DC2es_v6 | westus3 |
| Standard_DC2s_v3 | westus2 |
| Standard_DC32ads_v5 | eastus |
| Standard_DC48as_v5 | eastus |
| Standard_DC48s_v3 | westus2 |
| Standard_DC4ads_v6 | brazilsouth, eastus2 |
| Standard_DC4as_v6 | eastus |
| Standard_DC4ds_v3 | eastus2 |
| Standard_DC96ads_v6 | brazilsouth |
| Standard_DS1 | brazilsouth |
| Standard_DS13-4_v2 | brazilsouth |
| Standard_DS14-4_v2 | eastus2 |
| Standard_DS14-8_v2 | eastus2, westcentralus |
| Standard_DS14_v2 | eastus2 |
| Standard_DS15_v2 | westus2 |
| Standard_DS1_v2 | eastus2 |
| Standard_DS2_v2 | brazilsouth, northcentralus, westcentralus |
| Standard_E104i_v5 | westus2 |
| Standard_E104id_v5 | eastus |
| Standard_E112iads_v5 | eastus2, westcentralus |
| Standard_E112ias_v5 | eastus2, westus2 |
| Standard_E112ibs_v5 | eastus2 |
| Standard_E128-64s_v6 | westus2 |
| Standard_E128ds_v6 | northcentralus, westus2 |
| Standard_E128ns_v6 | westus2 |
| Standard_E128s_v6 | westus2 |
| Standard_E16-4as_v4 | westcentralus |
| Standard_E16-4as_v7 | uksouth |
| Standard_E16-4ds_v5 | canadacentral, southcentralus, westus2 |
| Standard_E16-4ds_v6 | westus2 |
| Standard_E16-4s_v3 | brazilsouth |
| Standard_E16-4s_v4 | eastus2, westus2 |
| Standard_E16-4s_v6 | brazilsoutheast, eastus2 |
| Standard_E16-8as_v5 | eastus |
| Standard_E16-8as_v7 | uksouth |
| Standard_E16-8s_v3 | mexicocentral |
| Standard_E16-8s_v4 | brazilsoutheast |
| Standard_E16-8s_v5 | westus2 |
| Standard_E16-8s_v6 | brazilsoutheast, northcentralus |
| Standard_E16_v3 | westus2 |
| Standard_E16_v4 | eastus2 |
| Standard_E16_v5 | eastus, uksouth |
| Standard_E16ads_v6 | mexicocentral |
| Standard_E16ads_v7 | brazilsouth, eastus2 |
| Standard_E16as_v4 | brazilsoutheast, westcentralus |
| Standard_E16as_v6 | eastus, westus2 |
| Standard_E16bds_v5 | brazilsoutheast |
| Standard_E16bs_v5 | canadacentral |
| Standard_E16d_v4 | brazilsouth |
| Standard_E16ds_v4 | brazilsoutheast, eastus2 |
| Standard_E16nds_v6 | brazilsouth |
| Standard_E16ps_v6 | canadacentral, eastus2 |
| Standard_E16s_v6 | brazilsoutheast |
| Standard_E20ads_v5 | westcentralus |
| Standard_E20ads_v6 | brazilsouth |
| Standard_E20as_v5 | eastus |
| Standard_E20as_v6 | eastus |
| Standard_E20ds_v4 | northcentralus, uksouth, westcentralus |
| Standard_E20ds_v5 | brazilsoutheast |
| Standard_E20ps_v5 | westus2 |
| Standard_E20s_v4 | westus2 |
| Standard_E20s_v5 | eastus2 |
| Standard_E20s_v6 | mexicocentral |
| Standard_E2_v3 | westus2 |
| Standard_E2_v4 | eastus2 |
| Standard_E2a_v4 | brazilsouth |
| Standard_E2ads_v6 | eastus |
| Standard_E2as_v4 | westus2 |
| Standard_E2as_v5 | brazilsoutheast |
| Standard_E2d_v4 | brazilsouth |
| Standard_E2d_v5 | eastus2 |
| Standard_E2ds_v4 | eastus2 |
| Standard_E2pds_v6 | eastus2 |
| Standard_E2ps_v6 | brazilsouth |
| Standard_E2s_v3 | eastus |
| Standard_E32-16ads_v5 | westcentralus |
| Standard_E32-16as_v4 | eastus2 |
| Standard_E32-16as_v5 | brazilsoutheast |
| Standard_E32-16as_v7 | westus2 |
| Standard_E32-16ds_v4 | eastus, westus2 |
| Standard_E32-16ds_v6 | canadacentral |
| Standard_E32-16s_v3 | brazilsouth |
| Standard_E32-8ads_v5 | brazilsouth, mexicocentral |
| Standard_E32-8as_v5 | westus2 |
| Standard_E32-8ds_v5 | westus2 |
| Standard_E32-8ds_v6 | eastus2, mexicocentral |
| Standard_E32-8s_v3 | westus2 |
| Standard_E32-8s_v4 | westcentralus, westus2 |
| Standard_E32-8s_v5 | mexicocentral |
| Standard_E32_v3 | westus2 |
| Standard_E32_v4 | mexicocentral, westcentralus, westus2 |
| Standard_E32ads_v5 | brazilsoutheast, westus2 |
| Standard_E32ads_v7 | brazilsouth |
| Standard_E32as_v4 | eastus2 |
| Standard_E32as_v6 | southcentralus, westus2 |
| Standard_E32as_v7 | westus2 |
| Standard_E32bs_v5 | mexicocentral, westus2 |
| Standard_E32d_v4 | brazilsouth |
| Standard_E32ds_v4 | westcentralus |
| Standard_E32ds_v5 | westus2 |
| Standard_E32ds_v7 | eastus2 |
| Standard_E32ns_v6 | eastus2 |
| Standard_E32s_v3 | westcentralus, westus2 |
| Standard_E32s_v4 | eastus2, westcentralus |
| Standard_E32s_v5 | brazilsoutheast |
| Standard_E4-2ads_v5 | westus2 |
| Standard_E4-2ds_v4 | eastus2 |
| Standard_E4-2ds_v6 | uksouth, westus2 |
| Standard_E4-2s_v4 | westus2 |
| Standard_E48_v3 | eastus |
| Standard_E48a_v4 | westcentralus |
| Standard_E48ds_v6 | westus2 |
| Standard_E48ns_v6 | westcentralus |
| Standard_E48pds_v6 | brazilsouth, eastus2, westus2 |
| Standard_E48ps_v6 | mexicocentral |
| Standard_E48s_v3 | westus2 |
| Standard_E48s_v5 | eastus, westus2 |
| Standard_E4_v3 | brazilsouth |
| Standard_E4_v4 | westus2 |
| Standard_E4_v5 | eastus2 |
| Standard_E4ads_v7 | westus2 |
| Standard_E4as_v4 | brazilsoutheast |
| Standard_E4as_v5 | westus2 |
| Standard_E4as_v6 | westcentralus |
| Standard_E4d_v4 | brazilsouth |
| Standard_E4ds_v4 | brazilsoutheast |
| Standard_E4ds_v6 | uksouth |
| Standard_E4nds_v6 | westus2 |
| Standard_E4ns_v6 | westus2 |
| Standard_E4ps_v5 | uksouth |
| Standard_E4ps_v6 | westus2 |
| Standard_E4s_v4 | eastus2 |
| Standard_E4s_v6 | northcentralus |
| Standard_E64-16ds_v4 | uksouth |
| Standard_E64-16ds_v6 | westus2 |
| Standard_E64-16s_v3 | canadacentral |
| Standard_E64-16s_v6 | westus2 |
| Standard_E64-32ads_v5 | eastus2 |
| Standard_E64-32as_v7 | brazilsouth |
| Standard_E64-32ds_v4 | westus2 |
| Standard_E64-32ds_v6 | westcentralus |
| Standard_E64-32s_v4 | brazilsoutheast, westus2 |
| Standard_E64-32s_v6 | westus2 |
| Standard_E64_v5 | eastus |
| Standard_E64ads_v6 | westcentralus |
| Standard_E64as_v6 | westus2 |
| Standard_E64bs_v5 | westus2 |
| Standard_E64d_v5 | eastus |
| Standard_E64ds_v4 | westus2 |
| Standard_E64ds_v6 | brazilsouth |
| Standard_E64pds_v6 | eastus2 |
| Standard_E64s_v3 | westus2 |
| Standard_E64s_v4 | brazilsoutheast |
| Standard_E64s_v5 | eastus2 |
| Standard_E64s_v6 | northcentralus |
| Standard_E8-2ads_v5 | brazilsoutheast |
| Standard_E8-2ads_v7 | eastus |
| Standard_E8-2as_v5 | eastus2, westus2 |
| Standard_E8-2ds_v5 | westus2 |
| Standard_E8-2s_v3 | mexicocentral |
| Standard_E8-2s_v4 | eastus2 |
| Standard_E8-2s_v5 | westus2 |
| Standard_E8-4as_v5 | southcentralus |
| Standard_E8-4ds_v4 | eastus |
| Standard_E8-4ds_v5 | brazilsouth |
| Standard_E8-4ds_v6 | westcentralus |
| Standard_E8-4s_v3 | westcentralus |
| Standard_E8-4s_v4 | westus2 |
| Standard_E80is_v4 | westus2 |
| Standard_E8_v3 | eastus |
| Standard_E8_v4 | brazilsouth, uksouth |
| Standard_E8a_v4 | uksouth, westus2 |
| Standard_E8ads_v5 | westus2 |
| Standard_E8as_v4 | brazilsouth, uksouth |
| Standard_E8as_v6 | westus2 |
| Standard_E8as_v7 | brazilsouth |
| Standard_E8d_v4 | westus2 |
| Standard_E8ds_v4 | westcentralus |
| Standard_E8ds_v5 | brazilsouth |
| Standard_E8ds_v6 | brazilsoutheast, eastus2 |
| Standard_E8pds_v6 | westus2 |
| Standard_E8ps_v6 | eastus2 |
| Standard_E8s_v3 | brazilsoutheast, uksouth |
| Standard_E8s_v4 | eastus, mexicocentral |
| Standard_E8s_v5 | eastus2 |
| Standard_E96-24ads_v5 | westus2 |
| Standard_E96-24ads_v6 | westus2 |
| Standard_E96-24as_v5 | brazilsouth |
| Standard_E96-24ds_v5 | eastus, westus2 |
| Standard_E96-24s_v5 | canadacentral |
| Standard_E96-48as_v4 | eastus2 |
| Standard_E96-48as_v5 | mexicocentral |
| Standard_E96-48as_v7 | brazilsouth |
| Standard_E96-48ds_v6 | westus2 |
| Standard_E96-48s_v5 | canadacentral, westus2 |
| Standard_E96-48s_v6 | brazilsouth |
| Standard_E96a_v4 | westcentralus |
| Standard_E96ads_v5 | northcentralus |
| Standard_E96ads_v7 | brazilsouth |
| Standard_E96bds_v5 | eastus2 |
| Standard_E96bs_v5 | brazilsouth |
| Standard_E96d_v5 | brazilsoutheast, westus2 |
| Standard_E96ds_v5 | eastus, eastus2 |
| Standard_E96ds_v7 | eastus2 |
| Standard_E96ns_v6 | canadacentral |
| Standard_E96s_v5 | westus2 |
| Standard_E96s_v6 | eastus2 |
| Standard_EC16ads_v6 | eastus2 |
| Standard_EC16as_v6 | uksouth |
| Standard_EC20ads_v5 | eastus |
| Standard_EC2as_v6 | eastus |
| Standard_EC32ads_v6 | eastus2 |
| Standard_EC48ads_v6 | eastus2 |
| Standard_EC4as_v6 | brazilsouth |
| Standard_EC64as_v6 | westus2 |
| Standard_EC8as_cc_v5 | eastus |
| Standard_EC96as_v6 | westus2 |
| Standard_F16-4amds_v7 | uksouth |
| Standard_F16-8ams_v7 | eastus, westus2 |
| Standard_F16amds_v7 | westus2 |
| Standard_F16as_v6 | eastus2, westcentralus |
| Standard_F16as_v7 | uksouth |
| Standard_F1alds_v7 | eastus2 |
| Standard_F1als_v7 | eastus2 |
| Standard_F1amds_v7 | eastus |
| Standard_F1s | brazilsoutheast |
| Standard_F2 | brazilsouth, northcentralus |
| Standard_F2als_v6 | mexicocentral |
| Standard_F2amds_v7 | eastus2 |
| Standard_F2as_v6 | mexicocentral |
| Standard_F2s | eastus2, westcentralus, westus2 |
| Standard_F32-16amds_v7 | uksouth |
| Standard_F32-16ams_v7 | eastus |
| Standard_F32-8amds_v7 | eastus |
| Standard_F32amds_v7 | eastus2 |
| Standard_F32ams_v6 | brazilsouth, northcentralus, westcentralus |
| Standard_F32as_v6 | westcentralus |
| Standard_F32as_v7 | eastus2, westus2 |
| Standard_F4 | westus2 |
| Standard_F4-1amds_v7 | uksouth |
| Standard_F4-2ams_v7 | westus2 |
| Standard_F48alds_v7 | eastus2 |
| Standard_F48als_v6 | eastus2 |
| Standard_F48als_v7 | westus2 |
| Standard_F4ads_v7 | westus2 |
| Standard_F4als_v6 | eastus2 |
| Standard_F4as_v7 | eastus, westus2 |
| Standard_F64-16ams_v7 | eastus2 |
| Standard_F64-32amds_v7 | westus2 |
| Standard_F64als_v6 | westcentralus |
| Standard_F64ams_v6 | westus2 |
| Standard_F64as_v6 | westcentralus |
| Standard_F72s_v2 | westus2 |
| Standard_F8 | eastus |
| Standard_F8-2amds_v7 | westus2 |
| Standard_F80alds_v7 | eastus2 |
| Standard_F8als_v6 | westus2 |
| Standard_F8amds_v7 | eastus2 |
| Standard_F8ams_v6 | brazilsoutheast |
| Standard_F8ams_v7 | westus2 |
| Standard_F8as_v6 | westcentralus |
| Standard_F8as_v7 | westus2 |
| Standard_FX12mds_v2 | westus3 |
| Standard_FX16-4mds_v2 | eastus2 |
| Standard_FX16-4ms_v2 | eastus |
| Standard_FX16-8mds_v2 | eastus2, uksouth, westcentralus |
| Standard_FX16-8ms_v2 | eastus |
| Standard_FX24-12mds_v2 | brazilsouth |
| Standard_FX24-6ms_v2 | eastus2 |
| Standard_FX24mds_v2 | brazilsouth |
| Standard_FX2mds_v2 | westus2 |
| Standard_FX32-16ms_v2 | eastus2 |
| Standard_FX32-8mds_v2 | brazilsouth |
| Standard_FX32-8ms_v2 | eastus, westcentralus |
| Standard_FX32mds_v2 | westus2 |
| Standard_FX32ms_v2 | brazilsouth |
| Standard_FX48-12mds_v2 | eastus2 |
| Standard_FX4ms_v2 | brazilsouth, westcentralus |
| Standard_FX64ms_v2 | westus2 |
| Standard_FX8-4mds_v2 | eastus |
| Standard_FX96-24mds_v2 | brazilsouth |
| Standard_FX96-48mds_v2 | westcentralus |
| Standard_FX96mds_v2 | westus2 |
| Standard_FX96ms_v2 | eastus2 |
| Standard_G2 | westus2 |
| Standard_GS1 | westus2 |
| Standard_GS2 | eastus2 |
| Standard_GS3 | eastus2, westus2 |
| Standard_HB120-32rs_v2 | canadacentral |
| Standard_HB120-64rs_v2 | brazilsouth, westus2 |
| Standard_HB176-24rs_v4 | eastus |
| Standard_HC44-16rs | eastus, eastus2 |
| Standard_HX176-24rs | eastus2 |
| Standard_L16s_v3 | eastus2 |
| Standard_L16s_v4 | westus2 |
| Standard_L2as_v4 | eastus, westus2 |
| Standard_L2s_v4 | brazilsouth |
| Standard_L32aos_v4 | brazilsouth |
| Standard_L32s_v4 | eastus2, mexicocentral |
| Standard_L48as_v3 | brazilsoutheast |
| Standard_L48as_v4 | uksouth |
| Standard_L48s_v2 | westus2 |
| Standard_L48s_v4 | mexicocentral |
| Standard_L64s_v3 | eastus2 |
| Standard_L80as_v3 | eastus2 |
| Standard_L8s_v2 | westus2 |
| Standard_L8s_v3 | eastus2, westus2 |
| Standard_M128 | westus2 |
| Standard_M128-32ms | eastus |
| Standard_M128-64ms | brazilsoutheast, westus2 |
| Standard_M128bds_v3 | brazilsouth |
| Standard_M128dms_v2 | eastus2 |
| Standard_M128ds_v2 | brazilsouth |
| Standard_M128ms | eastus |
| Standard_M128s | westus2 |
| Standard_M128s_v2 | uksouth |
| Standard_M16-4ms | mexicocentral |
| Standard_M16ms | eastus2 |
| Standard_M176-88bds_4_v3 | westus2 |
| Standard_M176-88bds_v3 | eastus2 |
| Standard_M176bds_4_v3 | westus2 |
| Standard_M176bs_v3 | uksouth |
| Standard_M176s_4_v3 | eastus2 |
| Standard_M192ims_v2 | eastus2 |
| Standard_M192is_v2 | westus2 |
| Standard_M208ms_v2 | eastus |
| Standard_M208s_v2 | westus2 |
| Standard_M24s_v3 | westus2 |
| Standard_M32bds_v3 | brazilsouth |
| Standard_M32ls | westus2 |
| Standard_M416-208s_v2 | brazilsouth, westus2 |
| Standard_M416ds_6_v3 | eastus2 |
| Standard_M416s_10_v2 | brazilsouth, eastus2 |
| Standard_M416s_10_v3 | brazilsouth |
| Standard_M416s_12_v3 | brazilsouth |
| Standard_M416s_8_v2 | brazilsouth |
| Standard_M416s_9_v3 | eastus2 |
| Standard_M624s_12_v3 | eastus2 |
| Standard_M64-16ms | uksouth |
| Standard_M64bds_1_v3 | uksouth |
| Standard_M64m | eastus2 |
| Standard_M832ds_12_v3 | eastus2 |
| Standard_M832is_16_v3 | westus2 |
| Standard_M8ms | brazilsouth |
| Standard_M96bds_v3 | eastus2 |
| Standard_M96s_1_v3 | uksouth |
| Standard_NC24ads_A100_v4 | westus2 |
| Standard_NC48ads_A100_v4 | uksouth |
| Standard_NC72lds_xl_RTXPRO6000BSE_v6 | westus2 |
| Standard_NC8as_T4_v3 | eastus2, westus2 |
| Standard_NV12s_v3 | brazilsouth, westus2 |
| Standard_NV16as_v4 | eastus2 |
| Standard_NV32as_v4 | eastus2 |
| Standard_NV6ads_A10_v5 | southcentralus, westus2 |

## LowPriorityCapable=False but HAS Spot Pricing

These SKUs claim NOT to support spot but have a spot price in the retail pricing API.

| SKU | Regions (358 unique SKUs, 636 total entries) |
|-----|--------|
| Standard_D16_v5 | francecentral, westcentralus |
| Standard_D16a_v4 | southindia, westcentralus |
| Standard_D16ads_v5 | westcentralus |
| Standard_D16as_v4 | southindia |
| Standard_D16as_v5 | westcentralus |
| Standard_D16d_v4 | qatarcentral |
| Standard_D16d_v5 | qatarcentral, westcentralus |
| Standard_D16ds_v5 | francecentral, qatarcentral |
| Standard_D16lds_v5 | qatarcentral |
| Standard_D16ls_v5 | qatarcentral |
| Standard_D16pds_v5 | northcentralus |
| Standard_D16plds_v5 | northcentralus |
| Standard_D16pls_v5 | northcentralus |
| Standard_D16ps_v5 | northcentralus |
| Standard_D16s_v5 | francecentral, qatarcentral |
| Standard_D2_v5 | francecentral, westcentralus |
| Standard_D2a_v4 | southindia, westcentralus |
| Standard_D2ads_v5 | westcentralus |
| Standard_D2as_v4 | southindia |
| Standard_D2as_v5 | westcentralus |
| Standard_D2d_v4 | qatarcentral |
| Standard_D2d_v5 | qatarcentral, westcentralus |
| Standard_D2ds_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D2lds_v5 | qatarcentral |
| Standard_D2ls_v5 | qatarcentral |
| Standard_D2pds_v5 | northcentralus |
| Standard_D2plds_v5 | northcentralus |
| Standard_D2pls_v5 | northcentralus |
| Standard_D2ps_v5 | northcentralus |
| Standard_D2s_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D32_v5 | francecentral, westcentralus |
| Standard_D32a_v4 | southindia, westcentralus |
| Standard_D32ads_v5 | westcentralus |
| Standard_D32as_v4 | southindia, westcentralus |
| Standard_D32as_v5 | westcentralus |
| Standard_D32d_v4 | qatarcentral |
| Standard_D32d_v5 | qatarcentral, westcentralus |
| Standard_D32ds_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D32lds_v5 | qatarcentral |
| Standard_D32ls_v5 | qatarcentral |
| Standard_D32pds_v5 | northcentralus |
| Standard_D32plds_v5 | northcentralus |
| Standard_D32pls_v5 | northcentralus |
| Standard_D32ps_v5 | northcentralus |
| Standard_D32s_v5 | francecentral, qatarcentral |
| Standard_D48_v5 | francecentral |
| Standard_D48a_v4 | southindia, westcentralus |
| Standard_D48ads_v5 | westcentralus |
| Standard_D48as_v4 | southindia, westcentralus |
| Standard_D48as_v5 | westcentralus |
| Standard_D48d_v4 | qatarcentral |
| Standard_D48d_v5 | qatarcentral, westcentralus |
| Standard_D48ds_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D48lds_v5 | qatarcentral |
| Standard_D48ls_v5 | qatarcentral |
| Standard_D48pds_v5 | northcentralus |
| Standard_D48plds_v5 | northcentralus |
| Standard_D48pls_v5 | northcentralus |
| Standard_D48ps_v5 | northcentralus |
| Standard_D48s_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D4_v5 | francecentral, westcentralus |
| Standard_D4a_v4 | southindia, westcentralus |
| Standard_D4ads_v5 | westcentralus |
| Standard_D4as_v4 | southindia, westcentralus |
| Standard_D4as_v5 | westcentralus |
| Standard_D4d_v4 | qatarcentral |
| Standard_D4d_v5 | qatarcentral, westcentralus |
| Standard_D4ds_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D4lds_v5 | qatarcentral |
| Standard_D4ls_v5 | qatarcentral |
| Standard_D4pds_v5 | northcentralus |
| Standard_D4plds_v5 | northcentralus |
| Standard_D4pls_v5 | northcentralus |
| Standard_D4ps_v5 | northcentralus |
| Standard_D4s_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D64_v5 | francecentral |
| Standard_D64a_v4 | southindia, westcentralus |
| Standard_D64ads_v5 | westcentralus |
| Standard_D64as_v4 | southindia, westcentralus |
| Standard_D64d_v4 | qatarcentral |
| Standard_D64d_v5 | qatarcentral, westcentralus |
| Standard_D64ds_v4 | westindia |
| Standard_D64ds_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D64lds_v5 | qatarcentral |
| Standard_D64ls_v5 | qatarcentral |
| Standard_D64pds_v5 | northcentralus |
| Standard_D64pds_v6 | brazilsouth |
| Standard_D64plds_v5 | northcentralus |
| Standard_D64plds_v6 | brazilsouth |
| Standard_D64pls_v5 | northcentralus |
| Standard_D64pls_v6 | brazilsouth |
| Standard_D64ps_v5 | northcentralus |
| Standard_D64ps_v6 | brazilsouth |
| Standard_D64s_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D8_v5 | francecentral, westcentralus |
| Standard_D8a_v4 | southindia, westcentralus |
| Standard_D8ads_v5 | westcentralus |
| Standard_D8as_v4 | southindia, westcentralus |
| Standard_D8as_v5 | westcentralus |
| Standard_D8d_v4 | qatarcentral |
| Standard_D8d_v5 | qatarcentral, westcentralus |
| Standard_D8ds_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D8lds_v5 | qatarcentral |
| Standard_D8ls_v5 | qatarcentral |
| Standard_D8pds_v5 | northcentralus |
| Standard_D8plds_v5 | northcentralus |
| Standard_D8pls_v5 | northcentralus |
| Standard_D8ps_v5 | northcentralus |
| Standard_D8s_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D96_v5 | francecentral, westcentralus |
| Standard_D96a_v4 | southindia |
| Standard_D96ads_v5 | westcentralus |
| Standard_D96as_v4 | southindia, westcentralus |
| Standard_D96as_v5 | westcentralus |
| Standard_D96d_v5 | qatarcentral, westcentralus |
| Standard_D96ds_v5 | francecentral, qatarcentral, westcentralus |
| Standard_D96lds_v5 | qatarcentral |
| Standard_D96ls_v5 | qatarcentral |
| Standard_D96pds_v6 | brazilsouth |
| Standard_D96plds_v6 | brazilsouth |
| Standard_D96pls_v6 | brazilsouth |
| Standard_D96ps_v6 | brazilsouth |
| Standard_D96s_v5 | francecentral, qatarcentral, westcentralus |
| Standard_DC16ds_v3 | westus |
| Standard_DC16s_v3 | westus |
| Standard_DC1ds_v3 | westus |
| Standard_DC1s_v3 | westus |
| Standard_DC24ds_v3 | westus |
| Standard_DC24s_v3 | westus |
| Standard_DC2ds_v3 | westus |
| Standard_DC2s_v3 | westus |
| Standard_DC32ds_v3 | westus |
| Standard_DC32s_v3 | westus |
| Standard_DC48ds_v3 | westus |
| Standard_DC48s_v3 | westus |
| Standard_DC4ds_v3 | westus |
| Standard_DC4s_v3 | westus |
| Standard_DC8ds_v3 | westus |
| Standard_DC8s_v3 | westus |
| Standard_E104i_v5 | westcentralus |
| Standard_E104id_v5 | qatarcentral, westcentralus |
| Standard_E104ids_v5 | westcentralus |
| Standard_E104is_v5 | westcentralus |
| Standard_E112iads_v5 | canadacentral, mexicocentral |
| Standard_E112ias_v5 | canadacentral, mexicocentral |
| Standard_E112ibds_v5 | westcentralus |
| Standard_E112ibs_v5 | westcentralus |
| Standard_E16-4as_v4 | southindia |
| Standard_E16-4s_v5 | westcentralus |
| Standard_E16-8as_v4 | southindia |
| Standard_E16-8ds_v5 | westcentralus |
| Standard_E16-8s_v5 | westcentralus |
| Standard_E16_v5 | westcentralus |
| Standard_E16a_v4 | southindia |
| Standard_E16as_v4 | southindia |
| Standard_E16bds_v5 | westcentralus |
| Standard_E16bs_v5 | westcentralus |
| Standard_E16d_v5 | qatarcentral, westcentralus |
| Standard_E16ds_v5 | westcentralus |
| Standard_E16pds_v5 | australiaeast, canadacentral, centralindia, centralus, eastasia, eastus, eastus2, francecentral, germanywestcentral, japaneast, northcentralus, northeurope, southcentralus, southeastasia, swedencentral, uksouth, westcentralus, westeurope, westus, westus2, westus3 |
| Standard_E16pds_v6 | brazilsouth |
| Standard_E16ps_v5 | northcentralus |
| Standard_E16ps_v6 | brazilsouth |
| Standard_E16s_v5 | westcentralus |
| Standard_E20_v5 | westcentralus |
| Standard_E20a_v4 | southindia |
| Standard_E20as_v4 | southindia |
| Standard_E20d_v5 | qatarcentral, westcentralus |
| Standard_E20ds_v5 | westcentralus |
| Standard_E20pds_v5 | australiaeast, canadacentral, centralindia, centralus, eastasia, eastus, eastus2, francecentral, germanywestcentral, japaneast, northcentralus, northeurope, southcentralus, southeastasia, swedencentral, uksouth, westcentralus, westeurope, westus, westus2, westus3 |
| Standard_E20ps_v5 | northcentralus |
| Standard_E20s_v5 | westcentralus |
| Standard_E2_v5 | westcentralus |
| Standard_E2a_v4 | southindia |
| Standard_E2as_v4 | southindia |
| Standard_E2bds_v5 | westcentralus |
| Standard_E2bs_v5 | westcentralus |
| Standard_E2d_v5 | qatarcentral, westcentralus |
| Standard_E2ds_v5 | westcentralus |
| Standard_E2pds_v5 | australiaeast, canadacentral, centralindia, centralus, eastasia, eastus, eastus2, francecentral, germanywestcentral, japaneast, northcentralus, northeurope, southcentralus, southeastasia, swedencentral, uksouth, westcentralus, westeurope, westus, westus3 |
| Standard_E2ps_v5 | northcentralus |
| Standard_E2s_v5 | westcentralus |
| Standard_E32-16as_v4 | southindia |
| Standard_E32-16ds_v5 | westcentralus |
| Standard_E32-16s_v5 | westcentralus |
| Standard_E32-8as_v4 | southindia |
| Standard_E32-8ds_v5 | westcentralus |
| Standard_E32-8s_v5 | westcentralus |
| Standard_E32_v5 | westcentralus |
| Standard_E32a_v4 | southindia |
| Standard_E32as_v4 | southindia |
| Standard_E32bds_v5 | westcentralus |
| Standard_E32bs_v5 | westcentralus |
| Standard_E32d_v5 | qatarcentral |
| Standard_E32pds_v5 | australiaeast, canadacentral, centralindia, centralus, eastasia, eastus, eastus2, francecentral, germanywestcentral, japaneast, northcentralus, northeurope, southcentralus, southeastasia, swedencentral, uksouth, westcentralus, westeurope, westus, westus2, westus3 |
| Standard_E32ps_v5 | northcentralus |
| Standard_E32s_v5 | westcentralus |
| Standard_E4-2as_v4 | southindia |
| Standard_E4-2ds_v5 | westcentralus |
| Standard_E4-2s_v5 | westcentralus |
| Standard_E48_v5 | westcentralus |
| Standard_E48a_v4 | southindia |
| Standard_E48as_v4 | southindia |
| Standard_E48bds_v5 | westcentralus |
| Standard_E48bs_v5 | westcentralus |
| Standard_E48d_v5 | qatarcentral, westcentralus |
| Standard_E4_v5 | westcentralus |
| Standard_E4a_v4 | southindia |
| Standard_E4as_v4 | southindia |
| Standard_E4bds_v5 | westcentralus |
| Standard_E4bs_v5 | westcentralus |
| Standard_E4d_v5 | qatarcentral, westcentralus |
| Standard_E4ds_v5 | westcentralus |
| Standard_E4pds_v5 | australiaeast, canadacentral, centralindia, centralus, eastasia, eastus, francecentral, germanywestcentral, japaneast, northcentralus, northeurope, southcentralus, southeastasia, swedencentral, uksouth, westcentralus, westeurope, westus, westus2, westus3 |
| Standard_E4ps_v5 | northcentralus |
| Standard_E4s_v5 | westcentralus |
| Standard_E64-16as_v4 | southindia |
| Standard_E64-16s_v5 | westcentralus |
| Standard_E64-32as_v4 | southindia |
| Standard_E64-32ds_v5 | westcentralus |
| Standard_E64-32s_v5 | westcentralus |
| Standard_E64_v5 | westcentralus |
| Standard_E64a_v4 | southindia |
| Standard_E64as_v4 | southindia |
| Standard_E64bds_v5 | westcentralus |
| Standard_E64bs_v5 | westcentralus |
| Standard_E64d_v5 | qatarcentral, westcentralus |
| Standard_E64ds_v5 | westcentralus |
| Standard_E64s_v5 | westcentralus |
| Standard_E8-2as_v4 | southindia |
| Standard_E8-2ds_v5 | westcentralus |
| Standard_E8-2s_v5 | westcentralus |
| Standard_E8-4as_v4 | southindia |
| Standard_E8-4s_v5 | westcentralus |
| Standard_E8_v5 | westcentralus |
| Standard_E8a_v4 | southindia |
| Standard_E8as_v4 | southindia |
| Standard_E8bds_v5 | westcentralus |
| Standard_E8bs_v5 | westcentralus |
| Standard_E8d_v5 | qatarcentral |
| Standard_E8ds_v5 | westcentralus |
| Standard_E8pds_v5 | australiaeast, canadacentral, centralindia, centralus, eastasia, eastus, eastus2, francecentral, germanywestcentral, japaneast, northcentralus, northeurope, southcentralus, southeastasia, swedencentral, uksouth, westcentralus, westeurope, westus, westus2, westus3 |
| Standard_E8ps_v5 | northcentralus |
| Standard_E8s_v5 | westcentralus |
| Standard_E96-24as_v4 | canadacentral, southindia |
| Standard_E96-24as_v5 | canadacentral |
| Standard_E96-24ds_v5 | westcentralus |
| Standard_E96-24s_v5 | westcentralus |
| Standard_E96-48as_v4 | canadacentral, southindia |
| Standard_E96-48as_v5 | canadacentral |
| Standard_E96-48ds_v5 | westcentralus |
| Standard_E96-48s_v5 | westcentralus |
| Standard_E96_v5 | westcentralus |
| Standard_E96a_v4 | canadacentral, southindia |
| Standard_E96as_v4 | canadacentral, southindia |
| Standard_E96as_v5 | canadacentral |
| Standard_E96as_v6 | canadacentral |
| Standard_E96bds_v5 | westcentralus |
| Standard_E96bs_v5 | westcentralus |
| Standard_E96d_v5 | qatarcentral, westcentralus |
| Standard_E96ds_v5 | westcentralus |
| Standard_E96ias_v4 | southindia |
| Standard_E96pds_v6 | brazilsouth |
| Standard_E96ps_v6 | brazilsouth |
| Standard_E96s_v5 | westcentralus |
| Standard_EC96ads_v6 | southafricawest |
| Standard_EC96as_v6 | southafricawest |
| Standard_HB120-16rs_v2 | qatarcentral |
| Standard_HB120-16rs_v3 | japaneast |
| Standard_HB120-32rs_v2 | qatarcentral |
| Standard_HB120-32rs_v3 | japaneast |
| Standard_HB120-64rs_v2 | qatarcentral |
| Standard_HB120-64rs_v3 | japaneast |
| Standard_HB120-96rs_v2 | qatarcentral |
| Standard_HB120-96rs_v3 | japaneast |
| Standard_HB120rs_v2 | qatarcentral |
| Standard_HB120rs_v3 | japaneast |
| Standard_HB368-144rs_v5 | eastus, southcentralus |
| Standard_HB368-192rs_v5 | eastus, southcentralus |
| Standard_HB368-240rs_v5 | eastus, southcentralus |
| Standard_HB368-288rs_v5 | eastus, southcentralus |
| Standard_HB368-336rs_v5 | eastus, southcentralus |
| Standard_HB368-48rs_v5 | eastus, southcentralus |
| Standard_HB368-96rs_v5 | southcentralus |
| Standard_HB368rs_v5 | eastus, southcentralus |
| Standard_HC44-32rs | brazilsouth |
| Standard_HC44rs | brazilsouth |
| Standard_M128 | qatarcentral |
| Standard_M128-32ms | qatarcentral |
| Standard_M128-64ms | qatarcentral |
| Standard_M128dms_v2 | qatarcentral, ukwest |
| Standard_M128ds_v2 | qatarcentral, ukwest |
| Standard_M128m | qatarcentral |
| Standard_M128ms | qatarcentral |
| Standard_M128ms_v2 | qatarcentral |
| Standard_M128s | qatarcentral |
| Standard_M128s_v2 | qatarcentral |
| Standard_M16-4ms | qatarcentral |
| Standard_M16-8ms | qatarcentral |
| Standard_M16ms | qatarcentral |
| Standard_M192idms_v2 | qatarcentral |
| Standard_M192ids_v2 | qatarcentral |
| Standard_M192ims_v2 | qatarcentral |
| Standard_M192is_v2 | qatarcentral |
| Standard_M208ms_v2 | switzerlandnorth |
| Standard_M208s_v2 | switzerlandnorth |
| Standard_M32-16ms | qatarcentral |
| Standard_M32-8ms | qatarcentral |
| Standard_M32dms_v2 | qatarcentral, ukwest |
| Standard_M32ls | qatarcentral |
| Standard_M32ms | qatarcentral |
| Standard_M32ms_v2 | qatarcentral |
| Standard_M32ts | qatarcentral |
| Standard_M416-208ms_v2 | switzerlandnorth |
| Standard_M416-208s_v2 | switzerlandnorth |
| Standard_M416ms_v2 | switzerlandnorth |
| Standard_M416s_v2 | switzerlandnorth |
| Standard_M64 | qatarcentral |
| Standard_M64-16ms | qatarcentral |
| Standard_M64-32ms | qatarcentral |
| Standard_M64dms_v2 | qatarcentral, ukwest |
| Standard_M64ds_v2 | qatarcentral, ukwest |
| Standard_M64ls | qatarcentral |
| Standard_M64m | qatarcentral |
| Standard_M64ms | qatarcentral |
| Standard_M64ms_v2 | qatarcentral |
| Standard_M64s | qatarcentral |
| Standard_M64s_v2 | qatarcentral |
| Standard_M8-2ms | qatarcentral |
| Standard_M8-4ms | qatarcentral |
| Standard_M8ms | qatarcentral |
| Standard_NC16ads_A10_v4 | westus3 |
| Standard_NC16as_T4_v3 | swedencentral, westus3 |
| Standard_NC24ads_A100_v4 | canadacentral, centralus, eastus, southeastasia |
| Standard_NC32ads_A10_v4 | westus3 |
| Standard_NC40ads_H100_v5 | centralindia, koreacentral |
| Standard_NC48ads_A100_v4 | brazilsouth, canadacentral, centralindia, centralus, eastus, germanywestcentral, japaneast, koreacentral, northeurope, southcentralus, southeastasia, swedencentral, switzerlandnorth, westeurope, westus, westus2, westus3 |
| Standard_NC4as_T4_v3 | swedencentral, westus3 |
| Standard_NC64as_T4_v3 | eastus, germanywestcentral, japanwest, southindia, spaincentral, swedencentral, westus3 |
| Standard_NC80adis_H100_v5 | centralindia, centralus, japaneast, koreacentral, westus2 |
| Standard_NC8ads_A10_v4 | westus3 |
| Standard_NC8as_T4_v3 | swedencentral, westus3 |
| Standard_NC96ads_A100_v4 | brazilsouth, canadacentral, centralindia, centralus, eastus, germanywestcentral, japaneast, koreacentral, northeurope, southcentralus, southeastasia, swedencentral, switzerlandnorth, westeurope, westus, westus2, westus3 |
| Standard_ND128isr_NDR_GB200_v6 | australiaeast, centralus, eastus, westus2, westus3 |
| Standard_ND96amsr_A100_v4 | eastus, eastus2, southcentralus, westeurope |
| Standard_ND96asr_v4 | eastus, southcentralus, westeurope, westus2 |
| Standard_ND96is_MI300X_v5 | francecentral |
| Standard_ND96isr_MI300X_v5 | francecentral |
| Standard_NV12ads_A10_v5 | brazilsouth, westeurope |
| Standard_NV12s_v2 | southcentralus, westus |
| Standard_NV18ads_A10_v5 | brazilsouth, centralus, westeurope |
| Standard_NV24s_v2 | southcentralus, westus |
| Standard_NV36adms_A10_v5 | brazilsouth, canadacentral, centralindia, centralus, westeurope, westus |
| Standard_NV36ads_A10_v5 | brazilsouth, centralindia, westeurope |
| Standard_NV6ads_A10_v5 | brazilsouth, westeurope |
| Standard_NV6s_v2 | southcentralus, westus |
| Standard_NV72ads_A10_v5 | brazilsouth, canadacentral, centralindia, centralus, westeurope |
| Standard_PB6s | eastus, westus2 |

## Per-Region Summary

| Region | VM SKUs | LowPriCapable | Has Spot Price | LowPri=True,NoSpot | LowPri=False,HasSpot |
|--------|---------|---------------|----------------|--------------------|-----------------------|
| australiacentral | 826 | 800 | 698 | 0 | 0 |
| australiacentral2 | 725 | 699 | 645 | 0 | 0 |
| australiaeast | 1235 | 1199 | 1205 | 0 | 7 |
| australiasoutheast | 1075 | 1048 | 854 | 0 | 0 |
| austriaeast | 733 | 717 | 607 | 10 | 0 |
| brazilsouth | 1044 | 991 | 913 | 79 | 22 |
| brazilsoutheast | 702 | 676 | 630 | 30 | 0 |
| canadacentral | 1053 | 1006 | 1016 | 11 | 21 |
| canadaeast | 914 | 827 | 811 | 0 | 0 |
| centralindia | 1124 | 1082 | 1074 | 0 | 13 |
| centralus | 1304 | 1263 | 1257 | 0 | 14 |
| chilecentral | 770 | 754 | 563 | 10 | 0 |
| eastasia | 1049 | 1016 | 948 | 0 | 6 |
| eastus | 1379 | 1328 | 1260 | 57 | 21 |
| eastus2 | 1327 | 1289 | 1130 | 139 | 6 |
| francecentral | 1145 | 1084 | 1098 | 0 | 32 |
| francesouth | 676 | 650 | 600 | 0 | 0 |
| germanynorth | 812 | 786 | 723 | 0 | 0 |
| germanywestcentral | 1301 | 1265 | 1268 | 0 | 9 |
| indonesiacentral | 856 | 839 | 769 | 10 | 0 |
| israelcentral | 717 | 691 | 645 | 0 | 0 |
| italynorth | 1134 | 1118 | 1017 | 10 | 0 |
| japaneast | 1289 | 1246 | 1258 | 0 | 14 |
| japanwest | 973 | 946 | 852 | 0 | 1 |
| jioindiacentral | 695 | 0 | 0 | 0 | 0 |
| jioindiawest | 680 | 0 | 0 | 0 | 0 |
| koreacentral | 986 | 956 | 925 | 0 | 4 |
| koreasouth | 656 | 630 | 579 | 0 | 0 |
| malaysiawest | 784 | 768 | 667 | 10 | 0 |
| mexicocentral | 811 | 793 | 685 | 36 | 2 |
| newzealandnorth | 844 | 828 | 742 | 10 | 0 |
| northcentralus | 1029 | 962 | 919 | 12 | 40 |
| northeurope | 1298 | 1260 | 1249 | 0 | 8 |
| norwayeast | 871 | 842 | 789 | 0 | 0 |
| norwaywest | 748 | 719 | 624 | 0 | 0 |
| polandcentral | 995 | 979 | 917 | 10 | 0 |
| qatarcentral | 391 | 265 | 365 | 0 | 100 |
| southafricanorth | 872 | 843 | 825 | 0 | 0 |
| southafricawest | 907 | 879 | 781 | 0 | 2 |
| southcentralus | 1292 | 1245 | 1243 | 4 | 21 |
| southeastasia | 1271 | 1233 | 1235 | 0 | 9 |
| southindia | 888 | 815 | 828 | 0 | 47 |
| spaincentral | 997 | 980 | 961 | 10 | 1 |
| swedencentral | 1252 | 1213 | 1224 | 0 | 12 |
| swedensouth | 523 | 497 | 448 | 0 | 0 |
| switzerlandnorth | 985 | 945 | 921 | 0 | 8 |
| switzerlandwest | 867 | 841 | 760 | 0 | 0 |
| uaecentral | 876 | 849 | 754 | 0 | 0 |
| uaenorth | 1106 | 1077 | 1077 | 0 | 0 |
| uksouth | 1227 | 1191 | 1159 | 35 | 6 |
| ukwest | 921 | 890 | 861 | 0 | 5 |
| westcentralus | 916 | 740 | 767 | 53 | 133 |
| westeurope | 1314 | 1269 | 1283 | 0 | 16 |
| westindia | 199 | 172 | 173 | 0 | 1 |
| westus | 1329 | 1275 | 1298 | 0 | 28 |
| westus2 | 1327 | 1285 | 1039 | 196 | 11 |
| westus3 | 1274 | 1231 | 1244 | 3 | 16 |
