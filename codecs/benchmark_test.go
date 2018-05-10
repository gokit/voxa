package codecs

import (
	"encoding/json"
	"testing"
	"time"
)

type address struct {
	Value string `id:"1"`
}

type model struct {
	Age        int       `id:"1"`
	Name       string    `id:"2"`
	Address    string    `id:"3"`
	Date       time.Time `id:"5"`
	OtherNames []address `id:"4"`
}

func BenchmarkListCodec_JSONNativeToBinary_Expanding(b *testing.B) {
	record := make([]model, 0, 1000)

	for i := 0; i < 1000; i++ {
		record = append(record, model{
			Age:     20,
			Name:    "bob",
			Address: "20. Classy Street",
			Date:    time.Now(),
			OtherNames: []address{
				{Value: "wreckage"},
				{Value: "moppers guild"},
				{Value: "Is His always Faithful!"},
			},
		})
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(record)
	}
	b.StopTimer()
}

func BenchmarkListCodec_NativeToBinary_Expanding(b *testing.B) {
	record := make([]model, 0, 1000)

	for i := 0; i < 1000; i++ {
		record = append(record, model{
			Age:     20,
			Name:    "bob",
			Address: "20. Classy Street",
			Date:    time.Now(),
			OtherNames: []address{
				{Value: "wreckage"},
				{Value: "moppers guild"},
				{Value: "Is His always Faithful!"},
			},
		})
	}

	var codec ListCodec

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codec.NativeToBinary(record, nil)
	}
	b.StopTimer()
}

func BenchmarkListCodec_NativeToBinary_Single(b *testing.B) {
	record := []model{
		{
			Age:     20,
			Name:    "bob",
			Address: "20. Classy Street",
			Date:    time.Now(),
			OtherNames: []address{
				{Value: "wreckage"},
				{Value: "moppers guild"},
				{Value: "Is His always Faithful!"},
			},
		},
	}

	var codec ListCodec

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codec.NativeToBinary(record, nil)
	}
	b.StopTimer()
}

func BenchmarkListCodec_JSONNativeToBinary_Single(b *testing.B) {
	record := []model{
		{
			Age:     20,
			Name:    "bob",
			Address: "20. Classy Street",
			Date:    time.Now(),
			OtherNames: []address{
				{Value: "wreckage"},
				{Value: "moppers guild"},
				{Value: "Is His always Faithful!"},
			},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(record)
	}
	b.StopTimer()
}

func BenchmarkRecordCodec_NativeToBinary(b *testing.B) {
	record := model{
		Age:     20,
		Name:    "bob",
		Address: "20. Classy Street",
		Date:    time.Now(),
		OtherNames: []address{
			{Value: "wreckage"},
			{Value: "moppers guild"},
			{Value: "Is His always Faithful!"},
		},
	}

	var codec RecordCodec

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codec.NativeToBinary(record, nil)
	}
	b.StopTimer()
}

func BenchmarkRecordCodec_JSONNativeToBinary(b *testing.B) {
	record := model{
		Age:     20,
		Name:    "bob",
		Address: "20. Classy Street",
		Date:    time.Now(),
		OtherNames: []address{
			{Value: "wreckage"},
			{Value: "moppers guild"},
			{Value: "Is His always Faithful!"},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(record)
	}
	b.StopTimer()
}
