package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func (m *mapper) mapFromTwo(src int) int {
	for _, d := range m.data {
		if src < d.src || src > d.src+d.rng {
			return src
		} else if src >= d.src && src < d.src+d.rng {
			return d.dst + (src - d.src)
		}
	}
	return src
}

func (m *multiMapper) mapFromTwo(src int) int {
	for _, mapper := range m.mappers {
		src = mapper.mapFrom(src)
	}
	return src
}

func part2() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	seedsSl := strings.Fields(strings.Split(lines[0], ": ")[1])

	seeds := strSlcInt(seedsSl)

	maps := mapPaths(lines)

	var minLocation int = math.MaxInt

	for i := 0; i < len(seeds); i += 2 {
		end := seeds[i] + seeds[i+1]
		for j := seeds[i]; j < end; j++ {
			minLocation = min(minLocation, maps.mapFrom(j))
		}
	}
	fmt.Println(minLocation)
}
