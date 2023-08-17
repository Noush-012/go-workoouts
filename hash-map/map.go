package hashmap

func Mapping() {

	mapp := make(map[string]int)

	for i := 'a'; i < 'j'; i++ {
		mapp[string(i)]++
	}

	// fmt.Println(mapp)

	// for k, v := range mapp {

	// }
}
