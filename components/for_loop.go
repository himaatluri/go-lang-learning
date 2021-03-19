
// for [index], [array/slice] := range array/slice {
// 	fmt.Println(array/slice)
// }

// index: Index of this element in the array
// array/slice: current list that is iterated
// IMPORTANT ELEMENT: range array/slice: Take the slice/array and loop over it
// fmt.Println(array/slice): Run this one time for each card in the slice
// range is a keyword that we use every time to iterate

c := []string{"card1", "card2", "card3"}

for i, card := range cards {
	fmt.Println(i, card)
}