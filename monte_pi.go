package main

import (
  "math/rand"
  "time"
  "runtime"
)

func PI(samples int) float64 {
  inside := 0

  for i := 0; i < samples; i++ {
    x := rand.Float64()
    y := rand.Float64()
    if (x*x + y*y) < 1 {
      inside++
    }
  }

  ratio := float64(inside) / float64(samples)

  return ratio * 4
}

func MultiPI(samples int) float64 {
  cpus := runtime.NumCPU()

  threadSamples := samples / cpus
  result := make(chan float64, cpus)

  for j := 0; j < cpus; j++ {
    go func() {
      var inside int
      r := rand.New(rand.NewSource(time.Now().UnixNano()))
      for i := 0; i < threadSamples; i++ {
        x, y := r.Float64(), r.Float64()

        if (x*x + y*y) < 1 {
          inside++
        }
      }
      result <- (float64(inside) / float64(threadSamples)) * 4
    }()
  }

  var total float64
  for i := 0; i < cpus; i++ {
    total += <-result
  }

  return total / float64(cpus)
}

func init() {
  runtime.GOMAXPROCS(runtime.NumCPU())
  rand.Seed(time.Now().UnixNano())
}
