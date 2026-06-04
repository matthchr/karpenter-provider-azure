# Pricing Mismatch Report

Generated: 2026-09-01T21:38:32Z

This report identifies VM sizes where:
- A region has **on-demand** pricing but **NO spot** pricing
- A region has **spot** pricing but **NO on-demand** pricing

## Summary

- Regions analyzed: 57
- SKUs with on-demand but NO spot: 6175 entries across regions
- SKUs with spot but NO on-demand: 818 entries across regions

## On-Demand Only (no spot pricing)

| SKU | Regions (1412 unique SKUs) | On-Demand Price (first region) |
|-----|---------|--------|
| Basic_A0 | brazilsouth, canadacentral | $0.0220 |
| Basic_A3 | brazilsouth | $0.2320 |
| DCadsv5 Type 1 | centralindia, centralus, eastus, jioindiacentral, jioindiawest (+4 more) | $4.1330 |
| DCadsv6_Type1 | australiacentral, australiaeast, australiasoutheast, brazilsouth, brazilsoutheast (+43 more) | $12.5140 |
| DCasv5 Type 1 | centralindia, centralus, eastus, jioindiacentral, jioindiawest (+3 more) | $3.4220 |
| DCasv6_Type1 | australiacentral, australiaeast, australiasoutheast, brazilsouth, brazilsoutheast (+43 more) | $9.9320 |
| DCdsv3 Type1 | australiaeast, canadacentral, centralindia, centralus, eastus (+17 more) | $7.4980 |
| DCsv2 Type 1 | australiaeast, australiasoutheast, canadacentral, canadaeast, eastus (+7 more) | $1.0560 |
| DCsv3 Type1 | australiaeast, canadacentral, centralindia, centralus, eastus (+17 more) | $6.3360 |
| Dadsv5_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+46 more) | $8.0080 |
| Dasv4_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+47 more) | $6.6000 |
| Dasv4_Type2 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+47 more) | $6.6000 |
| Dasv5_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+47 more) | $6.6530 |
| Dasv6_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+48 more) | $9.0290 |
| Ddsv4_Type 1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+46 more) | $4.9980 |
| Ddsv4_Type2 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+44 more) | $5.9360 |
| Ddsv5_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+49 more) | $7.4980 |
| Ddsv6_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, brazilsouth (+46 more) | $16.4460 |
| Dsv3_Type1 | israelcentral, italynorth, polandcentral | $3.9600 |
| Dsv3_Type2 | israelcentral, italynorth, polandcentral | $4.4600 |
| Dsv3_Type3 | australiacentral, australiacentral2, australiaeast, australiasoutheast, brazilsouth (+48 more) | $6.5990 |
| Dsv3_Type4 | australiacentral, australiacentral2, australiaeast, australiasoutheast, brazilsouth (+47 more) | $8.2490 |
| Dsv4_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+45 more) | $5.2800 |
| Dsv4_Type2 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+41 more) | $6.3360 |
| Dsv5_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+51 more) | $6.6000 |
| Dsv6_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, brazilsouth (+47 more) | $13.3060 |
| ECadsv5 Type 1 | centralindia, centralus, jioindiacentral, jioindiawest, northeurope (+2 more) | $4.8240 |
| ECadsv6_Type1 | australiacentral, australiaeast, australiasoutheast, brazilsouth, brazilsoutheast (+42 more) | $15.8730 |
| ECasv5 Type 1 | centralindia, centralus, eastus, jioindiacentral, jioindiawest (+3 more) | $4.1530 |
| ECasv6_Type1 | australiacentral, australiaeast, australiasoutheast, brazilsouth, brazilsoutheast (+42 more) | $13.0510 |
| Eadsv5_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+45 more) | $8.3950 |
| Easv4_Type1 | australiacentral, australiacentral2, australiaeast, austriaeast, brazilsouth (+45 more) | $7.9730 |
| Easv4_Type2 | australiacentral, australiacentral2, australiaeast, austriaeast, brazilsouth (+45 more) | $7.9730 |
| Easv5_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+45 more) | $7.1810 |
| Easv6_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+48 more) | $11.8800 |
| Ebdsv5-Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+51 more) | $7.0750 |
| Ebsv5-Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+50 more) | $6.2660 |
| Edsv4_Type 1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+46 more) | $6.1250 |
| Edsv4_Type2 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+42 more) | $8.7360 |
| Edsv5_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+50 more) | $9.1870 |
| Esv3_Type1 | israelcentral, italynorth, polandcentral | $5.2100 |
| Esv3_Type2 | israelcentral, italynorth, polandcentral | $5.2100 |
| Esv3_Type3 | australiacentral, australiacentral2, australiaeast, australiasoutheast, brazilsouth (+49 more) | $6.4340 |
| Esv3_Type4 | australiacentral, australiacentral2, australiaeast, australiasoutheast, brazilsouth (+49 more) | $8.7170 |
| Esv4_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+43 more) | $5.3150 |
| Esv4_Type2 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+43 more) | $6.9760 |
| Esv5_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+50 more) | $7.9730 |
| Esv6_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, brazilsouth (+46 more) | $18.4180 |
| FXmds Type1 | westus2 | $4.9100 |
| Fsv2 Type3 | australiacentral, australiaeast, australiasoutheast, austriaeast, brazilsouth (+43 more) | $5.1280 |
| Fsv2_Type2 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+50 more) | $4.3960 |
| Fsv2_Type4 | australiacentral, australiaeast, australiasoutheast, austriaeast, brazilsouth (+41 more) | $5.8610 |
| Lasv3_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, brazilsoutheast (+30 more) | $8.2280 |
| Lsv2_Type1 | australiaeast, centralindia, centralus, eastus, eastus2 (+12 more) | $8.2280 |
| Lsv3_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+48 more) | $9.1960 |
| Mdmsv2MedMem _Type1 | australiaeast, australiasoutheast, brazilsouth, brazilsoutheast, canadacentral (+36 more) | $42.5710 |
| Mdsv2MedMem_Type1 | australiaeast, australiasoutheast, brazilsouth, brazilsoutheast, canadacentral (+36 more) | $21.2770 |
| Mdsv3MedMem_Type1 | eastus, northeurope, westeurope, westus2 | $29.1170 |
| Mmsv2MedMem-Type1 | australiaeast, australiasoutheast, brazilsouth, brazilsoutheast, canadacentral (+34 more) | $41.8990 |
| Ms_Type1 | australiaeast, brazilsouth, canadacentral, centralindia, centralus (+27 more) | $21.2700 |
| Msm_Type1 | australiaeast, brazilsouth, canadacentral, centralindia, centralus (+28 more) | $42.5770 |
| Msmv2_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, brazilsouth (+31 more) | $158.1400 |
| Msv2MedMem Type1 | australiaeast, australiasoutheast, brazilsouth, brazilsoutheast, canadacentral (+34 more) | $20.6060 |
| Msv2_Type1 | australiacentral, australiacentral2, australiaeast, australiasoutheast, brazilsouth (+31 more) | $79.0800 |
| Msv3MedMem_Type1 | eastus, eastus2, northeurope, westeurope, westus2 | $27.4010 |
| NDamsrA100v4_Type1 | eastus, southcentralus, switzerlandnorth, westeurope, westus3 | $36.0470 |
| NDasrA100v4_Type1 | eastus, southcentralus, westeurope, westus2, westus3 | $29.9170 |
| NVadsA10v5_Type1 | australiaeast, brazilsouth, canadacentral, centralindia, centralus (+28 more) | $10.3990 |
| NVasv4_Type1 | australiaeast, canadacentral, centralindia, chilecentral, eastus (+16 more) | $10.4060 |
| NVsv3_Type1 | eastus, eastus2, southcentralus, westeurope, westus (+1 more) | $5.0160 |
| Standard_A1_v2 | eastus2, jioindiacentral, jioindiawest, westus2 | $0.0430 |
| Standard_A2_v2 | jioindiacentral, jioindiawest | $0.0980 |
| Standard_A2m_v2 | jioindiacentral, jioindiawest | $0.1300 |
| Standard_A4_v2 | brazilsouth, jioindiacentral, jioindiawest | $0.2700 |
| Standard_A4m_v2 | eastus, jioindiacentral, jioindiawest | $0.2370 |
| Standard_A5 | eastus | $0.2500 |
| Standard_A8_v2 | jioindiacentral, jioindiawest | $0.4330 |
| Standard_A8m_v2 | jioindiacentral, jioindiawest | $0.5190 |
| Standard_B12ms | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+52 more) | $0.6340 |
| Standard_B16als_v2 | jioindiacentral, jioindiawest, westus2 | $0.4220 |
| Standard_B16as_v2 | brazilsouth, jioindiacentral, jioindiawest, westus2 | $0.9680 |
| Standard_B16ls_v2 | jioindiacentral, jioindiawest | $0.6350 |
| Standard_B16ms | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+50 more) | $0.8450 |
| Standard_B16pls_v2 | jioindiacentral, jioindiawest | $0.3170 |
| Standard_B16ps_v2 | jioindiacentral, jioindiawest | $0.3580 |
| Standard_B16s_v2 | jioindiacentral, jioindiawest | $0.7900 |
| Standard_B1ls | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+51 more) | $0.0066 |
| Standard_B1ms | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+51 more) | $0.0264 |
| Standard_B1s | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+52 more) | $0.0132 |
| Standard_B20ms | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+51 more) | $1.0560 |
| Standard_B2als_v2 | jioindiacentral, jioindiawest | $0.0246 |
| Standard_B2as_v2 | eastus2, jioindiacentral, jioindiawest | $0.0844 |
| Standard_B2ats_v2 | jioindiacentral, jioindiawest | $0.0062 |
| Standard_B2ls_v2 | jioindiacentral, jioindiawest | $0.0540 |
| Standard_B2ms | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+52 more) | $0.1060 |
| Standard_B2pls_v2 | jioindiacentral, jioindiawest | $0.0224 |
| Standard_B2ps_v2 | jioindiacentral, jioindiawest | $0.0448 |
| Standard_B2pts_v2 | eastus2, jioindiacentral, jioindiawest, westcentralus | $0.0176 |
| Standard_B2s | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+51 more) | $0.0528 |
| Standard_B2s_v2 | jioindiacentral, jioindiawest | $0.0896 |
| Standard_B2ts_v2 | eastus2, jioindiacentral, jioindiawest | $0.0104 |
| Standard_B32als_v2 | jioindiacentral, jioindiawest | $0.8440 |
| Standard_B32as_v2 | jioindiacentral, jioindiawest | $0.7870 |
| Standard_B32ls_v2 | jioindiacentral, jioindiawest | $1.4170 |
| Standard_B32s_v2 | jioindiacentral, jioindiawest, westus2 | $1.5810 |
| Standard_B4als_v2 | jioindiacentral, jioindiawest | $0.0872 |
| Standard_B4as_v2 | jioindiacentral, jioindiawest | $0.0984 |
| Standard_B4ls_v2 | jioindiacentral, jioindiawest | $0.1770 |
| Standard_B4ms | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+51 more) | $0.2110 |
| Standard_B4pls_v2 | jioindiacentral, jioindiawest, uksouth | $0.0793 |
| Standard_B4ps_v2 | jioindiacentral, jioindiawest, westus2 | $0.0896 |
| Standard_B4s_v2 | jioindiacentral, jioindiawest | $0.1790 |
| Standard_B8als_v2 | jioindiacentral, jioindiawest | $0.1740 |
| Standard_B8as_v2 | jioindiacentral, jioindiawest, mexicocentral | $0.1970 |
| Standard_B8ls_v2 | brazilsouth, jioindiacentral, jioindiawest, westus2 | $0.5130 |
| Standard_B8ms | australiacentral, australiacentral2, australiaeast, australiasoutheast, austriaeast (+51 more) | $0.4220 |
| Standard_B8pls_v2 | eastus, jioindiacentral, jioindiawest | $0.2380 |
| Standard_B8ps_v2 | jioindiacentral, jioindiawest, westcentralus | $0.1790 |
| Standard_B8s_v2 | jioindiacentral, jioindiawest | $0.3580 |
| Standard_D1 | eastus | $0.0770 |
| Standard_D11_v2 | jioindiacentral, jioindiawest | $0.1890 |
| Standard_D11_v2_Promo | francecentral | $0.2340 |
| Standard_D128ads_v7 | jioindiawest | $4.7230 |
| Standard_D128alds_v7 | eastus, jioindiawest | $6.0930 |
| Standard_D128als_v7 | jioindiawest | $3.3280 |
| Standard_D128as_v7 | jioindiawest | $3.7500 |
| Standard_D128ds_v6 | jioindiacentral, jioindiawest | $8.3890 |
| Standard_D128lds_v6 | jioindiacentral, jioindiawest | $7.0600 |
| Standard_D128ls_v6 | jioindiacentral, jioindiawest, westus2 | $5.7130 |
| Standard_D128nds_v6 | jioindiacentral, jioindiawest, westcentralus | $10.8580 |
| Standard_D128nlds_v6 | brazilsouth, jioindiacentral, jioindiawest, westus2 | $13.7770 |
| Standard_D128nls_v6 | jioindiacentral, jioindiawest | $7.2580 |
| Standard_D128ns_v6 | jioindiacentral, jioindiawest | $9.2850 |
| Standard_D128s_v6 | jioindiacentral, jioindiawest | $6.7880 |
| Standard_D128s_v7 | eastus2, westus2 | $8.4670 |
| Standard_D12_v2 | eastus, jioindiacentral, jioindiawest | $0.3710 |
| Standard_D12_v2_Promo | francecentral | $0.4690 |
| Standard_D13 | brazilsoutheast | $1.2210 |
| Standard_D13_v2 | jioindiacentral, jioindiawest, westus2 | $0.7580 |
| Standard_D13_v2_Promo | francecentral | $0.9370 |
| Standard_D14 | brazilsoutheast | $2.4430 |
| Standard_D14_v2 | jioindiacentral, jioindiawest | $1.5160 |
| Standard_D14_v2_Promo | francecentral | $1.8740 |
| Standard_D15_v2 | jioindiacentral, jioindiawest | $1.8950 |
| Standard_D15i_v2 | brazilsouth, eastus2, jioindiacentral, jioindiawest | $2.3490 |
| Standard_D160ads_v7 | jioindiawest, mexicocentral | $5.9040 |
| Standard_D160alds_v7 | eastus, jioindiawest | $7.6160 |
| Standard_D160als_v7 | jioindiawest | $4.1600 |
| Standard_D160as_v7 | jioindiawest | $4.6880 |
| Standard_D16_v3 | jioindiacentral, jioindiawest | $0.8400 |
| Standard_D16_v4 | jioindiacentral, jioindiawest | $0.8080 |
| Standard_D16_v5 | jioindiacentral, jioindiawest | $0.8080 |
| Standard_D16a_v4 | eastus2, jioindiacentral, jioindiawest | $0.7680 |
| Standard_D16ads_v5 | jioindiacentral, jioindiawest, uksouth | $0.5370 |
| Standard_D16ads_v6 | eastus2, jioindiacentral, jioindiawest | $0.9120 |
| Standard_D16ads_v7 | eastus2, jioindiawest | $0.9120 |
| Standard_D16alds_v6 | jioindiacentral, jioindiawest, westus2 | $0.4930 |
| Standard_D16alds_v7 | jioindiawest | $0.4930 |
| Standard_D16als_v6 | jioindiacentral, jioindiawest | $0.4160 |
| Standard_D16als_v7 | jioindiawest, westus2 | $0.4160 |
| Standard_D16as_v4 | jioindiacentral, jioindiawest, westcentralus | $0.4930 |
| Standard_D16as_v5 | jioindiacentral, jioindiawest, mexicocentral | $0.4440 |
| Standard_D16as_v6 | jioindiacentral, jioindiawest, westcentralus | $0.4690 |
| Standard_D16as_v7 | jioindiawest | $0.4690 |
| Standard_D16d_v4 | jioindiacentral, jioindiawest | $0.9760 |
| Standard_D16d_v5 | jioindiacentral, jioindiawest | $0.9760 |
| Standard_D16ds_v4 | jioindiacentral, jioindiawest | $0.9760 |
| Standard_D16ds_v5 | jioindiacentral, jioindiawest, westcentralus | $0.9760 |
| Standard_D16ds_v6 | jioindiacentral, jioindiawest, westcentralus | $1.0490 |
| Standard_D16lds_v5 | jioindiacentral, jioindiawest | $0.7920 |
| Standard_D16lds_v6 | eastus2, jioindiacentral, jioindiawest | $0.8830 |
| Standard_D16ls_v5 | eastus2, jioindiacentral, jioindiawest | $0.6800 |
| Standard_D16ls_v6 | jioindiacentral, jioindiawest, uksouth | $0.7140 |
| Standard_D16ls_v7 | eastus2 | $0.9370 |
| Standard_D16nds_v6 | jioindiacentral, jioindiawest | $1.3570 |
| Standard_D16nlds_v6 | jioindiacentral, jioindiawest, westcentralus | $1.1200 |
| Standard_D16nls_v6 | jioindiacentral, jioindiawest | $0.9070 |
| Standard_D16ns_v6 | jioindiacentral, jioindiawest | $1.1610 |
| Standard_D16pds_v5 | jioindiacentral, jioindiawest | $0.4830 |
| Standard_D16plds_v5 | jioindiacentral, jioindiawest | $0.3920 |
| Standard_D16pls_v5 | jioindiacentral, jioindiawest | $0.3410 |
| Standard_D16pls_v6 | westus2 | $0.4960 |
| Standard_D16ps_v5 | jioindiacentral, jioindiawest | $0.4050 |
| Standard_D16s_v3 | jioindiacentral, jioindiawest | $0.8400 |
| Standard_D16s_v4 | jioindiacentral, jioindiawest | $0.8080 |
| Standard_D16s_v5 | jioindiacentral, jioindiawest, westcentralus | $0.8080 |
| Standard_D16s_v6 | jioindiacentral, jioindiawest | $0.8480 |
| Standard_D192ds_v6 | jioindiacentral, jioindiawest | $12.5840 |
| Standard_D192ds_v7 | northcentralus | $15.6630 |
| Standard_D192s_v6 | jioindiacentral, jioindiawest, westus2 | $10.1820 |
| Standard_D1_v2 | jioindiacentral, jioindiawest, westcentralus | $0.0840 |
| Standard_D2 | eastus2, northcentralus | $0.1340 |
| Standard_D248ds_v7 | brazilsouth | $32.3710 |
| Standard_D2_v2 | brazilsouth, jioindiacentral, jioindiawest | $0.1710 |
| Standard_D2_v2_Promo | francecentral | $0.1750 |
| Standard_D2_v3 | jioindiacentral, jioindiawest | $0.1050 |
| Standard_D2_v4 | jioindiacentral, jioindiawest | $0.1010 |
| Standard_D2_v5 | jioindiacentral, jioindiawest | $0.1010 |
| Standard_D2a_v4 | jioindiacentral, jioindiawest | $0.0622 |
| Standard_D2ads_v5 | jioindiacentral, jioindiawest, westus2 | $0.0671 |
| Standard_D2ads_v6 | jioindiacentral, jioindiawest, westcentralus | $0.0738 |
| Standard_D2ads_v7 | jioindiawest, westus2 | $0.0738 |
| Standard_D2alds_v6 | jioindiacentral, jioindiawest, westus2 | $0.0616 |
| Standard_D2alds_v7 | jioindiawest | $0.0616 |
| Standard_D2als_v6 | jioindiacentral, jioindiawest | $0.0520 |
| Standard_D2als_v7 | jioindiawest, westus2 | $0.0520 |
| Standard_D2as_v4 | jioindiacentral, jioindiawest, westcentralus | $0.0622 |
| Standard_D2as_v5 | jioindiacentral, jioindiawest | $0.0556 |
| Standard_D2as_v6 | brazilsouth, jioindiacentral, jioindiawest | $0.1460 |
| Standard_D2as_v7 | jioindiawest, westus2 | $0.0586 |
| Standard_D2d_v4 | jioindiacentral, jioindiawest, westus2 | $0.1220 |
| Standard_D2d_v5 | jioindiacentral, jioindiawest | $0.1220 |
| Standard_D2ds_v4 | jioindiacentral, jioindiawest | $0.1220 |
| Standard_D2ds_v5 | jioindiacentral, jioindiawest | $0.1220 |
| Standard_D2ds_v6 | eastus2, jioindiacentral, jioindiawest | $0.1250 |
| Standard_D2lds_v5 | eastus2, jioindiacentral, jioindiawest | $0.0960 |
| Standard_D2lds_v6 | jioindiacentral, jioindiawest, westus2 | $0.1100 |
| Standard_D2ls_v5 | jioindiacentral, jioindiawest | $0.0850 |
| Standard_D2ls_v6 | jioindiacentral, jioindiawest | $0.0893 |
| Standard_D2ls_v7 | eastus | $0.1170 |
| Standard_D2nds_v6 | jioindiacentral, jioindiawest | $0.1700 |
| Standard_D2nlds_v6 | jioindiacentral, jioindiawest, westus2 | $0.1400 |
| Standard_D2nls_v6 | jioindiacentral, jioindiawest, uksouth | $0.1130 |
| Standard_D2ns_v6 | jioindiacentral, jioindiawest, westcentralus, westus2 | $0.1450 |
| Standard_D2pds_v5 | jioindiacentral, jioindiawest | $0.0604 |
| Standard_D2plds_v5 | jioindiacentral, jioindiawest | $0.0490 |
| Standard_D2plds_v6 | westus2 | $0.0780 |
| Standard_D2pls_v5 | jioindiacentral, jioindiawest | $0.0426 |
| Standard_D2ps_v5 | jioindiacentral, jioindiawest | $0.0506 |
| Standard_D2ps_v6 | eastus2 | $0.0702 |
| Standard_D2s_v3 | eastus2, jioindiacentral, jioindiawest | $0.0960 |
| Standard_D2s_v4 | jioindiacentral, jioindiawest | $0.1010 |
| Standard_D2s_v5 | jioindiacentral, jioindiawest | $0.1010 |
| Standard_D2s_v6 | jioindiacentral, jioindiawest | $0.1060 |
| Standard_D32-16s_v3 | jioindiacentral, jioindiawest, westus2 | $1.6800 |
| Standard_D32-8s_v3 | jioindiacentral, jioindiawest, westus2 | $1.6800 |
| Standard_D32_v3 | jioindiacentral, jioindiawest | $1.6800 |
| Standard_D32_v4 | brazilsouth, jioindiacentral, jioindiawest | $2.4480 |
| Standard_D32_v5 | jioindiacentral, jioindiawest, westus2 | $1.6160 |
| Standard_D32a_v4 | brazilsouth, jioindiacentral, jioindiawest, mexicocentral | $2.4480 |
| Standard_D32ads_v5 | brazilsouth, jioindiacentral, jioindiawest, westus3 | $2.6400 |
| Standard_D32ads_v6 | jioindiacentral, jioindiawest | $1.1810 |
| Standard_D32ads_v7 | jioindiawest | $1.1810 |
| Standard_D32alds_v6 | jioindiacentral, jioindiawest | $0.9860 |
| Standard_D32alds_v7 | jioindiawest, westus2 | $0.9860 |
| Standard_D32als_v6 | jioindiacentral, jioindiawest | $0.8320 |
| Standard_D32als_v7 | eastus2, jioindiawest, mexicocentral | $1.2860 |
| Standard_D32as_v4 | jioindiacentral, jioindiawest | $0.9870 |
| Standard_D32as_v5 | jioindiacentral, jioindiawest | $2.3610 |
| Standard_D32as_v6 | brazilsouth, jioindiacentral, jioindiawest | $2.3300 |
| Standard_D32as_v7 | jioindiawest | $0.9380 |
| Standard_D32d_v4 | jioindiacentral, jioindiawest | $1.9520 |
| Standard_D32d_v5 | eastus, jioindiacentral, jioindiawest | $1.8080 |
| Standard_D32ds_v4 | jioindiacentral, jioindiawest | $1.9520 |
| Standard_D32ds_v5 | eastus, jioindiacentral, jioindiawest | $1.8080 |
| Standard_D32ds_v6 | jioindiacentral, jioindiawest | $2.0970 |
| Standard_D32lds_v5 | jioindiacentral, jioindiawest, westcentralus | $1.5840 |
| Standard_D32lds_v6 | jioindiacentral, jioindiawest | $1.7650 |
| Standard_D32ls_v5 | jioindiacentral, jioindiawest | $1.3600 |
| Standard_D32ls_v6 | jioindiacentral, jioindiawest, westus2 | $1.4280 |
| Standard_D32nds_v6 | eastus2, jioindiacentral, jioindiawest | $2.5460 |
| Standard_D32nlds_v6 | brazilsouth, jioindiacentral, jioindiawest, westcentralus | $3.4440 |
| Standard_D32nls_v6 | jioindiacentral, jioindiawest | $1.8140 |
| Standard_D32ns_v6 | jioindiacentral, jioindiawest | $2.3210 |
| Standard_D32pds_v5 | jioindiacentral, jioindiawest | $0.9660 |
| Standard_D32plds_v5 | eastus2, jioindiacentral, jioindiawest | $1.2290 |
| Standard_D32pls_v5 | jioindiacentral, jioindiawest | $0.6820 |
| Standard_D32ps_v5 | jioindiacentral, jioindiawest, westus2 | $0.8100 |
| Standard_D32ps_v6 | mexicocentral | $1.2350 |
| Standard_D32s_v3 | jioindiacentral, jioindiawest, westus2 | $1.6800 |
| Standard_D32s_v4 | jioindiacentral, jioindiawest | $1.6160 |
| Standard_D32s_v5 | jioindiacentral, jioindiawest, westcentralus | $1.6160 |
| Standard_D32s_v6 | jioindiacentral, jioindiawest, westus2 | $1.6970 |
| Standard_D32s_v7 | westus2 | $2.1170 |
| Standard_D3_v2 | eastus2, jioindiacentral, jioindiawest | $0.2290 |
| Standard_D3_v2_Promo | francecentral | $0.3510 |
| Standard_D4 | mexicocentral | $0.6780 |
| Standard_D48_v3 | jioindiacentral, jioindiawest | $2.5200 |
| Standard_D48_v4 | eastus2, jioindiacentral, jioindiawest | $2.3040 |
| Standard_D48_v5 | jioindiacentral, jioindiawest | $2.4240 |
| Standard_D48a_v4 | jioindiacentral, jioindiawest | $1.4810 |
| Standard_D48ads_v5 | jioindiacentral, jioindiawest | $1.6100 |
| Standard_D48ads_v6 | jioindiacentral, jioindiawest | $1.7710 |
| Standard_D48ads_v7 | jioindiawest, westus2 | $1.7710 |
| Standard_D48alds_v6 | jioindiacentral, jioindiawest, westus2 | $1.4780 |
| Standard_D48alds_v7 | jioindiawest | $1.4780 |
| Standard_D48als_v6 | jioindiacentral, jioindiawest, westcentralus, westus2 | $1.2480 |
| Standard_D48als_v7 | jioindiawest | $1.2480 |
| Standard_D48as_v4 | eastus, jioindiacentral, jioindiawest | $2.3040 |
| Standard_D48as_v5 | jioindiacentral, jioindiawest | $3.5410 |
| Standard_D48as_v6 | jioindiacentral, jioindiawest, westus2 | $1.4060 |
| Standard_D48as_v7 | jioindiawest | $1.4060 |
| Standard_D48d_v4 | jioindiacentral, jioindiawest | $2.9280 |
| Standard_D48d_v5 | jioindiacentral, jioindiawest | $2.9280 |
| Standard_D48ds_v4 | jioindiacentral, jioindiawest, westcentralus | $2.9280 |
| Standard_D48ds_v5 | jioindiacentral, jioindiawest | $2.9280 |
| Standard_D48ds_v6 | jioindiacentral, jioindiawest | $3.1460 |
| Standard_D48lds_v5 | jioindiacentral, jioindiawest, uksouth | $2.3760 |
| Standard_D48lds_v6 | jioindiacentral, jioindiawest, westus2 | $2.6480 |
| Standard_D48lds_v7 | westus2 | $3.3260 |
| Standard_D48ls_v5 | jioindiacentral, jioindiawest | $2.0400 |
| Standard_D48ls_v6 | brazilsouth, jioindiacentral, jioindiawest | $3.3010 |
| Standard_D48nds_v6 | brazilsouth, jioindiacentral, jioindiawest | $7.6380 |
| Standard_D48nlds_v6 | eastus2, jioindiacentral, jioindiawest | $3.3610 |
| Standard_D48nls_v6 | brazilsouth, eastus2, jioindiacentral, jioindiawest | $4.1830 |
| Standard_D48ns_v6 | brazilsouth, jioindiacentral, jioindiawest | $6.6830 |
| Standard_D48pds_v5 | jioindiacentral, jioindiawest | $1.4500 |
| Standard_D48plds_v5 | jioindiacentral, jioindiawest, westus2 | $1.1760 |
| Standard_D48plds_v6 | uksouth, westcentralus | $2.1600 |
| Standard_D48pls_v5 | jioindiacentral, jioindiawest, westus2 | $1.0220 |
| Standard_D48ps_v5 | jioindiacentral, jioindiawest | $1.2140 |
| Standard_D48s_v3 | brazilsouth, eastus2, jioindiacentral, jioindiawest | $3.8160 |
| Standard_D48s_v4 | jioindiacentral, jioindiawest | $2.4240 |
| Standard_D48s_v5 | eastus, eastus2, jioindiacentral, jioindiawest | $2.3040 |
| Standard_D48s_v6 | eastus, jioindiacentral, jioindiawest, westcentralus | $2.4190 |
| Standard_D4_v2 | jioindiacentral, jioindiawest, westcentralus | $0.6750 |
| Standard_D4_v2_Promo | francecentral | $0.7020 |
| Standard_D4_v3 | jioindiacentral, jioindiawest | $0.2100 |
| Standard_D4_v4 | jioindiacentral, jioindiawest, westcentralus, westus2 | $0.2020 |
| Standard_D4_v5 | brazilsoutheast, eastus2, jioindiacentral, jioindiawest | $0.3980 |
| Standard_D4a_v4 | jioindiacentral, jioindiawest, mexicocentral, westus2 | $0.1230 |
| Standard_D4ads_v5 | eastus2, jioindiacentral, jioindiawest | $0.3900 |
| Standard_D4ads_v6 | brazilsouth, jioindiacentral, jioindiawest | $0.3660 |
| Standard_D4ads_v7 | jioindiawest, westus2 | $0.1480 |
| Standard_D4alds_v6 | jioindiacentral, jioindiawest | $0.1230 |
| Standard_D4alds_v7 | eastus2, jioindiawest | $0.1900 |
| Standard_D4als_v6 | jioindiacentral, jioindiawest | $0.1040 |
| Standard_D4als_v7 | jioindiawest | $0.1040 |
| Standard_D4as_v4 | jioindiacentral, jioindiawest | $0.1230 |
| Standard_D4as_v5 | eastus2, jioindiacentral, jioindiawest | $0.3560 |
| Standard_D4as_v6 | jioindiacentral, jioindiawest, mexicocentral | $0.1170 |
| Standard_D4as_v7 | jioindiawest | $0.1170 |
| Standard_D4d_v4 | jioindiacentral, jioindiawest | $0.2440 |
| Standard_D4d_v5 | jioindiacentral, jioindiawest | $0.2440 |
| Standard_D4ds_v4 | jioindiacentral, jioindiawest | $0.2440 |
| Standard_D4ds_v5 | eastus, jioindiacentral, jioindiawest | $0.2260 |
| Standard_D4ds_v6 | brazilsouth, jioindiacentral, jioindiawest, westus2 | $0.3970 |
| Standard_D4lds_v5 | jioindiacentral, jioindiawest | $0.1980 |
| Standard_D4lds_v6 | jioindiacentral, jioindiawest | $0.2210 |
| Standard_D4ls_v5 | jioindiacentral, jioindiawest | $0.1700 |
| Standard_D4ls_v6 | eastus2, jioindiacentral, jioindiawest, mexicocentral | $0.1790 |
| Standard_D4ls_v7 | westus2 | $0.2340 |
| Standard_D4nds_v6 | jioindiacentral, jioindiawest, northcentralus | $0.3120 |
| Standard_D4nlds_v6 | jioindiacentral, jioindiawest, westus2 | $0.2800 |
| Standard_D4nls_v6 | eastus, jioindiacentral, jioindiawest | $0.2270 |
| Standard_D4ns_v6 | jioindiacentral, jioindiawest | $0.2900 |
| Standard_D4pds_v5 | jioindiacentral, jioindiawest, westus2 | $0.1210 |
| Standard_D4plds_v5 | jioindiacentral, jioindiawest | $0.0980 |
| Standard_D4plds_v6 | eastus2 | $0.1560 |
| Standard_D4pls_v5 | jioindiacentral, jioindiawest | $0.0852 |
| Standard_D4pls_v6 | brazilsouth | $0.1970 |
| Standard_D4ps_v5 | jioindiacentral, jioindiawest, westus2 | $0.1010 |
| Standard_D4ps_v6 | westus2 | $0.1400 |
| Standard_D4s_v3 | jioindiacentral, jioindiawest | $0.2100 |
| Standard_D4s_v4 | jioindiacentral, jioindiawest | $0.2020 |
| Standard_D4s_v5 | canadacentral, jioindiacentral, jioindiawest | $0.2140 |
| Standard_D4s_v6 | jioindiacentral, jioindiawest | $0.2120 |
| Standard_D4s_v7 | brazilsouth | $0.4230 |
| Standard_D5_v2 | jioindiacentral, jioindiawest | $1.3500 |
| Standard_D5_v2_Promo | francecentral | $1.4040 |
| Standard_D64-16s_v3 | jioindiacentral, jioindiawest | $3.3600 |
| Standard_D64-32s_v3 | jioindiacentral, jioindiawest | $3.3600 |
| Standard_D64_v3 | jioindiacentral, jioindiawest, westus2 | $3.3600 |
| Standard_D64_v4 | jioindiacentral, jioindiawest | $3.2320 |
| Standard_D64_v5 | brazilsoutheast, jioindiacentral, jioindiawest, westcentralus | $6.3650 |
| Standard_D64a_v4 | jioindiacentral, jioindiawest | $1.9750 |
| Standard_D64ads_v5 | jioindiacentral, jioindiawest, westus2 | $5.0910 |
| Standard_D64ads_v6 | jioindiacentral, jioindiawest | $2.3620 |
| Standard_D64ads_v7 | jioindiawest | $2.3620 |
| Standard_D64alds_v6 | jioindiacentral, jioindiawest | $1.9710 |
| Standard_D64alds_v7 | jioindiawest | $1.9710 |
| Standard_D64als_v6 | jioindiacentral, jioindiawest, uksouth | $1.6640 |
| Standard_D64als_v7 | jioindiawest | $1.6640 |
| Standard_D64as_v4 | jioindiacentral, jioindiawest | $1.9750 |
| Standard_D64as_v5 | brazilsoutheast, jioindiacentral, jioindiawest, westcentralus, westus2 | $8.6850 |
| Standard_D64as_v6 | brazilsoutheast, jioindiacentral, jioindiawest | $6.0420 |
| Standard_D64as_v7 | jioindiawest, mexicocentral, westus2 | $1.8750 |
| Standard_D64d_v4 | jioindiacentral, jioindiawest | $3.9040 |
| Standard_D64d_v5 | jioindiacentral, jioindiawest | $3.9040 |
| Standard_D64ds_v4 | jioindiacentral, jioindiawest | $3.9040 |
| Standard_D64ds_v5 | jioindiacentral, jioindiawest, westus2 | $3.9040 |
| Standard_D64ds_v6 | jioindiacentral, jioindiawest, uksouth, westus2 | $4.1950 |
| Standard_D64lds_v5 | brazilsouth, jioindiacentral, jioindiawest | $4.7680 |
| Standard_D64lds_v6 | brazilsouth, jioindiacentral, jioindiawest | $5.4400 |
| Standard_D64ls_v5 | jioindiacentral, jioindiawest, westus2 | $2.7200 |
| Standard_D64ls_v6 | jioindiacentral, jioindiawest | $2.8560 |
| Standard_D64nds_v6 | eastus, eastus2, jioindiacentral, jioindiawest | $5.0920 |
| Standard_D64nlds_v6 | eastus, jioindiacentral, jioindiawest, mexicocentral | $4.4820 |
| Standard_D64nls_v6 | jioindiacentral, jioindiawest | $3.6290 |
| Standard_D64ns_v6 | jioindiacentral, jioindiawest | $4.6430 |
| Standard_D64pds_v5 | jioindiacentral, jioindiawest | $1.9330 |
| Standard_D64pds_v6 | westus2 | $2.9380 |
| Standard_D64plds_v5 | eastus2, jioindiacentral, jioindiawest, westus2 | $2.4580 |
| Standard_D64pls_v5 | jioindiacentral, jioindiawest | $1.3630 |
| Standard_D64pls_v6 | eastus2 | $1.9840 |
| Standard_D64ps_v5 | jioindiacentral, jioindiawest, westus2 | $1.6190 |
| Standard_D64s_v3 | eastus2, jioindiacentral, jioindiawest | $3.0720 |
| Standard_D64s_v4 | jioindiacentral, jioindiawest | $3.2320 |
| Standard_D64s_v5 | jioindiacentral, jioindiawest | $3.2320 |
| Standard_D64s_v6 | jioindiacentral, jioindiawest, westus2 | $3.3940 |
| Standard_D64s_v7 | brazilsouth, westus2 | $6.7740 |
| Standard_D8_v3 | jioindiacentral, jioindiawest | $0.4200 |
| Standard_D8_v4 | jioindiacentral, jioindiawest | $0.4040 |
| Standard_D8_v5 | eastus, eastus2, jioindiacentral, jioindiawest, westus2 | $0.3840 |
| Standard_D8a_v4 | eastus2, jioindiacentral, jioindiawest | $0.3840 |
| Standard_D8ads_v5 | jioindiacentral, jioindiawest, westus2 | $0.6360 |
| Standard_D8ads_v6 | jioindiacentral, jioindiawest, mexicocentral | $0.2950 |
| Standard_D8ads_v7 | jioindiawest | $0.2950 |
| Standard_D8alds_v6 | jioindiacentral, jioindiawest | $0.2460 |
| Standard_D8alds_v7 | jioindiawest | $0.2460 |
| Standard_D8als_v6 | jioindiacentral, jioindiawest | $0.2080 |
| Standard_D8als_v7 | eastus2, jioindiawest | $0.3220 |
| Standard_D8as_v4 | brazilsouth, eastus, jioindiacentral, jioindiawest | $0.6120 |
| Standard_D8as_v5 | jioindiacentral, jioindiawest | $0.2220 |
| Standard_D8as_v5_Promo | eastus2 | $0.3080 |
| Standard_D8as_v6 | eastus2, jioindiacentral, jioindiawest, westus2 | $0.3630 |
| Standard_D8as_v7 | jioindiawest | $0.2340 |
| Standard_D8d_v4 | jioindiacentral, jioindiawest | $0.4880 |
| Standard_D8d_v5 | jioindiacentral, jioindiawest | $0.4880 |
| Standard_D8ds_v4 | jioindiacentral, jioindiawest, westcentralus | $0.4880 |
| Standard_D8ds_v5 | jioindiacentral, jioindiawest | $0.4880 |
| Standard_D8ds_v6 | eastus2, jioindiacentral, jioindiawest, westus2 | $0.4980 |
| Standard_D8lds_v5 | jioindiacentral, jioindiawest, uksouth | $0.3960 |
| Standard_D8lds_v6 | jioindiacentral, jioindiawest | $0.4410 |
| Standard_D8ls_v5 | eastus2, jioindiacentral, jioindiawest | $0.3400 |
| Standard_D8ls_v6 | jioindiacentral, jioindiawest | $0.3570 |
| Standard_D8nds_v6 | jioindiacentral, jioindiawest | $0.6790 |
| Standard_D8nlds_v6 | jioindiacentral, jioindiawest, westcentralus, westus2 | $0.5600 |
| Standard_D8nls_v6 | eastus, jioindiacentral, jioindiawest | $0.4540 |
| Standard_D8ns_v6 | brazilsouth, eastus2, jioindiacentral, jioindiawest | $1.1140 |
| Standard_D8pds_v5 | jioindiacentral, jioindiawest | $0.2420 |
| Standard_D8pds_v6 | westcentralus | $0.4410 |
| Standard_D8plds_v5 | canadacentral, jioindiacentral, jioindiawest, westus2 | $0.3390 |
| Standard_D8plds_v6 | eastus2 | $0.3120 |
| Standard_D8pls_v5 | jioindiacentral, jioindiawest | $0.1700 |
| Standard_D8pls_v6 | mexicocentral | $0.2730 |
| Standard_D8ps_v5 | jioindiacentral, jioindiawest | $0.2020 |
| Standard_D8s_v3 | jioindiacentral, jioindiawest | $0.4200 |
| Standard_D8s_v4 | jioindiacentral, jioindiawest | $0.4040 |
| Standard_D8s_v5 | jioindiacentral, jioindiawest | $0.4040 |
| Standard_D8s_v6 | eastus2, jioindiacentral, jioindiawest | $0.4030 |
| Standard_D8s_v7 | eastus2 | $0.5290 |
| Standard_D96_v5 | jioindiacentral, jioindiawest | $4.8480 |
| Standard_D96a_v4 | jioindiacentral, jioindiawest, westcentralus, westus2 | $2.9620 |
| Standard_D96ads_v5 | jioindiacentral, jioindiawest, westus2 | $3.2210 |
| Standard_D96ads_v6 | jioindiacentral, jioindiawest | $3.5420 |
| Standard_D96ads_v7 | jioindiawest | $3.5420 |
| Standard_D96alds_v6 | jioindiacentral, jioindiawest | $2.9570 |
| Standard_D96alds_v7 | jioindiawest | $2.9570 |
| Standard_D96als_v6 | jioindiacentral, jioindiawest, westus2 | $2.4960 |
| Standard_D96als_v7 | jioindiawest, westus2 | $2.4960 |
| Standard_D96as_v4 | jioindiacentral, jioindiawest | $2.9620 |
| Standard_D96as_v5 | brazilsouth, brazilsoutheast, jioindiacentral, jioindiawest | $6.6240 |
| Standard_D96as_v6 | brazilsouth, jioindiacentral, jioindiawest | $6.9890 |
| Standard_D96as_v7 | eastus2, jioindiawest | $4.3580 |
| Standard_D96d_v5 | jioindiacentral, jioindiawest | $5.8560 |
| Standard_D96ds_v5 | jioindiacentral, jioindiawest | $5.8560 |
| Standard_D96ds_v6 | jioindiacentral, jioindiawest, westus2 | $6.2920 |
| Standard_D96lds_v5 | jioindiacentral, jioindiawest | $4.7520 |
| Standard_D96lds_v6 | jioindiacentral, jioindiawest, uksouth, westus2 | $5.2950 |
| Standard_D96ls_v5 | jioindiacentral, jioindiawest, westus2 | $4.0800 |
| Standard_D96ls_v6 | jioindiacentral, jioindiawest | $4.2840 |
| Standard_D96nds_v6 | jioindiacentral, jioindiawest | $8.1430 |
| Standard_D96nlds_v6 | jioindiacentral, jioindiawest, westus2 | $6.7220 |
| Standard_D96nls_v6 | jioindiacentral, jioindiawest, westus2 | $5.4430 |
| Standard_D96ns_v6 | jioindiacentral, jioindiawest | $6.9640 |
| Standard_D96ps_v6 | eastus2 | $3.3700 |
| Standard_D96s_v5 | jioindiacentral, jioindiawest | $4.8480 |
| Standard_D96s_v6 | jioindiacentral, jioindiawest | $5.0910 |
| Standard_D96vds_v7 | eastus | $7.8320 |
| Standard_DC16ads_cc_v5 | jioindiacentral, jioindiawest | $0.5370 |
| Standard_DC16ads_v5 | jioindiacentral, jioindiawest | $0.5370 |
| Standard_DC16as_cc_v5 | eastus2, jioindiacentral, jioindiawest | $0.6880 |
| Standard_DC16as_v5 | jioindiacentral, jioindiawest | $0.4440 |
| Standard_DC16as_v6 | brazilsoutheast | $1.6610 |
| Standard_DC16ds_v3 | jioindiacentral, jioindiawest, mexicocentral | $1.9200 |
| Standard_DC16s_v3 | jioindiacentral, jioindiawest | $1.6160 |
| Standard_DC1ds_v3 | jioindiacentral, jioindiawest | $0.1200 |
| Standard_DC1s_v3 | jioindiacentral, jioindiawest | $0.1010 |
| Standard_DC24ds_v3 | eastus2, jioindiacentral, jioindiawest | $2.7120 |
| Standard_DC24s_v3 | jioindiacentral, jioindiawest, westus2 | $2.4240 |
| Standard_DC2ads_v5 | jioindiacentral, jioindiawest | $0.0671 |
| Standard_DC2ads_v6 | eastus2, northcentralus | $0.1250 |
| Standard_DC2as_v5 | eastus, eastus2, jioindiacentral, jioindiawest | $0.0860 |
| Standard_DC2as_v6 | eastus2, westus2 | $0.0999 |
| Standard_DC2ds_v3 | jioindiacentral, jioindiawest | $0.2400 |
| Standard_DC2es_v6 | westus3 | $0.1110 |
| Standard_DC2s_v3 | jioindiacentral, jioindiawest, westus2 | $0.2020 |
| Standard_DC32ads_cc_v5 | brazilsouth, jioindiacentral, jioindiawest | $2.6400 |
| Standard_DC32ads_v5 | eastus, jioindiacentral, jioindiawest | $1.6480 |
| Standard_DC32as_cc_v5 | eastus2, jioindiacentral, jioindiawest, westus2 | $1.3760 |
| Standard_DC32as_v5 | jioindiacentral, jioindiawest | $0.8890 |
| Standard_DC32ds_v3 | jioindiacentral, jioindiawest | $3.8400 |
| Standard_DC32s_v3 | jioindiacentral, jioindiawest | $3.2320 |
| Standard_DC48ads_cc_v5 | jioindiacentral, jioindiawest | $1.6100 |
| Standard_DC48ads_v5 | jioindiacentral, jioindiawest | $1.6100 |
| Standard_DC48as_cc_v5 | jioindiacentral, jioindiawest, westus2 | $1.3330 |
| Standard_DC48as_v5 | eastus, jioindiacentral, jioindiawest | $2.0640 |
| Standard_DC48ds_v3 | jioindiacentral, jioindiawest | $5.7600 |
| Standard_DC48s_v3 | jioindiacentral, jioindiawest, westus2 | $4.8480 |
| Standard_DC4ads_cc_v5 | jioindiacentral, jioindiawest | $0.1340 |
| Standard_DC4ads_v5 | jioindiacentral, jioindiawest | $0.1340 |
| Standard_DC4ads_v6 | brazilsouth, eastus2 | $0.4030 |
| Standard_DC4as_cc_v5 | jioindiacentral, jioindiawest, uksouth, westus2 | $0.1110 |
| Standard_DC4as_v5 | jioindiacentral, jioindiawest | $0.1110 |
| Standard_DC4as_v6 | eastus | $0.2000 |
| Standard_DC4ds_v3 | eastus2, jioindiacentral, jioindiawest | $0.4520 |
| Standard_DC4s_v2 | southcentralus | $0.4610 |
| Standard_DC4s_v3 | jioindiacentral, jioindiawest | $0.4040 |
| Standard_DC64ads_cc_v5 | jioindiacentral, jioindiawest | $2.1470 |
| Standard_DC64ads_v5 | jioindiacentral, jioindiawest | $2.1470 |
| Standard_DC64as_cc_v5 | jioindiacentral, jioindiawest | $1.7780 |
| Standard_DC64as_v5 | jioindiacentral, jioindiawest | $1.7780 |
| Standard_DC64eds_v5 | eastus2 | $3.6160 |
| Standard_DC8_v2 | eastus | $0.7680 |
| Standard_DC8ads_cc_v5 | jioindiacentral, jioindiawest | $0.2680 |
| Standard_DC8ads_v5 | jioindiacentral, jioindiawest | $0.2680 |
| Standard_DC8as_cc_v5 | jioindiacentral, jioindiawest | $0.2220 |
| Standard_DC8as_v5 | jioindiacentral, jioindiawest | $0.2220 |
| Standard_DC8ds_v3 | jioindiacentral, jioindiawest | $0.9600 |
| Standard_DC8s_v3 | jioindiacentral, jioindiawest | $0.8080 |
| Standard_DC96ads_cc_v5 | brazilsouth, jioindiacentral, jioindiawest, westus2 | $7.9200 |
| Standard_DC96ads_v5 | jioindiacentral, jioindiawest | $3.2210 |
| Standard_DC96ads_v6 | brazilsouth | $9.6620 |
| Standard_DC96as_cc_v5 | jioindiacentral, jioindiawest | $2.6660 |
| Standard_DC96as_v5 | jioindiacentral, jioindiawest | $2.6660 |
| Standard_DS1 | brazilsouth | $0.0950 |
| Standard_DS11-1_v2 | jioindiacentral, jioindiawest | $0.1890 |
| Standard_DS11_v2 | jioindiacentral, jioindiawest | $0.1890 |
| Standard_DS11_v2_Promo | francecentral, southafricawest | $0.2340 |
| Standard_DS12-1_v2 | jioindiacentral, jioindiawest | $0.3790 |
| Standard_DS12-2_v2 | jioindiacentral, jioindiawest | $0.3790 |
| Standard_DS12_v2 | jioindiacentral, jioindiawest | $0.3790 |
| Standard_DS12_v2_Promo | francecentral, southafricanorth | $0.4690 |
| Standard_DS13-2_v2 | jioindiacentral, jioindiawest | $0.7580 |
| Standard_DS13-4_v2 | brazilsouth, jioindiacentral, jioindiawest | $0.9400 |
| Standard_DS13_v2 | jioindiacentral, jioindiawest | $0.7580 |
| Standard_DS13_v2_Promo | francecentral, southafricanorth, southafricawest | $0.9370 |
| Standard_DS14-4_v2 | eastus2, jioindiacentral, jioindiawest | $1.1960 |
| Standard_DS14-8_v2 | eastus2, jioindiacentral, jioindiawest, westcentralus | $1.1960 |
| Standard_DS14_v2 | eastus2, jioindiacentral, jioindiawest | $1.1960 |
| Standard_DS14_v2_Promo | francecentral, southafricawest | $1.8740 |
| Standard_DS15_v2 | jioindiacentral, jioindiawest, westus2 | $1.8950 |
| Standard_DS15i_v2 | jioindiacentral, jioindiawest | $1.8950 |
| Standard_DS1_v2 | eastus2, jioindiacentral, jioindiawest | $0.0570 |
| Standard_DS2_v2 | brazilsouth, jioindiacentral, jioindiawest, northcentralus, westcentralus | $0.1710 |
| Standard_DS2_v2_Promo | francecentral | $0.1750 |
| Standard_DS3_v2 | jioindiacentral, jioindiawest | $0.3370 |
| Standard_DS3_v2_Promo | francecentral | $0.3510 |
| Standard_DS4_v2 | jioindiacentral, jioindiawest | $0.6750 |
| Standard_DS4_v2_Promo | francecentral, southafricanorth | $0.7020 |
| Standard_DS5_v2 | jioindiacentral, jioindiawest | $1.3500 |
| Standard_DS5_v2_Promo | francecentral, southafricanorth, southafricawest | $1.4040 |
| Standard_E104i_v5 | westus2 | $7.2070 |
| Standard_E104id_v5 | eastus | $8.2370 |
| Standard_E112iads_v5 | eastus2, jioindiacentral, jioindiawest, westcentralus | $8.0700 |
| Standard_E112ias_v5 | eastus2, jioindiacentral, jioindiawest, westus2 | $6.9610 |
| Standard_E112ibds_v5 | jioindiacentral, jioindiawest | $10.8420 |
| Standard_E112ibs_v5 | eastus2, jioindiacentral, jioindiawest | $9.1780 |
| Standard_E128-32ads_v7 | jioindiawest | $6.0030 |
| Standard_E128-32as_v7 | jioindiawest | $4.9280 |
| Standard_E128-32ds_v6 | jioindiacentral, jioindiawest | $10.7980 |
| Standard_E128-32s_v6 | jioindiacentral, jioindiawest | $8.7360 |
| Standard_E128-64ads_v7 | jioindiawest | $6.0030 |
| Standard_E128-64as_v7 | jioindiawest | $4.9280 |
| Standard_E128-64ds_v6 | jioindiacentral, jioindiawest | $10.7980 |
| Standard_E128-64s_v6 | jioindiacentral, jioindiawest, westus2 | $8.7360 |
| Standard_E128ads_v7 | jioindiawest | $6.0030 |
| Standard_E128as_v7 | jioindiawest | $4.9280 |
| Standard_E128bds_v6 | brazilsoutheast, eastus2, jioindiacentral, jioindiawest | $32.5150 |
| Standard_E128bs_v6 | eastus2, jioindiacentral, jioindiawest | $11.1620 |
| Standard_E128ds_v6 | jioindiacentral, jioindiawest, northcentralus, westus2 | $10.7980 |
| Standard_E128ds_v7 | westus2 | $13.3060 |
| Standard_E128nds_v6 | jioindiacentral, jioindiawest | $17.5070 |
| Standard_E128ns_v6 | jioindiacentral, jioindiawest, westus2 | $15.6200 |
| Standard_E128s_v6 | jioindiacentral, jioindiawest, westus2 | $8.7360 |
| Standard_E128s_v7 | mexicocentral | $12.2700 |
| Standard_E16-4ads_v5 | jioindiacentral, jioindiawest | $1.4000 |
| Standard_E16-4ads_v6 | jioindiacentral, jioindiawest | $0.7500 |
| Standard_E16-4ads_v7 | jioindiawest | $0.7500 |
| Standard_E16-4as_v4 | jioindiacentral, jioindiawest, westcentralus | $1.0400 |
| Standard_E16-4as_v5 | jioindiacentral, jioindiawest | $0.5720 |
| Standard_E16-4as_v6 | jioindiacentral, jioindiawest | $0.6160 |
| Standard_E16-4as_v7 | jioindiawest, uksouth | $0.6160 |
| Standard_E16-4ds_v4 | jioindiacentral, jioindiawest | $1.2080 |
| Standard_E16-4ds_v5 | canadacentral, jioindiacentral, jioindiawest, southcentralus, westcentralus (+1 more) | $1.2640 |
| Standard_E16-4ds_v6 | jioindiacentral, jioindiawest, westus2 | $1.3500 |
| Standard_E16-4s_v3 | brazilsouth, jioindiacentral, jioindiawest | $1.8790 |
| Standard_E16-4s_v4 | eastus2, jioindiacentral, jioindiawest, westus2 | $1.0080 |
| Standard_E16-4s_v5 | jioindiacentral, jioindiawest | $1.0400 |
| Standard_E16-4s_v6 | brazilsoutheast, eastus2, jioindiacentral, jioindiawest | $2.1950 |
| Standard_E16-8ads_v5 | jioindiacentral, jioindiawest | $1.4000 |
| Standard_E16-8ads_v6 | jioindiacentral, jioindiawest | $0.7500 |
| Standard_E16-8ads_v7 | jioindiawest | $0.7500 |
| Standard_E16-8as_v4 | jioindiacentral, jioindiawest | $1.0400 |
| Standard_E16-8as_v5 | eastus, jioindiacentral, jioindiawest | $0.9040 |
| Standard_E16-8as_v6 | eastus2, jioindiacentral, jioindiawest | $0.9540 |
| Standard_E16-8as_v7 | jioindiawest, uksouth | $0.6160 |
| Standard_E16-8ds_v4 | jioindiacentral, jioindiawest | $1.2080 |
| Standard_E16-8ds_v5 | jioindiacentral, jioindiawest | $1.2080 |
| Standard_E16-8ds_v6 | jioindiacentral, jioindiawest | $1.3500 |
| Standard_E16-8s_v3 | jioindiacentral, jioindiawest, mexicocentral | $1.0960 |
| Standard_E16-8s_v4 | brazilsoutheast, jioindiacentral, jioindiawest | $2.0900 |
| Standard_E16-8s_v5 | jioindiacentral, jioindiawest, westus2 | $1.0400 |
| Standard_E16-8s_v6 | brazilsoutheast, jioindiacentral, jioindiawest, northcentralus | $2.1950 |
| Standard_E160ads_v7 | jioindiawest | $7.5040 |
| Standard_E160as_v7 | jioindiawest | $6.1600 |
| Standard_E16_v3 | jioindiacentral, jioindiawest, westus2 | $1.0960 |
| Standard_E16_v4 | eastus2, jioindiacentral, jioindiawest | $1.0080 |
| Standard_E16_v5 | eastus, jioindiacentral, jioindiawest, uksouth | $1.0080 |
| Standard_E16a_v4 | jioindiacentral, jioindiawest | $0.6350 |
| Standard_E16ads_v5 | jioindiacentral, jioindiawest | $0.6640 |
| Standard_E16ads_v6 | jioindiacentral, jioindiawest, mexicocentral | $0.7500 |
| Standard_E16ads_v7 | brazilsouth, eastus2, jioindiawest | $1.8620 |
| Standard_E16as_v4 | brazilsoutheast, jioindiacentral, jioindiawest, westcentralus | $2.0900 |
| Standard_E16as_v5 | jioindiacentral, jioindiawest | $0.5720 |
| Standard_E16as_v6 | eastus, jioindiacentral, jioindiawest, westus2 | $0.9540 |
| Standard_E16as_v7 | jioindiawest | $0.6160 |
| Standard_E16bds_v5 | brazilsoutheast, jioindiacentral, jioindiawest | $2.7720 |
| Standard_E16bds_v6 | brazilsoutheast, eastus2, jioindiacentral, jioindiawest | $4.0640 |
| Standard_E16bs_v5 | canadacentral, jioindiacentral, jioindiawest | $1.3160 |
| Standard_E16bs_v6 | eastus, jioindiacentral, jioindiawest, westcentralus | $1.3950 |
| Standard_E16d_v4 | brazilsouth, jioindiacentral, jioindiawest | $1.8240 |
| Standard_E16d_v5 | jioindiacentral, jioindiawest | $1.2080 |
| Standard_E16ds_v4 | brazilsoutheast, eastus2, jioindiacentral, jioindiawest | $2.3890 |
| Standard_E16ds_v4_ADHType1 | jioindiacentral, jioindiawest | $1.2080 |
| Standard_E16ds_v5 | jioindiacentral, jioindiawest | $1.2080 |
| Standard_E16ds_v6 | jioindiacentral, jioindiawest | $1.3500 |
| Standard_E16ds_v7 | brazilsouth | $2.6610 |
| Standard_E16nds_v6 | brazilsouth, jioindiacentral, jioindiawest | $3.1260 |
| Standard_E16ns_v6 | jioindiacentral, jioindiawest | $1.9520 |
| Standard_E16pds_v5 | jioindiacentral, jioindiawest | $0.6560 |
| Standard_E16ps_v5 | jioindiacentral, jioindiawest | $0.5200 |
| Standard_E16ps_v6 | canadacentral, eastus2 | $0.8190 |
| Standard_E16s_v3 | jioindiacentral, jioindiawest | $1.0960 |
| Standard_E16s_v4 | jioindiacentral, jioindiawest | $1.0400 |
| Standard_E16s_v5 | jioindiacentral, jioindiawest | $1.0400 |
| Standard_E16s_v6 | brazilsoutheast, jioindiacentral, jioindiawest | $2.1950 |
| Standard_E192ds_v7 | brazilsouth | $31.9330 |
| Standard_E192ibds_v6 | jioindiacentral, jioindiawest, northcentralus | $28.8880 |
| Standard_E192ibs_v6 | canadacentral, jioindiacentral, jioindiawest | $22.0910 |
| Standard_E192ids_v6 | jioindiacentral, jioindiawest | $17.8170 |
| Standard_E192is_v6 | jioindiacentral, jioindiawest | $14.4140 |
| Standard_E192s_v7 | uksouth, westus2 | $20.1120 |
| Standard_E20_v3 | jioindiacentral, jioindiawest | $1.3700 |
| Standard_E20_v4 | jioindiacentral, jioindiawest | $1.3000 |
| Standard_E20_v5 | jioindiacentral, jioindiawest | $1.3000 |
| Standard_E20a_v4 | jioindiacentral, jioindiawest | $0.7940 |
| Standard_E20ads_v5 | jioindiacentral, jioindiawest, westcentralus | $0.8300 |
| Standard_E20ads_v6 | brazilsouth, jioindiacentral, jioindiawest | $2.3280 |
| Standard_E20as_v4 | jioindiacentral, jioindiawest | $0.7940 |
| Standard_E20as_v5 | eastus, jioindiacentral, jioindiawest | $1.1300 |
| Standard_E20as_v6 | eastus, jioindiacentral, jioindiawest | $1.1920 |
| Standard_E20d_v4 | jioindiacentral, jioindiawest | $1.5100 |
| Standard_E20d_v5 | jioindiacentral, jioindiawest | $1.5100 |
| Standard_E20ds_v4 | jioindiacentral, jioindiawest, northcentralus, uksouth, westcentralus | $1.5100 |
| Standard_E20ds_v5 | brazilsoutheast, jioindiacentral, jioindiawest | $2.9640 |
| Standard_E20ds_v6 | jioindiacentral, jioindiawest | $1.6870 |
| Standard_E20pds_v5 | jioindiacentral, jioindiawest | $0.8200 |
| Standard_E20ps_v5 | jioindiacentral, jioindiawest, westus2 | $0.6500 |
| Standard_E20s_v3 | jioindiacentral, jioindiawest | $1.3700 |
| Standard_E20s_v4 | jioindiacentral, jioindiawest, westus2 | $1.3000 |
| Standard_E20s_v5 | eastus2, jioindiacentral, jioindiawest | $1.2600 |
| Standard_E20s_v6 | jioindiacentral, jioindiawest, mexicocentral | $1.3650 |
| Standard_E2_v3 | jioindiacentral, jioindiawest, westus2 | $0.1370 |
| Standard_E2_v4 | eastus2, jioindiacentral, jioindiawest | $0.1260 |
| Standard_E2_v5 | jioindiacentral, jioindiawest | $0.1300 |
| Standard_E2a_v4 | brazilsouth, jioindiacentral, jioindiawest | $0.2010 |
| Standard_E2ads_v5 | jioindiacentral, jioindiawest | $0.0830 |
| Standard_E2ads_v6 | eastus, jioindiacentral, jioindiawest | $0.1450 |
| Standard_E2ads_v7 | jioindiawest | $0.0938 |
| Standard_E2as_v4 | jioindiacentral, jioindiawest, westus2 | $0.0800 |
| Standard_E2as_v5 | brazilsoutheast, jioindiacentral, jioindiawest | $0.2350 |
| Standard_E2as_v6 | jioindiacentral, jioindiawest | $0.0770 |
| Standard_E2as_v7 | jioindiawest | $0.0770 |
| Standard_E2bds_v5 | jioindiacentral, jioindiawest | $0.1760 |
| Standard_E2bds_v6 | jioindiacentral, jioindiawest | $0.2740 |
| Standard_E2bs_v5 | jioindiacentral, jioindiawest | $0.1550 |
| Standard_E2bs_v6 | brazilsoutheast, jioindiacentral, jioindiawest, westus2 | $0.4530 |
| Standard_E2d_v4 | brazilsouth, jioindiacentral, jioindiawest | $0.2280 |
| Standard_E2d_v5 | eastus2, jioindiacentral, jioindiawest | $0.1440 |
| Standard_E2ds_v4 | eastus2, jioindiacentral, jioindiawest | $0.1440 |
| Standard_E2ds_v5 | jioindiacentral, jioindiawest | $0.1510 |
| Standard_E2ds_v6 | jioindiacentral, jioindiawest | $0.1690 |
| Standard_E2nds_v6 | brazilsoutheast, jioindiacentral, jioindiawest | $0.5080 |
| Standard_E2ns_v6 | jioindiacentral, jioindiawest | $0.2440 |
| Standard_E2pds_v5 | jioindiacentral, jioindiawest, westus2 | $0.0820 |
| Standard_E2pds_v6 | eastus2 | $0.1170 |
| Standard_E2ps_v5 | jioindiacentral, jioindiawest | $0.0650 |
| Standard_E2ps_v6 | brazilsouth | $0.1470 |
| Standard_E2s_v3 | eastus, jioindiacentral, jioindiawest | $0.1260 |
| Standard_E2s_v4 | jioindiacentral, jioindiawest | $0.1300 |
| Standard_E2s_v5 | jioindiacentral, jioindiawest | $0.1300 |
| Standard_E2s_v6 | jioindiacentral, jioindiawest | $0.1370 |
| Standard_E32-16ads_v5 | jioindiacentral, jioindiawest, westcentralus | $2.8010 |
| Standard_E32-16ads_v6 | jioindiacentral, jioindiawest | $1.5010 |
| Standard_E32-16ads_v7 | jioindiawest | $1.5010 |
| Standard_E32-16as_v4 | eastus2, jioindiacentral, jioindiawest | $2.0160 |
| Standard_E32-16as_v5 | brazilsoutheast, jioindiacentral, jioindiawest | $5.2370 |
| Standard_E32-16as_v6 | jioindiacentral, jioindiawest, mexicocentral | $1.2320 |
| Standard_E32-16as_v7 | jioindiawest, westus2 | $1.2320 |
| Standard_E32-16ds_v4 | eastus, jioindiacentral, jioindiawest, westus2 | $2.3040 |
| Standard_E32-16ds_v5 | jioindiacentral, jioindiawest | $2.4160 |
| Standard_E32-16ds_v6 | canadacentral, jioindiacentral, jioindiawest | $2.8660 |
| Standard_E32-16s_v3 | brazilsouth, jioindiacentral, jioindiawest | $3.7580 |
| Standard_E32-16s_v4 | jioindiacentral, jioindiawest | $2.0800 |
| Standard_E32-16s_v5 | jioindiacentral, jioindiawest | $2.0800 |
| Standard_E32-16s_v6 | jioindiacentral, jioindiawest | $2.1840 |
| Standard_E32-8ads_v5 | brazilsouth, jioindiacentral, jioindiawest, mexicocentral | $4.8000 |
| Standard_E32-8ads_v6 | jioindiacentral, jioindiawest | $1.5010 |
| Standard_E32-8ads_v7 | jioindiawest | $1.5010 |
| Standard_E32-8as_v4 | jioindiacentral, jioindiawest | $2.0800 |
| Standard_E32-8as_v5 | jioindiacentral, jioindiawest, westus2 | $2.6160 |
| Standard_E32-8as_v6 | jioindiacentral, jioindiawest | $1.2320 |
| Standard_E32-8as_v7 | jioindiawest | $1.2320 |
| Standard_E32-8ds_v4 | jioindiacentral, jioindiawest | $2.4160 |
| Standard_E32-8ds_v5 | jioindiacentral, jioindiawest, westus2 | $2.4160 |
| Standard_E32-8ds_v6 | eastus2, jioindiacentral, jioindiawest, mexicocentral | $2.6160 |
| Standard_E32-8s_v3 | jioindiacentral, jioindiawest, westus2 | $2.1920 |
| Standard_E32-8s_v4 | jioindiacentral, jioindiawest, westcentralus, westus2 | $2.0800 |
| Standard_E32-8s_v5 | jioindiacentral, jioindiawest, mexicocentral | $2.0800 |
| Standard_E32-8s_v6 | jioindiacentral, jioindiawest | $2.1840 |
| Standard_E32_v3 | jioindiacentral, jioindiawest, westus2 | $2.1920 |
| Standard_E32_v4 | jioindiacentral, jioindiawest, mexicocentral, westcentralus, westus2 | $2.0800 |
| Standard_E32_v5 | jioindiacentral, jioindiawest | $2.0800 |
| Standard_E32a_v4 | jioindiacentral, jioindiawest | $1.2710 |
| Standard_E32ads_v5 | brazilsoutheast, jioindiacentral, jioindiawest, westus2 | $5.7980 |
| Standard_E32ads_v6 | jioindiacentral, jioindiawest | $1.5010 |
| Standard_E32ads_v7 | brazilsouth, jioindiawest | $3.7250 |
| Standard_E32as_v4 | eastus2, jioindiacentral, jioindiawest | $2.0160 |
| Standard_E32as_v5 | jioindiacentral, jioindiawest | $1.1440 |
| Standard_E32as_v6 | jioindiacentral, jioindiawest, southcentralus, westus2 | $1.2320 |
| Standard_E32as_v7 | jioindiawest, westus2 | $1.2320 |
| Standard_E32bds_v5 | jioindiacentral, jioindiawest | $2.8160 |
| Standard_E32bds_v6 | canadacentral, jioindiacentral, jioindiawest, northcentralus | $3.7510 |
| Standard_E32bs_v5 | jioindiacentral, jioindiawest, mexicocentral, westus2 | $2.4800 |
| Standard_E32bs_v6 | jioindiacentral, jioindiawest | $3.9070 |
| Standard_E32d_v4 | brazilsouth, jioindiacentral, jioindiawest | $3.6480 |
| Standard_E32d_v5 | jioindiacentral, jioindiawest, westcentralus | $2.4160 |
| Standard_E32ds_v4 | jioindiacentral, jioindiawest, westcentralus | $2.4160 |
| Standard_E32ds_v4_ADHType1 | jioindiacentral, jioindiawest | $2.4160 |
| Standard_E32ds_v5 | jioindiacentral, jioindiawest, westcentralus, westus2 | $2.4160 |
| Standard_E32ds_v6 | jioindiacentral, jioindiawest | $2.7000 |
| Standard_E32ds_v7 | eastus2 | $3.3260 |
| Standard_E32nds_v6 | jioindiacentral, jioindiawest | $4.3770 |
| Standard_E32ns_v6 | eastus, eastus2, jioindiacentral, jioindiawest | $2.7890 |
| Standard_E32pds_v5 | jioindiacentral, jioindiawest | $1.3120 |
| Standard_E32ps_v5 | jioindiacentral, jioindiawest | $1.0400 |
| Standard_E32s_v3 | jioindiacentral, jioindiawest, westcentralus, westus2 | $2.1920 |
| Standard_E32s_v4 | eastus2, jioindiacentral, jioindiawest, westcentralus | $2.0160 |
| Standard_E32s_v5 | brazilsoutheast, jioindiacentral, jioindiawest | $4.1810 |
| Standard_E32s_v6 | jioindiacentral, jioindiawest | $2.1840 |
| Standard_E372is_v7 | mexicocentral | $39.2260 |
| Standard_E4-2ads_v5 | jioindiacentral, jioindiawest, westus2 | $0.3500 |
| Standard_E4-2ads_v6 | jioindiacentral, jioindiawest | $0.1880 |
| Standard_E4-2ads_v7 | jioindiawest | $0.1880 |
| Standard_E4-2as_v4 | jioindiacentral, jioindiawest | $0.2600 |
| Standard_E4-2as_v5 | jioindiacentral, jioindiawest | $0.1430 |
| Standard_E4-2as_v6 | jioindiacentral, jioindiawest | $0.1540 |
| Standard_E4-2as_v7 | jioindiawest | $0.1540 |
| Standard_E4-2ds_v4 | eastus2, jioindiacentral, jioindiawest | $0.2880 |
| Standard_E4-2ds_v5 | jioindiacentral, jioindiawest | $0.3020 |
| Standard_E4-2ds_v6 | jioindiacentral, jioindiawest, uksouth, westus2 | $0.3370 |
| Standard_E4-2s_v3 | jioindiacentral, jioindiawest | $0.2740 |
| Standard_E4-2s_v4 | jioindiacentral, jioindiawest, westus2 | $0.2600 |
| Standard_E4-2s_v5 | jioindiacentral, jioindiawest | $0.2600 |
| Standard_E4-2s_v6 | jioindiacentral, jioindiawest | $0.2730 |
| Standard_E48_v3 | eastus, jioindiacentral, jioindiawest | $3.0240 |
| Standard_E48_v4 | jioindiacentral, jioindiawest | $3.1200 |
| Standard_E48_v5 | jioindiacentral, jioindiawest | $3.1200 |
| Standard_E48a_v4 | jioindiacentral, jioindiawest, westcentralus | $1.9060 |
| Standard_E48ads_v5 | jioindiacentral, jioindiawest | $4.2010 |
| Standard_E48ads_v6 | jioindiacentral, jioindiawest | $2.2510 |
| Standard_E48ads_v7 | jioindiawest | $2.2510 |
| Standard_E48as_v4 | jioindiacentral, jioindiawest | $1.9060 |
| Standard_E48as_v5 | jioindiacentral, jioindiawest | $1.7160 |
| Standard_E48as_v6 | jioindiacentral, jioindiawest | $1.8480 |
| Standard_E48as_v7 | jioindiawest | $1.8480 |
| Standard_E48bds_v5 | jioindiacentral, jioindiawest | $4.2240 |
| Standard_E48bds_v6 | eastus2, jioindiacentral, jioindiawest, westcentralus, westus2 | $4.6900 |
| Standard_E48bs_v5 | jioindiacentral, jioindiawest | $3.7200 |
| Standard_E48bs_v6 | eastus2, jioindiacentral, jioindiawest | $4.1860 |
| Standard_E48d_v4 | jioindiacentral, jioindiawest | $3.6240 |
| Standard_E48d_v5 | jioindiacentral, jioindiawest | $3.6240 |
| Standard_E48ds_v4 | jioindiacentral, jioindiawest | $3.6240 |
| Standard_E48ds_v5 | jioindiacentral, jioindiawest, westcentralus | $3.6240 |
| Standard_E48ds_v6 | jioindiacentral, jioindiawest, westus2 | $4.0490 |
| Standard_E48ds_v7 | westus2 | $4.9900 |
| Standard_E48nds_v6 | jioindiacentral, jioindiawest | $6.5650 |
| Standard_E48ns_v6 | jioindiacentral, jioindiawest, westcentralus | $5.8570 |
| Standard_E48pds_v6 | brazilsouth, eastus2, westus2 | $4.4590 |
| Standard_E48ps_v6 | mexicocentral | $2.4340 |
| Standard_E48s_v3 | jioindiacentral, jioindiawest, westus2 | $3.2810 |
| Standard_E48s_v4 | jioindiacentral, jioindiawest | $3.1200 |
| Standard_E48s_v5 | eastus, jioindiacentral, jioindiawest, westcentralus, westus2 | $3.0240 |
| Standard_E48s_v6 | jioindiacentral, jioindiawest | $3.2760 |
| Standard_E4_v3 | brazilsouth, jioindiacentral, jioindiawest | $0.4700 |
| Standard_E4_v4 | jioindiacentral, jioindiawest, westus2 | $0.2600 |
| Standard_E4_v5 | eastus2, jioindiacentral, jioindiawest | $0.2520 |
| Standard_E4a_v4 | jioindiacentral, jioindiawest | $0.1580 |
| Standard_E4ads_v5 | jioindiacentral, jioindiawest | $0.3500 |
| Standard_E4ads_v6 | jioindiacentral, jioindiawest | $0.1880 |
| Standard_E4ads_v7 | jioindiawest, westus2 | $0.1880 |
| Standard_E4as_v4 | brazilsoutheast, jioindiacentral, jioindiawest | $0.5230 |
| Standard_E4as_v5 | jioindiacentral, jioindiawest, westus2 | $0.1430 |
| Standard_E4as_v6 | jioindiacentral, jioindiawest, westcentralus | $0.1540 |
| Standard_E4as_v7 | jioindiawest | $0.1540 |
| Standard_E4bds_v5 | jioindiacentral, jioindiawest | $0.3520 |
| Standard_E4bds_v6 | jioindiacentral, jioindiawest | $0.5470 |
| Standard_E4bs_v5 | jioindiacentral, jioindiawest | $0.3100 |
| Standard_E4bs_v6 | eastus2, jioindiacentral, jioindiawest, westus2 | $0.3490 |
| Standard_E4d_v4 | brazilsouth, jioindiacentral, jioindiawest | $0.4560 |
| Standard_E4d_v5 | jioindiacentral, jioindiawest | $0.3020 |
| Standard_E4ds_v4 | brazilsoutheast, jioindiacentral, jioindiawest | $0.5970 |
| Standard_E4ds_v4_ADHType1 | jioindiacentral, jioindiawest | $0.3020 |
| Standard_E4ds_v5 | jioindiacentral, jioindiawest | $0.3020 |
| Standard_E4ds_v6 | jioindiacentral, jioindiawest, uksouth | $0.3370 |
| Standard_E4nds_v6 | jioindiacentral, jioindiawest, westus2 | $0.5470 |
| Standard_E4ns_v6 | jioindiacentral, jioindiawest, westus2 | $0.4880 |
| Standard_E4pds_v5 | eastus2, jioindiacentral, jioindiawest | $0.2300 |
| Standard_E4ps_v5 | jioindiacentral, jioindiawest, uksouth | $0.1300 |
| Standard_E4ps_v6 | westus2 | $0.1840 |
| Standard_E4s_v3 | jioindiacentral, jioindiawest | $0.2740 |
| Standard_E4s_v4 | eastus2, jioindiacentral, jioindiawest | $0.2520 |
| Standard_E4s_v5 | jioindiacentral, jioindiawest | $0.2600 |
| Standard_E4s_v6 | jioindiacentral, jioindiawest, northcentralus | $0.2730 |
| Standard_E64-16ads_v5 | jioindiacentral, jioindiawest | $2.6580 |
| Standard_E64-16ads_v6 | eastus2, jioindiacentral, jioindiawest | $4.6460 |
| Standard_E64-16ads_v7 | jioindiawest | $3.0020 |
| Standard_E64-16as_v4 | jioindiacentral, jioindiawest | $4.1600 |
| Standard_E64-16as_v5 | jioindiacentral, jioindiawest | $5.2320 |
| Standard_E64-16as_v6 | jioindiacentral, jioindiawest | $2.4640 |
| Standard_E64-16as_v7 | jioindiawest | $2.4640 |
| Standard_E64-16ds_v4 | jioindiacentral, jioindiawest, uksouth | $4.8320 |
| Standard_E64-16ds_v5 | jioindiacentral, jioindiawest, westcentralus | $4.8320 |
| Standard_E64-16ds_v6 | jioindiacentral, jioindiawest, westus2 | $5.3990 |
| Standard_E64-16s_v3 | canadacentral, jioindiacentral, jioindiawest | $3.9740 |
| Standard_E64-16s_v4 | jioindiacentral, jioindiawest | $4.1600 |
| Standard_E64-16s_v5 | jioindiacentral, jioindiawest | $4.1600 |
| Standard_E64-16s_v6 | jioindiacentral, jioindiawest, westus2 | $4.3680 |
| Standard_E64-32ads_v5 | eastus2, jioindiacentral, jioindiawest | $4.1920 |
| Standard_E64-32ads_v6 | jioindiacentral, jioindiawest, westcentralus | $3.0020 |
| Standard_E64-32ads_v7 | jioindiawest | $3.0020 |
| Standard_E64-32as_v4 | jioindiacentral, jioindiawest | $4.1600 |
| Standard_E64-32as_v5 | jioindiacentral, jioindiawest | $5.2320 |
| Standard_E64-32as_v6 | jioindiacentral, jioindiawest | $2.4640 |
| Standard_E64-32as_v7 | brazilsouth, jioindiawest | $6.1250 |
| Standard_E64-32ds_v4 | jioindiacentral, jioindiawest, westus2 | $4.8320 |
| Standard_E64-32ds_v5 | jioindiacentral, jioindiawest | $4.8320 |
| Standard_E64-32ds_v6 | jioindiacentral, jioindiawest, westcentralus | $5.3990 |
| Standard_E64-32s_v3 | jioindiacentral, jioindiawest | $3.9370 |
| Standard_E64-32s_v4 | brazilsoutheast, jioindiacentral, jioindiawest, westus2 | $8.3620 |
| Standard_E64-32s_v5 | jioindiacentral, jioindiawest | $4.1600 |
| Standard_E64-32s_v6 | jioindiacentral, jioindiawest, westus2 | $4.3680 |
| Standard_E64_v3 | jioindiacentral, jioindiawest | $3.9370 |
| Standard_E64_v4 | jioindiacentral, jioindiawest | $4.1600 |
| Standard_E64_v5 | eastus, jioindiacentral, jioindiawest | $4.0320 |
| Standard_E64a_v4 | jioindiacentral, jioindiawest | $2.5420 |
| Standard_E64ads_v5 | jioindiacentral, jioindiawest | $5.6020 |
| Standard_E64ads_v6 | jioindiacentral, jioindiawest, westcentralus | $3.0020 |
| Standard_E64ads_v7 | jioindiawest | $3.0020 |
| Standard_E64as_v4 | jioindiacentral, jioindiawest | $2.5420 |
| Standard_E64as_v5 | jioindiacentral, jioindiawest | $5.2320 |
| Standard_E64as_v6 | jioindiacentral, jioindiawest, westus2 | $2.4640 |
| Standard_E64as_v7 | jioindiawest | $2.4640 |
| Standard_E64bds_v5 | jioindiacentral, jioindiawest | $5.6320 |
| Standard_E64bds_v6 | jioindiacentral, jioindiawest, westcentralus | $8.7540 |
| Standard_E64bs_v5 | jioindiacentral, jioindiawest, westus2 | $4.9600 |
| Standard_E64bs_v6 | jioindiacentral, jioindiawest | $7.8130 |
| Standard_E64d_v4 | jioindiacentral, jioindiawest | $4.8320 |
| Standard_E64d_v5 | eastus, jioindiacentral, jioindiawest | $4.6080 |
| Standard_E64ds_v4 | jioindiacentral, jioindiawest, westus2 | $4.8320 |
| Standard_E64ds_v5 | jioindiacentral, jioindiawest | $4.8320 |
| Standard_E64ds_v6 | brazilsouth, jioindiacentral, jioindiawest | $8.3480 |
| Standard_E64ds_v7 | westus2 | $6.6530 |
| Standard_E64is_v3 | westcentralus | $4.1990 |
| Standard_E64nds_v6 | jioindiacentral, jioindiawest | $8.7530 |
| Standard_E64ns_v6 | jioindiacentral, jioindiawest | $7.8100 |
| Standard_E64pds_v6 | eastus2 | $3.7440 |
| Standard_E64s_v3 | jioindiacentral, jioindiawest, westus2 | $3.9370 |
| Standard_E64s_v4 | brazilsoutheast, jioindiacentral, jioindiawest | $8.3620 |
| Standard_E64s_v5 | eastus2, jioindiacentral, jioindiawest | $4.0320 |
| Standard_E64s_v6 | jioindiacentral, jioindiawest, northcentralus | $4.3680 |
| Standard_E8-2ads_v5 | brazilsoutheast, jioindiacentral, jioindiawest | $1.4500 |
| Standard_E8-2ads_v6 | eastus, jioindiacentral, jioindiawest | $0.5810 |
| Standard_E8-2ads_v7 | eastus, jioindiawest | $0.5810 |
| Standard_E8-2as_v4 | jioindiacentral, jioindiawest | $0.5200 |
| Standard_E8-2as_v5 | eastus2, jioindiacentral, jioindiawest, westus2 | $0.4520 |
| Standard_E8-2as_v6 | eastus, jioindiacentral, jioindiawest, westcentralus | $0.4770 |
| Standard_E8-2as_v7 | jioindiawest | $0.3080 |
| Standard_E8-2ds_v4 | jioindiacentral, jioindiawest | $0.6040 |
| Standard_E8-2ds_v5 | jioindiacentral, jioindiawest, westus2 | $0.6040 |
| Standard_E8-2ds_v6 | jioindiacentral, jioindiawest | $0.6750 |
| Standard_E8-2s_v3 | jioindiacentral, jioindiawest, mexicocentral | $0.5480 |
| Standard_E8-2s_v4 | eastus2, jioindiacentral, jioindiawest | $0.5040 |
| Standard_E8-2s_v5 | jioindiacentral, jioindiawest, westus2 | $0.5200 |
| Standard_E8-2s_v6 | jioindiacentral, jioindiawest | $0.5460 |
| Standard_E8-4ads_v5 | jioindiacentral, jioindiawest | $0.7000 |
| Standard_E8-4ads_v6 | brazilsoutheast, eastus2, jioindiacentral, jioindiawest | $1.2080 |
| Standard_E8-4ads_v7 | jioindiawest | $0.3750 |
| Standard_E8-4as_v4 | jioindiacentral, jioindiawest | $0.5200 |
| Standard_E8-4as_v5 | jioindiacentral, jioindiawest, southcentralus | $0.6540 |
| Standard_E8-4as_v6 | jioindiacentral, jioindiawest | $0.3080 |
| Standard_E8-4as_v7 | jioindiawest, mexicocentral | $0.3080 |
| Standard_E8-4ds_v4 | eastus, jioindiacentral, jioindiawest | $0.5760 |
| Standard_E8-4ds_v5 | brazilsouth, jioindiacentral, jioindiawest, westcentralus | $0.9120 |
| Standard_E8-4ds_v6 | jioindiacentral, jioindiawest, westcentralus | $0.6750 |
| Standard_E8-4s_v3 | jioindiacentral, jioindiawest, westcentralus | $0.5480 |
| Standard_E8-4s_v4 | jioindiacentral, jioindiawest, westus2 | $0.5200 |
| Standard_E8-4s_v5 | jioindiacentral, jioindiawest | $0.5200 |
| Standard_E8-4s_v6 | jioindiacentral, jioindiawest | $0.5460 |
| Standard_E80is_v4 | westus2 | $5.0400 |
| Standard_E8_v3 | eastus, jioindiacentral, jioindiawest | $0.5040 |
| Standard_E8_v4 | brazilsouth, jioindiacentral, jioindiawest, uksouth | $0.8040 |
| Standard_E8_v5 | jioindiacentral, jioindiawest | $0.5200 |
| Standard_E8a_v4 | jioindiacentral, jioindiawest, uksouth, westus2 | $0.3170 |
| Standard_E8ads_v5 | jioindiacentral, jioindiawest, westus2 | $0.7000 |
| Standard_E8ads_v6 | jioindiacentral, jioindiawest | $0.3750 |
| Standard_E8ads_v7 | jioindiawest | $0.3750 |
| Standard_E8as_v4 | brazilsouth, jioindiacentral, jioindiawest, uksouth | $0.8040 |
| Standard_E8as_v5 | jioindiacentral, jioindiawest | $0.6540 |
| Standard_E8as_v6 | jioindiacentral, jioindiawest, westus2 | $0.3080 |
| Standard_E8as_v7 | brazilsouth, jioindiawest | $0.7660 |
| Standard_E8bds_v5 | jioindiacentral, jioindiawest | $0.7040 |
| Standard_E8bds_v6 | jioindiacentral, jioindiawest | $1.0940 |
| Standard_E8bs_v5 | jioindiacentral, jioindiawest | $0.6200 |
| Standard_E8bs_v6 | jioindiacentral, jioindiawest | $0.9770 |
| Standard_E8d_v4 | jioindiacentral, jioindiawest, westus2 | $0.6040 |
| Standard_E8d_v5 | jioindiacentral, jioindiawest, westcentralus | $0.6040 |
| Standard_E8ds_v4 | jioindiacentral, jioindiawest, westcentralus | $0.6040 |
| Standard_E8ds_v4_ADHType1 | jioindiacentral, jioindiawest | $0.6040 |
| Standard_E8ds_v5 | brazilsouth, jioindiacentral, jioindiawest | $0.9120 |
| Standard_E8ds_v6 | brazilsoutheast, eastus2, jioindiacentral, jioindiawest | $1.3560 |
| Standard_E8nds_v6 | brazilsoutheast, jioindiacentral, jioindiawest | $2.0320 |
| Standard_E8ns_v6 | jioindiacentral, jioindiawest | $0.9760 |
| Standard_E8pds_v5 | jioindiacentral, jioindiawest | $0.3280 |
| Standard_E8pds_v6 | westus2 | $0.4680 |
| Standard_E8ps_v5 | jioindiacentral, jioindiawest | $0.2600 |
| Standard_E8ps_v6 | eastus2 | $0.3690 |
| Standard_E8s_v3 | brazilsoutheast, jioindiacentral, jioindiawest, uksouth | $1.2220 |
| Standard_E8s_v4 | eastus, jioindiacentral, jioindiawest, mexicocentral | $0.5040 |
| Standard_E8s_v5 | eastus2, jioindiacentral, jioindiawest | $0.5040 |
| Standard_E8s_v6 | jioindiacentral, jioindiawest | $0.5460 |
| Standard_E96-24ads_v5 | jioindiacentral, jioindiawest, westus2 | $3.9860 |
| Standard_E96-24ads_v6 | jioindiacentral, jioindiawest, westus2 | $4.5020 |
| Standard_E96-24ads_v7 | jioindiawest | $4.5020 |
| Standard_E96-24as_v4 | jioindiacentral, jioindiawest | $6.2400 |
| Standard_E96-24as_v5 | brazilsouth, jioindiacentral, jioindiawest | $13.1040 |
| Standard_E96-24as_v6 | brazilsouth, eastus2, jioindiacentral, jioindiawest, westus2 | $9.1870 |
| Standard_E96-24as_v7 | jioindiawest | $3.6960 |
| Standard_E96-24ds_v5 | eastus, jioindiacentral, jioindiawest, westus2 | $6.9120 |
| Standard_E96-24ds_v6 | jioindiacentral, jioindiawest | $8.0990 |
| Standard_E96-24s_v5 | canadacentral, jioindiacentral, jioindiawest | $6.6240 |
| Standard_E96-24s_v6 | jioindiacentral, jioindiawest | $6.5520 |
| Standard_E96-48ads_v5 | jioindiacentral, jioindiawest | $3.9860 |
| Standard_E96-48ads_v6 | jioindiacentral, jioindiawest | $4.5020 |
| Standard_E96-48ads_v7 | jioindiawest | $4.5020 |
| Standard_E96-48as_v4 | eastus2, jioindiacentral, jioindiawest | $6.0480 |
| Standard_E96-48as_v5 | jioindiacentral, jioindiawest, mexicocentral | $3.4320 |
| Standard_E96-48as_v6 | jioindiacentral, jioindiawest | $3.6960 |
| Standard_E96-48as_v7 | brazilsouth, jioindiawest | $9.1870 |
| Standard_E96-48ds_v5 | jioindiacentral, jioindiawest | $7.2480 |
| Standard_E96-48ds_v6 | jioindiacentral, jioindiawest, westus2 | $8.0990 |
| Standard_E96-48s_v5 | canadacentral, jioindiacentral, jioindiawest, westus2 | $6.6240 |
| Standard_E96-48s_v6 | brazilsouth, jioindiacentral, jioindiawest | $10.1310 |
| Standard_E96_v5 | jioindiacentral, jioindiawest | $6.2400 |
| Standard_E96a_v4 | jioindiacentral, jioindiawest, westcentralus | $3.8130 |
| Standard_E96ads_v5 | jioindiacentral, jioindiawest, northcentralus | $3.9860 |
| Standard_E96ads_v6 | jioindiacentral, jioindiawest | $4.5020 |
| Standard_E96ads_v7 | brazilsouth, jioindiawest | $11.1740 |
| Standard_E96as_v4 | jioindiacentral, jioindiawest | $3.8130 |
| Standard_E96as_v5 | jioindiacentral, jioindiawest | $7.8480 |
| Standard_E96as_v6 | jioindiacentral, jioindiawest | $3.6960 |
| Standard_E96as_v7 | jioindiawest | $3.6960 |
| Standard_E96bds_v5 | eastus2, jioindiacentral, jioindiawest | $8.0160 |
| Standard_E96bds_v6 | jioindiacentral, jioindiawest | $13.1310 |
| Standard_E96bs_v5 | brazilsouth, jioindiacentral, jioindiawest | $11.4960 |
| Standard_E96bs_v6 | eastus, jioindiacentral, jioindiawest | $8.3710 |
| Standard_E96d_v5 | brazilsoutheast, jioindiacentral, jioindiawest, westus2 | $14.2270 |
| Standard_E96ds_v5 | eastus, eastus2, jioindiacentral, jioindiawest | $6.9120 |
| Standard_E96ds_v6 | jioindiacentral, jioindiawest | $8.0990 |
| Standard_E96ds_v7 | eastus2 | $9.9790 |
| Standard_E96iads_v5 | jioindiacentral, jioindiawest | $4.3850 |
| Standard_E96ias_v5 | jioindiacentral, jioindiawest | $3.7750 |
| Standard_E96nds_v6 | jioindiacentral, jioindiawest | $13.1300 |
| Standard_E96ns_v6 | canadacentral, jioindiacentral, jioindiawest | $10.0410 |
| Standard_E96s_v5 | jioindiacentral, jioindiawest, westus2 | $6.2400 |
| Standard_E96s_v6 | eastus2, jioindiacentral, jioindiawest | $6.3500 |
| Standard_EC128eds_v5 | eastus2 | $9.2160 |
| Standard_EC128eds_v6 | eastus | $11.5120 |
| Standard_EC16ads_cc_v5 | jioindiacentral, jioindiawest | $0.6640 |
| Standard_EC16ads_v5 | jioindiacentral, jioindiawest | $0.6640 |
| Standard_EC16ads_v6 | eastus2 | $1.2780 |
| Standard_EC16as_cc_v5 | jioindiacentral, jioindiawest | $0.5720 |
| Standard_EC16as_v5 | eastus2, jioindiacentral, jioindiawest | $0.9040 |
| Standard_EC16as_v6 | uksouth | $1.2200 |
| Standard_EC20ads_cc_v5 | jioindiacentral, jioindiawest | $0.8310 |
| Standard_EC20ads_v5 | eastus, jioindiacentral, jioindiawest | $1.3100 |
| Standard_EC20as_cc_v5 | jioindiacentral, jioindiawest | $0.7150 |
| Standard_EC20as_v5 | jioindiacentral, jioindiawest | $0.7150 |
| Standard_EC2ads_v5 | eastus2, jioindiacentral, jioindiawest | $0.1310 |
| Standard_EC2as_v5 | jioindiacentral, jioindiawest | $0.0715 |
| Standard_EC2as_v6 | eastus | $0.1310 |
| Standard_EC32ads_cc_v5 | jioindiacentral, jioindiawest, westus2 | $1.3290 |
| Standard_EC32ads_v5 | jioindiacentral, jioindiawest | $1.3290 |
| Standard_EC32ads_v6 | eastus2 | $2.5560 |
| Standard_EC32as_cc_v5 | jioindiacentral, jioindiawest | $1.1440 |
| Standard_EC32as_v5 | eastus2, jioindiacentral, jioindiawest | $1.8080 |
| Standard_EC48ads_cc_v5 | jioindiacentral, jioindiawest | $1.9930 |
| Standard_EC48ads_v5 | jioindiacentral, jioindiawest | $1.9930 |
| Standard_EC48ads_v6 | eastus2 | $3.8330 |
| Standard_EC48as_cc_v5 | jioindiacentral, jioindiawest, westus2 | $1.7160 |
| Standard_EC48as_v5 | jioindiacentral, jioindiawest | $1.7160 |
| Standard_EC4ads_cc_v5 | jioindiacentral, jioindiawest | $0.1660 |
| Standard_EC4ads_v5 | jioindiacentral, jioindiawest | $0.1660 |
| Standard_EC4as_cc_v5 | jioindiacentral, jioindiawest, westus2 | $0.1430 |
| Standard_EC4as_v5 | jioindiacentral, jioindiawest | $0.1430 |
| Standard_EC4as_v6 | brazilsouth | $0.4210 |
| Standard_EC64ads_cc_v5 | jioindiacentral, jioindiawest | $2.6580 |
| Standard_EC64ads_v5 | jioindiacentral, jioindiawest | $2.6580 |
| Standard_EC64as_cc_v5 | jioindiacentral, jioindiawest | $2.2880 |
| Standard_EC64as_v5 | jioindiacentral, jioindiawest | $2.2880 |
| Standard_EC64as_v6 | westus2 | $4.1960 |
| Standard_EC8ads_cc_v5 | jioindiacentral, jioindiawest | $0.3320 |
| Standard_EC8ads_v5 | jioindiacentral, jioindiawest | $0.3320 |
| Standard_EC8as_cc_v5 | eastus, jioindiacentral, jioindiawest, westus2 | $0.4520 |
| Standard_EC8as_v5 | jioindiacentral, jioindiawest | $0.2860 |
| Standard_EC96ads_cc_v5 | jioindiacentral, jioindiawest | $3.9860 |
| Standard_EC96ads_v5 | jioindiacentral, jioindiawest | $3.9860 |
| Standard_EC96as_cc_v5 | jioindiacentral, jioindiawest | $3.4320 |
| Standard_EC96as_v5 | jioindiacentral, jioindiawest | $3.4320 |
| Standard_EC96as_v6 | westus2 | $6.2940 |
| Standard_EC96iads_v5 | jioindiacentral, jioindiawest | $4.3850 |
| Standard_EC96ias_v5 | jioindiacentral, jioindiawest | $3.7750 |
| Standard_F1 | jioindiacentral, jioindiawest | $0.0490 |
| Standard_F16 | jioindiacentral, jioindiawest | $0.7900 |
| Standard_F16-4amds_v7 | brazilsouth, jioindiawest, uksouth | $2.7970 |
| Standard_F16-4ams_v7 | jioindiawest | $0.9250 |
| Standard_F16-8amds_v7 | jioindiawest | $1.1260 |
| Standard_F16-8ams_v7 | brazilsouth, eastus, jioindiawest, westus2 | $2.2960 |
| Standard_F16ads_v7 | jioindiawest | $0.8860 |
| Standard_F16alds_v7 | jioindiawest | $0.7400 |
| Standard_F16als_v6 | jioindiacentral, jioindiawest | $0.6260 |
| Standard_F16als_v7 | jioindiawest | $0.6260 |
| Standard_F16amds_v7 | jioindiawest, westus2 | $1.1260 |
| Standard_F16ams_v6 | jioindiacentral, jioindiawest | $0.9250 |
| Standard_F16ams_v7 | jioindiawest | $0.9250 |
| Standard_F16as_v6 | eastus2, jioindiacentral, jioindiawest, westcentralus | $1.0930 |
| Standard_F16as_v7 | jioindiawest, uksouth | $0.7060 |
| Standard_F16s | jioindiacentral, jioindiawest | $0.7900 |
| Standard_F16s_v2 | jioindiacentral, jioindiawest | $0.6800 |
| Standard_F1ads_v7 | jioindiawest | $0.0554 |
| Standard_F1alds_v7 | eastus2, jioindiawest | $0.0715 |
| Standard_F1als_v6 | brazilsouth, brazilsoutheast, jioindiacentral, jioindiawest | $0.0970 |
| Standard_F1als_v7 | eastus2, jioindiawest | $0.0605 |
| Standard_F1amds_v7 | eastus, jioindiawest | $0.1090 |
| Standard_F1ams_v6 | jioindiacentral, jioindiawest | $0.0578 |
| Standard_F1ams_v7 | jioindiawest | $0.0578 |
| Standard_F1as_v6 | jioindiacentral, jioindiawest | $0.0441 |
| Standard_F1as_v7 | jioindiawest | $0.0441 |
| Standard_F1s | brazilsoutheast, jioindiacentral, jioindiawest | $0.0933 |
| Standard_F2 | brazilsouth, jioindiacentral, jioindiawest, northcentralus | $0.1440 |
| Standard_F2-1amds_v7 | jioindiawest | $0.1410 |
| Standard_F2-1ams_v7 | jioindiawest | $0.1160 |
| Standard_F2ads_v7 | jioindiawest | $0.1110 |
| Standard_F2alds_v7 | jioindiawest, mexicocentral | $0.0925 |
| Standard_F2als_v6 | jioindiacentral, jioindiawest, mexicocentral | $0.0782 |
| Standard_F2als_v7 | jioindiawest | $0.0782 |
| Standard_F2amds_v7 | eastus2, jioindiawest | $0.2180 |
| Standard_F2ams_v6 | jioindiacentral, jioindiawest | $0.1160 |
| Standard_F2ams_v7 | jioindiawest | $0.1160 |
| Standard_F2as_v6 | jioindiacentral, jioindiawest, mexicocentral | $0.0882 |
| Standard_F2as_v7 | jioindiawest | $0.0882 |
| Standard_F2s | eastus2, jioindiacentral, jioindiawest, westcentralus, westus2 | $0.0990 |
| Standard_F2s_v2 | jioindiacentral, jioindiawest | $0.0850 |
| Standard_F32-16amds_v7 | jioindiawest, uksouth | $2.2530 |
| Standard_F32-16ams_v7 | eastus, jioindiawest | $2.8640 |
| Standard_F32-8amds_v7 | eastus, jioindiawest | $3.4890 |
| Standard_F32-8ams_v7 | jioindiawest | $1.8500 |
| Standard_F32ads_v7 | jioindiawest | $1.7720 |
| Standard_F32alds_v7 | jioindiawest | $1.4800 |
| Standard_F32als_v6 | jioindiacentral, jioindiawest | $1.2510 |
| Standard_F32als_v7 | jioindiawest | $1.2510 |
| Standard_F32amds_v7 | eastus2, jioindiawest | $3.4890 |
| Standard_F32ams_v6 | brazilsouth, jioindiacentral, jioindiawest, northcentralus, westcentralus | $4.5920 |
| Standard_F32ams_v7 | jioindiawest | $1.8500 |
| Standard_F32as_v6 | jioindiacentral, jioindiawest, westcentralus | $1.4110 |
| Standard_F32as_v7 | eastus2, jioindiawest, westus2 | $2.1860 |
| Standard_F32s_v2 | jioindiacentral, jioindiawest | $1.3600 |
| Standard_F4 | jioindiacentral, jioindiawest, westus2 | $0.1980 |
| Standard_F4-1amds_v7 | jioindiawest, uksouth | $0.2820 |
| Standard_F4-1ams_v7 | jioindiawest | $0.2310 |
| Standard_F4-2amds_v7 | jioindiawest | $0.2820 |
| Standard_F4-2ams_v7 | jioindiawest, westus2 | $0.2310 |
| Standard_F48ads_v7 | jioindiawest | $2.6580 |
| Standard_F48alds_v7 | brazilsouth, eastus2, jioindiawest | $5.5060 |
| Standard_F48als_v6 | eastus2, jioindiacentral, jioindiawest | $2.9040 |
| Standard_F48als_v7 | jioindiawest, westus2 | $1.8770 |
| Standard_F48amds_v7 | jioindiawest | $3.3790 |
| Standard_F48ams_v6 | jioindiacentral, jioindiawest | $2.7740 |
| Standard_F48ams_v7 | jioindiawest | $2.7740 |
| Standard_F48as_v6 | jioindiacentral, jioindiawest | $2.1170 |
| Standard_F48as_v7 | jioindiawest | $2.1170 |
| Standard_F48s_v2 | jioindiacentral, jioindiawest | $2.0400 |
| Standard_F4ads_v7 | jioindiawest, westus2 | $0.2210 |
| Standard_F4alds_v7 | jioindiawest | $0.1850 |
| Standard_F4als_v6 | eastus2, jioindiacentral, jioindiawest | $0.2420 |
| Standard_F4als_v7 | brazilsouth, jioindiawest | $0.3880 |
| Standard_F4amds_v7 | jioindiawest | $0.2820 |
| Standard_F4ams_v6 | jioindiacentral, jioindiawest | $0.2310 |
| Standard_F4ams_v7 | jioindiawest | $0.2310 |
| Standard_F4as_v6 | jioindiacentral, jioindiawest | $0.1760 |
| Standard_F4as_v7 | eastus, jioindiawest, westus2 | $0.2730 |
| Standard_F4s | jioindiacentral, jioindiawest | $0.1980 |
| Standard_F4s_v2 | jioindiacentral, jioindiawest | $0.1700 |
| Standard_F64-16amds_v7 | jioindiawest | $4.5060 |
| Standard_F64-16ams_v7 | brazilsouth, eastus2, jioindiawest | $9.1840 |
| Standard_F64-32amds_v7 | jioindiawest, westus2 | $4.5060 |
| Standard_F64-32ams_v7 | jioindiawest | $3.6990 |
| Standard_F64ads_v7 | jioindiawest | $3.5440 |
| Standard_F64alds_v7 | brazilsouth, jioindiawest | $7.3420 |
| Standard_F64als_v6 | jioindiacentral, jioindiawest, westcentralus | $2.5020 |
| Standard_F64als_v7 | jioindiawest | $2.5020 |
| Standard_F64amds_v7 | jioindiawest, northcentralus | $4.5060 |
| Standard_F64ams_v6 | jioindiacentral, jioindiawest, westus2 | $3.6990 |
| Standard_F64ams_v7 | jioindiawest | $3.6990 |
| Standard_F64as_v6 | jioindiacentral, jioindiawest, westcentralus | $2.8220 |
| Standard_F64as_v7 | jioindiawest, mexicocentral | $2.8220 |
| Standard_F64s_v2 | jioindiacentral, jioindiawest | $2.7200 |
| Standard_F72als_v6 | jioindiacentral, jioindiawest, westcentralus | $2.8150 |
| Standard_F72as_v6 | brazilsoutheast, jioindiacentral, jioindiawest | $10.2310 |
| Standard_F72s_v2 | jioindiacentral, jioindiawest, westus2 | $3.0600 |
| Standard_F8 | eastus, jioindiacentral, jioindiawest | $0.3980 |
| Standard_F8-2amds_v7 | jioindiawest, westus2 | $0.5630 |
| Standard_F8-2ams_v7 | jioindiawest | $0.4620 |
| Standard_F8-4amds_v7 | jioindiawest | $0.5630 |
| Standard_F8-4ams_v7 | jioindiawest | $0.4620 |
| Standard_F80ads_v7 | jioindiawest | $4.4290 |
| Standard_F80alds_v7 | eastus2, jioindiawest | $5.7240 |
| Standard_F80als_v7 | jioindiawest | $3.1280 |
| Standard_F80amds_v7 | jioindiawest, mexicocentral | $5.6320 |
| Standard_F80ams_v7 | jioindiawest | $4.6240 |
| Standard_F80as_v7 | canadacentral, jioindiawest | $6.0960 |
| Standard_F8ads_v7 | jioindiawest | $0.4430 |
| Standard_F8alds_v7 | jioindiawest | $0.3700 |
| Standard_F8als_v6 | jioindiacentral, jioindiawest, westus2 | $0.3130 |
| Standard_F8als_v7 | jioindiawest | $0.3130 |
| Standard_F8amds_v7 | eastus2, jioindiawest | $0.8720 |
| Standard_F8ams_v6 | brazilsoutheast, jioindiacentral, jioindiawest | $1.4900 |
| Standard_F8ams_v7 | jioindiawest, westus2 | $0.4620 |
| Standard_F8as_v6 | jioindiacentral, jioindiawest, westcentralus | $0.3530 |
| Standard_F8as_v7 | jioindiawest, westus2 | $0.3530 |
| Standard_F8s | jioindiacentral, jioindiawest | $0.3950 |
| Standard_F8s_v2 | jioindiacentral, jioindiawest | $0.3400 |
| Standard_FX12mds_v2 | westus3 | $1.6410 |
| Standard_FX16-4mds_v2 | eastus2 | $2.1880 |
| Standard_FX16-4ms_v2 | eastus | $1.7710 |
| Standard_FX16-8mds_v2 | eastus2, uksouth, westcentralus | $2.1880 |
| Standard_FX16-8ms_v2 | eastus | $1.7710 |
| Standard_FX24-12mds_v2 | brazilsouth | $6.4030 |
| Standard_FX24-6ms_v2 | eastus2 | $2.6560 |
| Standard_FX24mds_v2 | brazilsouth | $6.4030 |
| Standard_FX2mds_v2 | westus2 | $0.2740 |
| Standard_FX32-16ms_v2 | eastus2 | $3.5410 |
| Standard_FX32-8mds_v2 | brazilsouth | $8.5370 |
| Standard_FX32-8ms_v2 | eastus, westcentralus | $3.5410 |
| Standard_FX32mds_v2 | westus2 | $4.3780 |
| Standard_FX32ms_v2 | brazilsouth | $6.9070 |
| Standard_FX36mds | brazilsouth | $5.0040 |
| Standard_FX48-12mds_v2 | eastus2 | $6.5650 |
| Standard_FX4ms_v2 | brazilsouth, westcentralus | $0.8630 |
| Standard_FX64ms_v2 | westus2 | $7.0840 |
| Standard_FX8-4mds_v2 | eastus | $1.0940 |
| Standard_FX96-24mds_v2 | brazilsouth | $25.6110 |
| Standard_FX96-48mds_v2 | westcentralus | $15.7560 |
| Standard_FX96mds_v2 | westus2 | $13.1330 |
| Standard_FX96ms_v2 | eastus2 | $10.6230 |
| Standard_G2 | westus2 | $0.9810 |
| Standard_GS1 | westus2 | $0.4900 |
| Standard_GS2 | eastus2 | $1.1000 |
| Standard_GS3 | eastus2, westus2 | $2.2000 |
| Standard_H16_Promo | australiaeast, centralindia, eastus, japaneast, northcentralus (+7 more) | $1.4220 |
| Standard_H16m_Promo | australiaeast, centralindia, eastus, japaneast, northcentralus (+7 more) | $1.9060 |
| Standard_H16mr_Promo | australiaeast, centralindia, eastus, japaneast, northcentralus (+7 more) | $2.0960 |
| Standard_H16r_Promo | australiaeast, centralindia, eastus, japaneast, northcentralus (+7 more) | $1.5650 |
| Standard_H8_Promo | australiaeast, centralindia, eastus, japaneast, northcentralus (+7 more) | $0.7110 |
| Standard_H8m | westus2 | $1.0660 |
| Standard_H8m_Promo | australiaeast, centralindia, eastus, japaneast, northcentralus (+7 more) | $0.9530 |
| Standard_HB120-32rs_v2 | canadacentral | $4.3200 |
| Standard_HB120-64rs_v2 | brazilsouth, westus2 | $7.2000 |
| Standard_HB120-96rs_v3 | mexicocentral | $3.9600 |
| Standard_HB176-24rs_v4 | eastus | $7.2000 |
| Standard_HB368-96rs_v5 | eastus | $19.8000 |
| Standard_HB60-30rs | eastus | $2.2800 |
| Standard_HC44-16rs | brazilsouth, eastus, eastus2 | $6.3360 |
| Standard_HX176-24rs | eastus2 | $8.6400 |
| Standard_L128aos_v5 | eastus2 | $19.5070 |
| Standard_L12aos_v5 | uksouth | $2.2860 |
| Standard_L16as_v3 | jioindiacentral, jioindiawest | $1.4160 |
| Standard_L16s_v3 | eastus2, jioindiacentral, jioindiawest | $1.3960 |
| Standard_L16s_v4 | westus2 | $1.3960 |
| Standard_L24aos_v5 | eastus2 | $3.6580 |
| Standard_L2aos_v5 | eastus | $0.3050 |
| Standard_L2as_v4 | eastus, westus2 | $0.1720 |
| Standard_L2s_v4 | brazilsouth | $0.2780 |
| Standard_L32aos_v4 | brazilsouth | $5.7600 |
| Standard_L32as_v3 | jioindiacentral, jioindiawest | $2.8320 |
| Standard_L32s_v3 | jioindiacentral, jioindiawest | $3.1680 |
| Standard_L32s_v4 | eastus2, mexicocentral | $2.7920 |
| Standard_L48as_v3 | brazilsoutheast, jioindiacentral, jioindiawest | $7.7690 |
| Standard_L48as_v4 | uksouth | $4.8000 |
| Standard_L48s_v2 | westus2 | $3.7440 |
| Standard_L48s_v3 | jioindiacentral, jioindiawest | $4.7520 |
| Standard_L48s_v4 | mexicocentral | $4.5940 |
| Standard_L64as_v3 | jioindiacentral, jioindiawest | $5.6640 |
| Standard_L64as_v5 | westus2 | $6.5120 |
| Standard_L64s_v3 | eastus2, jioindiacentral, jioindiawest | $5.5840 |
| Standard_L80as_v3 | eastus2, jioindiacentral, jioindiawest | $6.2400 |
| Standard_L80s_v3 | jioindiacentral, jioindiawest | $7.9200 |
| Standard_L8as_v3 | jioindiacentral, jioindiawest | $0.7080 |
| Standard_L8s_v2 | westus2 | $0.6240 |
| Standard_L8s_v3 | eastus2, jioindiacentral, jioindiawest, westus2 | $0.6980 |
| Standard_M128 | jioindiacentral, jioindiawest, westus2 | $13.7620 |
| Standard_M128-32ms | eastus, jioindiacentral, jioindiawest | $26.6880 |
| Standard_M128-64bds_3_v3 | jioindiacentral, jioindiawest | $45.9600 |
| Standard_M128-64bds_v3 | jioindiacentral, jioindiawest | $11.7100 |
| Standard_M128-64bs_v3 | jioindiacentral, jioindiawest | $11.1600 |
| Standard_M128-64ms | brazilsoutheast, jioindiacentral, jioindiawest, westus2 | $60.0480 |
| Standard_M128bds_2_v4 | eastus | $0.0000 |
| Standard_M128bds_3_v3 | jioindiacentral, jioindiawest | $45.9600 |
| Standard_M128bds_v3 | brazilsouth, jioindiacentral, jioindiawest | $23.4200 |
| Standard_M128bs_v3 | jioindiacentral, jioindiawest | $11.1600 |
| Standard_M128dms_v2 | eastus2, jioindiacentral, jioindiawest | $26.6900 |
| Standard_M128ds_v2 | brazilsouth, jioindiacentral, jioindiawest | $26.0130 |
| Standard_M128m | jioindiacentral, jioindiawest | $27.5240 |
| Standard_M128ms | eastus, jioindiacentral, jioindiawest | $26.6880 |
| Standard_M128ms_v2 | jioindiacentral, jioindiawest | $27.1020 |
| Standard_M128s | jioindiacentral, jioindiawest, westus2 | $13.7620 |
| Standard_M128s_v2 | jioindiacentral, jioindiawest, uksouth | $13.3290 |
| Standard_M12ds_v3 | jioindiawest | $1.6800 |
| Standard_M12s_v3 | jioindiawest | $1.5800 |
| Standard_M16-4ms | jioindiacentral, jioindiawest, mexicocentral | $3.1700 |
| Standard_M16-8ms | jioindiacentral, jioindiawest | $3.1700 |
| Standard_M16bds_v3 | jioindiacentral, jioindiawest | $1.4600 |
| Standard_M16bs_v3 | jioindiacentral, jioindiawest | $1.4000 |
| Standard_M16lds_v4 | eastus2 | $2.1895 |
| Standard_M16ls_v4 | westus2 | $2.0852 |
| Standard_M16ms | eastus2, jioindiacentral, jioindiawest | $3.0730 |
| Standard_M16s | jioindiacentral, jioindiawest | $2.4620 |
| Standard_M16s_v3 | jioindiacentral, jioindiawest | $1.6892 |
| Standard_M176-88bds_4_v3 | jioindiacentral, jioindiawest, westus2 | $62.8300 |
| Standard_M176-88bds_v3 | eastus2, jioindiacentral, jioindiawest | $16.1100 |
| Standard_M176-88bs_v3 | jioindiacentral, jioindiawest | $15.3400 |
| Standard_M176bds_4_v3 | jioindiacentral, jioindiawest, westus2 | $62.8300 |
| Standard_M176bds_v3 | jioindiacentral, jioindiawest | $16.1100 |
| Standard_M176bs_v3 | jioindiacentral, jioindiawest, uksouth | $15.3400 |
| Standard_M176ds-44_3_v3 | westus2 | $19.0000 |
| Standard_M176ds-44_4_v3 | westus2 | $26.4700 |
| Standard_M176ds_3_v3 | jioindiawest | $19.5700 |
| Standard_M176ds_4_v3 | jioindiawest | $27.2641 |
| Standard_M176ds_v3 | jioindiacentral, jioindiawest | $16.0667 |
| Standard_M176s-44_3_v3 | westus2 | $17.8800 |
| Standard_M176s-88_4_v3 | eastus2 | $24.9100 |
| Standard_M176s_3_v3 | jioindiawest | $18.4164 |
| Standard_M176s_4_v3 | eastus2, jioindiawest | $24.9100 |
| Standard_M176s_v3 | jioindiacentral, jioindiawest | $15.1118 |
| Standard_M192idms_v2 | jioindiacentral, jioindiawest | $33.0800 |
| Standard_M192ids_v2 | jioindiacentral, jioindiawest | $16.5400 |
| Standard_M192ims_v2 | eastus2, jioindiacentral, jioindiawest | $31.6430 |
| Standard_M192is_v2 | jioindiacentral, jioindiawest, westus2 | $16.1060 |
| Standard_M208-52ms_v2 | eastus2 | $44.6200 |
| Standard_M208ms_v2 | eastus, jioindiacentral, jioindiawest | $44.6200 |
| Standard_M208s_v2 | jioindiacentral, jioindiawest, westus2 | $22.9790 |
| Standard_M24ds_v3 | jioindiawest | $3.3578 |
| Standard_M24s-12_v3 | westus2 | $3.0700 |
| Standard_M24s_v3 | jioindiawest, westus2 | $3.1621 |
| Standard_M32-16ms | jioindiacentral, jioindiawest | $6.3410 |
| Standard_M32-8ms | jioindiacentral, jioindiawest | $6.3410 |
| Standard_M32bds_v3 | brazilsouth, jioindiacentral, jioindiawest | $5.8600 |
| Standard_M32bs_v3 | jioindiacentral, jioindiawest | $2.7900 |
| Standard_M32dms_v2 | jioindiacentral, jioindiawest | $6.3410 |
| Standard_M32lds_v4 | eastus2, westus2 | $4.1053 |
| Standard_M32ls | jioindiacentral, jioindiawest, westus2 | $2.9641 |
| Standard_M32mds_v4 | eastus2 | $8.7580 |
| Standard_M32ms | jioindiacentral, jioindiawest | $6.3410 |
| Standard_M32ms_v2 | jioindiacentral, jioindiawest | $6.2320 |
| Standard_M32s | eastus2, jioindiacentral, jioindiawest | $3.3350 |
| Standard_M32ts | jioindiacentral, jioindiawest | $2.7928 |
| Standard_M32ts_v3 | jioindiacentral, jioindiawest | $3.3784 |
| Standard_M416-208ms_v2 | jioindiacentral, jioindiawest | $149.7200 |
| Standard_M416-208s_v2 | brazilsouth, jioindiacentral, jioindiawest, westus2 | $96.6810 |
| Standard_M416ds-104_6_v3 | eastus2, jioindiacentral, jioindiawest | $58.0665 |
| Standard_M416ds-104_8_v3 | jioindiacentral, jioindiawest | $79.8793 |
| Standard_M416ds-208_6_v3 | jioindiacentral, jioindiawest | $59.9095 |
| Standard_M416ds-208_8_v3 | eastus2, jioindiacentral, jioindiawest | $77.4220 |
| Standard_M416ds_10_v3 | jioindiacentral, jioindiawest | $99.6388 |
| Standard_M416ds_12_v3 | jioindiacentral, jioindiawest | $119.5336 |
| Standard_M416ds_6_v3 | eastus2, jioindiacentral, jioindiawest | $58.0665 |
| Standard_M416ds_8_v3 | jioindiacentral, jioindiawest | $79.8750 |
| Standard_M416ds_9_v3 | jioindiacentral, jioindiawest | $89.9016 |
| Standard_M416is_v2 | westus2 | $49.5800 |
| Standard_M416ms_v2 | jioindiacentral, jioindiawest | $102.1200 |
| Standard_M416s-104_6_v3 | jioindiacentral, jioindiawest | $57.0567 |
| Standard_M416s-104_8_v3 | jioindiacentral, jioindiawest | $76.0756 |
| Standard_M416s-208_6_v3 | jioindiacentral, jioindiawest | $57.0567 |
| Standard_M416s-208_8_v3 | jioindiacentral, jioindiawest | $76.0756 |
| Standard_M416s_10_v2 | brazilsouth, eastus2, jioindiacentral, jioindiawest | $131.4000 |
| Standard_M416s_10_v3 | brazilsouth, jioindiacentral, jioindiawest | $147.8818 |
| Standard_M416s_12_v3 | brazilsouth, jioindiacentral, jioindiawest | $176.9280 |
| Standard_M416s_6_v3 | jioindiacentral, jioindiawest | $57.0600 |
| Standard_M416s_8_v2 | brazilsouth, jioindiacentral, jioindiawest | $105.1200 |
| Standard_M416s_8_v3 | jioindiacentral, jioindiawest | $76.0770 |
| Standard_M416s_9_v2 | jioindiacentral, jioindiawest | $104.3500 |
| Standard_M416s_9_v3 | eastus2, jioindiacentral, jioindiawest | $83.1267 |
| Standard_M416s_v2 | jioindiacentral, jioindiawest | $51.0670 |
| Standard_M48bds_v3 | jioindiacentral, jioindiawest | $4.3900 |
| Standard_M48bs_v3 | jioindiacentral, jioindiawest | $4.1800 |
| Standard_M48ds_1_v3 | jioindiawest | $6.8186 |
| Standard_M48ls_v4 | eastus2 | $6.2557 |
| Standard_M48s-24_1_v3 | eastus2 | $6.2300 |
| Standard_M48s_1_v3 | jioindiawest | $6.4169 |
| Standard_M624ds-156_12_v3 | jioindiacentral, jioindiawest | $119.8190 |
| Standard_M624ds-312_12_v3 | jioindiacentral, jioindiawest | $119.8190 |
| Standard_M624ds_12_v3 | jioindiacentral, jioindiawest | $119.8170 |
| Standard_M624s-156_12_v3 | jioindiacentral, jioindiawest | $114.1133 |
| Standard_M624s-312_12_v3 | jioindiacentral, jioindiawest | $114.1133 |
| Standard_M624s_12_v3 | eastus2, jioindiacentral, jioindiawest | $110.6028 |
| Standard_M64 | jioindiacentral, jioindiawest | $6.8810 |
| Standard_M64-16ms | jioindiacentral, jioindiawest, uksouth | $10.6660 |
| Standard_M64-32bds_1_v3 | jioindiacentral, jioindiawest | $22.9800 |
| Standard_M64-32ms | jioindiacentral, jioindiawest | $10.6660 |
| Standard_M64-32ms_v2 | eastus2 | $10.1300 |
| Standard_M64bds_1_v3 | jioindiacentral, jioindiawest, uksouth | $22.9800 |
| Standard_M64bds_v3 | jioindiacentral, jioindiawest | $5.8600 |
| Standard_M64bs_v3 | jioindiacentral, jioindiawest | $5.5800 |
| Standard_M64dms_v2 | jioindiacentral, jioindiawest | $10.6680 |
| Standard_M64ds_v2 | jioindiacentral, jioindiawest | $6.8800 |
| Standard_M64ds_v4 | westus2 | $13.1370 |
| Standard_M64ls | jioindiacentral, jioindiawest | $5.5867 |
| Standard_M64ls_v4 | westus2 | $7.9337 |
| Standard_M64m | eastus2, jioindiacentral, jioindiawest | $10.3370 |
| Standard_M64ms | jioindiacentral, jioindiawest | $10.6660 |
| Standard_M64ms_v2 | jioindiacentral, jioindiawest | $10.4510 |
| Standard_M64s | jioindiacentral, jioindiawest | $6.8810 |
| Standard_M64s_v2 | jioindiacentral, jioindiawest | $6.6630 |
| Standard_M8-2ms | jioindiacentral, jioindiawest | $1.5850 |
| Standard_M8-4ms | jioindiacentral, jioindiawest | $1.5850 |
| Standard_M832ds-208_12_v3 | jioindiacentral, jioindiawest | $120.2039 |
| Standard_M832ds-416_12_v3 | eastus2, jioindiacentral, jioindiawest | $116.5060 |
| Standard_M832ds_12_v3 | eastus2, jioindiacentral, jioindiawest | $116.5060 |
| Standard_M832ids-208_16_v3 | jioindiacentral, jioindiawest | $177.5096 |
| Standard_M832ids-416_16_v3 | jioindiacentral, jioindiawest | $177.5096 |
| Standard_M832ids_16_v3 | jioindiacentral, jioindiawest | $177.5100 |
| Standard_M832is-208_16_v3 | jioindiacentral, jioindiawest | $169.0568 |
| Standard_M832is-416_16_v3 | brazilsouth, jioindiacentral, jioindiawest | $261.3893 |
| Standard_M832is_16_v3 | jioindiacentral, jioindiawest, westus2 | $169.0600 |
| Standard_M832s-208_12_v3 | jioindiacentral, jioindiawest | $114.4798 |
| Standard_M832s-416_12_v3 | jioindiacentral, jioindiawest | $114.4798 |
| Standard_M832s_12_v3 | jioindiacentral, jioindiawest | $114.4784 |
| Standard_M8ms | brazilsouth, jioindiacentral, jioindiawest | $2.9960 |
| Standard_M96-48bds_2_v3 | jioindiacentral, jioindiawest | $34.4700 |
| Standard_M96bds_2_v3 | jioindiacentral, jioindiawest | $34.4700 |
| Standard_M96bds_v3 | eastus2, jioindiacentral, jioindiawest | $8.7900 |
| Standard_M96bs_v3 | jioindiacentral, jioindiawest | $8.3700 |
| Standard_M96ds-24_2_v3 | westus2 | $13.2300 |
| Standard_M96ds_1_v3 | jioindiawest | $10.2279 |
| Standard_M96ds_2_v3 | jioindiawest | $13.6269 |
| Standard_M96ds_v3 | jioindiacentral, jioindiawest | $10.7112 |
| Standard_M96s-48_1_v3 | eastus2 | $9.3400 |
| Standard_M96s_1_v3 | jioindiawest, uksouth | $9.6202 |
| Standard_M96s_2_v3 | jioindiawest | $12.8235 |
| Standard_M96s_v3 | jioindiacentral, jioindiawest | $10.0853 |
| Standard_NC12_Promo | australiaeast, eastus, eastus2, northcentralus, northeurope (+4 more) | $1.3660 |
| Standard_NC12s_v3 | jioindiawest | $8.4680 |
| Standard_NC132lds_xl_RTXPRO6000BSE_v6 | eastus2 | $5.5000 |
| Standard_NC24_Promo | australiaeast, eastus, eastus2, northcentralus, northeurope (+4 more) | $2.7320 |
| Standard_NC24ads_A100_v4 | jioindiacentral, jioindiawest, westus2 | $5.1420 |
| Standard_NC24r_Promo | australiaeast, eastus, northcentralus, northeurope, southcentralus (+3 more) | $3.0060 |
| Standard_NC24rs_v3 | jioindiawest | $18.6300 |
| Standard_NC24s_v3 | brazilsouth, jioindiawest | $24.4800 |
| Standard_NC256ds_xl_RTXPRO6000BSE_v6 | eastus2 | $12.7600 |
| Standard_NC256lds_xl_RTXPRO6000BSE_v6 | westus2 | $11.0000 |
| Standard_NC264ds_xl_RTXPRO6000BSE_v6 | westus2 | $12.7600 |
| Standard_NC32lds_xl_RTXPRO6000BSE_v6 | eastus2 | $1.2430 |
| Standard_NC40ads_H100_v5 | jioindiacentral, jioindiawest | $9.7720 |
| Standard_NC48ads_A100_v4 | jioindiacentral, jioindiawest, uksouth | $10.2840 |
| Standard_NC64as_T4_v3 | eastus2 | $4.3520 |
| Standard_NC64lds_xl_RTXPRO6000BSE_v6 | eastus | $2.4400 |
| Standard_NC6_Promo | australiaeast, eastus, eastus2, northcentralus, northeurope (+4 more) | $0.6830 |
| Standard_NC6s_v3 | jioindiawest | $4.2340 |
| Standard_NC72lds_xl_RTXPRO6000BSE_v6 | westus2 | $2.4400 |
| Standard_NC80adis_H100_v5 | jioindiacentral, jioindiawest | $19.5440 |
| Standard_NC8as_T4_v3 | eastus2, westus2 | $0.7520 |
| Standard_NC96ads_A100_v4 | jioindiacentral, jioindiawest | $20.5690 |
| Standard_ND96is_H100_v5 | westus2 | $88.4880 |
| Standard_ND96is_flex_H100_v5 | eastus2 | $88.4880 |
| Standard_ND96isrf_H100_v5 | westcentralus | $127.8160 |
| Standard_NG32adms_V620 | eastus2 | $3.3020 |
| Standard_NV12_Promo | australiaeast, centralindia, eastus, eastus2, japaneast (+8 more) | $1.9080 |
| Standard_NV12ads_A10_v5 | jioindiacentral, jioindiawest | $1.2710 |
| Standard_NV12s_v3 | brazilsouth, westus2 | $1.9380 |
| Standard_NV16as_v4 | eastus2 | $0.9320 |
| Standard_NV18ads_A10_v5 | jioindiacentral, jioindiawest | $2.2400 |
| Standard_NV24_Promo | australiaeast, centralindia, eastus, eastus2, japaneast (+8 more) | $3.8150 |
| Standard_NV32as_v4 | eastus2 | $1.8640 |
| Standard_NV36adms_A10_v5 | jioindiacentral, jioindiawest | $6.3280 |
| Standard_NV36ads_A10_v5 | jioindiacentral, jioindiawest | $4.4800 |
| Standard_NV6_Promo | australiaeast, centralindia, eastus, eastus2, japaneast (+8 more) | $0.9540 |
| Standard_NV6ads_A10_v5 | jioindiacentral, jioindiawest, southcentralus, westus2 | $0.6360 |
| Standard_NV72ads_A10_v5 | jioindiacentral, jioindiawest | $9.1280 |
| Standard_PB24s | brazilsoutheast | $3.3200 |
| Standard_SQLG7_NVME | jioindiacentral | $7.8520 |

## Spot Only (no on-demand pricing)

| SKU | Regions (679 unique SKUs) | Spot Price (first region) |
|-----|---------|--------|
| Basic_A1 | brazilsoutheast | $0.0480 |
| Basic_A2 | westus2 | $0.0612 |
| Basic_A3 | westus2 | $0.1584 |
| Basic_A4 | brazilsouth | $0.4176 |
| FXmds Type1 | eastus2 | $0.9820 |
| Standard_A0 | uksouth | $0.0198 |
| Standard_A10 | brazilsoutheast | $0.1170 |
| Standard_A1_v2 | eastus | $0.0084 |
| Standard_A2 | eastus2, westus2 | $0.1080 |
| Standard_A2m_v2 | eastus2 | $0.0570 |
| Standard_A4 | eastus2 | $0.4320 |
| Standard_A6 | westcentralus | $0.3960 |
| Standard_B16pls_v2 | westcentralus | $0.5139 |
| Standard_B16s_v2 | westus2 | $0.5994 |
| Standard_D1 | eastus2 | $0.0412 |
| Standard_D128ads_v7 | eastus | $1.4154 |
| Standard_D128als_v7 | mexicocentral, southcentralus | $1.0456 |
| Standard_D128as_v7 | mexicocentral | $1.1803 |
| Standard_D128ds_v6 | westcentralus | $1.7682 |
| Standard_D128ds_v7 | westus2 | $1.9297 |
| Standard_D128ls_v7 | westus2 | $1.3854 |
| Standard_D128nds_v6 | eastus | $2.1448 |
| Standard_D128nlds_v6 | brazilsoutheast, eastus2 | $3.3098 |
| Standard_D128nls_v6 | eastus2 | $1.3413 |
| Standard_D128ns_v6 | westus2 | $1.6468 |
| Standard_D12_v2 | eastus2 | $0.1839 |
| Standard_D13 | eastus2 | $0.4269 |
| Standard_D14_v2 | eastus2 | $0.7358 |
| Standard_D160ads_v7 | brazilsouth | $2.7055 |
| Standard_D160alds_v7 | eastus2 | $1.4074 |
| Standard_D160als_v7 | eastus2 | $1.1886 |
| Standard_D160as_v7 | westus2 | $1.3424 |
| Standard_D16_v5 | westus2 | $0.1682 |
| Standard_D16a_v4 | northcentralus, westus2 | $0.1419 |
| Standard_D16ads_v6 | westus2 | $0.1685 |
| Standard_D16als_v7 | eastus | $0.1247 |
| Standard_D16as_v6 | mexicocentral | $0.1475 |
| Standard_D16as_v7 | eastus | $0.1408 |
| Standard_D16d_v4 | westus2 | $0.1852 |
| Standard_D16ds_v4 | mexicocentral | $0.1837 |
| Standard_D16ds_v6 | brazilsouth | $0.2935 |
| Standard_D16lds_v5 | eastus | $0.1621 |
| Standard_D16lds_v6 | canadacentral | $0.1785 |
| Standard_D16lds_v7 | northcentralus | $0.1966 |
| Standard_D16ls_v5 | brazilsoutheast | $0.2517 |
| Standard_D16ls_v6 | eastus2, westus2 | $0.1319 |
| Standard_D16ls_v7 | westus2 | $0.1732 |
| Standard_D16nls_v6 | eastus2 | $0.1676 |
| Standard_D16ns_v6 | eastus2, westus2 | $0.2059 |
| Standard_D16pds_v5 | eastus | $0.1528 |
| Standard_D16pds_v6 | eastus | $0.2648 |
| Standard_D16plds_v5 | westus2 | $0.1135 |
| Standard_D16plds_v6 | eastus, eastus2 | $0.2251 |
| Standard_D16pls_v5 | westcentralus | $0.1207 |
| Standard_D16pls_v6 | uksouth | $0.0964 |
| Standard_D16ps_v5 | westus2 | $0.1138 |
| Standard_D16s_v5 | eastus, eastus2 | $0.1621 |
| Standard_D16s_v6 | eastus2 | $0.1489 |
| Standard_D16s_v7 | mexicocentral | $0.2158 |
| Standard_D192s_v7 | eastus | $2.3471 |
| Standard_D192vds_v7 | eastus | $3.1326 |
| Standard_D248lds_v7 | westus2 | $3.1760 |
| Standard_D2_v2 | northcentralus, uksouth | $0.0297 |
| Standard_D2ads_v5 | brazilsouth | $0.0306 |
| Standard_D2ads_v6 | eastus2, westus2 | $0.0211 |
| Standard_D2ads_v7 | eastus2 | $0.0211 |
| Standard_D2alds_v7 | northcentralus | $0.0176 |
| Standard_D2als_v6 | westcentralus | $0.0178 |
| Standard_D2ds_v4 | westus2 | $0.0232 |
| Standard_D2ds_v5 | eastus | $0.0239 |
| Standard_D2lds_v5 | westcentralus | $0.0213 |
| Standard_D2lds_v6 | westcentralus | $0.0244 |
| Standard_D2lds_v7 | westus2 | $0.0257 |
| Standard_D2ls_v6 | westus2 | $0.0165 |
| Standard_D2nlds_v6 | eastus2 | $0.0259 |
| Standard_D2ns_v6 | brazilsoutheast | $0.0669 |
| Standard_D2pds_v5 | canadacentral | $0.0187 |
| Standard_D2s_v3 | brazilsoutheast, mexicocentral | $0.0383 |
| Standard_D2s_v4 | eastus, westcentralus, westus2 | $0.0202 |
| Standard_D3 | brazilsoutheast | $0.1006 |
| Standard_D32_v5 | eastus2 | $0.6116 |
| Standard_D32ads_v6 | brazilsouth | $0.5411 |
| Standard_D32alds_v6 | eastus | $0.3210 |
| Standard_D32alds_v7 | southcentralus | $0.3371 |
| Standard_D32als_v7 | brazilsouth | $0.3814 |
| Standard_D32as_v4 | brazilsoutheast, mexicocentral, westcentralus | $0.5880 |
| Standard_D32as_v7 | uksouth | $0.3123 |
| Standard_D32d_v4 | eastus | $0.3809 |
| Standard_D32lds_v5 | canadacentral, eastus2 | $0.7094 |
| Standard_D32lds_v7 | mexicocentral | $0.4343 |
| Standard_D32ls_v5 | westus2 | $0.2978 |
| Standard_D32nds_v6 | mexicocentral | $0.5176 |
| Standard_D32nlds_v6 | westus2 | $0.4141 |
| Standard_D32nls_v6 | mexicocentral | $0.3689 |
| Standard_D32ns_v6 | eastus2 | $0.4117 |
| Standard_D32pds_v6 | westcentralus, westus2 | $0.3258 |
| Standard_D32plds_v6 | mexicocentral, westus2 | $0.2537 |
| Standard_D32pls_v6 | westcentralus, westus2 | $0.2199 |
| Standard_D32ps_v5 | eastus2 | $0.2277 |
| Standard_D32ps_v6 | westus2 | $0.2075 |
| Standard_D32s_v4 | westus2 | $0.3364 |
| Standard_D32s_v6 | westcentralus | $0.3576 |
| Standard_D32s_v7 | brazilsouth, eastus2 | $0.6259 |
| Standard_D372ds_v7 | westus2 | $5.6083 |
| Standard_D3_v2 | brazilsouth, eastus | $0.0698 |
| Standard_D48_v3 | eastus2, westcentralus | $1.1032 |
| Standard_D48_v5 | mexicocentral | $0.4683 |
| Standard_D48ads_v6 | westus2 | $0.5056 |
| Standard_D48ads_v7 | brazilsouth, eastus | $0.8116 |
| Standard_D48alds_v6 | southcentralus | $0.5056 |
| Standard_D48alds_v7 | eastus2 | $0.4223 |
| Standard_D48als_v7 | canadacentral | $0.3973 |
| Standard_D48as_v4 | brazilsouth, eastus2 | $0.6808 |
| Standard_D48as_v5 | eastus2 | $0.3820 |
| Standard_D48as_v7 | eastus | $0.4227 |
| Standard_D48ds_v5 | westus2 | $0.5706 |
| Standard_D48lds_v5 | brazilsoutheast | $0.8591 |
| Standard_D48lds_v6 | brazilsouth | $0.7540 |
| Standard_D48ls_v5 | eastus2, westcentralus | $0.8123 |
| Standard_D48ls_v6 | eastus | $0.4517 |
| Standard_D48nds_v6 | eastus2, westus2 | $0.7058 |
| Standard_D48pds_v6 | westus2 | $0.4071 |
| Standard_D48plds_v6 | eastus2 | $0.3459 |
| Standard_D48pls_v6 | westcentralus | $0.3301 |
| Standard_D48s_v4 | eastus, westus2 | $0.4859 |
| Standard_D4_v2 | brazilsouth | $0.1395 |
| Standard_D4_v3 | westus2 | $0.0355 |
| Standard_D4ads_v6 | eastus2 | $0.0421 |
| Standard_D4alds_v6 | westcentralus | $0.0421 |
| Standard_D4als_v6 | eastus, westus2 | $0.0337 |
| Standard_D4as_v7 | eastus2 | $0.0336 |
| Standard_D4ds_v5 | westus2 | $0.0476 |
| Standard_D4ds_v7 | westus2 | $0.0602 |
| Standard_D4lds_v7 | brazilsouth, westus2 | $0.0787 |
| Standard_D4nlds_v6 | canadacentral | $0.0565 |
| Standard_D4pds_v6 | westcentralus | $0.0407 |
| Standard_D4plds_v6 | eastus, westcentralus | $0.0563 |
| Standard_D4pls_v5 | westcentralus, westus2 | $0.0301 |
| Standard_D4ps_v5 | westcentralus | $0.0342 |
| Standard_D4s_v3 | eastus | $0.0376 |
| Standard_D4s_v4 | eastus2, westus2 | $0.0765 |
| Standard_D4s_v5 | westus2 | $0.0420 |
| Standard_D64_v3 | brazilsoutheast, westcentralus | $1.2223 |
| Standard_D64_v4 | brazilsoutheast, eastus2 | $1.1761 |
| Standard_D64als_v6 | brazilsouth, eastus2 | $0.7629 |
| Standard_D64as_v5_Promo | eastus2 | $0.4928 |
| Standard_D64as_v6 | westcentralus | $0.6422 |
| Standard_D64as_v7 | brazilsouth | $0.8610 |
| Standard_D64d_v5 | eastus, westcentralus | $0.7633 |
| Standard_D64ds_v4 | brazilsouth, brazilsoutheast, mexicocentral | $1.0644 |
| Standard_D64ds_v7 | eastus2 | $0.9648 |
| Standard_D64ls_v6 | eastus2 | $0.5278 |
| Standard_D64ns_v6 | eastus, westcentralus | $0.9396 |
| Standard_D64plds_v6 | eastus2 | $0.4613 |
| Standard_D64ps_v5 | northcentralus | $0.4553 |
| Standard_D64ps_v6 | brazilsouth, westcentralus | $0.6599 |
| Standard_D64s_v4 | eastus2 | $1.2233 |
| Standard_D8_v4 | westcentralus | $0.0852 |
| Standard_D8_v5 | mexicocentral | $0.0780 |
| Standard_D8a_v4 | brazilsoutheast, westus2 | $0.1471 |
| Standard_D8alds_v6 | brazilsouth, eastus | $0.1129 |
| Standard_D8alds_v7 | westus2 | $0.0704 |
| Standard_D8als_v6 | westus2 | $0.0595 |
| Standard_D8as_v7 | canadacentral, uksouth | $0.0748 |
| Standard_D8d_v4 | brazilsouth | $0.1331 |
| Standard_D8d_v5 | brazilsouth | $0.1331 |
| Standard_D8ds_v4 | eastus2, westus2 | $0.2310 |
| Standard_D8ds_v6 | westcentralus | $0.1105 |
| Standard_D8ds_v7 | uksouth | $0.1454 |
| Standard_D8lds_v5 | northcentralus | $0.0710 |
| Standard_D8ls_v6 | brazilsouth, eastus2 | $0.1016 |
| Standard_D8nds_v6 | westcentralus | $0.1412 |
| Standard_D8ns_v6 | westcentralus | $0.1234 |
| Standard_D8pds_v6 | westus2 | $0.0678 |
| Standard_D8ps_v6 | eastus2 | $0.0519 |
| Standard_D8s_v5 | northcentralus, westus2 | $0.0710 |
| Standard_D8s_v6 | mexicocentral, westus2 | $0.0821 |
| Standard_D96_v5 | westcentralus | $1.0219 |
| Standard_D96a_v4 | eastus2 | $0.8677 |
| Standard_D96ads_v6 | brazilsouth, westus2 | $1.6233 |
| Standard_D96alds_v6 | mexicocentral | $0.9295 |
| Standard_D96alds_v7 | eastus2 | $0.8445 |
| Standard_D96as_v7 | northcentralus | $0.8054 |
| Standard_D96ds_v5 | westcentralus | $1.2029 |
| Standard_D96ds_v7 | uksouth | $1.7458 |
| Standard_D96lds_v5 | eastus | $0.9727 |
| Standard_D96lds_v7 | brazilsouth | $1.8885 |
| Standard_D96ls_v5 | eastus2 | $1.6247 |
| Standard_D96ls_v7 | eastus2 | $1.0391 |
| Standard_D96ns_v6 | brazilsouth, westus2 | $2.4700 |
| Standard_D96pds_v6 | mexicocentral, westus2 | $0.8959 |
| Standard_D96plds_v6 | brazilsouth, eastus2 | $1.0999 |
| Standard_D96pls_v6 | eastus, mexicocentral | $1.0737 |
| Standard_D96s_v6 | westcentralus | $1.0729 |
| Standard_DC128eds_v6 | eastus | $1.6207 |
| Standard_DC16as_v6 | mexicocentral, westus2 | $0.1623 |
| Standard_DC16eds_v5 | eastus2 | $0.1671 |
| Standard_DC1s_v2 | westus2 | $0.0177 |
| Standard_DC1s_v3 | southcentralus | $0.0213 |
| Standard_DC24s_v3 | eastus2 | $0.4258 |
| Standard_DC2ads_v5 | uksouth | $0.0222 |
| Standard_DC2as_v6 | brazilsouth | $0.0296 |
| Standard_DC2s | eastus | $0.0534 |
| Standard_DC2s_v3 | eastus2 | $0.0355 |
| Standard_DC32ads_cc_v5 | westus2 | $0.3045 |
| Standard_DC32as_cc_v5 | uksouth | $0.2957 |
| Standard_DC32as_v5 | eastus | $0.2905 |
| Standard_DC32as_v6 | brazilsouth, brazilsoutheast, westus2 | $0.4736 |
| Standard_DC32eds_v5 | eastus2 | $0.3341 |
| Standard_DC48as_v5 | uksouth | $0.4435 |
| Standard_DC48ds_v3 | uksouth | $1.1620 |
| Standard_DC4ads_v6 | southcentralus | $0.0556 |
| Standard_DC4s_v3 | westus2 | $0.0710 |
| Standard_DC64as_v6 | northcentralus | $0.5906 |
| Standard_DC64es_v6 | eastus | $0.6557 |
| Standard_DC8as_v5 | eastus | $0.0726 |
| Standard_DC8as_v6 | southcentralus | $0.0883 |
| Standard_DC8ds_v3 | eastus, eastus2 | $0.1897 |
| Standard_DC8es_v5 | eastus2 | $0.0710 |
| Standard_DC96as_cc_v5 | uksouth | $0.8870 |
| Standard_DC96as_v6 | westus2 | $0.8859 |
| Standard_DS11 | brazilsouth, eastus2 | $0.0478 |
| Standard_DS12 | brazilsouth, eastus, eastus2 | $0.0957 |
| Standard_DS12-2_v2 | brazilsoutheast, westcentralus | $0.1244 |
| Standard_DS13-2_v2 | brazilsouth, eastus2, westcentralus, westus2 | $0.1914 |
| Standard_DS13-4_v2 | westus2 | $0.1218 |
| Standard_DS13_v2 | westus2 | $0.1218 |
| Standard_DS14 | mexicocentral | $0.3449 |
| Standard_DS14-8_v2 | eastus | $0.3201 |
| Standard_DS2 | brazilsouth, eastus2 | $0.0387 |
| Standard_DS4_v2 | brazilsouth, brazilsoutheast | $0.1395 |
| Standard_E104i_v5 | brazilsouth | $2.1246 |
| Standard_E104is_v5 | brazilsoutheast, westus2 | $2.7620 |
| Standard_E112iads_v5 | brazilsouth | $2.3678 |
| Standard_E112ias_v5 | westcentralus | $2.1851 |
| Standard_E112ibds_v5 | eastus2 | $5.7957 |
| Standard_E112ibs_v5 | eastus | $1.9329 |
| Standard_E128-32as_v7 | westus2 | $1.4098 |
| Standard_E128-64s_v6 | eastus | $1.7561 |
| Standard_E128ads_v7 | uksouth | $1.9964 |
| Standard_E128bs_v6 | brazilsoutheast, westus2 | $5.3629 |
| Standard_E128ns_v6 | mexicocentral | $2.2681 |
| Standard_E16-4as_v7 | mexicocentral, westus2 | $0.1940 |
| Standard_E16-4ds_v4 | eastus | $0.2422 |
| Standard_E16-4s_v3 | eastus | $0.1976 |
| Standard_E16-8as_v6 | westus2 | $0.1763 |
| Standard_E16-8ds_v6 | eastus | $0.2731 |
| Standard_E16-8s_v3 | eastus | $0.1976 |
| Standard_E16-8s_v4 | eastus, eastus2 | $0.2122 |
| Standard_E160ads_v7 | eastus2, westus2 | $2.1466 |
| Standard_E16a_v4 | westus2 | $0.1863 |
| Standard_E16ads_v6 | westus2 | $0.2147 |
| Standard_E16as_v4 | westus2 | $0.1863 |
| Standard_E16bs_v5 | brazilsouth, eastus2 | $0.3541 |
| Standard_E16bs_v6 | brazilsouth | $0.5156 |
| Standard_E16d_v5 | westus2 | $0.2371 |
| Standard_E16ds_v4 | westcentralus | $0.2554 |
| Standard_E16ds_v7 | uksouth | $0.3709 |
| Standard_E16nds_v6 | mexicocentral, westus2 | $0.3177 |
| Standard_E16ns_v6 | brazilsoutheast, westcentralus | $0.6701 |
| Standard_E16pds_v5 | eastus | $0.1949 |
| Standard_E16pds_v6 | westcentralus | $0.2075 |
| Standard_E16ps_v5 | westus2 | $0.1489 |
| Standard_E16ps_v6 | mexicocentral | $0.1499 |
| Standard_E16s_v4 | brazilsoutheast | $0.3862 |
| Standard_E16s_v6 | eastus, westus2 | $0.2194 |
| Standard_E192ibs_v6 | eastus | $3.4237 |
| Standard_E192ids_v6 | brazilsoutheast | $6.6175 |
| Standard_E20_v3 | brazilsouth | $0.4852 |
| Standard_E20_v4 | westus2 | $0.2524 |
| Standard_E20a_v4 | eastus | $0.2652 |
| Standard_E20ads_v6 | eastus, westus2 | $0.3043 |
| Standard_E20d_v4 | eastus2 | $0.6839 |
| Standard_E20d_v5 | southcentralus, westus2 | $0.3193 |
| Standard_E20ds_v4 | brazilsouth | $0.4213 |
| Standard_E20ds_v5 | westus2 | $0.2964 |
| Standard_E20pds_v5 | eastus, westus2 | $0.2435 |
| Standard_E20ps_v5 | westcentralus | $0.2236 |
| Standard_E20s_v4 | eastus | $0.2652 |
| Standard_E20s_v7 | brazilsouth | $0.5136 |
| Standard_E248ds_v7 | eastus | $4.7641 |
| Standard_E2_v3 | brazilsouth | $0.0407 |
| Standard_E2ads_v7 | mexicocentral | $0.0296 |
| Standard_E2as_v4 | westcentralus | $0.0279 |
| Standard_E2bds_v5 | brazilsoutheast, eastus2, mexicocentral | $0.0639 |
| Standard_E2bs_v5 | eastus2, northcentralus | $0.0760 |
| Standard_E2d_v4 | mexicocentral | $0.0292 |
| Standard_E2d_v5 | westus2 | $0.0296 |
| Standard_E2ds_v4 | eastus | $0.0303 |
| Standard_E2ds_v6 | mexicocentral, westus2 | $0.0333 |
| Standard_E2ds_v7 | eastus2, westus2 | $0.0384 |
| Standard_E2nds_v6 | westus2 | $0.0360 |
| Standard_E2pds_v5 | eastus | $0.0243 |
| Standard_E2ps_v5 | uksouth | $0.0171 |
| Standard_E2ps_v6 | eastus, mexicocentral | $0.0389 |
| Standard_E2s_v5 | northcentralus | $0.0233 |
| Standard_E2s_v6 | eastus, westus2 | $0.0274 |
| Standard_E32-16ads_v6 | brazilsoutheast | $0.8930 |
| Standard_E32-16as_v6 | eastus | $0.3997 |
| Standard_E32-16ds_v4 | brazilsoutheast | $0.8830 |
| Standard_E32-16s_v3 | mexicocentral | $0.4099 |
| Standard_E32-16s_v5 | westus2 | $0.4074 |
| Standard_E32-8ads_v6 | eastus2 | $0.4293 |
| Standard_E32-8ads_v7 | brazilsouth, eastus2, westus2 | $0.6884 |
| Standard_E32-8as_v4 | mexicocentral | $0.4099 |
| Standard_E32-8ds_v5 | brazilsouth | $0.6742 |
| Standard_E32-8s_v4 | eastus2 | $0.9888 |
| Standard_E32_v4 | eastus2 | $0.9888 |
| Standard_E32ads_v7 | eastus | $0.4335 |
| Standard_E32as_v6 | westcentralus | $0.4223 |
| Standard_E32as_v7 | westus3 | $0.3524 |
| Standard_E32d_v5 | westus2 | $0.4742 |
| Standard_E32ds_v4 | westus2 | $0.4742 |
| Standard_E32ds_v6 | eastus2 | $0.4834 |
| Standard_E32ns_v6 | brazilsouth | $1.0310 |
| Standard_E32pds_v5 | westcentralus | $0.4088 |
| Standard_E32pds_v6 | eastus2 | $0.3886 |
| Standard_E32s_v4 | eastus | $0.4244 |
| Standard_E32s_v5 | southcentralus | $0.4470 |
| Standard_E32s_v7 | uksouth | $0.6194 |
| Standard_E372is_v7 | eastus2 | $6.5661 |
| Standard_E4-2ads_v6 | mexicocentral | $0.0590 |
| Standard_E4-2as_v6 | eastus2 | $0.0440 |
| Standard_E4-2as_v7 | eastus | $0.0444 |
| Standard_E4-2ds_v4 | eastus | $0.0605 |
| Standard_E4-2ds_v6 | westcentralus | $0.0724 |
| Standard_E4-2s_v6 | westus2 | $0.0490 |
| Standard_E48_v4 | eastus2, westus2 | $1.4833 |
| Standard_E48a_v4 | mexicocentral, uksouth | $0.6146 |
| Standard_E48ads_v5 | westus2 | $0.5810 |
| Standard_E48ads_v6 | westcentralus | $0.7708 |
| Standard_E48as_v7 | eastus | $0.5339 |
| Standard_E48bds_v5 | eastus2 | $2.2581 |
| Standard_E48bs_v5 | westcentralus | $0.7930 |
| Standard_E48d_v5 | eastus2 | $1.6413 |
| Standard_E48ds_v4 | westcentralus | $0.7664 |
| Standard_E48nds_v6 | westus2 | $0.8665 |
| Standard_E48ns_v6 | westus2 | $0.7732 |
| Standard_E48s_v3 | northcentralus | $0.5588 |
| Standard_E48s_v4 | westcentralus | $0.6706 |
| Standard_E48s_v5 | eastus2 | $1.4833 |
| Standard_E48s_v6 | westus2 | $0.5867 |
| Standard_E48s_v7 | eastus2 | $0.7702 |
| Standard_E4a_v4 | eastus2 | $0.0466 |
| Standard_E4ads_v7 | eastus | $0.0541 |
| Standard_E4as_v7 | brazilsouth | $0.0708 |
| Standard_E4bds_v5 | westus2 | $0.0703 |
| Standard_E4bs_v5 | westus2 | $0.0653 |
| Standard_E4d_v4 | eastus2 | $0.1368 |
| Standard_E4d_v5 | brazilsoutheast | $0.1096 |
| Standard_E4ds_v4 | mexicocentral, westcentralus | $0.0586 |
| Standard_E4ds_v5 | eastus | $0.0606 |
| Standard_E4ds_v6 | westcentralus | $0.0724 |
| Standard_E4ds_v7 | eastus, mexicocentral, northcentralus | $0.0769 |
| Standard_E4ns_v6 | brazilsouth | $0.1288 |
| Standard_E4pds_v5 | westcentralus | $0.0510 |
| Standard_E4ps_v6 | westcentralus | $0.0408 |
| Standard_E4s_v5 | brazilsoutheast, eastus2, westus2 | $0.0966 |
| Standard_E4s_v7 | eastus2 | $0.0641 |
| Standard_E64-16as_v4 | eastus | $0.8487 |
| Standard_E64-16as_v6 | eastus2 | $0.7048 |
| Standard_E64-16ds_v4 | eastus2 | $2.1883 |
| Standard_E64-16ds_v5 | eastus | $0.9691 |
| Standard_E64-16ds_v6 | eastus2 | $0.9671 |
| Standard_E64-16s_v3 | mexicocentral | $0.7377 |
| Standard_E64-32as_v7 | canadacentral | $0.7865 |
| Standard_E64-32ds_v4 | eastus | $0.9686 |
| Standard_E64-32ds_v5 | brazilsouth | $1.3483 |
| Standard_E64-32s_v3 | westus2 | $0.6706 |
| Standard_E64-32s_v5 | westus2 | $0.8149 |
| Standard_E64ads_v5 | westcentralus | $1.3158 |
| Standard_E64ads_v7 | mexicocentral, westus2 | $0.9438 |
| Standard_E64as_v7 | uksouth | $0.8196 |
| Standard_E64bs_v6 | westcentralus | $1.2371 |
| Standard_E64d_v5 | westus2 | $0.9483 |
| Standard_E64ds_v4 | brazilsoutheast | $1.7659 |
| Standard_E64ds_v5 | brazilsouth, brazilsoutheast, canadacentral | $1.3483 |
| Standard_E64is_v3 | eastus2 | $1.7376 |
| Standard_E64nds_v6 | canadacentral, eastus2, westcentralus | $1.3866 |
| Standard_E64pds_v6 | westcentralus | $0.8303 |
| Standard_E64s_v4 | brazilsouth | $1.1886 |
| Standard_E64s_v5 | mexicocentral, westus2 | $0.8196 |
| Standard_E64s_v6 | eastus2 | $0.7824 |
| Standard_E64s_v7 | westus2 | $1.0269 |
| Standard_E8-2as_v6 | brazilsoutheast | $0.1986 |
| Standard_E8-2ds_v4 | westus2 | $0.1185 |
| Standard_E8-2ds_v5 | canadacentral | $0.2662 |
| Standard_E8-2s_v4 | mexicocentral, westus2 | $0.1024 |
| Standard_E8-2s_v5 | eastus2 | $0.2472 |
| Standard_E8-4ads_v7 | eastus2 | $0.1074 |
| Standard_E8-4as_v4 | eastus2 | $0.0931 |
| Standard_E8-4as_v7 | canadacentral, westus2 | $0.0983 |
| Standard_E8-4ds_v5 | eastus2 | $0.2735 |
| Standard_E8-4s_v3 | uksouth | $0.0907 |
| Standard_E8-4s_v5 | eastus | $0.1061 |
| Standard_E80ids_v4 | eastus2 | $2.7354 |
| Standard_E8_v3 | westcentralus, westus2 | $0.1077 |
| Standard_E8_v5 | westus2 | $0.1019 |
| Standard_E8ads_v6 | westus2 | $0.1074 |
| Standard_E8ads_v7 | westus2 | $0.1074 |
| Standard_E8as_v4 | westus2 | $0.0931 |
| Standard_E8bds_v6 | westus2 | $0.1445 |
| Standard_E8d_v5 | westus2 | $0.1185 |
| Standard_E8ds_v4 | westus2 | $0.1185 |
| Standard_E8ds_v6 | northeurope, westus2 | $0.1353 |
| Standard_E8ds_v7 | brazilsouth, westus2 | $0.2460 |
| Standard_E8s_v4 | westus2 | $0.1010 |
| Standard_E8s_v5 | westcentralus | $0.1118 |
| Standard_E8s_v6 | mexicocentral | $0.1076 |
| Standard_E96-24ads_v7 | westus2 | $1.2881 |
| Standard_E96-24as_v4 | westus2 | $1.1177 |
| Standard_E96-24as_v6 | mexicocentral | $1.1639 |
| Standard_E96-24as_v7 | brazilsouth, eastus | $1.6978 |
| Standard_E96-24ds_v5 | canadacentral | $3.1944 |
| Standard_E96-48ads_v6 | brazilsoutheast | $2.6789 |
| Standard_E96-48s_v5 | eastus2, westcentralus | $2.9665 |
| Standard_E96_v5 | brazilsoutheast, uksouth | $2.3178 |
| Standard_E96a_v4 | eastus2 | $1.1177 |
| Standard_E96ads_v7 | eastus2, mexicocentral | $1.2881 |
| Standard_E96as_v7 | mexicocentral | $1.1639 |
| Standard_E96bds_v5 | westcentralus | $1.7776 |
| Standard_E96bds_v6 | eastus2, uksouth | $1.7332 |
| Standard_E96bs_v5 | mexicocentral | $1.4538 |
| Standard_E96bs_v6 | eastus2, westcentralus | $1.5470 |
| Standard_E96ias_v5 | eastus2 | $1.1025 |
| Standard_E96pds_v6 | westus2 | $1.0378 |
| Standard_E96s_v5 | eastus2, westcentralus | $2.9665 |
| Standard_E96s_v6 | westus2 | $1.1735 |
| Standard_E96s_v7 | brazilsouth | $2.4647 |
| Standard_EC128es_v5 | eastus2 | $1.4902 |
| Standard_EC16ads_cc_v5 | eastus2 | $0.1937 |
| Standard_EC16ads_v6 | eastus | $0.2697 |
| Standard_EC16as_v6 | eastus2 | $0.1939 |
| Standard_EC20ads_cc_v5 | eastus2 | $0.2421 |
| Standard_EC48as_v5 | eastus2 | $0.5012 |
| Standard_EC48as_v6 | brazilsoutheast | $1.2108 |
| Standard_EC4ads_v6 | westus2 | $0.0590 |
| Standard_EC4as_v6 | westus2 | $0.0484 |
| Standard_EC64ads_cc_v5 | eastus2 | $0.7747 |
| Standard_EC64ads_v6 | mexicocentral | $1.0382 |
| Standard_EC64eds_v6 | eastus | $1.0637 |
| Standard_EC8ads_cc_v5 | eastus2 | $0.0968 |
| Standard_EC8ads_v6 | canadacentral | $0.1318 |
| Standard_EC8as_v6 | eastus2, westcentralus, westus2 | $0.0968 |
| Standard_EC96ads_v6 | eastus2, westus2 | $1.4169 |
| Standard_EC96eds_v5 | eastus2 | $1.2773 |
| Standard_EC96ias_v5 | eastus2 | $1.1025 |
| Standard_F16 | eastus2 | $0.4897 |
| Standard_F16-4amds_v7 | eastus2, westus2 | $0.3223 |
| Standard_F16-8amds_v7 | brazilsouth, eastus, mexicocentral | $0.5169 |
| Standard_F16als_v7 | westus2 | $0.1789 |
| Standard_F16as_v6 | brazilsouth | $0.3241 |
| Standard_F16s | brazilsouth | $0.2337 |
| Standard_F1ads_v7 | brazilsouth | $0.0255 |
| Standard_F1alds_v7 | westus2 | $0.0132 |
| Standard_F1als_v6 | eastus | $0.0128 |
| Standard_F1amds_v7 | eastus2 | $0.0201 |
| Standard_F1ams_v6 | westus2 | $0.0165 |
| Standard_F1as_v6 | eastus2 | $0.0126 |
| Standard_F1s | eastus, westus2 | $0.0107 |
| Standard_F2 | westus2 | $0.0202 |
| Standard_F2-1amds_v7 | westus2 | $0.0403 |
| Standard_F2alds_v7 | eastus2 | $0.0264 |
| Standard_F2als_v6 | westcentralus | $0.0268 |
| Standard_F2ams_v7 | westus2 | $0.0331 |
| Standard_F32-8amds_v7 | eastus2 | $0.6448 |
| Standard_F32als_v6 | brazilsoutheast, eastus2 | $0.7440 |
| Standard_F32als_v7 | westus2 | $0.3578 |
| Standard_F32ams_v6 | uksouth | $0.6150 |
| Standard_F32ams_v7 | westus2 | $0.5293 |
| Standard_F32as_v6 | westus2 | $0.4042 |
| Standard_F32s_v2 | eastus2, westus2 | $0.6726 |
| Standard_F4-1amds_v7 | mexicocentral | $0.0885 |
| Standard_F4-2ams_v7 | brazilsouth | $0.1061 |
| Standard_F48as_v6 | eastus2 | $0.6058 |
| Standard_F4ads_v7 | eastus | $0.0665 |
| Standard_F4als_v6 | mexicocentral | $0.0492 |
| Standard_F4amds_v7 | eastus2 | $0.0806 |
| Standard_F4ams_v6 | mexicocentral | $0.0728 |
| Standard_F4as_v6 | eastus2 | $0.0505 |
| Standard_F64-16amds_v7 | westus2 | $1.2893 |
| Standard_F64-16ams_v7 | westus2 | $1.0585 |
| Standard_F64-32ams_v7 | eastus | $1.1112 |
| Standard_F64als_v6 | eastus2, westus2 | $0.7155 |
| Standard_F64als_v7 | westus2 | $0.7155 |
| Standard_F64ams_v6 | brazilsoutheast | $2.2023 |
| Standard_F64as_v7 | westus2 | $0.8078 |
| Standard_F64s_v2 | eastus | $0.5953 |
| Standard_F72as_v6 | westus2 | $0.9093 |
| Standard_F8 | eastus2, westus2 | $0.2449 |
| Standard_F8-4amds_v7 | brazilsouth, westus2 | $0.2584 |
| Standard_F80als_v7 | brazilsouth, westus2 | $1.4340 |
| Standard_F80ams_v7 | eastus2 | $1.3232 |
| Standard_F80as_v7 | westus2 | $1.0097 |
| Standard_F8ads_v7 | brazilsouth | $0.2035 |
| Standard_F8alds_v7 | westus2 | $0.1057 |
| Standard_F8als_v7 | eastus2 | $0.0894 |
| Standard_F8ams_v6 | eastus, eastus2 | $0.1509 |
| Standard_F8ams_v7 | eastus2 | $0.1323 |
| Standard_F8as_v6 | westus2 | $0.1010 |
| Standard_F8s_v2 | eastus2 | $0.1680 |
| Standard_FX12-6mds_v2 | canadacentral | $0.3337 |
| Standard_FX12mds | eastus2 | $0.2062 |
| Standard_FX16-4mds_v2 | westus2 | $0.4045 |
| Standard_FX16-4ms_v2 | westus2 | $0.3273 |
| Standard_FX24-12ms_v2 | eastus2 | $0.4908 |
| Standard_FX24-6mds_v2 | westus2 | $0.6067 |
| Standard_FX24mds_v2 | eastus | $0.6889 |
| Standard_FX24ms_v2 | westus2 | $0.4910 |
| Standard_FX2mds_v2 | brazilsouth | $0.0987 |
| Standard_FX2ms_v2 | brazilsouth | $0.0798 |
| Standard_FX32-16mds_v2 | westus2 | $0.8091 |
| Standard_FX32-16ms_v2 | westus2 | $0.6546 |
| Standard_FX32-8mds_v2 | eastus2 | $0.8089 |
| Standard_FX32-8ms_v2 | eastus2, westus2 | $0.6544 |
| Standard_FX32ms_v2 | westus2 | $0.6546 |
| Standard_FX4-2ms_v2 | westus2 | $0.0819 |
| Standard_FX48-12mds_v2 | eastus | $1.3780 |
| Standard_FX48-12ms_v2 | westus2 | $0.9818 |
| Standard_FX48-24ms_v2 | eastus2 | $0.9817 |
| Standard_FX48mds | eastus2 | $0.8249 |
| Standard_FX48ms_v2 | westus2 | $0.9818 |
| Standard_FX4mds | westus2 | $0.0687 |
| Standard_FX4mds_v2 | westcentralus | $0.1212 |
| Standard_FX4ms_v2 | westus2 | $0.0819 |
| Standard_FX64-16ms_v2 | uksouth | $1.6495 |
| Standard_FX64-32ms_v2 | westus2 | $1.3091 |
| Standard_FX8-2mds_v2 | westus2 | $0.2022 |
| Standard_FX8-4ms_v2 | westcentralus | $0.1963 |
| Standard_FX8ms_v2 | brazilsouth, eastus2 | $0.3191 |
| Standard_FX96-24mds_v2 | eastus2 | $2.4264 |
| Standard_FX96-48mds_v2 | westus2 | $2.4270 |
| Standard_G5 | eastus2 | $2.6870 |
| Standard_GS5-16 | westus2 | $1.4451 |
| Standard_H16 | westus2 | $0.2940 |
| Standard_H16mr | northcentralus | $0.4925 |
| Standard_HB120-16rs_v2 | westus2 | $0.6653 |
| Standard_HB120-32rs_v2 | brazilsouth | $1.3306 |
| Standard_HB120-96rs_v2 | northcentralus | $0.7983 |
| Standard_HB120rs_v2 | brazilsouth | $1.3306 |
| Standard_HB368-144rs_v5 | eastus | $3.6590 |
| Standard_HC44-16rs | westus2 | $0.5854 |
| Standard_HC44-32rs | eastus | $0.5854 |
| Standard_HX176-96rs | eastus | $1.5967 |
| Standard_L128aos_v5 | eastus | $3.9014 |
| Standard_L16as_v3 | brazilsoutheast | $0.4786 |
| Standard_L16as_v4 | mexicocentral | $0.2790 |
| Standard_L16s_v2 | eastus2 | $0.2306 |
| Standard_L16s_v3 | brazilsouth | $0.4110 |
| Standard_L2aos_v4 | eastus2, westus2 | $0.0418 |
| Standard_L2as_v4 | brazilsouth | $0.0506 |
| Standard_L2as_v5 | westus2 | $0.0408 |
| Standard_L32as_v3 | eastus2, westus2 | $0.4613 |
| Standard_L32as_v5 | westus2 | $0.6017 |
| Standard_L32s_v3 | northcentralus | $0.5160 |
| Standard_L32s_v4 | brazilsouth | $0.8220 |
| Standard_L4aos_v4 | westus2 | $0.0835 |
| Standard_L4as_v4 | eastus2 | $0.0636 |
| Standard_L4as_v5 | westus2 | $0.0814 |
| Standard_L4s | westus2 | $0.0577 |
| Standard_L64as_v3 | southcentralus | $1.1070 |
| Standard_L64as_v4 | eastus | $1.1536 |
| Standard_L64s_v3 | eastus | $1.1721 |
| Standard_L64s_v4 | westus2 | $1.0319 |
| Standard_L80as_v5 | westus2 | $1.6280 |
| Standard_L80s_v3 | eastus2 | $3.9416 |
| Standard_L8aos_v4 | westus2 | $0.1671 |
| Standard_L8as_v3 | eastus2 | $0.1153 |
| Standard_L8s | eastus2 | $0.2357 |
| Standard_L8s_v2 | eastus | $0.1445 |
| Standard_L8s_v3 | brazilsouth | $0.2055 |
| Standard_L96as_v4 | brazilsouth, eastus | $2.4296 |
| Standard_M128-32ms | eastus2 | $4.9319 |
| Standard_M128-64bs_v3 | eastus2 | $2.0624 |
| Standard_M128bs_2_v4 | eastus | $0.0000 |
| Standard_M128bs_v3 | westus2 | $2.0624 |
| Standard_M128dms_v2 | brazilsouth | $9.6181 |
| Standard_M128ds_v4 | eastus2 | $5.2548 |
| Standard_M128lds_v4 | westus2 | $3.3287 |
| Standard_M128ms_v2 | westus2 | $4.8545 |
| Standard_M16-8ms | brazilsouth | $1.1073 |
| Standard_M160ds_v4 | eastus | $6.5685 |
| Standard_M16bds_v4 | eastus | $0.0000 |
| Standard_M16lds_v4 | eastus | $0.4379 |
| Standard_M16s_v3 | westus2 | $0.3280 |
| Standard_M176-88bds_v3 | westus2 | $3.2800 |
| Standard_M176bs_v3 | westus2 | $2.8348 |
| Standard_M176ds-88_3_v3 | eastus2 | $3.8000 |
| Standard_M176ds_3_v3 | westus2 | $3.5112 |
| Standard_M176ds_v3 | westus2 | $3.1320 |
| Standard_M176s-88_4_v3 | eastus | $4.9820 |
| Standard_M176s_4_v3 | westus2 | $4.6034 |
| Standard_M192ids_v2 | westus2 | $2.9627 |
| Standard_M192lds_v4 | eastus | $5.2548 |
| Standard_M208-52ms_v2 | brazilsoutheast | $21.4386 |
| Standard_M208-52s_v2 | eastus2 | $4.1229 |
| Standard_M208ms_v2 | westus2 | $8.2458 |
| Standard_M24ds-12_v3 | eastus | $0.6520 |
| Standard_M24ds_v3 | westus2 | $0.6024 |
| Standard_M24lds_v4 | eastus2 | $0.6568 |
| Standard_M24ls_v4 | westus2 | $0.6256 |
| Standard_M32-16ms | eastus2 | $1.1358 |
| Standard_M32bs_v3 | brazilsouth | $1.0312 |
| Standard_M32ls | eastus | $0.5309 |
| Standard_M32ms | westus2 | $1.1358 |
| Standard_M32ms_v4 | westus2 | $1.6682 |
| Standard_M32s_v4 | eastus2 | $1.2511 |
| Standard_M32ts_v3 | westus2 | $0.6560 |
| Standard_M416-104ms_v2 | eastus2, westus2 | $18.3229 |
| Standard_M416-208ms_v2 | eastus | $18.3229 |
| Standard_M416ds-104_8_v3 | westus2 | $15.4844 |
| Standard_M416ds-208_6_v3 | eastus | $11.6133 |
| Standard_M416ds_12_v3 | westus2 | $21.4464 |
| Standard_M416is_v2 | eastus | $9.1624 |
| Standard_M416s-104_6_v3 | eastus2 | $11.0603 |
| Standard_M416s-104_8_v3 | brazilsouth | $23.5250 |
| Standard_M416s_10_v3 | eastus | $18.4852 |
| Standard_M416s_8_v2 | eastus, westus2 | $12.2171 |
| Standard_M416s_9_v3 | westus2 | $16.6253 |
| Standard_M48ds-12_1_v3 | eastus | $1.3240 |
| Standard_M48ds-24_1_v3 | westus2 | $1.3240 |
| Standard_M48ds_1_v3 | westus2 | $1.2234 |
| Standard_M48ds_v4 | westus2 | $1.9705 |
| Standard_M48lds_v4 | eastus, westus2 | $1.3137 |
| Standard_M48s-12_1_v3 | eastus | $1.2460 |
| Standard_M624s_12_v3 | eastus | $20.4394 |
| Standard_M64 | mexicocentral | $1.3557 |
| Standard_M64-16ms | westus2 | $1.9103 |
| Standard_M64-32ms | westus2 | $1.9103 |
| Standard_M64dms_v2 | eastus2 | $1.9108 |
| Standard_M64ds_v2 | uksouth | $1.5529 |
| Standard_M64ls | brazilsoutheast | $1.8963 |
| Standard_M64m | mexicocentral | $2.1517 |
| Standard_M64ms | brazilsoutheast, westus2 | $4.2983 |
| Standard_M64s | eastus | $1.2324 |
| Standard_M64s_v2 | eastus | $1.1936 |
| Standard_M832ds-208_12_v3 | westus2 | $23.3012 |
| Standard_M832ds_12_v3 | brazilsouth | $34.3460 |
| Standard_M832ids-208_16_v3 | eastus2 | $34.4098 |
| Standard_M832ids-416_16_v3 | eastus2 | $34.4098 |
| Standard_M832ids_16_v3 | westus2 | $31.7946 |
| Standard_M896ixds_24_v3 | westus2 | $57.5999 |
| Standard_M896ixds_32_v3 | eastus2 | $63.2016 |
| Standard_M8ms | brazilsoutheast | $0.6389 |
| Standard_M96-48bds_2_v3 | westus2 | $7.0181 |
| Standard_M96bs_v3 | eastus | $1.5468 |
| Standard_M96ds-48_1_v3 | eastus2 | $1.9860 |
| Standard_M96ds_2_v3 | eastus, westus2 | $2.4449 |
| Standard_M96s-24_1_v3 | westus2 | $1.8680 |
| Standard_M96s-24_2_v3 | westus2 | $2.4900 |
| Standard_M96s-48_2_v3 | eastus2 | $2.4900 |
| Standard_M96s_1_v3 | canadacentral | $1.8987 |
| Standard_M96s_2_v3 | westus2 | $2.3008 |
| Standard_M96s_v3 | westus2 | $1.9660 |
| Standard_M96s_v4 | eastus, westus2 | $3.7534 |
| Standard_NC128lds_xl_RTXPRO6000BSE_v6 | eastus2, westus2 | $1.0164 |
| Standard_NC132ds_xl_RTXPRO6000BSE_v6 | eastus | $1.1790 |
| Standard_NC144lds_xl_RTXPRO6000BSE_v6 | eastus2 | $1.0164 |
| Standard_NC16as_T4_v3 | brazilsouth, eastus2 | $0.5802 |
| Standard_NC24s_v3 | eastus | $2.2620 |
| Standard_NC264lds_xl_RTXPRO6000BSE_v6 | eastus2 | $2.0328 |
| Standard_NC320lds_xl_RTXPRO6000BSE_v6 | westus2 | $2.6426 |
| Standard_NC36ds_xl_RTXPRO6000BSE_v6 | eastus2, westus2 | $0.2665 |
| Standard_NC40ads_H100_v5 | westus2 | $4.0205 |
| Standard_NC64ds_xl_RTXPRO6000BSE_v6 | westus2 | $0.5660 |
| Standard_NC6s_v3 | westus2 | $0.5655 |
| Standard_NC80adis_H100_v5 | westus2 | $8.0410 |
| Standard_ND128isr_NDR_GB200_v6 | northcentralus | $129.8000 |
| Standard_ND40rs_v2 | eastus | $4.0715 |
| Standard_ND96amsr_A100_v4 | brazilsouth | $16.8831 |
| Standard_ND96is_MI300X_v5 | uksouth | $11.0880 |
| Standard_NG16ads_V620_v1 | eastus2 | $0.2347 |
| Standard_NV12ads_A10_v5 | eastus2 | $0.1678 |
| Standard_NV12s_v3 | eastus2 | $0.2107 |
| Standard_NV32as_v4 | eastus | $0.3445 |
| Standard_NV36ads_A10_v5 | westus2 | $0.5914 |
| Standard_NV6ads_A10_v5 | eastus2 | $0.0839 |
| Standard_PB24s | westus2 | $0.6135 |
| Standard_PB6s | brazilsoutheast | $0.1534 |

## Per-Region Mismatch Counts

| Region | On-Demand Only | Spot Only | Total SKUs |
|--------|---------------|-----------|------------|
| australiacentral | 52 | 0 | 899 |
| australiacentral2 | 46 | 0 | 850 |
| australiaeast | 77 | 0 | 1536 |
| australiasoutheast | 55 | 0 | 1025 |
| austriaeast | 38 | 0 | 792 |
| brazilsouth | 157 | 76 | 1310 |
| brazilsoutheast | 94 | 40 | 840 |
| canadacentral | 74 | 15 | 1441 |
| canadaeast | 53 | 0 | 962 |
| centralindia | 75 | 0 | 1448 |
| centralus | 66 | 0 | 1566 |
| chilecentral | 37 | 0 | 671 |
| eastasia | 53 | 0 | 1315 |
| eastus | 155 | 115 | 1650 |
| eastus2 | 242 | 174 | 1488 |
| francecentral | 73 | 0 | 1282 |
| francesouth | 50 | 0 | 831 |
| germanynorth | 42 | 0 | 915 |
| germanywestcentral | 62 | 0 | 1536 |
| indonesiacentral | 39 | 0 | 969 |
| israelcentral | 54 | 0 | 844 |
| italynorth | 58 | 0 | 1296 |
| japaneast | 71 | 0 | 1510 |
| japanwest | 60 | 0 | 1229 |
| jioindiacentral | 1013 | 0 | 1013 |
| jioindiawest | 1193 | 0 | 1193 |
| koreacentral | 58 | 0 | 1391 |
| koreasouth | 52 | 0 | 866 |
| malaysiawest | 41 | 0 | 798 |
| mexicocentral | 93 | 52 | 1107 |
| newzealandnorth | 35 | 0 | 891 |
| northcentralus | 89 | 17 | 1358 |
| northeurope | 82 | 1 | 1650 |
| norwayeast | 48 | 0 | 1168 |
| norwaywest | 45 | 0 | 815 |
| polandcentral | 61 | 0 | 1235 |
| qatarcentral | 51 | 0 | 926 |
| southafricanorth | 63 | 0 | 1287 |
| southafricawest | 44 | 0 | 929 |
| southcentralus | 84 | 9 | 1584 |
| southeastasia | 73 | 0 | 1506 |
| southindia | 54 | 0 | 1082 |
| spaincentral | 50 | 0 | 1231 |
| swedencentral | 57 | 0 | 1461 |
| swedensouth | 41 | 0 | 757 |
| switzerlandnorth | 60 | 0 | 1385 |
| switzerlandwest | 44 | 0 | 891 |
| uaecentral | 50 | 0 | 912 |
| uaenorth | 61 | 0 | 1348 |
| uksouth | 109 | 25 | 1516 |
| ukwest | 54 | 0 | 1000 |
| westcentralus | 125 | 65 | 899 |
| westeurope | 86 | 0 | 1769 |
| westindia | 40 | 0 | 806 |
| westus | 77 | 0 | 1578 |
| westus2 | 295 | 228 | 1319 |
| westus3 | 64 | 1 | 1517 |
