package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type mapData struct {
	src int
	dst int
	rng int
}

type mapper struct {
	data []mapData
}

type multiMapper struct {
	mappers []mapper
}

func (m *mapper) mapFrom(src int) int {
	for _, d := range m.data {
		if src >= d.src && src < d.src+d.rng {
			return d.dst + (src - d.src)
		}
	}
	return src
}

func (m *multiMapper) mapFrom(src int) int {
	for _, mapper := range m.mappers {
		src = mapper.mapFrom(src)
	}
	return src
}

func part1() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	seedsSl := strings.Fields(strings.Split(lines[0], ": ")[1])

	seeds := strSlcInt(seedsSl)

	fmt.Println(seeds)

	maps := mapPaths(lines)

	var minLocation int = math.MaxInt

	for _, seed := range seeds {
		minLocation = min(minLocation, maps.mapFrom(seed))
	}

	fmt.Println(minLocation)
}

func mapPaths(data []string) multiMapper {
	mappers := make([]mapper, 0, 7)
	for _, dat := range data[1:] {
		if dat == "" {
			mappers = append(mappers, mapper{})
			continue
		}
		parts := strings.Split(dat, " ")
		if len(parts) == 2 {
			continue
		}

		dst, err := strconv.Atoi(parts[0])
		checkError(err)
		src, err := strconv.Atoi(parts[1])
		checkError(err)
		rng, err := strconv.Atoi(parts[2])
		mappers[len(mappers)-1].data = append(mappers[len(mappers)-1].data, mapData{src, dst, rng})
	}
	return multiMapper{mappers}
}

func checkSeed(seed int, maps multiMapper) {

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func strSlcInt(strSlice []string) []int {
	var ints []int
	for _, str := range strSlice {
		int, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, int)
	}
	return ints
}
