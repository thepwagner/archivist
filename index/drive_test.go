package index_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thepwagner/archivist/index"
)

func TestParseSmartctl_Linux(t *testing.T) {
	d, err := index.ParseSmartctl(smartctlLinux)
	require.NoError(t, err)
	assert.Equal(t, "ST10000VN0004-1ZD101", d.GetModelNumber())
	assert.Equal(t, "ZA2APHBL", d.GetSerialNumber())
	assert.Equal(t, uint64(10000831348736), d.GetCapacity())
	assert.Equal(t, uint64(1), d.GetPowerCycleCount())
	assert.Equal(t, uint64(789), d.GetPowerOnHours())
}

func TestParseSmartctl_Darwin(t *testing.T) {
	d, err := index.ParseSmartctl(smartctlDarwin)
	require.NoError(t, err)
	assert.Equal(t, "APPLE SSD SM0256L", d.GetModelNumber())
	assert.Equal(t, "C0271340010HCGH1M", d.GetSerialNumber())
	assert.Equal(t, uint64(0), d.GetCapacity())
	assert.Equal(t, uint64(24149), d.GetPowerCycleCount())
	assert.Equal(t, uint64(461), d.GetPowerOnHours())
}

const smartctlLinux = `smartctl 6.5 (build date May 10 2019) [x86_64-linux-3.10.105] (local build)
Copyright (C) 2002-16, Bruce Allen, Christian Franke, www.smartmontools.org

=== START OF INFORMATION SECTION ===
Model Family:     Seagate IronWolf
Device Model:     ST10000VN0004-1ZD101
Serial Number:    ZA2APHBL
LU WWN Device Id: 5 000c50 0b5318536
Firmware Version: SC60
User Capacity:    10,000,831,348,736 bytes [10.0 TB]
Sector Sizes:     512 bytes logical, 4096 bytes physical
Rotation Rate:    7200 rpm
Form Factor:      3.5 inches
Device is:        In smartctl database [for details use: -P show]
ATA Version is:   ACS-3 T13/2161-D revision 5
SATA Version is:  SATA 3.1, 6.0 Gb/s (current: 6.0 Gb/s)
Local Time is:    Mon Dec 30 14:46:58 2019 -05
SMART support is: Available - device has SMART capability.
SMART support is: Enabled

=== START OF READ SMART DATA SECTION ===
SMART overall-health self-assessment test result: PASSED

General SMART Values:
Offline data collection status:  (0x82)	Offline data collection activity
					was completed without error.
					Auto Offline Data Collection: Enabled.
Self-test execution status:      (   0)	The previous self-test routine completed
					without error or no self-test has ever
					been run.
Total time to complete Offline
data collection: 		(  567) seconds.
Offline data collection
capabilities: 			 (0x7b) SMART execute Offline immediate.
					Auto Offline data collection on/off support.
					Suspend Offline collection upon new
					command.
					Offline surface scan supported.
					Self-test supported.
					Conveyance Self-test supported.
					Selective Self-test supported.
SMART capabilities:            (0x0003)	Saves SMART data before entering
					power-saving mode.
					Supports SMART auto save timer.
Error logging capability:        (0x01)	Error logging supported.
					General Purpose Logging supported.
Short self-test routine
recommended polling time: 	 (   1) minutes.
Extended self-test routine
recommended polling time: 	 ( 847) minutes.
Conveyance self-test routine
recommended polling time: 	 (   2) minutes.
SCT capabilities: 	       (0x50bd)	SCT Status supported.
					SCT Error Recovery Control supported.
					SCT Feature Control supported.
					SCT Data Table supported.

SMART Attributes Data Structure revision number: 10
Vendor Specific SMART Attributes with Thresholds:
ID# ATTRIBUTE_NAME                                                   FLAG     VALUE WORST THRESH TYPE      UPDATED  WHEN_FAILED RAW_VALUE
  1 Raw_Read_Error_Rate                                              0x000f   080   064   044    Pre-fail  Always       -       89930826
  3 Spin_Up_Time                                                     0x0003   099   099   000    Pre-fail  Always       -       0
  4 Start_Stop_Count                                                 0x0032   100   100   020    Old_age   Always       -       1
  5 Reallocated_Sector_Ct                                            0x0033   100   100   010    Pre-fail  Always       -       0
  7 Seek_Error_Rate                                                  0x000f   083   060   045    Pre-fail  Always       -       178647929
  9 Power_On_Hours                                                   0x0032   100   100   000    Old_age   Always       -       789 (201 129 0)
 10 Spin_Retry_Count                                                 0x0013   100   100   097    Pre-fail  Always       -       0
 12 Power_Cycle_Count                                                0x0032   100   100   020    Old_age   Always       -       1
184 End-to-End_Error                                                 0x0032   100   100   099    Old_age   Always       -       0
187 Reported_Uncorrect                                               0x0032   100   100   000    Old_age   Always       -       0
188 Command_Timeout                                                  0x0032   100   100   000    Old_age   Always       -       0
189 High_Fly_Writes                                                  0x003a   093   093   000    Old_age   Always       -       7
190 Airflow_Temperature_Cel                                          0x0022   071   066   040    Old_age   Always       -       29 (Min/Max 19/34)
191 G-Sense_Error_Rate                                               0x0032   100   100   000    Old_age   Always       -       1806
192 Power-Off_Retract_Count                                          0x0032   100   100   000    Old_age   Always       -       0
193 Load_Cycle_Count                                                 0x0032   100   100   000    Old_age   Always       -       4
194 Temperature_Celsius                                              0x0022   029   040   000    Old_age   Always       -       29 (0 19 0 0 0)
195 Hardware_ECC_Recovered                                           0x001a   009   001   000    Old_age   Always       -       89930826
197 Current_Pending_Sector                                           0x0012   100   100   000    Old_age   Always       -       0
198 Offline_Uncorrectable                                            0x0010   100   100   000    Old_age   Offline      -       0
199 UDMA_CRC_Error_Count                                             0x003e   200   200   000    Old_age   Always       -       0
200 Multi_Zone_Error_Rate                                            0x0023   100   100   001    Pre-fail  Always       -       0
240 Head_Flying_Hours                                                0x0000   100   253   000    Old_age   Offline      -       789 (169 231 0)
241 Total_LBAs_Written                                               0x0000   100   253   000    Old_age   Offline      -       22300906612
242 Total_LBAs_Read                                                  0x0000   100   253   000    Old_age   Offline      -       1682633630

SMART Error Log Version: 1
No Errors Logged

SMART Self-test log structure revision number 1
Num  Test_Description    Status                  Remaining  LifeTime(hours)  LBA_of_first_error
# 1  Short offline       Completed without error       00%       630         -
# 2  Short offline       Completed without error       00%         0         -

SMART Selective self-test log data structure revision number 1
 SPAN  MIN_LBA  MAX_LBA  CURRENT_TEST_STATUS
    1        0        0  Not_testing
    2        0        0  Not_testing
    3        0        0  Not_testing
    4        0        0  Not_testing
    5        0        0  Not_testing
Selective self-test flags (0x0):
  After scanning selected spans, do NOT read-scan remainder of disk.
If Selective self-test is pending on power-up, resume after 0 minute delay.
`

const smartctlDarwin = `smartctl 7.0 2018-12-30 r4883 [Darwin 19.2.0 x86_64] (local build)
Copyright (C) 2002-18, Bruce Allen, Christian Franke, www.smartmontools.org

=== START OF INFORMATION SECTION ===
Model Number:                       APPLE SSD SM0256L
Serial Number:                      C0271340010HCGH1M
Firmware Version:                   CXS6AA0Q
PCI Vendor/Subsystem ID:            0x144d
IEEE OUI Identifier:                0x002538
Controller ID:                      2
Number of Namespaces:               1
Local Time is:                      Tue Dec 31 09:58:01 2019 EST
Firmware Updates (0x06):            3 Slots
Optional Admin Commands (0x0006):   Format Frmw_DL
Optional NVM Commands (0x001f):     Comp Wr_Unc DS_Mngmt Wr_Zero Sav/Sel_Feat
Maximum Data Transfer Size:         256 Pages

Supported Power States
St Op     Max   Active     Idle   RL RT WL WT  Ent_Lat  Ex_Lat
 0 +     6.00W       -        -    0  0  0  0        5       5
 1 -   0.0400W       -        -    1  1  1  1      210    1200
 2 -   0.0050W       -        -    2  2  2  2     1900    5300

=== START OF SMART DATA SECTION ===
SMART overall-health self-assessment test result: PASSED

SMART/Health Information (NVMe Log 0x02)
Critical Warning:                   0x00
Temperature:                        38 Celsius
Available Spare:                    100%
Available Spare Threshold:          10%
Percentage Used:                    2%
Data Units Read:                    27,627,372 [14.1 TB]
Data Units Written:                 27,460,731 [14.0 TB]
Host Read Commands:                 399,847,733
Host Write Commands:                344,865,053
Controller Busy Time:               962
Power Cycles:                       24,149
Power On Hours:                     461
Unsafe Shutdowns:                   163
Media and Data Integrity Errors:    0
Error Information Log Entries:      108

Read Error Information Log failed: NVMe admin command:0x02/page:0x01 is not supported
`
