package main

func main() {
	/*	p := make([]int, 5000000)
		for i := range p {
			p[i] = rand.Int()
		}
		file, err := os.Create("array.dat")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(p)
		if err != nil {
			panic(err)
		}*/
	/*p := make([]int, 5000000)
	file, err := os.Open("array.dat")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&p)
	if err != nil {
		panic(err)
	}*/
}
