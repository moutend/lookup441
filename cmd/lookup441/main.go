package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/moutend/lookup441/internal/database"
	"github.com/moutend/lookup441/internal/models"
	"github.com/moutend/lookup441/internal/synth"
	"github.com/moutend/lookup441/internal/util"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	if err := run(); err != nil {
		log.New(os.Stderr, "error: ", 0).Fatal(err)
	}
}

func run() error {
	sampleRateFlag := flag.Int("r", 44100, "sample rate (e.g. 44100)")
	bufferSizeFlag := flag.Int("b", 4096, "buffer size (e.g. 4096)")
	durationFlag := flag.Int("d", 2, "duration in second (e.g. 2)")

	flag.Parse()

	if sampleRateFlag == nil || bufferSizeFlag == nil || durationFlag == nil {
		return fmt.Errorf("invalid flag")
	}

	sampleRate := *sampleRateFlag
	bufferSize := *bufferSizeFlag
	duration := *durationFlag
	length := duration * sampleRate

	signal := make([]float32, length, length)

	for frequency := 200; frequency <= 44000; frequency += 1 {
		wave := synth.NewSin(float64(frequency)/10.0, float64(sampleRate))

		for i := range signal {
			signal[i] = float32(wave.Read())
		}

		c := Container{
			SampleRate: int64(sampleRate),
			BufferSize: int64(bufferSize),
			Frequency:  int64(frequency),
			Period:     countZeroCross(signal, bufferSize),
		}

		if err := c.save(context.Background()); err != nil {
			return err
		}
	}

	return nil
}

type Container struct {
	SampleRate int64
	BufferSize int64
	Frequency  int64
	Period     float64
}

func (v Container) save(ctx context.Context) error {
	if err := database.Setup("frequency.db3"); err != nil {
		return err
	}

	defer database.Teardown()

	if err := database.Transaction(ctx, func(ctx context.Context, tx boil.ContextTransactor) error {
		p, err := models.FindPeriod(ctx, tx, v.SampleRate, v.BufferSize, v.Frequency)

		if err != nil && err != sql.ErrNoRows {
			return err
		}

		found := p != nil

		if !found {
			p = &models.Period{
				SampleRate: v.SampleRate,
				BufferSize: v.BufferSize,
				Frequency:  v.Frequency,
				Period:     v.Period,
			}
		}
		if found {
			_, err = p.Update(ctx, tx, boil.Infer())
		} else {
			err = p.Insert(ctx, tx, boil.Infer())
		}

		return err
	}); err != nil {
		return err
	}

	fmt.Printf("DONE\t%d\n", v.Frequency)

	return nil
}

func countZeroCross(signal []float32, bufferSize int) float64 {
	average := &util.Average{}

	for i := 0; i < len(signal)/10; i++ {
		zc := &util.ZeroCross{}

		for j := i; j < (i + bufferSize); j++ {
			zc.Apply(signal[j])
		}

		average.Append(zc.Count())
	}

	return average.Value()
}
