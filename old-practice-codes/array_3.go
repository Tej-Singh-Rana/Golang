package main

import "fmt"

func main()  {
	// var x [7]float64
	// x[0] = 98
  // x[1] = 93
  // x[2] = 77
  // x[3] = 82
  // x[4] = 83
	// x[5] = 100
	// x[6] = 120
			// or
	// x := [7]float64{ 98, 93, 77, 82, 83, 100, 120}
			// or
		x := [7]float64{
			98,
			93,
			77,
			82,
			83,
			100,		// In every new line we have to use comma to initialize next one. If any element we want to remove from an array then just comment it.
			120,		// It will not impact other element while removing or truncating.
		}
	var total float64 = 0
	/* In this for loop "i" represents the current position(index value) in the array and value is the same as x[i]. 
	 We use the keyword range followed by the name of the variable we want to loop over.
	 we will get an error "i" declared but not used.
	 so in place of "i" we use a single _ (underscore) is used to tell the compiler that we don't need this.
	 (In this case we don't need the iterator variable).
	 
	*/
	//  for i, value := range x 
	for _, value := range x {
  total += value
}
	fmt.Println(total / float64(len(x)))
	
}