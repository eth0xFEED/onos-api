// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package types

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestPlmnIDEncode(t *testing.T) {
	plmnID := EncodePlmnID("208", "20")
	assert.Equal(t, "02F802", plmnID)

	mcc, mnc := DecodePlmnID(plmnID)
	assert.Equal(t, "208", mcc)
	assert.Equal(t, "20", mnc)

	plmnID = EncodePlmnID("310", "170")
	assert.Equal(t, "130071", plmnID)

	mcc, mnc = DecodePlmnID(plmnID)
	assert.Equal(t, "310", mcc)
	assert.Equal(t, "170", mnc)

	assert.Equal(t, PlmnID(0x130071), ToPlmnID("310", "170"))
	assert.Equal(t, PlmnID(0x130071), PlmnIDFromHexString("130071"))

	assert.Equal(t, PlmnID(0x130071), PlmnIDFromString("310170"))
	assert.Equal(t, "310170", PlmnIDToString(PlmnID(0x130071)))
}

func TestFullShiftBasics(t *testing.T) {
	plmnID := PlmnID(0xBBBCCC)
	cellID := CellID(0xAA)
	enbID := EnbID(0xFFFFF)

	assert.Equal(t, ECI(0xFFFFFAA), ToECI(enbID, cellID))
	assert.Equal(t, ECGI(0xBBBCCCFFFFFAA), ToECGI(plmnID, ToECI(enbID, cellID)))
}

func TestTypesWithFullShift(t *testing.T) {
	plmnID := PlmnID(0xbbbccc)
	cellID := CellID(0x12)
	enbID := EnbID(0xf8f8f)

	eci := ToECI(enbID, cellID)
	ecgi := ToECGI(plmnID, eci)
	genbID := ToGEnbID(plmnID, enbID)

	// NOTE: These work for the given values of cellID and enbID, but may fail with "lesser" values that result
	// in shifts less than 8 or 28 bits respectively.
	assert.Equal(t, plmnID, GetPlmnID(uint64(ecgi)), "incorrect PLMNID")
	assert.Equal(t, enbID, GetEnbID(uint64(genbID)), "incorrect EnbID")
	assert.Equal(t, eci, GetECI(uint64(ecgi)), "incorrect ECI")
	assert.Equal(t, cellID, GetCellID(uint64(ecgi)), "incorrect CID")
}

func TestTypesWithPartialShift(t *testing.T) {
	plmnID := PlmnID(0xbbbaaa)
	cellID := CellID(0x0f)
	enbID := EnbID(0xfffff)

	eci := ToECI(enbID, cellID)
	ecgi := ToECGI(plmnID, eci)
	genbID := ToGEnbID(plmnID, enbID)

	// NOTE: These work for the given values of cellID and enbID, but may fail with "lesser" values that result
	// in shifts less than 8 or 28 bits respectively.
	assert.Equal(t, plmnID, GetPlmnID(uint64(ecgi)), "incorrect PLMNID")
	assert.Equal(t, enbID, GetEnbID(uint64(genbID)), "incorrect EnbID")
	assert.Equal(t, eci, GetECI(uint64(ecgi)), "incorrect ECI")
	assert.Equal(t, cellID, GetCellID(uint64(ecgi)), "incorrect CID")
}

func TestSimValues(t *testing.T) {
	plmnID := ToPlmnID("310", "170")
	assert.Equal(t, PlmnID(0x130071), plmnID)

	enb1 := EnbID(144470)
	enb2 := EnbID(144471)
	ecgi11 := ToECGI(plmnID, ToECI(enb1, CellID(1)))
	ecgi12 := ToECGI(plmnID, ToECI(enb1, CellID(2)))
	ecgi21 := ToECGI(plmnID, ToECI(enb2, CellID(1)))
	ecgi22 := ToECGI(plmnID, ToECI(enb2, CellID(2)))

	fmt.Printf("%d\n%d\n%d\n%d\n", ecgi11, ecgi12, ecgi21, ecgi22)
}
